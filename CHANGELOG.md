<a name="unreleased"></a>
## [Unreleased]


<a name="v1.4.2"></a>
## [v1.4.2] - 2023-07-01
### Bug Fixes
- 0e5498a - dependabots security issue  ([#22](https://github.com/vdbulcke/cert-monitor/issues/22))


<a name="v1.4.1"></a>
## [v1.4.1] - 2023-03-25
### Bug Fixes
- e1f0572 - security issue ([#21](https://github.com/vdbulcke/cert-monitor/issues/21))


<a name="v1.4.0"></a>
## [v1.4.0] - 2023-03-04
### Bug Fixes
- eb9b418 - go lint workflow
- b25550e - dependabots security issues ([#20](https://github.com/vdbulcke/cert-monitor/issues/20))

### Features
- e45d149 - add cosign signature via goreleaser
- b86d921 - bump go version 1.20


<a name="v1.3.3"></a>
## [v1.3.3] - 2022-12-20
### Bug Fixes
- cd832a5 - build(deps): bump github.com/crewjam/saml from 0.4.6 to 0.4.9


<a name="v1.3.2"></a>
## [v1.3.2] - 2022-10-09
### Features
- 8f6fedb - switch to ghcr.io registry


<a name="v1.3.1"></a>
## [v1.3.1] - 2022-06-02
### Bug Fixes
- 8958680 - Dockerfile for gorleaser build ([#16](https://github.com/vdbulcke/cert-monitor/issues/16))

### Features
- 56217be - build/publish docker image in Action ([#16](https://github.com/vdbulcke/cert-monitor/issues/16))


<a name="v1.3.0"></a>
## [v1.3.0] - 2022-04-03
### Features
- 1ab2582 - add 'kty' filter ([#14](https://github.com/vdbulcke/cert-monitor/issues/14))


<a name="v1.2.1"></a>
## [v1.2.1] - 2022-01-16
### Bug Fixes
- 60475ec - Handle non pem cert in static dir [#13](https://github.com/vdbulcke/cert-monitor/issues/13)


<a name="v1.2.0"></a>
## [v1.2.0] - 2022-01-09
### Features
- a5d84d8 - Add build details in version command
- 5063054 - Add --no-color flag for fetch
- 24bb312 - Add skip TLS validation option

### BREAKING CHANGE


> Server TLS and TCP endpont skip-tls-validation
was implied before.
Add 'skip_tls_validation: true' to your config



<a name="v1.1.1"></a>
## [v1.1.1] - 2021-12-31
### Bug Fixes
- c0c08a4 - notlint false positive
- 7c9d487 - linting issues


<a name="v1.1.0"></a>
## [v1.1.0] - 2021-12-29
### Bug Fixes
- 5205c32 - avoid import loop

### Features
- 21e0bd2 - windows release
- 8626268 - jwk remote endpoint


<a name="v1.0.0"></a>
## [v1.0.0] - 2021-12-22
### Features
- e5e6289 - create cli commands


<a name="v0.4.1"></a>
## [v0.4.1] - 2021-11-14
### Bug Fixes
- f9f7788 - goreleaser univeral binary and release note


<a name="v0.4.0"></a>
## [v0.4.0] - 2021-11-14
### Features
- a90ee1c - update github action config for goreleaser


<a name="v0.3.0"></a>
## [v0.3.0] - 2021-05-31

<a name="v0.2.0"></a>
## [v0.2.0] - 2021-05-04

<a name="v0.1.0"></a>
## v0.1.0 - 2021-04-02

[Unreleased]: https://github.com/vdbulcke/cert-monitor/compare/v1.4.2...HEAD
[v1.4.2]: https://github.com/vdbulcke/cert-monitor/compare/v1.4.1...v1.4.2
[v1.4.1]: https://github.com/vdbulcke/cert-monitor/compare/v1.4.0...v1.4.1
[v1.4.0]: https://github.com/vdbulcke/cert-monitor/compare/v1.3.3...v1.4.0
[v1.3.3]: https://github.com/vdbulcke/cert-monitor/compare/v1.3.2...v1.3.3
[v1.3.2]: https://github.com/vdbulcke/cert-monitor/compare/v1.3.1...v1.3.2
[v1.3.1]: https://github.com/vdbulcke/cert-monitor/compare/v1.3.0...v1.3.1
[v1.3.0]: https://github.com/vdbulcke/cert-monitor/compare/v1.2.1...v1.3.0
[v1.2.1]: https://github.com/vdbulcke/cert-monitor/compare/v1.2.0...v1.2.1
[v1.2.0]: https://github.com/vdbulcke/cert-monitor/compare/v1.1.1...v1.2.0
[v1.1.1]: https://github.com/vdbulcke/cert-monitor/compare/v1.1.0...v1.1.1
[v1.1.0]: https://github.com/vdbulcke/cert-monitor/compare/v1.0.0...v1.1.0
[v1.0.0]: https://github.com/vdbulcke/cert-monitor/compare/v0.4.1...v1.0.0
[v0.4.1]: https://github.com/vdbulcke/cert-monitor/compare/v0.4.0...v0.4.1
[v0.4.0]: https://github.com/vdbulcke/cert-monitor/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/vdbulcke/cert-monitor/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/vdbulcke/cert-monitor/compare/v0.1.0...v0.2.0
