# Model
model:
  rest_name: cloudalertrule
  resource_name: cloudalertsrule
  entity_name: CloudAlertRule
  package: vargid
  group: pcn/infrastructure
  description: |-
    Represents an Alert rule along with policies and accounts associated
    with the Alert rule.
  get:
    description: Retrieves the Alert Rule with a given id.
    global_parameters:
    - $filtering
  update:
    description: Updates the Alert Rule with a given id.
  delete:
    description: Deletes the Alert Rule with a given id.
  extends:
  - '@base'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@timeable'
  - '@named'
  - '@zoned'
  - '@described'

# Indexes
indexes:
- - namespace
  - prismaCloudAlertRuleID
- - :unique
  - prismaCloudAlertRuleID

# Attributes
attributes:
  v1:
  - name: enabled
    description: Defines whether the Alert rule is enabled.
    type: boolean
    exposed: true
    subtype: boolean
    stored: true

  - name: prismaCloudAlertRuleID
    description: Prisma Cloud Alert Rule id.
    type: string
    exposed: true
    subtype: string
    stored: true

  - name: prismaCloudAlertRuleScopeLastChangedOn
    description: Result of the last changed time to the prisma cloud alert rule.
    type: integer
    exposed: true
    stored: true
    default_value: 0

  - name: prismaCloudPolicyIDs
    description: List of Policy IDs associated to an Alert rule.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: regions
    description: List of regions where the Alert rule is enforced.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: targetAccountIDs
    description: List of accounts associated to an Alert rule.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: tenantPrismaID
    description: Prisma ID of the tenant in which the Alert Rule is created.
    type: string
    exposed: true
    subtype: string
    stored: true
