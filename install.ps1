param([string]$ToolPath = "$env:LOCALAPPDATA\unity-agent-expert\bin", [string]$Version = "latest")
$BINARY = "unity-agent-expert.exe"
$REPO = "Ulysses-Alv/Unity-Agent-Expert"

Write-Host "[INFO] Fetching release info..." -ForegroundColor Cyan
$apiUrl = if ($Version -eq "latest") { "https://api.github.com/repos/$REPO/releases/latest" } else { "https://api.github.com/repos/$REPO/releases/tags/$Version" }

try {
    $release = Invoke-RestMethod $apiUrl -Headers @{Accept = "application/vnd.github.v3+json"}
} catch {
    Write-Host "[ERROR] Failed to fetch release: $_" -ForegroundColor Red
    exit 1
}

Write-Host "[INFO] Found release: $($release.name)" -ForegroundColor Cyan
Write-Host "[INFO] Available assets:" -ForegroundColor Cyan
$release.assets | ForEach-Object { Write-Host "  - $($_.name)" -ForegroundColor Gray }

$asset = $release.assets | Where-Object { $_.name -eq "unity-agent-expert_1.0.0_windows_amd64.zip" }
if (-not $asset) {
    Write-Host "[ERROR] Asset 'unity-agent-expert_1.0.0_windows_amd64.zip' not found" -ForegroundColor Red
    Write-Host "[INFO] Available assets:" -ForegroundColor Yellow
    $release.assets | ForEach-Object { Write-Host "  - $($_.name)" -ForegroundColor Yellow }
    exit 1
}

Write-Host "[INFO] Downloading: $($asset.name)" -ForegroundColor Cyan
$tmpZip = "$env:TEMP\unity-agent-expert.zip"
$tmpDir = "$env:TEMP\unity-agent-expert-extract"

try {
    Invoke-WebRequest -Uri $asset.browser_download_url -OutFile $tmpZip
} catch {
    Write-Host "[ERROR] Failed to download: $_" -ForegroundColor Red
    exit 1
}

if (-not (Test-Path $ToolPath)) { New-Item -ItemType Directory -Path $ToolPath -Force | Out-Null }
if (Test-Path $tmpDir) { Remove-Item $tmpDir -Recurse -Force }
New-Item -ItemType Directory -Path $tmpDir -Force | Out-Null

Write-Host "[INFO] Extracting..." -ForegroundColor Cyan
Expand-Archive -Path $tmpZip -DestinationPath $tmpDir -Force

$extractedExe = Get-ChildItem $tmpDir -Filter "*.exe" -Recurse | Select-Object -First 1

if (-not $extractedExe) {
    Write-Host "[ERROR] Executable not found in archive" -ForegroundColor Red
    Write-Host "[DEBUG] Contents of extracted dir:" -ForegroundColor Gray
    Get-ChildItem $tmpDir -Recurse | ForEach-Object { Write-Host "  $($_.FullName)" -ForegroundColor Gray }
    exit 1
}

Write-Host "[INFO] Moving to: $ToolPath\$BINARY" -ForegroundColor Cyan
Move-Item $extractedExe.FullName "$ToolPath\$BINARY" -Force

# Cleanup
Remove-Item $tmpZip -Force -ErrorAction SilentlyContinue
Remove-Item $tmpDir -Recurse -Force -ErrorAction SilentlyContinue

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host " unity-agent-expert installed!" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "Location: $ToolPath\$BINARY" -ForegroundColor White
Write-Host ""
Write-Host "To use, add to your PATH:" -ForegroundColor Yellow
Write-Host "  [Environment]::SetEnvironmentVariable('PATH', `$env:PATH + ';$ToolPath', 'User')" -ForegroundColor Cyan
Write-Host ""
Write-Host "Then restart PowerShell and run:" -ForegroundColor Yellow
Write-Host "  unity-agent-expert --help" -ForegroundColor Cyan