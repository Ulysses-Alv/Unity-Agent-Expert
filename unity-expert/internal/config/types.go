package config

import "time"

// OpenCodeConfig represents the structure of opencode.json.
type OpenCodeConfig struct {
	Agents map[string]AgentConfig `json:"agent,omitempty"`
}

// AgentConfig represents an agent configuration in opencode.json.
type AgentConfig struct {
	Description string                 `json:"description,omitempty"`
	Mode        string                 `json:"mode,omitempty"`
	Model       string                 `json:"model,omitempty"`
	Prompt      string                 `json:"prompt,omitempty"`
	Tools       map[string]bool        `json:"tools,omitempty"`
	Hidden      bool                   `json:"hidden,omitempty"`
	Permission  *PermissionConfig      `json:"permission,omitempty"`
	Raw         map[string]interface{} `json:"-"` // Preserves unknown fields
}

// PermissionConfig represents permission settings for an agent.
type PermissionConfig struct {
	Delegate map[string]string `json:"delegate,omitempty"`
	Task     map[string]string `json:"task,omitempty"`
}

// Preset represents a provider preset from presets.json.
type Preset struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Models      map[string]string `json:"models"`
}

// State represents the installation state persisted to state.json.
type State struct {
	Version          int               `json:"version"`
	InstalledAt      string            `json:"installedAt"`
	UpdatedAt        string            `json:"updatedAt"`
	InstalledAgents  []string          `json:"installedAgents"`
	Provider         string            `json:"provider"`
	ModelAssignments map[string]string `json:"modelAssignments"`
}

// InstallState represents the runtime installation state.
type InstallState struct {
	Provider        string    `json:"provider"`
	InstalledAt     time.Time `json:"installedAt"`
	InstalledAgents []string  `json:"installedAgents"`
	Version         string    `json:"version"`
}
