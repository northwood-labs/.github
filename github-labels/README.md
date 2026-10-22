# Northwood Labs — GitHub issue labels

* Uses <https://github.com/scttfrdmn/gh-label-sync> to sync.
* Tag list inspired by <https://github.com/warengonzaga/github-labels-template>.

## Export

```bash
gh label-sync export --repo northwood-labs/.github > labels.yml
```

## Import/overwrite

```bash
gh label-sync sync --file labels.yml --repo northwood-labs/.github --delete-unmanaged --force
```

## Apply to all

```bash
for repo in $(gh repo list northwood-labs --limit 1000 --json nameWithOwner -q '.[].nameWithOwner'); do
  gh label-sync sync --file labels.yml --repo "${repo}" --delete-unmanaged --force --yes
done
```

## Configure default tags for new repositories

<https://github.com/organizations/northwood-labs/settings/repository-defaults>
