# Model
model:
  rest_name: floodpayload
  resource_name: floodpayloads
  entity_name: FloodPayload
  package: yeul
  group: pcn/infrastructure
  description: Represents a neocna Payload.
  private: true
  detached: true

# Attributes
attributes:
  v1:
  - name: content
    description: The raw content of the payload.
    type: string
    exposed: true
    omit_empty: true

  - name: destinationPortEnd
    description: The end of the destination port range.
    type: integer
    exposed: true
    omit_empty: true

  - name: destinationPortStart
    description: The start of the destination port range.
    type: integer
    exposed: true
    omit_empty: true

  - name: protocol
    description: The protocol that is used for flooding.
    type: string
    exposed: true
    omit_empty: true

  - name: sourcePortEnd
    description: The end of the source port range.
    type: integer
    exposed: true
    omit_empty: true

  - name: sourcePortStart
    description: The start of the source port range.
    type: integer
    exposed: true
    omit_empty: true
