package minerindex

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/retrieval-market-spec/docs/components/contentrouting"
	"github.com/filecoin-project/retrieval-market-spec/docs/components/exchange"
	"github.com/ipld/go-ipld-prime"
	"github.com/libp2p/go-libp2p-core/peer"
)

// MinerRecord is a record of a single retrieval miner on the network
// It contains basic information about the miner as well as
type MinerRecord struct {
	// Address is the address for sending payments
	Address address.Address
	// Peer is the peer for this miner
	Peer peer.ID
	// ExchangeOptions are the available protocols for this miner
	ExchangeOptions []contentrouting.ExchangeOption
	// Regions are the regions this miner advertises serving
	Regions []contentrouting.GeographicRegion
	// SuccessRate is the decimal success rate for making retrieval deals with
	// this miner
	SuccessRate big.Int
	// BandwidthAvgBytes is a measurement of the miners average bandwidth
	BandwidthAvgBytes uint64
}

// ExchangeSearchOption specifies desired parameters for a given protocol
type ExchangeSearchOption struct {
	ProtocolName exchange.ExchangeProtocolName
	SearchParams ipld.Node
}

// SearchParams configures a search for an appropriate miner
type SearchParams struct {
	AcceptedExchanges    []ExchangeSearchOption
	AcceptedRegions      []contentrouting.GeographicRegion
	MinSuccessRate       big.Int
	MinBandwidthAvgBytes uint64
}

// MinerIndexAPI provides information to prospective retrieval clients and
// content distributors about miners on the network
type MinerIndexAPI interface {
	SearchMiners() (<-chan MinerRecord, error)
	ListMiners() (<-chan MinerRecord, error)
}
