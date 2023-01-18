package v0

import (
	"context"
	"time"

	"github.com/filecoin-project/go-state-types/builtin/v10/util/adt"
	"github.com/filecoin-project/go-state-types/store"
	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/lily/pkg/extract/actors"
	"github.com/filecoin-project/lily/tasks"
)

type StateDiff struct {
	DiffMethods []actors.ActorStateDiff
}

func (s *StateDiff) State(ctx context.Context, api tasks.DataSource, act *actors.ActorChange) (actors.ActorDiffResult, error) {
	start := time.Now()
	var stateDiff = new(StateDiffResult)
	for _, f := range s.DiffMethods {
		stateChange, err := f.Diff(ctx, api, act)
		if err != nil {
			return nil, err
		}
		if stateChange == nil {
			continue
		}
		switch stateChange.Kind() {
		case KindPowerClaims:
			stateDiff.ClaimsChanges = stateChange.(ClaimsChangeList)
		default:
			panic(stateChange.Kind())
		}
	}
	log.Infow("Extracted Power State Diff", "address", act.Address, "duration", time.Since(start))
	return stateDiff, nil
}

type StateDiffResult struct {
	ClaimsChanges ClaimsChangeList
}

func (sd *StateDiffResult) MarshalStateChange(ctx context.Context, s store.Store) (cbg.CBORMarshaler, error) {
	out := &StateChange{}
	if claims := sd.ClaimsChanges; claims != nil {
		root, err := claims.ToAdtMap(s, 5)
		if err != nil {
			return nil, err
		}
		out.Claims = &root
	}
	return out, nil
}

func (sd *StateDiffResult) Kind() string {
	return "power"
}

type StateChange struct {
	Claims *cid.Cid `cborgen:"claims"`
}

func (sc *StateChange) ToStateDiffResult(ctx context.Context, s store.Store) (*StateDiffResult, error) {
	out := &StateDiffResult{ClaimsChanges: ClaimsChangeList{}}

	if sc.Claims != nil {
		claimMap, err := adt.AsMap(s, *sc.Claims, 5)
		if err != nil {
			return nil, err
		}

		claims := ClaimsChangeList{}
		claimChange := new(ClaimsChange)
		if err := claimMap.ForEach(claimChange, func(key string) error {
			val := new(ClaimsChange)
			*val = *claimChange
			claims = append(claims, val)
			return nil
		}); err != nil {
			return nil, err
		}
		out.ClaimsChanges = claims
	}
	return out, nil
}
