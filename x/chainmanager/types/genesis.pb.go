// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chainmanager/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the chainmanager module's genesis state.
type GenesisState struct {
	// heimdall.types.chainmanager.x.params
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d42f085946771f, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

type ChainParams struct {
	BorChainID            string `protobuf:"bytes,1,opt,name=bor_chain_id,json=borChainId,proto3" json:"bor_chain_id,omitempty"`
	MaticTokenAddress     string `protobuf:"bytes,2,opt,name=matic_token_address,json=maticTokenAddress,proto3" json:"matic_token_address,omitempty"`
	StakingManagerAddress string `protobuf:"bytes,3,opt,name=staking_manager_address,json=stakingManagerAddress,proto3" json:"staking_manager_address,omitempty"`
	SlashManagerAddress   string `protobuf:"bytes,4,opt,name=slash_manager_address,json=slashManagerAddress,proto3" json:"slash_manager_address,omitempty"`
	RootChainAddress      string `protobuf:"bytes,5,opt,name=root_chain_address,json=rootChainAddress,proto3" json:"root_chain_address,omitempty"`
	StakingInfoAddress    string `protobuf:"bytes,6,opt,name=staking_info_address,json=stakingInfoAddress,proto3" json:"staking_info_address,omitempty"`
	StateSenderAddress    string `protobuf:"bytes,7,opt,name=state_sender_address,json=stateSenderAddress,proto3" json:"state_sender_address,omitempty"`
	StateReceiverAddress  string `protobuf:"bytes,8,opt,name=state_receiver_address,json=stateReceiverAddress,proto3" json:"state_receiver_address,omitempty"`
	ValidatorSetAddress   string `protobuf:"bytes,9,opt,name=validator_set_address,json=validatorSetAddress,proto3" json:"validator_set_address,omitempty"`
}

func (m *ChainParams) Reset()      { *m = ChainParams{} }
func (*ChainParams) ProtoMessage() {}
func (*ChainParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d42f085946771f, []int{1}
}
func (m *ChainParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainParams.Merge(m, src)
}
func (m *ChainParams) XXX_Size() int {
	return m.Size()
}
func (m *ChainParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainParams.DiscardUnknown(m)
}

var xxx_messageInfo_ChainParams proto.InternalMessageInfo

type Params struct {
	MainchainTxConfirmations  uint64       `protobuf:"varint,1,opt,name=mainchain_tx_confirmations,json=mainchainTxConfirmations,proto3" json:"mainchain_tx_confirmations,omitempty"`
	MaticchainTxConfirmations uint64       `protobuf:"varint,2,opt,name=maticchain_tx_confirmations,json=maticchainTxConfirmations,proto3" json:"maticchain_tx_confirmations,omitempty"`
	ChainParams               *ChainParams `protobuf:"bytes,3,opt,name=chain_params,json=chainParams,proto3" json:"chain_params,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d42f085946771f, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "chainmanager.v1beta1.GenesisState")
	proto.RegisterType((*ChainParams)(nil), "chainmanager.v1beta1.ChainParams")
	proto.RegisterType((*Params)(nil), "chainmanager.v1beta1.Params")
}

func init() {
	proto.RegisterFile("chainmanager/v1beta1/genesis.proto", fileDescriptor_57d42f085946771f)
}

var fileDescriptor_57d42f085946771f = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x93, 0x75, 0x94, 0xcd, 0xad, 0x10, 0x64, 0x1d, 0x94, 0x81, 0x52, 0xe8, 0x89, 0x03,
	0x4a, 0xb6, 0x31, 0x38, 0x20, 0x84, 0x44, 0x57, 0x09, 0xed, 0x80, 0x40, 0xe9, 0x4e, 0x5c, 0x22,
	0x27, 0x71, 0x53, 0xab, 0x8d, 0x5d, 0xd9, 0xa6, 0x94, 0x6f, 0xc0, 0x91, 0x23, 0xc7, 0x7d, 0x1c,
	0x8e, 0x3b, 0x21, 0x4e, 0x08, 0xd2, 0x2f, 0x82, 0xf2, 0xec, 0xa4, 0xd9, 0xd8, 0x6e, 0x6e, 0xff,
	0xbf, 0x9f, 0xf5, 0x5e, 0xde, 0x33, 0xea, 0xc7, 0x13, 0x4c, 0x59, 0x86, 0x19, 0x4e, 0x89, 0xf0,
	0x17, 0x07, 0x11, 0x51, 0xf8, 0xc0, 0x4f, 0x09, 0x23, 0x92, 0x4a, 0x6f, 0x2e, 0xb8, 0xe2, 0x4e,
	0xa7, 0xce, 0x78, 0x86, 0xd9, 0xeb, 0xa4, 0x3c, 0xe5, 0x00, 0xf8, 0xc5, 0x49, 0xb3, 0xfd, 0x21,
	0x6a, 0xbf, 0xd5, 0xf2, 0x48, 0x61, 0x45, 0x9c, 0x23, 0xd4, 0x9c, 0x63, 0x81, 0x33, 0xd9, 0xb5,
	0x1f, 0xd9, 0x4f, 0x5a, 0x87, 0x0f, 0xbd, 0xab, 0x2e, 0xf3, 0x3e, 0x00, 0x13, 0x18, 0xb6, 0xff,
	0xb7, 0x81, 0x5a, 0xc7, 0x05, 0xa7, 0xff, 0x77, 0xf6, 0x51, 0x3b, 0xe2, 0x22, 0x04, 0x35, 0xa4,
	0x09, 0xdc, 0xb5, 0x3d, 0xb8, 0x95, 0xff, 0xee, 0xa1, 0x01, 0x17, 0x40, 0x9e, 0x0c, 0x03, 0x14,
	0x95, 0xe7, 0xc4, 0xf1, 0xd0, 0x4e, 0x86, 0x15, 0x8d, 0x43, 0xc5, 0xa7, 0x84, 0x85, 0x38, 0x49,
	0x04, 0x91, 0xb2, 0xbb, 0x51, 0x88, 0xc1, 0x1d, 0x88, 0x4e, 0x8b, 0xe4, 0x8d, 0x0e, 0x9c, 0x17,
	0xe8, 0x9e, 0x54, 0x78, 0x4a, 0x59, 0x1a, 0x9a, 0xda, 0x2a, 0xa7, 0x01, 0xce, 0xae, 0x89, 0xdf,
	0xe9, 0xb4, 0xf4, 0x0e, 0xd1, 0xae, 0x9c, 0x61, 0x39, 0xf9, 0xcf, 0xda, 0x04, 0x6b, 0x07, 0xc2,
	0x4b, 0xce, 0x53, 0xe4, 0x08, 0xce, 0x95, 0x69, 0xa7, 0x14, 0x6e, 0x80, 0x70, 0xbb, 0x48, 0xa0,
	0x89, 0x92, 0xde, 0x47, 0x9d, 0xb2, 0x32, 0xca, 0xc6, 0xbc, 0xe2, 0x9b, 0xc0, 0x3b, 0x26, 0x3b,
	0x61, 0x63, 0x7e, 0xd1, 0x50, 0x24, 0x94, 0x84, 0x25, 0xb5, 0x92, 0x6e, 0x56, 0x86, 0x22, 0x23,
	0x88, 0x4a, 0xe3, 0x08, 0xdd, 0xd5, 0x86, 0x20, 0x31, 0xa1, 0x8b, 0x9a, 0xb3, 0x05, 0x8e, 0xbe,
	0x2f, 0x30, 0x61, 0xad, 0xf7, 0x05, 0x9e, 0xd1, 0x04, 0x2b, 0x2e, 0x42, 0x49, 0x54, 0x25, 0x6d,
	0xeb, 0xde, 0xab, 0x70, 0x44, 0x94, 0x71, 0x5e, 0x6e, 0x7d, 0x3d, 0xeb, 0x59, 0xdf, 0xcf, 0x7a,
	0x56, 0xff, 0xa7, 0x8d, 0x9a, 0x66, 0xbc, 0xaf, 0xd0, 0x5e, 0x86, 0x29, 0xd3, 0xdf, 0x43, 0x2d,
	0xc3, 0x98, 0xb3, 0x31, 0x15, 0xc5, 0x90, 0x38, 0xd3, 0x8b, 0xb3, 0x19, 0x74, 0x2b, 0xe2, 0x74,
	0x79, 0x5c, 0xcf, 0x9d, 0xd7, 0xe8, 0x01, 0xcc, 0xf3, 0x1a, 0x7d, 0x03, 0xf4, 0xfb, 0x6b, 0xe4,
	0xb2, 0x3f, 0x44, 0x6d, 0xad, 0x9a, 0x45, 0x6d, 0xc0, 0xa2, 0x3e, 0xbe, 0x7a, 0x51, 0x6b, 0x5b,
	0x19, 0xb4, 0xe2, 0xf5, 0x8f, 0x75, 0x63, 0x83, 0xf7, 0x3f, 0x72, 0xd7, 0x3e, 0xcf, 0x5d, 0xfb,
	0x4f, 0xee, 0xda, 0xdf, 0x56, 0xae, 0x75, 0xbe, 0x72, 0xad, 0x5f, 0x2b, 0xd7, 0xfa, 0xf8, 0x3c,
	0xa5, 0x6a, 0xf2, 0x29, 0xf2, 0x62, 0x9e, 0xf9, 0x50, 0x0f, 0x23, 0xea, 0x33, 0x17, 0x53, 0x7f,
	0x42, 0x68, 0x96, 0xe0, 0xd9, 0xcc, 0x5f, 0xfa, 0x17, 0xde, 0xa3, 0xfa, 0x32, 0x27, 0x32, 0x6a,
	0xc2, 0xd3, 0x7a, 0xf6, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x85, 0x6c, 0x13, 0xfc, 0xac, 0x03, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChainParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorSetAddress) > 0 {
		i -= len(m.ValidatorSetAddress)
		copy(dAtA[i:], m.ValidatorSetAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorSetAddress)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.StateReceiverAddress) > 0 {
		i -= len(m.StateReceiverAddress)
		copy(dAtA[i:], m.StateReceiverAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.StateReceiverAddress)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.StateSenderAddress) > 0 {
		i -= len(m.StateSenderAddress)
		copy(dAtA[i:], m.StateSenderAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.StateSenderAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.StakingInfoAddress) > 0 {
		i -= len(m.StakingInfoAddress)
		copy(dAtA[i:], m.StakingInfoAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.StakingInfoAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.RootChainAddress) > 0 {
		i -= len(m.RootChainAddress)
		copy(dAtA[i:], m.RootChainAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.RootChainAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SlashManagerAddress) > 0 {
		i -= len(m.SlashManagerAddress)
		copy(dAtA[i:], m.SlashManagerAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.SlashManagerAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.StakingManagerAddress) > 0 {
		i -= len(m.StakingManagerAddress)
		copy(dAtA[i:], m.StakingManagerAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.StakingManagerAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MaticTokenAddress) > 0 {
		i -= len(m.MaticTokenAddress)
		copy(dAtA[i:], m.MaticTokenAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.MaticTokenAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BorChainID) > 0 {
		i -= len(m.BorChainID)
		copy(dAtA[i:], m.BorChainID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.BorChainID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ChainParams != nil {
		{
			size, err := m.ChainParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MaticchainTxConfirmations != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaticchainTxConfirmations))
		i--
		dAtA[i] = 0x10
	}
	if m.MainchainTxConfirmations != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MainchainTxConfirmations))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *ChainParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BorChainID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.MaticTokenAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.StakingManagerAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.SlashManagerAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.RootChainAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.StakingInfoAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.StateSenderAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.StateReceiverAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ValidatorSetAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MainchainTxConfirmations != 0 {
		n += 1 + sovGenesis(uint64(m.MainchainTxConfirmations))
	}
	if m.MaticchainTxConfirmations != 0 {
		n += 1 + sovGenesis(uint64(m.MaticchainTxConfirmations))
	}
	if m.ChainParams != nil {
		l = m.ChainParams.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ChainParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChainParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BorChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaticTokenAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaticTokenAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingManagerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakingManagerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashManagerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SlashManagerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootChainAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootChainAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingInfoAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakingInfoAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateSenderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateSenderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateReceiverAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateReceiverAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSetAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorSetAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainchainTxConfirmations", wireType)
			}
			m.MainchainTxConfirmations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MainchainTxConfirmations |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaticchainTxConfirmations", wireType)
			}
			m.MaticchainTxConfirmations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaticchainTxConfirmations |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ChainParams == nil {
				m.ChainParams = &ChainParams{}
			}
			if err := m.ChainParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
