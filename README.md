# Cachi2 RPMs invalid checksum test

This scenario validates that Cachi2 will refuse to proceed the prefetch of RPMs when a checksum mismatch is found. This test contains two RPMs and their respective source RPMs, and one RPM has an invalid checksum.

The test steps are as follow:
- prefetch the dependencies with Cachi2
- validate that the prefetch fails with the appropriate error
