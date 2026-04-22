package cli

import (
	"fmt"
)

// PreconditionError is returned when a precondition check fails.
type PreconditionError struct {
	Message string
	Hint    string
}

func (e *PreconditionError) Error() string {
	return e.Message
}

// BackupError is returned when backup creation fails.
type BackupError struct {
	Path string
	Err  error
}

func (e *BackupError) Error() string {
	return fmt.Sprintf("backup failed at %s: %v", e.Path, e.Err)
}

// ApplyError is returned when an apply operation fails.
type ApplyError struct {
	Operation string
	Path      string
	Err       error
}

func (e *ApplyError) Error() string {
	return fmt.Sprintf("%s failed for %s: %v", e.Operation, e.Path, e.Err)
}

// RollbackError is returned when rollback fails.
type RollbackError struct {
	BackupPath string
	Err       error
}

func (e *RollbackError) Error() string {
	return fmt.Sprintf("rollback failed from %s: %v", e.BackupPath, e.Err)
}
