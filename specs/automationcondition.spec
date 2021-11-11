# Model
model:
  rest_name: automationcondition
  resource_name: automationconditions
  entity_name: AutomationCondition
  package: sephiroth
  group: integration/automation
  description: Condition that can be used in automations.
  aliases:
  - con
  get:
    description: Retrieves the automation condition with the given ID.
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
    description: Contains the unique identifier key for the condition.
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
