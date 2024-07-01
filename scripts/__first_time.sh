#!/usr/bin/env bash

GIT_CLONE_PROTECTION_ACTIVE=false

# Root directory of the repository.
# shellcheck disable=2312
ROOT_DIR="$(dirname "$(realpath "${BASH_SOURCE[0]}")")"

# Remove on the runner.
RUNNER_TEMP="/tmp/terraform-makefile"

# Clone repo into TMP directory.
rm -Rf "${RUNNER_TEMP}"
git clone \
    --depth 1 \
    --branch main \
    --single-branch \
    https://github.com/northwood-labs/.github.git \
    "${RUNNER_TEMP}" \
    ;

################################################################################
# FULL COPY

# Copy all "full-copy" files from the root into the repository.
find "${RUNNER_TEMP}/full-copy/" -maxdepth 1 -type f -print0 |
    xargs -0 -I% cp -Rfv "%" "${PWD}" ||
    true

# Folders to copy
FOLDERS=(
    ".githooks"
    ".github"
    ".vscode"
    "scripts"
)

for FOLDER in "${FOLDERS[@]}"; do
    # Copy all files from this directory into the root of the repository.
    mkdir -p "${PWD}/${FOLDER}"
    find "${RUNNER_TEMP}/full-copy/${FOLDER}/" -maxdepth 1 -type f -print0 |
        xargs -0 -I% cp -Rfv "%" "${PWD}/${FOLDER}" ||
        true
done

TYPES=()

# Pass GO=true when calling the script.
# shellcheck disable=2154
if [[ "${GO}" == "true" ]]; then
    TYPES+=("go")
fi

# Pass TF_MOD=true when calling the script.
# shellcheck disable=2154
if [[ "${TF_MOD}" == "true" ]]; then
    TYPES+=("go")
    TYPES+=("tf-mod")
fi

for TYPE in "${TYPES[@]}"; do
    # Copy all files from this directory into the root of the repository.
    mkdir -p "${PWD}"
    find "${RUNNER_TEMP}/full-copy/${TYPE}/" -maxdepth 1 -type f -not \( -name "*tmpl*" \) -print0 |
        xargs -0 -I% cp -Rfv "%" "${PWD}" ||
        true
done

cp -Rfv "${RUNNER_TEMP}/scripts/__update.sh" "${ROOT_DIR}/__update.sh"

################################################################################
# UPDATES

# Copy all "updates" files from the root into the repository.
find "${RUNNER_TEMP}/updates/" -maxdepth 1 -type f -not \( -name "*tmpl*" \) -print0 |
    xargs -0 -I% cp -Rfv "%" "${PWD}" ||
    true

# Folders to copy
FOLDERS=(
    ".github"
    ".vscode"
)

for FOLDER in "${FOLDERS[@]}"; do
    # Copy all files from this directory into the root of the repository.
    mkdir -p "${PWD}/${FOLDER}"
    find "${RUNNER_TEMP}/updates/${FOLDER}/" -maxdepth 1 -type f -not \( -name "*tmpl*" \) -print0 |
        xargs -0 -I% cp -Rfv "%" "${PWD}/${FOLDER}" ||
        true
done

TYPES=()

# Pass GO=true when calling the script.
# shellcheck disable=2154
if [[ "${GO}" == "true" ]]; then
    TYPES+=("go")
fi

# Pass TF_MOD=true when calling the script.
# shellcheck disable=2154
if [[ "${TF_MOD}" == "true" ]]; then
    TYPES+=("tf-mod")
fi

for TYPE in "${TYPES[@]}"; do
    # Copy all files from this directory into the root of the repository.
    mkdir -p "${PWD}"
    find "${RUNNER_TEMP}/updates/${TYPE}/" -maxdepth 1 -type f -not \( -name "*tmpl*" \) -print0 |
        xargs -0 -I% cp -Rfv "%" "${PWD}" ||
        true
done

################################################################################
# FINISH UP

# Generate files
cd "${RUNNER_TEMP}/generate/goplicate" && go run generate.go
cd "${ROOT_DIR}" || true
cp -vf "${RUNNER_TEMP}/generate/goplicate/.goplicate.yaml" "${ROOT_DIR}/.goplicate.yaml"

# Run Goplicate
goplicate run --allow-dirty --confirm --stash-changes --debug

# Generate .ecrc
tomljson ecrc.toml >.ecrc

# Make shell scripts executable
find "${PWD}" -type f -name "*.sh" -print0 |
    xargs -0 -I% chmod +x "%" ||
    true

# Delete self
# shellcheck disable=2154
if [[ "${TEST}" != "true" ]]; then
    rm -vf "${ROOT_DIR}/__first_time.sh"
fi
