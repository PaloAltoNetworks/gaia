# Model
model:
  rest_name: cloudalertrecordprocessor
  resource_name: cloudalertrecordprocessors
  entity_name: CloudAlertRecordProcessor
  package: vargid
  group: pcn/infrastructure
  description: Control message model to raise/resolve/generate cloud alert records.
  extends:
  - '@namespaced'
  validations:
  - $cloudalertrecordprocessor

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
