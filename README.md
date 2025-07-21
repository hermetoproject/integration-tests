# RPMs E2E test

This scenario covers the prefetch and building a container image that installs `vim`.

The test steps are as follow:
- prefetch the dependencies
- validate the prefetched output and the resulting SBOM
- build an image hermetically with `vim` installed
