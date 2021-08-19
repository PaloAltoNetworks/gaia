# Model
model:
  rest_name: validaterql
  resource_name: validaterql
  entity_name: ValidateRQL
  package: larl
  group: thought/crime
  description: TODO.

# Attributes
attributes:
  v1:
  - name: error
    description: TODO.
    type: string
    exposed: true

  - name: policyType
    description: TODO.
    type: string
    exposed: true

  - name: prismaId
    description: TODO.
    type: integer
    exposed: true

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

  - name: status
    description: TODO.
    type: integer
    exposed: true

  - name: timestamp
    description: TODO.
    type: integer
    exposed: true
