---
title: 'Filecoin Chain'
description: 'A chain contains functionality to interact with the Filecoin blockchain as needed for retrieval, and includes both reading state and pushing chain messages. For the purposes of retrieval, we may need to send funds between actors, and interact with filecoin payment channels. The rest of the chain is not used dirrectly by retrieval now. Any validator node on the Filecoin network should be sufficient to serve the chain API'
breadcrumb: 'Filecoin Chain'
confidenceLevel: 'Stable'
v0: |
  Lotus already contains two implementations of the above API:
  - A Lotus Full Node
  - The Lotus Gateway API, which wraps a JSON-RPC to the full node
v05:
  It's not clear whether the Lotus Gateway API works properly across an open network, and this may be important in a case where several parties in the retrieval market are not running a full lotus node.
v1:
  It's not clear we need anything else. It might make sense to extract the gateway API down to the above reduced package to remove any direct dependency on the lotus repo at some point.
---

<Header />

## Roadmap

<RoadMapPage />

## Preliminary API

This is extracted from the Lotus Gateway API, with dependencies reduced to what a validator node would serve as needed for a retrieval API

<<< @/docs/components/chain/chain.go
