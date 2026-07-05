#!/bin/bash
# This script fixes common Git and GitHub issues.

# 1. Fix for "fatal: ref refs/remotes/origin/HEAD is not a symbolic ref"
echo "Setting remote HEAD for origin..."
git remote set-head origin --auto > /dev/null 2>&1
echo "Done."

# 2. Fix for "GraphQL: Auto merge is not allowed for this repository"
echo "Attempting to enable auto-merge for the repository..."
REPO=$(git config --get remote.origin.url | sed -e 's/.*github.com[:|/]//' -e 's/\.git$//')
if [ -n "$REPO" ]; then
  echo "Detected repository: $REPO. Enabling auto-merge..."
  gh repo edit "$REPO" --enable-auto-merge > /dev/null 2>&1
fi
echo "Done."

# 3. Fix for "Diverging branches can't be fast-forwarded"
echo "Handling diverging branches by rebasing..."
git fetch origin > /dev/null 2>&1
BASE_BRANCH="main" # Assuming 'main' is the base branch
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
# Only rebase if the current branch is not the base branch
if [ "$CURRENT_BRANCH" != "$BASE_BRANCH" ]; then
    echo "Rebasing current branch '$CURRENT_BRANCH' onto 'origin/$BASE_BRANCH'..."
    git rebase "origin/$BASE_BRANCH"
    # Check for rebase failure
    if [ $? -ne 0 ]; then
      echo "---"
      echo "Error: Automatic rebase failed, likely due to merge conflicts."
      echo "Please resolve the conflicts, then run 'git add .' and 'git rebase --continue'."
      echo "To abort, run 'git rebase --abort'."
      echo "---"
      exit 1
    fi
fi
echo "Done."

# 4. Fix for "could not find any commits" when creating a PR
echo "Checking for new commits before creating a PR..."
COMMIT_COUNT=$(git rev-list --count "origin/$BASE_BRANCH".."$CURRENT_BRANCH")

if [ "$COMMIT_COUNT" -eq 0 ] && [ "$CURRENT_BRANCH" != "$BASE_BRANCH" ]; then
  echo "---"
  echo "Error: Your branch '$CURRENT_BRANCH' has no new commits compared to 'origin/$BASE_BRANCH'."
  echo "You must have at least one commit to create a pull request."
  echo
  echo "How to fix:"
  echo "  1. Make changes to your files."
  echo "  2. Stage them: git add <your-files>"
  echo "  3. Commit them: git commit -m \"Your commit message\""
  echo "  4. Re-run your previous command (e.g., ghop)."
  echo "---"
  exit 1
else
  echo "Found $COMMIT_COUNT new commit(s). Proceeding is safe."
fi

echo "All checks passed. Your Git state seems healthy."
