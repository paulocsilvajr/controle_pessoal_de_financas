#!/usr/bin/env bash

CURRENTBRANCH=$(git rev-parse --abbrev-ref HEAD)
BRANCH="web"

if [ "$BRANCH" = "$CURRENTBRANCH" ]; then
    git push origin $BRANCH && git push bitbucket $BRANCH && git push github $BRANCH
else
    echo "Altere para a branch '$BRANCH'(git checkout '$BRANCH') para executar esse script"
    exit 1
fi

