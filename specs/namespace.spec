# Model
model:
  rest_name: namespace
  resource_name: namespaces
  entity_name: Namespace
  package: squall
  group: core/namespace
  description: |-
    A namespace represents the core organizational unit of the system. All objects
    always exist in a single namespace. A namespace can also have child namespaces.
    They can be used to split the system into organizations, business units,
    applications, services or any combination you like.
  aliases:
  - ns
  get:
    description: Retrieves the namespace with the given ID.
  update:
    description: Updates the namespace with the given ID.
  delete:
    description: Deletes the namespace with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@timeable'

# Ordering
default_order:
- name

# Indexes
indexes:
- - name
- - namespace
  - name

# Attributes
attributes:
  v1:
  - name: JWTCertificateType
    description: |-
      JWTCertificateType defines the JWT signing certificate that must be created
      for this namespace. If the type is none no certificate will be created.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - RSA
    - EC
    - None
    default_value: None

  - name: JWTCertificates
    description: |-
      JWTCertificates hold the certificates used to sign tokens for this namespace.
      This is map indexed by the ID of the certificate.
    type: external
    exposed: true
    subtype: map[string]string
    stored: true
    read_only: true
    autogenerated: true

  - name: SSHCA
    description: The SSH certificate authority used by the namespace.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: SSHCAEnabled
    description: "If `true`, an SSH certificate authority (CA) will be generated for
      the\nnamespace. This CA \ncan be deployed in SSH server to validate SSH certificates
      issued by the\nplatform."
    type: boolean
    exposed: true
    stored: true
    orderable: true

  - name: associatedLocalCAID
    description: AssociatedLocalCAID holds the remote ID of the certificate authority
      to use.
    type: string
    stored: true
    read_only: true

  - name: associatedSSHCAID
    description: The remote ID of the SSH certificate authority to use.
    type: string
    exposed: true
    stored: true
    read_only: true

  - name: customZoning
    description: "Defines if the namespace should inherit its parent zone. If this
      property is set\nto `false`, \nthe `zoning` property will be ignored and the
      namespace will have the same zone\nas its parent."
    type: boolean
    exposed: true
    stored: true
    creation_only: true

  - name: localCA
    description: The certificate authority used by this namespace.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: localCAEnabled
    description: |-
      Defines if the namespace should use a local certificate
      authority (CA). Switching it off and on again will regenerate a new CA.
    type: boolean
    exposed: true
    stored: true
    orderable: true

  - name: name
    description: The name of the namespace.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    allowed_chars: ^[a-zA-Z0-9-_/]+$
    allowed_chars_message: must only contain alpha numerical characters, '-' or '_'
    example_value: mynamespace
    filterable: true
    getter: true
    setter: true
    orderable: true

  - name: networkAccessPolicyTags
    description: |-
      List of tags that will be added to every `or` clause of all network access
      policies in the namespace and its children.
    type: list
    exposed: true
    subtype: string
    stored: true
    deprecated: true
    validations:
    - $tags

  - name: serviceCertificateValidity
    description: |-
      Determines the length of validity of certificates issued in this namespace using
      [Golang duration syntax](https://golang.org/pkg/time/#example_Duration). Default
      value is `1h`.
    type: string
    exposed: true
    stored: true
    default_value: 1h
    validations:
    - $timeDuration

  - name: zoning
    description: Defines what zone the namespace should live in.
    type: integer
    exposed: true
    stored: true
    creation_only: true
    getter: true
    setter: true

# Relations
relations:
- rest_name: oauthinfo
  get:
    description: Retrieves the OAUTH info for this namespace.
    parameters:
      entries:
      - name: mode
        description: When set to type `OIDC` it will return the data as a raw JSON
          object and not an Aporeto compatible API.
        type: enum
        allowed_choices:
        - oidc

- rest_name: oauthkey
  get:
    description: Retrieves the OAUTH info for this namespace.
    parameters:
      entries:
      - name: mode
        description: When set to `OIDC` it will return the data as a raw JSON object
          and not an Aporeto compatible API.
        type: enum
        allowed_choices:
        - oidc

- rest_name: trustedca
  get:
    description: Returns the list of trusted CAs for this namespace.
    parameters:
      entries:
      - name: type
        description: Type of certificate to get.
        type: enum
        allowed_choices:
        - Any
        - X509
        - SSH
        - JWT
        default_value: Any
