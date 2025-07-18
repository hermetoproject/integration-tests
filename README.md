# Invalid RPM checksum test

This scenario validates that Hermeto will refuse to proceed the prefetch of RPMs when a checksum
mismatch is found

The test steps are as follow:
- prefetch the dependencies with Cachi2
- validate that the prefetch fails with the appropriate error
