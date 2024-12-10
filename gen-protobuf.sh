#!/usr/bin/env bash

rm -rf ./pkg
mkdir ./pkg

dish_version=$(cat VERSION)
echo "$dish_version"

protoset_file=protoset/"$dish_version"/dish.protoset

protoc --go_out=./pkg/ --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex/api/common/status/status.proto
protoc --go_out=./pkg/ --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex/api/device/command.proto
protoc --go_out=./pkg/ --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex/api/device/common.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ \
       --go_opt=Mspacex/api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go_opt=Mspacex/api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --go_opt=Mspacex/api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --go_opt=Mspacex/api/device/services/unlock/service.proto=spacex.com/api/device/services/unlock/service \
       --go-grpc_opt=Mspacex/api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go-grpc_opt=Mspacex/api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --go-grpc_opt=Mspacex/api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --go-grpc_opt=Mspacex/api/device/services/unlock/service.proto=spacex.com/api/device/services/unlock/service \
       --descriptor_set_in="$protoset_file" \
       spacex/api/device/device.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ \
       --go_opt=Mspacex/api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go_opt=Mspacex/api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --go-grpc_opt=Mspacex/api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go-grpc_opt=Mspacex/api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --descriptor_set_in="$protoset_file" \
       spacex/api/device/dish.proto

protoc --go_out=./pkg/ --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex/api/device/dish_config.proto
protoc --go_out=./pkg/ --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex/api/device/rssi_scan.proto

protoc --go_out=./pkg/ --go-grpc_out=./pkg/ \
       --go_opt=Mspacex/api/device/services/unlock/service.proto=spacex.com/api/device/services/unlock/service \
       --go-grpc_opt=Mspacex/api/device/services/unlock/service.proto=spacex.com/api/device/services/unlock/service \
       --descriptor_set_in="$protoset_file" \
       spacex/api/device/services/unlock/service.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ \
       --go_opt=Mspacex/api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go_opt=Mspacex/api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --go-grpc_opt=Mspacex/api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go-grpc_opt=Mspacex/api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --descriptor_set_in="$protoset_file" \
       spacex/api/device/transceiver.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ \
       --go_opt=Mspacex/api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go_opt=Mspacex/api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --go-grpc_opt=Mspacex/api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go-grpc_opt=Mspacex/api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --descriptor_set_in="$protoset_file" \
       spacex/api/satellites/network/ut_disablement_codes.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ \
       --go_opt=Mspacex/api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --go-grpc_opt=Mspacex/api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --descriptor_set_in="$protoset_file" \
       spacex/api/telemetron/public/common/time.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ \
       --go_opt=Mspacex/api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --go-grpc_opt=Mspacex/api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --descriptor_set_in="$protoset_file" \
       spacex/api/device/wifi.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex/api/device/wifi_config.proto
protoc --go_out=./pkg/  --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex/api/device/wifi_util.proto

find ./pkg/spacex.com -name "*.go" -print0 | xargs -0 sed -i 's|spacex.com/api|github.com/stepga/starlink-grpc-golang/pkg/spacex.com/api|g'
