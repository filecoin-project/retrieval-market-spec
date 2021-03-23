---
title: 'Marketplace Provider'
description: 
  MarketPlace Providers help other parties in he market find each other. MarketPlace Providers can host Content Bid Indexes, Miner Indexes, and Content Routing indexes. This allows clients to identify what content is hosted by what miners, and which miners will serve content quickest at the lowest price. It enables content providers to advertise content they want miners to host. And it enables miners to find the best and most profitable content to host.
confidenceLevel: 'Brainstorm'
breadcrumb: 'Marketplace Provider'
v0: |
  There is today no MarketPlace provider in the Filecoin Retrieval Market. The immediate task in front of us to enable a MarketPlace is to build the most critical indexes: first the Miner Index (for reputation) and second the initial Content Routing index to help Filecoin users find content.
v05: |
  Building out the IPFS Gateway to reliably locate and serve Filecoin content will lay the foundation for a robust MarketPlace Provider in the future. We have a number of tasks to get there:
  - identifying content in the network
  - benchmarking miners
  - unlocking these indexes not just on the IPFS Gateway but to regular Filecoin or IPFS clients.
v1: |
  The remaining task for the "MarketPlace MVP" is to build a Content Bids Index, and to provide a way for others to run other Marketplaces to begin to lay the foundations for decentalization.
v2: |
  Centralized indexing is likely a strategy with a limited lifespan as the network grows. The next steps is to bring in the results of ResNetLabs research to scale Content Routing, and likely there will be similar efforts neccessary around Miner and Content Bid Indexes.
---

<Header />

## Marketplace Provider Components

A MarketPlace Provider generally runs the following components themselves:

1. Miner Index
2. Content Bids Index

The MarketPlace Provider may also run an indexed version of Content Routing, though over time Content Routing may be come too large to run as a centralized index. Even so, a MarketPlace Provider might run localized index to augment a slower routing mechanism like a global DHT.

As Content Distribution evolves, MarketPlace Providers may be involved in executing automatic hosting deals and other more advanced deal making.

The MarketPlace may monitor a Chain to help assemble its indexes.

The MarketPlace Provider generally does not interact with the Transport, Data Transfer, Exchange, Payment Channel Manager, or Wallet components.

## Roadmap

<RoadMapPage />

::: tip
You've made it to the end of the Roles section. If you started with Roles, now's a good time to drop down to a lower level and get technical with [Components](../../components/)
:::
