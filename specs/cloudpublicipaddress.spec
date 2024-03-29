# Model
model:
  rest_name: cloudpublicipaddress
  resource_name: cloudpublicipaddresses
  entity_name: CloudPublicIPAddress
  package: yeul
  group: pcn/infrastructure
  description: |-
    A CloudPublicIPAddress represents a PublicIPAddress as defined in an Azure cloud
    provider.
  aliases:
  - publicipaddress
  - publicipaddresses
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
    description: PublicIPAddress related parameters.
    type: ref
    exposed: true
    subtype: cloudaddress
    stored: true
    extensions:
      refMode: pointer
