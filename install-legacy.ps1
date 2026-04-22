# Unity-Agent-Expert One-Liner Installer
# Usage: irm https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/install.ps1 | iex
# Or with flags: irm https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/install.ps1 -useb | iex -ArgumentList "-DryRun"
#
# Flags:
#   -DryRun      Preview changes without installing
#   -Force       Overwrite existing agents
#   -Dev         Install development dependencies
#   -Provider    Set model provider (opencode, claude, gpt, gemini, custom)
#   -Help        Show this help message

param(
    [switch]$DryRun,
    [switch]$Force,
    [switch]$Dev,
    [string]$Provider = "",
    [switch]$Help
)

$ErrorActionPreference = "Stop"

# ============================================
# CONFIGURATION
# ============================================

$REPO = "Ulysses-Alv/Unity-Agent-Expert"
$INSTALL_DIR = "$env:USERPROFILE\.config\opencode"
$AGENTS_SRC = "opencode\config\agents.json"
$PRESETS_SRC = "opencode\config\presets.json"
$PROMPTS_SRC = "opencode\prompts\unity"
$SKILLS_SRC = "skills"

# Colors
function Write-Banner {
    Write-Host ""
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host " Unity-Agent-Expert Installer" -ForegroundColor Cyan
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host ""
}

function Write-Help {
    Write-Host @"
Usage: $PSCommandPath [OPTIONS]

One-liner: irm https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/install.ps1 | iex

Options:
  -DryRun      Preview changes without installing
  -Force       Overwrite existing agents (default: skip)
  -Dev         Install development dependencies
  -Provider    Set model provider (opencode, claude, gpt, gemini, custom)
  -Help        Show this help message

Examples:
  $PSCommandPath                 # Interactive install
  $PSCommandPath -DryRun          # Preview what would be installed
  $PSCommandPath -Provider claude # Install with Claude models
  $PSCommandPath -Force -Dev      # Force reinstall with dev deps

"@
}

function Get-RepoDir {
    # If run via pipe from URL, $PSCommandPath will be temp location
    # In that case, we need to clone or extract the repo
    $scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
    $tempPath = $scriptDir

    # Check if we're in a valid repo
    $isValidRepo = $false
    if ($scriptDir -and (Test-Path $scriptDir)) {
        $gitDir = Join-Path $scriptDir ".git"
        $agentsFile = Join-Path $scriptDir $AGENTS_SRC
        if ((Test-Path $gitDir) -or (Test-Path $agentsFile)) {
            $isValidRepo = $true
        }
    }

    if (-not $isValidRepo) {
        # Clone to temp directory
        $tmpDir = Join-Path $env:TEMP "unity-agent-expert-install-$PID"
        Write-Host "[INFO] Cloning repository..." -ForegroundColor Cyan

        try {
            Git clone --depth 1 "https://github.com/$REPO.git" $tmpDir 2>&1 | Out-Null
            return $tmpDir
        }
        catch {
            Write-Host "[ERROR] Failed to clone repository: $_" -ForegroundColor Red
            Write-Host "[ERROR] git is required but not found or clone failed" -ForegroundColor Red
            exit 1
        }
    }

    return $scriptDir
}

function Test-Dependencies {
    Write-Host "[STEP] Checking dependencies..." -ForegroundColor Blue

    $missing = @()

    # Check git
    try {
        $null = Get-Command git -ErrorAction SilentlyContinue
        Write-Host "[OK] git found" -ForegroundColor Green
    }
    catch {
        $missing += "git"
    }

    # Check PowerShell version
    if ($PSVersionTable.PSVersion.Major -lt 5) {
        Write-Host "[WARN] PowerShell 5.1+ recommended, found $($PSVersionTable.PSVersion)" -ForegroundColor Yellow
    }

    if ($missing.Count -gt 0) {
        Write-Host "[ERROR] Missing dependencies: $($missing -join ', ')" -ForegroundColor Red
        exit 1
    }
}

function Test-Prerequisites {
    param([string]$RepoDir)

    Write-Host "[STEP] Checking prerequisites..." -ForegroundColor Blue

    # Check opencode.json
    $opencodeConfig = Join-Path $INSTALL_DIR "opencode.json"
    if (-not (Test-Path $opencodeConfig)) {
        Write-Host "[ERROR] OpenCode config not found at: $opencodeConfig" -ForegroundColor Red
        Write-Host "Please ensure OpenCode is installed and configured." -ForegroundColor Yellow
        Write-Host "Visit: https://opencode.ai" -ForegroundColor Cyan
        exit 1
    }
    Write-Host "[OK] Found opencode.json" -ForegroundColor Green

    # Check agents.json
    $agentsFile = Join-Path $RepoDir $AGENTS_SRC
    if (-not (Test-Path $agentsFile)) {
        Write-Host "[ERROR] agents.json not found at: $agentsFile" -ForegroundColor Red
        exit 1
    }
    Write-Host "[OK] Found agents.json" -ForegroundColor Green

    # Check presets.json
    $presetsFile = Join-Path $RepoDir $PRESETS_SRC
    if (-not (Test-Path $presetsFile)) {
        Write-Host "[ERROR] presets.json not found at: $presetsFile" -ForegroundColor Red
        exit 1
    }
    Write-Host "[OK] Found presets.json" -ForegroundColor Green

    # Check prompts
    $promptsDir = Join-Path $RepoDir $PROMPTS_SRC
    if (-not (Test-Path $promptsDir)) {
        Write-Host "[ERROR] Prompts folder not found at: $promptsDir" -ForegroundColor Red
        exit 1
    }
    $promptCount = (Get-ChildItem $promptsDir -Filter "*.md").Count
    Write-Host "[OK] Found $promptCount prompt files" -ForegroundColor Green

    # Check skills (optional)
    $skillsDir = Join-Path $RepoDir $SKILLS_SRC
    if (Test-Path $skillsDir) {
        $skillCount = (Get-ChildItem $skillsDir -Directory -Filter "unity-*").Count
        Write-Host "[OK] Found $skillCount skill folders" -ForegroundColor Green
    }
    else {
        Write-Host "[WARN] Skills folder not found (optional)" -ForegroundColor Yellow
    }

    # Validate JSON
    try {
        $null = Get-Content $agentsFile -Raw | ConvertFrom-Json
        $null = Get-Content $presetsFile -Raw | ConvertFrom-Json
        Write-Host "[OK] JSON files are valid" -ForegroundColor Green
    }
    catch {
        Write-Host "[ERROR] Invalid JSON in config files: $_" -ForegroundColor Red
        exit 1
    }
}

function Install-Agents {
    param(
        [string]$RepoDir,
        [string]$ProviderChoice
    )

    Write-Host ""
    Write-Host "Installing Unity agents..." -ForegroundColor Yellow

    $configFile = Join-Path $INSTALL_DIR "opencode.json"
    $agentsFile = Join-Path $RepoDir $AGENTS_SRC
    $presetsFile = Join-Path $RepoDir $PRESETS_SRC

    # Read configs
    $config = Get-Content $configFile -Raw | ConvertFrom-Json
    $agentsToAdd = Get-Content $agentsFile -Raw | ConvertFrom-Json
    $presets = Get-Content $presetsFile -Raw | ConvertFrom-Json

    # Get existing Unity agents
    $existingUnityAgents = @()
    if ($config.PSObject.Properties.Name -contains "agent") {
        $existingUnityAgents = $config.agent.PSObject.Properties.Name | Where-Object { $_ -like "unity-*" }
    }

    # Check for existing installation
    if ($existingUnityAgents.Count -gt 0 -and -not $Force) {
        Write-Host "[INFO] Unity agents already installed ($($existingUnityAgents.Count) found)" -ForegroundColor Cyan
        Write-Host "       Use -Force to reinstall" -ForegroundColor Cyan
        return
    }

    if ($DryRun) {
        Write-Host "  [DRY RUN] Would install $($agentsToAdd.PSObject.Properties.Name.Count) agent(s)" -ForegroundColor Cyan
        return
    }

    # Create backup
    $backupPath = "$configFile.backup.$(Get-Date -Format 'yyyyMMddHHmmss')"
    Copy-Item $configFile -Destination $backupPath
    Write-Host "[BACKUP] Created: $backupPath" -ForegroundColor Green

    # Add agents
    $added = 0
    foreach ($agentName in $agentsToAdd.PSObject.Properties.Name) {
        $agentEntry = $agentsToAdd.$agentName | Select-Object *

        # Get model from preset if available
        $presetModel = $null
        if ($ProviderChoice -ne "" -and $presets.PSObject.Properties.Name -contains $ProviderChoice) {
            $providerPreset = $presets.$ProviderChoice
            if ($providerPreset.models -and $providerPreset.models.PSObject.Properties.Name -contains $agentName) {
                $presetModel = $providerPreset.models.$agentName
            }
        }

        if ($presetModel) {
            $agentEntry.model = $presetModel
        }

        $config.agent | Add-Member -NotePropertyName $agentName -NotePropertyValue $agentEntry -Force
        $added++
    }

    # Save
    $jsonOutput = $config | ConvertTo-Json -Depth 20
    [System.IO.File]::WriteAllText($configFile, $jsonOutput, (New-Object System.Text.UTF8Encoding $false))

    Write-Host "[OK] Added $added agent(s)" -ForegroundColor Green
}

function Install-Prompts {
    param([string]$RepoDir)

    Write-Host ""
    Write-Host "Installing Unity prompts..." -ForegroundColor Yellow

    $promptsSource = Join-Path $RepoDir $PROMPTS_SRC
    $promptsTarget = Join-Path $INSTALL_DIR "prompts\unity"

    if ($DryRun) {
        $count = (Get-ChildItem $promptsSource -Filter "*.md").Count
        Write-Host "  [DRY RUN] Would copy $count prompt(s) to: $promptsTarget" -ForegroundColor Cyan
        return
    }

    # Create prompts directory
    if (-not (Test-Path $promptsTarget)) {
        New-Item -ItemType Directory -Path $promptsTarget -Force | Out-Null
    }

    # Copy prompts
    $copied = 0
    Get-ChildItem $promptsSource -Filter "*.md" | ForEach-Object {
        $targetFile = Join-Path $promptsTarget $_.Name
        Copy-Item $_.FullName -Destination $targetFile -Force
        $copied++
    }

    Write-Host "[OK] Copied $copied prompt(s) to: $promptsTarget" -ForegroundColor Green
}

function Install-Skills {
    param([string]$RepoDir)

    Write-Host ""
    Write-Host "Installing Unity skills..." -ForegroundColor Yellow

    $skillsSource = Join-Path $RepoDir $SKILLS_SRC
    $skillsTarget = Join-Path $INSTALL_DIR "skills\unity-6000"

    if (-not (Test-Path $skillsSource)) {
        Write-Host "[SKIP] Skills folder not found" -ForegroundColor Yellow
        return
    }

    if ($DryRun) {
        $count = (Get-ChildItem $skillsSource -Directory -Filter "unity-*").Count
        Write-Host "  [DRY RUN] Would copy $count skill(s) to: $skillsTarget" -ForegroundColor Cyan
        return
    }

    # Create target directory
    if (-not (Test-Path $skillsTarget)) {
        New-Item -ItemType Directory -Path $skillsTarget -Force | Out-Null
    }

    # Copy skills
    $copied = 0
    Get-ChildItem $skillsSource -Directory -Filter "unity-*" | ForEach-Object {
        $targetDir = Join-Path $skillsTarget $_.Name
        if (-not (Test-Path $targetDir)) {
            New-Item -ItemType Directory -Path $targetDir -Force | Out-Null
        }
        Copy-Item "$($_.FullName)\*" -Destination $targetDir -Recurse -Force
        $copied++
    }

    Write-Host "[OK] Copied $copied skill(s) to: $skillsTarget" -ForegroundColor Green
}

function Show-Status {
    Write-Host ""
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host " Installation Complete!" -ForegroundColor Green
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "Installed Components:" -ForegroundColor White
    Write-Host "  - unity-6000-expert (orchestrator)"
    Write-Host "  - 17 sub-agents for Unity development"
    $promptCount = (Get-ChildItem "$INSTALL_DIR\prompts\unity" -Filter "*.md" -ErrorAction SilentlyContinue | Measure-Object).Count
    $skillCount = (Get-ChildItem "$INSTALL_DIR\skills\unity-6000" -Directory -Filter "unity-*" -ErrorAction SilentlyContinue | Measure-Object).Count
    Write-Host "  - $promptCount prompts"
    Write-Host "  - $skillCount skills"
    Write-Host ""
    Write-Host "Locations:" -ForegroundColor White
    Write-Host "  Config: $INSTALL_DIR\opencode.json"
    Write-Host "  Prompts: $INSTALL_DIR\prompts\unity"
    Write-Host "  Skills: $INSTALL_DIR\skills\unity-6000"
    Write-Host ""
    Write-Host "Next Steps:" -ForegroundColor White
    Write-Host "  1. Restart OpenCode or reload config"
    Write-Host "  2. Type /agent unity-6000-expert to start"
    Write-Host ""
    Write-Host "To update after repo pull: re-run this script" -ForegroundColor Gray
    Write-Host "For preview: $PSCommandPath -DryRun" -ForegroundColor Gray
}

# ============================================
# MAIN
# ============================================

function Main {
    if ($Help) {
        Write-Banner
        Write-Help
        return
    }

    Write-Banner

    $os = $PSVersionTable.OS
    Write-Host "[INFO] OS: $os" -ForegroundColor Cyan
    Write-Host "[INFO] PowerShell: $($PSVersionTable.PSVersion)" -ForegroundColor Cyan
    Write-Host "[INFO] Install dir: $INSTALL_DIR" -ForegroundColor Cyan

    Test-Dependencies
    $repoDir = Get-RepoDir
    Write-Host "[INFO] Repository: $repoDir" -ForegroundColor Cyan

    Test-Prerequisites -RepoDir $repoDir

    Install-Agents -RepoDir $repoDir -ProviderChoice $Provider
    Install-Prompts -RepoDir $repoDir
    Install-Skills -RepoDir $repoDir

    Show-Status
}

Main