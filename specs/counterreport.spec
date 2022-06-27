# Model
model:
  rest_name: counterreport
  resource_name: counterreports
  entity_name: CounterReport
  package: zack
  group: core/enforcer
  description: Post a new counter tracing report.
  extends:
  - '@identifiable-stored'
  - '@zoned-monotonic'
  - '@migratable'
  validations:
  - $counterreport

# Ordering
default_order:
- :no-inherit
- timestamp

# Indexes
indexes:
- - namespace
  - timestamp

# Attributes
attributes:
  v1:
  - name: AckInUnknownState
    description: Counter for sending FIN ACK received in unknown connection state.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: a

  - name: AckInvalidFormat
    description: Counter for ACK packet dropped because of invalid format.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: b

  - name: AckRejected
    description: Counter for ACK packets rejected as per policy.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: c

  - name: AckSigValidationFailed
    description: Counter for ACK packet dropped because signature validation failed.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: d

  - name: AckTCPNoTCPAuthOption
    description: Counter for TCP authentication option not found.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: e

  - name: ConnectionsProcessed
    description: Counter for connections processed.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: f

  - name: ContextIDNotFound
    description: Counter for unable to find ContextID.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: g

  - name: DroppedExternalService
    description: |-
      Counter for no ACLs found for external services. Dropping application SYN
      packet.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: h

  - name: InvalidConnState
    description: Counter for invalid connection state.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: i

  - name: InvalidNetState
    description: Counter for invalid net state.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: j

  - name: InvalidProtocol
    description: Counter for invalid protocol.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: k

  - name: InvalidSynAck
    description: Counter for processing unit is already dead - drop SYN ACK packet.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: l

  - name: MarkNotFound
    description: Counter for processing unit mark not found.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: m

  - name: NetSynNotSeen
    description: Counter for network SYN packet was not seen.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: "n"

  - name: NoConnFound
    description: Counter for no context or connection found.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: o

  - name: NonPUTraffic
    description: Counter for traffic that belongs to a non-processing unit process.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: p

  - name: OutOfOrderSynAck
    description: Counter for SYN ACK for flow with processed FIN ACK.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: q

  - name: PortNotFound
    description: Counter for port not found.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: r

  - name: RejectPacket
    description: Counter for reject the packet as per policy.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: s

  - name: ServicePostprocessorFailed
    description: Counter for post service processing failed for network packet.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: t

  - name: ServicePreprocessorFailed
    description: Counter for network packets that failed preprocessing.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: u

  - name: SynAckBadClaims
    description: Counter for SYN ACK packet dropped because of bad claims.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: v

  - name: SynAckClaimsMisMatch
    description: Counter for SYN ACK packet dropped because of encryption mismatch.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: w

  - name: SynAckDroppedExternalService
    description: Counter for SYN ACK from external service dropped.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: x

  - name: SynAckInvalidFormat
    description: Counter for SYN ACK packet dropped because of invalid format.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: "y"

  - name: SynAckMissingClaims
    description: Counter for SYN ACK packet dropped because of no claims.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: z

  - name: SynAckMissingToken
    description: Counter for SYN ACK packet dropped because of missing token.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: aa

  - name: SynAckNoTCPAuthOption
    description: Counter for TCP authentication option not found.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ab

  - name: SynAckRejected
    description: Counter for dropping because of reject rule on transmitter.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ac

  - name: SynDroppedInvalidFormat
    description: Counter for SYN packet dropped because of invalid format.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ad

  - name: SynDroppedInvalidToken
    description: Counter for SYN packet dropped because of invalid token.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: af

  - name: SynDroppedNoClaims
    description: Counter for SYN packet dropped because of no claims.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ag

  - name: SynDroppedTCPOption
    description: Counter for TCP authentication option not found.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ah

  - name: SynRejectPacket
    description: Counter for SYN packet dropped due to policy.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ai

  - name: SynUnexpectedPacket
    description: Counter for received SYN packet from unknown processing unit.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: aj

  - name: TCPAuthNotFound
    description: Counter for TCP authentication option not found.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ak

  - name: UDPAckInvalidSignature
    description: Counter for UDP ACK packet dropped due to an invalid signature.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: al

  - name: UDPConnectionsProcessed
    description: Counter for number of processed UDP connections.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: am

  - name: UDPDropContextNotFound
    description: Counter for dropped UDP data packets with no context.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: an

  - name: UDPDropFin
    description: Counter for dropped UDP FIN handshake packets.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ao

  - name: UDPDropInNfQueue
    description: Counter for dropped UDP in NfQueue.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ap

  - name: UDPDropNoConnection
    description: Counter for dropped UDP data packets with no connection.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: aq

  - name: UDPDropPacket
    description: Counter for dropped UDP data packets.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ar

  - name: UDPDropQueueFull
    description: Counter for dropped UDP queue full.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: as

  - name: UDPDropSynAck
    description: Counter for dropped UDP SYN ACK handshake packets.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: at

  - name: UDPInvalidNetState
    description: Counter for UDP packets received in invalid network state.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: au

  - name: UDPPostProcessingFailed
    description: Counter for UDP packets failing postprocessing.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: av

  - name: UDPPreProcessingFailed
    description: Counter for UDP packets failing preprocessing.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: aw

  - name: UDPRejected
    description: Counter for UDP packets dropped due to policy.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ax

  - name: UDPSynAckDropBadClaims
    description: Counter for UDP SYN ACK packets dropped due to bad claims.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ay

  - name: UDPSynAckMissingClaims
    description: Counter for UDP SYN ACK packets dropped due to missing claims.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: az

  - name: UDPSynAckPolicy
    description: Counter for UDP SYN ACK packets dropped due to bad claims.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ba

  - name: UDPSynDrop
    description: Counter for dropped UDP SYN transmits.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: bb

  - name: UDPSynDropPolicy
    description: Counter for dropped UDP SYN policy.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: bc

  - name: UDPSynInvalidToken
    description: Counter for dropped UDP FIN handshake packets.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: bd

  - name: UDPSynMissingClaims
    description: Counter for UDP SYN packet dropped due to missing claims.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: be

  - name: UnknownError
    description: Counter for unknown error.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: bf

  - name: connectionsAnalyzed
    description: |-
      Non-zero counter indicates analyzed connections for unencrypted, encrypted,
      and packets from endpoint applications with the TCP Fast Open option set. These
      are not dropped counter.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: bg

  - name: connectionsDropped
    description: |-
      Non-zero counter indicates dropped connections because of invalid state,
      non-processing unit traffic, or out of order packets.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: bh

  - name: connectionsExpired
    description: |-
      Non-zero counter indicates expired connections because of response not being
      received within a certain amount of time after the request is made.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: bi

  - name: droppedPackets
    description: |-
      Non-zero counter indicates dropped packets that did not hit any of our iptables
      rules and queue drops.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: bj

  - name: encryptionFailures
    description: Non-zero counter indicates encryption processing failures of data packets.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: bk

  - name: enforcerID
    description: Identifier of the enforcer sending the report.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxxx-xxx-xxxx
    omit_empty: true
    extensions:
      bson_name: bl

  - name: enforcerNamespace
    description: |-
      Namespace of the enforcer sending the report. This field is deprecated. Use the
      'namespace' field instead.
      field instead.
    type: string
    exposed: true
    stored: true
    deprecated: true
    example_value: /my/namespace
    omit_empty: true
    extensions:
      bson_name: bm

  - name: externalNetworkConnections
    description: |-
      Non-zero counter indicates connections going to and from external networks.
      These may be drops or allowed counters.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: bn

  - name: namespace
    description: Namespace of the enforcer sending the report.
    type: string
    exposed: true
    stored: true
    read_only: true
    example_value: /my/namespace
    filterable: true
    getter: true
    setter: true
    omit_empty: true
    extensions:
      bson_name: bt

  - name: policyDrops
    description: Non-zero counter indicates packets dropped due to a reject policy.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: bo

  - name: processingUnitID
    description: PUID is the ID of the processing unit reporting the counter.
    type: string
    exposed: true
    stored: true
    example_value: xxx-xxx-xxx
    filterable: true
    omit_empty: true
    extensions:
      bson_name: bp

  - name: processingUnitNamespace
    description: Namespace of the processing unit reporting the counter.
    type: string
    exposed: true
    stored: true
    example_value: /my/namespace
    filterable: true
    omit_empty: true
    extensions:
      bson_name: bq

  - name: timestamp
    description: Timestamp is the date of the report.
    type: time
    exposed: true
    stored: true
    example_value: "2018-06-14T23:10:46.420397985Z"
    orderable: true
    omit_empty: true
    extensions:
      bson_name: br

  - name: tokenDrops
    description: |-
      Non-zero counter indicates packets rejected due to anything related to token
      creation/parsing failures.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: bs
