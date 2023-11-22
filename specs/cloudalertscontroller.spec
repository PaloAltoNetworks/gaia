# Model
model:
  rest_name: cloudalertscontroller
  resource_name: cloudalertscontrollers
  entity_name: CloudAlertsController
  package: vargid
  group: pcn/infrastructure
  description: Control message model to raise/resolve/generate cloud alerts.
  extends:
  - '@namespaced'
  validations:
  - $cloudalertscontroller

# Attributes
attributes:
  v1:
  - name: action
    description: Action type to perform.
    type: enum
    exposed: true
    allowed_choices:
    - Generate
    - Raise
    - Resolve
    default_value: Generate

  - name: alertResolveWaitDuration
    description: |-
      Duration to wait to resolve an alert. This attribute is only supported with
      action 'Resolve'.
    type: string
    exposed: true
    default_value: 0s
    validations:
    - $timeDuration

  - name: cloudAccountIDs
    description: |-
      IDs of cloud accounts to scan and generate alerts. When left empty all cloud
      accounts in the tenant are considered. This attribute is only supported with
      action 'Generate'.
    type: list
    exposed: true
    subtype: string

  - name: prismaCloudPolicyIDs
    description: |-
      Prisma Cloud policy IDs to scan and generate alerts. When left empty all
      policies
      in the tenant are considered. This attribute is only supported with action
      'Generate'.
    type: list
    exposed: true
    subtype: string
