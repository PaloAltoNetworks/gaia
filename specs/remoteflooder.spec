# Model
model:
  rest_name: remoteflooder
  resource_name: remoteflooders
  entity_name: RemoteFlooder
  package: yeul
  group: pcn/infrastructure
  description: Starts a flooder for a given source/destination/payload triplet.

# Attributes
attributes:
  v1:
  - name: FloodParams
    description: The parameters needed to create and start a flooder.
    type: ref
    exposed: true
    subtype: floodparam
    extensions:
      refMode: pointer

  - name: NodeMakerConfigs
    description: The options needed to create nodemakers that are registered with
      a cached mux.
    type: ref
    exposed: true
    subtype: floodnodemakerconfig
    extensions:
      refMode: pointer

  - name: Results
    description: The flooding results.
    type: ref
    exposed: true
    subtype: floodresult
    extensions:
      refMode: pointer

  - name: optionResultOmitTrails
    description: If set, trails will be omitted from the results.
    type: boolean
    exposed: true

  - name: sessionID
    description: |-
      a unique random string that is used to associate a cached mux nodemaker with
      successive requests.
    type: string
    exposed: true
    required: true
    example_value: 47a0479900504cb3ab4a1f626d174d2d
