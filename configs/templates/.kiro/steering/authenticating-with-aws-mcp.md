---
# @config-manager:start add-code-comments-frontmatter
inclusion: auto
name: authenticating-with-aws-mcp
description: Instructions for authenticating with AWS in order to leverage AWS MCP and Knowledge servers.
# @config-manager:end add-code-comments-frontmatter
---

<!-- @config-manager:start add-code-comments -->
# Authenticating with AWS MCP

The Agent Toolkit for AWS gives AI coding agents the tools, knowledge, and guardrails they need to build, deploy, and manage applications on AWS. It works with the coding agents that developers already use, such as Kiro, Claude Code, Cursor, and Codex. You don't need to switch tools or learn a new workflow.

AI coding agents can handle common AWS tasks like creating an Amazon S3 bucket or launching an Amazon EC2 instance. However, they often struggle with complex, multi-step workflows. They choose the wrong service for a use case, misconfigure resources, or repeatedly retry operations on newer services they were not trained on. The Agent Toolkit for AWS addresses these gaps. It gives agents secure access to AWS services, current documentation, and step-by-step guidance.

* [What is the Agent Toolkit for AWS?](https://docs.aws.amazon.com/agent-toolkit/latest/userguide/what-is-agent-toolkit.html)
* [Getting started](https://docs.aws.amazon.com/agent-toolkit/latest/userguide/quick-start.html)

IMPORTANT: If you are able to call the AWS MCP tools successfully, disregard the rest of this document as you are already authenticated.

## Reading the MCP configuration

If you are unable to call the AWS MCP tools successfully, it is likely that you need to authenticate.

Read both `~/.kiro/settings/mcp.json` and `<workspace>/.kiro/settings/mcp.json` to identify an MCP server which contains an arg containing `https://aws-mcp.us-east-1.api.aws/mcp`. This MCP configuration is the one which configures the AWS Agent Toolkit MCP server. If there is an arg containing a definition for `AWS_PROFILE` (e.g., `AWS_PROFILE=sandbox`), or an `env` entry for `AWS_PROFILE` (e.g., `"env":{"AWS_PROFILE": "sandbox"}`), then we will use that profile value in the next step.

## Alternative paths

If you cannot read the MCP configuration files, or if the MCP configuration files do not contain a definition for `AWS_PROFILE`, then you can provide a list of available profiles to the user by running `aws-vault list` and presenting the entries listed in the "Profile" column.

For example, if the output of `aws-vault list` is:

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

…then the list of profiles to present to the user would be:

* `iprod-admin`
* `iprod-ro`
* `iprod`
* `mgmt-admin`
* `msa-admin`
* `msa-ro`
* `msa`
* `personal`

CRITICAL: NEVER ASSUME WHICH PROFILE THE USER WANTS TO USE. MAKE THE USER EXPLICITLY SELECT THE ONE THEY WANT.

## Authenticating

On the shell, execute the following command. This will open a web browser window for the user to interact with for authentication. You should wait patiently for the user to complete their task before continuing. You will know that you have been successful when you can ask "What AWS Regions are available?" and get back a successful response.

```bash
aws sso login --profile <profile-value>
```

For example:

```bash
# If AWS_PROFILE=sandbox is set, then run:
aws sso login --profile sandbox
```
<!-- @config-manager:end add-code-comments -->
