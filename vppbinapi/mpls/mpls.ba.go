// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0
//  VPP:              23.02-rc0~230-g1d9780a43~b3131
// source: /usr/share/vpp/api/core/mpls.api.json

// Package mpls contains generated bindings for API file mpls.api.
//
// Contents:
//
//	 3 structs
//	16 messages
package mpls

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	fib_types "mygit.com/go-vpp-sr/vppbinapi/fib_types"
	interface_types "mygit.com/go-vpp-sr/vppbinapi/interface_types"
	ip_types "mygit.com/go-vpp-sr/vppbinapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "mpls"
	APIVersion = "1.1.1"
	VersionCrc = 0x46824f02
)

// MplsRoute defines type 'mpls_route'.
type MplsRoute struct {
	MrTableID     uint32              `binapi:"u32,name=mr_table_id" json:"mr_table_id,omitempty"`
	MrLabel       uint32              `binapi:"u32,name=mr_label" json:"mr_label,omitempty"`
	MrEos         uint8               `binapi:"u8,name=mr_eos" json:"mr_eos,omitempty"`
	MrEosProto    uint8               `binapi:"u8,name=mr_eos_proto" json:"mr_eos_proto,omitempty"`
	MrIsMulticast bool                `binapi:"bool,name=mr_is_multicast" json:"mr_is_multicast,omitempty"`
	MrNPaths      uint8               `binapi:"u8,name=mr_n_paths" json:"-"`
	MrPaths       []fib_types.FibPath `binapi:"fib_path[mr_n_paths],name=mr_paths" json:"mr_paths,omitempty"`
}

// MplsTable defines type 'mpls_table'.
type MplsTable struct {
	MtTableID uint32 `binapi:"u32,name=mt_table_id" json:"mt_table_id,omitempty"`
	MtName    string `binapi:"string[64],name=mt_name" json:"mt_name,omitempty"`
}

// MplsTunnel defines type 'mpls_tunnel'.
type MplsTunnel struct {
	MtSwIfIndex   interface_types.InterfaceIndex `binapi:"interface_index,name=mt_sw_if_index" json:"mt_sw_if_index,omitempty"`
	MtTunnelIndex uint32                         `binapi:"u32,name=mt_tunnel_index" json:"mt_tunnel_index,omitempty"`
	MtL2Only      bool                           `binapi:"bool,name=mt_l2_only" json:"mt_l2_only,omitempty"`
	MtIsMulticast bool                           `binapi:"bool,name=mt_is_multicast" json:"mt_is_multicast,omitempty"`
	MtTag         string                         `binapi:"string[64],name=mt_tag" json:"mt_tag,omitempty"`
	MtNPaths      uint8                          `binapi:"u8,name=mt_n_paths" json:"-"`
	MtPaths       []fib_types.FibPath            `binapi:"fib_path[mt_n_paths],name=mt_paths" json:"mt_paths,omitempty"`
}

// MplsIPBindUnbind defines message 'mpls_ip_bind_unbind'.
type MplsIPBindUnbind struct {
	MbMplsTableID uint32          `binapi:"u32,name=mb_mpls_table_id" json:"mb_mpls_table_id,omitempty"`
	MbLabel       uint32          `binapi:"u32,name=mb_label" json:"mb_label,omitempty"`
	MbIPTableID   uint32          `binapi:"u32,name=mb_ip_table_id" json:"mb_ip_table_id,omitempty"`
	MbIsBind      bool            `binapi:"bool,name=mb_is_bind" json:"mb_is_bind,omitempty"`
	MbPrefix      ip_types.Prefix `binapi:"prefix,name=mb_prefix" json:"mb_prefix,omitempty"`
}

func (m *MplsIPBindUnbind) Reset()               { *m = MplsIPBindUnbind{} }
func (*MplsIPBindUnbind) GetMessageName() string { return "mpls_ip_bind_unbind" }
func (*MplsIPBindUnbind) GetCrcString() string   { return "c7533b32" }
func (*MplsIPBindUnbind) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MplsIPBindUnbind) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.MbMplsTableID
	size += 4      // m.MbLabel
	size += 4      // m.MbIPTableID
	size += 1      // m.MbIsBind
	size += 1      // m.MbPrefix.Address.Af
	size += 1 * 16 // m.MbPrefix.Address.Un
	size += 1      // m.MbPrefix.Len
	return size
}
func (m *MplsIPBindUnbind) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.MbMplsTableID)
	buf.EncodeUint32(m.MbLabel)
	buf.EncodeUint32(m.MbIPTableID)
	buf.EncodeBool(m.MbIsBind)
	buf.EncodeUint8(uint8(m.MbPrefix.Address.Af))
	buf.EncodeBytes(m.MbPrefix.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.MbPrefix.Len)
	return buf.Bytes(), nil
}
func (m *MplsIPBindUnbind) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.MbMplsTableID = buf.DecodeUint32()
	m.MbLabel = buf.DecodeUint32()
	m.MbIPTableID = buf.DecodeUint32()
	m.MbIsBind = buf.DecodeBool()
	m.MbPrefix.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.MbPrefix.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.MbPrefix.Len = buf.DecodeUint8()
	return nil
}

// MplsIPBindUnbindReply defines message 'mpls_ip_bind_unbind_reply'.
type MplsIPBindUnbindReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *MplsIPBindUnbindReply) Reset()               { *m = MplsIPBindUnbindReply{} }
func (*MplsIPBindUnbindReply) GetMessageName() string { return "mpls_ip_bind_unbind_reply" }
func (*MplsIPBindUnbindReply) GetCrcString() string   { return "e8d4e804" }
func (*MplsIPBindUnbindReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MplsIPBindUnbindReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *MplsIPBindUnbindReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *MplsIPBindUnbindReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// MplsRouteAddDel defines message 'mpls_route_add_del'.
type MplsRouteAddDel struct {
	MrIsAdd       bool      `binapi:"bool,name=mr_is_add" json:"mr_is_add,omitempty"`
	MrIsMultipath bool      `binapi:"bool,name=mr_is_multipath" json:"mr_is_multipath,omitempty"`
	MrRoute       MplsRoute `binapi:"mpls_route,name=mr_route" json:"mr_route,omitempty"`
}

func (m *MplsRouteAddDel) Reset()               { *m = MplsRouteAddDel{} }
func (*MplsRouteAddDel) GetMessageName() string { return "mpls_route_add_del" }
func (*MplsRouteAddDel) GetCrcString() string   { return "8e1d1e07" }
func (*MplsRouteAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MplsRouteAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.MrIsAdd
	size += 1 // m.MrIsMultipath
	size += 4 // m.MrRoute.MrTableID
	size += 4 // m.MrRoute.MrLabel
	size += 1 // m.MrRoute.MrEos
	size += 1 // m.MrRoute.MrEosProto
	size += 1 // m.MrRoute.MrIsMulticast
	size += 1 // m.MrRoute.MrNPaths
	for j2 := 0; j2 < len(m.MrRoute.MrPaths); j2++ {
		var s2 fib_types.FibPath
		_ = s2
		if j2 < len(m.MrRoute.MrPaths) {
			s2 = m.MrRoute.MrPaths[j2]
		}
		size += 4      // s2.SwIfIndex
		size += 4      // s2.TableID
		size += 4      // s2.RpfID
		size += 1      // s2.Weight
		size += 1      // s2.Preference
		size += 4      // s2.Type
		size += 4      // s2.Flags
		size += 4      // s2.Proto
		size += 1 * 16 // s2.Nh.Address
		size += 4      // s2.Nh.ViaLabel
		size += 4      // s2.Nh.ObjID
		size += 4      // s2.Nh.ClassifyTableIndex
		size += 1      // s2.NLabels
		for j3 := 0; j3 < 16; j3++ {
			size += 1 // s2.LabelStack[j3].IsUniform
			size += 4 // s2.LabelStack[j3].Label
			size += 1 // s2.LabelStack[j3].TTL
			size += 1 // s2.LabelStack[j3].Exp
		}
	}
	return size
}
func (m *MplsRouteAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.MrIsAdd)
	buf.EncodeBool(m.MrIsMultipath)
	buf.EncodeUint32(m.MrRoute.MrTableID)
	buf.EncodeUint32(m.MrRoute.MrLabel)
	buf.EncodeUint8(m.MrRoute.MrEos)
	buf.EncodeUint8(m.MrRoute.MrEosProto)
	buf.EncodeBool(m.MrRoute.MrIsMulticast)
	buf.EncodeUint8(uint8(len(m.MrRoute.MrPaths)))
	for j1 := 0; j1 < len(m.MrRoute.MrPaths); j1++ {
		var v1 fib_types.FibPath // MrPaths
		if j1 < len(m.MrRoute.MrPaths) {
			v1 = m.MrRoute.MrPaths[j1]
		}
		buf.EncodeUint32(v1.SwIfIndex)
		buf.EncodeUint32(v1.TableID)
		buf.EncodeUint32(v1.RpfID)
		buf.EncodeUint8(v1.Weight)
		buf.EncodeUint8(v1.Preference)
		buf.EncodeUint32(uint32(v1.Type))
		buf.EncodeUint32(uint32(v1.Flags))
		buf.EncodeUint32(uint32(v1.Proto))
		buf.EncodeBytes(v1.Nh.Address.XXX_UnionData[:], 16)
		buf.EncodeUint32(v1.Nh.ViaLabel)
		buf.EncodeUint32(v1.Nh.ObjID)
		buf.EncodeUint32(v1.Nh.ClassifyTableIndex)
		buf.EncodeUint8(v1.NLabels)
		for j2 := 0; j2 < 16; j2++ {
			buf.EncodeUint8(v1.LabelStack[j2].IsUniform)
			buf.EncodeUint32(v1.LabelStack[j2].Label)
			buf.EncodeUint8(v1.LabelStack[j2].TTL)
			buf.EncodeUint8(v1.LabelStack[j2].Exp)
		}
	}
	return buf.Bytes(), nil
}
func (m *MplsRouteAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.MrIsAdd = buf.DecodeBool()
	m.MrIsMultipath = buf.DecodeBool()
	m.MrRoute.MrTableID = buf.DecodeUint32()
	m.MrRoute.MrLabel = buf.DecodeUint32()
	m.MrRoute.MrEos = buf.DecodeUint8()
	m.MrRoute.MrEosProto = buf.DecodeUint8()
	m.MrRoute.MrIsMulticast = buf.DecodeBool()
	m.MrRoute.MrNPaths = buf.DecodeUint8()
	m.MrRoute.MrPaths = make([]fib_types.FibPath, m.MrRoute.MrNPaths)
	for j1 := 0; j1 < len(m.MrRoute.MrPaths); j1++ {
		m.MrRoute.MrPaths[j1].SwIfIndex = buf.DecodeUint32()
		m.MrRoute.MrPaths[j1].TableID = buf.DecodeUint32()
		m.MrRoute.MrPaths[j1].RpfID = buf.DecodeUint32()
		m.MrRoute.MrPaths[j1].Weight = buf.DecodeUint8()
		m.MrRoute.MrPaths[j1].Preference = buf.DecodeUint8()
		m.MrRoute.MrPaths[j1].Type = fib_types.FibPathType(buf.DecodeUint32())
		m.MrRoute.MrPaths[j1].Flags = fib_types.FibPathFlags(buf.DecodeUint32())
		m.MrRoute.MrPaths[j1].Proto = fib_types.FibPathNhProto(buf.DecodeUint32())
		copy(m.MrRoute.MrPaths[j1].Nh.Address.XXX_UnionData[:], buf.DecodeBytes(16))
		m.MrRoute.MrPaths[j1].Nh.ViaLabel = buf.DecodeUint32()
		m.MrRoute.MrPaths[j1].Nh.ObjID = buf.DecodeUint32()
		m.MrRoute.MrPaths[j1].Nh.ClassifyTableIndex = buf.DecodeUint32()
		m.MrRoute.MrPaths[j1].NLabels = buf.DecodeUint8()
		for j2 := 0; j2 < 16; j2++ {
			m.MrRoute.MrPaths[j1].LabelStack[j2].IsUniform = buf.DecodeUint8()
			m.MrRoute.MrPaths[j1].LabelStack[j2].Label = buf.DecodeUint32()
			m.MrRoute.MrPaths[j1].LabelStack[j2].TTL = buf.DecodeUint8()
			m.MrRoute.MrPaths[j1].LabelStack[j2].Exp = buf.DecodeUint8()
		}
	}
	return nil
}

// MplsRouteAddDelReply defines message 'mpls_route_add_del_reply'.
type MplsRouteAddDelReply struct {
	Retval     int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	StatsIndex uint32 `binapi:"u32,name=stats_index" json:"stats_index,omitempty"`
}

func (m *MplsRouteAddDelReply) Reset()               { *m = MplsRouteAddDelReply{} }
func (*MplsRouteAddDelReply) GetMessageName() string { return "mpls_route_add_del_reply" }
func (*MplsRouteAddDelReply) GetCrcString() string   { return "1992deab" }
func (*MplsRouteAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MplsRouteAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.StatsIndex
	return size
}
func (m *MplsRouteAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.StatsIndex)
	return buf.Bytes(), nil
}
func (m *MplsRouteAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.StatsIndex = buf.DecodeUint32()
	return nil
}

// MplsRouteDetails defines message 'mpls_route_details'.
type MplsRouteDetails struct {
	MrRoute MplsRoute `binapi:"mpls_route,name=mr_route" json:"mr_route,omitempty"`
}

func (m *MplsRouteDetails) Reset()               { *m = MplsRouteDetails{} }
func (*MplsRouteDetails) GetMessageName() string { return "mpls_route_details" }
func (*MplsRouteDetails) GetCrcString() string   { return "9b5043dc" }
func (*MplsRouteDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MplsRouteDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.MrRoute.MrTableID
	size += 4 // m.MrRoute.MrLabel
	size += 1 // m.MrRoute.MrEos
	size += 1 // m.MrRoute.MrEosProto
	size += 1 // m.MrRoute.MrIsMulticast
	size += 1 // m.MrRoute.MrNPaths
	for j2 := 0; j2 < len(m.MrRoute.MrPaths); j2++ {
		var s2 fib_types.FibPath
		_ = s2
		if j2 < len(m.MrRoute.MrPaths) {
			s2 = m.MrRoute.MrPaths[j2]
		}
		size += 4      // s2.SwIfIndex
		size += 4      // s2.TableID
		size += 4      // s2.RpfID
		size += 1      // s2.Weight
		size += 1      // s2.Preference
		size += 4      // s2.Type
		size += 4      // s2.Flags
		size += 4      // s2.Proto
		size += 1 * 16 // s2.Nh.Address
		size += 4      // s2.Nh.ViaLabel
		size += 4      // s2.Nh.ObjID
		size += 4      // s2.Nh.ClassifyTableIndex
		size += 1      // s2.NLabels
		for j3 := 0; j3 < 16; j3++ {
			size += 1 // s2.LabelStack[j3].IsUniform
			size += 4 // s2.LabelStack[j3].Label
			size += 1 // s2.LabelStack[j3].TTL
			size += 1 // s2.LabelStack[j3].Exp
		}
	}
	return size
}
func (m *MplsRouteDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.MrRoute.MrTableID)
	buf.EncodeUint32(m.MrRoute.MrLabel)
	buf.EncodeUint8(m.MrRoute.MrEos)
	buf.EncodeUint8(m.MrRoute.MrEosProto)
	buf.EncodeBool(m.MrRoute.MrIsMulticast)
	buf.EncodeUint8(uint8(len(m.MrRoute.MrPaths)))
	for j1 := 0; j1 < len(m.MrRoute.MrPaths); j1++ {
		var v1 fib_types.FibPath // MrPaths
		if j1 < len(m.MrRoute.MrPaths) {
			v1 = m.MrRoute.MrPaths[j1]
		}
		buf.EncodeUint32(v1.SwIfIndex)
		buf.EncodeUint32(v1.TableID)
		buf.EncodeUint32(v1.RpfID)
		buf.EncodeUint8(v1.Weight)
		buf.EncodeUint8(v1.Preference)
		buf.EncodeUint32(uint32(v1.Type))
		buf.EncodeUint32(uint32(v1.Flags))
		buf.EncodeUint32(uint32(v1.Proto))
		buf.EncodeBytes(v1.Nh.Address.XXX_UnionData[:], 16)
		buf.EncodeUint32(v1.Nh.ViaLabel)
		buf.EncodeUint32(v1.Nh.ObjID)
		buf.EncodeUint32(v1.Nh.ClassifyTableIndex)
		buf.EncodeUint8(v1.NLabels)
		for j2 := 0; j2 < 16; j2++ {
			buf.EncodeUint8(v1.LabelStack[j2].IsUniform)
			buf.EncodeUint32(v1.LabelStack[j2].Label)
			buf.EncodeUint8(v1.LabelStack[j2].TTL)
			buf.EncodeUint8(v1.LabelStack[j2].Exp)
		}
	}
	return buf.Bytes(), nil
}
func (m *MplsRouteDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.MrRoute.MrTableID = buf.DecodeUint32()
	m.MrRoute.MrLabel = buf.DecodeUint32()
	m.MrRoute.MrEos = buf.DecodeUint8()
	m.MrRoute.MrEosProto = buf.DecodeUint8()
	m.MrRoute.MrIsMulticast = buf.DecodeBool()
	m.MrRoute.MrNPaths = buf.DecodeUint8()
	m.MrRoute.MrPaths = make([]fib_types.FibPath, m.MrRoute.MrNPaths)
	for j1 := 0; j1 < len(m.MrRoute.MrPaths); j1++ {
		m.MrRoute.MrPaths[j1].SwIfIndex = buf.DecodeUint32()
		m.MrRoute.MrPaths[j1].TableID = buf.DecodeUint32()
		m.MrRoute.MrPaths[j1].RpfID = buf.DecodeUint32()
		m.MrRoute.MrPaths[j1].Weight = buf.DecodeUint8()
		m.MrRoute.MrPaths[j1].Preference = buf.DecodeUint8()
		m.MrRoute.MrPaths[j1].Type = fib_types.FibPathType(buf.DecodeUint32())
		m.MrRoute.MrPaths[j1].Flags = fib_types.FibPathFlags(buf.DecodeUint32())
		m.MrRoute.MrPaths[j1].Proto = fib_types.FibPathNhProto(buf.DecodeUint32())
		copy(m.MrRoute.MrPaths[j1].Nh.Address.XXX_UnionData[:], buf.DecodeBytes(16))
		m.MrRoute.MrPaths[j1].Nh.ViaLabel = buf.DecodeUint32()
		m.MrRoute.MrPaths[j1].Nh.ObjID = buf.DecodeUint32()
		m.MrRoute.MrPaths[j1].Nh.ClassifyTableIndex = buf.DecodeUint32()
		m.MrRoute.MrPaths[j1].NLabels = buf.DecodeUint8()
		for j2 := 0; j2 < 16; j2++ {
			m.MrRoute.MrPaths[j1].LabelStack[j2].IsUniform = buf.DecodeUint8()
			m.MrRoute.MrPaths[j1].LabelStack[j2].Label = buf.DecodeUint32()
			m.MrRoute.MrPaths[j1].LabelStack[j2].TTL = buf.DecodeUint8()
			m.MrRoute.MrPaths[j1].LabelStack[j2].Exp = buf.DecodeUint8()
		}
	}
	return nil
}

// MplsRouteDump defines message 'mpls_route_dump'.
type MplsRouteDump struct {
	Table MplsTable `binapi:"mpls_table,name=table" json:"table,omitempty"`
}

func (m *MplsRouteDump) Reset()               { *m = MplsRouteDump{} }
func (*MplsRouteDump) GetMessageName() string { return "mpls_route_dump" }
func (*MplsRouteDump) GetCrcString() string   { return "935fdefa" }
func (*MplsRouteDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MplsRouteDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4  // m.Table.MtTableID
	size += 64 // m.Table.MtName
	return size
}
func (m *MplsRouteDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Table.MtTableID)
	buf.EncodeString(m.Table.MtName, 64)
	return buf.Bytes(), nil
}
func (m *MplsRouteDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Table.MtTableID = buf.DecodeUint32()
	m.Table.MtName = buf.DecodeString(64)
	return nil
}

// MplsTableAddDel defines message 'mpls_table_add_del'.
type MplsTableAddDel struct {
	MtIsAdd bool      `binapi:"bool,name=mt_is_add,default=true" json:"mt_is_add,omitempty"`
	MtTable MplsTable `binapi:"mpls_table,name=mt_table" json:"mt_table,omitempty"`
}

func (m *MplsTableAddDel) Reset()               { *m = MplsTableAddDel{} }
func (*MplsTableAddDel) GetMessageName() string { return "mpls_table_add_del" }
func (*MplsTableAddDel) GetCrcString() string   { return "57817512" }
func (*MplsTableAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MplsTableAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1  // m.MtIsAdd
	size += 4  // m.MtTable.MtTableID
	size += 64 // m.MtTable.MtName
	return size
}
func (m *MplsTableAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.MtIsAdd)
	buf.EncodeUint32(m.MtTable.MtTableID)
	buf.EncodeString(m.MtTable.MtName, 64)
	return buf.Bytes(), nil
}
func (m *MplsTableAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.MtIsAdd = buf.DecodeBool()
	m.MtTable.MtTableID = buf.DecodeUint32()
	m.MtTable.MtName = buf.DecodeString(64)
	return nil
}

// MplsTableAddDelReply defines message 'mpls_table_add_del_reply'.
type MplsTableAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *MplsTableAddDelReply) Reset()               { *m = MplsTableAddDelReply{} }
func (*MplsTableAddDelReply) GetMessageName() string { return "mpls_table_add_del_reply" }
func (*MplsTableAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*MplsTableAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MplsTableAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *MplsTableAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *MplsTableAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// MplsTableDetails defines message 'mpls_table_details'.
type MplsTableDetails struct {
	MtTable MplsTable `binapi:"mpls_table,name=mt_table" json:"mt_table,omitempty"`
}

func (m *MplsTableDetails) Reset()               { *m = MplsTableDetails{} }
func (*MplsTableDetails) GetMessageName() string { return "mpls_table_details" }
func (*MplsTableDetails) GetCrcString() string   { return "f03ecdc8" }
func (*MplsTableDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MplsTableDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4  // m.MtTable.MtTableID
	size += 64 // m.MtTable.MtName
	return size
}
func (m *MplsTableDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.MtTable.MtTableID)
	buf.EncodeString(m.MtTable.MtName, 64)
	return buf.Bytes(), nil
}
func (m *MplsTableDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.MtTable.MtTableID = buf.DecodeUint32()
	m.MtTable.MtName = buf.DecodeString(64)
	return nil
}

// MplsTableDump defines message 'mpls_table_dump'.
type MplsTableDump struct{}

func (m *MplsTableDump) Reset()               { *m = MplsTableDump{} }
func (*MplsTableDump) GetMessageName() string { return "mpls_table_dump" }
func (*MplsTableDump) GetCrcString() string   { return "51077d14" }
func (*MplsTableDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MplsTableDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *MplsTableDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *MplsTableDump) Unmarshal(b []byte) error {
	return nil
}

// MplsTunnelAddDel defines message 'mpls_tunnel_add_del'.
type MplsTunnelAddDel struct {
	MtIsAdd  bool       `binapi:"bool,name=mt_is_add,default=true" json:"mt_is_add,omitempty"`
	MtTunnel MplsTunnel `binapi:"mpls_tunnel,name=mt_tunnel" json:"mt_tunnel,omitempty"`
}

func (m *MplsTunnelAddDel) Reset()               { *m = MplsTunnelAddDel{} }
func (*MplsTunnelAddDel) GetMessageName() string { return "mpls_tunnel_add_del" }
func (*MplsTunnelAddDel) GetCrcString() string   { return "44350ac1" }
func (*MplsTunnelAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MplsTunnelAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1  // m.MtIsAdd
	size += 4  // m.MtTunnel.MtSwIfIndex
	size += 4  // m.MtTunnel.MtTunnelIndex
	size += 1  // m.MtTunnel.MtL2Only
	size += 1  // m.MtTunnel.MtIsMulticast
	size += 64 // m.MtTunnel.MtTag
	size += 1  // m.MtTunnel.MtNPaths
	for j2 := 0; j2 < len(m.MtTunnel.MtPaths); j2++ {
		var s2 fib_types.FibPath
		_ = s2
		if j2 < len(m.MtTunnel.MtPaths) {
			s2 = m.MtTunnel.MtPaths[j2]
		}
		size += 4      // s2.SwIfIndex
		size += 4      // s2.TableID
		size += 4      // s2.RpfID
		size += 1      // s2.Weight
		size += 1      // s2.Preference
		size += 4      // s2.Type
		size += 4      // s2.Flags
		size += 4      // s2.Proto
		size += 1 * 16 // s2.Nh.Address
		size += 4      // s2.Nh.ViaLabel
		size += 4      // s2.Nh.ObjID
		size += 4      // s2.Nh.ClassifyTableIndex
		size += 1      // s2.NLabels
		for j3 := 0; j3 < 16; j3++ {
			size += 1 // s2.LabelStack[j3].IsUniform
			size += 4 // s2.LabelStack[j3].Label
			size += 1 // s2.LabelStack[j3].TTL
			size += 1 // s2.LabelStack[j3].Exp
		}
	}
	return size
}
func (m *MplsTunnelAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.MtIsAdd)
	buf.EncodeUint32(uint32(m.MtTunnel.MtSwIfIndex))
	buf.EncodeUint32(m.MtTunnel.MtTunnelIndex)
	buf.EncodeBool(m.MtTunnel.MtL2Only)
	buf.EncodeBool(m.MtTunnel.MtIsMulticast)
	buf.EncodeString(m.MtTunnel.MtTag, 64)
	buf.EncodeUint8(uint8(len(m.MtTunnel.MtPaths)))
	for j1 := 0; j1 < len(m.MtTunnel.MtPaths); j1++ {
		var v1 fib_types.FibPath // MtPaths
		if j1 < len(m.MtTunnel.MtPaths) {
			v1 = m.MtTunnel.MtPaths[j1]
		}
		buf.EncodeUint32(v1.SwIfIndex)
		buf.EncodeUint32(v1.TableID)
		buf.EncodeUint32(v1.RpfID)
		buf.EncodeUint8(v1.Weight)
		buf.EncodeUint8(v1.Preference)
		buf.EncodeUint32(uint32(v1.Type))
		buf.EncodeUint32(uint32(v1.Flags))
		buf.EncodeUint32(uint32(v1.Proto))
		buf.EncodeBytes(v1.Nh.Address.XXX_UnionData[:], 16)
		buf.EncodeUint32(v1.Nh.ViaLabel)
		buf.EncodeUint32(v1.Nh.ObjID)
		buf.EncodeUint32(v1.Nh.ClassifyTableIndex)
		buf.EncodeUint8(v1.NLabels)
		for j2 := 0; j2 < 16; j2++ {
			buf.EncodeUint8(v1.LabelStack[j2].IsUniform)
			buf.EncodeUint32(v1.LabelStack[j2].Label)
			buf.EncodeUint8(v1.LabelStack[j2].TTL)
			buf.EncodeUint8(v1.LabelStack[j2].Exp)
		}
	}
	return buf.Bytes(), nil
}
func (m *MplsTunnelAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.MtIsAdd = buf.DecodeBool()
	m.MtTunnel.MtSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.MtTunnel.MtTunnelIndex = buf.DecodeUint32()
	m.MtTunnel.MtL2Only = buf.DecodeBool()
	m.MtTunnel.MtIsMulticast = buf.DecodeBool()
	m.MtTunnel.MtTag = buf.DecodeString(64)
	m.MtTunnel.MtNPaths = buf.DecodeUint8()
	m.MtTunnel.MtPaths = make([]fib_types.FibPath, m.MtTunnel.MtNPaths)
	for j1 := 0; j1 < len(m.MtTunnel.MtPaths); j1++ {
		m.MtTunnel.MtPaths[j1].SwIfIndex = buf.DecodeUint32()
		m.MtTunnel.MtPaths[j1].TableID = buf.DecodeUint32()
		m.MtTunnel.MtPaths[j1].RpfID = buf.DecodeUint32()
		m.MtTunnel.MtPaths[j1].Weight = buf.DecodeUint8()
		m.MtTunnel.MtPaths[j1].Preference = buf.DecodeUint8()
		m.MtTunnel.MtPaths[j1].Type = fib_types.FibPathType(buf.DecodeUint32())
		m.MtTunnel.MtPaths[j1].Flags = fib_types.FibPathFlags(buf.DecodeUint32())
		m.MtTunnel.MtPaths[j1].Proto = fib_types.FibPathNhProto(buf.DecodeUint32())
		copy(m.MtTunnel.MtPaths[j1].Nh.Address.XXX_UnionData[:], buf.DecodeBytes(16))
		m.MtTunnel.MtPaths[j1].Nh.ViaLabel = buf.DecodeUint32()
		m.MtTunnel.MtPaths[j1].Nh.ObjID = buf.DecodeUint32()
		m.MtTunnel.MtPaths[j1].Nh.ClassifyTableIndex = buf.DecodeUint32()
		m.MtTunnel.MtPaths[j1].NLabels = buf.DecodeUint8()
		for j2 := 0; j2 < 16; j2++ {
			m.MtTunnel.MtPaths[j1].LabelStack[j2].IsUniform = buf.DecodeUint8()
			m.MtTunnel.MtPaths[j1].LabelStack[j2].Label = buf.DecodeUint32()
			m.MtTunnel.MtPaths[j1].LabelStack[j2].TTL = buf.DecodeUint8()
			m.MtTunnel.MtPaths[j1].LabelStack[j2].Exp = buf.DecodeUint8()
		}
	}
	return nil
}

// MplsTunnelAddDelReply defines message 'mpls_tunnel_add_del_reply'.
type MplsTunnelAddDelReply struct {
	Retval      int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex   interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	TunnelIndex uint32                         `binapi:"u32,name=tunnel_index" json:"tunnel_index,omitempty"`
}

func (m *MplsTunnelAddDelReply) Reset()               { *m = MplsTunnelAddDelReply{} }
func (*MplsTunnelAddDelReply) GetMessageName() string { return "mpls_tunnel_add_del_reply" }
func (*MplsTunnelAddDelReply) GetCrcString() string   { return "afb01472" }
func (*MplsTunnelAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MplsTunnelAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	size += 4 // m.TunnelIndex
	return size
}
func (m *MplsTunnelAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint32(m.TunnelIndex)
	return buf.Bytes(), nil
}
func (m *MplsTunnelAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.TunnelIndex = buf.DecodeUint32()
	return nil
}

// MplsTunnelDetails defines message 'mpls_tunnel_details'.
type MplsTunnelDetails struct {
	MtTunnel MplsTunnel `binapi:"mpls_tunnel,name=mt_tunnel" json:"mt_tunnel,omitempty"`
}

func (m *MplsTunnelDetails) Reset()               { *m = MplsTunnelDetails{} }
func (*MplsTunnelDetails) GetMessageName() string { return "mpls_tunnel_details" }
func (*MplsTunnelDetails) GetCrcString() string   { return "57118ae3" }
func (*MplsTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MplsTunnelDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4  // m.MtTunnel.MtSwIfIndex
	size += 4  // m.MtTunnel.MtTunnelIndex
	size += 1  // m.MtTunnel.MtL2Only
	size += 1  // m.MtTunnel.MtIsMulticast
	size += 64 // m.MtTunnel.MtTag
	size += 1  // m.MtTunnel.MtNPaths
	for j2 := 0; j2 < len(m.MtTunnel.MtPaths); j2++ {
		var s2 fib_types.FibPath
		_ = s2
		if j2 < len(m.MtTunnel.MtPaths) {
			s2 = m.MtTunnel.MtPaths[j2]
		}
		size += 4      // s2.SwIfIndex
		size += 4      // s2.TableID
		size += 4      // s2.RpfID
		size += 1      // s2.Weight
		size += 1      // s2.Preference
		size += 4      // s2.Type
		size += 4      // s2.Flags
		size += 4      // s2.Proto
		size += 1 * 16 // s2.Nh.Address
		size += 4      // s2.Nh.ViaLabel
		size += 4      // s2.Nh.ObjID
		size += 4      // s2.Nh.ClassifyTableIndex
		size += 1      // s2.NLabels
		for j3 := 0; j3 < 16; j3++ {
			size += 1 // s2.LabelStack[j3].IsUniform
			size += 4 // s2.LabelStack[j3].Label
			size += 1 // s2.LabelStack[j3].TTL
			size += 1 // s2.LabelStack[j3].Exp
		}
	}
	return size
}
func (m *MplsTunnelDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.MtTunnel.MtSwIfIndex))
	buf.EncodeUint32(m.MtTunnel.MtTunnelIndex)
	buf.EncodeBool(m.MtTunnel.MtL2Only)
	buf.EncodeBool(m.MtTunnel.MtIsMulticast)
	buf.EncodeString(m.MtTunnel.MtTag, 64)
	buf.EncodeUint8(uint8(len(m.MtTunnel.MtPaths)))
	for j1 := 0; j1 < len(m.MtTunnel.MtPaths); j1++ {
		var v1 fib_types.FibPath // MtPaths
		if j1 < len(m.MtTunnel.MtPaths) {
			v1 = m.MtTunnel.MtPaths[j1]
		}
		buf.EncodeUint32(v1.SwIfIndex)
		buf.EncodeUint32(v1.TableID)
		buf.EncodeUint32(v1.RpfID)
		buf.EncodeUint8(v1.Weight)
		buf.EncodeUint8(v1.Preference)
		buf.EncodeUint32(uint32(v1.Type))
		buf.EncodeUint32(uint32(v1.Flags))
		buf.EncodeUint32(uint32(v1.Proto))
		buf.EncodeBytes(v1.Nh.Address.XXX_UnionData[:], 16)
		buf.EncodeUint32(v1.Nh.ViaLabel)
		buf.EncodeUint32(v1.Nh.ObjID)
		buf.EncodeUint32(v1.Nh.ClassifyTableIndex)
		buf.EncodeUint8(v1.NLabels)
		for j2 := 0; j2 < 16; j2++ {
			buf.EncodeUint8(v1.LabelStack[j2].IsUniform)
			buf.EncodeUint32(v1.LabelStack[j2].Label)
			buf.EncodeUint8(v1.LabelStack[j2].TTL)
			buf.EncodeUint8(v1.LabelStack[j2].Exp)
		}
	}
	return buf.Bytes(), nil
}
func (m *MplsTunnelDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.MtTunnel.MtSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.MtTunnel.MtTunnelIndex = buf.DecodeUint32()
	m.MtTunnel.MtL2Only = buf.DecodeBool()
	m.MtTunnel.MtIsMulticast = buf.DecodeBool()
	m.MtTunnel.MtTag = buf.DecodeString(64)
	m.MtTunnel.MtNPaths = buf.DecodeUint8()
	m.MtTunnel.MtPaths = make([]fib_types.FibPath, m.MtTunnel.MtNPaths)
	for j1 := 0; j1 < len(m.MtTunnel.MtPaths); j1++ {
		m.MtTunnel.MtPaths[j1].SwIfIndex = buf.DecodeUint32()
		m.MtTunnel.MtPaths[j1].TableID = buf.DecodeUint32()
		m.MtTunnel.MtPaths[j1].RpfID = buf.DecodeUint32()
		m.MtTunnel.MtPaths[j1].Weight = buf.DecodeUint8()
		m.MtTunnel.MtPaths[j1].Preference = buf.DecodeUint8()
		m.MtTunnel.MtPaths[j1].Type = fib_types.FibPathType(buf.DecodeUint32())
		m.MtTunnel.MtPaths[j1].Flags = fib_types.FibPathFlags(buf.DecodeUint32())
		m.MtTunnel.MtPaths[j1].Proto = fib_types.FibPathNhProto(buf.DecodeUint32())
		copy(m.MtTunnel.MtPaths[j1].Nh.Address.XXX_UnionData[:], buf.DecodeBytes(16))
		m.MtTunnel.MtPaths[j1].Nh.ViaLabel = buf.DecodeUint32()
		m.MtTunnel.MtPaths[j1].Nh.ObjID = buf.DecodeUint32()
		m.MtTunnel.MtPaths[j1].Nh.ClassifyTableIndex = buf.DecodeUint32()
		m.MtTunnel.MtPaths[j1].NLabels = buf.DecodeUint8()
		for j2 := 0; j2 < 16; j2++ {
			m.MtTunnel.MtPaths[j1].LabelStack[j2].IsUniform = buf.DecodeUint8()
			m.MtTunnel.MtPaths[j1].LabelStack[j2].Label = buf.DecodeUint32()
			m.MtTunnel.MtPaths[j1].LabelStack[j2].TTL = buf.DecodeUint8()
			m.MtTunnel.MtPaths[j1].LabelStack[j2].Exp = buf.DecodeUint8()
		}
	}
	return nil
}

// MplsTunnelDump defines message 'mpls_tunnel_dump'.
type MplsTunnelDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index,default=4294967295" json:"sw_if_index,omitempty"`
}

func (m *MplsTunnelDump) Reset()               { *m = MplsTunnelDump{} }
func (*MplsTunnelDump) GetMessageName() string { return "mpls_tunnel_dump" }
func (*MplsTunnelDump) GetCrcString() string   { return "f9e6675e" }
func (*MplsTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MplsTunnelDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *MplsTunnelDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *MplsTunnelDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// SwInterfaceSetMplsEnable defines message 'sw_interface_set_mpls_enable'.
type SwInterfaceSetMplsEnable struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Enable    bool                           `binapi:"bool,name=enable,default=true" json:"enable,omitempty"`
}

func (m *SwInterfaceSetMplsEnable) Reset()               { *m = SwInterfaceSetMplsEnable{} }
func (*SwInterfaceSetMplsEnable) GetMessageName() string { return "sw_interface_set_mpls_enable" }
func (*SwInterfaceSetMplsEnable) GetCrcString() string   { return "ae6cfcfb" }
func (*SwInterfaceSetMplsEnable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceSetMplsEnable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.Enable
	return size
}
func (m *SwInterfaceSetMplsEnable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.Enable)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSetMplsEnable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Enable = buf.DecodeBool()
	return nil
}

// SwInterfaceSetMplsEnableReply defines message 'sw_interface_set_mpls_enable_reply'.
type SwInterfaceSetMplsEnableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SwInterfaceSetMplsEnableReply) Reset() { *m = SwInterfaceSetMplsEnableReply{} }
func (*SwInterfaceSetMplsEnableReply) GetMessageName() string {
	return "sw_interface_set_mpls_enable_reply"
}
func (*SwInterfaceSetMplsEnableReply) GetCrcString() string { return "e8d4e804" }
func (*SwInterfaceSetMplsEnableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceSetMplsEnableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SwInterfaceSetMplsEnableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSetMplsEnableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_mpls_binapi_init() }
func file_mpls_binapi_init() {
	api.RegisterMessage((*MplsIPBindUnbind)(nil), "mpls_ip_bind_unbind_c7533b32")
	api.RegisterMessage((*MplsIPBindUnbindReply)(nil), "mpls_ip_bind_unbind_reply_e8d4e804")
	api.RegisterMessage((*MplsRouteAddDel)(nil), "mpls_route_add_del_8e1d1e07")
	api.RegisterMessage((*MplsRouteAddDelReply)(nil), "mpls_route_add_del_reply_1992deab")
	api.RegisterMessage((*MplsRouteDetails)(nil), "mpls_route_details_9b5043dc")
	api.RegisterMessage((*MplsRouteDump)(nil), "mpls_route_dump_935fdefa")
	api.RegisterMessage((*MplsTableAddDel)(nil), "mpls_table_add_del_57817512")
	api.RegisterMessage((*MplsTableAddDelReply)(nil), "mpls_table_add_del_reply_e8d4e804")
	api.RegisterMessage((*MplsTableDetails)(nil), "mpls_table_details_f03ecdc8")
	api.RegisterMessage((*MplsTableDump)(nil), "mpls_table_dump_51077d14")
	api.RegisterMessage((*MplsTunnelAddDel)(nil), "mpls_tunnel_add_del_44350ac1")
	api.RegisterMessage((*MplsTunnelAddDelReply)(nil), "mpls_tunnel_add_del_reply_afb01472")
	api.RegisterMessage((*MplsTunnelDetails)(nil), "mpls_tunnel_details_57118ae3")
	api.RegisterMessage((*MplsTunnelDump)(nil), "mpls_tunnel_dump_f9e6675e")
	api.RegisterMessage((*SwInterfaceSetMplsEnable)(nil), "sw_interface_set_mpls_enable_ae6cfcfb")
	api.RegisterMessage((*SwInterfaceSetMplsEnableReply)(nil), "sw_interface_set_mpls_enable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*MplsIPBindUnbind)(nil),
		(*MplsIPBindUnbindReply)(nil),
		(*MplsRouteAddDel)(nil),
		(*MplsRouteAddDelReply)(nil),
		(*MplsRouteDetails)(nil),
		(*MplsRouteDump)(nil),
		(*MplsTableAddDel)(nil),
		(*MplsTableAddDelReply)(nil),
		(*MplsTableDetails)(nil),
		(*MplsTableDump)(nil),
		(*MplsTunnelAddDel)(nil),
		(*MplsTunnelAddDelReply)(nil),
		(*MplsTunnelDetails)(nil),
		(*MplsTunnelDump)(nil),
		(*SwInterfaceSetMplsEnable)(nil),
		(*SwInterfaceSetMplsEnableReply)(nil),
	}
}
