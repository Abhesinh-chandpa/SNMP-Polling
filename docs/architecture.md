# Architecture — Batch 0

## Goals
A lightweight Go CLI and Linux SNMP collector, with independent monitoring UI and reusable NMP contracts.

## Layers
1. CLI/application bootstrap: `cmd/nmp-snmp`, `internal/app`.
2. SNMP transport adapter: `internal/snmp` (interface in Batch 0; GoSNMP implementation in Batch 1).
3. Device registry, discovery, scheduler and diagnostics: future `internal/` packages.
4. Snapshot provider: `internal/snapshot` — read-only data for the dashboard.
5. Stable telemetry contract: `pkg/model`.
6. Output sinks: JSONL, CSV, SQLite and future NMP transport.

## Terminal dashboard
Use Bubble Tea, Bubbles and Lip Gloss as dependencies when Batch 4.5 starts.
Do not embed btop or fork its C++ source. btop remains an optional external Linux host monitor.
The dashboard must not initiate duplicate polling; it reads snapshots from the collector, via a local Unix socket when running as a separate process.

## Security
Prefer SNMPv3 authPriv, use read-only credentials, never log secrets.
Avoid secrets in command-line arguments. Do not expose collector control to untrusted network clients.

## Interfaces
`internal/snmp.Client` isolates protocol transport for mocking and future implementation.
`internal/snapshot.Provider` isolates UI from polling.
`pkg/model.Observation` preserves large counters as decimal strings.

## Non-goals of Batch 0
No real SNMP traffic, device discovery, trap receiver, dashboard implementation or long-running service yet.
