# Copyright 2019 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# Creates a container that generates gRPC json config that can be used directly
# by local ESP container. These two containers share the same volume that holds
# the mixer-grpc.json.

FROM golang:1.19 AS base

WORKDIR /workspace

# Install protoc
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y unzip

RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip
RUN unzip protoc-3.19.4-linux-x86_64.zip
RUN chmod +x /workspace/bin/protoc
ENV PATH="${PATH}:/workspace/bin/"

# Install protobuf go plugin
ENV GO111MODULE=on
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0


WORKDIR /workspace

RUN git clone https://github.com/googleapis/api-compiler

WORKDIR /workspace/api-compiler

RUN git submodule update --recursive --init

# Build protobuf
COPY proto/ proto
RUN protoc \
    --proto_path=proto \
    --include_source_info \
    --include_imports \
    --descriptor_set_out mixer-grpc.pb \
    proto/*.proto proto/v1/*.proto


FROM openjdk:8-slim

COPY --from=base /workspace/api-compiler /workspace/api-compiler
WORKDIR /workspace/api-compiler

COPY esp/endpoints-dev.yaml .
RUN mkdir /output
RUN ./run.sh \
    --configs endpoints-dev.yaml \
    --descriptor mixer-grpc.pb \
    --json_out /output/mixer-grpc.json

CMD [ "/bin/sh" ]