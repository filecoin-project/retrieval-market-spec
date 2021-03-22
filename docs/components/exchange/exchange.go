package exchange

import (
	"github.com/filecoin-project/retrieval-market-spec/docs/components/datatransfer"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	peer "github.com/libp2p/go-libp2p-core/peer"
)

type DealID uint64

// ClientDealState is the state of an in progress deal
// fields are dependent on the exchange protocol
type ClientDealState ipld.Node

// ClientEvent is an event that occurs in a deal lifecycle on the client
type ClientEvent uint64

type ClientEventData struct {
	Event ClientEvent
	State ClientDealState
}

// ExchangeProtocolName identifies a given exchange protocol
// This probably will need to be more than a string in actuality,
// this is just a placeholder
type ExchangeProtocolName string

// ExchangeClientAPI is exposed the party paying to receive or send data
type ExchangeClientAPI interface {
	ProtocolName() ExchangeProtocolName

	// RequestData requests data in exchange for payment (a traditional
	// retrieval
	// deal)
	RequestData(
		payloadCID cid.Cid,
		selector ipld.Node,
		parameters ipld.Node,
	) (DealID, error)

	// SendData requests to send someone data in excahgne for payment - probably
	// unused for some time
	SendData(
		payloadCID cid.Cid,
		selector ipld.Node,
		parameters ipld.Node,
	) (DealID, error)

	// CancelDeal attempts to cancel an inprogress deal
	CancelDeal(id DealID) error

	// GetDeal returns a given deal by deal ID, if it exists
	GetDeal(dealID DealID) (ClientDealState, error)

	// ListDeals returns all deals
	ListDeals() (map[DealID]ClientDealState, error)

	// get notified when certain types of events happen
	EventsNotify() <-chan ClientEventData
}

// ProviderDealState is the state of an in progress deal
// fields are dependent on the exchange protocol
type ProviderDealState ipld.Node

// ProviderEvent is an event that occurs in a deal lifecycle on the provider
type ProviderEvent uint64

type ProviderEventData struct {
	Event ProviderEvent
	State ProviderDealState
}

// ProviderDealID is a value that uniquely identifies a deal for a provider
type ProviderDealID struct {
	Receiver peer.ID
	DealID   DealID
}

// Ask specifies the parameters that will be accepted for incoming deals
// the specific fields will depend on the exchange protocol
type Ask ipld.Node

// ExchangeProviderAPI is exposed to the party sending or receiving data in
// exchange for payment
type ExchangeProviderAPI interface {
	ProtocolName() ExchangeProtocolName

	// SetAsk sets the payment parameters for this
	SetAsk(ask Ask)

	// GetAsk returns the retrieval providers pricing information
	GetAsk() Ask

	// get notified when certain types of events happen
	EventsNotify() <-chan ProviderEventData

	// ListDeals shows the in progress deals for this exchange protocol
	ListDeals() map[ProviderDealID]ProviderDealState
}

// ExchangeValidatorAPI is an interface provided to data transfer to validate
// incoming transfers by the party sending or receiving data in exchange for
// payment
type ExchangeValidatorAPI interface {
	// ValidatePush validates an incoming data transfer push request received
	// for
	// this exchange protocol
	ValidatePush(
		sender peer.ID,
		voucher datatransfer.Voucher,
		baseCid cid.Cid,
		selector ipld.Node,
	) (datatransfer.VoucherResult, error)

	// ValidatePull validates an incoming data transfer pull request received
	// for
	// this exchange protocol
	ValidatePull(
		receiver peer.ID,
		voucher datatransfer.Voucher,
		baseCid cid.Cid,
		selector ipld.Node,
	) (datatransfer.VoucherResult, error)
}

// ExchangeCheckpointValidatorAPI is for exchange providers that want to
// checkpoint payments. Data transfer uses this to pause and resume
// This is based on the existing data transfer revalidator interface but
// made slightly more high leve For ease of use
type ExchangeCheckpointValidatorAPI interface {

	// NextCheckPoint specifies the next point at which data transfer should
	// pause transfer
	// return values are:
	// byteOffset = minimum next offset at which checkpoint reached should be
	// called
	// err = abort request if not nil
	NextCheckPoint(
		channelID datatransfer.ChannelID,
	) (byteOffset uint64, err error)

	// CheckpointReached allows an exchange to produce a request for payment
	// after a checkpoint has
	// been passed
	// Note: checkpoint is not exact for some protocols -- the byteOffs may be
	// slightly larger than
	// the one requested
	// return values are:
	// voucher result = information about payment requested, if any
	// err = error
	// - nil = no payment needed, continue request
	// - error = abort this request
	// - err == ErrPause - pause for payment
	CheckpointReached(
		channelID datatransfer.ChannelID,
		actualByteOffset uint64,
	) (datatransfer.VoucherResult, error)

	// EndOfDataReached allows for a final last request for payment
	// return values are:
	// voucher result = information about payment requested, if any
	// err = error
	// - nil = no payment needed, completed request
	// - error = terminate request with error
	// - err == ErrPause - pause for final payment
	EndOfDataReached(
		channelID datatransfer.ChannelID,
		actualByteOffset uint64,
	) (datatransfer.VoucherResult, error)

	// Revalidate processes a new voucher (likely a payment) intended to resume
	// transfer
	// return values are:
	// voucher result = information about why the payment succeeded or failed
	// err = error
	// - nil = resume the request
	// - error = abort this request
	// - err == ErrPause - leave this request paused
	Revalidate(
		channelID datatransfer.ChannelID,
		voucher datatransfer.Voucher,
	) (datatransfer.VoucherResult, error)
}
