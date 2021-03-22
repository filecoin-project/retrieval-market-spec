package contentdistribution

import (
	"time"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/retrieval-market-spec/docs/components/contentrouting"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/libp2p/go-libp2p-core/peer"
)

// ClientHostingDealState is the state of an in progress hosting deal
// fields TBD
type ClientHostingDealState ipld.Node

// ClientHostingEvent is an event that occurs in a hosting deal lifecycle on the client
type ClientHostingEvent uint64

type ClientHostingEventData struct {
	Event ClientHostingEvent
	State ClientHostingDealState
}

// ClientHostingDealID is a value that uniquely identifies a hosting deal for a client
type ClientHostingDealID struct {
	Provider      peer.ID
	HostingDealID HostingDealID
}

type DesiredRequestsServed struct {
	Min uint64
	Max uint64
}

type HostingAsk struct {
	// Root is the root CID of the content the provider would like to make available
	Root cid.Cid
	// Selector indicates the rest of the DAG of the content the provider would like to make available
	Selector ipld.Node
	// PricePerByteServed is amount paid for each retrieval on a per-byte basis
	PricePerByteServed abi.TokenAmount
	// PricePerRequestServed is amount paid for each retrieval on a per-byte basis
	PricePerRequestServed abi.TokenAmount
	// MinimumBandwidthBytes indicates the expected transfer speed in bytes
	DesiredAvgBandwidthBytes uint64
	// MaximumTimeToFirstByte specifies the maximum miner latency
	MaximumTimeToFirstByte time.Duration
	// RequestsServed indicates how many requests you are bidding on by Geographic region
	RequestsServed map[contentrouting.GeographicRegion]DesiredRequestsServed
	// ExpiryTime indicates when this deal will go stale
	ExpiryTime time.Time
}

type HostingClientAPI interface {

	// SetAsks sets the list of content the provider would like to pay to host
	SetAsks(asks []HostingAsk)

	// GetAsks returns the content providers content and what they're willing to pay for it.
	GetAsks() []HostingAsk

	// get notified when certain types of events happen
	EventsNotify() <-chan ClientHostingEventData

	// ListDeals shows the in progress deals for this exchange protocol
	ListDeals() map[ClientHostingDealID]ClientHostingDealState
}

type HostingOffer struct {
	// Root is the root CID of the content the provider would like to make available
	Root cid.Cid
	// Selector indicates the rest of the DAG of the content the provider would like to make available
	Selector ipld.Node
	// PricePerByteServed is amount paid for each retrieval on a per-byte basis
	PricePerByteServed abi.TokenAmount
	// PricePerRequestServed is amount paid for each retrieval on a per-byte basis
	PricePerRequestServed abi.TokenAmount
	// MinimumBandwidthBytes indicates the expected transfer speed in bytes
	// TODO: is this possible to actually measure?
	MinimumAvgBandwidthBytes uint64
	// MaximumTimeToFirstByte specifies the maximum miner latency
	// TODO: is this possible to actually measure?
	MaximumTimeToFirstByte time.Duration
	// RequestsServed indicates how many requests will be served  by region
	RequestsServed map[contentrouting.GeographicRegion]uint64
}

type HostingDealID uint64

// ProviderHostingDealState is the state of an in progress hosting deal
// fields TBD
type ProviderHostingDealState ipld.Node

// ProviderHostingEvent is an event that occurs in a deal lifecycle on the provider
type ProviderHostingEvent uint64

type ProviderHostingEventData struct {
	Event ProviderHostingEvent
	State ProviderHostingDealState
}

type HostingProviderAPI interface {
	// OfferHosting initiates a new request to host data based on parameters discovered on
	// the content bid index. It initiates a new data transfer request to the content provider to
	// send content to host
	// TODO: Is this ok? Is the content provider possibly behind a firewall?
	OfferHostingDeal(p peer.ID, offer HostingOffer) (HostingDealID, error)

	// GetDeal returns a given deal by deal ID, if it exists
	GetHostingDeal(dealID HostingDealID) (ProviderHostingDealState, error)

	// ListDeals returns all deals
	ListHostingDeals() (map[HostingDealID]ProviderHostingDealState, error)

	// get notified when certain types of events happen
	EventsNotify() <-chan ProviderHostingEventData
}
