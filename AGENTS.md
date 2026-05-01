# Profile Service Agents Guide

This service aggregates profile data from other services.

## Scope

- profile read models and aggregation logic
- composition of data coming from auth, user, subscription, or related services

## Working Rules

- Keep this service focused on aggregation and profile-facing composition.
- Avoid duplicating source-of-truth logic that belongs to upstream services.
- When changing aggregated fields, verify every upstream dependency and downstream consumer.
