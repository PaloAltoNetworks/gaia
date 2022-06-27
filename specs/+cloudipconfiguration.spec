# Model
model:
  rest_name: cloudipconfiguration
  resource_name: cloudipconfiguration
  entity_name: CloudIPConfiguration
  package: yeul
  group: pcn/infrastructure
  description: Manages the properties associated to an IP configuration.
  detached: true

# Attributes
attributes:
  v1:
  - name: IPConfigName
    description: IP configuration of the NICs associated to the Scale Set.
    type: string
    exposed: true
    stored: true

  - name: addresses
    description: |-
      List of IP addresses/subnets (IPv4 or IPv6) associated with the
      interface.
    type: refList
    exposed: true
    subtype: cloudaddress
    stored: true
    extensions:
      refMode: pointer

  - name: backendPool
    description: Backend pool of the load balancer (if any) fronting the Scale Set.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: hasPublicIP
    description: If the IP Configuration has a public IP.
    type: boolean
    exposed: true
    stored: true

  - name: subnet
    description: Subnet of the NIC associated to the Scale Set.
    type: string
    exposed: true
    stored: true
