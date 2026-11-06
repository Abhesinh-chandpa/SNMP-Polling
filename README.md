# SNMP-Polling

Go-based SNMP tester and future NMP collector foundation.

> Development is underway. Batch 0 establishes architecture, contracts, CI and a starter CLI. SNMP polling and the interactive dashboard are planned for subsequent batches.

## Quick start

```sh
go run ./cmd/nmp-snmp version
go run ./cmd/nmp-snmp help
go test ./...
```

## Roadmap

- Batch 0: foundation and CI
- Batch 1: SNMP v2c/v3 protocol adapter
- Batch 2: device registry and configuration
- Batch 3: capability discovery
- Batch 4: scheduler and telemetry
- Batch 4.5: Bubble Tea interactive CLI dashboard
- Batch 5: storage and exports
- Batch 6: traps and diagnostics
- Batch 7: release and Linux packaging

See `docs/architecture.md` and `docs/development-plan.md`.

No credentials or device secrets should be committed.
