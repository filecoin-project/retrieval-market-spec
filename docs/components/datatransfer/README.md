---
title: 'Data Transfer'
description: 'Data Transfer coordinates transfers across the wire, negotiating transport mechanism, maintaining state about a transfer, and coordinating with an Exchange protocol to validate requests and manage payments.'
confidenceLevel: Draft/WIP
breadcrumb: 'Data Transfer'
v0: >
  Data Transfer has been implemented, but only supports a single Transport mechanism. Moreover, the current implementation has architectural issues, as documented in [go-data-transfer#158](https://github.com/filecoin-project/go-data-transfer/issues/158) 
v05: >
  Team Ignite is implementing a second transport for data transfer as part of the F3 Project. As part of this, Team Ignite will need to improve Data Transfer to support multiple protocols, and determine how Data Transfer will choose between protocols.
---

<Header />

## Dependencies

| Name | Kind | APIs |
| ---- | ---- | ---- |
| Transport | multiple | default |
| Exchange | multiple | ExchangeValidatorAPI, ExchangeCheckpointValidatorAPI |

## Roadmap

<RoadMapPage />

## Preliminary API

This is an extraction of the existing API from [go-data-transfer](https://github.com/filecoin-project/go-data-transfer) with only one minor revision -- the events method returns a channel instead of accepting a callback (per API spec rules). Missing from this API proposal is any parameters around choosing a transport. I've also removed all of the configuration methods setup voucher handlers on the theory these details that could differ per-implementation 

<<< @/docs/components/datatransfer/datatransfer.go
