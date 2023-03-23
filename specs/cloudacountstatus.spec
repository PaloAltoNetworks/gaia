# Model
model:
  rest_name: cloudacountstatus
  resource_name: cloudacountstatuses
  entity_name: CloudAccountStatus
  package: vargid
  group: pcn/infrastructure
  description: CloudAccountStatus represents the status for a cloud account.
  get:
    description: Retrieve the cloud account status with the given ID.
  update:
    description: Updates the cloud account status with the given ID.
  delete:
    description: Deletes the cloud account status with the given ID.
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
