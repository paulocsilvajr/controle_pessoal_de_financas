#!/usr/bin/env bash

CURRENTBRANCH=$(git rev-parse --abbrev-ref HEAD)
BRANCH="api-gorm"

if [ "$BRANCH" = "$CURRENTBRANCH" ]; then
    git push origin api-gorm && git push bitbucket api-gorm && git push github api-gorm
else
    echo "Altere para a branch '$BRANCH'(git checkout '$BRANCH') para executar esse script"
    exit 1
fi
