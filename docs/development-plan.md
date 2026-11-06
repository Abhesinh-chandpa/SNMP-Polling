# Batch-based delivery plan

| Batch | Goal | Acceptance |
| --- | --- | --- |
| 0 | Repo, module, CLI, contracts, docs, CI | go test, go vet, go build pass |
| 1 | GoSNMP adapter, GET/WALK/BULK, v2c/v3 | mock and lab-device tests |
| 2 | Device registry, YAML config, credential references | safe add/list/test/remove |
| 3 | IF-MIB and system discovery | supported/unsupported capability report |
| 4 | Poll scheduler, counter rates, bounded workers | no overlapping polls, reset detection |
| 4.5 | Bubble Tea/Bubbles/Lip Gloss monitor | devices, interfaces, poll status; read-only snapshots |
| 5 | JSONL, CSV, SQLite and history | reliable exports, bounded retention |
| 6 | Traps, health and diagnostics | authenticated traps, event visibility |
| 7 | Linux systemd packaging and release | documented install, load testing |

## Working agreement
- Branch `batch/NN-topic` from `main`, use small feature branches as needed.
- All changes via pull requests, squash merge after review.
- Require CI and tests. Do not merge Batch 0 without a clean build.
- Protect `main` with repository ruleset when available.
- Keep hardware-dependent tests out of untrusted public PR runners.
- Assign one owner to shared interfaces and dependency files.

## Dependencies
Batch 0 uses only the Go standard library (Go 1.23+).
Batch 1 will pin a reviewed GoSNMP version.
Batch 4.5 will pin compatible Bubble Tea/Bubbles/Lip Gloss versions.
Avoid unused dependencies in the foundation commit.

## Initial CLI contract
`nmp-snmp help`, `nmp-snmp version` are implemented.
`get`, `walk`, `bulk`, `discover`, `poll`, `monitor`, `traps` and `stats` are future commands.
