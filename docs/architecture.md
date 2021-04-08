---
title: 'Achitectural Overview'
description:
  The Retrieval Market is a decentralized data market made of of components, nodes, and participants performing high level roles.
breadcrumb: 'Architecture'
---

## Achitectural Overview

The Retrieval Market is a decentralized data market made of of components, nodes, and participants performing high level roles.

A Component is a unit of common software functionality needed for the Filecoin Retrieval Market. Components always specify what other components, if any, they depend on. Components APIs must conform to a set of rules that allow the API to translate clearly between native function call, FFI, RPC, and in certain cases an HTTP REST API. This enables components to communicate with their dependents across language, process, or network boundaries.

A Node is a single libp2p host (or other type of network bounded address) running in the retrieval network, typically (though not always) in a single OS process. It can run one or more components.

An participant is an person or organization running one or more nodes in order to perform one or more high level roles neccesary to make the overall Retrieval Market function. Roles describe a typical composition of components that participants run together on a single node or across multiple nodes to perform a distinct higher level function in the network. The currently defined Roles in the Retrieval Market are Retrieval Client, Retrieval Provider, Content Provider, Payment Provider, and Marketplace Provider.

## Architecture Diagram

<br />
<br />

@startuml

node "Retrieval Client" as RetrievalClient {
  [Exchange] as CExchange
  [CExchange] -- ExchangeClientAPI 
  () ManagerAPI as CManagerAPI
  [CExchange] ----( CManagerAPI
  [Data Transfer] as CDataTransfer
  CManagerAPI -- [CDataTransfer]
  () EventAPI as CEventAPI
  () TransportAPI as CTransportAPI
  [Transport] as CTransport
  [CDataTransfer] -- CEventAPI
  [CDataTransfer] --( CTransportAPI
  CEventAPI )-- [CTransport]
  CTransportAPI -- [CTransport]
  () "WalletAPI" as CWalletAPI
  [Wallet] as CWallet
  CWalletAPI -- [CWallet]
}
node "Retrieval Provider" as RetrievalProvider {
  () "WalletAPI" as RPWalletAPI
  [Wallet] as RPWallet
  RPWalletAPI -- [RPWallet]
  [Exchange] as RPExchange
  [RPExchange] -- ExchangeProviderAPI 
  ExchangeProviderAPI .left. ExchangeClientAPI
  () ExhangeValidatorAPI as RPExhangeValidatorAPI
  () ExchangeCheckpointValidatorAPI as RPExchangeCheckpointValidatorAPI
  [RPExchange] -- RPExhangeValidatorAPI
  [RPExchange] -- RPExchangeCheckpointValidatorAPI
  () ManagerAPI as RPManagerAPI
  [RPExchange] ----( RPManagerAPI
  [Content Distribution] as RPContentDistribution
  [RPContentDistribution] -up- HostingProviderAPI
  [RPContentDistribution] -left-( RPManagerAPI
  [Data Transfer] as RPDataTransfer
  [RPDataTransfer] -left-  [CDataTransfer]
  RPManagerAPI -- [RPDataTransfer]
  RPExhangeValidatorAPI )-- [RPDataTransfer]
  RPExchangeCheckpointValidatorAPI )-- [RPDataTransfer]
  () EventAPI as RPEventAPI
  () TransportAPI as RPTransportAPI
  [Transport] as RPTransport
  [RPTransport] -left- [CTransport]
  [RPDataTransfer] -- RPEventAPI
  [RPDataTransfer] --( RPTransportAPI
  RPEventAPI )-- [RPTransport]
  RPTransportAPI -- [RPTransport]
}
node "Content Provider" as ContentProvider {
  [Exchange] as CPExchange
  () ManagerAPI as CPManagerAPI
  [CPExchange] ----( CPManagerAPI
  [Content Distribution] as CPContentDistribution
  [CPContentDistribution] -- HostingClientAPI
  HostingClientAPI .left. HostingProviderAPI
  [CPContentDistribution] ---- HostingValidatorAPI
  [CPContentDistribution] ----( CPManagerAPI
  [Data Transfer] as CPDataTransfer
  [CPDataTransfer] - [RPDataTransfer]
  CPManagerAPI -- [CPDataTransfer]
  HostingValidatorAPI )-- [CPDataTransfer]
  () EventAPI as CPEventAPI
  () TransportAPI as CPTransportAPI
  [Transport] as CPTransport
  [CPTransport] -right- [RPTransport]
  [CPDataTransfer] -- CPEventAPI
  [CPDataTransfer] --( CPTransportAPI
  CPEventAPI )-- [CPTransport]
  CPTransportAPI -- [CPTransport]
  () "WalletAPI" as CPWalletAPI
  [Wallet] as CPWallet
  CPWalletAPI -- [CPWallet]
}
node "Payment Provider" {
  () "WalletAPI" as PPWalletAPI
  [Wallet] as PPWallet
  [CExchange] -up-( PaymentChannelManagerAPI
  [RPExchange] -up-( PaymentChannelManagerAPI
  [CPExchange] -up-( PaymentChannelManagerAPI
  [Payment Channel Manager] --( ChainAPI
  [Payment Channel Manager] -( CWalletAPI
  [Payment Channel Manager] -( RPWalletAPI
  [Payment Channel Manager] -( CPWalletAPI
  [Payment Channel Manager] --( PPWalletAPI
  PPWalletAPI -- [PPWallet]
  ChainAPI -- [Chain]
  PaymentChannelManagerAPI -up- [Payment Channel Manager]
}
node "MarketPlaceProvider" {
  RetrievalClient -------down-( MinerIndexAPI
  RetrievalProvider -------down-( ContentBidIndexAPI
  ContentProvider -------down-( ContentBidIndexAPI
  RetrievalProvider -------down-( ContentRoutingAPI
  RetrievalClient -------down-( ContentRoutingAPI
  MinerIndexAPI -- [Miner Index]
  ChainAPI .....down. [Miner Index]
  ContentBidIndexAPI -- [Content Bid Index]
  ContentRoutingAPI -- [Content Routing (Indexed)]
}

() ContentRoutingAPI as DContentRoutingAPI
RetrievalProvider -------down-( DContentRoutingAPI
RetrievalClient -------down-( DContentRoutingAPI
DContentRoutingAPI -- [Content Routing (Distributed - DHT)]
@enduml