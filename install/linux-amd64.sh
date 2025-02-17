#!/bin/bash

# Variables
REPO="st-osi/krow"
VERSION="v0.0.2-dev.3"
BINARY_NAME="krow"
TAR_FILE="${BINARY_NAME}-${VERSION}-linux-amd64.tar.gz"
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/${TAR_FILE}"

# Download the tar.gz file
echo "Downloading ${TAR_FILE} from ${DOWNLOAD_URL}..."
curl -L -o /tmp/${TAR_FILE} ${DOWNLOAD_URL}

# Extract the tar.gz file
echo "Extracting ${TAR_FILE}..."
tar -xzf /tmp/${TAR_FILE} -C /tmp

# Move the binary to /usr/local/bin
echo "Installing ${BINARY_NAME}..."
sudo mv /tmp/${BINARY_NAME} /usr/local/bin/${BINARY_NAME}

# Clean up
echo "Cleaning up..."
rm /tmp/${TAR_FILE}
rm /tmp/${BINARY_NAME}

echo "Installation completed. You can now run '${BINARY_NAME}' from anywhere."
