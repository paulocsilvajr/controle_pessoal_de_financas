#!/usr/bin/env bash

CURRENTBRANCH=$(git rev-parse --abbrev-ref HEAD)
BRANCH="api"

if [ "$BRANCH" = "$CURRENTBRANCH" ]; then
    git push origin api && git push bitbucket api && git push github api
else
    echo "Altere para a branch '$BRANCH'(git checkout '$BRANCH') para executar esse script"
    exit 1
fi

