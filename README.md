# AWS ECR Login

A tiny utility to simplify docker login process with AWS ECR.

## Usage

### Installation

Install using Go 1.6+ toolchain:

```shell
go install github.com/MOZGIII/aws-ecr-login/cmd/aws-ecr-login
```

### Configuration

All configuration comes from standard sources for AWS CLI and other tooling.

The most straightforward way is to set `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` and `AWS_REGION` environment variables.

### Invocation

Just run `aws-ecr-login`.

### Example

```shell
export AWS_ACCESS_KEY_ID=...
export AWS_SECRET_ACCESS_KEY=...
export AWS_REGION=...
aws-ecr-login
```

It will print full docker login command to the standard output.

### Usage with docker

This app comes with a docker image: `mozgiii/aws-ecr-login`.
Usage pattern is much the same: configure environment variables and run the command.

```shell
docker run --rm -it \
  -e AWS_ACCESS_KEY_ID=... \
  -e AWS_SECRET_ACCESS_KEY=... \
  -e AWS_REGION=... \
  mozgiii/aws-ecr-login
```
