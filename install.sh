#!/bin/bash

# --- Find a suitable installation directory from the user's PATH ---
suggested_path=""
# Prioritize ~/.local/bin if it exists and is in the PATH
if [[ ":$PATH:" == *":$HOME/.local/bin:"* ]] && [ -d "$HOME/.local/bin" ]; then
    suggested_path="$HOME/.local/bin"
else
    # Otherwise, find the first user-writable directory in the PATH
    IFS=':' read -ra path_dirs <<< "$PATH"
    for dir in "${path_dirs[@]}"; do
        # Expand tilde if present
        eval expanded_dir=$dir
        if [ -d "$expanded_dir" ] && [ -w "$expanded_dir" ]; then
            suggested_path=$expanded_dir
            break
        fi
    done
fi

if [ -z "$suggested_path" ]; then
    echo "Could not find a suitable installation directory in your PATH."
    echo "Please add a directory to your PATH (e.g., ~/.local/bin) and try again."
    exit 1
fi

# --- Prompt the user for the installation path ---
echo "We suggest installing gofigirl to the following directory in your PATH:"
echo "  $suggested_path"
read -p "Press ENTER to accept, or specify a different path: " user_path

# If the user provides a path, use it. Otherwise, use the suggestion.
install_path="${user_path:-$suggested_path}"

# Expand tilde in user-provided path
eval install_path=$install_path

# --- Verify the chosen directory exists and is in the PATH ---
if [[ ":$PATH:" != *":$install_path:"* ]]; then
    echo "Warning: The chosen directory '$install_path' is not in your PATH."
    read -p "Continue anyway? (y/N) " confirm
    if [[ "$confirm" != "y" ]] && [[ "$confirm" != "Y" ]]; then
        echo "Installation cancelled."
        exit 1
    fi
fi

if [ ! -d "$install_path" ]; then
    echo "Directory '$install_path' does not exist."
    read -p "Create it now? (y/N) " create_confirm
    if [[ "$create_confirm" == "y" ]] || [[ "$create_confirm" == "Y" ]]; then
        mkdir -p "$install_path"
        if [ $? -ne 0 ]; then
            echo "Failed to create directory. Installation cancelled."
            exit 1
        fi
    else
        echo "Installation cancelled."
        exit 1
    fi
fi

# --- Build the binary first ---
echo "Building the 'gofigirl' binary..."
go build -o gofigirl
if [ $? -ne 0 ]; then
    echo "Failed to build the binary. Please make sure you have Go installed."
    exit 1
fi


# --- Copy files and set permissions ---
echo "Installing to $install_path..."

cp gofigirl "$install_path/"
cp presets.json "$install_path/"

if [ $? -ne 0 ]; then
    echo "Failed to copy files. Do you have the correct permissions?"
    exit 1
fi

chmod +x "$install_path/gofigirl"

echo ""
echo "✅ Installation complete!"
echo "You can now run 'gofigirl' from anywhere in your terminal."

