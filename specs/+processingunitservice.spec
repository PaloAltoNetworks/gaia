# Model
model:
  rest_name: processingunitservice
  resource_name: processingunitservices
  entity_name: ProcessingUnitService
  package: squall
  group: policy/processingunits
  description: Represents a service attached to a processing unit.
  detached: true
  extensions:
    commentFlags:
    - +k8s:openapi-gen=true

# Attributes
attributes:
  v1:
  - name: ports
    description: Contains the list of allowed ports and ranges.
    type: string
    exposed: true
    stored: true
    read_only: true
    deprecated: true
    orderable: true

  - name: protocol
    description: Protocol used by the service.
    type: integer
    exposed: true
    stored: true

  - name: targetPorts
    description: List of single ports or range (xx:yy).
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $ports
