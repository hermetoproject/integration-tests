# Cachi2 RPMs multiple input packages test

This test the prefetching and reporting of RPMs present in two different Cachi2 input packages. It contains the following scenarios:
- Same package in two lockfiles (gzip)
- Same package in two lockfiles, but is missing checksums in one of them (glibc-common)
- Package with different versions (glibc-minimal-langpack)
