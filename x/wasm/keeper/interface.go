package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CosmWasm/wasmd/x/wasm/types"
)

type WasmKeeper interface {
	IterateContractInfo(ctx sdk.Context, cb func(sdk.AccAddress, types.ContractInfo) bool)
	StoreKey() storetypes.StoreKey
	CDC() codec.Codec
	StoreCodeInfo(ctx sdk.Context, codeID uint64, codeInfo types.CodeInfo)
	AddToContractCreatorSecondaryIndex(ctx sdk.Context, creatorAddress sdk.AccAddress, position *types.AbsoluteTxPosition, contractAddress sdk.AccAddress)
	ClassicAddressGenerator() AddressGenerator
	GetAuthority() string
	Instantiate(ctx sdk.Context, codeID uint64, creator, admin sdk.AccAddress, initMsg []byte, label string, deposit sdk.Coins, addressGenerator AddressGenerator, authPolicy types.AuthorizationPolicy) (sdk.AccAddress, []byte, error)
	Execute(ctx sdk.Context, contractAddress, caller sdk.AccAddress, msg []byte, coins sdk.Coins) ([]byte, error)
	Migrate(ctx sdk.Context, contractAddress, caller sdk.AccAddress, newCodeID uint64, msg []byte, authZ types.AuthorizationPolicy) ([]byte, error)
	SetContractAdmin(ctx sdk.Context, contractAddress, caller, newAdmin sdk.AccAddress, authZ types.AuthorizationPolicy) error
	SetAccessConfig(ctx sdk.Context, codeID uint64, caller sdk.AccAddress, newConfig types.AccessConfig, authz types.AuthorizationPolicy) error
	PinCode(ctx sdk.Context, codeID uint64) error
	UnpinCode(ctx sdk.Context, codeID uint64) error
	Sudo(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
	Create(ctx sdk.Context, creator sdk.AccAddress, wasmCode []byte, instantiateAccess *types.AccessConfig, authZ types.AuthorizationPolicy) (codeID uint64, checksum []byte, err error)
	GetParams(ctx sdk.Context) types.Params
	SetParams(ctx sdk.Context, params types.Params) error
	SetContractInfoExtension(ctx sdk.Context, contract sdk.AccAddress, extra types.ContractInfoExtension) error
	PropagateGovAuthorization() map[types.AuthorizationPolicyAction]struct{}
	ImportCode(ctx sdk.Context, codeID uint64, codeInfo types.CodeInfo, wasmCode []byte) error
	ImportContract(ctx sdk.Context, contractAddr sdk.AccAddress, c *types.ContractInfo, state []types.Model, entries []types.ContractCodeHistoryEntry) error
	ImportAutoIncrementID(ctx sdk.Context, sequenceKey []byte, val uint64) error
	PeekAutoIncrementID(ctx sdk.Context, sequenceKey []byte) uint64
	HasContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) bool
	GetByteCode(ctx sdk.Context, codeID uint64) ([]byte, error)
	IsPinnedCode(ctx sdk.Context, codeID uint64) bool
	GetContractHistory(ctx sdk.Context, contractAddr sdk.AccAddress) []types.ContractCodeHistoryEntry
	types.ViewKeeper
	QueryGasLimit() uint64
}

func (k Keeper) ImportAutoIncrementID(ctx sdk.Context, sequenceKey []byte, val uint64) error {
	return k.importAutoIncrementID(ctx, sequenceKey, val)
}

func (k Keeper) ImportContract(ctx sdk.Context, contractAddr sdk.AccAddress, c *types.ContractInfo, state []types.Model, entries []types.ContractCodeHistoryEntry) error {
	return k.importContract(ctx, contractAddr, c, state, entries)
}

func (k Keeper) ImportCode(ctx sdk.Context, codeID uint64, codeInfo types.CodeInfo, wasmCode []byte) error {
	return k.importCode(ctx, codeID, codeInfo, wasmCode)
}

func (k Keeper) SetContractInfoExtension(ctx sdk.Context, contractAddr sdk.AccAddress, ext types.ContractInfoExtension) error {
	return k.setContractInfoExtension(ctx, contractAddr, ext)
}
func (k Keeper) PropagateGovAuthorization() map[types.AuthorizationPolicyAction]struct{} {
	return k.propagateGovAuthorization
}

func (k Keeper) Create(ctx sdk.Context, creator sdk.AccAddress, wasmCode []byte, instantiateAccess *types.AccessConfig, authZ types.AuthorizationPolicy) (codeID uint64, checksum []byte, err error) {
	return k.create(ctx, creator, wasmCode, instantiateAccess, authZ)
}

func (k Keeper) PinCode(ctx sdk.Context, codeID uint64) error {
	return k.pinCode(ctx, codeID)
}

func (k Keeper) UnpinCode(ctx sdk.Context, codeID uint64) error {
	return k.unpinCode(ctx, codeID)
}

func (k Keeper) SetContractAdmin(ctx sdk.Context, contractAddress, caller, newAdmin sdk.AccAddress, authZ types.AuthorizationPolicy) error {
	return k.setContractAdmin(ctx, contractAddress, caller, newAdmin, authZ)
}

func (k Keeper) SetAccessConfig(ctx sdk.Context, codeID uint64, caller sdk.AccAddress, newConfig types.AccessConfig, authz types.AuthorizationPolicy) error {
	return k.setAccessConfig(ctx, codeID, caller, newConfig, authz)
}

func (k Keeper) Migrate(ctx sdk.Context, contractAddress, caller sdk.AccAddress, newCodeID uint64, msg []byte, authZ types.AuthorizationPolicy) ([]byte, error) {
	return k.migrate(ctx, contractAddress, caller, newCodeID, msg, authZ)
}

func (k Keeper) Instantiate(ctx sdk.Context, codeID uint64, creator, admin sdk.AccAddress, initMsg []byte, label string, deposit sdk.Coins, addressGenerator AddressGenerator, authPolicy types.AuthorizationPolicy) (sdk.AccAddress, []byte, error) {
	return k.instantiate(ctx, codeID, creator, admin, initMsg, label, deposit, addressGenerator, authPolicy)
}

func (k Keeper) Execute(ctx sdk.Context, contractAddress, caller sdk.AccAddress, msg []byte, coins sdk.Coins) ([]byte, error) {
	return k.execute(ctx, contractAddress, caller, msg, coins)
}

func (k Keeper) StoreKey() storetypes.StoreKey {
	return k.storeKey
}

func (k Keeper) CDC() codec.Codec {
	return k.cdc
}

func (k Keeper) StoreCodeInfo(ctx sdk.Context, codeID uint64, codeInfo types.CodeInfo) {
	k.storeCodeInfo(ctx, codeID, codeInfo)
}

func (k Keeper) AddToContractCreatorSecondaryIndex(ctx sdk.Context, creatorAddress sdk.AccAddress, position *types.AbsoluteTxPosition, contractAddress sdk.AccAddress) {
	k.addToContractCreatorSecondaryIndex(ctx, creatorAddress, position, contractAddress)
}
