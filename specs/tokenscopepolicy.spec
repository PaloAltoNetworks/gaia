# Model
model:
  rest_name: tokenscopepolicy
  resource_name: tokenscopepolicies
  entity_name: TokenScopePolicy
  package: squall
  group: policy/services
  description: |-
    Defines a set of policies that allow customization of the authorization tokens
    issued by the Microsegmentation Console. This allows Microsegmentation tokens to
    be used by external applications.
  aliases:
  - tsp
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
  - name: allowedAudiences
    description: |-
      A list of audience values that are allowed when issuing a service token. An
      empty list will allow any audience values.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: assignedAudience
    description: |-
      The audience that should be assigned to a request if the caller is not
      requesting any specific audience.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: assignedScopes
    description: The list of scopes that the policy will assign.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true
    validations:
    - $stringlistnotempty

  - name: expirationTime
    description: If set the policy will be automatically deleted after the given time.
    type: time
    exposed: true
    stored: true
    getter: true
    setter: true

  - name: inheritedClaimKeys
    description: |-
      A list of claim keys that should be inherited from the claims of the caller to
      the assigned token. In this case, some of the caller claims will be propagated
      to resolved token.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: subject
    description: |-
      Defines the selection criteria that this policy must match on identity
      and scope request information.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    orderable: true
    validations:
    - $tagsExpression
