# Model
model:
  rest_name: cloudgraphasset
  resource_name: cloudgraphassets
  entity_name: CloudGraphAsset
  package: yeul
  group: pcn/infrastructure
  description: |-
    Returns a data structure representing the graph of all cloud nodes and their
    connections for the given source unified asset IDs.

# Attributes
attributes:
  v1:
  - name: cloudGraphs
    description: The set of nodes and paths from the given source to the given destination.
    type: refMap
    exposed: true
    subtype: cloudgraph
    read_only: true
    autogenerated: true
    extensions:
      refMode: pointer

  - name: direction
    description: Direction of the network path.
    type: enum
    exposed: true
    allowed_choices:
    - Inbound
    - Outbound
    default_value: Outbound

  - name: errors
    description: The error message if encountered any.
    type: list
    exposed: true
    subtype: string
    read_only: true
    autogenerated: true
    omit_empty: true

  - name: executionTimestamp
    description: Result of the graph execution timestamp.
    type: time
    exposed: true
    read_only: true
    autogenerated: true
    orderable: true

  - name: prismaCloudPolicyID
    description: Prisma Cloud Policy ID.
    type: string
    exposed: true
    required: true
    example_value: ad23603d-754e-4499-8988-b80178785898
    omit_empty: true

  - name: unifiedAssetIDs
    description: Prisma Cloud Unified Asset IDs.
    type: list
    exposed: true
    subtype: string
    required: true
    example_value: 4b882952f1a26c8a3ce16ee90af3cac9
    omit_empty: true
    validations:
    - $validatemaxunifiedassetids
