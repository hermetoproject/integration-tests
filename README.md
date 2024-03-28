# Cachi2 RPMs multiple architectures test

This test the prefetching and reporting of RPMs present in two different architectures: x86_64 and aarch65. It contains the following scenarios:

- Same package in two archs (vim-minimal)
- Same package in two archs, but is missing checksums in one of them (vim-data)
- One package in only one arch (vim-common)