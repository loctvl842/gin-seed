#!/bin/bash

# Function to install a Go package if it's not already installed
install_if_not_exists() {
    local package_name="$1"
    local binary_name="$2"

    if ! command -v "$binary_name" &> /dev/null
    then
        echo "Installing $package_name..."
        go install "$package_name"
    else
        echo "$package_name is already installed."
    fi
}

# Ensure Go is installed
if ! command -v go &> /dev/null
then
    echo "Go is not installed. Please install Go first."
    exit 1
fi

# Set Go bin directory to PATH (if it's not already)
export PATH=$PATH:$(go env GOPATH)/bin

# Install air for hot reload
install_if_not_exists "github.com/cosmtrek/air@latest" "air"

# Install swag for Swagger documentation
install_if_not_exists "github.com/swaggo/swag/cmd/swag@latest" "swag"

echo "All tools are installed and ready to use."
