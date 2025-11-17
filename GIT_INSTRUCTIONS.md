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

## Step 3: Create a New Public Repository on GitHub

Now, you need to create a new repository on GitHub.

1.  Go to [GitHub](https://github.com) and log in to your account.
2.  In the upper-right corner of any page, click the `+` icon, and then click **New repository**.
3.  Name your repository (e.g., `gofigirl-tui`).
4.  Ensure that **Public** is selected.
5.  **Do not** initialize the repository with a `README`, `.gitignore`, or `license`. You've already created these files locally.
6.  Click **Create repository**.

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
