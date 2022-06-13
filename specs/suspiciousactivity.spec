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

  - name: lastOccurrenceTime
    description: The timestamp when the suspicious activity last occurred.
    type: time
    exposed: true
    stored: true

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
    creation_only: true

  - name: sourceName
    description: The name of the source.
    type: string
    exposed: true
    stored: true
    creation_only: true

  - name: sourceNamespace
    description: The namespace of the source.
    type: string
    exposed: true
    stored: true
    creation_only: true

  - name: sourceType
    description: The type of the source.
    type: string
    exposed: true
    stored: true
    creation_only: true
