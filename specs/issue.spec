# Model
model:
  rest_name: issue
  resource_name: issue
  entity_name: Issue
  package: midgard
  description: This API issues a new token according to given data.
  extends:
  - '@realm'

# Attributes
attributes:
  v1:
  - name: data
    description: Data contains additional data. The value depends on the issuer type.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: metadata
    description: Metadata contains various additional information. Meaning depends
      on the realm.
    type: external
    exposed: true
    subtype: metadata
    orderable: true

  - name: token
    description: Token is the token to use for the registration.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: validity
    description: |-
      Validity configures the max validity time for a token. If it is bigger than the
      configured max validity, it will be capped.
    type: string
    exposed: true
    stored: true
    allowed_chars: ^([0-9]+h[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+h[0-9]+s|[0-9]+h[0-9]+m|[0-9]+s|[0-9]+h|[0-9]+m)$
    allowed_chars_message: Must be a valid duration like Ns or Ns or Nh
    default_value: 24h
    orderable: true
