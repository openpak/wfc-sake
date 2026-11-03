# Changelog

## v0.2.2 - 2026-11-03

- Internal `/internal/` polls are counted and summarised once a minute instead of logged per call (wfc-common v0.3.1).

## v0.2.1 - 2026-11-03

- Every HTTP request now logs one line to stderr, `http <method> <path> -> <status> <ms>ms` (path only, never the query, headers or bodies; `/healthz` skipped), through wfc-common v0.3.0 `httplog`.

## v0.2.0

- Requires wfc-common v0.2.0: every accepted Mario Kart Wii ghost upload is also archived with its
  RKG header metadata (see [docs/mario-kart-wii.md](docs/mario-kart-wii.md)); the archive schema is
  created and existing ghosts backfilled at start.
- Unknown query parameters on the ghost download and upload are logged (names only).
- Game-facing behaviour is unchanged.

## v0.1.0

First split release.
