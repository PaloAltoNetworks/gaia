# Model
model:
  rest_name: suspiciousactivity
  resource_name: suspiciousactivities
  entity_name: SuspiciousActivity
  package: karl
  group: pc/alerts
  description: Represents a suspicious activity found on the platform.
  aliases:
  - susact
  get:
    description: Retrieves the suspicious activity with the given ID.
  update:
    description: Updates the suspicious activity with the given ID.
  delete:
    description: Deletes the suspicious activity with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@identifiable-stored'
  - '@migratable'
  - '@namespaced'
  - '@timeable'
  - '@zoned'

# Indexes
indexes:
- - namespace
  - sourceID
- - namespace
  - sourceName

# Attributes
attributes:
  v1:
  - name: categories
    description: The list of category names.
    type: list
    exposed: true
    subtype: integer
    stored: true
    required: true
    example_value:
    - 56
    - 100

  - name: counter
    description: Number of times this suspicious activity is found.
    type: integer
    exposed: true
    stored: true

  - name: firstOccurrenceTime
    description: The timestamp when the suspicious activity first occurred.
    type: time
    exposed: true
    stored: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"

  - name: lastOccurrenceTime
    description: The timestamp when the suspicious activity last occurred.
    type: time
    exposed: true
    stored: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"

  - name: risk
    description: The level of risk.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Low
    - Medium
    - High
    - Unknown
    default_value: Unknown

  - name: sourceID
    description: The identifier of the source.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: 62aa2888a76fe8dc9efa1000

  - name: sourceName
    description: The name of the source.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: google.com

  - name: sourceNamespace
    description: The namespace of the source.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: /my/namespace

  - name: sourceType
    description: The type of the source.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: flowreport
