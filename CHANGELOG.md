# Changelog
All notable changes to this project will be documented in this file.

**ATTN**: This project uses [semantic versioning](http://semver.org/).

## [Unreleased]
### Added
- Documented ARK: Survival Ascended as a supported game. It speaks the same basic Source RCON protocol as ARK: Survival Evolved, so no code changes were needed.

### Changed
- Simplified `logger.OpenFile` to a single `os.OpenFile` call with `O_CREATE|O_APPEND` instead of stat-then-branch, removing a TOCTOU race and tightening file/directory permissions (0666/0766 -> 0644/0755).
- Used `filepath.Join` instead of manual string concatenation for the default config path, fixing path separator handling on Windows.
- Reduced repeated map lookups in `Executor.NewSession` by reading the config environment once instead of on every field.

### Security
- Fixed a resource leak in `config_test.go` where the log file was not closed on write error.
- Bumped Go toolchain to go1.26.5, fixing 3 standard library CVEs (crypto/tls, net/textproto, crypto/x509) reachable from the RCON client dial/handshake path.
- Removed the unused `golang.org/x/net` indirect dependency (was pinned at v0.20.0 with 9 known vulnerabilities).
- Updated Docker build image from `golang:1.19.3-alpine` to `golang:1.26.5-alpine3.24`, and the runtime image from unpinned `alpine` to `alpine:3.24`.

### Updated
- Updated Go toolchain requirement to go1.26.5 (was go1.21); the declared language version (`go` directive) stays at 1.23, the actual minimum required by dependencies.
- Updated `github.com/gorcon/rcon` v1.3.5 -> v1.4.0.
- Updated `github.com/gorilla/websocket` v1.5.1 -> v1.5.3.
- Updated `github.com/urfave/cli/v2` v2.27.1 -> v2.27.7.
- Updated `github.com/stretchr/testify` v1.7.1 -> v1.11.1.
- Updated `github.com/cpuguy83/go-md2man/v2` v2.0.3 -> v2.0.7.
- Updated CI workflow actions (`actions/checkout`, `actions/setup-go`, `golangci-lint-action`) to current major versions and golangci-lint to v1.61.0.
- Added a `workflow_dispatch` trigger so the build job can be run manually from the Actions tab.

### Removed
- Removed the `docker-release` job from the CI workflow (published images to a Docker Hub account this fork doesn't control).

## [v0.10.3] - 2023-03-11
### Added
- Added folder creation for logs.

### Fixed
- Fixed default value to config flag.

### Updated
- Updated Go modules (go1.19).
- Updated golang-ci linter (1.50.1).
- Updated dependencies.

## [v0.10.2] - 2022-03-09
### Added
- Added `--timeout, -T` flag, allowed to set dial and execute timeout.
- Added Makefile.

## [v0.10.1] - 2021-11-13
### Fixed
- Fixed close connection panic #19 

## [v0.10.0] - 2021-09-09
### Added
- Added `--skip, -s` flag, allowed to skip error on multiple commands or in terminal mode.

### Changed
- Create one keep alive connection in terminal mode.

## [v0.9.1] - 2021-03-14
### Added
- Added Dockerfile. 

### Changed
- Disabled CGO in release script.

## [v0.9.0] - 2020-12-20
### Added
- Added config validation. 
- Added config JSON format supporting. 
- Added protocol type asking in interactive mode.
- Added protocol type validation in interactive mode.
- Added the ability to send several commands in a row with one request #10

### Changed
- Removed `-c value, --command value` flag. It replaced to commands without flags.
- Flag `--cfg value` replaced to `-c value, --config value` flag. 

## [v0.8.1] - 2020-11-17
### Added
- Added tests for real servers. Servers list: `Project Zomboid`, `7 Days to Die`, `Rust`. 
- Added removing part of the constantly repeated data from `7 Days to Die` response ([details](https://github.com/gorcon/telnet/issues/1)). 

## [v0.8.0-beta.2] - 2020-10-18
### Added
- Added interactive mode for Web RCON #12.

### Fixed
- Fixed response for another request for Rust server #13.

## [v0.8.0-beta] - 2020-10-18
### Added
- Added Rust Web RCON support #8. Add `-t web` argument when execute `rcon` cli.

### Changed
- Code and tests refactoring.

### Fixed
- Fixed changelog.

## [v0.7.0] - 2020-10-10
### Added
- Added support amd64 darwin compilation.
- Added 7 Days to Die support #5. Add `-t telnet` argument when execute `rcon` cli.

## [v0.6.0] - 2020-07-10
### Fixed
- Fix `Conan Exiles "response for another request" error` bug #6.

## [v0.5.0] - 2020-06-18
### Fixed
- Correction of text after AutoCorrect.

### Added
- More tests.
- Add Go modules (go1.13).
- Add golang-ci linter.

## [v0.4.0] - 2019-08-05
### Added
- Added argument `-l`, `--log` to pass custom log path/name. Argument has higher priority .
than entry in configuration file.
- Added interactive opportunity to enter server address and password.

## [v0.3.0] - 2019-07-28
### Added
- Print error messages for missed address and password arguments.
- Added argument `--cfg` to pass custom config path/name.
- Check of parameters of authorization of a remote server before launching interactive mode. 
- Added rcon.yaml as config sample.
- Added log variable to config file that enables logs for requests and responses for remote server.

### Removed
- Remove `rcon-upx` from release.
- Remove 'cli' command to run in interactive mode. For use interactive mode run `rcon` without `-c` argument.

## [v0.2.0] - 2019-07-27
### Added
- Added environments to config.

### Fixed
- Fix global options in interactive mode.

## v0.1.0 - 2019-07-22
### Added
- Initial implementation.

[Unreleased]: https://github.com/gorcon/rcon-cli/compare/v0.10.3...HEAD
[v0.10.3]: https://github.com/gorcon/rcon-cli/compare/v0.10.2...v0.10.3
[v0.10.2]: https://github.com/gorcon/rcon-cli/compare/v0.10.1...v0.10.2
[v0.10.1]: https://github.com/gorcon/rcon-cli/compare/v0.10.0...v0.10.1
[v0.10.0]: https://github.com/gorcon/rcon-cli/compare/v0.9.1...v0.10.0
[v0.9.1]: https://github.com/gorcon/rcon-cli/compare/v0.9.0...v0.9.1
[v0.9.0]: https://github.com/gorcon/rcon-cli/compare/v0.8.1...v0.9.0
[v0.8.1]: https://github.com/gorcon/rcon-cli/compare/v0.8.0-beta.2...v0.8.1
[v0.8.0-beta.2]: https://github.com/gorcon/rcon-cli/compare/v0.8.0-beta...v0.8.0-beta.2
[v0.8.0-beta]: https://github.com/gorcon/rcon-cli/compare/v0.7.0...v0.8.0-beta
[v0.7.0]: https://github.com/gorcon/rcon-cli/compare/v0.6.0...v0.7.0
[v0.6.0]: https://github.com/gorcon/rcon-cli/compare/0.5.0...v0.6.0
[v0.5.0]: https://github.com/gorcon/rcon-cli/compare/v0.4.0...0.5.0
[v0.4.0]: https://github.com/gorcon/rcon-cli/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/gorcon/rcon-cli/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/gorcon/rcon-cli/compare/v0.1.0...v0.2.0
