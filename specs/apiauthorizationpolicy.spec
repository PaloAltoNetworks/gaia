# Model
model:
  rest_name: apiauthorizationpolicy
  resource_name: apiauthorizationpolicies
  entity_name: APIAuthorizationPolicy
  package: squall
  group: policy/authorization
  description: |-
    An API authorization defines the operations a user can perform in a
    namespace: `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, and/or `HEAD`.
    It is also possible to restrict the user to a subset of the APIs in the
    namespace by setting `authorizedIdentities`. An API authorization always
    propagates down to all the children of the current namespace.
  aliases:
  - apiauth
  - apiauths
  get:
    description: Retrieves the authorization with the given ID.
  update:
    description: Updates the authorization with the given ID.
  delete:
    description: Deletes the authorization with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@namespaced'
  - '@described'
  - '@disabled'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@hidden'
  - '@fallback'
  - '@schedulable'
  - '@timeable'
  - '@zoned'

# Indexes
indexes:
- - namespace
  - allObjectTags
  - disabled
- - namespace
  - allSubjectTags
  - disabled
- - namespace
  - allObjectTags
  - propagate
- - namespace
  - allSubjectTags
  - propagate

# Attributes
attributes:
  v1:
  - name: allSubjectTags
    description: This is a set of all subject tags for matching in the DB.
    type: list
    subtype: string
    stored: true

  - name: authorizedIdentities
    description: A list of roles assigned to the user.
    type: list
    exposed: true
    subtype: string
    stored: true
    required: true
    example_value:
    - '@auth:role=namespace.administrator'
    validations:
    - $validateFormatForAuthorizedIdentities

  - name: authorizedNamespace
    description: Defines the namespace the user is authorized to access.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /namespace

  - name: authorizedSubnets
    description: |-
      If set, the API authorization will only be valid if the request comes from one
      the declared subnets.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $optionalcidrs

  - name: expirationTime
    description: If set, the policy will be automatically deleted after the given time.
    type: time
    exposed: true
    stored: true
    getter: true
    setter: true

  - name: propagate
    description: Propagates the api authorization to all of its children.
    type: boolean
    stored: true
    read_only: true
    default_value: true
    getter: true
    setter: true
    orderable: true

  - name: subject
    description: A tag or tag expression that identifies the authorized user(s).
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    orderable: true
    validations:
    - $tagsExpression
    - $apiauthsubject
