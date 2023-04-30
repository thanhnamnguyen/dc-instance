# Data Commons Mixer Developer Guide

The developement uses
[Kustomize](https://kubectl.docs.kubernetes.io/guides/introduction/kustomize/)
to manage yaml templating and composition. Detailed deploy folder structure can
be found [here](../deploy/README.md).

Local development uses [Skaffold](https://skaffold.dev) to manage the build,
deploy and port forwarding.

## Prerequisite

- Contact DataCommons team to get data access to Cloud Bigtable and BigQuery.

- Install the following tools to develop mixer locally (with Docker):

  - [`Docker`](https://www.docker.com/products/docker-desktop)
  - [`Minikube`](https://minikube.sigs.k8s.io/docs/start/)
  - [`Skaffold`](https://skaffold.dev/docs/install/)
  - [`gcloud`](https://cloud.google.com/sdk/docs/install)
  - [`kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

- If you prefer to do it locally without Docker, then need to install the
  following:

  - [`Golang`](https://golang.org/doc/install)
  - [`protoc`](http://google.github.io/proto-lens/installing-protoc.html)

  Make sure to add `GOPATH` and update `PATH`:

  ```bash
  # Use the actual path of your Go installation
  export GOPATH=/Users/<USER>/go/
  export PATH=$PATH:$GOPATH/bin
  ```

- Authenticate to GCP

  ```bash
  gcloud components update
  gcloud auth login
  gcloud auth application-default login
  ```

## Develop mixer locally as a Go server (Recommended)

**NOTE** This can only develop and test the gRPC server. Since the
[ESP](https://cloud.google.com/endpoints/docs/grpc/running-esp-localdev) is not
brought up here, can not test the REST API.

### Generate Go proto files

Install the following packages as a one-time action.

```bash
cd ~/   # Be sure there is no go.mod in the local directory
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
```

Run the following command to generate Go proto files.

```bash
# In repo root directory
protoc \
  --proto_path=proto \
  --go_out=internal \
  --go-grpc_out=internal \
  --go-grpc_opt=require_unimplemented_servers=false \
  proto/*.proto proto/v1/*.proto
```

### Start Mixer as a gRPC server backed by Cloud BigTable (BigQuery)

Run the following code to start mixer gRPC server (without branch cache)

```bash
# In repo root directory
go run --tags sqlite_fts5 cmd/main.go \
    --mixer_project=datcom-mixer-staging \
    --bq_dataset=$(head -1 deploy/storage/bigquery.version) \
    --base_bigtable_info="$(cat deploy/storage/base_bigtable_info.yaml)" \
    --custom_bigtable_info="$(cat test/custom_bigtable_info.yaml)" \
    --schema_path=$PWD/deploy/mapping/ \
    --use_branch_bt=false

go run examples/main.go
```

### Start Mixer as a gRPC server backed by TMCF + CSV files

Mixer can load and serve TMCF + CSV files. This is used for a private Data Commons
instance. This requires to set the following flag

- `--use_tmcf_csv_data=true`
- `--tmcf_csv_bucket=<bucket-name>`
- `--tmcf_csv_folder=<folder-name>`

Prerequists:

- Create a GCS bucket <BUCKET_NAME>
- Create a folder in the bucket <FOLDER_NAME> to host all the data files

Run the following code to start mixer gRPC server with TMCF + CSV files stored in GCS

```bash
# In repo root directory
go run --tags sqlite_fts5 cmd/main.go \
    --mixer_project=datcom-mixer-dev-316822 \
    --tmcf_csv_bucket=datcom-mixer-dev-resources \
    --tmcf_csv_folder=test \
    --use_tmcf_csv_data=true \
    --use_bigquery=false \
    --use_base_bt=false \
    --use_branch_bt=false
```

### Running ESP locally

Mixer is a gRPC service but callers (website, API clients) are normally http
clients. Therefore developing and testing mixer locally often requires both the
mixer gRPC server and its corresponding json transcoding server. HTTP to gRPC
translation can be done locally thorugh [envoy
proxy](https://www.envoyproxy.io/docs/envoy/latest/intro/what_is_envoy). To
install envoy, please follow the [official
doc](https://www.envoyproxy.io/docs/envoy/latest/start/install).

Before running envoy proxy, please make sure mixer service
definition(mixer-grpc.pb) is available by running the follow from repo root.

```sh
protoc --proto_path=proto \
  --include_source_info --include_imports \
  --descriptor_set_out mixer-grpc.pb \
  proto/*.proto proto/v1/*.proto
```

Assuming mixer gRPC server is at `localhost:12345`, run the following from repo
root to spin up envoy proxy. This exposes the http mixer service at
`localhost:8081`.

```sh
envoy --config-path esp/envoy-config.yaml
```

Mixer services in live clusters do not use envoy directly. To replicate the live
k8s setup locally, please see [the guide
below](#fully-replicating-mixer-k8s-setup-locally).

### Update Go package dependencies

To view possible updates:

```bash
go list -m -u all
```

To update:

```bash
go get -u ./...
go mod tidy
```

### Run Tests (Go)

```bash
./scripts/run_test.sh
```

### Update e2e test golden files (Go)

```bash
./scripts/update_golden.sh
```

### Run import group latency tests

In root directory, run:

```bash
./test/e2e/run_latency.sh
```

### Profile a program

Install [Graphgiz](https://graphviz.org/).

```bash
go test -v -parallel 1 -cpuprofile cpu.prof -memprofile mem.prof XXX_test.go
go tool pprof -png cpu.prof
go tool pprof -png mem.prof
```

### Profile Mixer Startup Memory

Run the regular `go run cmd/main.go` command that you'd like to profile with the
flag `--startup_memprof=<output_file_path>`. This will save the memory profile
to that path, and you can use `go tool pprof` to analyze it. For example;

```bash
# Command from ### Start Mixer as a gRPC server backed by TMCF + CSV files
# In repo root directory
go run --tags sqlite_fts5 cmd/main.go \
    --mixer_project=datcom-mixer-staging \
    --bq_dataset=$(head -1 deploy/storage/bigquery.version) \
    --base_bigtable_info="$(cat deploy/storage/base_bigtable_info.yaml)" \
    --schema_path=$PWD/deploy/mapping/ \
    --use_branch_bt=true
    --startup_memprof=grpc.memprof     # <-- note the additional flag here

# -sample_index=alloc_space reports on all memory allocations, including those
# that have been garbage collected. use -sample_index=inuse_space for memory
# still in use after garbage collection
go tool pprof -sample_index=alloc_space -png grpc.memprof
```

### Profile API Requests against a Running Mixer Instance

Run the regular `go run cmd/main.go` command that you'd like to profile with the
flag `--httpprof_port=<port, recommended 6060>`. This will run the mixer server
with an HTTP handler at that port serving memory and CPU profiles of the running
server.

```bash
# Command from ### Start Mixer as a gRPC server backed by TMCF + CSV files
# In repo root directory
go run cmd/main.go \
    --mixer_project=datcom-mixer-staging \
    --bq_dataset=$(head -1 deploy/storage/bigquery.version) \
    --base_bigtable_info="$(cat deploy/storage/base_bigtable_info.yaml)" \
    --schema_path=$PWD/deploy/mapping/ \
    --use_branch_bt=true
    --httpprof_port=6060     # <-- note the additional flag here
```

Once this server is ready to serve requests, you can send it requests and use
the profile handler to retrieve memory and CPU profiles. `test/http_memprof/http_memprof.go`
is a program that automatically sends and profiles the memory usage of given
gRPC calls. You can update this file to your profiling needs or use it as a
starting point for an independent script that will automatically run a suite of
tests.

```bash
# in another process...
go run test/http_memprof/http_memprof.go \
  --grpc_addr=127.0.0.1:12345 \ # default is given; where to find the Mixer server
  --prof_addr=127.0.0.1:6060 # default is given; where to find the live profile handler
```

`go tool pprof` also supports ad-hoc profiling of servers started as described
above. To use, specify the URL at which the HTTP handler can be found as the
input file argument. `pprof` will download a profile from the handler and open
in interactive mode to run queries.

```bash
# ?gc=1 triggers a garbage collection run before the memory profile is served
# See net/http/pprof for other URLs and profiles available https://pkg.go.dev/net/http/pprof
# with no flags specifying output, pprof goes into interactive mode
go tool pprof -sample_index=alloc_space 127.0.0.1:6060/debug/pprof/heap?gc=1
```

## Fully replicating mixer k8s setup locally

Mixer and [ESP](https://cloud.google.com/endpoints/docs/grpc/running-esp-localdev)
is deployed on a local Minikube cluster.
To avoid using Endpoints API management and talking to GCP,
local deployment uses Json API configuration,
which is compiled using [API Compiler](https://github.com/googleapis/api-compiler).

### Start mixer in minikube

Mixer in Minikube (local Kubernetes) brings up mixer with REST endpoints through
ESP. This requires docker and minikube to be installed in local environment. To
start Mixer in Minikube, run:

```bash
minikube start
minikube addons enable gcp-auth
eval $(minikube docker-env)
kubectl config use-context minikube
skaffold dev --port-forward -n mixer
```

This exposes the local mixer service at `localhost:8081`.

To verify the server serving request:

```bash
curl http://localhost:8081/v1/bulk/observations/series?entities=geoId/06025&variables=Count_Person
```

After code edit, the container images are automatically rebuilt and re-deployed
to the local cluster.

Starting Mixer and ESP locally allows development that relies on a custom mixer.
For example, Data Commons website local development may need a mixer with custom
Bigtables. To do this, update
[deploy/local/custom_bigtable_info.yaml](../deploy/local/custom_bigtable_info.yaml)
and the local mixer would be able to serve the custom Bigtable.

### Run Tests

```bash
./scripts/run_test.sh -d
```

### Update e2e test golden files

```bash
./scripts/update_golden.sh -d
```

## Update prod golden files

Run the following commands to update prod golden files from staging golden files

```bash
./scripts/update_golden_prod.sh
```
