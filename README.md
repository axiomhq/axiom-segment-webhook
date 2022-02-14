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

_Axiom Segment Webhook_ allows you to forward Segment data to Axiom.

## Installation

### Download the pre-compiled and archived binary manually

Binary releases are available on [GitHub Releases][2].

  [2]: https://github.com/axiomhq/axiom-segement-webhook/releases/latest

### Install using [Homebrew](https://brew.sh)

```shell
brew tap axiomhq/tap
brew install axiom-segement-webhook
```

To update:

```shell
brew update
brew upgrade axiom-segement-webhook
```

### Install using `go get`

```shell
go get -u github.com/axiomhq/axiom-segement-webhook/cmd/axiom-segement-webhook
```

### Install from source

```shell
git clone https://github.com/axiomhq/axiom-segement-webhook.git
cd axiom-segement-webhook
make install
```

### Run the Docker image

Docker images are available on [DockerHub][docker].

## Usage

1. Set the following environment variables to connect to Axiom Cloud:

* `AXIOM_TOKEN`: **Personal Access** or **Ingest** token. Can be created under
  `Profile` or `Settings > Ingest Tokens`. For security reasons it is advised to
  use an Ingest Token with minimal privileges only.
* `AXIOM_ORG_ID`: The organization identifier of the organization to use

When using Axiom Selfhost:

* `AXIOM_TOKEN`: **Personal Access** or **Ingest** token. Can be created under
  `Profile` or `Settings > Ingest Tokens`. For security reasons it is advised to
  use an Ingest Token with minimal privileges only.
* `AXIOM_URL`: URL of the Axiom deployment to use

2. Run it: `./axiom-segement-webhook` or using Docker:

```shell
docker run -p8080:8080/tcp \
  -e=AXIOM_TOKEN=<YOUR_AXIOM_TOKEN> \
  axiomhq/axiom-segement-webhook
```

3. Set the public URL as a Webhook Desination in Segment Control Panel.

## Contributing

Feel free to submit PRs or to fill issues. Every kind of help is appreciated. 

Before committing, `make` should run without any issues.

Kindly check our [Contributing](Contributing.md) guide on how to propose
bugfixes and improvements, and submitting pull requests to the project.

## License

&copy; Axiom, Inc., 2021

Distributed under MIT License (`The MIT License`).

See [LICENSE](LICENSE) for more information.

<!-- Badges -->

[go_workflow]: https://github.com/axiomhq/axiom-segement-webhook/actions/workflows/push.yml
[go_workflow_badge]: https://img.shields.io/github/workflow/status/axiomhq/axiom-segement-webhook/Push?style=flat-square&ghcache=unused
[coverage]: https://codecov.io/gh/axiomhq/axiom-segement-webhook
[coverage_badge]: https://img.shields.io/codecov/c/github/axiomhq/axiom-segement-webhook.svg?style=flat-square&ghcache=unused
[report]: https://goreportcard.com/report/github.com/axiomhq/axiom-segement-webhook
[report_badge]: https://goreportcard.com/badge/github.com/axiomhq/axiom-segement-webhook?style=flat-square&ghcache=unused
[release]: https://github.com/axiomhq/axiom-segement-webhook/releases/latest
[release_badge]: https://img.shields.io/github/release/axiomhq/axiom-segement-webhook.svg?style=flat-square&ghcache=unused
[license]: https://opensource.org/licenses/MIT
[license_badge]: https://img.shields.io/github/license/axiomhq/axiom-segement-webhook.svg?color=blue&style=flat-square&ghcache=unused
[docker]: https://hub.docker.com/r/axiomhq/axiom-segement-webhook
[docker_badge]: https://img.shields.io/docker/pulls/axiomhq/axiom-segement-webhook.svg?style=flat-square&ghcache=unused
