# Model
model:
  rest_name: cloudscaleset
  resource_name: cloudscalesets
  entity_name: CloudScaleSet
  package: yeul
  group: pcn/infrastructure
  description: |-
    A Cloud Scale Set represents a group of homogeneous VMs which are scaled up and
    down based on the load.
  aliases:
  - scaleset
  - scalesets
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
    description: Scale set related parameters.
    type: ref
    exposed: true
    subtype: cloudscalesetdata
    stored: true
    extensions:
      refMode: pointer
