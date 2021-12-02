# Model
model:
  rest_name: currentconnection
  resource_name: currentconnections
  entity_name: CurrentConnection
  package: guy
  group: core/enforcer
  description: A current connection.
  extends:
  - '@namespaced'
  - '@identifiable-not-stored'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: startTime
    description: The time the enforcer started tracking a connection.
    type: time
    exposed: true
    omit_empty: true

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