# Model
model:
  rest_name: enforcer
  resource_name: enforcers
  entity_name: Enforcer
  package: squall
  group: core/enforcer
  description: |-
    An Enforcer contains all parameters associated with a registered enforcer. The
    object is mainly by maintained by the enforcers themselves. Users can read the
    object in order to understand the current status of the enforcers.
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
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@zonable'
  - '@timeable'

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
    orderable: true

  - name: certificate
    description: Certificate is the certificate of the enforcer.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    orderable: true

  - name: certificateExpirationDate
    description: |-
      CertificateExpirationDate is the expiration date of the certificate. This is an
      internal attribute, not exposed in the API.
    type: time
    read_only: true
    autogenerated: true

  - name: certificateKey
    description: |-
      CertificateKey is the certificate key of the enforcer. This is an internal
      attribute, not exposed in the API.
    type: string
    read_only: true
    autogenerated: true

  - name: certificateRequest
    description: |-
      If not empty during a create or update generation, the provided CSR will be
      validated and signed by the backend providing a renewed certificate.
    type: string
    exposed: true
    example_value: |-
      -----BEGIN CERTIFICATE REQUEST-----
      MIICvDCCAaQCAQAwdzELMAkGA1UEBhMCVVMxDTALBgNVBAgMBFV0YWgxDzANBgNV
      BAcMBkxpbmRvbjEWMBQGA1UECgwNRGlnaUNlcnQgSW5jLjERMA8GA1UECwwIRGln
      aUNlcnQxHTAbBgNVBAMMFGV4YW1wbGUuZGlnaWNlcnQuY29tMIIBIjANBgkqhkiG
      9w0BAQEFAAOCAQ8AMIIBCgKCAQEA8+To7d+2kPWeBv/orU3LVbJwDrSQbeKamCmo
      wp5bqDxIwV20zqRb7APUOKYoVEFFOEQs6T6gImnIolhbiH6m4zgZ/CPvWBOkZc+c
      1Po2EmvBz+AD5sBdT5kzGQA6NbWyZGldxRthNLOs1efOhdnWFuhI162qmcflgpiI
      WDuwq4C9f+YkeJhNn9dF5+owm8cOQmDrV8NNdiTqin8q3qYAHHJRW28glJUCZkTZ
      wIaSR6crBQ8TbYNE0dc+Caa3DOIkz1EOsHWzTx+n0zKfqcbgXi4DJx+C1bjptYPR
      BPZL8DAeWuA8ebudVT44yEp82G96/Ggcf7F33xMxe0yc+Xa6owIDAQABoAAwDQYJ
      KoZIhvcNAQEFBQADggEBAB0kcrFccSmFDmxox0Ne01UIqSsDqHgL+XmHTXJwre6D
      hJSZwbvEtOK0G3+dr4Fs11WuUNt5qcLsx5a8uk4G6AKHMzuhLsJ7XZjgmQXGECpY
      Q4mC3yT3ZoCGpIXbw+iP3lmEEXgaQL0Tx5LFl/okKbKYwIqNiyKWOMj7ZR/wxWg/
      ZDGRs55xuoeLDJ/ZRFf9bI+IaCUd1YrfYcHIl3G87Av+r49YVwqRDT0VDV7uLgqn
      29XI1PpVUNCPQGn9p/eX6Qo7vpDaPybRtA2R7XLKjQaF9oXWeCUqy1hvJac9QFO2
      97Ob1alpHPoZ7mWiEuJwjBPii6a9M9G30nUo39lBi1w=
      -----END CERTIFICATE REQUEST-----
    transient: true

  - name: collectInfo
    description: CollectInfo indicates to the enforcer it needs to collect information.
    type: boolean
    exposed: true
    stored: true

  - name: collectedInfo
    description: CollectedInfo represents the latest info collected by the enforcer.
    type: external
    exposed: true
    subtype: map[string]string
    stored: true

  - name: currentVersion
    description: |-
      CurrentVersion holds the enforcerd binary version that is currently associated
      to this object.
    type: string
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: enforcementStatus
    description: Status of enforcement for PU managed directly by enforcerd, like
      Host PUs.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Inactive
    - Active
    - Failed
    default_value: Inactive
    filterable: true

  - name: lastCollectionTime
    description: |-
      LastCollectionTime represents the date and time when the info have been
      collected.
    type: time
    exposed: true
    stored: true

  - name: lastPokeTime
    description: Last poke is the time when the enforcer got last poked.
    type: time
    stored: true

  - name: lastSyncTime
    description: LastSyncTime holds the last heart beat time.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: lastValidHostServices
    description: |-
      LastValidHostServices is a read only attribute that stores the list valid host
      services that have been applied to this enforcer. This list might be different
      from the list retrieved through policy, if the dynamically calculated list leads
      into conflicts.
    type: refList
    subtype: hostservice
    stored: true
    autogenerated: true

  - name: localCA
    description: |-
      LocalCA contains the initial chain of trust for the enforcer. This valud is only
      given when you retrieve a single enforcer.
    type: string
    exposed: true
    autogenerated: true
    transient: true

  - name: machineID
    description: |-
      MachineID holds a unique identifier for every machine as detected by the
      enforcer. It is based on hardware information such as the SMBIOS UUID, MAC
      addresses of interfaces or cloud provider IDs.
    type: string
    exposed: true
    stored: true
    example_value: 3F23E8DF-C56D-45CF-89B8-A867F3956409
    filterable: true

  - name: operationalStatus
    description: OperationalStatus tells the status of the enforcer.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Registered
    - Connected
    - Disconnected
    - Initialized
    default_value: Registered
    filterable: true

  - name: publicToken
    description: |-
      PublicToken is the public token of the server that will be included in the
      datapath and its signed by the private CA.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    transient: true

  - name: startTime
    description: |-
      startTime holds the time this enforcerd was started. This time-stamp is reported
      by the enforcer and is is preserved across disconnects.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: subnets
    description: Local subnets of the enforcer running.
    type: list
    exposed: true
    subtype: string

  - name: unreachable
    description: |-
      Control plane will set this value to true if it hasn't heard from the enforcer
      for more than 5m.
    type: boolean
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: updateAvailable
    description: Tells if the the version of enforcerd is outdated and should be updated.
    type: boolean
    exposed: true
    stored: true
    orderable: true

# Relations
relations:
- rest_name: auditprofile
  get:
    description: Returns a list of the audit profiles that must be applied to this
      enforcer.

- rest_name: enforcerprofile
  get:
    description: Returns the enforcer profile that must be used by an enforcer.

- rest_name: hostservice
  get:
    description: Returns a list of the host services policies that apply to this enforcer.
    parameters:
      entries:
      - name: appliedServices
        description: Valid when retrieved for a given enforcer and returns the applied
          services.
        type: boolean

      - name: setServices
        description: Instructs the backend to cache the services that were resolved.
          services.
        type: boolean

- rest_name: poke
  get:
    description: Sends a poke empty object. This is used to ensure an enforcer is
      up and running.
    parameters:
      entries:
      - name: cpuload
        description: Deprecated.
        type: float
        example_value: 1000

      - name: enforcementStatus
        description: If set, changes the enforcement status of the enforcer alongside
          with the poke.
        type: enum
        allowed_choices:
        - Failed
        - Inactive
        - Active

      - name: forceFullPoke
        description: If set, it will trigger a full poke (slower).
        type: boolean

      - name: memory
        description: Deprecated.
        type: integer
        example_value: 1000

      - name: processes
        description: Deprecated.
        type: integer
        example_value: 10

      - name: sessionClose
        description: If set, terminates a session for an enforcer.
        type: boolean

      - name: sessionID
        description: If set, sends the current session ID of an enforcer.
        type: string
        example_value: "1233"

      - name: status
        description: If set, changes the status of the enforcer alongside with the
          poke.
        type: enum
        allowed_choices:
        - Registered
        - Connected
        - Disconnected

      - name: ts
        description: time of report. If not set, local server time will be used.
        type: time

      - name: version
        description: If set, version of the current running enforcer.
        type: string
        example_value: v1.10
