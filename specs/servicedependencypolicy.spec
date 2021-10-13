# Model
model:
  rest_name: servicedependencypolicy
  resource_name: servicedependencypolicies
  entity_name: ServiceDependencyPolicy
  package: squall
  group: policy/services
  description: |-
    Allows you to define a service dependency policy where a set of processing units
    as
    defined
    by their tags require access to specific services.
  aliases:
  - srvdep
  - srvdeps
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
  - '@namespaced'
  - '@described'
  - '@disabled'
  - '@identifiable-not-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'
  - '@schedulable'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: object
    description: Object of the service dependency.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

  - name: subject
    description: Subject of the service dependency.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

# Relations
relations:
- rest_name: processingunit
  get:
    description: Returns the list of processing units that depend on an service.

- rest_name: service
  get:
    description: Returns the list of external services that are targets of service dependency.
