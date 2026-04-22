param([string]$ToolPath = "$env:LOCALAPPDATA\unity-agent-expert\bin", [string]$Version = "latest")
$BINARY = "unity-agent-expert.exe"
$REPO = "Ulysses-Alv/Unity-Agent-Expert"

$apiUrl = if ($Version -eq "latest") { "https://api.github.com/repos/$REPO/releases/latest" } else { "https://api.github.com/repos/$REPO/releases/tags/$Version" }
$asset = (Invoke-RestMethod $apiUrl -Headers @{Accept = "application/vnd.github.v3+json"}).assets | Where-Object { $_.name -eq "unity-agent-expert_1.0.0_windows_amd64.zip" }
if (-not $asset) { Write-Host "[ERROR] Asset not found"; exit 1 }

$tmpZip = "$env:TEMP\unity-agent-expert.zip"
$tmpDir = "$env:TEMP\unity-agent-expert-extract"
Invoke-WebRequest -Uri $asset.browser_download_url -OutFile $tmpZip

if (-not (Test-Path $ToolPath)) { New-Item -ItemType Directory -Path $ToolPath -Force | Out-Null }
if (Test-Path $tmpDir) { Remove-Item $tmpDir -Recurse -Force }
New-Item -ItemType Directory -Path $tmpDir -Force | Out-Null

Expand-Archive -Path $tmpZip -DestinationPath $tmpDir -Force
$extractedExe = Get-ChildItem $tmpDir -Filter "*.exe" -Recurse | Select-Object -First 1

if (-not $extractedExe) {
    Write-Host "[ERROR] Executable not found in archive"; exit 1
}

Move-Item $extractedExe.FullName "$ToolPath\$BINARY" -Force
Remove-Item $tmpZip -Force
Remove-Item $tmpDir -Recurse -Force

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