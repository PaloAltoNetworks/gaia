# Model
model:
  rest_name: processingunit
  resource_name: processingunits
  entity_name: ProcessingUnit
  package: squall
  group: core/processingunit
  description: |-
    A processing unit represents anything that can compute. It can be a Docker
    container or a simple UNIX process. Processing units are created, updated, and
    deleted by
    the system as they come and go. You can only modify their tags. Processing units
    use network policies to define which other processing units or external
    networks they can communicate with and file access policies to define what file
    paths they can use.
  aliases:
  - pu
  - pus
  get:
    description: Retrieves the processing unit with the given ID.
    global_parameters:
    - $archivable
  update:
    description: Updates the processing unit with the given ID.
  delete:
    description: Deletes the processing unit with the given ID.
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
  - '@metadatable'
  - '@named'
  - '@timeable'
  - '@controllerable'
  extensions:
    commentFlags:
    - +k8s:openapi-gen=true

# Indexes
indexes:
- - namespace
  - operationalStatus
  - archived
- - namespace
  - normalizedTags
  - archived
- - namespace
  - nativeContextID
- - enforcerID
- - archived
  - updatetime
- - createtime

# Attributes
attributes:
  v1:
  - name: clientLocalID
    description: |-
      The local PUID set by enforcer. Enforcer may create a local PU if it cannot
      communicate with the Microsegmentation Console. When eventually the
      Microsegmentation Console is able to create the PU, the `clientLocalID` will be
      used to convert a CachedFlowReport containing a local PUID to a real FlowReport.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: collectInfo
    description: |-
      A value of `true` indicates to the enforcer that it needs to collect information
      for this processing unit.
    type: boolean
    exposed: true
    stored: true

  - name: collectedInfo
    description: |-
      Represents the latest information collected by the enforcer for this processing
      unit.
    type: external
    exposed: true
    subtype: map[string]string
    stored: true

  - name: datapathType
    description: |-
      The datapath type that processing units are implementing:
      - `Aporeto`: The enforcer is managing and handling the datapath.
      - `EnvoyAuthorizer`: The enforcer is serving Envoy-compatible gRPC APIs
      that for example can be used by an Envoy proxy to use the Microsegmentation PKI
      and implement Microsegmentation network policies. NOTE: The enforcer is not
      owning the datapath in this case. It is merely providing an authorizer API.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Aporeto
    - EnvoyAuthorizer
    default_value: Aporeto
    filterable: true

  - name: enforcementStatus
    description: |-
      Contains the state of the enforcer for the processing unit. `Inactive`
      (default): the enforcer is not enforcing any host service. `Active`: the
      enforcer is enforcing a host service. `Failed`: an error occurred during the
      enforcement attempt.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Active
    - Failed
    - Inactive
    default_value: Inactive
    filterable: true

  - name: enforcerID
    description: The ID of the enforcer associated with the processing unit.
    type: string
    exposed: true
    stored: true
    filterable: true

  - name: enforcerNamespace
    description: The namespace of the enforcer associated with the processing unit.
    type: string
    exposed: true
    stored: true
    filterable: true

  - name: signedHashOfHashedTags
    description: This field contains the signature of the hash of the hashed tags.
      field.
    type: string
    exposed: true
    stored: true

  - name: associatedLocalCAID
    description: AssociatedLocalCAID holds the remote ID of the certificate authority that has been used to sign
     the tags.
    type: string
    stored: true

  - name: image
    description: |-
      This field is deprecated and it is there for backward compatibility. Use
      `images` instead.
    type: string
    exposed: true
    deprecated: true

  - name: images
    description: List of images or executable paths used by the processing unit.
    type: list
    exposed: true
    subtype: string
    stored: true
    creation_only: true
    filterable: true

  - name: lastCollectionTime
    description: The date and time when the information was collected.
    type: time
    exposed: true
    stored: true

  - name: lastLocalTimestamp
    description: Time and date of the processing unit set by the enforcer.
    type: time
    exposed: true

  - name: lastSyncTime
    description: The date and time of the last policy resolution.
    type: time
    exposed: true
    stored: true
    autogenerated: true
    orderable: true

  - name: nativeContextID
    description: The Docker UUID or service PID.
    type: string
    exposed: true
    stored: true

  - name: networkServices
    description: |-
      The list of services that this processing unit has declared that it will be
      listening to,
      either in its activation command or by exposing the ports in a container
      manifest.
    type: refList
    exposed: true
    subtype: processingunitservice
    stored: true
    orderable: true
    validations:
    - $processingunitservices
    extensions:
      refMode: pointer

  - name: operationalStatus
    description: |-
      Operational status of the processing unit: `Initialized` (default), `Paused`,
      `Running`,
      `Stopped`, or `Terminated`.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Initialized
    - Paused
    - Running
    - Stopped
    - Terminated
    default_value: Initialized
    filterable: true

  - name: previousOperationalStatus
    description: Holds the previous operational status if it has changed.
    type: external
    exposed: true
    subtype: _pu_operational_status
    read_only: true
    autogenerated: true
    transient: true
    omit_empty: true

  - name: tracing
    description: Indicates if this processing unit must be placed in tracing mode.
    type: ref
    exposed: true
    subtype: tracemode
    extensions:
      refMode: pointer

  - name: type
    description: |-
      Type of processing unit: `APIGateway`, `Docker`, `Host`, `HostService`,
      `LinuxService`, `WindowsService`, `RKT`, `User`, or `SSHSession`.
    type: enum
    exposed: true
    stored: true
    creation_only: true
    allowed_choices:
    - APIGateway
    - Docker
    - Host
    - HostService
    - LinuxService
    - WindowsService
    - RKT
    - User
    - SSHSession
    - ECS
    example_value: Docker
    filterable: true

  - name: vulnerabilities
    description: List of vulnerabilities affecting this processing unit.
    type: list
    subtype: string
    stored: true
    deprecated: true
    filterable: true

  - name: vulnerabilityLevel
    description: List of vulnerabilities affecting this processing unit.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    deprecated: true
    filterable: true
    transient: true

  - name: wireTagToHashTag
    description: |-
      Contains the list of wire tags that must go on the wire and their mapping to
      corresponding hashes.
    type: external
    exposed: true
    subtype: map[string]string

  - name: hashOfHashedTags
    description: |-
      Contains the hash of the hashed wire tags.
    type: string
    exposed: true
    stored: true

# Relations
relations:
- rest_name: pingprobe
  create:
    description: Create a ping probe.

- rest_name: poke
  get:
    description: |-
      Sends a poke empty object. This will send a snapshot of the processing unit to
      the time series database.
    parameters:
      entries:
      - name: enforcementStatus
        description: If set, changes the enforcement status of the processing unit
          alongside with the poke.
        type: enum
        allowed_choices:
        - Failed
        - Inactive
        - Active

      - name: forceFullPoke
        description: If set, it will trigger a full poke (slower).
        type: boolean

      - name: notify
        description: Can be sent to trigger a `ProcessingUnitRefresh` event that will
          be handled by the enforcer. If this is set, all other additional parameters
          will be ignored.
        type: boolean

      - name: status
        description: If set, changes the status of the processing unit alongside with
          the poke.
        type: enum
        allowed_choices:
        - Initialized
        - Paused
        - Running
        - Stopped
        example_value: Running

      - name: ts
        description: time of report. If not set, local server time will be used.
        type: time

      - name: zhash
        description: Can be set to help backend target the correct shard where the
          processing unit is stored.
        type: integer

- rest_name: processingunitrefresh
  create:
    description: Sends a Processing Unit Refresh command.

- rest_name: renderedpolicy
  get:
    description: Retrieves the policies for the processing unit.
    parameters:
      entries:
      - name: renderer
        description: Select the network policy renderer to use.
        type: enum
        allowed_choices:
        - v1
        - v2

- rest_name: service
  get:
    description: Retrieves the services used by a processing unit.

- rest_name: vulnerability
  get:
    description: Retrieves the vulnerabilities affecting the processing unit.
    global_parameters:
    - $propagatable
