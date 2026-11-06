// Package snapshot defines read-only access to collector state for the future
// Bubble Tea terminal dashboard and CLI exports.
package snapshot

import (
	"context"
	"time"
)

type DeviceStatus struct {
	ID        string
	Address   string
	Reachable bool
	LastPoll  time.Time
	LastError string
}

type State struct {
	CapturedAt time.Time
	Devices    []DeviceStatus
}

type Provider interface {
	Snapshot(ctx context.Context) (State, error)
}
