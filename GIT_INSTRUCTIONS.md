# How to Create a Public Git Repository and Protect Your Main Branch

This guide will walk you through the steps to create a new public Git repository on GitHub for your project. By default, public repositories on GitHub do not allow users who are not explicitly added as collaborators to push code directly to your branches. This setup is ideal for open-source projects where you want to encourage community contributions through pull requests while maintaining control over the main codebase.

## Step 1: Initialize a Local Git Repository

First, you need to initialize a Git repository in your project's root directory.

```bash
# Make sure you are in your project's root directory
# (the one containing your go.mod file)

# Initialize the git repository
git init -b main
```

This command creates a new `.git` subdirectory in your project directory, which contains all the necessary repository files. The `-b main` flag sets the default branch name to `main`.

## Step 2: Add and Commit Your Project Files

Next, you'll want to add your project's files to the staging area and then commit them to the repository's history.

```bash
# Add all files in the current directory to the staging area
git add .

# Commit the staged files with a descriptive message
git commit -m "Initial commit: Create GoFigirl TUI application"
```

This creates the first snapshot of your project in the repository's history.

## Step 3: Create a New Public Repository on GitHub using `gh` CLI

You can create a new repository directly from your terminal using the GitHub CLI (`gh`). If you don't have `gh` installed, you can find installation instructions on the [GitHub CLI documentation](https://cli.github.com/).

```bash
# Make sure you are logged in to GitHub with the gh CLI
gh auth login

# Create a new public repository
# Replace 'gofigirl-tui' with your desired repository name
# The '--public' flag makes the repository public
# The '--source .' flag uploads the current directory to the repository
# The '--remote upstream' flag sets the remote name to 'upstream' instead of 'origin'
# The '--push' flag pushes the local branch to the remote repository
gh repo create gofigirl-tui --public --source=. --remote=upstream --push
```

**Explanation of flags:**

*   `gofigirl-tui`: This is the name your new repository will have on GitHub. Change it if you want a different name.
*   `--public`: This makes your repository public. If you wanted a private repository, you would use `--private`.
*   `--source=.`: This tells `gh` to create the repository from the current directory. It will automatically add and commit your local files.
*   `--remote=upstream`: This sets the name of the remote to `upstream`. You can choose any name you like, but `origin` is a common convention for your main remote. Since we are using `gh` to push everything in one go, we will rename the default `origin` to `upstream` to avoid any conflicts with the default remote name `origin` that `gh` might create internally.
*   `--push`: This flag pushes your local `main` branch to the newly created GitHub repository. You don't need to manually run `git push` afterwards.

After running this command, your local repository will be created, pushed to GitHub, and linked to the remote. You can then skip Step 4 (pushing to GitHub).


## Step 4: Push Your Local Repository to GitHub

After creating the repository on GitHub, you will be on the repository's main page. GitHub will provide you with a URL for your new repository.

Now, you need to link your local repository to the remote repository on GitHub and push your code.

```bash
# Replace <repository-url> with the URL you copied from GitHub
# It will look something like: https://github.com/your-username/gofigirl-tui.git

git remote add origin <repository-url>

# Push your 'main' branch to the remote repository named 'origin'
git push -u origin main
```

Your code is now on GitHub!

## How Public Commits are Handled

By default, when you create a public repository on GitHub:

-   **Anyone can see your code.**
-   **Only you and any collaborators you explicitly add can push directly to the repository.**

Other users who want to contribute must first **fork** your repository, make their changes in their own copy, and then submit a **pull request** to your repository. You can then review their changes and decide whether to merge them into your `main` branch. This is the standard workflow for open-source projects and provides a secure way to manage contributions.
