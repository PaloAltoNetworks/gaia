# Model
model:
  rest_name: ngfwreport
  resource_name: ngfwreports
  entity_name: NGFWReport
  package: tbd
  group: tbd
  description: Defines a Next-Generation Firewall (NGFW) report.
  extends:
  - '@identifiable-stored'
  - '@namespaced'
  - '@zoned'
  - '@migratable'

# Indexes
indexes:
- - namespace
  - timestamp

# Attributes
attributes:
  v1:
  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
    stored: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"

  - name: instanceID
    description: Cloud NGFW instance for generated report.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 12345678

  - name: reportType
    description: Type of Cloud NGFW report stored.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - NetworkingReport
    - SecurityReport
    default_value:
    - NetworkingReport

  - name: topSources
    description: Top IP sources for generated networking report.
    type: external
    subtype: []map[string]int
    exposed: true
    stored: true
    omit_empty: true
    example_value: [{"127.0.0.1": 200}, {"0.0.0.0": 100}]

  - name: topDestinations
    description: Top IP destinations for generated networking report.
    type: external
    subtype: []map[string]int
    exposed: true
    stored: true
    omit_empty: true
    example_value: [{"127.0.0.1": 200}, {"0.0.0.0": 100}]

  - name: topApplications
    description: Top applications used for generated networking report.
    type: external
    subtype: []map[string]int
    exposed: true
    stored: true
    omit_empty: true
    example_value: [{"app1": 200}, {"app2": 100}]

  - name: totalIngress
    description: Total ingress traffic in bytes for generated networking report.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    example_value: 1000000

  - name: totalEgress
    description: Total egress traffic in bytes for generated networking report.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    example_value: 1000000

  - name: topAttackedApplications
    description: Top applications attacked used for generated security report.
    type: external
    subtype: []map[string]int
    exposed: true
    stored: true
    omit_empty: true
    example_value: [{"app1": 200}, {"app2": 100}]

  - name: topSourceAttackers
    description: Top source IPs in attacks used for generated security report.
    type: external
    subtype: []map[string]int
    exposed: true
    stored: true
    omit_empty: true
    example_value: [{"127.0.0.1": 200}, {"0.0.0.0": 100}]

  - name: topDestinationVictims
    description: Top destination IPs for victims used for generated security report.
    type: external
    subtype: []map[string]int
    exposed: true
    stored: true
    omit_empty: true
    example_value: [{"127.0.0.1": 200}, {"0.0.0.0": 100}]

  - name: topAttackCategories
    description: Top attack categories used for generated security report.
    type: external
    subtype: []map[string]int
    exposed: true
    stored: true
    omit_empty: true
    example_value: [{"Phishing": 200}, {"Malware": 100}]

  - name: topDetectedAttacks
    description: Top detected attacks used for generated security report.
    type: external
    subtype: []map[string]int
    exposed: true
    stored: true
    omit_empty: true
    example_value: [{"Log4j": 200}, {"TrojanX": 100}]
