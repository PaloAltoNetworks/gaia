# Model
model:
  rest_name: policyrefresh
  resource_name: policyrefreshs
  entity_name: PolicyRefresh
  package: squall
  group: core/policy
  description: Sent to a client as a push event when a policy refresh is needed on their side.

# Attributes
attributes:
  v1:
  - name: sourceID
    description: Contains the original ID of the updated object.
    type: string
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: sourceNamespace
    description: Contains the original namespace of the updated object.
    type: string
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: type
    description: Contains the policy type that is affected.
    type: string
    exposed: true
    stored: true
    filterable: true
    orderable: true
