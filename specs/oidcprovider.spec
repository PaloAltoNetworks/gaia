# Model
model:
  rest_name: oidcprovider
  resource_name: oidcproviders
  entity_name: OIDCProvider
  package: cactuar
  group: core/authentication
  description: |-
    Allows to declare a generic OpenID Connect provider that can be used in exchange
    for a Midgard token.
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@identifiable-stored'
  - '@timeable'

# Indexes
indexes:
- - :unique
  - parentID
  - name

# Attributes
attributes:
  v1:
  - name: clientID
    description: Unique client ID.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 6195189841830-0644ee9d89ef0644ee9d89examle.apps.googleusercontent.com

  - name: clientSecret
    description: Client secret associated with the client ID.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: Ytgbfjtj4652jHDFGls99jF

  - name: default
    description: |-
      If set, this will be the default OIDCProvider. There can be only one default
      provider in your account. When logging in with OIDC, if not provider name is
      given, the default will be used.
    type: boolean
    exposed: true
    stored: true

  - name: endpoint
    description: OIDC information endpoint.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: https://accounts.google.com

  - name: name
    description: Name of the provider.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: google

  - name: parentID
    description: ParentID contains the parent Vince account ID.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    filterable: true

  - name: parentName
    description: ParentName contains the name of the Vince parent Account.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    filterable: true

  - name: scopes
    description: List of scopes to allow.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - email
    - profile
