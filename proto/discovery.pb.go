// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: discovery.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ServiceInfoProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	ClientId string `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *ServiceInfoProto) Reset() {
	*x = ServiceInfoProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInfoProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInfoProto) ProtoMessage() {}

func (x *ServiceInfoProto) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInfoProto.ProtoReflect.Descriptor instead.
func (*ServiceInfoProto) Descriptor() ([]byte, []int) {
	return file_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceInfoProto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceInfoProto) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ServiceInfoProto) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type ServiceInterest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName []string `protobuf:"bytes,1,rep,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *ServiceInterest) Reset() {
	*x = ServiceInterest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInterest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInterest) ProtoMessage() {}

func (x *ServiceInterest) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInterest.ProtoReflect.Descriptor instead.
func (*ServiceInterest) Descriptor() ([]byte, []int) {
	return file_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceInterest) GetServiceName() []string {
	if x != nil {
		return x.ServiceName
	}
	return nil
}

type AgentStopped struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId string `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
}

func (x *AgentStopped) Reset() {
	*x = AgentStopped{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentStopped) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentStopped) ProtoMessage() {}

func (x *AgentStopped) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentStopped.ProtoReflect.Descriptor instead.
func (*AgentStopped) Descriptor() ([]byte, []int) {
	return file_discovery_proto_rawDescGZIP(), []int{2}
}

func (x *AgentStopped) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

type ServicesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*ServiceInfoProto `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ServicesList) Reset() {
	*x = ServicesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicesList) ProtoMessage() {}

func (x *ServicesList) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicesList.ProtoReflect.Descriptor instead.
func (*ServicesList) Descriptor() ([]byte, []int) {
	return file_discovery_proto_rawDescGZIP(), []int{3}
}

func (x *ServicesList) GetServices() []*ServiceInfoProto {
	if x != nil {
		return x.Services
	}
	return nil
}

var File_discovery_proto protoreflect.FileDescriptor

var file_discovery_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a,
	0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x28, 0x5a,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x74, 0x6f,
	0x62, 0x69, 0x74, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_discovery_proto_rawDescOnce sync.Once
	file_discovery_proto_rawDescData = file_discovery_proto_rawDesc
)

func file_discovery_proto_rawDescGZIP() []byte {
	file_discovery_proto_rawDescOnce.Do(func() {
		file_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_discovery_proto_rawDescData)
	})
	return file_discovery_proto_rawDescData
}

var file_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_discovery_proto_goTypes = []interface{}{
	(*ServiceInfoProto)(nil), // 0: proto.ServiceInfoProto
	(*ServiceInterest)(nil),  // 1: proto.ServiceInterest
	(*AgentStopped)(nil),     // 2: proto.AgentStopped
	(*ServicesList)(nil),     // 3: proto.ServicesList
}
var file_discovery_proto_depIdxs = []int32{
	0, // 0: proto.ServicesList.services:type_name -> proto.ServiceInfoProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_discovery_proto_init() }
func file_discovery_proto_init() {
	if File_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInfoProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInterest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_discovery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentStopped); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_discovery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicesList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_discovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_discovery_proto_goTypes,
		DependencyIndexes: file_discovery_proto_depIdxs,
		MessageInfos:      file_discovery_proto_msgTypes,
	}.Build()
	File_discovery_proto = out.File
	file_discovery_proto_rawDesc = nil
	file_discovery_proto_goTypes = nil
	file_discovery_proto_depIdxs = nil
}
