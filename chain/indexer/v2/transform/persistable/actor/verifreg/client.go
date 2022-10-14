package verifreg

import (
	"context"
	"fmt"
	"reflect"

	logging "github.com/ipfs/go-log/v2"

	"github.com/filecoin-project/lily/chain/indexer/v2/extract"
	"github.com/filecoin-project/lily/chain/indexer/v2/transform"
	"github.com/filecoin-project/lily/chain/indexer/v2/transform/persistable"
	"github.com/filecoin-project/lily/chain/indexer/v2/transform/persistable/actor"
	"github.com/filecoin-project/lily/model"
	verifregmodel "github.com/filecoin-project/lily/model/actors/verifreg"
	v2 "github.com/filecoin-project/lily/model/v2"
	"github.com/filecoin-project/lily/model/v2/actors/verifreg"
)

var log = logging.Logger("transform/verifreg")

type VerifiedClientTransform struct {
	meta     v2.ModelMeta
	taskName string
}

func NewVerifiedClientTransform(taskName string) *VerifiedClientTransform {
	info := verifreg.VerifiedClient{}
	return &VerifiedClientTransform{meta: info.Meta(), taskName: taskName}
}

func (s *VerifiedClientTransform) Run(ctx context.Context, reporter string, in chan *extract.ActorStateResult, out chan transform.Result) error {
	log.Debugf("run %s", s.Name())
	for res := range in {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			report := actor.ToProcessingReport(s.taskName, reporter, res)
			data := model.PersistableList{report}
			log.Debugw("received data", "count", len(res.Results.Models()))
			sqlModels := make(verifregmodel.VerifiedRegistryVerifiedClientsList, 0, len(res.Results.Models()))
			for _, modeldata := range res.Results.Models() {
				vc := modeldata.(*verifreg.VerifiedClient)
				sqlModels = append(sqlModels, &verifregmodel.VerifiedRegistryVerifiedClient{
					Height:    int64(vc.Height),
					StateRoot: vc.StateRoot.String(),
					Address:   vc.Client.String(),
					Event:     vc.Event.String(),
					DataCap:   vc.DataCap.String(),
				})
			}
			if len(sqlModels) > 0 {
				data = append(data, sqlModels)
			}
			out <- &persistable.Result{Model: data}
		}
	}
	return nil
}

func (s *VerifiedClientTransform) ModelType() v2.ModelMeta {
	return s.meta
}

func (s *VerifiedClientTransform) Name() string {
	info := VerifiedClientTransform{}
	return reflect.TypeOf(info).Name()
}

func (s *VerifiedClientTransform) Matcher() string {
	return fmt.Sprintf("^%s$", s.meta.String())
}
