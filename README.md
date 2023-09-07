# Yarnberry [Zero-Installs](https://v3.yarnpkg.com/features/zero-installs) Integration Test

This integration test case showcases Yarnberry's zero-installs feature, enabling dependency management without the need to install packages locally.

Combination of [Plug'n'Play](https://v3.yarnpkg.com/features/pnp) and [Offline Cache](https://v3.yarnpkg.com/features/offline-cache).
Each time a package is downloaded from a remote location a copy will be stored within the cache. By default, it is `.yarn/cache` directory.
There is no need to run `yarn install`.

## [Protocols](https://v3.yarnpkg.com/features/protocols)

- [x] Semver
- [ ] Tag
- [x] Npm alias
- [x] Git
- [x] GitHub
- [x] File
- [x] Link
- [x] Patch
- [x] Portal
- [x] Workspace
- [x] Exec
- [x] Https

## Dependencies

- [x] dev dependencies
- [x] optional dependencies
- [x] peer dependencies
- [x] scoped dependencies
- [x] compiled dependencies

### Scoped dependencies

Starting with _@_.

```bash
"@types/node": "^20.4.5",
"@types/yargs": "^17",
```

### Optional dependencies

See this [commit](https://github.com/cachito-testing/cachi2-yarn-berry/commit/4326fbbde1b1770e752138a9347e13fcfaabc9be).

```bash
"fsevents": "^2.3.2", # only for MacOS
```

### Compiled dependencies

See this [commit](https://github.com/cachito-testing/cachi2-yarn-berry/commit/4066baa6ff91ce6d4491370479d8b8c23ed7dd37).

Let the container build to do the compilation, not cachi2.

```bash
"tree-sitter-json": "^0.20.0",
```
