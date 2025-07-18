# Missing RPM checksum test

This scenario covers the reporting of missing checksums for RPM dependencies in the resulting SBOM.
The test exercises an RPM that doesn't have checksum declared in the lockfile.

The test steps are as follow:
- prefetch the dependencies
- validate that the resulting SBOM specifies which files had missing checksums in the lockfile
