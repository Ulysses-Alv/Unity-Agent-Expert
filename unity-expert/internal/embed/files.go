package embed

import (
	"embed"
	"io/fs"
)

// ConfigFS is the embed.FS instance for configuration files.
//go:embed config
var ConfigFS embed.FS

// PromptsFS is the embed.FS instance for prompt files.
//go:embed prompts
var PromptsFS embed.FS

// SkillsFS is the embed.FS instance for skill files.
//go:embed skills
var SkillsFS embed.FS

// EmbeddedConfigFS returns the ConfigFS for external use.
func EmbeddedConfigFS() embed.FS {
	return ConfigFS
}

// EmbeddedPromptsFS returns the PromptsFS for external use.
func EmbeddedPromptsFS() embed.FS {
	return PromptsFS
}

// EmbeddedSkillsFS returns the SkillsFS for external use.
func EmbeddedSkillsFS() embed.FS {
	return SkillsFS
}

// ListPrompts lists all files in the prompts directory.
func ListPrompts() ([]fs.DirEntry, error) {
	return PromptsFS.ReadDir(".")
}

// ListSkills lists all directories in the skills directory.
func ListSkills() ([]fs.DirEntry, error) {
	return SkillsFS.ReadDir(".")
}

// EmbeddedAgentsJSON returns the embedded agents.json content.
func EmbeddedAgentsJSON() ([]byte, error) {
	return ConfigFS.ReadFile("config/agents.json")
}

// EmbeddedPresetsJSON returns the embedded presets.json content.
func EmbeddedPresetsJSON() ([]byte, error) {
	return ConfigFS.ReadFile("config/presets.json")
}
