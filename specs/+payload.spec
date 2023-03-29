# Model
model:
  rest_name: payload
  resource_name: payloads
  entity_name: Payload
  package: yeul
  group: pcn/infrastructure
  description: Represents a neocna Payload.
  detached: true

# Attributes
attributes:
  v1:
  - name: Content
    description: The raw content of the payload.
    type: string
    exposed: true
    omit_empty: true

  - name: DestinationPortEnd
    description: The end of the destination port range.
    type: integer
    exposed: true
    omit_empty: true

  - name: DestinationPortStart
    description: The start of the destination port range.
    type: integer
    exposed: true
    omit_empty: true

  - name: SourcePortEnd
    description: The end of the source port range.
    type: integer
    exposed: true
    omit_empty: true

  - name: SourcePortStart
    description: The start of the source port range.
    type: integer
    exposed: true
    omit_empty: true

  - name: protocol
    description: The protocol that is used for flooding.
    type: string
    exposed: true
    omit_empty: true
