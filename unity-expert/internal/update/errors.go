package update

import "fmt"

// NetworkError is returned when a network request fails.
type NetworkError struct {
	Err error
}

func (e *NetworkError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("network error: %v", e.Err)
	}
	return "network error: no connection"
}

// RateLimitError is returned when the GitHub API rate limit is exceeded.
type RateLimitError struct{}

func (e *RateLimitError) Error() string {
	return "GitHub API rate limit alcanzado. Esperá unos minutos e intentá de nuevo."
}

// AssetNotFoundError is returned when no matching asset is found for the platform.
type AssetNotFoundError struct {
	OS   string
	Arch string
}

func (e *AssetNotFoundError) Error() string {
	return fmt.Sprintf("No hay binario disponible para %s/%s", e.OS, e.Arch)
}

// ReplaceError is returned when binary replacement fails.
type ReplaceError struct {
	Path string
	Err  error
}

func (e *ReplaceError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("No se pudo reemplazar el binario en %s: %v", e.Path, e.Err)
	}
	return fmt.Sprintf("No se pudo reemplazar el binario en %s", e.Path)
}

// DownloadError is returned when asset download fails.
type DownloadError struct {
	URL string
	Err error
}

func (e *DownloadError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("Error descargando la actualización desde %s: %v", e.URL, e.Err)
	}
	return fmt.Sprintf("Error descargando la actualización desde %s", e.URL)
}

// ExtractError is returned when zip extraction fails.
type ExtractError struct {
	Path string
	Err  error
}

func (e *ExtractError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("Error extrayendo el binario desde %s: %v", e.Path, e.Err)
	}
	return fmt.Sprintf("Error extrayendo el binario desde %s", e.Path)
}

// GentleAIError is returned when running in a gentle-ai managed environment.
type GentleAIError struct{}

func (e *GentleAIError) Error() string {
	return "Este entorno es gestionado por gentle-ai. Ejecutá:\n  irm https://... | iex\nUna vez instalado, ejecutá install de nuevo."
}

// VersionParseError is returned when version string parsing fails.
type VersionParseError struct {
	Version string
	Err     error
}

func (e *VersionParseError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("invalid version '%s': %v", e.Version, e.Err)
	}
	return fmt.Sprintf("invalid version '%s'", e.Version)
}
