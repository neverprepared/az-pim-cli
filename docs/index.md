# az-pim-cli

CLI utility for managing Azure Privileged Identity Management (PIM) role assignments from the command line.

## Install via Homebrew

```bash
brew tap mindmorass/az-pim-cli https://github.com/mindmorass/az-pim-cli
brew install az-pim-cli
```

To install the latest development build directly from `main`:

```bash
brew install --HEAD az-pim-cli
```

## Prerequisites

Authentication uses `az-cli`. Log in once before using this tool:

```bash
az login
```

## Token for Entra Roles and Groups

Azure resource roles are handled automatically. For **Entra groups** and **Entra roles**, you need a scoped access token:

```bash
az account get-access-token \
  --resource https://api.azrbac.mspim.azure.com \
  --query accessToken -o tsv
```

Export it for the session:

```bash
export PIM_TOKEN=$(az account get-access-token \
  --resource https://api.azrbac.mspim.azure.com \
  --query accessToken -o tsv)
```

## Usage Examples

### List eligible assignments

```bash
# Azure resource roles
az-pim-cli list resources

# Entra groups
az-pim-cli list groups

# Entra roles
az-pim-cli list roles
```

### Activate a role

```bash
# Activate by subscription prefix
az-pim-cli activate resource --prefix S100

# Activate a specific role
az-pim-cli activate resource --prefix S100 --role Owner

# Activate with ticket number
az-pim-cli activate resource --name S100-Example-Subscription --role Owner \
  --ticket-system Jira --ticket-number T-1337

# Activate an Entra group
az-pim-cli activate group --name my-entra-id-group --duration 60
```

## Configuration

Create `~/.az-pim-cli.yaml` with defaults:

```yaml
token: eyJ0...        # Entra groups/roles token
reason: my-reason
ticketSystem: Jira
ticketNumber: T-1337
duration: 480         # minutes
cloud: global         # global | usgov | china
```

Or use environment variables prefixed with `PIM_`:

```bash
export PIM_TOKEN=eyJ0...
export PIM_CLOUD=global
```

## Links

- [Full README](https://github.com/mindmorass/az-pim-cli/blob/main/README.md)
- [Upstream project](https://github.com/netr0m/az-pim-cli)
- [Releases](https://github.com/mindmorass/az-pim-cli/releases)
