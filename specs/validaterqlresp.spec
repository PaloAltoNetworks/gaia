# Model
model:
  rest_name: validaterqlresp
  resource_name: validaterqlresp
  entity_name: ValidateRQLResp
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
    example_value: invalid quert

  - name: status
    description: TODO.
    type: integer
    exposed: true
    example_value: 200

  - name: timestamp
    description: TODO.
    type: integer
    exposed: true
    example_value: 1232432432
