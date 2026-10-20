#!/usr/bin/env bash

gfind configs/templates/.kiro/steering/ -type f -exec basename {} \; \
    | gxargs -I% bash -c 'set -x; stubs -output configs/stubs/.kiro/steering/% configs/templates/.kiro/steering/%'
