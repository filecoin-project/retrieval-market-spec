---
title: 'Understanding Components'
description: 'Components are basic building blocks of software functionality within the retrieval market. Components generally run in a single process on a single node in the network (though some of the more brainstorm-ey components represent potentially large groups of functionality that may break up later into smaller components). Components may be dependent on one or more other components. Components may expose a single API or multiple APIs that give access to only portions of the components functionality.'
---

## Filecoin Specificity

Significant portions of the retrieval market are in no way specific or dependent on the Filecoin Blockchain. The only components that are in specific to the Filecoin Blockchain at all are Wallet, Chain, and PaymentChannelManager. Even within that, there is very little functionality that could not be implemented with a different blockchain. At a later time, even these components might be generalized so that they supported arbitrary blockchains.

In the farther future however, Filecoin may implement chain based incentives for retrievial, which would make the market potentially operate more efficiently on the Filecoin blockchain.

## Dependency Specification

Every component depends on other components.

Each type of dependent component is categorized as either:

- *required*: The component always depends on a singled instance of the dependent component type
- *optional*: Optional or one instance depending on implementation (0..1 or optional depedncy)
- *multiple*: The component can interact with multiple instances of this component type.

When a dependent component type exposes multiple APIs, the component must specific whic API(s) it dependents on. When a dependent component type specifies only one API, the component is assumed to receive that API.

## API Specifications

For the time being all component APIs are specified as `golang` interfaces. However, the interfaces MUST follow an additional set of rules to enable relatively straightforward translation to FFI, RPC, and HTTP reset API.

1. Parameters passed to API functions must conform to roughly the IPLD data model. Concretely, this means no Golang interfaces, channels, or functions. The one exception to this is when the data conforms to IPLD data model, but its specific schema is not known, it can be referenced as a `go-ipld-prime` `ipld.Node`. `go-ipld-prime`'s `ipld.Link` may also be used in place of `cid.Cid`.

2. API functions that return once much return as most one data type conforming to the IPLD data model and one `golang` error.

3. API functions that stream results back must return a channel to a data type conforming to the IPLD data model and optionally, a `golang` error.

::: tip
These constraints are largely drawn from the [go-jsonrpc library](https://github.com/filecoin-project/go-jsonrpc) used by Lotus. This means translating them to a JSON-RPC API that operates across boundarings should be trivial. The only addition is support for `go-ipld-prime` nodes. Given that all `go-ipld-prime` nodes should be encodable to and decodable from DAG-JSON, this should not be hard to add.
:::

::: warning
I haven't though through the translation between these APIs and FFI... though I think it should be doable, with various work to convert structs to C-structs and encoding IPLD prime nodes as DAG-CBOR byte arrays.

Translation to HTTP REST is not immediately clear, and may not be doable in these constraints, though HTTP only based JSON-RPC should support all but streaming results, and HTTP+Websockects should support all of JSON-RPC
:::

::: danger TODO
Specified API permissions
:::

::: danger TODO
Think through authentication schemes across network boundaries
:::

::: danger TODO
In the future, we may want to translate this to something more nuetral like JSON-RPC. Alternatively, we may want to describe all of the types in IPLD schemas -- maybe it would make sense to formally define an IPLD schema named "IPLD-RPC" based on JSON-RPC.
:::

