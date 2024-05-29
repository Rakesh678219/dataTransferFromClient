#!/bin/bash

# Increment the version number
# CURRENT_VERSION=$(git tag --sort=-v:refname | head -n 1)
# IFS='.' read -r -a version_parts <<< "$(echo "$CURRENT_VERSION")"
# major="${version_parts[0]}"
# minor="${version_parts[1]}"
# new_minor=$((minor + 1))
NEW_VERSION="v1.31.0"

# Commit changes
git add .
git commit -m "Version $NEW_VERSION"
git tag $NEW_VERSION

# Push commits and tags
git push origin main
git push origin $NEW_VERSION

# Ensure Go module integrity
go mod tidy

# Verify module version availability
if GOPROXY=proxy.golang.org go list -m github.com/Rakesh678219/dataTransferFromClient@$NEW_VERSION &> /dev/null; then
    echo "Module version $NEW_VERSION is available."
else
    echo "Module version $NEW_VERSION is not available."
fi
