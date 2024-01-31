# Go GitHub Actions Demo

> Demo app demonstrating how to use GitHub actions to implement a CI Pipeline for a Go based CLI application.

## Features

* Simple "hello world" golang app.
* Popular Golang quality tools like Go Lint and Go vet executing in the Pipeline.
* Code coverage publishing using [Codecov](https://codecov.io/).
* Binary release with [Go Releaseer](https://goreleaser.com/)

Add the following secret to GitHub under Settings -> Secrets and variables > Actions -> Repository secrets:
```
GHUBTOKEN
```

Run the following commands to trigger a release that generates the release artifacts in Github:
```
$ git tag v0.3.0
$ git push --tags
```
