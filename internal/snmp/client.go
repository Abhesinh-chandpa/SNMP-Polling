// Package snmp contains interfaces for a future GoSNMP adapter.
// This package intentionally has no protocol implementation in Batch 0.
package snmp

import "context"

type Variable struct {
 OID string
 Type string
 Value string
}

type Client interface {
 Get(ctx context.Context, oids []string) ([]Variable, error)
 Walk(ctx context.Context, rootOID string, visit func(Variable) error) error
 Close() error
}
