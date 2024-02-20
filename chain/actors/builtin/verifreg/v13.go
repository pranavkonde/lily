// Code generated by: `make actors-gen`. DO NOT EDIT.
package verifreg

import (
	"crypto/sha256"
	"fmt"

	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	actorstypes "github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/go-state-types/big"
	builtin13 "github.com/filecoin-project/go-state-types/builtin"
	adt13 "github.com/filecoin-project/go-state-types/builtin/v13/util/adt"
	verifreg13 "github.com/filecoin-project/go-state-types/builtin/v13/verifreg"
	verifreg9 "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/lily/chain/actors/adt"

	"github.com/filecoin-project/lotus/chain/actors"
)

var _ State = (*state13)(nil)

func load13(store adt.Store, root cid.Cid) (State, error) {
	out := state13{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state13 struct {
	verifreg13.State
	store adt.Store
}

func (s *state13) ActorKey() string {
	return manifest.VerifregKey
}

func (s *state13) ActorVersion() actorstypes.Version {
	return actorstypes.Version13
}

func (s *state13) Code() cid.Cid {
	code, ok := actors.GetActorCodeID(s.ActorVersion(), s.ActorKey())
	if !ok {
		panic(fmt.Errorf("didn't find actor %v code id for actor version %d", s.ActorKey(), s.ActorVersion()))
	}

	return code
}

func (s *state13) VerifiedClientsMapBitWidth() int {

	return builtin13.DefaultHamtBitwidth

}

func (s *state13) VerifiedClientsMapHashFunction() func(input []byte) []byte {

	return func(input []byte) []byte {
		res := sha256.Sum256(input)
		return res[:]
	}

}

func (s *state13) VerifiedClientsMap() (adt.Map, error) {

	return nil, fmt.Errorf("unsupported in actors v13")

}

func (s *state13) VerifiersMap() (adt.Map, error) {
	return adt13.AsMap(s.store, s.Verifiers, builtin13.DefaultHamtBitwidth)
}

func (s *state13) VerifiersMapBitWidth() int {

	return builtin13.DefaultHamtBitwidth

}

func (s *state13) VerifiersMapHashFunction() func(input []byte) []byte {

	return func(input []byte) []byte {
		res := sha256.Sum256(input)
		return res[:]
	}

}

func (s *state13) RootKey() (address.Address, error) {
	return s.State.RootKey, nil
}

func (s *state13) VerifiedClientDataCap(addr address.Address) (bool, abi.StoragePower, error) {

	return false, big.Zero(), fmt.Errorf("unsupported in actors v13")

}

func (s *state13) VerifierDataCap(addr address.Address) (bool, abi.StoragePower, error) {
	return getDataCap(s.store, actorstypes.Version13, s.VerifiersMap, addr)
}

func (s *state13) RemoveDataCapProposalID(verifier address.Address, client address.Address) (bool, uint64, error) {
	return getRemoveDataCapProposalID(s.store, actorstypes.Version13, s.removeDataCapProposalIDs, verifier, client)
}

func (s *state13) ForEachVerifier(cb func(addr address.Address, dcap abi.StoragePower) error) error {
	return forEachCap(s.store, actorstypes.Version13, s.VerifiersMap, cb)
}

func (s *state13) ForEachClient(cb func(addr address.Address, dcap abi.StoragePower) error) error {

	return fmt.Errorf("unsupported in actors v13")

}

func (s *state13) removeDataCapProposalIDs() (adt.Map, error) {
	return adt13.AsMap(s.store, s.RemoveDataCapProposalIDs, builtin13.DefaultHamtBitwidth)
}

func (s *state13) GetState() interface{} {
	return &s.State
}

func (s *state13) GetAllocation(clientIdAddr address.Address, allocationId verifreg9.AllocationId) (*Allocation, bool, error) {

	alloc, ok, err := s.FindAllocation(s.store, clientIdAddr, verifreg13.AllocationId(allocationId))
	return (*Allocation)(alloc), ok, err
}

func (s *state13) GetAllocations(clientIdAddr address.Address) (map[AllocationId]Allocation, error) {

	v13Map, err := s.LoadAllocationsToMap(s.store, clientIdAddr)

	retMap := make(map[AllocationId]Allocation, len(v13Map))
	for k, v := range v13Map {
		retMap[AllocationId(k)] = Allocation(v)
	}

	return retMap, err

}

func (s *state13) GetClaim(providerIdAddr address.Address, claimId verifreg9.ClaimId) (*Claim, bool, error) {

	claim, ok, err := s.FindClaim(s.store, providerIdAddr, verifreg13.ClaimId(claimId))
	return (*Claim)(claim), ok, err

}

func (s *state13) GetClaims(providerIdAddr address.Address) (map[ClaimId]Claim, error) {

	v13Map, err := s.LoadClaimsToMap(s.store, providerIdAddr)

	retMap := make(map[ClaimId]Claim, len(v13Map))
	for k, v := range v13Map {
		retMap[ClaimId(k)] = Claim(v)
	}

	return retMap, err

}

func (s *state13) ClaimsMap() (adt.Map, error) {

	return adt13.AsMap(s.store, s.Claims, builtin13.DefaultHamtBitwidth)

}

// TODO this could return an error since not all versions have a claims map
func (s *state13) ClaimsMapBitWidth() int {

	return builtin13.DefaultHamtBitwidth

}

// TODO this could return an error since not all versions have a claims map
func (s *state13) ClaimsMapHashFunction() func(input []byte) []byte {

	return func(input []byte) []byte {
		res := sha256.Sum256(input)
		return res[:]
	}

}

func (s *state13) ClaimMapForProvider(providerIdAddr address.Address) (adt.Map, error) {

	innerHamtCid, err := s.getInnerHamtCid(s.store, abi.IdAddrKey(providerIdAddr), s.Claims, builtin13.DefaultHamtBitwidth)
	if err != nil {
		return nil, err
	}
	return adt13.AsMap(s.store, innerHamtCid, builtin13.DefaultHamtBitwidth)

}

func (s *state13) getInnerHamtCid(store adt.Store, key abi.Keyer, mapCid cid.Cid, bitwidth int) (cid.Cid, error) {

	actorToHamtMap, err := adt13.AsMap(store, mapCid, bitwidth)
	if err != nil {
		return cid.Undef, fmt.Errorf("couldn't get outer map: %x", err)
	}

	var innerHamtCid cbg.CborCid
	if found, err := actorToHamtMap.Get(key, &innerHamtCid); err != nil {
		return cid.Undef, fmt.Errorf("looking up key: %s: %w", key, err)
	} else if !found {
		return cid.Undef, fmt.Errorf("did not find key: %s", key)
	}

	return cid.Cid(innerHamtCid), nil

}
