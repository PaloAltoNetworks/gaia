# Model
model:
  rest_name: longlivedconnectionsreport
  resource_name: longlivedconnectionsreport
  entity_name: LonglLivedConnectionsReport
  package: guy
  group: core/enforcer
  description: new longlived connections report.
  extends:
  - '@flow'
  - '@namespaced'
  - '@identifiable-not-stored'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: collectID
    description: Corelated to the enforcer refresh that this was generated in response to.
    type: string
    exposed: true
    omit_empty: true

  - name: connectionStartTime
    description: The time the connection was estalished.
    type: time
    exposed: true
    omit_empty: true
