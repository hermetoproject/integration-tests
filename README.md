# berryscary

Exploring the features of Yarn :yarn: version [Berry][berry] :blueberries:.

It's pretty scary.

## Usage

Build the container and run it (requires `podman`):

```shell
make build
make run
```

## Hermetic builds

Note that the `make build` command cuts off podman's network connection and the
build works anyway. That is thanks to Yarnberry's [Zero-Installs][0i] feature.

See also the other branches, which push Zero-Installs to its limits or test
alternative install strategies.

[berry]: https://github.com/yarnpkg/berry
[0i]: https://yarnpkg.com/features/zero-installs
