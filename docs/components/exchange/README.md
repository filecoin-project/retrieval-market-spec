---
title: 'Exchange'
description:
  Exchange is a protocol for performing an optimistic fair exchange of payment for data. The protocol coordinates sending of data (with Data Transfer) and payments (via the Payment Channel Manager) during a paid retrieval, and attempts to insure that both parties fulfill their obligations completely. (Fair exchange is generally a hard problem and there may be no perfect solutions)
confidenceLevel: Draft/WIP
breadcrumb: 'Exchange'
v0: |
  Retrieval Protocol V1 is the current implementation of Exchange in Filecoin. It operates through a system of incremental payments, with the data provider sending small chunks then asking for payment before sending more. 
v05: >
  The proposed API below represents a slightly more semantic and clean design than the current implementation. It may make sense ot conform to the proposed API in preparation for adding other exchange protocols in the future.
v1:
  In a more traditional CDN context, a content provider is often pays for bandwidth and resources on a servers so that a client is able to download data quickly. The client themselves often does not pay anything for this data. To support this kind of usage, we will likely have to augment the Exchange Protocol to support this kind of 3-way transaction. A client might in this scenario instead send a proof to the content provider that they have received the data causing the content provider to unlock payment to the miner.
v2:
  Fair exchange is a hard problem and out existing retrieval protocol is very much an "MVP". While it may be sufficient for some time, improved protocols might unlock superior performance, particularly if they become tied to crtypographic incentires.
---

<Header />

## Dependencies

| Name | Kind | APIs |
| ---- | ---- | ---- |
| Data Transfer | required | ManagerAPI |
| Payment Channel Manager | required | default |

## Roadmap

<RoadMapPage />

## Preliminary API

This API is primarily based on the Filecoin Retrieval Protocol implementation in [go-fil-markets](https://github.com/filecoin-project/go-fil-markets/tree/master/retrievalmarket)

<<< @/docs/components/exchange/exchange.go
