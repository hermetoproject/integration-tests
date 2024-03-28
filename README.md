# Cachi2 RPMs E2E test

This scenario covers the prefetch and building of a Fedora 39 based container image that has `vim` installed from a set of RPMs. It provides two different sets of RPMs, one for x86_64 and the other for aarch64, as well as source RPMs.

The test steps are as follow:
- prefetch the dependencies with Cachi2
- validate the prefetched output and the resulting SBOM
- build an image hermetically with `vim` installed
