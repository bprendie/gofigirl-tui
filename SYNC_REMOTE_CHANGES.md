To sync your local repository with the changes you made on the remote (GitHub) side, you can use the `gh repo sync` command. This command is designed to fetch and integrate changes from your remote branch into your local branch.

Here's how to do it:

```bash
gh repo sync upstream --branch main
```

**Explanation:**

*   `gh repo sync`: This command synchronizes your current local branch with its upstream remote.
*   `upstream`: This is the name of your remote repository. If you used `origin` when setting up your remote, you would use `gh repo sync origin --branch main` instead.
*   `--branch main`: Specifies that you want to sync the `main` branch.

Alternatively, you can use the standard Git command:

```bash
git pull upstream main
```

Both commands will fetch the changes from your remote `main` branch and merge them into your local `main` branch, making your local repository match the remote. Choose the command you feel most comfortable with.