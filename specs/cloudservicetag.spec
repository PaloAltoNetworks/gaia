# Model
model:
  rest_name: cloudservicetag
  resource_name: cloudserviceta
  entity_name: CloudServiceTag
  package: yeul
  group: pcn/infrastructure
  description: |-
    Provides the expanded list of network prefixes for a given Service Tag provided
    by the platform.
  get:
    description: Retrieves the IP prefixes mapping for a given service tag.
    global_parameters:
    - $filtering
  update:
    description: Updates the IP prefixes for a given service tag.
  delete:
    description: Deletes the IP prefixes for a service tag.
  extends:
  - '@zoned'
  - '@base'
  - '@migratable'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@named'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: IPPrefixList
    description: The IP Prefixes associated with a Service Tag.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $optionalcidrs

  - name: cloudType
    description: Cloud Type of the Service Tag.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - AZURE
    - AWS
    default_value: AWS
