# Model
model:
  rest_name: testresource
  resource_name: testresources
  entity_name: TestResource
  package: basefetcher
  group: neocna/test
  description: Represents a resource meant to test neocna basefetcher.
  private: true
  extends:
  - '@identifiable-stored'
  - '@migratable'
  - '@namespaced'
  - '@zoned'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: UID
    description: The unique identifier of the resource.
    type: string
    exposed: true
    stored: true
    example_value: some-unique-identifier

  - name: rawData
    description: The raw data that represents the resource.
    type: external
    exposed: true
    subtype: '[]byte'
    stored: true
    required: true
    example_value: |-
      {
        "UID": "some-unique-identifier",
        "kind": "VPN-Gateway"
        ...
      }
