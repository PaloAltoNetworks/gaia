# Model
model:
  rest_name: tracemode
  resource_name: tracemode
  entity_name: TraceMode
  package: squall
  group: core/enforcer
  description: Represents the tracing mode to apply to a processing unit.
  detached: true
  extensions:
    commentFlags:
    - +k8s:openapi-gen=true

# Attributes
attributes:
  v1:
  - name: IPTables
    description: Instructs the enforcers to provide an iptables trace for a processing
      unit.
    type: boolean
    exposed: true
    stored: true

  - name: applicationConnections
    description: |-
      Instructs the enforcer to send records for all
      application-initiated connections.
    type: boolean
    exposed: true
    stored: true

  - name: interval
    description: |-
      Determines the length of the time interval that the trace must be
      enabled, using [Golang duration
      syntax](https://golang.org/pkg/time/#example_Duration).
    type: string
    exposed: true
    stored: true
    default_value: 10s

  - name: networkConnections
    description: |-
      Instructs the enforcer to send records for all
      network-initiated connections.
    type: boolean
    exposed: true
    stored: true
