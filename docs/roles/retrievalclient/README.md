---
title: 'Retrieval Client'
description: 'Retrieval Clients pay to retrieve content. They find content through the a Marketplace Provider, and initiate Data Transfer through an Exchange Protocol to retrieve data from a Retrieval Provider/Miner. For paid retrievals, a Retrieval Client must have a Wallet and a connection to either a local or a remote instance of the Filecoin Chain'
confidenceLevel: 'Draft/WIP'
breadcrumb: 'Retrieval Client'
v0: 
  Currently, the only instance of a retrieval client is a Lotus Node, or a Lotus Lite Node combined with a Lotus Gateway. This makes for a very high bar to entry. Running a full Lotus Node is often impossible even on a powerful computer, and even running Lotus Lite you're cut off from much of the IPFS stack unless you also run PowerGate.
v05: |
  Why's [Estuary](https://github.com/application-research/estuary) is a prototype instance of a retrieval client taken out of the Lotus stack. It will be interesting to see where this heads in the future
v1: |
  Running retrieval through go-ipfs might produce a smaller, more functional retrieval client. To do this we would:
  - Put go-data-transfer in go-ipfs
  - Connect go-ipfs to index-style Content Routing for Filecoin
  - Implement clear rules for negotiation between IPFS transfer vs Exchange based paid retrieval. Figure out how if at all IPFS will interface with a Miner Index to determine the most reputable miners to pull from.
  - To keep go-ipfs open and nuetral, we may want leave Filecoin specific exchange out of the go-ipfs binary. Instead, we could allow IPFS to communicate with any Exchange Client running locally but not in the same process (enabling potentially non-Filecoin paid retrievals)
  - The smallest version of this second process consists of an Exchange Client and a Wallet, talking to a remote Payment Provider, if such a provider exists.
  - A slightly larger version would run a Payment Channel Manager locally but talk to a remote Chain

  **Alterative 1**: Myel is exploring support for retrieval in IPFS through a plugin that runs Data Transfer and a custom Exchange/Content Routing protocol to turn IPFS nodes into both Retrieval Clients and Retrieval Providers

  **Alternative 2**: build a custom go IPFS/Filecoin retrieval client outside of go-ipfs. A single process IPFS/Filecoin node build from the group up might prove lighter and more nimble than trying to build on top of go-IPFS.

v2: |
  The biggest pool of potential clients for the Retrieval Market are web browsers. To get retrieval running in the browser however, we must first convert several parts of the Filecoin stack currently in Go to javascript (or Web Assembly):
  - The following components currently have no JS equivalent: Data Transfer, Exchange, and a Wallet.
  - We'd need to implement a Transport, but it's possible we could use STP, an in development version of Transport based on HTTP, that might be easier to implement in Javascript
  - Hopefully, by the time we built retrieval in the browser, the JS client could talk to Payment Providers and MarketPlace Providers to access the remaining components remotely
---

<Header />

## Retrieval Client Components

A retrieval client must run the following components themselves:

1. Transport(s)
2. Data Transfer
3. Exchange Protocol(s) (Client)
4. Wallet

A retrieval client can run the following components locally, but will often instead interact with them remotely through a different provider:
1. Content Routing
2. Payment Channel Manager
3. Chain

A retrieval client may also access a remote Miner Index to identify Retrieval Providers with good reputations.

A retrieval client generally does not interact with the Content Bid Index or Content Distribution components.

## Roadmap

<RoadMapPage />
