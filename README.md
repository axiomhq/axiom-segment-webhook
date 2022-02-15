# Axiom Segment Webhook

[![Go Workflow][go_workflow_badge]][go_workflow]
[![Coverage Status][coverage_badge]][coverage]
[![Go Report][report_badge]][report]
[![Latest Release][release_badge]][release]
[![License][license_badge]][license]
[![Docker][docker_badge]][docker]

---

## Table of Contents

1. [Introduction](#introduction)
1. [Usage](#usage)
1. [Contributing](#contributing)
1. [License](#license)

## Introduction

_Axiom Segment Webhook_ provides a [Segment][1] compatible webhook that sends
data to Axiom.

  [1]: segment.io

## Installation

### Download the pre-compiled and archived binary manually

Binary releases are available on [GitHub Releases][2].

  [2]: https://github.com/axiomhq/axiom-segment-webhook/releases/latest

### Install using [Homebrew](https://brew.sh)

```shell
brew tap axiomhq/tap
brew install axiom-segment-webhook
```

To update:

```shell
brew update
brew upgrade axiom-segment-webhook
```

### Install using `go get`

```shell
go get -u github.com/axiomhq/axiom-segment-webhook/cmd/axiom-segment-webhook
```

### Install from source

```shell
git clone https://github.com/axiomhq/axiom-segment-webhook.git
cd axiom-segment-webhook
make install
```

### Run the Docker image

Docker images are available on [DockerHub][docker].

## Usage

1. Set the following environment variables to connect to **Axiom Cloud**:

* `AXIOM_TOKEN`: **Personal Access** token which can be created under
  `Setting -> Profile`.
* `AXIOM_ORG_ID`: The organization identifier of the organization to use.

When using **Axiom Selfhost**:

* `AXIOM_TOKEN`: **Personal Access** token which can be created under
  `Setting -> Profile`.
* `AXIOM_URL`: URL of the Axiom deployment to use.

2. Run it: `./axiom-segment-webhook` or using Docker:

```shell
docker run -p8080:8080/tcp \
  -e=AXIOM_TOKEN=<YOU_AXIOM_TOKEN> \
  axiomhq/axiom-segment-webhook
```

## Contributing

Feel free to submit PRs or to fill issues. Every kind of help is appreciated. 

Before committing, `make` should run without any issues.

Kindly check our [Contributing](Contributing.md) guide on how to propose
bugfixes and improvements, and submitting pull requests to the project.

## License

&copy; Axiom, Inc., 2022

Distributed under MIT License (`The MIT License`).

See [LICENSE](LICENSE) for more information.

<!-- Badges -->

[go_workflow]: https://github.com/axiomhq/axiom-segment-webhook/actions/workflows/push.yml
[go_workflow_badge]: https://img.shields.io/github/workflow/status/axiomhq/axiom-segment-webhook/Push?style=flat-square&ghcache=unused
[coverage]: https://codecov.io/gh/axiomhq/axiom-segment-webhook
[coverage_badge]: https://img.shields.io/codecov/c/github/axiomhq/axiom-segment-webhook.svg?style=flat-square&ghcache=unused
[report]: https://goreportcard.com/report/github.com/axiomhq/axiom-segment-webhook
[report_badge]: https://goreportcard.com/badge/github.com/axiomhq/axiom-segment-webhook?style=flat-square&ghcache=unused
[release]: https://github.com/axiomhq/axiom-segment-webhook/releases/latest
[release_badge]: https://img.shields.io/github/release/axiomhq/axiom-segment-webhook.svg?style=flat-square&ghcache=unused
[license]: https://opensource.org/licenses/MIT
[license_badge]: https://img.shields.io/github/license/axiomhq/axiom-segment-webhook.svg?color=blue&style=flat-square&ghcache=unused
[docker]: https://hub.docker.com/r/axiomhq/axiom-segment-webhook
[docker_badge]: https://img.shields.io/docker/pulls/axiomhq/axiom-segment-webhook.svg?style=flat-square&ghcache=unused
