# Model
model:
  rest_name: firewalllog
  resource_name: firewalllogs
  entity_name: FirewallLog
  package: larl
  group: core/rql
  description: Results of a firewall log query specified in RQL.
  detached: true

# Attributes
attributes:
  v1:
  - name: logLines
    description: Firewall log lines.
    type: refList
    exposed: true
    subtype: firewalllogline
    extensions:
      refMode: pointer
