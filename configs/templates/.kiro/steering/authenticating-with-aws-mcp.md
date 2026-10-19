---
# @config-manager:start add-code-comments-frontmatter
inclusion: auto
name: authenticating-with-aws-mcp
description: Instructions for authenticating with AWS in order to leverage AWS MCP and Knowledge servers.
# @config-manager:end add-code-comments-frontmatter
---

<!-- @config-manager:start add-code-comments -->
# Authenticating with AWS MCP

The Agent Toolkit for AWS gives AI coding agents the tools, knowledge, and guardrails to build, deploy, and manage applications on AWS. It works with existing coding agents (Kiro, Claude Code, Cursor, Codex) — no need to switch tools or learn a new workflow. It gives agents secure access to AWS services, current documentation, and step-by-step guidance, which helps avoid common failures: choosing the wrong service, misconfiguring resources, or retrying operations on newer services the agent was not trained on.

* [What is the Agent Toolkit for AWS?](https://docs.aws.amazon.com/agent-toolkit/latest/userguide/what-is-agent-toolkit.html)
* [Getting started](https://docs.aws.amazon.com/agent-toolkit/latest/userguide/quick-start.html)

**IMPORTANT:** This document is only relevant when AWS MCP calls fail due to authentication. If you can call the AWS MCP tools successfully, you are already authenticated — disregard the rest of this document.

Follow the steps below in order only when authentication is needed.

## Step 1: Find the AWS profile from the MCP configuration

Read both `~/.kiro/settings/mcp.json` and `<workspace>/.kiro/settings/mcp.json`. Locate the MCP server whose args contain `https://aws-mcp.us-east-1.api.aws/mcp` — this is the AWS Agent Toolkit MCP server.

In that server's configuration, look for an `AWS_PROFILE` value in either location:

* An arg of the form `AWS_PROFILE=<value>` (e.g., `AWS_PROFILE=sandbox`).
* An `env` entry such as `"env": { "AWS_PROFILE": "sandbox" }`.

If found, use that profile value in Step 3 and skip Step 2.

## Step 2: List profiles when none is configured

Perform this step only if you cannot read the MCP configuration files, or they do not define `AWS_PROFILE`.

Run `aws-vault list` and present the user the entries from the "Profile" column.

For example, given this `aws-vault list` output:

```text
Profile                  Credentials              Sessions
=======                  ===========              ========
personal                 personal                 -
iprod                    -                        oidc:nwl
iprod-admin              -                        oidc:nwl
iprod-ro                 -                        oidc:nwl
mgmt-admin               -                        oidc:nwl
msa                      -                        oidc:nwl
msa-admin                -                        oidc:nwl
msa-ro                   -                        oidc:nwl, sso.GetRoleCredentials:7h58m47s
```

…present these profiles to the user:

* `iprod-admin`
* `iprod-ro`
* `iprod`
* `mgmt-admin`
* `msa-admin`
* `msa-ro`
* `msa`
* `personal`

**CRITICAL:** NEVER ASSUME WHICH PROFILE THE USER WANTS. The user MUST explicitly select one before you continue.

## Step 3: Authenticate

Run the command below, substituting the profile value from Step 1 or Step 2. It opens a web browser for the user to complete authentication — wait patiently for them to finish before continuing.

```bash
aws sso login --profile <profile-value>
```

For example, if the profile is `sandbox`:

```bash
aws sso login --profile sandbox
```

## Step 4: Verify

Confirm success by making an AWS MCP call (for example, asking "What AWS Regions are available?"). A successful response means authentication worked and you can proceed with the user's original task.
<!-- @config-manager:end add-code-comments -->
