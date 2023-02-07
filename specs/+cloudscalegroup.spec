# Model
model:
  rest_name: cloudscalegroup
  resource_name: cloudscalegroups
  entity_name: CloudScaleGroup
  package: yeul
  group: pcn/infrastructure
  description: |-
    A Cloud Scale Group represents a group of VMs which are scaled up and
    down based on the load.
  aliases:
  - scalegroup
  - scalegroups
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
  extends:
  - '@base'
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@prismabase'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: parameters
    description: Scale group related parameters.
    type: ref
    exposed: true
    subtype: cloudscalegroupdata
    stored: true
    extensions:
      refMode: pointer
