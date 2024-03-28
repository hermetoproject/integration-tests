# Cachi2 RPMs missing checksums test

This scenario covers the reporting of missing checksums in the resulting SBOM. The test contains 2 RPMs and their corresponding source RPMs, and one of them does not have its checksum declared in the lockfile (neither for the RPM or the SRPM).

The test steps are as follow:
- prefetch the dependencies with Cachi2
- validate that the resulting SBOM specifies which files had missing checksums in the lockfile
