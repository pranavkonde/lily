// Code generated by: `make actors-gen`. DO NOT EDIT.
package verifreg

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lily/chain/actors/adt"

	"crypto/sha256"

	builtin8 "github.com/filecoin-project/go-state-types/builtin"
	adt8 "github.com/filecoin-project/go-state-types/builtin/v8/util/adt"
	verifreg8 "github.com/filecoin-project/go-state-types/builtin/v8/verifreg"

	verifreg9 "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"

	actorstypes "github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/lotus/chain/actors"
)

var _ State = (*state8)(nil)

func load8(store adt.Store, root cid.Cid) (State, error) {
	out := state8{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state8 struct {
	verifreg8.State
	store adt.Store
}

func (s *state8) ActorKey() string {
	return manifest.VerifregKey
}

func (s *state8) ActorVersion() actorstypes.Version {
	return actorstypes.Version8
}

func (s *state8) Code() cid.Cid {
	code, ok := actors.GetActorCodeID(s.ActorVersion(), s.ActorKey())
	if !ok {
		panic(fmt.Errorf("didn't find actor %v code id for actor version %d", s.ActorKey(), s.ActorVersion()))
	}

	return code
}

func (s *state8) VerifiedClientsMapBitWidth() int {

	return builtin8.DefaultHamtBitwidth

}

func (s *state8) VerifiedClientsMapHashFunction() func(input []byte) []byte {

	return func(input []byte) []byte {
		res := sha256.Sum256(input)
		return res[:]
	}

}

func (s *state8) VerifiedClientsMap() (adt.Map, error) {

	return adt8.AsMap(s.store, s.VerifiedClients, builtin8.DefaultHamtBitwidth)

}

func (s *state8) VerifiersMap() (adt.Map, error) {
	return adt8.AsMap(s.store, s.Verifiers, builtin8.DefaultHamtBitwidth)
}

func (s *state8) VerifiersMapBitWidth() int {

	return builtin8.DefaultHamtBitwidth

}

func (s *state8) VerifiersMapHashFunction() func(input []byte) []byte {

	return func(input []byte) []byte {
		res := sha256.Sum256(input)
		return res[:]
	}

}

func (s *state8) RootKey() (address.Address, error) {
	return s.State.RootKey, nil
}

func (s *state8) VerifiedClientDataCap(addr address.Address) (bool, abi.StoragePower, error) {

	return getDataCap(s.store, actorstypes.Version8, s.VerifiedClientsMap, addr)

}

func (s *state8) VerifierDataCap(addr address.Address) (bool, abi.StoragePower, error) {
	return getDataCap(s.store, actorstypes.Version8, s.VerifiersMap, addr)
}

func (s *state8) RemoveDataCapProposalID(verifier address.Address, client address.Address) (bool, uint64, error) {
	return getRemoveDataCapProposalID(s.store, actorstypes.Version8, s.removeDataCapProposalIDs, verifier, client)
}

func (s *state8) ForEachVerifier(cb func(addr address.Address, dcap abi.StoragePower) error) error {
	return forEachCap(s.store, actorstypes.Version8, s.VerifiersMap, cb)
}

func (s *state8) ForEachClient(cb func(addr address.Address, dcap abi.StoragePower) error) error {

	return forEachCap(s.store, actorstypes.Version8, s.VerifiedClientsMap, cb)

}

func (s *state8) removeDataCapProposalIDs() (adt.Map, error) {
	return adt8.AsMap(s.store, s.RemoveDataCapProposalIDs, builtin8.DefaultHamtBitwidth)
}

func (s *state8) GetState() interface{} {
	return &s.State
}

func (s *state8) GetAllocation(clientIdAddr address.Address, allocationId verifreg9.AllocationId) (*Allocation, bool, error) {

	return nil, false, fmt.Errorf("unsupported in actors v8")

}

func (s *state8) GetAllocations(clientIdAddr address.Address) (map[AllocationId]Allocation, error) {

	return nil, fmt.Errorf("unsupported in actors v8")

}

func (s *state8) GetClaim(providerIdAddr address.Address, claimId verifreg9.ClaimId) (*Claim, bool, error) {

	return nil, false, fmt.Errorf("unsupported in actors v8")

}

func (s *state8) GetClaims(providerIdAddr address.Address) (map[ClaimId]Claim, error) {

	return nil, fmt.Errorf("unsupported in actors v8")

}

func (s *state8) ClaimsMap() (adt.Map, error) {

	return nil, fmt.Errorf("unsupported in actors v8")

}

// TODO this could return an error since not all versions have a claims map
func (s *state8) ClaimsMapBitWidth() int {

	return builtin8.DefaultHamtBitwidth

}

// TODO this could return an error since not all versions have a claims map
func (s *state8) ClaimsMapHashFunction() func(input []byte) []byte {

	return func(input []byte) []byte {
		res := sha256.Sum256(input)
		return res[:]
	}

}

func (s *state8) ClaimMapForProvider(providerIdAddr address.Address) (adt.Map, error) {

	return nil, fmt.Errorf("unsupported in actors v8")

}

func (s *state8) getInnerHamtCid(store adt.Store, key abi.Keyer, mapCid cid.Cid, bitwidth int) (cid.Cid, error) {

	return cid.Undef, fmt.Errorf("unsupported in actors v8")

}
