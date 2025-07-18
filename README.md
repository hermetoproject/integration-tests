# Incorrect RPM size test

This scenario validates that Hermeto will refuse to proceed the prefetch of RPMs when an incorrect
file size is found.

The test steps are as follow:
- prefetch the dependencies
- validate that the prefetch fails with the appropriate error
