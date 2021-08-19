# Model
model:
  rest_name: validaterql
  resource_name: validaterql
  entity_name: ValidateRQL
  package: larl
  group: core/rql
  description: Used to validate RQL queries.

# Attributes
attributes:
  v1:
  - name: error
    description: The error message if the query fails validation.
    type: string
    exposed: true

  - name: policyType
    description: The type of the policy.
    type: string
    exposed: true

  - name: prismaId
    description: The Prisma Cloud ID.
    type: integer
    exposed: true

  - name: query
    description: The query to validate.
    type: string
    exposed: true
    required: true
    example_value: network from microseg.flow_record where id == '1'

  - name: searchType
    description: The search type.
    type: string
    exposed: true

  - name: status
    description: The http status code of the response.
    type: integer
    exposed: true

  - name: timestamp
    description: The timestamp of the query in nanoseconds.
    type: integer
    exposed: true
