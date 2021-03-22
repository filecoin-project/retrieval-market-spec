---
title: Retrieval Market Spec
description: 'Specification For Components In A Retrieval Decentralized Data Market'
breadcrumb: Home
---

# Retrieval Market Specification

## Goals

This documention site describes a preliminary specification for the Filecoin Retrieval Market. The work here is synthesized from several previous efforts at building a Retrieval market. Previous efforts have gone deep into "how" to build a specific part of the Retrieval Market. This document aims to give the "40,000 foot" view of "what" the Retrieval Market actually is. The primary goal here is to describe, in technical detail, all the moving parts that will make up the Filecoin Retrieval Market and how they relate to each other. 

The secondary goal of this specification is to describe how to build the system as a whole, in the form of a roadmap towards various stages of development moving towards the broad goal of the Filecoin Retrieval Market serving as "CDN for the decentralized web". This roadmap provides recommendations for various milestones during the estimated "first year" of development, then offers suggestions for future development.

The final goal of this specification is to decentralized development of the Filecoin retrieval market, enabling it to be completed faster by allowing more work to happen in parallel, much of it outside Protocol Labs. By drawing clear boundaries around units of functionality in the Filecoin Retrieval Market, we can increase the ability of third parties, PL Research, or Filecoin dev grantees to go deep on building a particular part of the specification.

## Roadmap Overview

The roadmap in this document is broken into 3 milestones covering the estimated "first year" of development up to a "V1" milestone, and an unbounded "V2" milestone covering everything after that.

### **V0 Milestone**: Existing Or Finished In 3 Months

The V0 milestone describes simply what exists now, or will exist in 3 months. It contains primarily work that is either done or actively in development. For any new efforts not already started, the primary goal is just to fix extreme problem areas in the Filecoin Retrieval Market *AS IT STANDS*. Large areas of functionality in this specification will be missing or incomplete.

### **V0.5 Milestone**: Six Months

The V0.5 milestone is focused on lowering barriers to access, and making existing functionality work better. There will still be parts missing that make the Retrieval Market incomplete as true content delivery network, but we aim to provide a network that lots of folks can be part of and exchange content and coins easily in.

### **V1 Milestone**: 1 Year

The goal of the V1 1 year-milestone is "MVP for a web3 content delivery network". By V1, each major part of the Retrieval market is built, if not built in ideal or feature complete ways. The market should deliver subsecond retrieval times to off chain dapp users, unlocking significant product market fit for the web 3 stack for this segment

The V1 Milestone is defined as much by what's not in it as what is. Specifically, the V1 Milestone is comprised only of solutions for which there is a clear path to implementation. In other words, nothing that *needs research*. This includes:
- Any kind of on chain incentives or other complex cryptoeconomics
- Future content routing solutions (see PL Research 2021 Endeavor)
- Fancy geographic on demand content distribution and content marketplaces (see 3DM docs)

### **V2 Milestone**: Future Directions

With the above areas of research hopefully producing more well defined solutions in a year or so, V2 encompasses everything concrete we'd like to build next. It has no specific timeframe for implementation or delivery. The goal here is to

::: danger
This ends to non-technical introduction to this specification. From here on out, things are going to get engineer-ey
:::

## Achitectural Overview

The Retrieval Market is a decentralized data market made of of components, nodes, and actors performing high level roles.

A Component is a unit of common software functionality needed for the Filecoin Retrieval Market. Components always specify what other components, if any, they depend on. Components APIs must conform to a set of rules that allow the API to translate clearly between native function call, FFI, RPC, and in certain cases an HTTP REST API. This enables components to communicate with their dependents across language, process, or network boundaries.

A Node is a single libp2p host (or other type of network bounded address) running in the retrieval network, typically (though not always) in a single OS process. It can run one or more components.

An actor is an person or organization running one or more nodes in order to perform one or more high level roles neccesary to make the overall Retrieval Market function. Roles describe a typical composition of components that actors run together on a single node or across multiple nodes to perform a distinct higher level function in the network. The currently defined Roles in the Retrieval Market are Retrieval Client, Retrieval Provider, Content Provider, Payment Provider, and Marketplace Provider.

## How To Read This Specification

The rest of this specification describes in detail the components and role that make up the retrieval market. Each page will always contain at least the following distinct sections.

- **Name**: The name of the component or role
- **Description**: A brief prose description of what the component or role does
- **Technical Description**: For components, a API(s) for the components. For roles, a list of the components that make up the role.
- **Status**: A rating of how confident we are that the Technical Description is correct or close to where it will ultimately end up. The ratings follow the the Filecoin spec with one additional status to reflect the early state of this specification:
   - *Stable* - Unlikely to change in the foreseeable future.
   - *Reliable* - All content is correct. Important details are covered.	
   - *Draft/WIP* - All content is correct. Details are being worked on.
   - *Brainstorm* - No implementation or only a basic prototype exists, and the most of the content is very likely to change
   - *Incorrect* - Do not follow. Important things have changed.	
   - *Missing* - No work has been done yet.
- **Roadmap**: How we'd like to implement this component or role within the Roadmap milestones. Where possible, V0 will point to existing examples in the web3 ecosystem which either follow the specification closely or perform a broadly similar function.

# Compilable Spec

As a tool for enforcing correct typing, the repository containing this specification is also a go module, and all go code should compile correctly within this module (external dependencies intentionally do not include large existing monoliths like Lotus). This spec does not contain and will not contain implementations.