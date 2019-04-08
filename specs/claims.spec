# Model
model:
  rest_name: claims
  resource_name: claims
  entity_name: Claims
  package: guy
  group: policy/networking
  description: This API represents the claims that accessed a service.
  get:
    description: Retrieves the object with the given ID.
  extends:
  - '@base'
  - '@identifiable-stored'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: content
    description: Content contains the raw JWT claims.
    type: external
    exposed: true
    subtype: map[string]string
    stored: true
    creation_only: true
    example_value:
      exp: 1553899021
      iat: 1553888221
      iss: https://accounts.acme.com
      sub: alice@acme.com

  - name: firstSeen
    description: firstSeen contains the date of the first appearance of the claims.
    type: time
    stored: true
    read_only: true
    autogenerated: true

  - name: hash
    description: |-
      XXH64 of the claims content. It will be used as ID. To compute a correct hash,
      you must first clob Content as an string array in the form `key=value`, sort it
      then apply the xxhash function.
    type: string
    exposed: true
    required: true
    example_value: "1134423925458173049"

  - name: lastSeen
    description: lastSeen contains the date of the last appearance of the claims.
    type: time
    stored: true
    read_only: true
    autogenerated: true
