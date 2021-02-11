module github.com/cachito-testing/cachito-gomod-local-deps

go 1.15

require github.com/cachito-testing/some-module v0.0.0

replace github.com/cachito-testing/some-module => ./staging/src/some-module
