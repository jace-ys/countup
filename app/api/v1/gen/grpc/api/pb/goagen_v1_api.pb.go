// Code generated with goa v3.19.1, DO NOT EDIT.
//
// api protocol buffer definition
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.1
// source: goagen_v1_api.proto

package apipb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Provider      string                 `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	AccessToken   string                 `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthTokenRequest) Reset() {
	*x = AuthTokenRequest{}
	mi := &file_goagen_v1_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthTokenRequest) ProtoMessage() {}

func (x *AuthTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_v1_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthTokenRequest.ProtoReflect.Descriptor instead.
func (*AuthTokenRequest) Descriptor() ([]byte, []int) {
	return file_goagen_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *AuthTokenRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *AuthTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type AuthTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthTokenResponse) Reset() {
	*x = AuthTokenResponse{}
	mi := &file_goagen_v1_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthTokenResponse) ProtoMessage() {}

func (x *AuthTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_v1_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthTokenResponse.ProtoReflect.Descriptor instead.
func (*AuthTokenResponse) Descriptor() ([]byte, []int) {
	return file_goagen_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *AuthTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CounterGetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CounterGetRequest) Reset() {
	*x = CounterGetRequest{}
	mi := &file_goagen_v1_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CounterGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterGetRequest) ProtoMessage() {}

func (x *CounterGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_v1_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterGetRequest.ProtoReflect.Descriptor instead.
func (*CounterGetRequest) Descriptor() ([]byte, []int) {
	return file_goagen_v1_api_proto_rawDescGZIP(), []int{2}
}

type CounterGetResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Count           int32                  `protobuf:"zigzag32,1,opt,name=count,proto3" json:"count,omitempty"`
	LastIncrementBy string                 `protobuf:"bytes,2,opt,name=last_increment_by,json=lastIncrementBy,proto3" json:"last_increment_by,omitempty"`
	LastIncrementAt string                 `protobuf:"bytes,3,opt,name=last_increment_at,json=lastIncrementAt,proto3" json:"last_increment_at,omitempty"`
	NextFinalizeAt  string                 `protobuf:"bytes,4,opt,name=next_finalize_at,json=nextFinalizeAt,proto3" json:"next_finalize_at,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CounterGetResponse) Reset() {
	*x = CounterGetResponse{}
	mi := &file_goagen_v1_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CounterGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterGetResponse) ProtoMessage() {}

func (x *CounterGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_v1_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterGetResponse.ProtoReflect.Descriptor instead.
func (*CounterGetResponse) Descriptor() ([]byte, []int) {
	return file_goagen_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *CounterGetResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CounterGetResponse) GetLastIncrementBy() string {
	if x != nil {
		return x.LastIncrementBy
	}
	return ""
}

func (x *CounterGetResponse) GetLastIncrementAt() string {
	if x != nil {
		return x.LastIncrementAt
	}
	return ""
}

func (x *CounterGetResponse) GetNextFinalizeAt() string {
	if x != nil {
		return x.NextFinalizeAt
	}
	return ""
}

type CounterIncrementRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CounterIncrementRequest) Reset() {
	*x = CounterIncrementRequest{}
	mi := &file_goagen_v1_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CounterIncrementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterIncrementRequest) ProtoMessage() {}

func (x *CounterIncrementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_v1_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterIncrementRequest.ProtoReflect.Descriptor instead.
func (*CounterIncrementRequest) Descriptor() ([]byte, []int) {
	return file_goagen_v1_api_proto_rawDescGZIP(), []int{4}
}

type CounterIncrementResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Count           int32                  `protobuf:"zigzag32,1,opt,name=count,proto3" json:"count,omitempty"`
	LastIncrementBy string                 `protobuf:"bytes,2,opt,name=last_increment_by,json=lastIncrementBy,proto3" json:"last_increment_by,omitempty"`
	LastIncrementAt string                 `protobuf:"bytes,3,opt,name=last_increment_at,json=lastIncrementAt,proto3" json:"last_increment_at,omitempty"`
	NextFinalizeAt  string                 `protobuf:"bytes,4,opt,name=next_finalize_at,json=nextFinalizeAt,proto3" json:"next_finalize_at,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CounterIncrementResponse) Reset() {
	*x = CounterIncrementResponse{}
	mi := &file_goagen_v1_api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CounterIncrementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterIncrementResponse) ProtoMessage() {}

func (x *CounterIncrementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_v1_api_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterIncrementResponse.ProtoReflect.Descriptor instead.
func (*CounterIncrementResponse) Descriptor() ([]byte, []int) {
	return file_goagen_v1_api_proto_rawDescGZIP(), []int{5}
}

func (x *CounterIncrementResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CounterIncrementResponse) GetLastIncrementBy() string {
	if x != nil {
		return x.LastIncrementBy
	}
	return ""
}

func (x *CounterIncrementResponse) GetLastIncrementAt() string {
	if x != nil {
		return x.LastIncrementAt
	}
	return ""
}

func (x *CounterIncrementResponse) GetNextFinalizeAt() string {
	if x != nil {
		return x.NextFinalizeAt
	}
	return ""
}

var File_goagen_v1_api_proto protoreflect.FileDescriptor

var file_goagen_v1_api_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x67, 0x6f, 0x61, 0x67, 0x65, 0x6e, 0x5f, 0x76, 0x31, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x51, 0x0a, 0x10, 0x41, 0x75,
	0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x29, 0x0a,
	0x11, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xac, 0x01,
	0x0a, 0x12, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x11, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69,
	0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65,
	0x78, 0x74, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x41, 0x74, 0x22, 0x19, 0x0a, 0x17,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb2, 0x01, 0x0a, 0x18, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x11, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69,
	0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65,
	0x78, 0x74, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x41, 0x74, 0x32, 0xd1, 0x01, 0x0a,
	0x03, 0x41, 0x50, 0x49, 0x12, 0x3a, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x65, 0x74, 0x12, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4f, 0x0a, 0x10, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x08, 0x5a, 0x06, 0x2f, 0x61, 0x70, 0x69, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_goagen_v1_api_proto_rawDescOnce sync.Once
	file_goagen_v1_api_proto_rawDescData []byte
)

func file_goagen_v1_api_proto_rawDescGZIP() []byte {
	file_goagen_v1_api_proto_rawDescOnce.Do(func() {
		file_goagen_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_goagen_v1_api_proto_rawDesc), len(file_goagen_v1_api_proto_rawDesc)))
	})
	return file_goagen_v1_api_proto_rawDescData
}

var file_goagen_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_goagen_v1_api_proto_goTypes = []any{
	(*AuthTokenRequest)(nil),         // 0: api.AuthTokenRequest
	(*AuthTokenResponse)(nil),        // 1: api.AuthTokenResponse
	(*CounterGetRequest)(nil),        // 2: api.CounterGetRequest
	(*CounterGetResponse)(nil),       // 3: api.CounterGetResponse
	(*CounterIncrementRequest)(nil),  // 4: api.CounterIncrementRequest
	(*CounterIncrementResponse)(nil), // 5: api.CounterIncrementResponse
}
var file_goagen_v1_api_proto_depIdxs = []int32{
	0, // 0: api.API.AuthToken:input_type -> api.AuthTokenRequest
	2, // 1: api.API.CounterGet:input_type -> api.CounterGetRequest
	4, // 2: api.API.CounterIncrement:input_type -> api.CounterIncrementRequest
	1, // 3: api.API.AuthToken:output_type -> api.AuthTokenResponse
	3, // 4: api.API.CounterGet:output_type -> api.CounterGetResponse
	5, // 5: api.API.CounterIncrement:output_type -> api.CounterIncrementResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goagen_v1_api_proto_init() }
func file_goagen_v1_api_proto_init() {
	if File_goagen_v1_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_goagen_v1_api_proto_rawDesc), len(file_goagen_v1_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goagen_v1_api_proto_goTypes,
		DependencyIndexes: file_goagen_v1_api_proto_depIdxs,
		MessageInfos:      file_goagen_v1_api_proto_msgTypes,
	}.Build()
	File_goagen_v1_api_proto = out.File
	file_goagen_v1_api_proto_goTypes = nil
	file_goagen_v1_api_proto_depIdxs = nil
}
