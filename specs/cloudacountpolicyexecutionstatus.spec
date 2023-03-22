# Model
model:
  rest_name: cloudacountpolicyexecutionstatus
  resource_name: cloudacountpolicyexecutionstatuses
  entity_name: CloudAccountPolicyExecutionStatus
  package: vargid
  group: pcn/infrastructure
  description: |-
    CloudAccountPolicyExecutionStatus represents the execution status of cloud
    policies relates to the cloud account.
  extends:
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@timeable'
  - '@named'

# Indexes
indexes:
- - namespace
  - accountID

# Attributes
attributes:
  v1:
  - name: accountID
    description: Cloud account ID of the account.
    type: string
    exposed: true
    stored: true

  - name: successfulPolicyExecutionTimestampMap
    description: Mapping of last successful execution timestamp for every cloud policy.
    type: external
    exposed: true
    subtype: map[string]time
    stored: true
