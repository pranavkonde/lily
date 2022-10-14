package multisig

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
	multisigmodel "github.com/filecoin-project/lily/model/actors/multisig"
	v2 "github.com/filecoin-project/lily/model/v2"
	"github.com/filecoin-project/lily/model/v2/actors/multisig"
)

var log = logging.Logger("transform/multisig")

type TransactionTransform struct {
	meta     v2.ModelMeta
	taskName string
}

func NewTransactionTransform(taskName string) *TransactionTransform {
	info := multisig.MultisigTransaction{}
	return &TransactionTransform{meta: info.Meta(), taskName: taskName}
}

func (s *TransactionTransform) Run(ctx context.Context, reporter string, in chan *extract.ActorStateResult, out chan transform.Result) error {
	log.Debugf("run %s", s.Name())
	for res := range in {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			report := actor.ToProcessingReport(s.taskName, reporter, res)
			data := model.PersistableList{report}
			log.Debugw("received data", "count", len(res.Results.Models()))
			sqlModels := make(multisigmodel.MultisigTransactionList, 0, len(res.Results.Models()))
			for _, modeldata := range res.Results.Models() {
				tx := modeldata.(*multisig.MultisigTransaction)
				if tx.Event == multisig.Added || tx.Event == multisig.Modified {
					approved := make([]string, len(tx.Approved))
					for i, addr := range tx.Approved {
						approved[i] = addr.String()
					}
					sqlModels = append(sqlModels, &multisigmodel.MultisigTransaction{
						MultisigID:    tx.Multisig.String(),
						StateRoot:     tx.StateRoot.String(),
						Height:        int64(tx.Height),
						TransactionID: tx.TransactionID,
						To:            tx.To.String(),
						Value:         tx.Value.String(),
						Method:        uint64(tx.Method),
						Params:        tx.Params,
						Approved:      approved,
					})
				}
			}
			if len(sqlModels) > 0 {
				data = append(data, sqlModels)
			}
			out <- &persistable.Result{Model: data}
		}
	}
	return nil
}

func (s *TransactionTransform) ModelType() v2.ModelMeta {
	return s.meta
}

func (s *TransactionTransform) Name() string {
	info := TransactionTransform{}
	return reflect.TypeOf(info).Name()
}

func (s *TransactionTransform) Matcher() string {
	return fmt.Sprintf("^%s$", s.meta.String())
}
