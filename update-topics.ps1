$headers = @{
    'Authorization' = 'token $env:GITHUB_TOKEN'
    'Accept' = 'application/vnd.github.v3+json'
}
$body = @'
{"topics": ["unity", "unity3d", "unity-6000", "opencode", "ai-agents", "game-development", "csharp", "agent-systems", "spec-driven-development", "vibe-coding"]}
'@
Invoke-RestMethod -Uri 'https://api.github.com/repos/Ulysses-Alv/Unity-Agent-Expert' -Method PATCH -Headers $headers -ContentType 'application/json' -Body $body