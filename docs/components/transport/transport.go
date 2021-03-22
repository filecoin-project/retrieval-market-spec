package transport

import (
	"context"

	datatransfer "github.com/filecoin-project/retrieval-market-spec/docs/components/datatransfer"
	"github.com/ipld/go-ipld-prime"
	peer "github.com/libp2p/go-libp2p-peer"
)

// Capabilities describes what a transport can and cannot do
type Capabilities struct {
	DispatchesDataSentEvents     bool
	DispatchesDataQueuedEvents   bool
	DispatchesDataReceivedEvents bool
	DispatchesTimeoutsEvents     bool
	DispatchesDisconnectedEvents bool
	SupportsRestarts             bool
	SupportsPauseResume          bool
	Verifiable                   bool
	IncrementallyVerifiable      bool
}

type TransportAPI interface {
	Capabilities() Capabilities
	// OpenChannel initiates an outgoing request for the other peer to send data
	// to us on this channel
	OpenChannel(ctx context.Context,
		dataSender peer.ID,
		channelID datatransfer.ChannelID,
		root ipld.Link,
		stor ipld.Node,
		resumeState ipld.Node,
		msg datatransfer.Message) error
	// CloseChannel closes the given channel
	CloseChannel(ctx context.Context, chid datatransfer.ChannelID) error
	// CleanupChannel is called on the otherside of a cancel - removes any
	// associated
	// data for the channel
	CleanupChannel(chid datatransfer.ChannelID)
	PauseChannel(ctx context.Context,
		chid datatransfer.ChannelID,
	) error
	// ResumeChannel resumes the given channel
	ResumeChannel(ctx context.Context,
		msg datatransfer.Message,
		chid datatransfer.ChannelID,
	) error
	Shutdown(ctx context.Context) error
}
