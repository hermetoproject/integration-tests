# Missing hashes test (pylock)

This test covers a scenario where `pylock.toml` is missing hashes. Hermeto should still
successfully prefetch all packages and report missing hashes (checksums) in the SBOM.
