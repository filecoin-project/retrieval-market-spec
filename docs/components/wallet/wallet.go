package payments

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)

type MsgType string

const (
	MTUnknown = "unknown"

	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes
	MTChainMsg = "message"

	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is
	// empty)
	MTBlock = "block"

	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra
	// is empty)
	MTDealProposal = "dealproposal"

	// TODO: Deals, Vouchers, VRF
)

type MsgMeta struct {
	Type MsgType

	// Additional data related to what is signed. Should be verifiable with the
	// signed bytes (e.g. CID(Extra).Bytes() == toSign)
	Extra []byte
}

// KeyType defines a type of a key
type KeyType string

const (
	KTBLS             KeyType = "bls"
	KTSecp256k1       KeyType = "secp256k1"
	KTSecp256k1Ledger KeyType = "secp256k1-ledger"
)

// KeyInfo is used for storing keys in KeyStore
type KeyInfo struct {
	Type       KeyType
	PrivateKey []byte
}

type WalletAPI interface {
	WalletNew(context.Context, KeyType) (address.Address, error)
	WalletHas(context.Context, address.Address) (bool, error)
	WalletList(context.Context) ([]address.Address, error)

	WalletSign(
		ctx context.Context,
		signer address.Address,
		toSign []byte,
		meta MsgMeta,
	) (*crypto.Signature, error)

	WalletExport(context.Context, address.Address) (*KeyInfo, error)
	WalletImport(context.Context, *KeyInfo) (address.Address, error)
	WalletDelete(context.Context, address.Address) error
}
