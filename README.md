# t8rctl

[Documentation can be found here](./docs/content.md)

## Install

```sh
curl -fsSL https://github.com/ylallemant/t8rctl/raw/0.0.11/install.sh | bash
```

## Azure authentication

By default, t8rctl uses developer credentials: Azure CLI first, then Azure
Developer CLI. Log in once before running Azure commands:

```bash
az login
go run ./cmd/t8rctl/main.go cluster ps
```

No environment variable is required for local use. When `AZURE_TOKEN_CREDENTIALS`
is unset, t8rctl sets it to `dev` within its own process. Explicit values are
preserved and validated by the Azure SDK.

To fetch fresh cluster and subscription data instead of using cached results:

```bash
go run ./cmd/t8rctl/main.go cluster ps --disable-cache
```

The short form is `-c`. Fresh results still refresh the cache for subsequent
commands.

For automation using environment, workload identity, or managed identity
credentials instead of a developer login, set `AZURE_TOKEN_CREDENTIALS=prod`
and configure the chosen identity. This setting controls the authentication
chain, not which subscriptions or clusters are listed.

## Tests

Development and releases use Go 1.27.2, as specified in `go.mod`. Go can
download the required toolchain automatically when `GOTOOLCHAIN=auto`.
Dependencies are vendored; after changing them, run `go mod tidy` and
`go mod vendor`. The release workflow runs tests before publishing the existing
Linux, Windows, and macOS binaries, using the Go version from `go.mod`.

### Run

```bash
go test -cover ./...
```

### Coverage

Enable visual coverage feedback in `vscode` :

```json
{
    "go.coverOnSave": true
}
```
