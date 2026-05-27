# Directory dependency test (pylock)

This test checks that an attempt to fetch a package with a directory
dependency in `pylock.toml` will result in failure. Hermeto does not
support directory dependencies because they cannot be fetched or verified.
