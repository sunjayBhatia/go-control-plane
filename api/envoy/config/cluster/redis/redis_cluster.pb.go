// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: envoy/config/cluster/redis/redis_cluster.proto

package redis

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// [#next-free-field: 7]
type RedisClusterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Interval between successive topology refresh requests. If not set, this defaults to 5s.
	ClusterRefreshRate *duration.Duration `protobuf:"bytes,1,opt,name=cluster_refresh_rate,json=clusterRefreshRate,proto3" json:"cluster_refresh_rate,omitempty"`
	// Timeout for topology refresh request. If not set, this defaults to 3s.
	ClusterRefreshTimeout *duration.Duration `protobuf:"bytes,2,opt,name=cluster_refresh_timeout,json=clusterRefreshTimeout,proto3" json:"cluster_refresh_timeout,omitempty"`
	// The minimum interval that must pass after triggering a topology refresh request before a new
	// request can possibly be triggered again. Any errors received during one of these
	// time intervals are ignored. If not set, this defaults to 5s.
	RedirectRefreshInterval *duration.Duration `protobuf:"bytes,3,opt,name=redirect_refresh_interval,json=redirectRefreshInterval,proto3" json:"redirect_refresh_interval,omitempty"`
	// The number of redirection errors that must be received before
	// triggering a topology refresh request. If not set, this defaults to 5.
	// If this is set to 0, topology refresh after redirect is disabled.
	RedirectRefreshThreshold *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=redirect_refresh_threshold,json=redirectRefreshThreshold,proto3" json:"redirect_refresh_threshold,omitempty"`
	// The number of failures that must be received before triggering a topology refresh request.
	// If not set, this defaults to 0, which disables the topology refresh due to failure.
	FailureRefreshThreshold uint32 `protobuf:"varint,5,opt,name=failure_refresh_threshold,json=failureRefreshThreshold,proto3" json:"failure_refresh_threshold,omitempty"`
	// The number of hosts became degraded or unhealthy before triggering a topology refresh request.
	// If not set, this defaults to 0, which disables the topology refresh due to degraded or
	// unhealthy host.
	HostDegradedRefreshThreshold uint32 `protobuf:"varint,6,opt,name=host_degraded_refresh_threshold,json=hostDegradedRefreshThreshold,proto3" json:"host_degraded_refresh_threshold,omitempty"`
}

func (x *RedisClusterConfig) Reset() {
	*x = RedisClusterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_cluster_redis_redis_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisClusterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisClusterConfig) ProtoMessage() {}

func (x *RedisClusterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_cluster_redis_redis_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisClusterConfig.ProtoReflect.Descriptor instead.
func (*RedisClusterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_cluster_redis_redis_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *RedisClusterConfig) GetClusterRefreshRate() *duration.Duration {
	if x != nil {
		return x.ClusterRefreshRate
	}
	return nil
}

func (x *RedisClusterConfig) GetClusterRefreshTimeout() *duration.Duration {
	if x != nil {
		return x.ClusterRefreshTimeout
	}
	return nil
}

func (x *RedisClusterConfig) GetRedirectRefreshInterval() *duration.Duration {
	if x != nil {
		return x.RedirectRefreshInterval
	}
	return nil
}

func (x *RedisClusterConfig) GetRedirectRefreshThreshold() *wrappers.UInt32Value {
	if x != nil {
		return x.RedirectRefreshThreshold
	}
	return nil
}

func (x *RedisClusterConfig) GetFailureRefreshThreshold() uint32 {
	if x != nil {
		return x.FailureRefreshThreshold
	}
	return 0
}

func (x *RedisClusterConfig) GetHostDegradedRefreshThreshold() uint32 {
	if x != nil {
		return x.HostDegradedRefreshThreshold
	}
	return 0
}

var File_envoy_config_cluster_redis_redis_cluster_proto protoreflect.FileDescriptor

var file_envoy_config_cluster_redis_redis_cluster_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2f, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x03, 0x0a, 0x12, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x55, 0x0a, 0x14, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a, 0x00, 0x52, 0x12,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x5b, 0x0a, 0x17, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a, 0x00, 0x52, 0x15, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x55, 0x0a, 0x19, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x17, 0x72,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x5a, 0x0a, 0x1a, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x18, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x12, 0x3a, 0x0a, 0x19, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x45,
	0x0a, 0x1f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x5f,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1c, 0x68, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x42, 0x8a, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01,
	0x0a, 0x28, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x42, 0x11, 0x52, 0x65, 0x64, 0x69,
	0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_cluster_redis_redis_cluster_proto_rawDescOnce sync.Once
	file_envoy_config_cluster_redis_redis_cluster_proto_rawDescData = file_envoy_config_cluster_redis_redis_cluster_proto_rawDesc
)

func file_envoy_config_cluster_redis_redis_cluster_proto_rawDescGZIP() []byte {
	file_envoy_config_cluster_redis_redis_cluster_proto_rawDescOnce.Do(func() {
		file_envoy_config_cluster_redis_redis_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_cluster_redis_redis_cluster_proto_rawDescData)
	})
	return file_envoy_config_cluster_redis_redis_cluster_proto_rawDescData
}

var file_envoy_config_cluster_redis_redis_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_cluster_redis_redis_cluster_proto_goTypes = []interface{}{
	(*RedisClusterConfig)(nil),   // 0: envoy.config.cluster.redis.RedisClusterConfig
	(*duration.Duration)(nil),    // 1: google.protobuf.Duration
	(*wrappers.UInt32Value)(nil), // 2: google.protobuf.UInt32Value
}
var file_envoy_config_cluster_redis_redis_cluster_proto_depIdxs = []int32{
	1, // 0: envoy.config.cluster.redis.RedisClusterConfig.cluster_refresh_rate:type_name -> google.protobuf.Duration
	1, // 1: envoy.config.cluster.redis.RedisClusterConfig.cluster_refresh_timeout:type_name -> google.protobuf.Duration
	1, // 2: envoy.config.cluster.redis.RedisClusterConfig.redirect_refresh_interval:type_name -> google.protobuf.Duration
	2, // 3: envoy.config.cluster.redis.RedisClusterConfig.redirect_refresh_threshold:type_name -> google.protobuf.UInt32Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_config_cluster_redis_redis_cluster_proto_init() }
func file_envoy_config_cluster_redis_redis_cluster_proto_init() {
	if File_envoy_config_cluster_redis_redis_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_cluster_redis_redis_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisClusterConfig); i {
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
			RawDescriptor: file_envoy_config_cluster_redis_redis_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_cluster_redis_redis_cluster_proto_goTypes,
		DependencyIndexes: file_envoy_config_cluster_redis_redis_cluster_proto_depIdxs,
		MessageInfos:      file_envoy_config_cluster_redis_redis_cluster_proto_msgTypes,
	}.Build()
	File_envoy_config_cluster_redis_redis_cluster_proto = out.File
	file_envoy_config_cluster_redis_redis_cluster_proto_rawDesc = nil
	file_envoy_config_cluster_redis_redis_cluster_proto_goTypes = nil
	file_envoy_config_cluster_redis_redis_cluster_proto_depIdxs = nil
}
