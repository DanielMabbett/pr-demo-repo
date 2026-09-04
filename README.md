# pr-demo-repo

A minimal Go HTTP service used to demonstrate Jumpstart's **pull-request
trigger + GitHub Checks status reporting** feature.

## What's in the box

```
.jumpstart.yml                  # Jumpstart pipeline definition (on.pull_request)
cmd/api/main.go                 # HTTP server entry point
internal/status/status.go       # Business logic
internal/status/status_test.go
go.mod
```

## Running locally

```sh
go run ./cmd/api
# → pr-demo-repo listening on :8091

curl http://localhost:8091/status
# → pr-demo-api: OK
```

## Running the tests

```sh
go test -v ./...
```

## Walkthrough: linking this repo to a pipeline

1. **Install the Jumpstart GitHub App** on this repository (or a fork of
   it), granting the permissions described in the main repo's
   [README — Pull-request triggers and GitHub Checks status](../../README.md#pull-request-triggers-and-github-checks-status).
2. **Create a pipeline** in Jumpstart pointing at this repo, using
   `.jumpstart.yml` as its definition (either via the "sync from repo"
   definition syncer, or by pasting the file's contents when creating the
   pipeline).
3. **Open a pull request** against `main` in this repo (any branch, any
   trivial change — e.g. edit this README).
4. **Watch the PR's "Checks" tab** — a `jumpstart / pr-demo-api` check
   appears within a few seconds, showing "in progress", then turns green or
   red as the run completes.
5. **Push a second commit** to the same PR. You'll see:
   - The *first* check run gets marked **cancelled**, with the title
     "Superseded by a newer commit" — this demonstrates the job
     cancellation behavior from the prerequisite
     [job-cancellation-design.md](../../docs/design/job-cancellation-design.md)
     work, applied end-to-end via the PR trigger.
   - A *fresh* check run starts for the new commit's SHA.

## The `.jumpstart.yml` pipeline

```yaml
on:
  pull_request:
    types: [opened, synchronize, reopened]
    branches: [main]

jobs:
  build-and-test:
    runs_on: vm
    steps:
      - run: go build ./...
      - run: go test -v -count=1 ./...
      - run: go vet ./...
```

- `types` restricts which PR webhook actions fire the pipeline. Left empty,
  this defaults to `[opened, synchronize, reopened]`.
- `branches` restricts by the PR's **base** branch. Left empty, any base
  branch matches.

At least one VM agent must be running and registered with the Jumpstart
control plane for the run to actually dispatch and execute.
