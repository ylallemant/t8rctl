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

For automation using environment, workload identity, or managed identity
credentials instead of a developer login, set `AZURE_TOKEN_CREDENTIALS=prod`
and configure the chosen identity. This setting controls the authentication
chain, not which subscriptions or clusters are listed.

## Tests

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
