# Changelog

## v0.4.0 - 2026-11-05

- Link-first (prds#139), behind `WFC_LINK_FIRST=1`: storage requests and Mario Kart Wii ghost uploads of an unlinked console are refused (`LoginTicketInvalid`, missing-parameter result). The leaderboards adapter reads the account from the local link cache first. Vendors wfc-common v0.8.0.

## v0.3.0 - 2026-11-04

- Mario Kart Wii times go to the neutral leaderboards core (wfc-common v0.4.0 `mkwboard`): each accepted ghost upload is also submitted as a time on the course board (worldwide and the player's region) for the account linked to the console. Needs `WFC_LEADERBOARDS_ADDR`/`_KEY` and `WFC_CORE_ADDRESS`/`_KEY`; unset leaves the service as before.
- At start and every ten minutes the times in `mario_kart_wii_sake` that the core lacks are imported for linked consoles (idempotent, the table is never changed). `mario_kart_wii_sake` stays the ghost-file store.

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
