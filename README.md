# ChefBook Backend Profile Service

The profile service owns aggregated profile read models. It composes data from downstream services instead of owning all profile facts directly.

## Responsibilities

- Return current and public profile read models.
- Combine account/auth state, public user data, and subscription state.
- Provide minimal profile information to recipe, encryption, and shopping-list flows.

## Main RPC Families

- `GetProfile`
- `GetProfilesMinInfo`

## Dependencies

- Calls `auth` for account/auth details.
- Calls `user` for public profile fields and avatar data.
- Calls `subscription` for subscription state.
- Does not own another service's data. IDs are logical cross-service references.

## Change Guidance

- Change this service when profile read composition changes.
- Change `auth` for credentials, sessions, OAuth, nickname, or deletion state behavior.
- Change `user` for public profile fields and avatar behavior.
- Change `subscription` for plan/source/expiration behavior.
