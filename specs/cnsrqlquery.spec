# Model
model:
  rest_name: cnsrqlquery
  resource_name: cnsrqlquery
  entity_name: CNSrqlquery
  package: vargid
  group: pcn/infrastructure
  description: A CNS endpoint which will generate the RQL query for the given alert-id.

# Attributes
attributes:
  v1:
  - name: alertID
    exposed_name: alertID
    description: ID of the alert request.
    type: string
    exposed: true

  - name: policyID
    exposed_name: policyID
    description: ID of the alert request.
    type: string
    exposed: true

  - name: query
    exposed_name: query
    description: CNS rql query for alertID.
    type: string
    exposed: true

  - name: valid
    exposed_name: valid
    description: Whether the response to request is valid.
    type: boolean
    exposed: true
    omit_empty: true
