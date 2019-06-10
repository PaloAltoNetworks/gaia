# Model
model:
  rest_name: policyrenderer
  resource_name: policyrenderers
  entity_name: PolicyRenderer
  package: squall
  group: core
  description: |-
    Render is a low level api that allows to render policies of given tyoe for a
    given set of tags.

# Attributes
attributes:
  v1:
  - name: policies
    description: List of policies rendered for the given set of tags.
    type: refList
    exposed: true
    subtype: policyrule
    read_only: true
    autogenerated: true

  - name: processMode
    description: |-
      Define if the processMode should be using the object or subject. This only has
      effect when rendering a SSHAuthorizationPolicy for now.
    type: enum
    exposed: true
    allowed_choices:
    - Subject
    - Object
    default_value: Subject

  - name: tags
    description: List of tags of the object to render the hook policy for.
    type: list
    exposed: true
    subtype: string
    required: true
    example_value:
    - a=a
    - b=b

  - name: type
    description: Type of the policy to render.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - APIAuthorization
    - EnforcerProfile
    - File
    - Hook
    - Infrastructure
    - NamespaceMapping
    - Network
    - ProcessingUnit
    - Quota
    - Syscall
    - TokenScope
    - SSHAuthorization
    example_value: APIAuthorization
