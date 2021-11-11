# Model
model:
  rest_name: automationaction
  resource_name: automationactions
  entity_name: AutomationAction
  package: sephiroth
  group: integration/automation
  description: Action that can be used in automations.
  aliases:
  - act
  get:
    description: Retrieves the automation action with the given ID.
  extends:
  - '@described'
  - '@named'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: entitlements
    description: Contains the entitlements needed for executing the function.
    type: external
    exposed: true
    subtype: _automation_entitlements

  - name: function
    description: Function contains the code.
    type: string
    exposed: true

  - name: key
    description: Contains the unique identifier key for the action.
    type: string
    exposed: true

  - name: parameters
    description: Contains the computed parameters.
    type: external
    exposed: true
    subtype: map[string]interface{}

  - name: steps
    description: Contains all the steps with parameters.
    type: refList
    exposed: true
    subtype: uistep
    extensions:
      refMode: pointer
