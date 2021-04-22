---
title: 'Content Provider'
description: Content providers pay to host their data so that clients can retrieve it. Content providers list their bids in a Content Marketplace, and Retrieval Providers make offers to fulfill those bids. Content Providers then transfer their data to a Retrieval Provider, who then serves it to a Retrieval Client, usually at no cost to the Retrieval Client. Instead, the Content Provider pays the Retrieval Provider when their data is served.
confidenceLevel: 'Brainstorm'
breadcrumb: 'Content Provider'
v0: |
  There is no current mechanism for a content provider to participate in the retrieval market, other than to use a storage miner and hope that they serve retrievals quickly.
v1: |
  Building a successful MVP for content providing in the retrieval market is likely a year long project. All of the following much be in place for content providing to work:
  
  - We need to build first iterations of the Content Bid Index and Content Distribution components. Both of these components are proposed in this document as brainstorms. They will likely need revision along the way.
  - We need to deliver a 3 way exchange mechanism, where a miner sends data to a client, in exchange for the client verifying to the content provider they received the data, in exchange for the content provider sending payment to the miner. Making sure we do this in a way so that all parties receive fair exchange in a trustless environment may not be easy.
  - We need to find an appropriate piece of software in which to run the content providing system. It's possible go-IPFS is that system. Alternatively, it's possible the first customers of content providing are IPFS pinning serves that want to pay to insure the content people pay them to store is served quickly.
  
  All of this makes for a lot of work for a single year. However, enabling content providers is key to the economics of a retrieval market CDN. People browsing the web expect to browse for free in most cases. To deliver a functioning retrieval market, we need to unlock payments for parties other than the client.
v2: |
  Likely the above set of work may bleed into 2022. Beyond that, the proposed design makes content providers somewhat passive -- they make bids, then wait for miners to offer. It probably makes sense to offer other kinds of negotiation, and more complex ways of putting together hosting packages.
---

<Header />

## Content Provider Components

A content provider must run the following components themselves:

1. Transport(s)
2. Data Transfer
3. Content Distribution (HostingClient)
4. Wallet
5. Exchange Protocol (a "to be developed" 3-way exchange protocol)

A content provider can run the following components locally, but will often instead interact with them remotely through a Payment Provider:
1. Payment Channel Manager
2. Chain

A content provider also makes remote calls to list their bids on a Content Bid Index and may call a Reputational Index to check the reputation of the miners they are storing with

A content provider generally does not interact with Content Routing.

## Roadmap

<RoadMapPage />