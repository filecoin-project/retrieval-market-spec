package datatransfer

import (
	"context"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	peer "github.com/libp2p/go-libp2p-peer"
)

// TypeIdentifier is a unique string identifier for a type of encodable object
// in a
// registry
type TypeIdentifier string

// Registerable is a type of object in a registry. It must be encodable and must
// have a single method that uniquely identifies its type
type Registerable struct {
	Data ipld.Node // cbor-encoded IPLD
	// Type is a unique string identifier for this voucher type
	Type TypeIdentifier
}

// Voucher is used to validate
// a data transfer request against the underlying storage or retrieval deal
// that precipitated it. The only requirement is a voucher can read and write
// from bytes, and has a string identifier type
type Voucher Registerable

// VoucherResult is used to provide option additional information about a
// voucher being rejected or accepted
type VoucherResult Registerable

// TransferID is an identifier for a data transfer, shared between
// request/responder and unique to the requester
type TransferID uint64

// ChannelID is a unique identifier for a channel, distinct by both the other
// party's peer ID + the transfer ID
type ChannelID struct {
	Initiator peer.ID
	Responder peer.ID
	ID        TransferID
}

type EventCode int

// Event is a struct containing information about a data transfer event
type Event struct {
	Code      EventCode // What type of event it is
	Message   string    // Any clarifying information about the event
	Timestamp time.Time // when the event happened
}

type EventData struct {
	event        Event
	channelState ChannelState
}

type Status uint64

// ChannelState is the internal representation on disk for the channel fsm
type ChannelState struct {
	// PeerId of the manager peer
	SelfPeer peer.ID
	// an identifier for this channel shared by request and responder, set by
	// requester through
	// protocol
	TransferID TransferID
	// Initiator is the person who intiated this datatransfer request
	Initiator peer.ID
	// Responder is the person who is responding to this datatransfer request
	Responder peer.ID
	// base CID for the piece being transferred
	BaseCid cid.Cid
	// cbor-encoded IPLD selector node
	Selector []byte
	// the party that is sending the data (not who initiated the request)
	Sender peer.ID
	// the party that is receiving the data (not who initiated the request)
	Recipient peer.ID
	// expected amount of data to be transferred
	TotalSize uint64
	// current status of this deal
	Status Status
	// total bytes read from this node and queued for sending (0 if receiver)
	Queued uint64
	// total bytes sent from this node (0 if receiver)
	Sent uint64
	// total bytes received by this node (0 if sender)
	Received uint64
	// more informative status on a channel
	Message        string
	Vouchers       []Voucher
	VoucherResults []VoucherResult
}

type Message struct {
	IsRequest bool

	Request  *Request
	Response *Response
}

type MessageType uint64

type Request struct {
	BaseCID        cid.Cid
	Type           MessageType
	IsPaused       bool
	IsPull         bool
	Selector       ipld.Node
	Voucher        Voucher
	TransferID     TransferID
	RestartChannel ChannelID
}

type Response struct {
	Type       MessageType
	Accepted   bool
	IsPaused   bool
	Voucher    Voucher
	TransferID TransferID
}

// ManagerAPI is the core interface for interacting with the data transfer
// system
type ManagerAPI interface {

	// open a data transfer that will send data to the recipient peer and
	// transfer parts of the piece that match the selector
	OpenPushDataChannel(
		ctx context.Context,
		to peer.ID,
		voucher Voucher,
		baseCid cid.Cid,
		selector ipld.Node,
	) (ChannelID, error)

	// open a data transfer that will request data from the sending peer and
	// transfer parts of the piece that match the selector
	OpenPullDataChannel(
		ctx context.Context,
		to peer.ID,
		voucher Voucher,
		baseCid cid.Cid,
		selector ipld.Node,
	) (ChannelID, error)

	// send an intermediate voucher as needed when the receiver sends a request
	// for revalidation
	SendVoucher(ctx context.Context, chid ChannelID, voucher Voucher) error

	// close an open channel (effectively a cancel)
	CloseDataTransferChannel(ctx context.Context, chid ChannelID) error

	// pause a data transfer channel (only allowed if transport supports it)
	PauseDataTransferChannel(ctx context.Context, chid ChannelID) error

	// resume a data transfer channel (only allowed if transport supports it)
	ResumeDataTransferChannel(ctx context.Context, chid ChannelID) error

	// get status of a transfer
	TransferChannelStatus(ctx context.Context, x ChannelID) Status

	// get notified when certain types of events happen
	EventsNotify() <-chan EventData

	// get all in progress transfers
	InProgressChannels(ctx context.Context) (map[ChannelID]ChannelState, error)
}

// EventAPI is used by the transport interface to tell data transfer about
// events happening on the transport
type EventAPI interface {
	// OnChannelOpened is called when we ask the other peer to send us data on
	// the
	// given channel ID
	// return values are:
	// - nil = this channel is recognized
	// - error = ignore incoming data for this channel
	OnChannelOpened(chid ChannelID) error
	// OnResponseReceived is called when we receive a response to a request
	// - nil = continue receiving data
	// - error = cancel this request
	OnResponseReceived(chid ChannelID, msg Response) error
	// OnDataReceive is called when we receive data for the given channel ID
	// return values are:
	// - nil = proceed with sending data
	// - error = cancel this request
	// - err == ErrPause - pause this request
	OnDataReceived(chid ChannelID, link ipld.Link, size uint64) error

	// OnDataQueued is called when data is queued for sending for the given
	// channel ID
	// return values are:
	// message = data transfer message along with data
	// err = error
	// - nil = proceed with sending data
	// - error = cancel this request
	// - err == ErrPause - pause this request
	OnDataQueued(chid ChannelID, link ipld.Link, size uint64) (Message, error)

	// OnDataSent is called when we send data for the given channel ID
	OnDataSent(chid ChannelID, link ipld.Link, size uint64) error

	// OnRequestReceived is called when we receive a new request to send data
	// for the given channel ID
	// return values are:
	// message = data transfer message along with reply
	// err = error
	// - nil = proceed with sending data
	// - error = cancel this request
	// - err == ErrPause - pause this request (only for new requests)
	// - err == ErrResume - resume this request (only for update requests)
	OnRequestReceived(chid ChannelID, msg Request) (Response, error)
	// OnResponseCompleted is called when we finish sending data for the given
	// channel ID
	// Error returns are logged but otherwise have not effect
	OnChannelCompleted(chid ChannelID, success bool) error

	// OnRequestTimedOut is called when a request we opened (with the given
	// channel Id) to receive
	// data times out.
	// Error returns are logged but otherwise have no effect
	OnRequestTimedOut(ctx context.Context, chid ChannelID) error

	// OnRequestDisconnected is called when a network error occurs in a
	// graphsync request
	// or we appear to stall while receiving data
	OnRequestDisconnected(ctx context.Context, chid ChannelID) error
}
