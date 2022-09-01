# Model
model:
  rest_name: firewalllogline
  resource_name: firewallloglines
  entity_name: FirewallLogLine
  package: larl
  group: core/rql
  description: Log line from the results of a firewall log query specified in RQL.
  detached: true

# Attributes
attributes:
  v1:
  - name: fields
    description: The fields and values for the log line.
    type: external
    exposed: true
    subtype: map[string]string

  - name: timestamp
    description: The time for the log message.
    type: time
    exposed: true
