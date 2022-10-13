package message

import (
	"context"
	"fmt"
	"reflect"

	"github.com/filecoin-project/lily/chain/indexer/v2/transform"
	"github.com/filecoin-project/lily/chain/indexer/v2/transform/persistable"
	"github.com/filecoin-project/lily/lens/util"
	messages2 "github.com/filecoin-project/lily/model/messages"
	v2 "github.com/filecoin-project/lily/model/v2"
	"github.com/filecoin-project/lily/model/v2/messages"
)

type ImplicitMessageTransform struct {
	meta v2.ModelMeta
}

func NewImplicitMessageTransform() *ImplicitMessageTransform {
	info := messages.VMMessage{}
	return &ImplicitMessageTransform{meta: info.Meta()}
}

func (s *ImplicitMessageTransform) Run(ctx context.Context, in chan transform.IndexState, out chan transform.Result) error {
	log.Debugf("run %s", s.Name())
	for res := range in {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			sqlModels := make(messages2.InternalMessageList, 0, len(res.Models()))
			for _, modeldata := range res.Models() {
				vm := modeldata.(*messages.VMMessage)
				if !vm.Implicit {
					continue
				}
				name, family, err := util.ActorNameAndFamilyFromCode(vm.ToActorCode)
				if err != nil {
					return err
				}
				sqlModels = append(sqlModels, &messages2.InternalMessage{
					Height:        int64(vm.Height),
					Cid:           vm.MessageCID.String(),
					StateRoot:     vm.StateRoot.String(),
					SourceMessage: "", // source of implicit messages DNE.
					From:          vm.From.String(),
					To:            vm.To.String(),
					Value:         vm.Value.String(),
					Method:        uint64(vm.Method),
					ActorName:     name,
					ActorFamily:   family,
					ExitCode:      int64(vm.ExitCode),
					GasUsed:       vm.GasUsed,
				})
			}
			if len(sqlModels) > 0 {
				out <- &persistable.Result{Model: sqlModels}
			}
		}
	}
	return nil
}

func (s *ImplicitMessageTransform) ModelType() v2.ModelMeta {
	return s.meta
}

func (s *ImplicitMessageTransform) Name() string {
	info := ImplicitMessageTransform{}
	return reflect.TypeOf(info).Name()
}

func (s *ImplicitMessageTransform) Matcher() string {
	return fmt.Sprintf("^%s$", s.meta.String())
}
