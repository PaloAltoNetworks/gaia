# Model
model:
  rest_name: floodparam
  resource_name: floodparams
  entity_name: FloodParam
  package: yeul
  group: pcn/infrastructure
  description: Parameters to start a flooder.
  detached: true

# Attributes
attributes:
  v1:
  - name: destination
    description: The destination NodeUID where the flooder should stop.
    type: ref
    exposed: true
    subtype: nodeuid
    extensions:
      refMode: pointer

  - name: optionContinueOnErrNotPermitted
    description: If set, the flooder will continue if a node returns ErrNotPermitted.
    type: boolean
    exposed: true

  - name: optionOnlyReportResultsWithErrNotPermitted
    description: |-
      If set, only denied paths will be reported. This will also force-set
      optionContinueOnErrNotPermitted.
    type: boolean
    exposed: true

  - name: payload
    description: The payload which the flooder should use.
    type: ref
    exposed: true
    subtype: payload
    extensions:
      refMode: pointer

  - name: source
    description: The source NodeUID where the flooder should start.
    type: ref
    exposed: true
    subtype: nodeuid
    extensions:
      refMode: pointer
