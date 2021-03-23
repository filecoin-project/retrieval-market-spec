---
title: 'How To Read This Spec'
description:
  A guide to understanding the specification
breadcrumb: 'How To Read'
---

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

## Compilable Spec

As a tool for enforcing correct typing, the repository containing this specification is also a go module, and all go code should compile correctly within this module (external dependencies intentionally do not include large existing monoliths like Lotus). This spec does not contain and will not contain implementations.

::: tip NEXT STEPS
Now that you've familiarized yourself with how this specification works, you're ready to read about the Retrieval Market's low-level [Components](./components/) and high-level [Roles](./roles)
:::