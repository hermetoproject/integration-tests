# Yarnberry Missing Yarn lockfile Integration Test

Its purpose is to verify that cachi2 will correctly fail a request.

There is a different lockfile specified in `.yarnrc.yml` with `lockfileFilename` variable. But that lockfile is missing in the repository, cachi2 will not process a request.
