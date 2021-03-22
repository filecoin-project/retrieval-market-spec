---
title: 'Wallet'
description: 'A Wallet contains one or more private keys to addresses in a crypto currency network. It is used to verify origination for chain messages, off chain vouchers, and other important verified pieces of data'
breadcrumb: 'Wallet'
confidenceLevel: 'Stable'
v0:
  Lotus implements a LocalWallet based on a local Keystore, and a remote wallet that works on a JSON-RPC endpoint. It's not clear whether the remote wallet contains well fleshed out authentication controls
v1:
  We may need to be able to delegate remote signing on a per-address basis. Lotus 
---

<Header />

## Roadmap

<RoadMapPage />

## Preliminary API

This is an extraction of the Lotus wallet API

<<< @/docs/components/wallet/wallet.go
