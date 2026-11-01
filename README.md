# Wii/DS Sake Server

OpenPak's Wii/DS Sake Server.

| | |
| --- | --- |
| Network | Wii, DS |
| Role | SAKE storage and file HTTP protocols |
| Depends on | No remote WFC service; the existing WFC database and shared `WFC_TOKEN_KEY_FILE` |
| Image | `ghcr.io/openpak/wfc-sake` |
| Licence | AGPL-3.0-only |

## What it serves

| Endpoint | What |
| --- | --- |
| Internal HTTP (`WFC_HTTP_ADDR`, default `:20207`) | Protocol and internal paths, which require `WFC_INTERNAL_KEY` |
| `GET /healthz` | Unauthenticated health check |

## Run

Configure the service from `example.env`, using the existing WFC database and
shared `WFC_TOKEN_KEY_FILE`. Build the image:

```sh
docker build -t wfc-sake .
```

Test commands, ports, environment variables and release detail are in
[docs/deploy.md](docs/deploy.md).

## Docs

- [Mario Kart Wii](docs/mario-kart-wii.md)
- [Changelog](CHANGELOG.md)
- [Protocol](docs/protocol.md)
- [Evidence](docs/evidence.md)
- [Deploy](docs/deploy.md)
- [Notes](docs/notes.md)

## Licence

AGPL-3.0-only. See [LICENSE](LICENSE).
