#!/usr/bin/env bash

git remote -v
git remote add upstream https://github.com/hidevopsio/hioak.git
git remote -v
sleep 5
git fetch upstream
git rebase upstream/master
