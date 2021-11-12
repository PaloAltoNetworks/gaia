# Model
model:
  rest_name: enforcerprofile
  resource_name: enforcerprofiles
  entity_name: EnforcerProfile
  package: squall
  group: policy/enforcerconfig
  description: |-
    Allows you to create reusable configuration profiles for your enforcers.
    Enforcer profiles contain various startup information that can (for some) be
    updated live. Enforcer profiles are assigned to enforcers using a enforcer
    profile mapping.
  aliases:
  - profile
  - profiles
  get:
    description: Retrieves the enforcer profile with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the enforcer profile with the given ID.
  delete:
    description: Deletes the enforcer profile with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@named'
  - '@metadatable'
  - '@propagated'
  - '@timeable'
  validations:
  - $enforcerprofile
  extensions:
    commentFlags:
    - +k8s:openapi-gen=true

# Attributes
attributes:
  v1:
  - name: excludedInterfaces
    description: |-
      Ignore traffic with a source or destination matching the specified
      interfaces.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: excludedNetworks
    description: |-
      Ignore any networks specified here and do not even report any flows.
      This can be useful for excluding localhost loopback traffic, ignoring
      traffic to the Kubernetes API, and using Microsegmentation for SSH only.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: ignoreExpression
    description: |-
      A tag expression that identifies processing units to ignore. This can be
      useful to exclude `kube-system` pods, AWS EC2 agent pods, and third-party
      agents.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    validations:
    - $tagsExpression

  - name: kubernetesMetadataExtractor
    description: This field is kept for backward compatibility for enforcers <= 3.5.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - KubeSquall
    - PodAtomic
    - PodContainers
    default_value: PodAtomic
    deprecated: true

  - name: kubernetesSupportEnabled
    description: This field is kept for backward compatibility for enforcers <= 3.5.
    type: boolean
    exposed: true
    stored: true
    deprecated: true

  - name: metadataExtractor
    description: This field is kept for backward compatibility for enforcers <= 3.5.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Docker
    - ECS
    - Kubernetes
    default_value: Docker
    deprecated: true

  - name: syslogEnabled
    description: Enables syslog functionality of enforcers using this enforcerprofile.
    type: boolean
    exposed: true
    stored: true

  - name: syslogEndpoint
    description: Contains the remote endpoint to dispatch the syslog messages.
    type: string
    exposed: true
    stored: true

  - name: syslogEndpointTLSClientCertificate
    description: PEM-encoded certificate to expose to the clients for TLS.
    type: string
    exposed: true
    stored: true

  - name: syslogEndpointTLSClientCertificateKey
    description: |-
      PEM-encoded certificate key associated with
      `syslogEndpointTLSClientCertificate`.
    type: string
    exposed: true
    stored: true
    encrypted: true

  - name: syslogEndpointTLSServerCA
    description: PEM-encoded server CA certificate.
    type: string
    exposed: true
    stored: true

  - name: syslogFacility
    description: Contains the list of supported syslog facilities.
    type: integer
    exposed: true
    stored: true
    default_value: 1
    max_value: 23

  - name: syslogFormat
    description: Contains the list of supported syslog message format.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Auto
    - BSD
    - IETF
    default_value: Auto

  - name: targetNetworks
    description: |-
      If empty, the enforcer auto-discovers the TCP networks. Auto-discovery
      works best in Kubernetes and OpenShift deployments. You may need to manually
      specify the TCP networks if middle boxes exist that do not comply with
      [TCP Fast Open RFC 7413](https://tools.ietf.org/html/rfc7413).
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: targetUDPNetworks
    description: |-
      If empty, the enforcer enforces all UDP networks. This works best when all UDP
      networks have enforcers. If some UDP networks do not have enforcers, you
      may need to manually specify the UDP networks that should be enforced.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: trustedCAs
    description: |-
      List of trusted certificate authorities. If empty, the main chain of trust
      will be used.
    type: list
    exposed: true
    subtype: string
    stored: true
