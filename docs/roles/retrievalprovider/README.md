---
title: 'Retrieval Provider/Miner'
description: 'Retrieval Providers/Miners host content and provide it for retrieval. Retrieval Providers advertise their content through Content Marketplaces, accept bids to host content from Content Providers, and serve content to Retieval Clients'
confidenceLevel: 'Draft/WIP'
breadcrumb: 'Retrieval Provider'
v0: 
  Currently, the only way to be a Filecoin Retrieval Miner is to be a Filecoin Storage Miner, an incredibly high bar to entry that among other things doesn't incentivize doing the job of retrieval mining well (given storage rewards produce greater profit). 
v1: |
  Our immediate MVP goal for the retrieval miner is to enable miners to serve retrievals without running Lotus Storage Mining software. Similar to the client, go-ipfs might produce a smaller, more functional retrieval provider. To do this we could:
  - Put go-data-transfer in go-ipfs
  - Have the miner  advertise its content to index-style Content Routing for Filecoin
  - Accept incoming requests and call out to an Exchange Provider to manage checkpointing and verification of payment
  - Run the payment channel manager locally in a Lotus Lite node or potentially talk to a remote Payment Provider, if such a provider exists.

  **Alterative 1**: As with the Retrieval Client, Myel's work to enable retrieval through a plugin that runs Data Transfer and a custom Exchange/Content Routing protocol seems like an avenue worth exploring further

v2: |
  The biggest long term risk for Retrieval Miners is making sure it's profitable to do only retrieval mining. This means we may need to research cryptoeconomic incentives for retrieval mining over the long term.
---

<Header />

## Retrieval Provider Components

A retrieval provider must run the following components themselves:

1. Transport(s)
2. Data Transfer
3. Exchange Protocol(s) (Provider)
4. Wallet
5. Content Distribution (Provider - if they want to get paid to host)

A retrieval provider can run the following components locally, but will often instead interact with them remotely through a different provider:
1. Content Routing
2. Payment Channel Manager
3. Chain
4. Content Bid Index

While a retrieval provider is listed on a Miner Index, they generally don't act directly with the index.

## Roadmap

<RoadMapPage />
