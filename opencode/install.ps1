# Unity-Expert OpenCode Installer
# Run this after cloning the repo to install Unity agents into your OpenCode config
# Usage: .\install.ps1 [-DryRun] [-Uninstall] [-Provider <provider>]

param(
    [switch]$DryRun,
    [switch]$Uninstall,
    [string]$Provider = ""
)

$ErrorActionPreference = "Stop"

# Paths
$OPENCODE_CONFIG = "$env:USERPROFILE\.config\opencode\opencode.json"
$PROMPTS_SOURCE = Join-Path $PSScriptRoot "prompts\unity"
$PROMPTS_TARGET = "$env:USERPROFILE\.config\opencode\prompts\unity"
$SKILLS_SOURCE = Join-Path $PSScriptRoot "..\skills"
$SKILLS_TARGET = "$env:USERPROFILE\.config\opencode\skills\unity-6000"
$AGENTS_CONFIG = Join-Path $PSScriptRoot "config\agents.json"
$PRESETS_CONFIG = Join-Path $PSScriptRoot "config\presets.json"

function Write-Banner {
    Write-Host ""
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host " Unity-Expert OpenCode Installer" -ForegroundColor Cyan
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host ""
}

function Test-Prerequisites {
    # Check opencode.json exists
    if (-not (Test-Path $OPENCODE_CONFIG)) {
        Write-Host "[ERROR] opencode.json not found at: $OPENCODE_CONFIG" -ForegroundColor Red
        Write-Host "Please ensure OpenCode is installed and configured." -ForegroundColor Yellow
        exit 1
    }
    Write-Host "[OK] Found opencode.json" -ForegroundColor Green

    # Check agents.json exists
    if (-not (Test-Path $AGENTS_CONFIG)) {
        Write-Host "[ERROR] agents.json not found at: $AGENTS_CONFIG" -ForegroundColor Red
        exit 1
    }
    Write-Host "[OK] Found agents.json" -ForegroundColor Green

    # Check presets.json exists
    if (-not (Test-Path $PRESETS_CONFIG)) {
        Write-Host "[ERROR] presets.json not found at: $PRESETS_CONFIG" -ForegroundColor Red
        exit 1
    }
    Write-Host "[OK] Found presets.json" -ForegroundColor Green

    # Check prompts source exists
    if (-not (Test-Path $PROMPTS_SOURCE)) {
        Write-Host "[ERROR] Prompts folder not found at: $PROMPTS_SOURCE" -ForegroundColor Red
        exit 1
    }
    $promptCount = (Get-ChildItem $PROMPTS_SOURCE -Filter "*.md").Count
    Write-Host "[OK] Found $promptCount prompt files" -ForegroundColor Green

    # Check skills source exists (optional warning)
    if (-not (Test-Path $SKILLS_SOURCE)) {
        Write-Host "[WARN] Skills folder not found at: $SKILLS_SOURCE" -ForegroundColor Yellow
        Write-Host "       Skills will not be installed." -ForegroundColor Yellow
    } else {
        $skillCount = (Get-ChildItem $SKILLS_SOURCE -Directory -Filter "unity-*").Count
        Write-Host "[OK] Found $skillCount skill folders" -ForegroundColor Green
    }

    # Validate agents.json is valid JSON
    try {
        $null = Get-Content $AGENTS_CONFIG -Raw | ConvertFrom-Json
        Write-Host "[OK] agents.json is valid JSON" -ForegroundColor Green
    } catch {
        Write-Host "[ERROR] agents.json is not valid JSON: $_" -ForegroundColor Red
        exit 1
    }

    # Validate presets.json is valid JSON
    try {
        $null = Get-Content $PRESETS_CONFIG -Raw | ConvertFrom-Json
        Write-Host "[OK] presets.json is valid JSON" -ForegroundColor Green
    } catch {
        Write-Host "[ERROR] presets.json is not valid JSON: $_" -ForegroundColor Red
        exit 1
    }
}

function Select-Provider {
    param(
        [string]$RequestedProvider
    )

    $providers = @(
        @{ Name = "opencode"; Label = "OpenCode Models"; Default = $true }
        @{ Name = "claude"; Label = "Claude Models"; Default = $false }
        @{ Name = "gpt"; Label = "GPT Models"; Default = $false }
        @{ Name = "gemini"; Label = "Gemini Models"; Default = $false }
        @{ Name = "custom"; Label = "User-defined models"; Default = $false }
    )

    $selectedProvider = ""

    if ($RequestedProvider -ne "") {
        # Validate requested provider
        $validNames = $providers | ForEach-Object { $_.Name }
        if ($RequestedProvider -in $validNames) {
            $selectedProvider = $RequestedProvider
        } else {
            Write-Host "[WARN] Provider '$RequestedProvider' not recognized." -ForegroundColor Yellow
            Write-Host "       Falling back to 'opencode'." -ForegroundColor Yellow
            $selectedProvider = "opencode"
        }
    } else {
        # Show provider selection menu
        Write-Host "Select model provider:" -ForegroundColor Yellow
        $i = 1
        foreach ($p in $providers) {
            $defaultMark = ""
            if ($p.Default) { $defaultMark = " (default)" }
            Write-Host "  [$i] $($p.Name) - $($p.Label)$defaultMark" -ForegroundColor White
            $i++
        }
        Write-Host ""

        $choice = Read-Host "Enter selection (1-5) or provider name"
        $choice = $choice.Trim().ToLower()

        # Try to match by number
        $number = 0
        if ([int]::TryParse($choice, [ref]$number)) {
            if ($number -ge 1 -and $number -le $providers.Count) {
                $selectedProvider = $providers[$number - 1].Name
            } else {
                Write-Host "[WARN] Invalid selection. Using 'opencode'." -ForegroundColor Yellow
                $selectedProvider = "opencode"
            }
        } else {
            # Try to match by name
            $validNames = $providers | ForEach-Object { $_.Name }
            if ($choice -in $validNames) {
                $selectedProvider = $choice
            } else {
                Write-Host "[WARN] Provider '$choice' not recognized." -ForegroundColor Yellow
                Write-Host "       Falling back to 'opencode'." -ForegroundColor Yellow
                $selectedProvider = "opencode"
            }
        }
    }

    return $selectedProvider
}

function Get-ModelAssignments {
    param(
        [string]$Provider,
        [PSCustomObject]$AgentsToAdd
    )

    $presets = Get-Content $PRESETS_CONFIG -Raw | ConvertFrom-Json
    $assignments = @{}

    # Check if provider exists in presets
    if ($presets.PSObject.Properties.Name -notcontains $Provider) {
        Write-Host "[WARN] Provider '$Provider' not found in presets.json." -ForegroundColor Yellow
        Write-Host "       Falling back to 'opencode' preset." -ForegroundColor Yellow
        $Provider = "opencode"
    }

    # Get the provider preset
    $providerPreset = $presets.$Provider

    # Check if it has models property (new nested structure)
    $modelLookup = $null
    if ($providerPreset.PSObject.Properties.Name -contains "models") {
        $modelLookup = $providerPreset.models
    } else {
        # Fallback to old flat structure (direct property access)
        $modelLookup = $providerPreset
    }

    # Warn for custom provider
    if ($Provider -eq "custom") {
        Write-Host "[WARN] Using 'custom' provider." -ForegroundColor Yellow
        Write-Host "       Models must be configured manually in presets.json under 'custom.models'." -ForegroundColor Yellow
        Write-Host "       Agents without model assignments will use agents.json defaults." -ForegroundColor Yellow
    }

    Write-Host "[PRESET] Using provider: $Provider" -ForegroundColor Cyan

    foreach ($agentName in $AgentsToAdd.PSObject.Properties.Name) {
        $modelFromPreset = $null
        if ($modelLookup.PSObject.Properties.Name -contains $agentName) {
            $modelFromPreset = $modelLookup.$agentName
        }

        $assignments[$agentName] = $modelFromPreset
    }

    return $assignments
}

function Install-Agents {
    param(
        [string]$Provider
    )

    Write-Host ""
    Write-Host "Installing Unity agents..." -ForegroundColor Yellow

    # Read current config
    $configContent = Get-Content $OPENCODE_CONFIG -Raw
    $config = $configContent | ConvertFrom-Json

    # Read agents to add
    $agentsToAdd = Get-Content $AGENTS_CONFIG -Raw | ConvertFrom-Json

    # Get model assignments from presets
    $modelAssignments = Get-ModelAssignments -Provider $Provider -AgentsToAdd $agentsToAdd

    # Check for existing Unity agents (don't duplicate)
    $existingUnityAgents = @()
    if ($config.PSObject.Properties.Name -contains "agent") {
        $existingUnityAgents = $config.agent.PSObject.Properties.Name | Where-Object { $_ -like "unity-*" }
    }

    # Add missing agents
    $added = 0
    $skipped = 0
    foreach ($agentName in $agentsToAdd.PSObject.Properties.Name) {
        if ($agentName -in $existingUnityAgents) {
            Write-Host "  [SKIP] $agentName (already exists)" -ForegroundColor Gray
            $skipped++
            continue
        }

        # Build agent entry from agents.json
        $agentEntry = $agentsToAdd.$agentName | Select-Object *

        # Override model from preset if available
        $presetModel = $modelAssignments[$agentName]
        if ($presetModel -and $presetModel -ne "") {
            $agentEntry.model = $presetModel
            Write-Host "  [ADD] $agentName (model: $presetModel)" -ForegroundColor Green
        } else {
            # Fall back to agents.json default model
            $defaultModel = $agentsToAdd.$agentName.model
            Write-Host "  [ADD] $agentName (model: $defaultModel [preset default])" -ForegroundColor Green
        }

        $config.agent | Add-Member -NotePropertyName $agentName -NotePropertyValue $agentEntry -Force
        $added++
    }

    if ($added -eq 0) {
        Write-Host "  All Unity agents already installed. Nothing to do." -ForegroundColor Yellow
        return
    }

    if ($DryRun) {
        Write-Host "  [DRY RUN] Would add $added agent(s)" -ForegroundColor Cyan
        return
    }

    # Create backup before modifying
    $backupPath = "$OPENCODE_CONFIG.backup.$(Get-Date -Format 'yyyyMMddHHmmss')"
    Copy-Item $OPENCODE_CONFIG -Destination $backupPath
    Write-Host "  [BACKUP] Created: $backupPath" -ForegroundColor Green

    # Save updated config with UTF8 without BOM
    $jsonOutput = $config | ConvertTo-Json -Depth 20
    [System.IO.File]::WriteAllText($OPENCODE_CONFIG, $jsonOutput, (New-Object System.Text.UTF8Encoding $false))

    # Validate the written JSON is valid
    try {
        $null = Get-Content $OPENCODE_CONFIG -Raw | ConvertFrom-Json
        Write-Host "  Added $added agent(s), skipped $skipped" -ForegroundColor Green
        Write-Host "[OK] Agents installed" -ForegroundColor Green
    } catch {
        Write-Host "[ERROR] Written config is invalid JSON! Restoring backup..." -ForegroundColor Red
        Copy-Item $backupPath -Destination $OPENCODE_CONFIG -Force
        exit 1
    }
}

function Install-Prompts {
    Write-Host ""
    Write-Host "Installing Unity prompts..." -ForegroundColor Yellow

    if ($DryRun) {
        $count = (Get-ChildItem $PROMPTS_SOURCE -Filter "*.md").Count
        Write-Host "  [DRY RUN] Would copy $count prompt(s) to: $PROMPTS_TARGET" -ForegroundColor Cyan
        return
    }

    # Create prompts directory if needed
    if (-not (Test-Path $PROMPTS_TARGET)) {
        New-Item -ItemType Directory -Path $PROMPTS_TARGET -Force | Out-Null
    }

    # Copy prompts
    $copied = 0
    Get-ChildItem $PROMPTS_SOURCE -Filter "*.md" | ForEach-Object {
        $targetFile = Join-Path $PROMPTS_TARGET $_.Name
        Copy-Item $_.FullName -Destination $targetFile -Force
        $copied++
    }

    Write-Host "  Copied $copied prompt file(s)"
    Write-Host "[OK] Prompts installed at: $PROMPTS_TARGET" -ForegroundColor Green
}

function Install-Skills {
    Write-Host ""
    Write-Host "Installing Unity skills..." -ForegroundColor Yellow

    # Check if skills exist in repo
    if (-not (Test-Path $SKILLS_SOURCE)) {
        Write-Host "  [SKIP] Skills folder not found at: $SKILLS_SOURCE" -ForegroundColor Yellow
        return
    }

    if ($DryRun) {
        $count = (Get-ChildItem $SKILLS_SOURCE -Directory -Filter "unity-*").Count
        Write-Host "  [DRY RUN] Would copy $count skill(s) to: $SKILLS_TARGET" -ForegroundColor Cyan
        return
    }

    # Create target directory
    if (-not (Test-Path $SKILLS_TARGET)) {
        New-Item -ItemType Directory -Path $SKILLS_TARGET -Force | Out-Null
    }

    # Copy skills
    $copied = 0
    Get-ChildItem $SKILLS_SOURCE -Directory | Where-Object { $_.Name -like "unity-*" } | ForEach-Object {
        $targetDir = Join-Path $SKILLS_TARGET $_.Name
        if (-not (Test-Path $targetDir)) {
            New-Item -ItemType Directory -Path $targetDir -Force | Out-Null
        }
        Copy-Item "$($_.FullName)\*" -Destination $targetDir -Recurse -Force
        $copied++
    }

    Write-Host "  Copied $copied skill(s)"
    Write-Host "[OK] Skills installed at: $SKILLS_TARGET" -ForegroundColor Green
}

function Show-Status {
    param(
        [string]$Provider
    )

    Write-Host ""
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host " Installation Complete!" -ForegroundColor Green
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "Installed Unity agents: unity-6000-expert (orchestrator) + 16 sub-agents" -ForegroundColor White
    Write-Host "Model provider: $Provider" -ForegroundColor White
    Write-Host "Prompts location: $PROMPTS_TARGET" -ForegroundColor White
    Write-Host "Skills location: $SKILLS_TARGET" -ForegroundColor White
    Write-Host ""
    Write-Host "To use Unity Expert, type: /agent unity-6000-expert" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "To update after repo pull: re-run this script" -ForegroundColor Gray
    Write-Host "To preview changes: .\install.ps1 -DryRun" -ForegroundColor Gray
    Write-Host "To use different provider: .\install.ps1 -Provider <provider>" -ForegroundColor Gray
}

# Main
Write-Banner

if ($Uninstall) {
    Write-Host "Uninstall not implemented yet. Remove agents manually from opencode.json" -ForegroundColor Yellow
    Write-Host "Backup location: $OPENCODE_CONFIG.backup.*" -ForegroundColor Gray
    exit 0
}

Test-Prerequisites

# Select provider
$selectedProvider = Select-Provider -RequestedProvider $Provider

Install-Agents -Provider $selectedProvider
Install-Prompts
Install-Skills
Show-Status -Provider $selectedProvider
