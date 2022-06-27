# Model
model:
  rest_name: cloudnetworkconfigdata
  resource_name: cloudnetworkconfigdata
  entity_name: CloudNetworkConfigData
  package: yeul
  group: pcn/infrastructure
  description: Manages the network properties associated to a Virtual Machine Scale Set.
  detached: true

# Attributes
attributes:
  v1:
  - name: ipConfigurations
    description: IP configurations of the NICs in the Scale Set.
    type: refList
    exposed: true
    subtype: cloudipconfiguration
    stored: true
    extensions:
      refMode: pointer

  - name: networkConfigName
    description: Network configuration name of the NIC associated in the Scale Set.
    type: string
    exposed: true
    stored: true
