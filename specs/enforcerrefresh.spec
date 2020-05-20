# Model
model:
  rest_name: enforcerrefresh
  resource_name: enforcerrefreshes
  entity_name: EnforcerRefresh
  package: gaga
  group: core/policy
  description: |-
    Sent to client when a poke has been triggered using the
    parameter `?notify=true`. This is used by instances of enforcerd to notify an
    external change on the processing unit must be processed.

# Attributes
attributes:
  v1:
  - name: ID
    description: Contains the ID of the target enforcer.
    type: string
    exposed: true
    read_only: true
    getter: true
    setter: true
    identifier: true

  - name: debug
    description: Set the debug information collected by the enforcer.
    type: enum
    exposed: true
    allowed_choices:
    - Counters
    - Logs
    - Packets
    - PUState
    - Pcap
    - CoreDump
    default_value: Counters
    omit_empty: true

  - name: namespace
    description: Contains the original namespace of the enforcer.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: debugProcessingUnitID
    description: Isolates debug information to a given processing unit, where possible.
    type: string
    exposed: true
    omit_empty: true

  - name: debugPcapFilter
    description: Packet capture filter, syntax varying by platform.
    type: string
    exposed: true
    omit_empty: true

  - name: debugID
    description: Can be used to correlate with a DebugBundle.
    type: string
    exposed: true
    omit_empty: true
    