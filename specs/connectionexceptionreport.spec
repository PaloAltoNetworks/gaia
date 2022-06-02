# Model
model:
  rest_name: connectionexceptionreport
  resource_name: connectionexceptionreports
  entity_name: ConnectionExceptionReport
  package: zack
  group: policy/networking
  description: Post a new flow log.
  extends:
  - '@identifiable-stored'
  - '@zoned-monotonic'
  - '@migratable'
  validations:
  - $connectionexceptionreport

# Indexes
indexes:
- - processingunitnamespace
  - timestamp
- - enforcernamespace
  - timestamp
- - namespace
  - timestamp

# Attributes
attributes:
  v1:
  - name: destinationController
    description: |-
      Identifier of the destination controller. This should be set in
      SynAckTransmitted state.
    type: string
    exposed: true
    stored: true
    deprecated: true
    example_value: api.west.acme.com
    omit_empty: true
    extensions:
      bson_name: a

  - name: destinationIP
    description: Destination IP address.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: b

  - name: destinationPort
    description: Port of the destination.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: c

  - name: destinationProcessingUnitID
    description: |-
      ID of the destination processing unit. This should be set in SynAckTransmitted
      state.
    type: string
    exposed: true
    stored: true
    example_value: xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: d

  - name: enforcerID
    description: ID of the enforcer.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: e

  - name: enforcerNamespace
    description: Namespace of the enforcer.
    type: string
    exposed: true
    stored: true
    deprecated: true
    example_value: /my/namespace
    omit_empty: true
    extensions:
      bson_name: f

  - name: namespace
    description: Namespace of the processing unit that encountered this exception.
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
      bson_name: p

  - name: processingUnitID
    description: ID of the processing unit encountered this exception.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: g

  - name: processingUnitNamespace
    description: Namespace of the processing unit encountered this exception.
    type: string
    exposed: true
    stored: true
    deprecated: true
    example_value: /my/namespace
    omit_empty: true
    extensions:
      bson_name: h

  - name: protocol
    description: Protocol number.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 6
    omit_empty: true
    extensions:
      bson_name: i

  - name: reason
    description: It specifies the reason for the exception.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: j

  - name: serviceType
    description: Type of the service.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - L3
    - HTTP
    - TCP
    default_value: L3
    omit_empty: true
    extensions:
      bson_name: o

  - name: sourceIP
    description: Source IP address.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: k

  - name: state
    description: Represents the current state this report was generated.
    type: enum
    exposed: true
    subtype: string
    stored: true
    required: true
    allowed_choices:
    - SynTransmitted
    - SynAckTransmitted
    - AckTransmitted
    - SynReceived
    - SynAckReceived
    - AckProcessed
    - Data
    - Unknown
    - TFOProbe
    example_value:
    - Unknown
    extensions:
      bson_name: l

  - name: timestamp
    description: Time and date of the report.
    type: time
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: m

  - name: value
    description: Number of packets hit.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 1
    omit_empty: true
    extensions:
      bson_name: "n"
