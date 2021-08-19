# Model
model:
  rest_name: validaterql
  resource_name: validaterql
  entity_name: ValidateRQLReq
  package: larl
  group: thought/crime
  description: TODO.

# Attributes
attributes:
  v1:
  - name: policyType
    description: TODO.
    type: string
    exposed: true
    required: true
    example_value: network

  - name: prismaId
    description: TODO.
    type: integer
    exposed: true
    required: true
    example_value: 0

  - name: query
    description: TODO.
    type: string
    exposed: true
    required: true
    example_value: config from network where ...

  - name: searchType
    description: TODO.
    type: string
    exposed: true
    required: true
    example_value: network_config
