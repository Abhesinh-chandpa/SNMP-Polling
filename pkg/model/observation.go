// Package model defines protocol-independent telemetry contracts for future NMP integration.
package model

import "time"

// Observation is a single measured value. RawValue preserves large SNMP
// counters exactly as decimal strings rather than lossy JSON floating points.
type Observation struct {
 SchemaVersion string    `json:"schema_version"`
 DeviceID      string    `json:"device_id"`
 CollectorID   string    `json:"collector_id"`
 Timestamp     time.Time `json:"timestamp"`
 Protocol      string    `json:"protocol"`
 OID           string    `json:"oid,omitempty"`
 Metric        string    `json:"metric"`
 InterfaceIndex *uint32  `json:"interface_index,omitempty"`
 ValueType     string    `json:"value_type"`
 RawValue      string    `json:"value"`
 Unit          string    `json:"unit,omitempty"`
 Quality       string    `json:"quality"`
}
