// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: spacex/api/satellites/network/ut_disablement_codes.proto

package ut_disablement_codes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UtDisablementCode int32

const (
	UtDisablementCode_UNKNOWN_STATE                UtDisablementCode = 0
	UtDisablementCode_OKAY                         UtDisablementCode = 1
	UtDisablementCode_NO_ACTIVE_ACCOUNT            UtDisablementCode = 2
	UtDisablementCode_TOO_FAR_FROM_SERVICE_ADDRESS UtDisablementCode = 3
	UtDisablementCode_IN_OCEAN                     UtDisablementCode = 4
	UtDisablementCode_INVALID_COUNTRY              UtDisablementCode = 5
	UtDisablementCode_BLOCKED_COUNTRY              UtDisablementCode = 6
	UtDisablementCode_DATA_OVERAGE_SANDBOX_POLICY  UtDisablementCode = 7
	UtDisablementCode_CELL_IS_DISABLED             UtDisablementCode = 8
	UtDisablementCode_ROAM_RESTRICTED              UtDisablementCode = 10
	UtDisablementCode_UNKNOWN_LOCATION             UtDisablementCode = 11
	UtDisablementCode_ACCOUNT_DISABLED             UtDisablementCode = 12
	UtDisablementCode_UNSUPPORTED_VERSION          UtDisablementCode = 13
)

// Enum value maps for UtDisablementCode.
var (
	UtDisablementCode_name = map[int32]string{
		0:  "UNKNOWN_STATE",
		1:  "OKAY",
		2:  "NO_ACTIVE_ACCOUNT",
		3:  "TOO_FAR_FROM_SERVICE_ADDRESS",
		4:  "IN_OCEAN",
		5:  "INVALID_COUNTRY",
		6:  "BLOCKED_COUNTRY",
		7:  "DATA_OVERAGE_SANDBOX_POLICY",
		8:  "CELL_IS_DISABLED",
		10: "ROAM_RESTRICTED",
		11: "UNKNOWN_LOCATION",
		12: "ACCOUNT_DISABLED",
		13: "UNSUPPORTED_VERSION",
	}
	UtDisablementCode_value = map[string]int32{
		"UNKNOWN_STATE":                0,
		"OKAY":                         1,
		"NO_ACTIVE_ACCOUNT":            2,
		"TOO_FAR_FROM_SERVICE_ADDRESS": 3,
		"IN_OCEAN":                     4,
		"INVALID_COUNTRY":              5,
		"BLOCKED_COUNTRY":              6,
		"DATA_OVERAGE_SANDBOX_POLICY":  7,
		"CELL_IS_DISABLED":             8,
		"ROAM_RESTRICTED":              10,
		"UNKNOWN_LOCATION":             11,
		"ACCOUNT_DISABLED":             12,
		"UNSUPPORTED_VERSION":          13,
	}
)

func (x UtDisablementCode) Enum() *UtDisablementCode {
	p := new(UtDisablementCode)
	*p = x
	return p
}

func (x UtDisablementCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UtDisablementCode) Descriptor() protoreflect.EnumDescriptor {
	return file_spacex_api_satellites_network_ut_disablement_codes_proto_enumTypes[0].Descriptor()
}

func (UtDisablementCode) Type() protoreflect.EnumType {
	return &file_spacex_api_satellites_network_ut_disablement_codes_proto_enumTypes[0]
}

func (x UtDisablementCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UtDisablementCode.Descriptor instead.
func (UtDisablementCode) EnumDescriptor() ([]byte, []int) {
	return file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescGZIP(), []int{0}
}

type AccountDisablementReason int32

const (
	AccountDisablementReason_NO_RESTRICTION              AccountDisablementReason = 0
	AccountDisablementReason_KNOW_YOUR_CUSTOMER_REQUIRED AccountDisablementReason = 1
)

// Enum value maps for AccountDisablementReason.
var (
	AccountDisablementReason_name = map[int32]string{
		0: "NO_RESTRICTION",
		1: "KNOW_YOUR_CUSTOMER_REQUIRED",
	}
	AccountDisablementReason_value = map[string]int32{
		"NO_RESTRICTION":              0,
		"KNOW_YOUR_CUSTOMER_REQUIRED": 1,
	}
)

func (x AccountDisablementReason) Enum() *AccountDisablementReason {
	p := new(AccountDisablementReason)
	*p = x
	return p
}

func (x AccountDisablementReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountDisablementReason) Descriptor() protoreflect.EnumDescriptor {
	return file_spacex_api_satellites_network_ut_disablement_codes_proto_enumTypes[1].Descriptor()
}

func (AccountDisablementReason) Type() protoreflect.EnumType {
	return &file_spacex_api_satellites_network_ut_disablement_codes_proto_enumTypes[1]
}

func (x AccountDisablementReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountDisablementReason.Descriptor instead.
func (AccountDisablementReason) EnumDescriptor() ([]byte, []int) {
	return file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescGZIP(), []int{1}
}

var File_spacex_api_satellites_network_ut_disablement_codes_proto protoreflect.FileDescriptor

var file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDesc = []byte{
	0x0a, 0x38, 0x73, 0x70, 0x61, 0x63, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x61, 0x74,
	0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f,
	0x75, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x53, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65,
	0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2a, 0xcc, 0x02, 0x0a, 0x11, 0x55, 0x74,
	0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x4b, 0x41, 0x59, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11,
	0x4e, 0x4f, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x4f, 0x4f, 0x5f, 0x46, 0x41, 0x52, 0x5f, 0x46,
	0x52, 0x4f, 0x4d, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x52,
	0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x5f, 0x4f, 0x43, 0x45, 0x41,
	0x4e, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43,
	0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x06, 0x12, 0x1f, 0x0a,
	0x1b, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x41,
	0x4e, 0x44, 0x42, 0x4f, 0x58, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x07, 0x12, 0x14,
	0x0a, 0x10, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x49, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c,
	0x45, 0x44, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x4f, 0x41, 0x4d, 0x5f, 0x52, 0x45, 0x53,
	0x54, 0x52, 0x49, 0x43, 0x54, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0b, 0x12,
	0x14, 0x0a, 0x10, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42,
	0x4c, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f,
	0x52, 0x54, 0x45, 0x44, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x0d, 0x22, 0x04,
	0x08, 0x09, 0x10, 0x09, 0x2a, 0x12, 0x55, 0x4e, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x44,
	0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x2a, 0x4f, 0x0a, 0x18, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x52,
	0x49, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x4b, 0x4e, 0x4f, 0x57,
	0x5f, 0x59, 0x4f, 0x55, 0x52, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescOnce sync.Once
	file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescData = file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDesc
)

func file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescGZIP() []byte {
	file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescOnce.Do(func() {
		file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescData)
	})
	return file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescData
}

var file_spacex_api_satellites_network_ut_disablement_codes_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spacex_api_satellites_network_ut_disablement_codes_proto_goTypes = []any{
	(UtDisablementCode)(0),        // 0: SpaceX.API.Satellites.Network.UtDisablementCode
	(AccountDisablementReason)(0), // 1: SpaceX.API.Satellites.Network.AccountDisablementReason
}
var file_spacex_api_satellites_network_ut_disablement_codes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spacex_api_satellites_network_ut_disablement_codes_proto_init() }
func file_spacex_api_satellites_network_ut_disablement_codes_proto_init() {
	if File_spacex_api_satellites_network_ut_disablement_codes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spacex_api_satellites_network_ut_disablement_codes_proto_goTypes,
		DependencyIndexes: file_spacex_api_satellites_network_ut_disablement_codes_proto_depIdxs,
		EnumInfos:         file_spacex_api_satellites_network_ut_disablement_codes_proto_enumTypes,
	}.Build()
	File_spacex_api_satellites_network_ut_disablement_codes_proto = out.File
	file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDesc = nil
	file_spacex_api_satellites_network_ut_disablement_codes_proto_goTypes = nil
	file_spacex_api_satellites_network_ut_disablement_codes_proto_depIdxs = nil
}
