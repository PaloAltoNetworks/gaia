# Model
model:
  rest_name: cnsrqlquery
  resource_name: cnsrqlquery
  entity_name: CNSRQLQuery
  package: vargid
  group: pcn/infrastructure
  description: A CNS endpoint which will generate the RQL query for the given alert-id.

# Attributes
attributes:
  v1:
  - name: alertID
    exposed_name: alertID
    description: ID of the query request.
    type: string
    exposed: true

  - name: policyID
    exposed_name: policyID
    description: The policy for which the alert was generated.
    type: string
    exposed: true

  - name: query
    exposed_name: query
    description: The rql query for the alert.
    type: string
    exposed: true

  - name: valid
    exposed_name: valid
    description: Whether the response to request is valid.
    type: boolean
    exposed: true
    omit_empty: true
