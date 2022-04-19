# Model
model:
  rest_name: enforcerprofilemappingpolicy
  resource_name: enforcerprofilemappingpolicies
  entity_name: EnforcerProfileMappingPolicy
  package: squall
  group: policy/enforcerconfig
  description: |-
    Allows you to map an enforcer profile to one or more enforcers.
    The mapping can also be propagated down to the child namespace.
  aliases:
  - enfpols
  - enfpol
  - epm
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
  - '@propagated'
  - '@fallback'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: object
    description: The tag or tag expression that identifies the enforcer profile to
      be mapped.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    example_value:
    - - a=a
      - b=b
    - - c=c
    validations:
    - $tagsExpression

  - name: subject
    description: |-
      The tag or tag expression that identifies the enforcers that should implement
      the mapped profile.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    example_value:
    - - a=a
      - b=b
    - - c=c
    validations:
    - $tagsExpression

# Relations
relations:
- rest_name: enforcer
  get:
    description: Returns the list of enforcers affected by an enforcer profile mapping.

- rest_name: enforcerprofile
  get:
    description: |-
      Returns the list of enforcer profiles that an enforcer profile mapping
      matches.
