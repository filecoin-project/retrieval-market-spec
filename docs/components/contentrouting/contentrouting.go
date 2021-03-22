package contentrouting

import (
	"context"

	"github.com/filecoin-project/retrieval-market-spec/docs/components/exchange"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/libp2p/go-libp2p-core/peer"
)

type ExchangeOption struct {
	// Name of exchange protocol
	ProtocolName exchange.ExchangeProtocolName
	// Offer parameters
	Parameters ipld.Node
}

// GeographicRegion content is located in
// TODO: what should this look like?
type GeographicRegion ipld.Node

// ContentProvider is someone who is providing the content requested
// on the given exchange and parameters
type ContentProvider struct {
	Peer         peer.ID
	ProtocolName exchange.ExchangeProtocolName
	Parameters   ipld.Node
	Region       GeographicRegion
}

// ContentRoutingAPI provides mechanisms to find providers of content,
// and offer content to the network
type ContentRouting interface {
	// ProvideCID is called by a provider to announce that they are offering
	// the given content to a network, on the given exchange protocols
	// and parameters
	ProvideCID(ctx context.Context,
		cid cid.Cid,
		regions []GeographicRegion,
		exchangeOptions []ExchangeOption) error

	// ProvideDAG is called by a provider to announce that they are offering
	// a DAG, expressed by a CID+Selector, to a network, on the given exchange
	//  protocols and parameters
	ProvideDAG(ctx context.Context,
		cid cid.Cid,
		selector ipld.Node,
		regions []GeographicRegion,
		exchangeOptions []ExchangeOption) error

	// Search for peers who are able to provide the given cid
	//
	// When count is 0, this method will return an unbounded number of
	// results.
	FindProvidersCID(ctx context.Context,
		cid cid.Cid,
		desiredRegion GeographicRegion,
		acceptedExchanges []exchange.ExchangeProtocolName,
		count int) (<-chan ContentProvider, error)

	// Search for peers who are able to provide the given DAG, as expressed
	// by CID + Selector
	//
	// When count is 0, this method will return an unbounded number of
	// results.
	FindProvidersDAG(ctx context.Context,
		cid cid.Cid,
		selector ipld.Node,
		desiredRegion GeographicRegion,
		acceptedExchanges []exchange.ExchangeProtocolName,
		count int) (<-chan ContentProvider, error)
}
