# Model
model:
  rest_name: namespacemappingpolicy
  resource_name: namespacemappingpolicies
  entity_name: NamespaceMappingPolicy
  package: squall
  group: core/namespace
  description: |-
    A namespace mapping defines the namespace a processing unit should
    be placed when it is created, based on its tags.  When an enforcer creates
    a new processing unit, the system will place it in its own namespace if no
    matching namespace mapping can be found. If one match is found, then the
    processing unit will be bumped down to the namespace declared in the namespace
    mapping. If it
    finds in that child namespace another matching namespace mapping, then
    the processing unit will be bumped down again, until it reaches a namespace with
    no matching namespace mappings.  This is very useful to dispatch processes and
    containers
    into a particular namespace, based on a lot of factors. For example, you can put
    in place a
    quarantine namespace mapping that will grab all processing units with excessive
    vulnerabilities.
  aliases:
  - nspolicy
  - nspolicies
  - nsmap
  - nsmaps
  get:
    description: Retrieves the mapping with the given ID.
  update:
    description: Updates the mapping with the given ID.
  delete:
    description: Deletes the mapping with the given ID.
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
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: mappedNamespace
    description: The namespace to map the `subject` to.
    type: string
    exposed: true
    stored: true
    required: true
    allowed_chars: ^[a-zA-Z0-9-_/]+$
    allowed_chars_message: must only contain alpha numerical characters, '-' or '_'
    example_value: /blue/namespace
    orderable: true

  - name: subject
    description: A tag or tag expression identifying the entity to be mapped.
    type: external
    exposed: true
    subtype: '[][]string'
    example_value:
    - - color=blue
    orderable: true
    validations:
    - $tagsExpression
