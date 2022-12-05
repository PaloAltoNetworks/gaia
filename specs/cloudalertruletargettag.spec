# Model
model:
  rest_name: cloudalertruletargettag
  resource_name: cloudalertruletargettags
  entity_name: CloudAlertRuleTargetTag
  package: vargid
  group: pcn/infrastructure
  description: Represents an alert rule target resource tag.
  detached: true

# Attributes
attributes:
  v1:
  - name: key
    description: Alert Rule target tag key.
    type: string
    exposed: true
    subtype: string
    stored: true

  - name: values
    description: List of Alert Rule target tag values.
    type: list
    exposed: true
    subtype: string
    stored: true
