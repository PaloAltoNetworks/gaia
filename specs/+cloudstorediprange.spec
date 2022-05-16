# Model
model:
  rest_name: cloudstorediprange
  resource_name: cloudstorediprange
  entity_name: CloudStoredIPRange
  package: yeul
  group: pcn/infrastructure
  description: Parameters associated with a stored cloud ips range.
  detached: true

# Attributes
attributes:
  v1:
  - name: endIP
    description: End IP for the range.
    type: external
    exposed: true
    subtype: network
    stored: true

  - name: startIP
    description: Starting IP for the range.
    type: external
    exposed: true
    subtype: network
    stored: true
