# Model
model:
  rest_name: bucketentry
  resource_name: bucketentries
  entity_name: BucketEntry
  package: agrias
  group: applicationinsight/buckets
  description: Represent an entry within a bucket.
  detached: true

# Attributes
attributes:
  v1:
  - name: action
    description: Action applied to the type.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Accept
    - Reject
    example_value: Accept
    omit_empty: true

  - name: occurrences
    description: Number of times the entry was logged.
    type: integer
    exposed: true
    stored: true
    required: true
    default_value: 1

  - name: timestamp
    description: Time and date of the entry.
    type: time
    exposed: true
    stored: true
    orderable: true
    omit_empty: true
