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
  - name: attachedEntities
    description: ID of associated objects with this interface configuration.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - lb1-backendpool-1
    - lb2-backendpool-3

  - name: networkConfigs
    description: Scale set related parameters.
    type: refList
    exposed: true
    subtype: cloudnetworkconfigdata
    stored: true
    extensions:
      refMode: pointer
