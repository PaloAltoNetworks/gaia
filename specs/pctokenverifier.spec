# Model
model:
  rest_name: pctokenverifier
  resource_name: pctokenverifiers
  entity_name: PCTokenVerifier
  package: larl
  group: pc/tokenverification
  description: |-
    Verifies the given PC token (provided in the password field) and returns its
    claims.
  private: true

# Attributes
attributes:
  v1:
  - name: password
    description: The provided token (name is to comply with HTTP source requests).
    type: string
    exposed: true
    required: true
    example_value: secret
