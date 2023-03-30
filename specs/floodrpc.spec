# Model
model:
  rest_name: floodrpc
  resource_name: floodrpcs
  entity_name: FloodRPC
  package: yeul
  group: pcn/infrastructure
  description: |-
    Starts a flood remote procedural call for a given source/destination/payload
    triplet.
  private: true

# Attributes
attributes:
  v1:
  - name: floodParams
    description: The parameters needed to create and start a flooder.
    type: ref
    exposed: true
    subtype: floodparam
    extensions:
      refMode: pointer

  - name: nodeMakerConfigs
    description: The options needed to create nodemakers that are registered with
      a cached mux.
    type: ref
    exposed: true
    subtype: floodnodemakerconfig
    extensions:
      refMode: pointer

  - name: optionResultOmitTrails
    description: If set, trails will be omitted from the results.
    type: boolean
    exposed: true

  - name: results
    description: The flooding results.
    type: ref
    exposed: true
    subtype: floodresult
    extensions:
      refMode: pointer

  - name: sessionID
    description: |-
      a unique random string that is used to associate a cached mux nodemaker with
      successive requests.
    type: string
    exposed: true
    required: true
    example_value: 47a0479900504cb3ab4a1f626d174d2d
