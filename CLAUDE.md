# CLAUDE.md

Guidance for Claude Code when working in this repository.

## Overview

`az-pim-cli` is a single-binary Go CLI for listing, activating and deactivating
Azure Entra ID **Privileged Identity Management (PIM)** role assignments. It
covers three assignment families:

- **Azure resources** (subscriptions, resource groups, …) via the ARM PIM API
- **Entra groups** (`aadGroups`) via the Azure RBAC governance API
- **Entra roles** (`aadroles`) via the same governance API

It also ships an **MCP (Model Context Protocol) server** (`az-pim-cli mcp`) that
exposes the same operations as tools over stdio for AI assistants.

Module: `github.com/neverprepared/az-pim-cli` — Go 1.24.

## Architecture

```
main.go            -> cmd.Execute()
cmd/               cobra command tree, flag/config wiring (viper)
  root.go          root cmd, global flags, config + env binding, logger init
  list.go          list {resource|group|role}  (eligible assignments)
  list_active.go   list active {resource|group|role}
  activate.go      activate {resource|group|role}
  deactivate.go    deactivate {resource|group|role}
  token.go         token {resource|governance}
  mcp.go           mcp (stdio MCP server)
  version.go       version
pkg/pim/           Azure PIM API layer
  client.go        Client interface + AzureClient impl (azidentity AzureCLICredential)
  const.go         base URLs per cloud, role types, defaults
  models.go        request/response types
  utils.go         request builders, status predicates, wait/poll helpers
  test_data.go     fixtures used by tests
pkg/utils/         table/JSON output + assignment lookup by name/prefix/role
pkg/common/        slog setup, ANSI color helpers, Error type
pkg/mcp/server.go  MCP tool registration + handlers (wraps pkg/pim + pkg/utils)
```

Key design points:

- **Auth is delegated to the Azure CLI.** `AzureClient.GetAccessToken` uses
  `azidentity.NewAzureCLICredential`, so `az login` must have been run. There is
  no credential storage in this repo.
- **Two APIs, two tokens.** Resource commands fetch an ARM token automatically.
  Group/role (governance) commands require an explicit `--token` / `-t` for the
  `api.azrbac.mspim.azure.com` scope; `az-pim-cli token governance` prints one.
- **Multi-cloud** via the global `--cloud` flag (`global`, `usgov`, `china`),
  resolved through `pim.ARM_BASE_URLS` in `PersistentPreRun`. Only the ARM base
  URL varies; the governance API base URL is a single constant.
- **`pim.Client` is an interface**, and every exported `pim.X(..., c Client)`
  helper takes it — that is the seam tests use (see `pkg/pim/client_test.go`).
- Fatal errors call `slog.Error` then `os.Exit(1)` inside `pkg/pim`; commands do
  not propagate errors upward.

## Commands

```bash
go build ./...          # build (binary: go build -o az-pim-cli .)
go test ./...           # unit tests (go test -v ./... in CI)
go vet ./...
pre-commit run --all-files   # lint: golangci-lint + whitespace/yaml hooks
go run . <subcommand>   # run locally
```

Lint config lives in `golangci.yaml`; hook versions in `.pre-commit-config.yaml`
(golangci-lint v1.64.5). CI runs lint, build (5 GOOS/GOARCH combos), tests,
conventional-commit checks and Semgrep — see `.github/workflows/`.

## CLI surface

```
az-pim-cli list       {resource|group|role}
az-pim-cli list active {resource|group|role}
az-pim-cli activate   {resource|group|role}
az-pim-cli deactivate {resource|group|role}
az-pim-cli token      {resource|governance}
az-pim-cli mcp
az-pim-cli version
```

Global flags: `--debug`, `--json/-j` (logs move to stderr), `--config/-c`,
`--cloud`.
Activate flags: `--name/-n` (repeatable), `--prefix/-p`, `--role/-r`,
`--duration/-d` (default 480 min), `--start-date`, `--start-time/-s`,
`--reason`, `--ticket-system`, `--ticket-number/-T`, `--dry-run`,
`--validate-only/-v`, `--wait`, `--timeout` (default 300 s), `--all`
(resource/group/role subcommands only; mutually exclusive with name/prefix).
Group/role subcommands additionally require `--token/-t`.

Config: `$HOME/.az-pim-cli.yaml` (or `--config`), env vars prefixed `PIM_`
(e.g. `PIM_TOKEN`, `PIM_DURATION`). Flags win over config.

## MCP tools

`az-pim-cli mcp` serves JSON-RPC over stdio and registers 12 tools in
`pkg/mcp/server.go`:

`list_eligible_resources`, `list_eligible_groups`, `list_eligible_entra_roles`,
`list_active_resources`, `list_active_groups`, `list_active_entra_roles`,
`activate_resource`, `activate_group`, `activate_entra_role`,
`deactivate_resource`, `deactivate_group`, `deactivate_entra_role`.

Group/role tools require a `token` argument; all accept `cloud`. Activate tools
accept `name`/`prefix`/`role`/`duration`/`reason`/`ticket_system`/
`ticket_number`/`start_date`/`start_time`.

## Conventions

- **Conventional Commits are enforced** in CI and by a commit-msg hook. Allowed
  types: `feat, fix, refactor, chore, test, ci, build, lint, docs`.
- Releases are automated: release-please opens the release PR on `main`,
  goreleaser publishes on tag. Do not hand-edit `CHANGELOG.md` or
  `.release-please-manifest.json`.
- New CLI flags must be added to the `bindFlags(...)` list in
  `cmd/root.go:initConfig` to be readable from config/env.
- New subcommands follow the existing shape: a `cobra.Command` var plus an
  `init()` that wires it into its parent.
- Human output goes through `pkg/utils` printers; `--json` output through
  `utils.PrintJSON`. Keep both paths in sync when adding fields.
- Tests use `testify` and the fixtures in `pkg/pim/test_data.go`; mock the
  `pim.Client` interface rather than hitting the network.
- Never log or print access tokens outside `az-pim-cli token`; `--debug` output
  is already documented as sensitive.
