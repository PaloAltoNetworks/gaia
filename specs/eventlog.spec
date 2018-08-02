# Model
model:
  rest_name: eventlog
  resource_name: eventlogs
  entity_name: EventLog
  package: leon
  description: This api allows to report various event on any objects.

# Attributes
attributes:
  v1:
  - name: category
    description: Category of the log.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: enforcerd:policy

  - name: content
    description: Content of the log.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: Unable to activate docker container xyz because abc.

  - name: date
    description: Creation date of the eventlog.
    type: time
    exposed: true
    stored: true
    creation_only: true
    autogenerated: true
    orderable: true

  - name: level
    description: Represent the level of the log .
    type: enum
    exposed: true
    stored: true
    creation_only: true
    allowed_choices:
    - Debug
    - Info
    - Warning
    - Error
    - Critical
    default_value: Info

  - name: namespace
    description: Namespace tag attached to an entity.
    type: string
    exposed: true
    stored: true
    read_only: true
    creation_only: true
    autogenerated: true
    getter: true
    setter: true
    orderable: true
    primary_key: true

  - name: targetID
    description: |-
      ID of the object this eventlog is attached to. The object must be in the same
      namespace than the eventlog.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: xxx-xxx-xxx-xxx

  - name: targetIdentity
    description: Identity of the object this eventlog is attached to.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: processingunit

  - name: title
    description: Title of the eventlog.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: Error while activating processing unit.
