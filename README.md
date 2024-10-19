# starlink-grpc-golang

Archive of the Starlink gRPC protobuf API for Golang.

[CHANGELOG](./CHANGELOG.md)

The [pkg](./pkg/) directory contains the latest version of the autogenerated gRPC code from the Starlink gRPC protosets.

The [protoset](./protoset/) directory contains the archived protosets from the Starlink firmware for the `dish`.

## Get started

Install [grpcurl](https://github.com/fullstorydev/grpcurl) from https://github.com/fullstorydev/grpcurl/releases/tag/v1.9.1.

On Debian-based systems, you can install the required dependencies with:

```bash
sudo apt install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Run [`get-protoset.sh`](./get-protoset.sh) to download the latest dish protoset.

Run [`gen-protobuf.sh`](./gen-protobuf.sh) to auto-generate the gRPC code for Golang. Note that the auto-generated code has not been tested.

## Useful `grpcurl` commands

```bash
grpcurl -plaintext -d {\"get_status\":{}} 192.168.100.1:9200 SpaceX.API.Device.Device/Handle
```

```bash
grpcurl -plaintext -d {\"dish_get_obstruction_map\":{}} 192.168.100.1:9200 SpaceX.API.Device.Device/Handle
```

```bash
grpcurl -plaintext -d {\"get_diagnostics\":{}} 192.168.100.1:9200 SpaceX.API.Device.Device/Handle
```

```bash
grpcurl -plaintext -d {\"get_location\":{}} 192.168.100.1:9200 SpaceX.API.Device.Device/Handle
```

```bash
grpcurl -plaintext -d {\"get_device_info\":{}} 192.168.100.1:9200 SpaceX.API.Device.Device/Handle
```

```bash
grpcurl -plaintext -d {\"get_history\":{}} 192.168.100.1:9200 SpaceX.API.Device.Device/Handle
```

Dump the compiled protoset binary files for `dish`, `mesh` and `router`. In this repository, we only focused on the `dish` protoset.

```bash
grpcurl -plaintext -protoset-out dish.protoset 192.168.100.1:9200 describe SpaceX.API.Device.Device
```

```bash
grpcurl -plaintext -protoset-out mesh.protoset 192.168.100.1:9200 describe SpaceX.API.Device.Mesh
```

```bash
grpcurl -plaintext -protoset-out router.protoset 192.168.1.1:9000 describe SpaceX.API.Device.Device
```

## Disclaimer

This repository is not affiliated with, endorsed by, or in any way connected to Starlink, SpaceX Inc., or any of their subsidiaries. All content provided here is either independently developed or obtained from publicly available sources on the Internet and is for educational and informational purposes only.
