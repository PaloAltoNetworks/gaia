package protocols

import (
	"strconv"
)

// ANY defines all protocols
// NOTE: Keeping it uppercase for consistency
const ANY = "ANY"

// All L4 Protocols
// (Ref https://en.wikipedia.org/wiki/List_of_IP_protocol_numbers)
const (
	L4ProtocolHOPOPT            = "HOPOPT"
	L4ProtocolICMP              = "ICMP"
	L4ProtocolIGMP              = "IGMP"
	L4ProtocolGGP               = "GGP"
	L4ProtocolIPIP              = "IPIP"
	L4ProtocolST                = "ST"
	L4ProtocolTCP               = "TCP"
	L4ProtocolCBT               = "CBT"
	L4ProtocolEGP               = "EGP"
	L4ProtocolIGP               = "IGP"
	L4ProtocolBBNRCCMON         = "BBN-RCC-MON"
	L4ProtocolNVPII             = "NVP-II"
	L4ProtocolPUP               = "PUP"
	L4ProtocolARGUS             = "ARGUS"
	L4ProtocolEMCON             = "EMCON"
	L4ProtocolXNET              = "XNET"
	L4ProtocolCHAOS             = "CHAOS"
	L4ProtocolUDP               = "UDP"
	L4ProtocolMUX               = "MUX"
	L4ProtocolDCNMEAS           = "DCN-MEAS"
	L4ProtocolHMP               = "HMP"
	L4ProtocolPRM               = "PRM"
	L4ProtocolXNSIDP            = "XNS-IDP"
	L4ProtocolTRUNK1            = "TRUNK-1"
	L4ProtocolTRUNK2            = "TRUNK-2"
	L4ProtocolLEAF1             = "LEAF-1"
	L4ProtocolLEAF2             = "LEAF-2"
	L4ProtocolRDP               = "RDP"
	L4ProtocolIRTP              = "IRTP"
	L4ProtocolISOTP4            = "ISO-TP4"
	L4ProtocolNETBLT            = "NETBLT"
	L4ProtocolMFENSP            = "MFE-NSP"
	L4ProtocolMERITINP          = "MERIT-INP"
	L4ProtocolDCCP              = "DCCP"
	L4Protocol3PC               = "3PC"
	L4ProtocolIDPR              = "IDPR"
	L4ProtocolXTP               = "XTP"
	L4ProtocolDDP               = "DDP"
	L4ProtocolIDPRCMTP          = "IDPR-CMTP"
	L4ProtocolTP                = "TP++"
	L4ProtocolIL                = "IL"
	L4ProtocolIPv6              = "IPv6"
	L4ProtocolSDRP              = "SDRP"
	L4ProtocolIPv6Route         = "IPv6-Route"
	L4ProtocolIPv6Frag          = "IPv6-Frag"
	L4ProtocolIDRP              = "IDRP"
	L4ProtocolRSVP              = "RSVP"
	L4ProtocolGRE               = "GRE"
	L4ProtocolDSR               = "DSR"
	L4ProtocolBNA               = "BNA"
	L4ProtocolESP               = "ESP"
	L4ProtocolAH                = "AH"
	L4ProtocolINLSP             = "I-NLSP"
	L4ProtocolSWIPE             = "SWIPE"
	L4ProtocolNARP              = "NARP"
	L4ProtocolMOBILE            = "MOBILE"
	L4ProtocolTLSP              = "TLSP"
	L4ProtocolSKIP              = "SKIP"
	L4ProtocolIPv6ICMP          = "IPV6-ICMP"
	L4ProtocolIPv6NoNxt         = "IPV6-NoNxt"
	L4ProtocolIPv6Opts          = "IPV6-Opts"
	L4ProtocolAnyLocalHost      = "ANYLOCALHOST"
	L4ProtocolAnyLocalNetwork   = "ANYLOCALNETWORK"
	L4ProtocolCFTP              = "CFTP"
	L4ProtocolSATEXPAK          = "SAT-EXPAK"
	L4ProtocolKRYPTOLAN         = "KRYPTOLAN"
	L4ProtocolRVD               = "RVD"
	L4ProtocolIPPC              = "IPPC"
	L4ProtocolSATMON            = "SAT-MON"
	L4ProtocolVISA              = "VISA"
	L4ProtocolIPCU              = "IPCU"
	L4ProtocolCPNX              = "CPNX"
	L4ProtocolCPHB              = "CPHB"
	L4ProtocolWSN               = "WSN"
	L4ProtocolPVP               = "PVP"
	L4ProtocolBRSATMON          = "BR-SAT-MON"
	L4ProtocolSUNND             = "SUN-ND"
	L4ProtocolWBMON             = "WB-MON"
	L4ProtocolWBEXPAK           = "WB-EXPAK"
	L4ProtocolISOIP             = "ISO-IP"
	L4ProtocolVMTP              = "VMTP"
	L4ProtocolSECUREVMTP        = "SECURE-VMTP"
	L4ProtocolVINES             = "VINES"
	L4ProtocolTTP               = "TTP"
	L4ProtocolNSFNETIGP         = "NSFNET-IGP"
	L4ProtocolDGP               = "DGP"
	L4ProtocolTCF               = "TCF"
	L4ProtocolEIGRP             = "EIGRP"
	L4ProtocolOSPF              = "OSPF"
	L4ProtocolSpriteRPC         = "Sprite-RPC"
	L4ProtocolLARP              = "LARP"
	L4ProtocolMTP               = "MTP"
	L4ProtocolAX25              = "AX.25"
	L4ProtocolOS                = "OS"
	L4ProtocolMICP              = "MICP"
	L4ProtocolSCCSP             = "SCC-SP"
	L4ProtocolETHERIP           = "ETHERIP"
	L4ProtocolENCAP             = "ENCAP"
	L4ProtocolPrivateEncryption = "PrivateEncryption"
	L4ProtocolGMTP              = "GMTP"
	L4ProtocolIFMP              = "IFMP"
	L4ProtocolPNNI              = "PNNI"
	L4ProtocolPIM               = "PIM"
	L4ProtocolARIS              = "ARIS"
	L4ProtocolSCPS              = "SCPS"
	L4ProtocolQNX               = "QNX"
	L4ProtocolAN                = "A/N"
	L4ProtocolIPComp            = "IPCOMP"
	L4ProtocolSNP               = "SNP"
	L4ProtocolCompaqPeer        = "COMPAQ-PEER"
	L4ProtocolIPXinIP           = "IPX-IN-IP"
	L4ProtocolVRRP              = "VRRP"
	L4ProtocolPGM               = "PGM"
	L4ProtocolL2TP              = "L2TP"
	L4ProtocolDDX               = "DDX"
	L4ProtocolIATP              = "IATP"
	L4ProtocolSTP               = "STP"
	L4ProtocolSRP               = "SRP"
	L4ProtocolUTI               = "UTI"
	L4ProtocolSMP               = "SMP"
	L4ProtocolSM                = "SM"
	L4ProtocolPTP               = "PTP"
	L4ProtocolISIS              = "ISIS"
	L4ProtocolFIRE              = "FIRE"
	L4ProtocolCRTP              = "CRTP"
	L4ProtocolCRUDP             = "CRUDP"
	L4ProtocolSSCOPMCE          = "SSCOPMCE"
	L4ProtocolIPLT              = "IPLT"
	L4ProtocolSPS               = "SPS"
	L4ProtocolPIPE              = "PIPE"
	L4ProtocolSCTP              = "SCTP"
	L4ProtocolFC                = "FC"
	L4ProtocolRSVPE2EIGNORE     = "RSVP-E2E-IGNORE"
	L4ProtocolMobilityHeader    = "MOBILITY-HEADER"
	L4ProtocolUDPLite           = "UDPLITE"
	L4ProtocolMPLSinIP          = "MPLS-IN-IP"
	L4Protocolmanet             = "MANET"
	L4ProtocolHIP               = "HIP"
	L4ProtocolShim6             = "SHIM6"
	L4ProtocolWESP              = "WESP"
	L4ProtocolROHC              = "ROHC"
	L4ProtocolICMP6             = "ICMP6"
)

// L4ProtocolNumbers contains the list of all IANA protocols organized by numbers
var L4ProtocolNumbers = make([]string, 150)

// L4ProtocolNumbersMap contains the list of all IANA protocols organized by numbers
var L4ProtocolNumbersMap = map[int]string{}

// L4ProtocolNames contains the list of all IANA protocols organized by names
var L4ProtocolNames = map[string]int{}

// init initializes variables
func init() {
	L4ProtocolNumbers[0] = L4ProtocolHOPOPT
	L4ProtocolNumbers[1] = L4ProtocolICMP
	L4ProtocolNumbers[2] = L4ProtocolIGMP
	L4ProtocolNumbers[3] = L4ProtocolGGP
	L4ProtocolNumbers[4] = L4ProtocolIPIP
	L4ProtocolNumbers[5] = L4ProtocolST
	L4ProtocolNumbers[6] = L4ProtocolTCP
	L4ProtocolNumbers[7] = L4ProtocolCBT
	L4ProtocolNumbers[8] = L4ProtocolEGP
	L4ProtocolNumbers[9] = L4ProtocolIGP
	L4ProtocolNumbers[10] = L4ProtocolBBNRCCMON
	L4ProtocolNumbers[11] = L4ProtocolNVPII
	L4ProtocolNumbers[12] = L4ProtocolPUP
	L4ProtocolNumbers[13] = L4ProtocolARGUS
	L4ProtocolNumbers[14] = L4ProtocolEMCON
	L4ProtocolNumbers[15] = L4ProtocolXNET
	L4ProtocolNumbers[16] = L4ProtocolCHAOS
	L4ProtocolNumbers[17] = L4ProtocolUDP
	L4ProtocolNumbers[18] = L4ProtocolMUX
	L4ProtocolNumbers[19] = L4ProtocolDCNMEAS
	L4ProtocolNumbers[20] = L4ProtocolHMP
	L4ProtocolNumbers[21] = L4ProtocolPRM
	L4ProtocolNumbers[22] = L4ProtocolXNSIDP
	L4ProtocolNumbers[23] = L4ProtocolTRUNK1
	L4ProtocolNumbers[24] = L4ProtocolTRUNK2
	L4ProtocolNumbers[25] = L4ProtocolLEAF1
	L4ProtocolNumbers[26] = L4ProtocolLEAF2
	L4ProtocolNumbers[27] = L4ProtocolRDP
	L4ProtocolNumbers[28] = L4ProtocolIRTP
	L4ProtocolNumbers[29] = L4ProtocolISOTP4
	L4ProtocolNumbers[30] = L4ProtocolNETBLT
	L4ProtocolNumbers[31] = L4ProtocolMFENSP
	L4ProtocolNumbers[32] = L4ProtocolMERITINP
	L4ProtocolNumbers[33] = L4ProtocolDCCP
	L4ProtocolNumbers[34] = L4Protocol3PC
	L4ProtocolNumbers[35] = L4ProtocolIDPR
	L4ProtocolNumbers[36] = L4ProtocolXTP
	L4ProtocolNumbers[37] = L4ProtocolDDP
	L4ProtocolNumbers[38] = L4ProtocolIDPRCMTP
	L4ProtocolNumbers[39] = L4ProtocolTP
	L4ProtocolNumbers[40] = L4ProtocolIL
	L4ProtocolNumbers[41] = L4ProtocolIPv6
	L4ProtocolNumbers[42] = L4ProtocolSDRP
	L4ProtocolNumbers[43] = L4ProtocolIPv6Route
	L4ProtocolNumbers[44] = L4ProtocolIPv6Frag
	L4ProtocolNumbers[45] = L4ProtocolIDRP
	L4ProtocolNumbers[46] = L4ProtocolRSVP
	L4ProtocolNumbers[47] = L4ProtocolGRE
	L4ProtocolNumbers[48] = L4ProtocolDSR
	L4ProtocolNumbers[49] = L4ProtocolBNA
	L4ProtocolNumbers[50] = L4ProtocolESP
	L4ProtocolNumbers[51] = L4ProtocolAH
	L4ProtocolNumbers[52] = L4ProtocolINLSP
	L4ProtocolNumbers[53] = L4ProtocolSWIPE
	L4ProtocolNumbers[54] = L4ProtocolNARP
	L4ProtocolNumbers[55] = L4ProtocolMOBILE
	L4ProtocolNumbers[56] = L4ProtocolTLSP
	L4ProtocolNumbers[57] = L4ProtocolSKIP
	L4ProtocolNumbers[58] = L4ProtocolIPv6ICMP
	L4ProtocolNumbers[59] = L4ProtocolIPv6NoNxt
	L4ProtocolNumbers[60] = L4ProtocolIPv6Opts
	L4ProtocolNumbers[61] = L4ProtocolAnyLocalHost
	L4ProtocolNumbers[62] = L4ProtocolCFTP
	L4ProtocolNumbers[63] = L4ProtocolAnyLocalNetwork
	L4ProtocolNumbers[64] = L4ProtocolSATEXPAK
	L4ProtocolNumbers[65] = L4ProtocolKRYPTOLAN
	L4ProtocolNumbers[66] = L4ProtocolRVD
	L4ProtocolNumbers[67] = L4ProtocolIPPC
	L4ProtocolNumbers[69] = L4ProtocolSATMON
	L4ProtocolNumbers[70] = L4ProtocolVISA
	L4ProtocolNumbers[71] = L4ProtocolIPCU
	L4ProtocolNumbers[72] = L4ProtocolCPNX
	L4ProtocolNumbers[73] = L4ProtocolCPHB
	L4ProtocolNumbers[74] = L4ProtocolWSN
	L4ProtocolNumbers[75] = L4ProtocolPVP
	L4ProtocolNumbers[76] = L4ProtocolBRSATMON
	L4ProtocolNumbers[77] = L4ProtocolSUNND
	L4ProtocolNumbers[78] = L4ProtocolWBMON
	L4ProtocolNumbers[79] = L4ProtocolWBEXPAK
	L4ProtocolNumbers[80] = L4ProtocolISOIP
	L4ProtocolNumbers[81] = L4ProtocolVMTP
	L4ProtocolNumbers[82] = L4ProtocolSECUREVMTP
	L4ProtocolNumbers[83] = L4ProtocolVINES
	L4ProtocolNumbers[84] = L4ProtocolTTP
	L4ProtocolNumbers[85] = L4ProtocolNSFNETIGP
	L4ProtocolNumbers[86] = L4ProtocolDGP
	L4ProtocolNumbers[87] = L4ProtocolTCF
	L4ProtocolNumbers[88] = L4ProtocolEIGRP
	L4ProtocolNumbers[89] = L4ProtocolOSPF
	L4ProtocolNumbers[90] = L4ProtocolSpriteRPC
	L4ProtocolNumbers[91] = L4ProtocolLARP
	L4ProtocolNumbers[92] = L4ProtocolMTP
	L4ProtocolNumbers[93] = L4ProtocolAX25
	L4ProtocolNumbers[94] = L4ProtocolOS
	L4ProtocolNumbers[95] = L4ProtocolMICP
	L4ProtocolNumbers[96] = L4ProtocolSCCSP
	L4ProtocolNumbers[97] = L4ProtocolETHERIP
	L4ProtocolNumbers[98] = L4ProtocolENCAP
	L4ProtocolNumbers[99] = L4ProtocolPrivateEncryption
	L4ProtocolNumbers[100] = L4ProtocolGMTP
	L4ProtocolNumbers[101] = L4ProtocolIFMP
	L4ProtocolNumbers[102] = L4ProtocolPNNI
	L4ProtocolNumbers[103] = L4ProtocolPIM
	L4ProtocolNumbers[104] = L4ProtocolARIS
	L4ProtocolNumbers[105] = L4ProtocolSCPS
	L4ProtocolNumbers[106] = L4ProtocolQNX
	L4ProtocolNumbers[107] = L4ProtocolAN
	L4ProtocolNumbers[108] = L4ProtocolIPComp
	L4ProtocolNumbers[109] = L4ProtocolSNP
	L4ProtocolNumbers[110] = L4ProtocolCompaqPeer
	L4ProtocolNumbers[111] = L4ProtocolIPXinIP
	L4ProtocolNumbers[112] = L4ProtocolVRRP
	L4ProtocolNumbers[113] = L4ProtocolPGM
	L4ProtocolNumbers[115] = L4ProtocolL2TP
	L4ProtocolNumbers[116] = L4ProtocolDDX
	L4ProtocolNumbers[117] = L4ProtocolIATP
	L4ProtocolNumbers[118] = L4ProtocolSTP
	L4ProtocolNumbers[119] = L4ProtocolSRP
	L4ProtocolNumbers[120] = L4ProtocolUTI
	L4ProtocolNumbers[121] = L4ProtocolSMP
	L4ProtocolNumbers[122] = L4ProtocolSM
	L4ProtocolNumbers[123] = L4ProtocolPTP
	L4ProtocolNumbers[124] = L4ProtocolISIS
	L4ProtocolNumbers[125] = L4ProtocolFIRE
	L4ProtocolNumbers[126] = L4ProtocolCRTP
	L4ProtocolNumbers[127] = L4ProtocolCRUDP
	L4ProtocolNumbers[128] = L4ProtocolSSCOPMCE
	L4ProtocolNumbers[129] = L4ProtocolIPLT
	L4ProtocolNumbers[130] = L4ProtocolSPS
	L4ProtocolNumbers[131] = L4ProtocolPIPE
	L4ProtocolNumbers[132] = L4ProtocolSCTP
	L4ProtocolNumbers[133] = L4ProtocolFC
	L4ProtocolNumbers[134] = L4ProtocolRSVPE2EIGNORE
	L4ProtocolNumbers[135] = L4ProtocolMobilityHeader
	L4ProtocolNumbers[136] = L4ProtocolUDPLite
	L4ProtocolNumbers[137] = L4ProtocolMPLSinIP
	L4ProtocolNumbers[138] = L4Protocolmanet
	L4ProtocolNumbers[139] = L4ProtocolHIP
	L4ProtocolNumbers[140] = L4ProtocolShim6
	L4ProtocolNumbers[141] = L4ProtocolWESP
	L4ProtocolNumbers[142] = L4ProtocolROHC
	L4ProtocolNumbers[143] = L4ProtocolICMP6

	L4ProtocolNumbersMap = map[int]string{
		0:   L4ProtocolHOPOPT,
		1:   L4ProtocolICMP,
		2:   L4ProtocolIGMP,
		3:   L4ProtocolGGP,
		4:   L4ProtocolIPIP,
		5:   L4ProtocolST,
		6:   L4ProtocolTCP,
		7:   L4ProtocolCBT,
		8:   L4ProtocolEGP,
		9:   L4ProtocolIGP,
		10:  L4ProtocolBBNRCCMON,
		11:  L4ProtocolNVPII,
		12:  L4ProtocolPUP,
		13:  L4ProtocolARGUS,
		14:  L4ProtocolEMCON,
		15:  L4ProtocolXNET,
		16:  L4ProtocolCHAOS,
		17:  L4ProtocolUDP,
		18:  L4ProtocolMUX,
		19:  L4ProtocolDCNMEAS,
		20:  L4ProtocolHMP,
		21:  L4ProtocolPRM,
		22:  L4ProtocolXNSIDP,
		23:  L4ProtocolTRUNK1,
		24:  L4ProtocolTRUNK2,
		25:  L4ProtocolLEAF1,
		26:  L4ProtocolLEAF2,
		27:  L4ProtocolRDP,
		28:  L4ProtocolIRTP,
		29:  L4ProtocolISOTP4,
		30:  L4ProtocolNETBLT,
		31:  L4ProtocolMFENSP,
		32:  L4ProtocolMERITINP,
		33:  L4ProtocolDCCP,
		34:  L4Protocol3PC,
		35:  L4ProtocolIDPR,
		36:  L4ProtocolXTP,
		37:  L4ProtocolDDP,
		38:  L4ProtocolIDPRCMTP,
		39:  L4ProtocolTP,
		40:  L4ProtocolIL,
		41:  L4ProtocolIPv6,
		42:  L4ProtocolSDRP,
		43:  L4ProtocolIPv6Route,
		44:  L4ProtocolIPv6Frag,
		45:  L4ProtocolIDRP,
		46:  L4ProtocolRSVP,
		47:  L4ProtocolGRE,
		48:  L4ProtocolDSR,
		49:  L4ProtocolBNA,
		50:  L4ProtocolESP,
		51:  L4ProtocolAH,
		52:  L4ProtocolINLSP,
		53:  L4ProtocolSWIPE,
		54:  L4ProtocolNARP,
		55:  L4ProtocolMOBILE,
		56:  L4ProtocolTLSP,
		57:  L4ProtocolSKIP,
		58:  L4ProtocolIPv6ICMP,
		59:  L4ProtocolIPv6NoNxt,
		60:  L4ProtocolIPv6Opts,
		61:  L4ProtocolAnyLocalHost,
		62:  L4ProtocolCFTP,
		63:  L4ProtocolAnyLocalNetwork,
		64:  L4ProtocolSATEXPAK,
		65:  L4ProtocolKRYPTOLAN,
		66:  L4ProtocolRVD,
		67:  L4ProtocolIPPC,
		69:  L4ProtocolSATMON,
		70:  L4ProtocolVISA,
		71:  L4ProtocolIPCU,
		72:  L4ProtocolCPNX,
		73:  L4ProtocolCPHB,
		74:  L4ProtocolWSN,
		75:  L4ProtocolPVP,
		76:  L4ProtocolBRSATMON,
		77:  L4ProtocolSUNND,
		78:  L4ProtocolWBMON,
		79:  L4ProtocolWBEXPAK,
		80:  L4ProtocolISOIP,
		81:  L4ProtocolVMTP,
		82:  L4ProtocolSECUREVMTP,
		83:  L4ProtocolVINES,
		84:  L4ProtocolTTP,
		85:  L4ProtocolNSFNETIGP,
		86:  L4ProtocolDGP,
		87:  L4ProtocolTCF,
		88:  L4ProtocolEIGRP,
		89:  L4ProtocolOSPF,
		90:  L4ProtocolSpriteRPC,
		91:  L4ProtocolLARP,
		92:  L4ProtocolMTP,
		93:  L4ProtocolAX25,
		94:  L4ProtocolOS,
		95:  L4ProtocolMICP,
		96:  L4ProtocolSCCSP,
		97:  L4ProtocolETHERIP,
		98:  L4ProtocolENCAP,
		99:  L4ProtocolPrivateEncryption,
		100: L4ProtocolGMTP,
		101: L4ProtocolIFMP,
		102: L4ProtocolPNNI,
		103: L4ProtocolPIM,
		104: L4ProtocolARIS,
		105: L4ProtocolSCPS,
		106: L4ProtocolQNX,
		107: L4ProtocolAN,
		108: L4ProtocolIPComp,
		109: L4ProtocolSNP,
		110: L4ProtocolCompaqPeer,
		111: L4ProtocolIPXinIP,
		112: L4ProtocolVRRP,
		113: L4ProtocolPGM,
		115: L4ProtocolL2TP,
		116: L4ProtocolDDX,
		117: L4ProtocolIATP,
		118: L4ProtocolSTP,
		119: L4ProtocolSRP,
		120: L4ProtocolUTI,
		121: L4ProtocolSMP,
		122: L4ProtocolSM,
		123: L4ProtocolPTP,
		124: L4ProtocolISIS,
		125: L4ProtocolFIRE,
		126: L4ProtocolCRTP,
		127: L4ProtocolCRUDP,
		128: L4ProtocolSSCOPMCE,
		129: L4ProtocolIPLT,
		130: L4ProtocolSPS,
		131: L4ProtocolPIPE,
		132: L4ProtocolSCTP,
		133: L4ProtocolFC,
		134: L4ProtocolRSVPE2EIGNORE,
		135: L4ProtocolMobilityHeader,
		136: L4ProtocolUDPLite,
		137: L4ProtocolMPLSinIP,
		138: L4Protocolmanet,
		139: L4ProtocolHIP,
		140: L4ProtocolShim6,
		141: L4ProtocolWESP,
		142: L4ProtocolROHC,
		143: L4ProtocolICMP6,
	}

	L4ProtocolNames = map[string]int{
		L4ProtocolHOPOPT:            0,
		L4ProtocolICMP:              1,
		L4ProtocolIGMP:              2,
		L4ProtocolGGP:               3,
		L4ProtocolIPIP:              4,
		L4ProtocolST:                5,
		L4ProtocolTCP:               6,
		L4ProtocolCBT:               7,
		L4ProtocolEGP:               8,
		L4ProtocolIGP:               9,
		L4ProtocolBBNRCCMON:         10,
		L4ProtocolNVPII:             11,
		L4ProtocolPUP:               12,
		L4ProtocolARGUS:             13,
		L4ProtocolEMCON:             14,
		L4ProtocolXNET:              15,
		L4ProtocolCHAOS:             16,
		L4ProtocolUDP:               17,
		L4ProtocolMUX:               18,
		L4ProtocolDCNMEAS:           19,
		L4ProtocolHMP:               20,
		L4ProtocolPRM:               21,
		L4ProtocolXNSIDP:            22,
		L4ProtocolTRUNK1:            23,
		L4ProtocolTRUNK2:            24,
		L4ProtocolLEAF1:             25,
		L4ProtocolLEAF2:             26,
		L4ProtocolRDP:               27,
		L4ProtocolIRTP:              28,
		L4ProtocolISOTP4:            29,
		L4ProtocolNETBLT:            30,
		L4ProtocolMFENSP:            31,
		L4ProtocolMERITINP:          32,
		L4ProtocolDCCP:              33,
		L4Protocol3PC:               34,
		L4ProtocolIDPR:              35,
		L4ProtocolXTP:               36,
		L4ProtocolDDP:               37,
		L4ProtocolIDPRCMTP:          38,
		L4ProtocolTP:                39,
		L4ProtocolIL:                40,
		L4ProtocolIPv6:              41,
		L4ProtocolSDRP:              42,
		L4ProtocolIPv6Route:         43,
		L4ProtocolIPv6Frag:          44,
		L4ProtocolIDRP:              45,
		L4ProtocolRSVP:              46,
		L4ProtocolGRE:               47,
		L4ProtocolDSR:               48,
		L4ProtocolBNA:               49,
		L4ProtocolESP:               50,
		L4ProtocolAH:                51,
		L4ProtocolINLSP:             52,
		L4ProtocolSWIPE:             53,
		L4ProtocolNARP:              54,
		L4ProtocolMOBILE:            55,
		L4ProtocolTLSP:              56,
		L4ProtocolSKIP:              57,
		L4ProtocolIPv6ICMP:          58,
		L4ProtocolIPv6NoNxt:         59,
		L4ProtocolIPv6Opts:          60,
		L4ProtocolAnyLocalHost:      61,
		L4ProtocolCFTP:              62,
		L4ProtocolAnyLocalNetwork:   63,
		L4ProtocolSATEXPAK:          64,
		L4ProtocolKRYPTOLAN:         65,
		L4ProtocolRVD:               66,
		L4ProtocolIPPC:              67,
		L4ProtocolSATMON:            69,
		L4ProtocolVISA:              70,
		L4ProtocolIPCU:              71,
		L4ProtocolCPNX:              72,
		L4ProtocolCPHB:              73,
		L4ProtocolWSN:               74,
		L4ProtocolPVP:               75,
		L4ProtocolBRSATMON:          76,
		L4ProtocolSUNND:             77,
		L4ProtocolWBMON:             78,
		L4ProtocolWBEXPAK:           79,
		L4ProtocolISOIP:             80,
		L4ProtocolVMTP:              81,
		L4ProtocolSECUREVMTP:        82,
		L4ProtocolVINES:             83,
		L4ProtocolTTP:               84,
		L4ProtocolNSFNETIGP:         85,
		L4ProtocolDGP:               86,
		L4ProtocolTCF:               87,
		L4ProtocolEIGRP:             88,
		L4ProtocolOSPF:              89,
		L4ProtocolSpriteRPC:         90,
		L4ProtocolLARP:              91,
		L4ProtocolMTP:               92,
		L4ProtocolAX25:              93,
		L4ProtocolOS:                94,
		L4ProtocolMICP:              95,
		L4ProtocolSCCSP:             96,
		L4ProtocolETHERIP:           97,
		L4ProtocolENCAP:             98,
		L4ProtocolPrivateEncryption: 99,
		L4ProtocolGMTP:              100,
		L4ProtocolIFMP:              101,
		L4ProtocolPNNI:              102,
		L4ProtocolPIM:               103,
		L4ProtocolARIS:              104,
		L4ProtocolSCPS:              105,
		L4ProtocolQNX:               106,
		L4ProtocolAN:                107,
		L4ProtocolIPComp:            108,
		L4ProtocolSNP:               109,
		L4ProtocolCompaqPeer:        110,
		L4ProtocolIPXinIP:           111,
		L4ProtocolVRRP:              112,
		L4ProtocolPGM:               113,
		L4ProtocolL2TP:              115,
		L4ProtocolDDX:               116,
		L4ProtocolIATP:              117,
		L4ProtocolSTP:               118,
		L4ProtocolSRP:               119,
		L4ProtocolUTI:               120,
		L4ProtocolSMP:               121,
		L4ProtocolSM:                122,
		L4ProtocolPTP:               123,
		L4ProtocolISIS:              124,
		L4ProtocolFIRE:              125,
		L4ProtocolCRTP:              126,
		L4ProtocolCRUDP:             127,
		L4ProtocolSSCOPMCE:          128,
		L4ProtocolIPLT:              129,
		L4ProtocolSPS:               130,
		L4ProtocolPIPE:              131,
		L4ProtocolSCTP:              132,
		L4ProtocolFC:                133,
		L4ProtocolRSVPE2EIGNORE:     134,
		L4ProtocolMobilityHeader:    135,
		L4ProtocolUDPLite:           136,
		L4ProtocolMPLSinIP:          137,
		L4Protocolmanet:             138,
		L4ProtocolHIP:               139,
		L4ProtocolShim6:             140,
		L4ProtocolWESP:              141,
		L4ProtocolROHC:              142,
		L4ProtocolICMP6:             143,
	}
}

// L4ProtocolNameFromNumber returns the IANA name for the given protocol number.
func L4ProtocolNameFromNumber(n int64) string {

	if n < 0 {
		return ""
	}

	if n >= int64(len(L4ProtocolNumbers)) {
		return strconv.FormatInt(n, 10)
	}

	if protocol := L4ProtocolNumbers[n]; protocol != "" {
		return protocol
	}

	return strconv.FormatInt(n, 10)
}

// L4ProtocolNumberFromName returns the protocol number given a protocol name.
func L4ProtocolNumberFromName(name string) int {
	if n, ok := L4ProtocolNames[name]; ok {
		return n
	}

	return -1
}
