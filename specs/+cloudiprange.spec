# Model
model:
  rest_name: cloudiprange
  resource_name: cloudiprange
  entity_name: CloudIPRange
  package: yeul
  group: pcn/infrastructure
  description: Parameters associated with a cloud ips range.
  detached: true

# Attributes
attributes:
  v1:
  - name: endIP
    description: End IP for the range.
    type: string
    exposed: true
    stored: true

  - name: startIP
    description: Starting IP for the range.
    type: string
    exposed: true
    stored: true
