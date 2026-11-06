# SNMP protocol security

Batch 1 uses GoSNMP v1.47.0 or later compatible patched release. Use read-only accounts and SNMPv3 authPriv (SHA/SHA256 and AES) wherever possible.

Do not commit device passwords or community strings. Do not place credentials in command-line flags; future CLI commands should load protected credential profiles or use a secure prompt.

SNMPv2c has no encryption and is intended only for isolated, authorized lab networks.

Current Batch 1 limitations: no trap receiver, no cancellation of in-flight UDP reads (context checked before and between operations), no dynamic credential provider, no real device integration results yet.
