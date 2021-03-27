package contentbidindex

import (
	"context"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/retrieval-market-spec/docs/components/contentdistribution"
	"github.com/filecoin-project/retrieval-market-spec/docs/components/contentrouting"
	peer "github.com/libp2p/go-libp2p-peer"
)

type ContentBid struct {
	// Address is the address that will sent payments
	Address address.Address
	// Peer is the peer for the content provider
	Peer peer.ID
	// HostingAsk describe the parameters of the content bid
	HostingAsk contentdistribution.HostingAsk
}

type BidSearchParams struct {
	// MinimumPricePerByteServed is the minimum amount paid for each retrieval
	// on a per-byte basis
	MinimumPricePerByteServed abi.TokenAmount
	// MinimumPricePerRequestServed is the minimum amount paid for each
	// retrieval on a per-byte basis
	MinimumPricePerRequestServed abi.TokenAmount
	// MaximumBandwidthBytes indicates the maximum expected transfer speed in
	// bytes
	MaximumBandwidthBytes uint64
	// MinimumTimeToFirstByte specifies the minimum time to serve the first byte
	// (miner latency)
	MaximumTimeToFirstByte time.Duration
	// RequestsServed indicates the range of acceptable requests served
	RequestsServed map[contentrouting.GeographicRegion]contentdistribution.DesiredRequestsServed
}

// ContentBidIndexAPI provides a machinism for making and discovering bids for
// different
// content
type ContentBidIndexAPI interface {
	// MakeBid lists a new bid on the content bid index
	MakeBid(ctx context.Context, bid ContentBid) error
	// ListBids returns all bids on the content bid index (TODO: Practical?)
	ListBids(ctx context.Context) (<-chan ContentBid, error)
	// Search bids searches the content index to find bids that matche the
	// search criteria
	SearchBids(
		ctx context.Context,
		params BidSearchParams,
	) (<-chan ContentBid, error)
}
