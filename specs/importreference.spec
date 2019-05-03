# Model
model:
  rest_name: importreference
  resource_name: importreferences
  entity_name: ImportReference
  package: Yuna
  group: core
  description: Import and keep a reference.
  aliases:
  - importref
  - impref
  get:
    description: Retrieves the object with the given ID.
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
  - name: data
    description: The data to import.
    type: ref
    exposed: true
    subtype: export
    stored: true
    required: true
    example_value:
      externalnetworks:
      - associatedTags:
        - ext:net=tcp
        description: Represents all TCP traffic on any port
        entries:
        - 0.0.0.0/0
        name: all-tcp
        protocols:
        - tcp
      - associatedTags:
        - ext:net=udp
        description: Represents all UDP traffic on any port
        entries:
        - 0.0.0.0/0
        name: all-udp
        protocols:
        - udp
      networkaccesspolicies:
      - action: Allow
        description: Allows all communication from pu to pu, tcp and udp
        logsEnabled: true
        name: allow-all-communication
        object:
        - - $identity=processingunit
        - - ext:net=tcp
        - - ext:net=udp
        subject:
        - - $identity=processingunit
    extensions:
      refMode: pointer
