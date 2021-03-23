---
title: 'Content Distribution'
description: 'The Content Distribution component allows retrieval miners and and content providers to negotiate content distribution. Miners offer hosting to providers, and providers send data then pay when its retrieved by clients'
confidenceLevel: 'Brainstorm'
breadcrumb: 'Content Distribution'
v0:
  The storage API is arguably somewaht similar to this distribution API, except the providers make the offers and money is paid for retrieval, not storage.
v5:
  We need to research design for this system further, and it may be worth offering a dev grant to someone to deep dive into content distribution.
v1:
  It may make sense for hosting and distribution to operate more like an automatic ledger -- content provider asks get automaticallys matched with miner pararms and content providers may not be allowed to reject deals that fall within the parameters of the ask.
v2:
  Ultimately this is a wide open research problem and hopefully during 2021 ResNetLab can look at potential apporaches.
---

<Header />

## Dependencies

| Name | Kind | APIs |
| ---- | ---- | ---- |
| Data Transfer | required | ManagerAPI |

## Roadmap

<RoadMapPage />

## Preliminary API

This API is purely speculative

::: warning 
This is based of the idea that retrieval miners offer hosting deals to content providers based off of the information in the content bid index. It's not clear if that is the right direction-- content providers may be behind a VPN/NAT, meaning they may be undialable.
:::

<<< @/docs/components/contentdistribution/contentdistribution.go

::: tip
You've made it to the end of the Components section. Now maybe you'd like to learn how all these components come together to perform higher level [Roles](../../roles/)
:::