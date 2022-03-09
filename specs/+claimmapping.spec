# Model
model:
  rest_name: claimmapping
  resource_name: claimmappings
  entity_name: ClaimMapping
  package: squall
  group: policy/services
  description: |-
    Allows you to map a claim in a token to an HTTP header. This can be useful when
    offloading authentication and authorization to Microsegmentation.
    Some applications may expect to receive information in the HTTP header.
  detached: true

# Attributes
attributes:
  v1:
  - name: claimName
    description: The name of the claim to map to the HTTP header. header.
    type: string
    exposed: true
    stored: true
    required: true
    allowed_chars: ^[a-zA-Z0-9-_/*#&@\+\$~:]+$
    allowed_chars_message: must be an alpha numerical character or '-', '_', '/',
      '*', '#', '&', '@', '_', '$' ~ or ':'
    example_value: email

  - name: targetHTTPHeader
    description: The HTTP header that will be the destination of the mapped claim.
    type: string
    exposed: true
    stored: true
    required: true
    allowed_chars: ^[a-zA-Z0-9-_/*#&@\+\$~:]+$
    allowed_chars_message: must be an alpha numerical character or '-', '_', '/',
      '*', '#', '&', '@', '_', '$' ~ or ':'
    example_value: X-Username
