# Model
model:
  rest_name: import
  resource_name: import
  entity_name: Import
  package: squall
  group: core
  description: Imports an export of policies and related objects into the namespace.

# Attributes
attributes:
  v1:
  - name: data
    description: Data to import.
    type: ref
    exposed: true
    subtype: export
    required: true
    example_value:
      externalnetworks:
      - associatedTags:
        - ext:net=tcp
        description: Represents all TCP traffic on any port
        entries:
        - 0.0.0.0/0
        name: all-tcp
        servicePorts:
        - tcp/1:65535
      - associatedTags:
        - ext:net=udp
        description: Represents all UDP traffic on any port
        entries:
        - 0.0.0.0/0
        name: all-udp
        servicePorts:
        - udp/1:65535
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

  - name: mode
    description: |-
      The mode defines how the import reacts:
      - `Import` (default): Creates all objects within the request. Exiting objects
      that do not match their hash are destroyed and recreated. Objects that fail to
      import are dropped.
      - `ImportRemoveOnFailure`: Creates all objects within the request. Removes all
      objects touched if any object fails to import.
      - `Remove`: Removes all objects within the request. Removed objects that fail
      are ignored.
    type: enum
    exposed: true
    allowed_choices:
    - Import
    - ImportRemoveOnFailure
    - Remove
    default_value: Import
