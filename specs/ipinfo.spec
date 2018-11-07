# Model
model:
  rest_name: ipinfo
  resource_name: ipinfos
  entity_name: IPInfo
  package: canyon
  description: This apis allows to resolve information from an IP address.

# Attributes
attributes:
  v1:
  - name: IP
    description: The IP resolved.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: error
    description: Eventual error that happened during resolution.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: records
    description: List of DNS records associated to that IP.
    type: external
    exposed: true
    subtype: whois_info
    read_only: true
    autogenerated: true
