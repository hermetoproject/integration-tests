# Cachi2 RPMs incorrect size test

This scenario validates that Cachi2 will refuse to proceed the prefetch of RPMs when an incorrect file size is found. This test contains two RPMs and their respective source RPMs, and one RPM has an incorrect file size.

The test steps are as follow:
- prefetch the dependencies with Cachi2
- validate that the prefetch fails with the appropriate error
