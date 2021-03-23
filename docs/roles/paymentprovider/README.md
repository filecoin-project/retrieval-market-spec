---
title: 'Payment Provider'
description: 
  Payment Providers provide access to a payment system (i.e the Filecoin Chain + Payment Channels) via which people can pay for retrievals. Payments are the only part of the retrieval market that interfaces with a blockchain. As major "hub and spoke" payment providers emerge, actual interactions with the Filecoin Blockchain may simply go through them. Everyone else would then be able to participate without having to run a full Filecoin validator node.
confidenceLevel: 'Draft/WIP'
breadcrumb: 'Payment Provider'
v0: 
  There is currently no way for someone who wants to participate as a payment provider in the retrieval market to make money doings so. Instead, payments are run directly on client and provider nodes.
v05: |
  The Filecoin blockchain contains primitives to support more complex payment channel schemes like "hub and spoke". However, none of these primitives have ever fully been used on mainnet. Before any work towards more robust payment provider system begins in earnest, the first step is exercise the Filecoin chain primitives and determine what, if anything, needs to change to support complex systems of payment in the retrieval market. If changes are required, they would have to go through a FIP process and an actors upgrade, so it makes sense to start early.
v1: |
  The value of more complex payment systems is that they scale better than a system where every time two parties transact they setup a direct channel. The moment when this scaling becomes a neccesity for the retrieval market is a bit hard to predict. In the immediate sense, more can be done to unblock the existing simple PaymentChannelManager. At the same time, if retrievals are ever to compete in the CDN space, we simply can't block on chain operations, ever, in the context of retrieval.
v2: |
  The long term question in terms of payments is how the chain itself can incentivize retrieval. This is an semi-open research question, which it probably makes sense to devote time to sooner rather than later, even if actual development doesn't start for a year or two. 
---

<Header />

## Payment Provider Components

A payment provider must run the following components themselves:

1. Payment Channel Manager
2. Chain
3. Wallet (their own)

A payment provider will often interact with several remote wallets. Determining the appropriate authentication scheme to enable this is an open problem.

A payment provider generally does not interact with the Content Bid Index,  Content Distribution, Transport, Data Transfer, Exchange, Content Routing, or Miner Index components. Many of these components utilize a Payment Provider to make financial transactions. But from the payment provider's point of view, it doesn't know anything about the nature of the financial transactions that run through it.

## Roadmap

<RoadMapPage />