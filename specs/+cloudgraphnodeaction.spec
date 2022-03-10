# Model
model:
  rest_name: cloudgraphnodeaction
  resource_name: cloudgraphactions
  entity_name: CloudGraphNodeAction
  package: yeul
  group: pcn/infrastructure
  description: Describes the action and corresponding policy that resulted in this
    decision.
  detached: true

# Attributes
attributes:
  v1:
  - name: action
    description: The action that is been applied for the particular destination.
    type: string
    exposed: true

  - name: policyID
    description: The ID of the policies that were used in the path.
    type: string
    exposed: true
