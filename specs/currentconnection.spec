# Model
model:
  rest_name: currentconnection
  resource_name: currentconnections
  entity_name: CurrentConnection
  package: guy
  group: core/enforcer
  description: A current connection.
  extends:
  - '@identifiable-stored'
  detached: true

# Attributes
attributes:
  v1:
  - name: duration
    description: The duration of the tracked connection.
    type: string
    exposed: true
    omit_empty: true
    validations:
    - $timeDuration

  - name: existing
    description: Was the connection existing when the enforcer started.
    type: boolean
    exposed: true
    omit_empty: true

  - name: flow
    description: The flow report for this connection.
    type: ref
    exposed: true
    subtype: flowreport
    omit_empty: true
    extensions:
      refMode: pointer

  - name: sourcePort
    description: Port of the source.
    type: integer
    exposed: true
    max_value: 65536
    omit_empty: true

  - name: startTime
    description: The time the enforcer started tracking the connection.
    type: time
    exposed: true
    omit_empty: true
