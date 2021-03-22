---
title: 'Content Bid Index'
description: 
  The Content Bid Index tracks bids from content providers who want to store content. Providers list what they want to store and how they want to store it. Retrieval miners can search bid indexes to identify potential data they can host and get paid for.
confidenceLevel: 'Brainstorm'
breadcrumb: 'Content Bid Index'
v0: 
  There is no current content bid index for retrieval in Filecoin. Operation Dumbo Drop vaguely resembed a content bid index for Filecoin storage. 
v05: |
  As a starting point, a single global content bid index seems straightforward to build. PL could build one itself, or offer a dev grant to a third party.

  However, the "content distribution" side of the retrieval market needs further research. We are competing with web 2 CDN providers, so we need to make sure our offerings are comparable.
v2:
  Eventually, an efficient Retrieval Market would have several Content Bid Indexes, perhaps serving unique geographic regions. 
---

<Header />

## Dependencies

| Name | Kind | APIs |
| ---- | ---- | ---- |
| Content Distribution | optional | HostingClientAPI |

## Roadmap

<RoadMapPage />

## Preliminary API

This API is purely speculative

::: danger TODO
Continue to investigate existing CDN offerings
:::

<<< @/docs/components/contentbidindex/contentbidindex.go
