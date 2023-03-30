# Model
model:
  rest_name: floodresult
  resource_name: floodresults
  entity_name: FloodResult
  package: yeul
  group: pcn/infrastructure
  description: The results of flooding.
  private: true
  detached: true

# Attributes
attributes:
  v1:
  - name: errors
    description: Errors that were encountered while flooding.
    type: list
    exposed: true
    subtype: string
    read_only: true
    autogenerated: true

  - name: pathExists
    description: Set if flooder found at least one trail.
    type: boolean
    exposed: true
    read_only: true
    autogenerated: true

  - name: trails
    description: |-
      A list of trails the flooder found. WARNING: this will eventually go away as we
      should transmit the tree. We keep it this way for backwards compatibility with
      existing code for the sake of speed.
    type: external
    exposed: true
    subtype: '[][]detacheddecoy'
    read_only: true
    autogenerated: true
    extensions:
      refMode: pointer
