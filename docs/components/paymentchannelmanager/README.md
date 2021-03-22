---
title: 'Payments'
description: 'The Payments Component provides APIs that enable payment for retrieval via payment channels.'
breadcrumb: ''
---
# {{ $frontmatter.title }}

{{ $frontmatter.description }}

## Payment Channel

A payment channel is a direct payment connection between two actors in a cryptocurrency network. Through on-chain opening and closing transactions, payment channels leverage the security of the underlying blockchain to provide trustless payments between two parties without needing to submit each intermediate payment on chain. 

Because we assume that retrievals need to happen fast, we further assume that we cannot block any step of retrieval for an on chain transaction. This must happen later. As such, we need some kind of payment channel system that operates as an L2 solution to the underlying L1 blockchain.

# Example/Preliminary API

```golang
// for chain level operations, consumers may want to record a chain operation
// has been submitted in a persistent manner that survives a restart
type CompletionToken ipld.Node
type Address ipld.Node
type TokenAmount ipld.Node
type PaymentChannel interface{
  ipld.Node
  From() Address
  To() Address
  Address() Address
}

type Voucher interface{
  ipld.Node
  PaymentChannel() PaymentChannel
}

type Provider interface {
  CreatePaymentChannel(from, to Address, amtReserved TokenAmount) (PaymentChannel, CompletionToken, error) 
  LookupPaymentChanne(from, to Address) (PaymentChannel, error)
  AddToOnChainBalance(paymentChannel PaymentChannel) (CompletionToken, error) 
  OnChainState(ctx context.Context, paymentChannel PaymentChannel)
  CreateOffchainVoucher(ctx context.Context, paymentChannel PaymentChannel, amount TokenAmount) (Voucher, error)
  AddPaymentVoucherToOffChainState(ctx context.Context, paymentVoucher )
  WaitForCompletion(ctx context.Context, CompletionToken) error
}
```
Examples: early bitcoin lightning network (https://tik-old.ee.ethz.ch/file/716b955c130e6c703fac336ea17b1670/duplex-micropayment-channels.pdf)

## Virtual Payment Channel

A virtual payment channel is a second level abstraction on top of on chain payment channels that allows two parties with no on chain payment channel to transact through an intermediary, either through hash locked payments or in some cases through direct exchange of value until the moment the virtual channel is closed.

Examples: PeRun (https://eprint.iacr.org/2017/635.pdf), Set-Payment Channel (https://finalitylabs.io/static/media/SetPaymentChannels.8a29d449.pdf)

## Payments API

Payments are a system that provides the ability to open payment channels, create and validate vouchers, enter a settling period, and close the channel. Whether the channel created is ultimately an on-chain ledger channel or a virtual channel is an implementation detail of the payment systems. 

::: warning TODO
Is it really possible to abstract whether a payment channel is a legdger channel or virtual? Not clear if research in this field is clear. If not, how would we support something like hash locked payment channels where every single transaction must go through an intermediary?
:::

::: warning TODO
For now, the payments system is hard coded to Filecoin payment channels and the Filecoin network. A future improvement might be to support other crytocurrency networks
:::

