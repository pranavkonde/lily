package transform

import (
	"context"
	"fmt"

	"github.com/filecoin-project/lotus/chain/types"
	logging "github.com/ipfs/go-log/v2"
	evntbus "github.com/mustafaturan/bus/v3"
	"golang.org/x/sync/errgroup"

	"github.com/filecoin-project/lily/chain/indexer/v2/bus"
	v22 "github.com/filecoin-project/lily/chain/indexer/v2/extract"
	v2 "github.com/filecoin-project/lily/model/v2"
	"github.com/filecoin-project/lily/tasks"
)

var log = logging.Logger("transform")

type Kind string

type Result interface {
	Kind() Kind
	Data() interface{}
}

type IndexState interface {
	Task() v2.ModelMeta
	Current() *types.TipSet
	Executed() *types.TipSet
	Complete() bool
	State() *v22.StateResult
}

type Handler interface {
	Run(ctx context.Context, api tasks.DataSource, in chan IndexState, out chan Result) error
	Name() string
	ModelType() v2.ModelMeta
}

func NewRouter(handlers ...Handler) (*Router, error) {
	b, err := bus.NewBus()
	if err != nil {
		return nil, err
	}
	handlerChans := make([]chan IndexState, len(handlers))
	routerHandlers := make([]Handler, len(handlers))
	registry := make(map[v2.ModelMeta][]Handler)
	for i, handler := range handlers {
		// map of model types to handlers for said type
		registry[handler.ModelType()] = append(registry[handler.ModelType()], handler)
		// maintain list of handlers
		routerHandlers[i] = handler
		// initialize handler channel
		handlerChans[i] = make(chan IndexState, 8) // TODO buffer channel
		// register the handle topic with the bus
		b.Bus.RegisterTopics(handler.ModelType().String())
		// register handler for its required model, all models the hander can process are sent on its channel
		hch := handlerChans[i]
		b.Bus.RegisterHandler(handler.Name(), evntbus.Handler{
			Handle: func(ctx context.Context, e evntbus.Event) {
				hch <- e.Data.(IndexState)
			},
			// TODO fix this annoying shit, make your own damn bus, this one is falling a bit short..
			Matcher: fmt.Sprintf("^%s$", handler.ModelType().String()),
		})
	}
	return &Router{
		registry:        registry,
		bus:             b,
		resultCh:        make(chan Result, len(routerHandlers)), // TODO buffer channel
		handlerChannels: handlerChans,
		handlerGrp:      &errgroup.Group{},
		handlers:        routerHandlers,
	}, nil
}

type Router struct {
	registry        map[v2.ModelMeta][]Handler
	bus             *bus.Bus
	resultCh        chan Result
	handlerChannels []chan IndexState
	handlerGrp      *errgroup.Group
	handlers        []Handler
}

func (r *Router) Start(ctx context.Context, api tasks.DataSource) {
	log.Infow("starting router", "topics", r.bus.Bus.Topics())
	for i, handler := range r.handlers {
		log.Infow("start handler", "type", handler.Name())
		i := i
		handler := handler
		r.handlerGrp.Go(func() error {
			return handler.Run(ctx, api, r.handlerChannels[i], r.resultCh)
		})
	}
}

func (r *Router) Stop() error {
	log.Info("stopping router")
	// close all channel feeding handlers
	for _, c := range r.handlerChannels {
		close(c)
	}
	log.Info("closed handler channels")
	// wait for handlers to complete and drain their now closed channel
	err := r.handlerGrp.Wait()
	log.Infow("handlers completed", "error", err)
	// close the output channel signaling there are no more results to handle.
	close(r.resultCh)
	log.Info("router stopped")
	return err
}

func (r *Router) Route(ctx context.Context, data IndexState) error {
	log.Debugw("routing data", "type", data.Task().String())
	return r.bus.Bus.Emit(ctx, data.Task().String(), data)
}

func (r *Router) Results() chan Result {
	return r.resultCh
}
