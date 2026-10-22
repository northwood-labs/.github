#!/usr/bin/env bash

gum style --foreground='#fffc67' --border='rounded' --border-foreground='12' --padding='1 2' --bold \
    "This will break multi-source stubs. Fix them before committing."

cd configs/templates/ && \
fd --unrestricted --exclude='.DS_Store' --exclude='__*' --type=file --color=never --strip-cwd-prefix=always --threads=1 \
    --exec stubs -output "../../configs/stubs/{}" "../../configs/templates/{}" \;
