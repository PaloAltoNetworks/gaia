# Model
model:
  rest_name: apiservice
  resource_name: apiservices
  entity_name: APIService
  package: squall
  description: |-
    APIService descibes a L4/L7 service and the corresponding implementation. It
    allows users to define their services, the APIs that they expose, the
    implementation of the service. These definitions can be used by network policy
    in order to define advanced controls based on the APIs.
  get: true
  update: true
  delete: true
  extends:
  - '@archivable'
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@metadatable'
  - '@named'

# Attributes
attributes:
  v1:
  - name: FQDN
    description: |-
      FQDN is the fully qualified domain name of the service. It is required
      for external API services. For HTTP services, FQND must match the host part
      of the URI that is used to call a service. It will be used for automatically
      generating service certificates for internal services.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: IPList
    description: |-
      IPList is the list of ip address or subnets of the service if available. This is
      an
      optional field and it can be automatically populated at runtime by the enforcers
      if DNS resolution is available.
    type: external
    exposed: true
    subtype: ip_list
    stored: true

  - name: JWTSigningCertificate
    description: |-
      JWTSigningCertificate is a certificate that can be used to validate user JWT in
      HTTP requests. This is an optional field, needed only if user JWT validation is
      required for this service. The certificate must be in PEM format.
    type: string
    exposed: true
    stored: true
    format: free

  - name: allServiceTags
    description: |-
      AllServiceTags is an internal object that summarizes all the implementedBy tags
      to accelerate database searches. It is not exposed.
    type: external
    subtype: tags_list
    stored: true

  - name: exposedAPIs
    description: ExposedAPIs is a list of API endpoints that are exposed for the service.
    type: external
    exposed: true
    subtype: exposed_api_list
    stored: true

  - name: external
    description: External is a boolean that indicates if this is an external service.
    type: boolean
    exposed: true
    stored: true
    default_value: "false"
    filterable: true
    orderable: true

  - name: externalServiceCA
    description: |-
      ExternalServiceCA is the certificate authority that the service is using. This
      is needed for external API services with private certificate authorities. The
      field is optional. If provided, this must be a valid PEM CA file.
    type: string
    exposed: true
    stored: true
    format: free

  - name: listeningPort
    description: |-
      listeningPort is the port that the application is listening to and
      it can be different than the ports describing the service. This is needed for
      port mapping use cases where there is private and public ports.
    type: string
    exposed: true
    stored: true
    allowed_chars: ^([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535)$
    example_value:
    - 80
    - 445:448

  - name: networkProtocol
    description: NetworkProtocol is the network protocol of the service.
    type: integer
    exposed: true
    stored: true
    default_value: "6"
    filterable: true
    max_value: 255
    min_value: 1
    orderable: true

  - name: ports
    description: |-
      Ports is a list of the public ports for the service. Ports are either
      exact match, or a range portMin:portMax. For HTTP services only exact match
      ports aresupported. These should be the ports that are used by other services
      to communicate with the defined service.
    type: external
    exposed: true
    subtype: port_list
    stored: true
    required: true
    example_value:
    - 80
    - 445:448

  - name: runtimeSelectors
    description: |-
      RuntimeSelectors is a list of tag selectors that identifies the Processing
      Units that will implement this service. The list can be empty for external
      services.
    type: external
    exposed: true
    subtype: target_tags
    stored: true
    example_value:
    - - a=a
      - b=b
    - - c=c

  - name: type
    description: |-
      Type is the type of the service (HTTP, TCP, etc). More types will be added to
      the system.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - HTTP
    - L3
    - TCP
    default_value: L3
    filterable: true
    orderable: true
