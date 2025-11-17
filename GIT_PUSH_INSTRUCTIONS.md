To push your local changes to the `main` branch of your Git repository, follow these steps:

1.  **Stage your changes:**
    This command adds all your modified and new files to the staging area.

    ```bash
    git add .
    ```

2.  **Commit your changes:**
    This command records your changes to the repository with a descriptive message.

    ```bash
    git commit -m "feat: Add direct URL playback and update README.md"
    ```

3.  **Push your changes to GitHub:**
    This command uploads your committed changes to your `main` branch on the remote repository (which is typically named `origin` or `upstream`).

    ```bash
    git push origin main
    ```

    *If you followed the `gh repo create` instructions with `--remote=upstream`, you might need to use:*

    ```bash
    git push upstream main
    ```

    After running this, your changes will be live on your GitHub repository.
