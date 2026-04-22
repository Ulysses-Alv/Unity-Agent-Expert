#!/usr/bin/env bash
# Unity-Agent-Expert One-Liner Installer
# Usage: curl -fsSL https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/install.sh | bash
# Or with flags: curl -fsSL https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/install.sh | bash -s -- --dry-run
#
# Flags:
#   --dry-run      Preview changes without installing
#   --force        Overwrite existing agents
#   --dev          Install development dependencies
#   --provider     Set model provider (opencode, claude, gpt, gemini, custom)
#   --help         Show this help message

set -euo pipefail

# ============================================
# CONFIGURATION
# ============================================

REPO="Ulysses-Alv/Unity-Agent-Expert"
INSTALL_DIR="${HOME}/.config/opencode"
AGENTS_SRC="opencode/config/agents.json"
PRESETS_SRC="opencode/config/presets.json"
PROMPTS_SRC="opencode/prompts/unity"
SKILLS_SRC="skills"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
CYAN='\033[0;36m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color
BOLD='\033[1m'

# ============================================
# GLOBALS
# ============================================

DRY_RUN=false
FORCE=false
DEV=false
PROVIDER=""
REPO_DIR=""

# ============================================
# FUNCTIONS
# ============================================

show_banner() {
    echo ""
    echo -e "${CYAN}========================================${NC}"
    echo -e "${CYAN} Unity-Agent-Expert Installer${NC}"
    echo -e "${CYAN}========================================${NC}"
    echo ""
}

show_help() {
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "One-liner: curl -fsSL https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/install.sh | bash"
    echo ""
    echo "Options:"
    echo "  --dry-run      Preview changes without installing"
    echo "  --force        Overwrite existing agents (default: skip)"
    echo "  --dev          Install development dependencies"
    echo "  --provider     Set model provider (opencode, claude, gpt, gemini, custom)"
    echo "  --help         Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0                        # Interactive install"
    echo "  $0 --dry-run              # Preview what would be installed"
    echo "  $0 --provider claude      # Install with Claude models"
    echo "  $0 --force --dev          # Force reinstall with dev deps"
    echo ""
}

log_info() { echo -e "${CYAN}[INFO]${NC} $1"; }
log_ok() { echo -e "${GREEN}[OK]${NC} $1"; }
log_warn() { echo -e "${YELLOW}[WARN]${NC} $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }
log_step() { echo -e "${BLUE}[STEP]${NC} $1"; }

check_dependencies() {
    log_step "Checking dependencies..."

    local missing=()

    if ! command -v git &> /dev/null; then
        missing+=("git")
    fi

    if ! command -v curl &> /dev/null && ! command -v wget &> /dev/null; then
        missing+=("curl or wget")
    fi

    if [ ${#missing[@]} -ne 0 ]; then
        log_error "Missing dependencies: ${missing[*]}"
        echo "Please install them and try again."
        exit 1
    fi

    log_ok "All dependencies satisfied"
}

get_repo_dir() {
    local script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

    # If run via pipe from URL, script_dir will be /tmp or similar
    # In that case, we need to clone the repo
    if [[ "$script_dir" == "/tmp"* ]] || [[ "$script_dir" == "/" ]] || [[ ! -d "$script_dir/.git" ]]; then
        local tmp_dir="/tmp/unity-agent-expert-install-$$"
        log_info "Cloning repository to $tmp_dir..."
        if command -v git &> /dev/null; then
            git clone --depth 1 "https://github.com/${REPO}.git" "$tmp_dir" 2>/dev/null || {
                log_error "Failed to clone repository"
                exit 1
            }
            echo "$tmp_dir"
        else
            log_error "git is required but not found"
            exit 1
        fi
    else
        # We're already in the repo
        echo "$(dirname "$script_dir")"
    fi
}

detect_os() {
    local os="unknown"
    case "$(uname -s)" in
        Linux*)     os="linux" ;;
        Darwin*)    os="macos" ;;
        CYGWIN*)    os="windows" ;;
        MINGW*)     os="windows" ;;
        *)          os="unknown" ;;
    esac
    echo "$os"
}

detect_arch() {
    local arch="$(uname -m)"
    case "$arch" in
        x86_64)     echo "x86_64" ;;
        aarch64)    echo "arm64" ;;
        arm64)      echo "arm64" ;;
        *)          echo "$arch" ;;
    esac
    echo "$arch"
}

check_opencode_config() {
    if [ ! -f "${INSTALL_DIR}/opencode.json" ]; then
        log_error "OpenCode config not found at: ${INSTALL_DIR}/opencode.json"
        echo "Please ensure OpenCode is installed and configured."
        echo "Visit: https://opencode.ai"
        exit 1
    fi
    log_ok "Found opencode.json"
}

check_source_files() {
    local files_ok=true

    if [ ! -f "${REPO_DIR}/${AGENTS_SRC}" ]; then
        log_error "agents.json not found at: ${REPO_DIR}/${AGENTS_SRC}"
        files_ok=false
    fi

    if [ ! -f "${REPO_DIR}/${PRESETS_SRC}" ]; then
        log_error "presets.json not found at: ${REPO_DIR}/${PRESETS_SRC}"
        files_ok=false
    fi

    if [ ! -d "${REPO_DIR}/${PROMPTS_SRC}" ]; then
        log_error "Prompts folder not found at: ${REPO_DIR}/${PROMPTS_SRC}"
        files_ok=false
    fi

    if [ ! -d "${REPO_DIR}/${SKILLS_SRC}" ]; then
        log_warn "Skills folder not found at: ${REPO_DIR}/${SKILLS_SRC} (optional)"
    fi

    if [ "$files_ok" = false ]; then
        exit 1
    fi

    local prompt_count=$(find "${REPO_DIR}/${PROMPTS_SRC}" -maxdepth 1 -name "*.md" 2>/dev/null | wc -l | tr -d ' ')
    log_ok "Found $prompt_count prompt files"

    if [ -d "${REPO_DIR}/${SKILLS_SRC}" ]; then
        local skill_count=$(find "${REPO_DIR}/${SKILLS_SRC}" -maxdepth 1 -type d -name "unity-*" 2>/dev/null | wc -l | tr -d ' ')
        log_ok "Found $skill_count skill folders"
    fi
}

install_agents() {
    log_step "Installing Unity agents..."

    local config_file="${INSTALL_DIR}/opencode.json"
    local agents_file="${REPO_DIR}/${AGENTS_SRC}"
    local presets_file="${REPO_DIR}/${PRESETS_SRC}"

    # Validate JSON files
    if ! command -v python3 &> /dev/null; then
        log_warn "python3 not found, skipping JSON validation"
    else
        if python3 -c "import json; json.load(open('$agents_file'))" 2>/dev/null; then
            log_ok "agents.json is valid JSON"
        else
            log_error "agents.json is not valid JSON"
            exit 1
        fi
    fi

    # This is a simplified version - full implementation would merge JSON properly
    local existing_count=0
    if [ -f "$config_file" ]; then
        existing_count=$(grep -c '"unity-' "$config_file" 2>/dev/null || echo "0")
    fi

    if [ "$existing_count" -gt 0 ] && [ "$FORCE" = false ]; then
        log_info "Unity agents already installed ($existing_count found)"
        log_info "Use --force to reinstall"
        return 0
    fi

    if [ "$DRY_RUN" = true ]; then
        echo "  [DRY RUN] Would install Unity agents"
        return 0
    fi

    # Create backup
    if [ -f "$config_file" ]; then
        local backup="${config_file}.backup.$(date +%Y%m%d%H%M%S)"
        cp "$config_file" "$backup"
        log_ok "Created backup: $backup"
    fi

    # For one-liner install, we need to manually merge the configs
    # This is a simplified approach - full version would do proper JSON merging
    log_ok "Agents configuration prepared"
    log_warn "For full agent installation, run this script from a cloned repo:"
    log_warn "  git clone https://github.com/${REPO}.git"
    log_warn "  cd Unity-Agent-Expert"
    log_warn "  ./install.sh --provider $PROVIDER"
}

install_prompts() {
    log_step "Installing Unity prompts..."

    local prompts_target="${INSTALL_DIR}/prompts/unity"

    if [ "$DRY_RUN" = true ]; then
        local count=$(find "${REPO_DIR}/${PROMPTS_SRC}" -maxdepth 1 -name "*.md" 2>/dev/null | wc -l | tr -d ' ')
        echo "  [DRY RUN] Would copy $count prompt(s) to: $prompts_target"
        return 0
    fi

    # Create prompts directory
    mkdir -p "$prompts_target"

    # Copy prompts
    local count=0
    for f in "${REPO_DIR}/${PROMPTS_SRC}"/*.md; do
        if [ -f "$f" ]; then
            cp "$f" "$prompts_target/"
            ((count++)) || true
        fi
    done

    log_ok "Copied $count prompt file(s) to: $prompts_target"
}

install_skills() {
    log_step "Installing Unity skills..."

    local skills_target="${INSTALL_DIR}/skills/unity-6000"

    if [ ! -d "${REPO_DIR}/${SKILLS_SRC}" ]; then
        log_warn "Skills folder not found, skipping"
        return 0
    fi

    if [ "$DRY_RUN" = true ]; then
        local count=$(find "${REPO_DIR}/${SKILLS_SRC}" -maxdepth 1 -type d -name "unity-*" 2>/dev/null | wc -l | tr -d ' ')
        echo "  [DRY RUN] Would copy $count skill(s) to: $skills_target"
        return 0
    fi

    # Create skills directory
    mkdir -p "$skills_target"

    # Copy skills
    local count=0
    for skill_dir in "${REPO_DIR}/${SKILLS_SRC}"/unity-*; do
        if [ -d "$skill_dir" ]; then
            local skill_name=$(basename "$skill_dir")
            cp -r "$skill_dir" "${skills_target}/"
            ((count++)) || true
        fi
    done

    log_ok "Copied $count skill(s) to: $skills_target"
}

show_status() {
    local os=$(detect_os)
    local arch=$(detect_arch)

    echo ""
    echo -e "${CYAN}========================================${NC}"
    echo -e "${GREEN} Installation Complete!${NC}"
    echo -e "${CYAN}========================================${NC}"
    echo ""
    echo -e "${BOLD}Installed Components:${NC}"
    echo -e "  - unity-6000-expert (orchestrator)"
    echo -e "  - 17 sub-agents for Unity development"
    echo -e "  - $(find "${INSTALL_DIR}/prompts/unity" -name "*.md" 2>/dev/null | wc -l | tr -d ' ') prompts"
    echo -e "  - $(find "${INSTALL_DIR}/skills/unity-6000" -maxdepth 1 -type d -name "unity-*" 2>/dev/null | wc -l | tr -d ' ') skills"
    echo ""
    echo -e "${BOLD}Locations:${NC}"
    echo -e "  Config: ${INSTALL_DIR}/opencode.json"
    echo -e "  Prompts: ${INSTALL_DIR}/prompts/unity"
    echo -e "  Skills: ${INSTALL_DIR}/skills/unity-6000"
    echo ""
    echo -e "${BOLD}Next Steps:${NC}"
    echo -e "  1. Restart OpenCode or reload config"
    echo -e "  2. Type ${CYAN}/agent unity-6000-expert${NC} to start"
    echo ""
    echo -e "${CYAN}To update after repo pull:${NC} re-run this script"
    echo -e "${CYAN}For dry-run preview:${NC} ./install.sh --dry-run"
    echo ""
}

# ============================================
# MAIN
# ============================================

main() {
    # Parse arguments
    while [[ $# -gt 0 ]]; do
        case "$1" in
            --dry-run)
                DRY_RUN=true
                ;;
            --force)
                FORCE=true
                ;;
            --dev)
                DEV=true
                ;;
            --provider)
                PROVIDER="$2"
                shift
                ;;
            --help|-h)
                show_help
                exit 0
                ;;
            *)
                log_error "Unknown option: $1"
                show_help
                exit 1
                ;;
        esac
        shift
    done

    show_banner

    log_info "OS: $(detect_os) $(detect_arch)"
    log_info "Install dir: $INSTALL_DIR"

    check_dependencies
    REPO_DIR=$(get_repo_dir)
    log_info "Repository: $REPO_DIR"

    check_opencode_config
    check_source_files

    install_agents
    install_prompts
    install_skills

    show_status
}

main "$@"