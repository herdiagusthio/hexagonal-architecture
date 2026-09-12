# ADR 001: Adopting Hexagonal Architecture

**Date:** 2026-09-12
**Status:** Accepted

## Context
As a system grows, business logic often becomes entwined with database queries and HTTP handlers. This "Big Ball of Mud" makes it difficult to change databases, upgrade frameworks, or test logic without a full environment.

## Decision
We adopted the **Hexagonal Architecture** (Ports and Adapters) pattern.

## Trade-offs
- **Pros**:
  - **Independent of Frameworks**: The core logic is pure Go.
  - **Independent of UI**: We can swap an HTTP API for a CLI or a Message Queue without changing the business rules.
  - **Testability**: High-speed unit tests via mocks.
- **Cons**:
  - **Boilerplate**: More files and interfaces compared to a simple "MVC" approach.
  - **Complexity**: Requires a steeper learning curve for junior developers.

## Final Verdict
The long-term maintainability and reliability of the system far outweigh the initial boilerplate cost.
