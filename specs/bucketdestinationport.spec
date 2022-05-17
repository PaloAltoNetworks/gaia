# Model
model:
  rest_name: bucketdestinationport
  resource_name: bucketdestinationports
  entity_name: BucketDestinationPort
  package: agrias
  group: applicationinsight/buckets
  description: |-
    Holds the destination ports of flows across 24 hours (beginning of a day to
    the end).
  extends:
  - '@identifiable-stored'
  - '@namespaced'
  - '@zoned'

# Indexes
indexes:
- - namespace
  - date

# Attributes
attributes:
  v1:
  - name: date
    description: The date that the bucket tracks.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: "2022-05-09"

  - name: ports
    description: Ports encountered and number of occurrences seen.
    type: external
    exposed: true
    subtype: map[string]int
    stored: true
    omit_empty: true
