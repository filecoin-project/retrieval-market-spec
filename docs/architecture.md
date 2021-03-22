---
title: 'Achitectural Overview'
description:
  The Retrieval Market is a decentralized data market made of of components, nodes, and actors performing high level roles.
breadcrumb: 'Architecture'
---

## Achitectural Overview

The Retrieval Market is a decentralized data market made of of components, nodes, and actors performing high level roles.

A Component is a unit of common software functionality needed for the Filecoin Retrieval Market. Components always specify what other components, if any, they depend on. Components APIs must conform to a set of rules that allow the API to translate clearly between native function call, FFI, RPC, and in certain cases an HTTP REST API. This enables components to communicate with their dependents across language, process, or network boundaries.

A Node is a single libp2p host (or other type of network bounded address) running in the retrieval network, typically (though not always) in a single OS process. It can run one or more components.

An actor is an person or organization running one or more nodes in order to perform one or more high level roles neccesary to make the overall Retrieval Market function. Roles describe a typical composition of components that actors run together on a single node or across multiple nodes to perform a distinct higher level function in the network. The currently defined Roles in the Retrieval Market are Retrieval Client, Retrieval Provider, Content Provider, Payment Provider, and Marketplace Provider.
