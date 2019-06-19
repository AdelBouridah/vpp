// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: /usr/share/vpp/api/core/ipsec.api.json

/*
 Package ipsec is a generated from VPP binary API module 'ipsec'.

 It contains following objects:
	 14 services
	  7 enums
	  2 aliases
	  8 types
	  1 union
	 28 messages
*/
package ipsec

import api "git.fd.io/govpp.git/api"
import struc "github.com/lunixbochs/struc"
import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// Services represents VPP binary API services:
type Services interface {
	DumpIpsecBackend(*IpsecBackendDump) ([]*IpsecBackendDetails, error)
	DumpIpsecSa(*IpsecSaDump) ([]*IpsecSaDetails, error)
	DumpIpsecSpd(*IpsecSpdDump) ([]*IpsecSpdDetails, error)
	DumpIpsecSpdInterface(*IpsecSpdInterfaceDump) ([]*IpsecSpdInterfaceDetails, error)
	DumpIpsecSpds(*IpsecSpdsDump) ([]*IpsecSpdsDetails, error)
	IpsecInterfaceAddDelSpd(*IpsecInterfaceAddDelSpd) (*IpsecInterfaceAddDelSpdReply, error)
	IpsecSaSetKey(*IpsecSaSetKey) (*IpsecSaSetKeyReply, error)
	IpsecSadEntryAddDel(*IpsecSadEntryAddDel) (*IpsecSadEntryAddDelReply, error)
	IpsecSelectBackend(*IpsecSelectBackend) (*IpsecSelectBackendReply, error)
	IpsecSpdAddDel(*IpsecSpdAddDel) (*IpsecSpdAddDelReply, error)
	IpsecSpdEntryAddDel(*IpsecSpdEntryAddDel) (*IpsecSpdEntryAddDelReply, error)
	IpsecTunnelIfAddDel(*IpsecTunnelIfAddDel) (*IpsecTunnelIfAddDelReply, error)
	IpsecTunnelIfSetKey(*IpsecTunnelIfSetKey) (*IpsecTunnelIfSetKeyReply, error)
	IpsecTunnelIfSetSa(*IpsecTunnelIfSetSa) (*IpsecTunnelIfSetSaReply, error)
}

/* Enums */

// AddressFamily represents VPP binary API enum 'address_family':
type AddressFamily uint32

const (
	ADDRESS_IP4 AddressFamily = 0
	ADDRESS_IP6 AddressFamily = 1
)

// IPProto represents VPP binary API enum 'ip_proto':
type IPProto uint32

const (
	IP_API_PROTO_TCP IPProto = 6
	IP_API_PROTO_UDP IPProto = 17
)

// IpsecCryptoAlg represents VPP binary API enum 'ipsec_crypto_alg':
type IpsecCryptoAlg uint32

const (
	IPSEC_API_CRYPTO_ALG_NONE        IpsecCryptoAlg = 0
	IPSEC_API_CRYPTO_ALG_AES_CBC_128 IpsecCryptoAlg = 1
	IPSEC_API_CRYPTO_ALG_AES_CBC_192 IpsecCryptoAlg = 2
	IPSEC_API_CRYPTO_ALG_AES_CBC_256 IpsecCryptoAlg = 3
	IPSEC_API_CRYPTO_ALG_AES_CTR_128 IpsecCryptoAlg = 4
	IPSEC_API_CRYPTO_ALG_AES_CTR_192 IpsecCryptoAlg = 5
	IPSEC_API_CRYPTO_ALG_AES_CTR_256 IpsecCryptoAlg = 6
	IPSEC_API_CRYPTO_ALG_AES_GCM_128 IpsecCryptoAlg = 7
	IPSEC_API_CRYPTO_ALG_AES_GCM_192 IpsecCryptoAlg = 8
	IPSEC_API_CRYPTO_ALG_AES_GCM_256 IpsecCryptoAlg = 9
	IPSEC_API_CRYPTO_ALG_DES_CBC     IpsecCryptoAlg = 10
	IPSEC_API_CRYPTO_ALG_3DES_CBC    IpsecCryptoAlg = 11
)

// IpsecIntegAlg represents VPP binary API enum 'ipsec_integ_alg':
type IpsecIntegAlg uint32

const (
	IPSEC_API_INTEG_ALG_NONE        IpsecIntegAlg = 0
	IPSEC_API_INTEG_ALG_MD5_96      IpsecIntegAlg = 1
	IPSEC_API_INTEG_ALG_SHA1_96     IpsecIntegAlg = 2
	IPSEC_API_INTEG_ALG_SHA_256_96  IpsecIntegAlg = 3
	IPSEC_API_INTEG_ALG_SHA_256_128 IpsecIntegAlg = 4
	IPSEC_API_INTEG_ALG_SHA_384_192 IpsecIntegAlg = 5
	IPSEC_API_INTEG_ALG_SHA_512_256 IpsecIntegAlg = 6
)

// IpsecProto represents VPP binary API enum 'ipsec_proto':
type IpsecProto uint32

const (
	IPSEC_API_PROTO_ESP IpsecProto = 1
	IPSEC_API_PROTO_AH  IpsecProto = 2
)

// IpsecSadFlags represents VPP binary API enum 'ipsec_sad_flags':
type IpsecSadFlags uint32

const (
	IPSEC_API_SAD_FLAG_NONE            IpsecSadFlags = 0
	IPSEC_API_SAD_FLAG_USE_ESN         IpsecSadFlags = 1
	IPSEC_API_SAD_FLAG_USE_ANTI_REPLAY IpsecSadFlags = 2
	IPSEC_API_SAD_FLAG_IS_TUNNEL       IpsecSadFlags = 4
	IPSEC_API_SAD_FLAG_IS_TUNNEL_V6    IpsecSadFlags = 8
	IPSEC_API_SAD_FLAG_UDP_ENCAP       IpsecSadFlags = 16
)

// IpsecSpdAction represents VPP binary API enum 'ipsec_spd_action':
type IpsecSpdAction uint32

const (
	IPSEC_API_SPD_ACTION_BYPASS  IpsecSpdAction = 0
	IPSEC_API_SPD_ACTION_DISCARD IpsecSpdAction = 1
	IPSEC_API_SPD_ACTION_RESOLVE IpsecSpdAction = 2
	IPSEC_API_SPD_ACTION_PROTECT IpsecSpdAction = 3
)

/* Aliases */

// IP4Address represents VPP binary API alias 'ip4_address':
type IP4Address [4]uint8

// IP6Address represents VPP binary API alias 'ip6_address':
type IP6Address [16]uint8

/* Types */

// Address represents VPP binary API type 'address':
type Address struct {
	Af AddressFamily
	Un AddressUnion
}

func (*Address) GetTypeName() string {
	return "address"
}
func (*Address) GetCrcString() string {
	return "09f11671"
}

// IP4Prefix represents VPP binary API type 'ip4_prefix':
type IP4Prefix struct {
	Prefix IP4Address
	Len    uint8
}

func (*IP4Prefix) GetTypeName() string {
	return "ip4_prefix"
}
func (*IP4Prefix) GetCrcString() string {
	return "ea8dc11d"
}

// IP6Prefix represents VPP binary API type 'ip6_prefix':
type IP6Prefix struct {
	Prefix IP6Address
	Len    uint8
}

func (*IP6Prefix) GetTypeName() string {
	return "ip6_prefix"
}
func (*IP6Prefix) GetCrcString() string {
	return "779fd64f"
}

// IpsecSadEntry represents VPP binary API type 'ipsec_sad_entry':
type IpsecSadEntry struct {
	SadID              uint32
	Spi                uint32
	Protocol           IpsecProto
	CryptoAlgorithm    IpsecCryptoAlg
	CryptoKey          Key
	IntegrityAlgorithm IpsecIntegAlg
	IntegrityKey       Key
	Flags              IpsecSadFlags
	TunnelSrc          Address
	TunnelDst          Address
	TxTableID          uint32
	Salt               uint32
}

func (*IpsecSadEntry) GetTypeName() string {
	return "ipsec_sad_entry"
}
func (*IpsecSadEntry) GetCrcString() string {
	return "559c6abb"
}

// IpsecSpdEntry represents VPP binary API type 'ipsec_spd_entry':
type IpsecSpdEntry struct {
	SpdID              uint32
	Priority           int32
	IsOutbound         uint8
	SaID               uint32
	Policy             IpsecSpdAction
	Protocol           uint8
	RemoteAddressStart Address
	RemoteAddressStop  Address
	LocalAddressStart  Address
	LocalAddressStop   Address
	RemotePortStart    uint16
	RemotePortStop     uint16
	LocalPortStart     uint16
	LocalPortStop      uint16
}

func (*IpsecSpdEntry) GetTypeName() string {
	return "ipsec_spd_entry"
}
func (*IpsecSpdEntry) GetCrcString() string {
	return "876fdb2c"
}

// Key represents VPP binary API type 'key':
type Key struct {
	Length uint8
	Data   []byte `struc:"[128]byte"`
}

func (*Key) GetTypeName() string {
	return "key"
}
func (*Key) GetCrcString() string {
	return "f3d0c4fd"
}

// Mprefix represents VPP binary API type 'mprefix':
type Mprefix struct {
	Af               AddressFamily
	GrpAddressLength uint16
	GrpAddress       AddressUnion
	SrcAddress       AddressUnion
}

func (*Mprefix) GetTypeName() string {
	return "mprefix"
}
func (*Mprefix) GetCrcString() string {
	return "1c4cba05"
}

// Prefix represents VPP binary API type 'prefix':
type Prefix struct {
	Address       Address
	AddressLength uint8
}

func (*Prefix) GetTypeName() string {
	return "prefix"
}
func (*Prefix) GetCrcString() string {
	return "0403aebc"
}

/* Unions */

// AddressUnion represents VPP binary API union 'address_union':
type AddressUnion struct {
	Union_data [16]byte
}

func (*AddressUnion) GetTypeName() string {
	return "address_union"
}
func (*AddressUnion) GetCrcString() string {
	return "d68a2fb4"
}

func AddressUnionIP4(a IP4Address) (u AddressUnion) {
	u.SetIP4(a)
	return
}
func (u *AddressUnion) SetIP4(a IP4Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.Union_data[:], b.Bytes())
}
func (u *AddressUnion) GetIP4() (a IP4Address) {
	var b = bytes.NewReader(u.Union_data[:])
	struc.Unpack(b, &a)
	return
}

func AddressUnionIP6(a IP6Address) (u AddressUnion) {
	u.SetIP6(a)
	return
}
func (u *AddressUnion) SetIP6(a IP6Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.Union_data[:], b.Bytes())
}
func (u *AddressUnion) GetIP6() (a IP6Address) {
	var b = bytes.NewReader(u.Union_data[:])
	struc.Unpack(b, &a)
	return
}

/* Messages */

// IpsecBackendDetails represents VPP binary API message 'ipsec_backend_details':
type IpsecBackendDetails struct {
	Name     []byte `struc:"[128]byte"`
	Protocol IpsecProto
	Index    uint8
	Active   uint8
}

func (*IpsecBackendDetails) GetMessageName() string {
	return "ipsec_backend_details"
}
func (*IpsecBackendDetails) GetCrcString() string {
	return "3341f485"
}
func (*IpsecBackendDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpsecBackendDump represents VPP binary API message 'ipsec_backend_dump':
type IpsecBackendDump struct{}

func (*IpsecBackendDump) GetMessageName() string {
	return "ipsec_backend_dump"
}
func (*IpsecBackendDump) GetCrcString() string {
	return "51077d14"
}
func (*IpsecBackendDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpsecInterfaceAddDelSpd represents VPP binary API message 'ipsec_interface_add_del_spd':
type IpsecInterfaceAddDelSpd struct {
	IsAdd     uint8
	SwIfIndex uint32
	SpdID     uint32
}

func (*IpsecInterfaceAddDelSpd) GetMessageName() string {
	return "ipsec_interface_add_del_spd"
}
func (*IpsecInterfaceAddDelSpd) GetCrcString() string {
	return "1e3b8286"
}
func (*IpsecInterfaceAddDelSpd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpsecInterfaceAddDelSpdReply represents VPP binary API message 'ipsec_interface_add_del_spd_reply':
type IpsecInterfaceAddDelSpdReply struct {
	Retval int32
}

func (*IpsecInterfaceAddDelSpdReply) GetMessageName() string {
	return "ipsec_interface_add_del_spd_reply"
}
func (*IpsecInterfaceAddDelSpdReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IpsecInterfaceAddDelSpdReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpsecSaDetails represents VPP binary API message 'ipsec_sa_details':
type IpsecSaDetails struct {
	Entry          IpsecSadEntry
	SwIfIndex      uint32
	Salt           uint32
	SeqOutbound    uint64
	LastSeqInbound uint64
	ReplayWindow   uint64
	TotalDataSize  uint64
}

func (*IpsecSaDetails) GetMessageName() string {
	return "ipsec_sa_details"
}
func (*IpsecSaDetails) GetCrcString() string {
	return "62f0eab8"
}
func (*IpsecSaDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpsecSaDump represents VPP binary API message 'ipsec_sa_dump':
type IpsecSaDump struct {
	SaID uint32
}

func (*IpsecSaDump) GetMessageName() string {
	return "ipsec_sa_dump"
}
func (*IpsecSaDump) GetCrcString() string {
	return "2076c2f4"
}
func (*IpsecSaDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpsecSaSetKey represents VPP binary API message 'ipsec_sa_set_key':
type IpsecSaSetKey struct {
	SaID         uint32
	CryptoKey    Key
	IntegrityKey Key
}

func (*IpsecSaSetKey) GetMessageName() string {
	return "ipsec_sa_set_key"
}
func (*IpsecSaSetKey) GetCrcString() string {
	return "f407f496"
}
func (*IpsecSaSetKey) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpsecSaSetKeyReply represents VPP binary API message 'ipsec_sa_set_key_reply':
type IpsecSaSetKeyReply struct {
	Retval int32
}

func (*IpsecSaSetKeyReply) GetMessageName() string {
	return "ipsec_sa_set_key_reply"
}
func (*IpsecSaSetKeyReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IpsecSaSetKeyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpsecSadEntryAddDel represents VPP binary API message 'ipsec_sad_entry_add_del':
type IpsecSadEntryAddDel struct {
	IsAdd uint8
	Entry IpsecSadEntry
}

func (*IpsecSadEntryAddDel) GetMessageName() string {
	return "ipsec_sad_entry_add_del"
}
func (*IpsecSadEntryAddDel) GetCrcString() string {
	return "05747d5b"
}
func (*IpsecSadEntryAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpsecSadEntryAddDelReply represents VPP binary API message 'ipsec_sad_entry_add_del_reply':
type IpsecSadEntryAddDelReply struct {
	Retval    int32
	StatIndex uint32
}

func (*IpsecSadEntryAddDelReply) GetMessageName() string {
	return "ipsec_sad_entry_add_del_reply"
}
func (*IpsecSadEntryAddDelReply) GetCrcString() string {
	return "9ffac24b"
}
func (*IpsecSadEntryAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpsecSelectBackend represents VPP binary API message 'ipsec_select_backend':
type IpsecSelectBackend struct {
	Protocol IpsecProto
	Index    uint8
}

func (*IpsecSelectBackend) GetMessageName() string {
	return "ipsec_select_backend"
}
func (*IpsecSelectBackend) GetCrcString() string {
	return "b36bcff3"
}
func (*IpsecSelectBackend) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpsecSelectBackendReply represents VPP binary API message 'ipsec_select_backend_reply':
type IpsecSelectBackendReply struct {
	Retval int32
}

func (*IpsecSelectBackendReply) GetMessageName() string {
	return "ipsec_select_backend_reply"
}
func (*IpsecSelectBackendReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IpsecSelectBackendReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpsecSpdAddDel represents VPP binary API message 'ipsec_spd_add_del':
type IpsecSpdAddDel struct {
	IsAdd uint8
	SpdID uint32
}

func (*IpsecSpdAddDel) GetMessageName() string {
	return "ipsec_spd_add_del"
}
func (*IpsecSpdAddDel) GetCrcString() string {
	return "9ffdf5da"
}
func (*IpsecSpdAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpsecSpdAddDelReply represents VPP binary API message 'ipsec_spd_add_del_reply':
type IpsecSpdAddDelReply struct {
	Retval int32
}

func (*IpsecSpdAddDelReply) GetMessageName() string {
	return "ipsec_spd_add_del_reply"
}
func (*IpsecSpdAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IpsecSpdAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpsecSpdDetails represents VPP binary API message 'ipsec_spd_details':
type IpsecSpdDetails struct {
	Entry IpsecSpdEntry
}

func (*IpsecSpdDetails) GetMessageName() string {
	return "ipsec_spd_details"
}
func (*IpsecSpdDetails) GetCrcString() string {
	return "928e5fcc"
}
func (*IpsecSpdDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpsecSpdDump represents VPP binary API message 'ipsec_spd_dump':
type IpsecSpdDump struct {
	SpdID uint32
	SaID  uint32
}

func (*IpsecSpdDump) GetMessageName() string {
	return "ipsec_spd_dump"
}
func (*IpsecSpdDump) GetCrcString() string {
	return "afefbf7d"
}
func (*IpsecSpdDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpsecSpdEntryAddDel represents VPP binary API message 'ipsec_spd_entry_add_del':
type IpsecSpdEntryAddDel struct {
	IsAdd uint8
	Entry IpsecSpdEntry
}

func (*IpsecSpdEntryAddDel) GetMessageName() string {
	return "ipsec_spd_entry_add_del"
}
func (*IpsecSpdEntryAddDel) GetCrcString() string {
	return "bbab53da"
}
func (*IpsecSpdEntryAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpsecSpdEntryAddDelReply represents VPP binary API message 'ipsec_spd_entry_add_del_reply':
type IpsecSpdEntryAddDelReply struct {
	Retval    int32
	StatIndex uint32
}

func (*IpsecSpdEntryAddDelReply) GetMessageName() string {
	return "ipsec_spd_entry_add_del_reply"
}
func (*IpsecSpdEntryAddDelReply) GetCrcString() string {
	return "9ffac24b"
}
func (*IpsecSpdEntryAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpsecSpdInterfaceDetails represents VPP binary API message 'ipsec_spd_interface_details':
type IpsecSpdInterfaceDetails struct {
	SpdIndex  uint32
	SwIfIndex uint32
}

func (*IpsecSpdInterfaceDetails) GetMessageName() string {
	return "ipsec_spd_interface_details"
}
func (*IpsecSpdInterfaceDetails) GetCrcString() string {
	return "2c54296d"
}
func (*IpsecSpdInterfaceDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpsecSpdInterfaceDump represents VPP binary API message 'ipsec_spd_interface_dump':
type IpsecSpdInterfaceDump struct {
	SpdIndex      uint32
	SpdIndexValid uint8
}

func (*IpsecSpdInterfaceDump) GetMessageName() string {
	return "ipsec_spd_interface_dump"
}
func (*IpsecSpdInterfaceDump) GetCrcString() string {
	return "8971de19"
}
func (*IpsecSpdInterfaceDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpsecSpdsDetails represents VPP binary API message 'ipsec_spds_details':
type IpsecSpdsDetails struct {
	SpdID     uint32
	Npolicies uint32
}

func (*IpsecSpdsDetails) GetMessageName() string {
	return "ipsec_spds_details"
}
func (*IpsecSpdsDetails) GetCrcString() string {
	return "a04bb254"
}
func (*IpsecSpdsDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpsecSpdsDump represents VPP binary API message 'ipsec_spds_dump':
type IpsecSpdsDump struct{}

func (*IpsecSpdsDump) GetMessageName() string {
	return "ipsec_spds_dump"
}
func (*IpsecSpdsDump) GetCrcString() string {
	return "51077d14"
}
func (*IpsecSpdsDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpsecTunnelIfAddDel represents VPP binary API message 'ipsec_tunnel_if_add_del':
type IpsecTunnelIfAddDel struct {
	IsAdd              uint8
	Esn                uint8
	AntiReplay         uint8
	LocalIP            Address
	RemoteIP           Address
	LocalSpi           uint32
	RemoteSpi          uint32
	CryptoAlg          uint8
	LocalCryptoKeyLen  uint8
	LocalCryptoKey     []byte `struc:"[128]byte"`
	RemoteCryptoKeyLen uint8
	RemoteCryptoKey    []byte `struc:"[128]byte"`
	IntegAlg           uint8
	LocalIntegKeyLen   uint8
	LocalIntegKey      []byte `struc:"[128]byte"`
	RemoteIntegKeyLen  uint8
	RemoteIntegKey     []byte `struc:"[128]byte"`
	Renumber           uint8
	ShowInstance       uint32
	UDPEncap           uint8
	TxTableID          uint32
	Salt               uint32
}

func (*IpsecTunnelIfAddDel) GetMessageName() string {
	return "ipsec_tunnel_if_add_del"
}
func (*IpsecTunnelIfAddDel) GetCrcString() string {
	return "94c63a9e"
}
func (*IpsecTunnelIfAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpsecTunnelIfAddDelReply represents VPP binary API message 'ipsec_tunnel_if_add_del_reply':
type IpsecTunnelIfAddDelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*IpsecTunnelIfAddDelReply) GetMessageName() string {
	return "ipsec_tunnel_if_add_del_reply"
}
func (*IpsecTunnelIfAddDelReply) GetCrcString() string {
	return "fda5941f"
}
func (*IpsecTunnelIfAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpsecTunnelIfSetKey represents VPP binary API message 'ipsec_tunnel_if_set_key':
type IpsecTunnelIfSetKey struct {
	SwIfIndex uint32
	KeyType   uint8
	Alg       uint8
	KeyLen    uint8
	Key       []byte `struc:"[128]byte"`
}

func (*IpsecTunnelIfSetKey) GetMessageName() string {
	return "ipsec_tunnel_if_set_key"
}
func (*IpsecTunnelIfSetKey) GetCrcString() string {
	return "326169a8"
}
func (*IpsecTunnelIfSetKey) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpsecTunnelIfSetKeyReply represents VPP binary API message 'ipsec_tunnel_if_set_key_reply':
type IpsecTunnelIfSetKeyReply struct {
	Retval int32
}

func (*IpsecTunnelIfSetKeyReply) GetMessageName() string {
	return "ipsec_tunnel_if_set_key_reply"
}
func (*IpsecTunnelIfSetKeyReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IpsecTunnelIfSetKeyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpsecTunnelIfSetSa represents VPP binary API message 'ipsec_tunnel_if_set_sa':
type IpsecTunnelIfSetSa struct {
	SwIfIndex  uint32
	SaID       uint32
	IsOutbound uint8
}

func (*IpsecTunnelIfSetSa) GetMessageName() string {
	return "ipsec_tunnel_if_set_sa"
}
func (*IpsecTunnelIfSetSa) GetCrcString() string {
	return "6ab567f2"
}
func (*IpsecTunnelIfSetSa) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpsecTunnelIfSetSaReply represents VPP binary API message 'ipsec_tunnel_if_set_sa_reply':
type IpsecTunnelIfSetSaReply struct {
	Retval int32
}

func (*IpsecTunnelIfSetSaReply) GetMessageName() string {
	return "ipsec_tunnel_if_set_sa_reply"
}
func (*IpsecTunnelIfSetSaReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IpsecTunnelIfSetSaReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*IpsecBackendDetails)(nil), "ipsec.IpsecBackendDetails")
	api.RegisterMessage((*IpsecBackendDump)(nil), "ipsec.IpsecBackendDump")
	api.RegisterMessage((*IpsecInterfaceAddDelSpd)(nil), "ipsec.IpsecInterfaceAddDelSpd")
	api.RegisterMessage((*IpsecInterfaceAddDelSpdReply)(nil), "ipsec.IpsecInterfaceAddDelSpdReply")
	api.RegisterMessage((*IpsecSaDetails)(nil), "ipsec.IpsecSaDetails")
	api.RegisterMessage((*IpsecSaDump)(nil), "ipsec.IpsecSaDump")
	api.RegisterMessage((*IpsecSaSetKey)(nil), "ipsec.IpsecSaSetKey")
	api.RegisterMessage((*IpsecSaSetKeyReply)(nil), "ipsec.IpsecSaSetKeyReply")
	api.RegisterMessage((*IpsecSadEntryAddDel)(nil), "ipsec.IpsecSadEntryAddDel")
	api.RegisterMessage((*IpsecSadEntryAddDelReply)(nil), "ipsec.IpsecSadEntryAddDelReply")
	api.RegisterMessage((*IpsecSelectBackend)(nil), "ipsec.IpsecSelectBackend")
	api.RegisterMessage((*IpsecSelectBackendReply)(nil), "ipsec.IpsecSelectBackendReply")
	api.RegisterMessage((*IpsecSpdAddDel)(nil), "ipsec.IpsecSpdAddDel")
	api.RegisterMessage((*IpsecSpdAddDelReply)(nil), "ipsec.IpsecSpdAddDelReply")
	api.RegisterMessage((*IpsecSpdDetails)(nil), "ipsec.IpsecSpdDetails")
	api.RegisterMessage((*IpsecSpdDump)(nil), "ipsec.IpsecSpdDump")
	api.RegisterMessage((*IpsecSpdEntryAddDel)(nil), "ipsec.IpsecSpdEntryAddDel")
	api.RegisterMessage((*IpsecSpdEntryAddDelReply)(nil), "ipsec.IpsecSpdEntryAddDelReply")
	api.RegisterMessage((*IpsecSpdInterfaceDetails)(nil), "ipsec.IpsecSpdInterfaceDetails")
	api.RegisterMessage((*IpsecSpdInterfaceDump)(nil), "ipsec.IpsecSpdInterfaceDump")
	api.RegisterMessage((*IpsecSpdsDetails)(nil), "ipsec.IpsecSpdsDetails")
	api.RegisterMessage((*IpsecSpdsDump)(nil), "ipsec.IpsecSpdsDump")
	api.RegisterMessage((*IpsecTunnelIfAddDel)(nil), "ipsec.IpsecTunnelIfAddDel")
	api.RegisterMessage((*IpsecTunnelIfAddDelReply)(nil), "ipsec.IpsecTunnelIfAddDelReply")
	api.RegisterMessage((*IpsecTunnelIfSetKey)(nil), "ipsec.IpsecTunnelIfSetKey")
	api.RegisterMessage((*IpsecTunnelIfSetKeyReply)(nil), "ipsec.IpsecTunnelIfSetKeyReply")
	api.RegisterMessage((*IpsecTunnelIfSetSa)(nil), "ipsec.IpsecTunnelIfSetSa")
	api.RegisterMessage((*IpsecTunnelIfSetSaReply)(nil), "ipsec.IpsecTunnelIfSetSaReply")
}

var Messages = []api.Message{
	(*IpsecBackendDetails)(nil),
	(*IpsecBackendDump)(nil),
	(*IpsecInterfaceAddDelSpd)(nil),
	(*IpsecInterfaceAddDelSpdReply)(nil),
	(*IpsecSaDetails)(nil),
	(*IpsecSaDump)(nil),
	(*IpsecSaSetKey)(nil),
	(*IpsecSaSetKeyReply)(nil),
	(*IpsecSadEntryAddDel)(nil),
	(*IpsecSadEntryAddDelReply)(nil),
	(*IpsecSelectBackend)(nil),
	(*IpsecSelectBackendReply)(nil),
	(*IpsecSpdAddDel)(nil),
	(*IpsecSpdAddDelReply)(nil),
	(*IpsecSpdDetails)(nil),
	(*IpsecSpdDump)(nil),
	(*IpsecSpdEntryAddDel)(nil),
	(*IpsecSpdEntryAddDelReply)(nil),
	(*IpsecSpdInterfaceDetails)(nil),
	(*IpsecSpdInterfaceDump)(nil),
	(*IpsecSpdsDetails)(nil),
	(*IpsecSpdsDump)(nil),
	(*IpsecTunnelIfAddDel)(nil),
	(*IpsecTunnelIfAddDelReply)(nil),
	(*IpsecTunnelIfSetKey)(nil),
	(*IpsecTunnelIfSetKeyReply)(nil),
	(*IpsecTunnelIfSetSa)(nil),
	(*IpsecTunnelIfSetSaReply)(nil),
}
