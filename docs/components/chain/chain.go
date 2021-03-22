package payments

import (
	"context"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/go-state-types/network"
	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"
)

type TipSetKey struct {
	value string
}

type Ticket struct {
	VRFProof []byte
}

type BeaconEntry struct {
	Round uint64
	Data  []byte
}

type ElectionProof struct {
	WinCount int64
	VRFProof []byte
}

type BlockHeader struct {
	Miner                 address.Address    // 0 unique per block/miner
	Ticket                *Ticket            // 1 unique per block/miner: should be a valid VRF
	ElectionProof         *ElectionProof     // 2 unique per block/miner: should be a valid VRF
	BeaconEntries         []BeaconEntry      // 3 identical for all blocks in same tipset
	WinPoStProof          []proof2.PoStProof // 4 unique per block/miner
	Parents               []cid.Cid          // 5 identical for all blocks in same tipset
	ParentWeight          big.Int            // 6 identical for all blocks in same tipset
	Height                abi.ChainEpoch     // 7 identical for all blocks in same tipset
	ParentStateRoot       cid.Cid            // 8 identical for all blocks in same tipset
	ParentMessageReceipts cid.Cid            // 9 identical for all blocks in same tipset
	Messages              cid.Cid            // 10 unique per block
	BLSAggregate          *crypto.Signature  // 11 unique per block: aggrregate of BLS messages from above
	Timestamp             uint64             // 12 identical for all blocks in same tipset / hard-tied to the value of Height above
	BlockSig              *crypto.Signature  // 13 unique per block/miner: miner signature
	ForkSignaling         uint64             // 14 currently unused/undefined
	ParentBaseFee         abi.TokenAmount    // 15 identical for all blocks in same tipset: the base fee after executing parent tipset

	validated bool // internal, true if the signature has been validated
}

type TipSet struct {
	cids   []cid.Cid
	blks   []*BlockHeader
	height abi.ChainEpoch
}

type Message struct {
	Version uint64

	To   address.Address
	From address.Address

	Nonce uint64

	Value abi.TokenAmount

	GasLimit   int64
	GasFeeCap  abi.TokenAmount
	GasPremium abi.TokenAmount

	Method abi.MethodNum
	Params []byte
}

type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64
}

type MsgLookup struct {
	Message   cid.Cid // Can be different than requested, in case it was replaced, but only gas values changed
	Receipt   MessageReceipt
	ReturnDec interface{}
	TipSet    TipSetKey
	Height    abi.ChainEpoch
}

type MessageSendSpec struct {
	MaxFee abi.TokenAmount
}

type SignedMessage struct {
	Message   Message
	Signature crypto.Signature
}

type Actor struct {
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance abi.TokenAmount
}

type BlockMessages struct {
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage

	Cids []cid.Cid
}

type HeadChange struct {
	Type string
	Val  *TipSet
}

type FilecoinChainAPI interface {
	ChainHead(context.Context) (*TipSet, error)
	ChainHasObj(context.Context, cid.Cid) (bool, error)
	ChainReadObj(context.Context, cid.Cid) ([]byte, error)
	ChainGetBlockMessages(context.Context, cid.Cid) (*BlockMessages, error)
	ChainGetMessage(ctx context.Context, mc cid.Cid) (*Message, error)
	ChainGetTipSet(ctx context.Context, tsk TipSetKey) (*TipSet, error)
	ChainGetTipSetByHeight(
		ctx context.Context,
		h abi.ChainEpoch,
		tsk TipSetKey,
	) (*TipSet, error)
	ChainNotify(context.Context) (<-chan []*HeadChange, error)
	GasEstimateMessageGas(
		ctx context.Context,
		msg *Message,
		spec *MessageSendSpec,
		tsk TipSetKey,
	) (*Message, error)
	MpoolPush(ctx context.Context, sm *SignedMessage) (cid.Cid, error)
	StateAccountKey(
		ctx context.Context,
		addr address.Address,
		tsk TipSetKey,
	) (address.Address, error)
	StateGetActor(
		ctx context.Context,
		actor address.Address,
		ts TipSetKey,
	) (*Actor, error)
	StateGetReceipt(
		context.Context,
		cid.Cid,
		TipSetKey,
	) (*MessageReceipt, error)
	StateLookupID(
		ctx context.Context,
		addr address.Address,
		tsk TipSetKey,
	) (address.Address, error)
	StateNetworkVersion(context.Context, TipSetKey) (network.Version, error)
	StateWaitMsg(
		ctx context.Context,
		msg cid.Cid,
		confidence uint64,
	) (*MsgLookup, error)
	StateSearchMsg(ctx context.Context, msg cid.Cid) (*MsgLookup, error)
}
