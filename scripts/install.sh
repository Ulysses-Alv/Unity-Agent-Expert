#!/usr/bin/env bash
# Unity-Agent-Expert CLI Installer (Bootstrapper)
# Usage: curl -fsSL https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.sh | bash
#
# This script downloads and installs the unity-agent-expert binary.
# After installation, run: unity-agent-expert --help

set -euo pipefail

REPO="Ulysses-Alv/Unity-Agent-Expert"
TOOL_NAME="unity-agent-expert"
INSTALL_DIR="${HOME}/.local/bin"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
CYAN='\033[0;36m'
NC='\033[0m'

detect_os() {
    case "$(uname -s)" in
        Linux*)     echo "linux" ;;
        Darwin*)    echo "macos" ;;
        *)          echo "unsupported" ;;
    esac
}

detect_arch() {
    local arch="$(uname -m)"
    case "$arch" in
        x86_64)     echo "amd64" ;;
        aarch64)    echo "arm64" ;;
        arm64)      echo "arm64" ;;
        *)          echo "$arch" ;;
    esac
}

main() {
    local os=$(detect_os)
    local arch=$(detect_arch)

    if [ "$os" == "unsupported" ]; then
        echo -e "${RED}Error: Unsupported OS${NC}"
        echo "Supported: Linux, macOS"
        exit 1
    fi

    echo ""
    echo -e "${CYAN}========================================${NC}"
    echo -e "${CYAN} unity-agent-expert installer${NC}"
    echo -e "${CYAN}========================================${NC}"
    echo ""

    # Get latest release
    echo -e "${CYAN}[info] Fetching latest release...${NC}"
    local api_url="https://api.github.com/repos/${REPO}/releases/latest"
    local asset_name="${TOOL_NAME}-${os}-${arch}.exe"

    local response=$(curl -fsSL "$api_url" -H "Accept: application/vnd.github.v3+json")
    local download_url=$(echo "$response" | grep -o '"browser_download_url": "[^"]*'" | grep "$asset_name" | cut -d'"' -f4)

    if [ -z "$download_url" ]; then
        echo -e "${RED}[error] Asset not found: $asset_name${NC}"
        echo "If this error persists, please open an issue at:"
        echo "  https://github.com/${REPO}/issues"
        exit 1
    fi

    # Download
    echo -e "${CYAN}[info] Downloading $asset_name...${NC}"
    local tmp_file="/tmp/${TOOL_NAME}-$$"
    curl -fsSL "$download_url" -o "$tmp_file"

    # Install
    mkdir -p "$INSTALL_DIR"
    mv "$tmp_file" "${INSTALL_DIR}/${TOOL_NAME}"
    chmod +x "${INSTALL_DIR}/${TOOL_NAME}"

    echo ""
    echo -e "${GREEN}========================================${NC}"
    echo -e "${GREEN} unity-agent-expert installed!${NC}"
    echo -e "${GREEN}========================================${NC}"
    echo ""
    echo -e "Location: ${INSTALL_DIR}/${TOOL_NAME}"
    echo ""
    echo -e "${YELLOW}To use, add to your PATH:${NC}"
    echo ""
    echo -e "  Linux/macOS: export PATH=\"\$PATH:${INSTALL_DIR}\""
    echo ""
    echo -e "Add this to your ~/.bashrc or ~/.zshrc to persist across sessions."
    echo ""
    echo -e "${YELLOW}Then run:${NC}"
    echo -e "  ${CYAN}${TOOL_NAME} --help${NC}"
    echo ""
}

main "$@"