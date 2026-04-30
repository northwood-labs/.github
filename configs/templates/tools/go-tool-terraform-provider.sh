#!/usr/bin/env bash
set -euo pipefail

# @config-manager:start terraform-provider
go get -tool -modfile go.tools.mod github.com/hashicorp/terraform-mcp-server/cmd/terraform-mcp-server@latest
go get -tool -modfile go.tools.mod github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest
# @config-manager:end terraform-provider
