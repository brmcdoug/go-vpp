// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0
//  VPP:              23.02-rc0~230-g1d9780a43~b3131
// source: /usr/share/vpp/api/plugins/dhcp6_ia_na_client_cp.api.json

// Package dhcp6_ia_na_client_cp contains generated bindings for API file dhcp6_ia_na_client_cp.api.
//
// Contents:
//
//	2 messages
package dhcp6_ia_na_client_cp

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "mygit.com/myproject/vppbinapi/interface_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "dhcp6_ia_na_client_cp"
	APIVersion = "1.0.1"
	VersionCrc = 0x6e8abdfb
)

// DHCP6ClientEnableDisable defines message 'dhcp6_client_enable_disable'.
type DHCP6ClientEnableDisable struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Enable    bool                           `binapi:"bool,name=enable" json:"enable,omitempty"`
}

func (m *DHCP6ClientEnableDisable) Reset()               { *m = DHCP6ClientEnableDisable{} }
func (*DHCP6ClientEnableDisable) GetMessageName() string { return "dhcp6_client_enable_disable" }
func (*DHCP6ClientEnableDisable) GetCrcString() string   { return "ae6cfcfb" }
func (*DHCP6ClientEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DHCP6ClientEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.Enable
	return size
}
func (m *DHCP6ClientEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.Enable)
	return buf.Bytes(), nil
}
func (m *DHCP6ClientEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Enable = buf.DecodeBool()
	return nil
}

// DHCP6ClientEnableDisableReply defines message 'dhcp6_client_enable_disable_reply'.
type DHCP6ClientEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *DHCP6ClientEnableDisableReply) Reset() { *m = DHCP6ClientEnableDisableReply{} }
func (*DHCP6ClientEnableDisableReply) GetMessageName() string {
	return "dhcp6_client_enable_disable_reply"
}
func (*DHCP6ClientEnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*DHCP6ClientEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DHCP6ClientEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *DHCP6ClientEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *DHCP6ClientEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_dhcp6_ia_na_client_cp_binapi_init() }
func file_dhcp6_ia_na_client_cp_binapi_init() {
	api.RegisterMessage((*DHCP6ClientEnableDisable)(nil), "dhcp6_client_enable_disable_ae6cfcfb")
	api.RegisterMessage((*DHCP6ClientEnableDisableReply)(nil), "dhcp6_client_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*DHCP6ClientEnableDisable)(nil),
		(*DHCP6ClientEnableDisableReply)(nil),
	}
}