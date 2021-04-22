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

// ContentRoutingAPI provides mechanisms to find providers of content
type ContentRoutingAPI interface {

	// Search for peers who are able to provide the given cid
	//
	// When count is 0, this method will return an unbounded number of
	// results.
	FindProvidersCID(ctx context.Context,
		cid cid.Cid,
		desiredRegion GeographicRegion,
		acceptedExchanges []exchange.ExchangeProtocolName,
		count int) (<-chan ContentProvider, error)
}

// DAGContentRoutingAPI provides mechanisms to find providers of content by
// CID+Selector
type DAGContentRoutingAPI interface {
	ContentRoutingAPI
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

// ContentProvidingAPI offers tools for providers to annouce their content to
// the network
type ContentProvidingAPI interface {
	// ProvideCID is called by a provider to announce that they are offering
	// the given content to a network, on the given exchange protocols
	// and parameters
	ProvideCID(ctx context.Context,
		cid cid.Cid,
		regions []GeographicRegion,
		exchangeOptions []ExchangeOption) error
}

type PrioritizedCID struct {
	Priority uint64
	Cid      cid.Cid
}

// BatchConcentProvidingAPI offers tools for providers to announce several CIDs at once,
// including identifying those that are most likely semantically important
type BatchConcentProvidingAPI interface {
	ContentProvidingAPI

	// ProvideCIDs is called by a provider to announce that they are offering
	// a set of CIDS to the network
	ProvideCIDs(ctx context.Context,
		prioritizedCIDs []PrioritizedCID,
		regions []GeographicRegion,
		exchangeOptions []ExchangeOption) error
}

// DAGContentProvidingAPI offers tools for providers to annouce their content to
// the network by CID+Selector
type DAGContentProvidingAPI interface {
	ContentProvidingAPI

	// ProvideDAG is called by a provider to announce that they are offering
	// a DAG, expressed by a CID+Selector, to a network, on the given exchange
	//  protocols and parameters
	ProvideDAG(ctx context.Context,
		cid cid.Cid,
		selector ipld.Node,
		regions []GeographicRegion,
		exchangeOptions []ExchangeOption) error
}
