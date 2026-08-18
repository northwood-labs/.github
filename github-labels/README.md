# Northwood Labs — GitHub issue labels

Uses <https://github.com/scttfrdmn/gh-label-sync> and <https://github.com/warengonzaga/github-labels-template>.

## Export

```bash
gh label-sync export --repo northwood-labs/.github > labels.yml
```

## Import/overwrite

```bash
gh label-sync sync --file labels.yml --repo northwood-labs/.github --delete-unmanaged --force
```
