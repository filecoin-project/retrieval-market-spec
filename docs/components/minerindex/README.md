---
title: 'Miner Index'
description:
  Miner Index is a reputational index for miners in the market. It allows prospective retrieval clients and content providers to search for the best miners from which to store and retrieve their data.
confidenceLevel: Brainstorm
breadcrumb: 'Miner Index'
v0: |
  Not having a reputational index has made retrieval in the current Filecoin network largely unusable. During the next three months, Protocol Labs will likely build a prototype reputation index in the form of a retrieval deal bot. The bot will make deals with storage miners and begin to asses the reliability of different miners. 
v05: >
  Once we have a prototype index, we can make it available as an API as well as a web portal, and begin to integrate it into retrieval clients.
v2:
  A single point of failure index is insufficent to provide reputation for the network for the long term. Over time, third parties can create competetive indexes.
---

<Header />

## Dependencies

| Name | Kind | APIs |
| ---- | ---- | ---- |
| Chain | optional | default |
| Exchange | multiple | ExchangeProviderAPI |

## Roadmap

<RoadMapPage />

## Preliminary API

This API is purely speculative

::: danger TODO
Sync up with current designs for retrieval deal bot
:::

<<< @/docs/components/minerindex/minerindex.go
