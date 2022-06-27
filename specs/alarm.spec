# Model
model:
  rest_name: alarm
  resource_name: alarms
  entity_name: Alarm
  package: sephiroth
  group: core/monitoring
  description: Represents an event requiring attention.
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
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
  - '@timeable'

# Ordering
default_order:
- :no-inherit
- updateTime

# Indexes
indexes:
- - kind
- - status
- - namespace
  - kind
- - namespace
  - kind
  - status

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
    description: Data represent user data related to the alarms.
    type: external
    exposed: true
    subtype: '[]map[string]string'
    stored: true

  - name: emails
    description: |-
      A list of recipients that should be emailed when this alarm is
      created.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - amir@aporeto.com
    - john@aporeto.com

  - name: kind
    description: |-
      Identifies the kind of alarm. If two alarms are created with the same
      identifier, then only the occurrence will be incremented.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: aporeto.alarm.kind
    orderable: true

  - name: lastLocalTimestamp
    description: Time and date of the alarm set by the enforcer.
    type: time
    exposed: true

  - name: occurrences
    description: Number of times this alarm has been seen.
    type: external
    exposed: true
    subtype: '[]time.Time'
    stored: true
    creation_only: true
    autogenerated: true

  - name: severity
    description: Severity of the alarm.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Low
    - Medium
    - High
    - Critical
    example_value: Low

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
    orderable: true
