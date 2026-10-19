---
# @config-manager:start using-github-cli-frontmatter
inclusion: auto
name: using-github-cli
description: Use when you need to understand how to use the `gh` CLI for interacting with GitHub.
# @config-manager:end using-github-cli-frontmatter
---

<!-- @config-manager:start using-github-cli -->
# Using GitHub CLI (`gh`)

[`gh`](https://cli.github.com/manual/) is GitHub's command-line interface for interacting with GitHub from a terminal or script. Prefer it over raw `curl` for authenticated requests, since it reuses the user's existing credentials and handles pagination and headers.

## When to use what

* Use a built-in command (`gh repo`, `gh pr`, `gh issue`, `gh release`, `gh secret`, `gh variable`, `gh workflow`, `gh run`) whenever one exists for the task.
* Fall back to [`gh api`](https://cli.github.com/manual/gh_api) only for operations with no built-in command.
* Pass `graphql` as the endpoint to reach the v4 (GraphQL) API; pass a REST path (for example, `repos/{owner}/{repo}/issues`) to reach the v3 (REST) API.

## Choosing between REST and GraphQL

Both APIs are reachable through `gh api`, and [global node IDs](https://docs.github.com/en/graphql/guides/using-global-node-ids) let you move between them. Pick the one that fits the task.

* Prefer GraphQL when you need to fetch related data in one round trip or want to request exactly the fields you need. It replaces several chained REST calls with a single query.
* Prefer REST for simple, single-resource reads and writes using standard HTTP verbs.

### REST (v3)

* [About the REST API](https://docs.github.com/en/rest/about-the-rest-api/about-the-rest-api)
* [Getting started with the REST API](https://docs.github.com/en/rest/using-the-rest-api/getting-started-with-the-rest-api)
* [API versions](https://docs.github.com/en/rest/about-the-rest-api/api-versions)
* [OpenAPI definition](https://raw.githubusercontent.com/github/rest-api-description/refs/heads/main/descriptions/api.github.com/api.github.com.yaml)
* [OpenAPI extensions](https://raw.githubusercontent.com/github/rest-api-description/refs/heads/main/extensions.md)

### GraphQL (v4)

* [About the GraphQL API](https://docs.github.com/en/graphql/overview/about-the-graphql-api)
* [Introduction to GraphQL](https://docs.github.com/en/graphql/guides/introduction-to-graphql)
* [GraphQL schema](https://docs.github.com/public/fpt/schema.docs.graphql)

## Examples

Create an issue with the built-in command (preferred over a raw REST call):

```bash
gh issue create --title "Bug with feature X" --body "If you do A, then B happens"
```

Call the REST API directly when no built-in command fits:

```bash
gh api --method POST repos/octocat/Spoon-Knife/issues \
  --field title="Bug with feature X" \
  --field body="If you do A, then B happens"
```

Query related data in a single GraphQL request (here, ten followers and ten followers of each):

```bash
gh api graphql --raw-field query='
  query {
    viewer {
      followers(first: 10) {
        nodes {
          login
          followers(first: 10) {
            nodes {
              login
            }
          }
        }
      }
    }
  }
'
```

The response is a JSON object mirroring the shape of the request. Achieving the same result over REST would require one call to `GET /user/followers` followed by one `GET /users/{username}/followers` per follower.
<!-- @config-manager:end using-github-cli -->
