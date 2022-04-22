# Model
model:
  rest_name: cloudscalesetdata
  resource_name: cloudscalesetdata
  entity_name: CloudScaleSetData
  package: yeul
  group: pcn/infrastructure
  description: Manages the properties associated to a Virtual Machine Scale Set.
  detached: true

# Attributes
attributes:
  v1:
  - name: networkConfigs
    description: Scale set related parameters.
    type: refList
    exposed: true
    subtype: cloudnetworkconfigdata
    stored: true
    extensions:
      refMode: pointer
