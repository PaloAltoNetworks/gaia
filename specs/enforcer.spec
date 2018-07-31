# Model
model:
  rest_name: enforcer
  resource_name: enforcers
  entity_name: Enforcer
  package: squall
  description: |-
    An Enforcer Profile contains a configuration for a Enforcer. It contains various
    parameters, like what network should not policeds, what processing units should
    be ignored based on their tags and so on. It also contains more advanced
    parameters to fine tune the Agent. A Enforcer will decide what profile to apply
    using aEnforcer Profile Mapping Policy. This policy will decide according the
    Enforcer's tags what profile to use. If an Enforcer tags are matching more than
    a single policy, it will refuse to start. Some parameters will be applied
    directly to a running agent, some will need to restart it.
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
  extends:
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@metadatable'
  - '@named'

# Attributes
attributes:
  v1:
  - name: FQDN
    description: FQDN contains the fqdn of the server where the enforcer is running.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: server1.domain.com
    filterable: true
    format: free
    orderable: true

  - name: certificate
    description: Certificate is the certificate of the enforcer.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    format: free
    orderable: true

  - name: certificateExpirationDate
    description: CertificateExpirationDate is the expiration date of the certificate.
    type: time
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: certificateKey
    description: |-
      CertificateKey is the secret key of the enforcer. Returned only when a enforcer
      is created or the certificate is updated.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    format: free
    orderable: true

  - name: certificateRequest
    description: |-
      CertificateRequest can be used to generate a certificate from that CSR instead
      of letting the server generate your private key for you.
    type: string
    exposed: true
    format: free
    transient: true

  - name: certificateRequestEnabled
    description: |-
      If set during creation,the server will not initially generate a certificate. In
      that case you have to provide a valid CSR through certificateRequest during an
      update.
    type: boolean
    exposed: true
    stored: true
    creation_only: true

  - name: collectInfo
    description: CollectInfo indicates to the enforcer it needs to collect information.
    type: boolean
    exposed: true
    stored: true

  - name: collectedInfo
    description: CollectedInfo represents the latest info collected by the enforcer.
    type: external
    exposed: true
    subtype: collected_info
    stored: true

  - name: currentVersion
    description: |-
      CurrentVersion holds the enforcerd binary version that is currently associated
      to this object.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: enforcerProfileID
    description: Contains the ID of the profile used by the instance of enforcerd.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: lastCollectionTime
    description: |-
      LastCollectionTime represents the date and time when the info have been
      collected.
    type: time
    exposed: true
    stored: true

  - name: lastSyncTime
    description: LastSyncTime holds the last heart beat time.
    type: time
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: localCA
    description: |-
      LocalCA contains the initial chain of trust for the enforcer. This valud is only
      given when you retrieve a single enforcer.
    type: string
    exposed: true
    autogenerated: true
    format: free
    transient: true

  - name: operationalStatus
    description: OperationalStatus tells the status of the enforcer.
    type: enum
    exposed: true
    allowed_choices:
    - Registered
    - Connected
    - Disconnected
    - Initialized
    - Unknown
    autogenerated: true
    default_value: Registered
    filterable: true
    transient: true

  - name: publicToken
    description: |-
      PublicToken is the public token of the server that will be included in the
      datapath and its signed by the private CA.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    format: free
    transient: true

  - name: startTime
    description: |-
      startTime holds the time this enforcerd was started. This time-stamp is reported
      by the enforcer and is is preserved across disconnects.
    type: time
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: updateAvailable
    description: Tells if the the version of enforcerd is outdated and should be updated.
    type: boolean
    exposed: true
    stored: true
    filterable: true
    orderable: true

# Relations
relations:
- rest_name: enforcerprofile
  get:
    description: Returns the enforcer profile that must be used by an enforcer.

- rest_name: poke
  get:
    description: Sends a poke empty object. This is used to ensure an enforcer is
      up and running.
