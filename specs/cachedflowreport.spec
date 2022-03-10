# Model
model:
  rest_name: cachedflowreport
  resource_name: cachedflowreports
  entity_name: CachedFlowReport
  package: angeal
  group: policy/networking
  description: Post a new cached flow report.
  extends:
  - '@identifiable-stored'
  - '@zoned-monotonic'
  - '@migratable'
  - '@flow'
  validations:
  - $cachedflowreport

# Indexes
indexes:
- - destinationID
- - sourceID

# Attributes
attributes:
  v1:
  - name: isLocalDestinationID
    description: Indicates if the destination endpoint is an enforcer-local processing
      unit.
    type: boolean
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ai

  - name: isLocalSourceID
    description: Indicates if the source endpoint is an enforcer-local processing
      unit.
    type: boolean
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: aj
