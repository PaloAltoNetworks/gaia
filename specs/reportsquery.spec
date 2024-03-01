# Model
model:
  rest_name: reportsquery
  resource_name: reportsqueries
  entity_name: ReportsQuery
  package: jenova
  group: visualization/reportsquery
  description: |-
    Supports querying Aporeto reports. All queries are protected within the
    namespace of the user.
  aliases:
  - rq

# Attributes
attributes:
  v1:
  - name: DNSLookupReports
    description: List of DNSLookupReports.
    type: refList
    exposed: true
    subtype: dnslookupreport
    omit_empty: true

  - name: accessReports
    description: List of AccessReports.
    type: refList
    exposed: true
    subtype: accessreport
    omit_empty: true

  - name: connectionExceptionReports
    description: List of ConnectionExceptionReports.
    type: refList
    exposed: true
    subtype: connectionexceptionreport
    omit_empty: true

  - name: counterReports
    description: List of CounterReports.
    type: refList
    exposed: true
    subtype: counterreport
    omit_empty: true

  - name: enforcerReports
    description: List of EnforcerReports.
    type: refList
    exposed: true
    subtype: enforcerreport
    omit_empty: true

  - name: eventLogs
    description: List of EventLogs.
    type: refList
    exposed: true
    subtype: eventlog
    omit_empty: true

  - name: firewallLogs
    description: List of FirewallLogs.
    type: refList
    exposed: true
    subtype: firewalllog
    omit_empty: true

  - name: flowReports
    description: List of FlowReports.
    type: refList
    exposed: true
    subtype: flowreport
    omit_empty: true

  - name: packetReports
    description: List of PacketReports.
    type: refList
    exposed: true
    subtype: packetreport
    omit_empty: true

  - name: report
    description: Name of the report type to query.
    type: enum
    exposed: true
    allowed_choices:
    - Accesses
    - Flows
    - Enforcers
    - EventLogs
    - Packets
    - Counters
    - DNSLookups
    - ConnectionExceptions
    - FirewallLogs
    default_value: Flows
