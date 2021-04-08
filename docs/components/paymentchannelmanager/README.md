---
title: 'Payment Channel Manager'
description: 'The Payments Channel Manager provides APIs that enable payment for retrieval via payment channels.'
confidenceLevel: 'Draft/WIP'
breadcrumb: 'Payment Channel Manager'
v0: >
  A Payment Channel Channel manager is currently implemented in Lotus and used for retrievals. It works, but it does have one significant pain point: as written, it always creates a blocking chain message to deposit funds in a channel at the beginning of a retrieval. This prevents performant retrieval
v05:
  The API specified below is designed to remove some of the ambiguity in the current payment channel manager, and to efficiently unlock paying for lots of retrievals ahead of time, so that they can proceed in batchs without being blocked by the chain. This fulfills the original purpose of a payment channel.
v1:
  It's fairly ineffecient to create a payment channel between every party that wants to transact on a network. A common technique for managing payment channels in a large scale crypto currency network is a "hub and spoke pattern". In this scenario, two parties who do not yet have a direct payment channel can go through an intermediary they've both already established a channel with. The mechanism to do this called is "Hash-Locked-Time-Codes", and the Filecoin blockchain has the machinery to do this. This is possibly something we need to implement to get payment channels to work at scale, even for an "MVP" retrieval market CDN. Both Bitcoin and Ethereum offer networks that have this functionality
v2:
  There are several further innovations in the world of payment channel mechanisms (https://finalitylabs.io/static/media/SetPaymentChannels.8a29d449.pdf, https://eprint.iacr.org/2017/635.pdf). Third parties looking to innovate in this space can offer custom implementations of the payment channel manager as a way of becoming service providers for fast payments
---

<Header />

## What is a payment channel?

A payment channel is a direct payment connection between two participants in a cryptocurrency network. Through on-chain opening and closing transactions, payment channels leverage the security of the underlying blockchain to provide trustless payments between two parties without needing to submit each intermediate payment on chain. 

Because we assume that retrievals need to happen fast, we further assume that we cannot block any step of retrieval for an on chain transaction. This must happen later. As such, we need some kind of payment channel system that operates as an L2 solution to the underlying L1 blockchain.

## What does the payment channel manager do?

The payment channel manager has two responsibilies: it communicates with the block chain to manage the on-chain component of payment channels, and it maintains state about payment channels that has not yet been submitted on on chain (such as outstanding vouchers).

## Dependencies

| Name | Kind | APIs |
| ---- | ---- | ---- |
| Chain | required | default |
| Wallet | required | default |

Note: a chain and a wallet are sufficient as dependencies to a payment channel manager, but there are a couple steps involved to hook up the current Lotus payment channel manager to those APIs. A good example the intermediate code required can be found here:
https://github.com/application-research/estuary/blob/master/filclient/msgpusher.go
and here:
https://github.com/filecoin-project/lotus/blob/master/chain/stmgr/rpc/rpcstatemanager.go

## Roadmap

<RoadMapPage />

## Preliminary API

The API below is extracted and modified from Lotus's Payment Channel Manager API
Notable changes include:

- removing the highly complicated and opaque method PaychGet (essentially a "get, create, and add" operation that is highly unpredicatble) and replacing it
with methods that perform simpler clearer operations
- combining two seperate methods we offer for getting status on payment channels into one
- adding a PaychReserveFunds and PaychReleaseFunds that mirror the methods on the Lotus FundManager -- by allowing other processes to essentially "lock up" outstanding funds in a payment channel for an ongoing transaction, we have a sense of how much is actually available to spend. Not having this information has meant we always assume we need a deposit -- an create a blocking chain message to create a channel or add funds. This way, we'll know if there are free funds for a retrieval, and avoid a chain message if we do
- Removing the concept of a "wait sentinel" as a thin abstraction on a chain message. The waiting methods just wrapped the Chain StateWaitMsg call and seem redundant.

<<< @/docs/components/paymentchannelmanager/paymentchannelmanager.go



