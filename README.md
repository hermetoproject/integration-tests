# VCS dependency test (pylock)

This integration test covers a scenario with a git VCS dependency
alongside PyPI index packages in `pylock.toml`. Index packages
have no hashes to avoid the require-hashes conflict with VCS
dependencies, which cannot be hashed.

See [pylock.toml](pylock.toml).
