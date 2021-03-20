# Model
model:
  rest_name: cloudvpc
  resource_name: cloudvpcs
  entity_name: CloudVPC
  package: yeul
  group: prisma/infrastructure
  description: |-
    A CloudVPC represents a VPC as defined in an cloud provider (AWS/Azure/GCP etc).
    The VPC is essentially an L3 routing domain with at least one subnet attached
    and it defines an isolated network.
  aliases:
  - vpc
  - vpcs
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@prismabase'

# Attributes
attributes:
  v1:
  - name: parameters
    description: VPC related parameters.
    type: ref
    exposed: true
    subtype: cloudvpcdata
    stored: true
    extensions:
      refMode: pointer
