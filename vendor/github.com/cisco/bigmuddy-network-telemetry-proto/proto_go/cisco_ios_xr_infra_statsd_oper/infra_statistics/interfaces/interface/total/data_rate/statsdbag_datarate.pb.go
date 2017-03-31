// Code generated by protoc-gen-go.
// source: statsdbag_datarate.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_total_data_rate is a generated protocol buffer package.

It is generated from these files:
	statsdbag_datarate.proto

It has these top-level messages:
	StatsdbagDatarate_KEYS
	StatsdbagDatarate
*/
package cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_total_data_rate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Datarate information
type StatsdbagDatarate_KEYS struct {
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *StatsdbagDatarate_KEYS) Reset()                    { *m = StatsdbagDatarate_KEYS{} }
func (m *StatsdbagDatarate_KEYS) String() string            { return proto.CompactTextString(m) }
func (*StatsdbagDatarate_KEYS) ProtoMessage()               {}
func (*StatsdbagDatarate_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StatsdbagDatarate_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type StatsdbagDatarate struct {
	// Input data rate in 1000's of bps
	InputDataRate uint64 `protobuf:"varint,50,opt,name=input_data_rate,json=inputDataRate" json:"input_data_rate,omitempty"`
	// Input packets per second
	InputPacketRate uint64 `protobuf:"varint,51,opt,name=input_packet_rate,json=inputPacketRate" json:"input_packet_rate,omitempty"`
	// Output data rate in 1000's of bps
	OutputDataRate uint64 `protobuf:"varint,52,opt,name=output_data_rate,json=outputDataRate" json:"output_data_rate,omitempty"`
	// Output packets per second
	OutputPacketRate uint64 `protobuf:"varint,53,opt,name=output_packet_rate,json=outputPacketRate" json:"output_packet_rate,omitempty"`
	// Peak input data rate
	PeakInputDataRate uint64 `protobuf:"varint,54,opt,name=peak_input_data_rate,json=peakInputDataRate" json:"peak_input_data_rate,omitempty"`
	// Peak input packet rate
	PeakInputPacketRate uint64 `protobuf:"varint,55,opt,name=peak_input_packet_rate,json=peakInputPacketRate" json:"peak_input_packet_rate,omitempty"`
	// Peak output data rate
	PeakOutputDataRate uint64 `protobuf:"varint,56,opt,name=peak_output_data_rate,json=peakOutputDataRate" json:"peak_output_data_rate,omitempty"`
	// Peak output packet rate
	PeakOutputPacketRate uint64 `protobuf:"varint,57,opt,name=peak_output_packet_rate,json=peakOutputPacketRate" json:"peak_output_packet_rate,omitempty"`
	// Bandwidth (in kbps)
	Bandwidth uint32 `protobuf:"varint,58,opt,name=bandwidth" json:"bandwidth,omitempty"`
	// Number of 30-sec intervals less one
	LoadInterval uint32 `protobuf:"varint,59,opt,name=load_interval,json=loadInterval" json:"load_interval,omitempty"`
	// Output load as fraction of 255
	OutputLoad uint32 `protobuf:"varint,60,opt,name=output_load,json=outputLoad" json:"output_load,omitempty"`
	// Input load as fraction of 255
	InputLoad uint32 `protobuf:"varint,61,opt,name=input_load,json=inputLoad" json:"input_load,omitempty"`
	// Reliability coefficient
	Reliability uint32 `protobuf:"varint,62,opt,name=reliability" json:"reliability,omitempty"`
}

func (m *StatsdbagDatarate) Reset()                    { *m = StatsdbagDatarate{} }
func (m *StatsdbagDatarate) String() string            { return proto.CompactTextString(m) }
func (*StatsdbagDatarate) ProtoMessage()               {}
func (*StatsdbagDatarate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StatsdbagDatarate) GetInputDataRate() uint64 {
	if m != nil {
		return m.InputDataRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetInputPacketRate() uint64 {
	if m != nil {
		return m.InputPacketRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetOutputDataRate() uint64 {
	if m != nil {
		return m.OutputDataRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetOutputPacketRate() uint64 {
	if m != nil {
		return m.OutputPacketRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetPeakInputDataRate() uint64 {
	if m != nil {
		return m.PeakInputDataRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetPeakInputPacketRate() uint64 {
	if m != nil {
		return m.PeakInputPacketRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetPeakOutputDataRate() uint64 {
	if m != nil {
		return m.PeakOutputDataRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetPeakOutputPacketRate() uint64 {
	if m != nil {
		return m.PeakOutputPacketRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *StatsdbagDatarate) GetLoadInterval() uint32 {
	if m != nil {
		return m.LoadInterval
	}
	return 0
}

func (m *StatsdbagDatarate) GetOutputLoad() uint32 {
	if m != nil {
		return m.OutputLoad
	}
	return 0
}

func (m *StatsdbagDatarate) GetInputLoad() uint32 {
	if m != nil {
		return m.InputLoad
	}
	return 0
}

func (m *StatsdbagDatarate) GetReliability() uint32 {
	if m != nil {
		return m.Reliability
	}
	return 0
}

func init() {
	proto.RegisterType((*StatsdbagDatarate_KEYS)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.total.data_rate.statsdbag_datarate_KEYS")
	proto.RegisterType((*StatsdbagDatarate)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.total.data_rate.statsdbag_datarate")
}

func init() { proto.RegisterFile("statsdbag_datarate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4f, 0xab, 0xd3, 0x40,
	0x14, 0xc5, 0x09, 0x14, 0xa1, 0xb7, 0xb6, 0xda, 0xb1, 0xda, 0x2c, 0x14, 0x43, 0x45, 0x09, 0x22,
	0x11, 0xad, 0xf5, 0xbf, 0xe2, 0x42, 0x17, 0x45, 0x51, 0xa9, 0x6e, 0x5c, 0x0d, 0x37, 0xc9, 0xd4,
	0x37, 0x34, 0xcd, 0x84, 0xe4, 0xf6, 0xfd, 0xf9, 0x60, 0xef, 0xfb, 0x3d, 0xe6, 0x4e, 0x48, 0xa7,
	0xaf, 0xbb, 0x72, 0xce, 0xef, 0x9c, 0x33, 0x97, 0x06, 0xc2, 0x86, 0x90, 0x9a, 0x3c, 0xc5, 0xff,
	0x32, 0x47, 0xc2, 0x1a, 0x49, 0x25, 0x55, 0x6d, 0xc8, 0x88, 0xbf, 0x99, 0x6e, 0x32, 0x23, 0xb5,
	0x69, 0xe4, 0x79, 0x2d, 0x75, 0xb9, 0xae, 0x51, 0x3a, 0x58, 0x9a, 0x4a, 0xd5, 0xc9, 0x5e, 0xd1,
	0x0d, 0xe9, 0xac, 0x49, 0x74, 0x49, 0xaa, 0x5e, 0x63, 0xa6, 0xbc, 0x9f, 0x09, 0x19, 0xc2, 0x22,
	0xb1, 0xd5, 0xd2, 0x76, 0xcf, 0xbe, 0xc0, 0xf4, 0x78, 0x51, 0x7e, 0xff, 0xf6, 0xef, 0x8f, 0x78,
	0x0c, 0xa3, 0x2e, 0x27, 0x4b, 0xdc, 0xaa, 0x30, 0x88, 0x82, 0xb8, 0xbf, 0x1a, 0x76, 0xea, 0x4f,
	0xdc, 0xaa, 0xd9, 0x65, 0x0f, 0xc4, 0x71, 0x85, 0x78, 0x02, 0xb7, 0x74, 0x59, 0xed, 0x48, 0x76,
	0x5b, 0xe1, 0xcb, 0x28, 0x88, 0x7b, 0x36, 0x5e, 0xed, 0xe8, 0x2b, 0x12, 0xae, 0x2c, 0xf7, 0x14,
	0xc6, 0x8e, 0xab, 0x30, 0xdb, 0x28, 0x72, 0xe4, 0x9c, 0x49, 0x57, 0xf0, 0x9b, 0x75, 0x66, 0x63,
	0xb8, 0x6d, 0x76, 0x74, 0x58, 0xfa, 0x8a, 0xd1, 0x91, 0xd3, 0xbb, 0xd6, 0x67, 0x20, 0x5a, 0xd2,
	0xaf, 0x5d, 0x30, 0xdb, 0x76, 0x78, 0xbd, 0xcf, 0x61, 0x52, 0x29, 0xdc, 0xc8, 0xeb, 0x0f, 0x7e,
	0xcd, 0xfc, 0xd8, 0x7a, 0xcb, 0x83, 0x47, 0xcf, 0xe1, 0x9e, 0x17, 0xf0, 0x27, 0xde, 0x70, 0xe4,
	0x4e, 0x17, 0xf1, 0x56, 0x5e, 0xc0, 0x5d, 0x0e, 0x1d, 0x9d, 0xf0, 0x96, 0x33, 0xc2, 0x9a, 0xbf,
	0x0e, 0xcf, 0x58, 0xc0, 0xd4, 0x8f, 0xf8, 0x43, 0xef, 0x38, 0x34, 0xd9, 0x87, 0xbc, 0xa5, 0xfb,
	0xd0, 0x4f, 0xb1, 0xcc, 0xcf, 0x74, 0x4e, 0x27, 0xe1, 0xfb, 0x28, 0x88, 0x87, 0xab, 0xbd, 0x20,
	0x1e, 0xc1, 0xb0, 0x30, 0x98, 0x4b, 0xfe, 0x1b, 0x4f, 0xb1, 0x08, 0x3f, 0x30, 0x71, 0xd3, 0x8a,
	0xcb, 0x56, 0x13, 0x0f, 0x61, 0xd0, 0x8e, 0x5a, 0x39, 0xfc, 0xc8, 0x08, 0x38, 0xe9, 0x87, 0xc1,
	0x5c, 0x3c, 0x00, 0x70, 0xd7, 0xb3, 0xff, 0xc9, 0x8d, 0xb0, 0xc2, 0x76, 0x04, 0x83, 0x5a, 0x15,
	0x1a, 0x53, 0x5d, 0x68, 0xba, 0x08, 0x3f, 0xb3, 0xef, 0x4b, 0xe9, 0x0d, 0xfe, 0xac, 0xe7, 0x57,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x32, 0xe7, 0xa5, 0x61, 0xf2, 0x02, 0x00, 0x00,
}