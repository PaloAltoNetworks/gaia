# Model
model:
  rest_name: processingunitrefresh
  resource_name: processingunitrefreshes
  entity_name: ProcessingUnitRefresh
  package: gaga
  group: core/policy
  description: |-
    Sent to client when a poke has been triggered using the
    parameter `?notify=true`. This is used to notify a enforcer of an
    external change on the processing unit must be processed.

# Attributes
attributes:
  v1:
  - name: ID
    description: Contains the ID of the target processing unit.
    type: string
    exposed: true
    read_only: true
    getter: true
    setter: true
    identifier: true

  - name: debug
    description: |-
      If set to true, start reporting debug information for the target processing
      unit.
    type: boolean
    exposed: true
    omit_empty: true

  - name: namespace
    description: Contains the original namespace of the processing unit.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: pingAddress
    description: Destination address to run ping.
    type: string
    exposed: true
    omit_empty: true

  - name: pingEnabled
    description: If set to true, start ping to the destination.
    type: boolean
    exposed: true
    omit_empty: true

  - name: pingIterations
    description: Number of iterations to run a ping probe.
    type: integer
    exposed: true
    default_value: 1
    min_value: 1
    omit_empty: true

  - name: pingMode
    description: Represents the mode of ping to be used.
    type: enum
    exposed: true
    allowed_choices:
    - Auto
    - L3
    - L4
    - L7
    default_value: Auto

  - name: pingPort
    description: Destination port to run ping.
    type: integer
    exposed: true
    omit_empty: true

  - name: refreshID
    description: ID unique per purefresh event.
    type: string
    exposed: true
    read_only: true

  - name: refreshPolicy
    description: If set to true, the target processing unit will refresh its policy
      immediately.
    type: boolean
    exposed: true
    omit_empty: true

  - name: traceApplicationConnections
    description: |-
      Instructs the enforcer to send records for all
      application-initiated connections for the target processing unit.
    type: boolean
    exposed: true
    omit_empty: true

  - name: traceDuration
    description: |-
      Determines the length of the time interval that the trace must be
      enabled, using [Golang duration
      syntax](https://golang.org/pkg/time/#example_Duration).
    type: string
    exposed: true
    default_value: 10s
    omit_empty: true

  - name: traceIPTables
    description: |-
      Instructs the enforcers to provide an iptables trace for the target processing
      unit.
    type: boolean
    exposed: true
    omit_empty: true

  - name: traceNetworkConnections
    description: |-
      Instructs the enforcer to send records for all
      network-initiated connections for the target processing unit.
    type: boolean
    exposed: true
    omit_empty: true
