#!/usr/bin/env bash
set -euo pipefail

##
# Only rely on these tools for things which are loaded by the partials which are
# also included for this language or type of project.
##

# @config-manager:start terraform-provider
go get -tool -modfile go.tools.mod github.com/hashicorp/terraform-mcp-server/cmd/terraform-mcp-server@latest
go get -tool -modfile go.tools.mod github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest
# @config-manager:end terraform-provider
