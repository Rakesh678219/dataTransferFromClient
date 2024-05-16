#!/bin/bash

# Step 1: Code changes
# Make your desired changes to the code

# Step 2: Versioning
VERSION=$(TZ=Asia/Kolkata date +"%Y-%m-%d-%H-%M")
echo "Version: $VERSION"

# Step 3: Commit changes
echo "Committing changes..."
git add .
git commit -m "Changes for version $VERSION"
echo "Changes committed successfully."

# Step 4: Tagging
echo "Tagging commit with version v-$VERSION..."
git tag v-$VERSION
echo "Tag created successfully."

# Step 5: Push commits and tags
echo "Pushing commits to remote repository..."
git push origin main
echo "Commits pushed successfully."

echo "Pushing tags to remote repository..."
git push origin v-$VERSION
echo "Tags pushed successfully."

# Step 6: Ensure Go module integrity
echo "Ensuring Go module integrity..."
go mod tidy
echo "Go module integrity ensured."
