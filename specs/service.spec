# Model
model:
  rest_name: service
  resource_name: services
  entity_name: Service
  package: squall
  group: policy/services
  description: |-
    Defines a generic service object at layer 4 or layer 7 that encapsulates the
    description of a microservice. A service exposes APIs and can be implemented
    through third-party entities (such as a cloud provider) or through processing
    units.
  aliases:
  - srv
  get:
    description: Retrieves the service with the given ID.
    global_parameters:
    - $archivable
    - $propagatable
  update:
    description: Updates the service with the given ID.
  delete:
    description: Deletes the service with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@archivable'
  - '@described'
  - '@identifiable-stored'
  - '@named'
  - '@metadatable'
  - '@disabled'
  - '@timeable'
  - '@propagated'
  validations:
  - $serviceEntity

# Indexes
indexes:
- - allAPITags
- - namespace
  - allAPITags
- - allProcessingUnitsTags
- - namespace
  - allProcessingUnitsTags

# Attributes
attributes:
  v1:
  - name: IPs
    description: |-
      The list of IP addresses where the service can be accessed. This is an optional
      attribute and is only required if no host names are provided. The system will
      automatically resolve IP addresses from host names otherwise.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $optionalipaddresslist

  - name: JWTSigningCertificate
    description: |-
      PEM-encoded certificate that will be used to validate the user's JSON web token
      (JWT)
      in HTTP requests. This is an optional field, needed only if the
      `authorizationType`
      is set to `JWT`.
    type: string
    exposed: true
    stored: true

  - name: MTLSCertificateAuthority
    description: |-
      PEM-encoded certificate authority to use to verify client certificates. This
      only applies if `authorizationType` is set to `MTLS`. If it is not set,
      Microsegmentation Console's public signing certificate authority will be used.
    type: string
    exposed: true
    stored: true

  - name: OIDCCallbackURL
    description: |-
      This is an advanced setting. Optional OIDC callback URL. If you don't set it,
      the enforcer will autodiscover it. It will be
      `https://<hosts[0]|IPs[0]>/aporeto/oidc/callback`.
    type: string
    exposed: true
    stored: true

  - name: OIDCClientID
    description: OIDC Client ID. Only has effect if the `authorizationType` is set
      to `OIDC`.
    type: string
    exposed: true
    stored: true

  - name: OIDCClientSecret
    description: OIDC Client Secret. Only has effect if the `authorizationType` is
      set to `OIDC`.
    type: string
    exposed: true
    stored: true

  - name: OIDCProviderURL
    description: |-
      OIDC discovery endpoint. Only has effect if the `authorizationType`
      is set to `OIDC`.
    type: string
    exposed: true
    stored: true
    example_value: https://accounts.google.com

  - name: OIDCScopes
    description: |-
      Configures the scopes you want to request from the OIDC provider. Only has
      effect
      if `authorizationType` is set to `OIDC`.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - email
    - profile

  - name: TLSCertificate
    description: |-
      PEM-encoded certificate to expose to the clients for TLS. Only has effect and
      required if `TLSType` is set to `External`.
    type: string
    exposed: true
    stored: true

  - name: TLSCertificateKey
    description: |-
      PEM-encoded certificate key associated with `TLSCertificate`. Only has effect
      and required if `TLSType` is set to `External`.
    type: string
    exposed: true
    stored: true
    encrypted: true

  - name: TLSType
    description: |-
      Set how to provide a server certificate to the service.

      - `Aporeto`: Generate a certificate signed by the Microsegmentation
      Console public CA.
      - `LetsEncrypt`: Issue a certificate from Let's Encrypt.
      - `External`: Let you define your own certificate and key to use.
      - `None`: TLS is disabled (not recommended).
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Aporeto
    - LetsEncrypt
    - External
    - None
    default_value: Aporeto

  - name: allAPITags
    description: This is a set of all API tags for matching in the database.
    type: list
    subtype: string
    stored: true
    read_only: true

  - name: allProcessingUnitsTags
    description: This is a set of all selector tags for matching in the database.
    type: list
    subtype: string
    stored: true
    read_only: true

  - name: authorizationType
    description: |-
      Defines the user authorization type that should be used.

      - `None` (default): No authorization.
      - `JWT`:  Configures a simple JWT verification from the HTTP `Authorization`
      header.
      - `OIDC`: Configures OIDC authorization. You must then set
      `OIDCClientID`,`OIDCClientSecret`, `OIDCProviderURL`.
      - `MTLS`: Configures client certificate authorization. Then you can optionally
      use `MTLSCertificateAuthority`, otherwise Microsegmentation Console's public
      signing certificate will be used.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - None
    - JWT
    - OIDC
    - MTLS
    default_value: None

  - name: claimsToHTTPHeaderMappings
    description: |-
      Defines a list of mappings between claims and HTTP headers. When these mappings
      are defined, the enforcer will copy the values of the claims to the
      corresponding HTTP headers.
    type: refList
    exposed: true
    subtype: claimmapping
    stored: true
    extensions:
      refMode: pointer

  - name: endpoints
    description: |-
      Resolves the API endpoints that the service is exposing. Only valid during
      policy rendering.
    type: refList
    exposed: true
    subtype: endpoint
    read_only: true
    extensions:
      refMode: pointer

  - name: exposedAPIs
    description: |-
      Contains a tag expression that will determine which APIs a service is exposing.
      The APIs can be defined as the `RESTAPISpec` or similar specifications for other
      layer 7 protocols.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    example_value:
    - - package=p1
    validations:
    - $tagsExpression

  - name: exposedPort
    description: |-
      The port that the service can be accessed on. Note that this is different from
      the `port` attribute that describes the port that the service is actually
      listening on. For example if a load balancer is used, the `exposedPort` is
      the port that the load balancer is listening for the service, whereas the
      port that the implementation is listening on can be different.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 443
    max_value: 65535

  - name: exposedServiceIsTLS
    description: |-
      Indicates that the exposed service is TLS. This means that the enforcer has to
      initiate a TLS session in order to forward traffic to the service.
    type: boolean
    exposed: true
    stored: true
    default_value: false
    filterable: true
    orderable: true

  - name: external
    description: Indicates if this is an external service.
    type: boolean
    exposed: true
    stored: true
    default_value: false
    filterable: true
    orderable: true

  - name: hosts
    description: The host names that the service can be accessed on.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: port
    description: |-
      The port that the implementation of the service is listening to. It can be
      different than `exposedPort`. This is needed for port mapping use cases
      where there are private and public ports.
    type: integer
    exposed: true
    stored: true
    example_value: 443
    max_value: 65535

  - name: proxyProtocolEnabled
    description: Enable trust in proxy protocols header.
    type: boolean
    exposed: true
    stored: true

  - name: proxyProtocolPreserveHeader
    description: Do not strip proxy protocol header.
    type: boolean
    exposed: true
    stored: true

  - name: proxyProtocolSubnets
    description: Only allow proxy protocols from the given subnets .
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $optionalcidrs

  - name: publicApplicationPort
    description: |-
      A new virtual port that the service can be accessed on using HTTPS. Since the
      enforcer transparently inserts TLS in the application path, you might want
      to declare a new port where the enforcer listens for TLS. However, the
      application does not need to be modified and the enforcer will map the
      traffic to the correct application port. This is useful when
      an application is being accessed from a public network.
    type: integer
    exposed: true
    stored: true
    example_value: 443
    max_value: 65535

  - name: redirectURLOnAuthorizationFailure
    description: |-
      If this is set, the user will be redirected to that URL in case of any
      authorization failure, allowing you to provide a nice message to the user.
      The query parameter `?failure_message=<message>` will be added to that
      URL explaining the possible reason for the failure.
    type: string
    exposed: true
    stored: true

  - name: selectors
    description: |-
      A tag or tag expression that identifies the processing unit that implements this
      particular service.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    example_value:
    - - $identity=processingunit
    validations:
    - $tagsExpression

  - name: trustedCertificateAuthorities
    description: |-
      PEM-encoded certificate authorities to trust when additional hops are needed. It
      must be set if the service must reach a service marked as `external` or must go
      through an additional TLS termination point like a layer 7 load balancer.
    type: string
    exposed: true
    stored: true

  - name: type
    description: Type of service.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - HTTP
    - TCP
    - KubernetesSecrets
    - VaultSecrets
    default_value: HTTP

# Relations
relations:
- rest_name: httpresourcespec
  get:
    description: Retrieves the HTTP Resource exposed by this service.

- rest_name: processingunit
  get:
    description: Retrieves the processing units that implement this service.
