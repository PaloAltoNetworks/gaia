# Model
model:
  rest_name: alarm
  resource_name: alarms
  entity_name: Alarm
  package: sephiroth
  description: An alarm represents an event requiring attention.
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
  - '@named'

# Attributes
attributes:
  v1:
  - name: content
    description: Content of the alarm.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: This is an alarm

  - name: data
    description: Data represent user data related to the alams.
    type: external
    exposed: true
    subtype: alarm_data
    stored: true

  - name: kind
    description: |-
      Kind identifies the kind of alarms. If two alarms are created with the same
      identifier, then only the occurrence will be incremented.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: aporeto.alarm.kind
    filterable: true
    orderable: true

  - name: occurrences
    description: Number of time this alarm have been seen.
    type: external
    exposed: true
    subtype: alarm_occurrences
    stored: true
    creation_only: true
    autogenerated: true

  - name: status
    description: Status of the alarm.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Acknowledged
    - Open
    - Resolved
    default_value: Open
    filterable: true
    orderable: true
