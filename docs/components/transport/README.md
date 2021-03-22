---
title: 'Transport'
description: 'A transport is a mechanism for moving data over the wire with Data Transfer. Different transports provide different capabilities to data transfer. For example, some transports may have the ability to pause and resume requests, while others may not. Some may support incremental data events, while others may only support message events.'
confidenceLevel: Draft/WIP
breadcrumb: 'Transport'
v0: >
  Currently, a single Transport interface is imported for Data Transfer: Graphsync. The interface does not look exactly like the API described below, and is somewhat Graphsync specific.
v05: >
  Team Ignite is implementing a second transport for data transfer as part of the F3 Project: STP, or simple transport protocol. To support both protocols, the Transport interface will need to be made fully protocol nuetral. The API below is a suggested revision to the interface.
---

<Header />

## Dependencies

| Name | Kind | APIs |
| ---- | ---- | ---- |
| Data Transfer | single | EventAPI |

## Roadmap

<RoadMapPage />

## Preliminary API

This is an extraction and revision of data transfer's Transport interface

<<< @/docs/components/transport/transport.go
