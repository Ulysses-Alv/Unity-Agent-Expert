param([string]$ToolPath = "$env:LOCALAPPDATA\unity-agent-expert\bin", [string]$Version = "latest")
$BINARY = "unity-agent-expert.exe"
$REPO = "Ulysses-Alv/Unity-Agent-Expert"

$apiUrl = if ($Version -eq "latest") { "https://api.github.com/repos/$REPO/releases/latest" } else { "https://api.github.com/repos/$REPO/releases/tags/$Version" }
$asset = (Invoke-RestMethod $apiUrl -Headers @{Accept = "application/vnd.github.v3+json"}).assets | Where-Object { $_.name -eq "unity-agent-expert-windows-amd64.exe" }
if (-not $asset) { Write-Host "[ERROR] Asset not found"; exit 1 }

$tmpFile = "$env:TEMP\$BINARY"
Invoke-WebRequest -Uri $asset.browser_download_url -OutFile $tmpFile
if (-not (Test-Path $ToolPath)) { New-Item -ItemType Directory -Path $ToolPath -Force | Out-Null }
Move-Item $tmpFile "$ToolPath\$BINARY" -Force

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host " unity-agent-expert installed!" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "Location: $ToolPath\$BINARY" -ForegroundColor White
Write-Host ""
Write-Host "To use, add to your PATH:" -ForegroundColor Yellow
Write-Host "  [Environment]::SetEnvironmentVariable('PATH', \$env:PATH + ';$ToolPath', 'User')" -ForegroundColor Cyan
Write-Host ""
Write-Host "Then restart PowerShell and run:" -ForegroundColor Yellow
Write-Host "  unity-agent-expert --help" -ForegroundColor Cyan
