# Model
model:
  rest_name: automationaction
  resource_name: automationactions
  entity_name: AutomationAction
  package: sephiroth
  group: integration/automation
  description: Action that can be used in automations.
  aliases:
  - autoact
  get:
    description: Retrieves the automation action with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the automation action with the given ID.
  delete:
    description: Deletes the automation action with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@named'
  - '@metadatable'
  - '@propagated'
  - '@timeable'

# Indexes
indexes:
- - namespace
  - key

# Attributes
attributes:
  v1:
  - name: entitlements
    description: Contains the entitlements needed for executing the function.
    type: external
    exposed: true
    subtype: _automation_entitlements
    stored: true

  - name: function
    description: Function contains the code.
    type: string
    exposed: true
    stored: true

  - name: key
    description: Contains the unique identifier key for the action.
    type: string
    exposed: true
    stored: true
    creation_only: true

  - name: parameters
    description: Contains the computed parameters.
    type: external
    exposed: true
    subtype: map[string]interface{}
    stored: true

  - name: steps
    description: Contains all the steps with parameters.
    type: refList
    exposed: true
    subtype: uistep
    stored: true
    extensions:
      refMode: pointer
