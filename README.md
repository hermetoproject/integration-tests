# Local path test (pylock)

This test checks that an attempt to fetch a package with a local path
source in `pylock.toml` will result in failure. Hermeto does not support
local path sources because they cannot be fetched or verified.
