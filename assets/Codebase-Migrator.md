

# Cross-Repository Synchronization with GitHub Actions

## Overview

This guide shows you how to automatically sync code changes from one GitHub repository to another using GitHub Actions. Whenever you push changes to a specific branch in your source repository, those changes will automatically be copied to a branch in a different repository (even if it's owned by a different GitHub account).

## Use Case

**Scenario**: You have:
- **Source Repository**: `your-account/repo-x` with branches `a` and `b`
- **Target Repository**: `different-account/repo-y` with branch `c`

**Goal**: When you push changes to branch `a` in `repo-x`, automatically update branch `c` in `repo-y`.

## Prerequisites

Before starting, ensure you have:

1. **Write access** to the source repository (`repo-x`)
2. **Write access** to the target repository (`repo-y`) or collaboration permissions
3. **GitHub account** with permissions to create Personal Access Tokens

## Step-by-Step Setup Guide

### Step 1: Create a Personal Access Token (PAT)

A Personal Access Token is like a password that allows GitHub Actions to access other repositories on your behalf.

1. **Log into the GitHub account** that owns the target repository (`repo-y`)
2. **Navigate to Settings**:
   - Click your profile picture (top-right corner)
   - Select "Settings" from dropdown menu
3. **Access Developer Settings**:
   - Scroll down to "Developer settings" in the left sidebar
   - Click "Personal access tokens"
   - Select "Tokens (classic)"
4. **Generate New Token**:
   - Click "Generate new token (classic)"
   - Add a descriptive note: "Cross-repo sync for repo-x to repo-y"
   - Set expiration (recommend 90 days for security)
   - **Select scopes**: Check `repo` (this gives full repository access)
5. **Copy the token**: Save it somewhere safe - you won't see it again!

> **Security Note**: Treat this token like a password. Never share it or commit it to your repository.

### Step 2: Add Token as Repository Secret

Repository secrets are encrypted and only accessible to GitHub Actions in your repository.

1. **Go to your source repository** (`repo-x`)
2. **Navigate to Settings**:
   - Click the "Settings" tab in your repository
3. **Access Secrets**:
   - In the left sidebar, click "Secrets and variables"
   - Select "Actions"
4. **Create New Secret**:
   - Click "New repository secret"
   - **Name**: `TARGET_REPO_TOKEN`
   - **Secret**: Paste the Personal Access Token you created
   - Click "Add secret"

### Step 3: Create the GitHub Actions Workflow

This workflow file tells GitHub Actions what to do when changes are pushed.

1. **Create the workflow directory** in your source repository (`repo-x`):
   ```
   .github/
   └── workflows/
   ```

2. **Create the workflow file**: `.github/workflows/sync-to-external-repo.yml`

```yaml
name: Sync to External Repository

# This workflow runs when code is pushed to branch 'a'
on:
  push:
    branches:
      - a  # Only trigger when pushing to branch 'a'

jobs:
  sync:
    runs-on: ubuntu-latest
    
    steps:
      # Step 1: Download the source code from branch 'a'
      - name: Checkout source repository
        uses: actions/checkout@v4
        with:
          fetch-depth: 0  # Get complete history for proper sync
          
      # Step 2: Set up Git with a name and email for commits
      - name: Configure Git
        run: |
          git config --global user.name "GitHub Actions Bot"
          git config --global user.email "actions@github.com"
          
      # Step 3: Connect to the target repository
      - name: Add target repository as remote
        run: |
          git remote add target https://${{ secrets.TARGET_REPO_TOKEN }}@github.com/different-account/repo-y.git
          
      # Step 4: Push changes to target repository
      - name: Push to target repository
        run: |
          git push target a:c --force
        env:
          GITHUB_TOKEN: ${{ secrets.TARGET_REPO_TOKEN }}

      # Step 5: Clean up the connection
      - name: Cleanup
        run: |
          git remote remove target
```

### Step 4: Customize for Your Repositories

Update the workflow file with your actual repository information:

1. **Replace repository names**:
   - Change `different-account/repo-y` to your actual target repository
   - Example: `johnsmith/my-target-repo`

2. **Adjust branch names if needed**:
   - Source branch: Change `a` to your desired source branch
   - Target branch: Change `c` to your desired target branch
   - The format `a:c` means "push from branch `a` to branch `c`"

## 🔄 Alternative Workflow: File-by-File Copy

If you need more control over what gets synchronized, use this alternative approach:

```yaml
name: Sync Files to External Repository

on:
  push:
    branches:
      - a

jobs:
  sync:
    runs-on: ubuntu-latest
    
    steps:
      # Download source repository
      - name: Checkout source repository
        uses: actions/checkout@v4
        with:
          path: source-repo
          
      # Download target repository
      - name: Checkout target repository
        uses: actions/checkout@v4
        with:
          repository: different-account/repo-y
          token: ${{ secrets.TARGET_REPO_TOKEN }}
          path: target-repo
          ref: c  # Checkout branch 'c'
          
      # Copy files from source to target
      - name: Copy files
        run: |
          # Remove existing files in target (optional)
          rm -rf target-repo/*
          # Copy files (excluding .git directory)
          cp -r source-repo/* target-repo/ || true
          cp -r source-repo/.[!.]* target-repo/ || true
          
      # Commit and push changes
      - name: Commit and push changes
        run: |
          cd target-repo
          git add .
          # Only commit if there are changes
          git diff --staged --quiet || git commit -m "Sync from repo-x branch a - $(date)"
          git push origin c
```

## Testing Your Setup

1. **Make a test change** in branch `a` of your source repository
2. **Commit and push** the changes:
   ```bash
   git add .
   git commit -m "Test sync workflow"
   git push origin a
   ```
3. **Check the Actions tab** in your source repository to see if the workflow runs
4. **Verify** that changes appear in branch `c` of your target repository

## Troubleshooting

### Common Issues and Solutions

#### "Permission denied" or "403 Forbidden"
- **Cause**: Invalid or expired Personal Access Token
- **Solution**: 
  - Regenerate the PAT with proper `repo` permissions
  - Update the `TARGET_REPO_TOKEN` secret

#### "Repository not found"
- **Cause**: Incorrect repository name in workflow
- **Solution**: Double-check the repository name format: `owner/repository-name`

#### "Branch not found"
- **Cause**: Target branch doesn't exist
- **Solution**: The workflow will create the branch automatically, or create it manually first

#### Workflow doesn't trigger
- **Cause**: Workflow file in wrong location or incorrect branch name
- **Solution**: 
  - Ensure file is at `.github/workflows/sync-to-external-repo.yml`
  - Check that you're pushing to the correct branch (default is `a`)

### Debugging Steps

1. **Check Actions Log**:
   - Go to your source repository
   - Click "Actions" tab
   - Click on the failed workflow run
   - Expand each step to see detailed error messages

2. **Verify Permissions**:
   - Ensure your PAT has `repo` scope
   - Confirm you have write access to the target repository

3. **Test Manually**:
   ```bash
   # Clone target repository with your token
   git clone https://YOUR_TOKEN@github.com/target-account/repo-y.git
   # If this fails, your token or permissions are incorrect
   ```

## Configuration Options

### Modify Trigger Conditions

```yaml
# Trigger on multiple branches
on:
  push:
    branches:
      - a
      - main
      - develop

# Trigger on specific file changes
on:
  push:
    branches:
      - a
    paths:
      - 'src/**'
      - '*.md'
```

### Add File Filtering

```yaml
# Exclude certain files/directories
- name: Copy files with exclusions
  run: |
    rsync -av --exclude='.git' --exclude='node_modules' --exclude='*.log' source-repo/ target-repo/
```

### Custom Commit Messages

```yaml
- name: Push to target repository
  run: |
    git push target a:c
    # Custom commit message based on source commit
    git log -1 --pretty=format:"Sync: %s (from repo-x)" | git commit --amend -F -
```

## Security Best Practices

1. **Use least privilege**: Only grant necessary permissions to PATs
2. **Set token expiration**: Don't create tokens that never expire
3. **Rotate tokens regularly**: Update tokens every 90 days
4. **Monitor usage**: Check which repositories are using your tokens
5. **Use repository secrets**: Never hardcode tokens in workflow files

## Additional Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Managing Personal Access Tokens](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token)
- [Encrypted Secrets](https://docs.github.com/en/actions/security-guides/encrypted-secrets)

## FAQ

**Q: Can I sync multiple branches?**
A: Yes, modify the `branches` section in the workflow trigger and add multiple push commands.

**Q: What happens if both repositories are modified simultaneously?**
A: The workflow uses `--force` push, so the source repository changes will overwrite target changes. Remove `--force` if you want to handle conflicts manually.

**Q: Can I sync to multiple target repositories?**
A: Yes, add multiple remote repositories and push commands in your workflow.

**Q: How do I sync only specific files?**
A: Use the file-by-file copy approach with rsync or cp commands to selectively copy files.

---

**You're all set!** Your repositories will now automatically stay in sync whenever you push changes to your specified source branch.