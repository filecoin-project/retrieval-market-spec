---
title: 'Content Routing'
description: 'Content Routing is the mechanism by which parties looking for content in the network identify parties providing that content'
confidenceLevel: Brainstorm
breadcrumb: 'Content Routing'
v0: |
  At the moment there are several in progress proposals and prototypes for content routing:
  - ChainSafe built [a prototype of content-routing based on GossipSub](https://github.com/ChainSafe/fil-secondary-retrieval-markets)
  - Pegasys has proposed [a DHT based content routing system](https://docs.google.com/document/d/13pWpEU7wmWR0t4fnrGqKDnG-idiDrJwfBDi7z366Dfg/edit#heading=h.xl8ulbm31o3h)
  - ResNetLab is currently researching major changes to IPFS's DHT

  All of the current prototypes diverge quite a bit from the API interface described below.
v05: >
  The simplest path towards building a functioning prototype for content-routing for the retrieval market is to build a global index of content (i.e. a gateway). There are some very clear directions we can go to assemble such an index:
  - we can examine the Filecoin chain to look at current deals
  - we can encourage miners to use the ProvideCID/Provide DAG API calls to post CIDs for data they are storing
  - we can create retrieval deals for payload CIDs discovered on the Filecoin chain, then add other CIDs contained in the retrieved piece to the index
v1:
  If we can build a functioning index/gateway, as true "retrieval miners" come online, they can also use the existing APIs continue to augment the index above.
v2:
  An index will likely hit scaling limits fairly quickly. In a year, hopefully ResNetLab's content routing efforts will produce clear next steps for more scalable solutions
---

<Header />

## Dependencies

| Name | Kind | APIs |
| ---- | ---- | ---- |
| Chain | optional | default |
| Miner Index | optional | default |

## Roadmap

<RoadMapPage />

## Preliminary API

This API is purely speculative, based on the likely needs for people seeking content they can retrieve quickly

::: tip
ResNetLab is working on a new content routing system. The new interface will be very general -- allowing storage of fairly arbitrary data on composoble routing platforms (DHT, index, etc). When their work done, we will either switch to use their interface directly or build the interface described below on top of it.
:::

<<< @/docs/components/contentrouting/contentrouting.go
