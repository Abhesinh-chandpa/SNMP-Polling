# Linux installation

Go is **not required** on the target Linux machine when using a compiled release binary.
Go is required to build from source.

## Install from CI artifact (development builds)
1. Open the GitHub Actions **Linux binary builds** workflow.
2. Download the artifact matching your Linux CPU architecture (amd64 or arm64).
3. Verify its SHA256 checksum and extract it.
4. Install the executable:

```sh
sudo install -m 0755 nmp-snmp-linux-amd64 /usr/local/bin/nmp-snmp
nmp-snmp version
nmp-snmp help
```

For ARM64, replace `amd64` with `arm64`. These are development builds; verify checksums against the artifact's SHA256SUMS before installing.

## Build from source

```sh
git clone https://github.com/Abhesinh-chandpa/SNMP-Polling.git
cd SNMP-Polling
go build -o nmp-snmp ./cmd/nmp-snmp
./nmp-snmp version
```

## Roadmap
A dedicated unprivileged systemd service, config directory and package installer will be delivered after continuous polling is implemented. The Batch 0 binary currently supports only help/version; SNMP commands are being implemented in Batch 1.
