# Model
model:
  rest_name: enforcerreport
  resource_name: enforcerreports
  entity_name: EnforcerReport
  package: zack
  group: core/enforcer
  description: Post a new Defender statistics report.
  extends:
  - '@identifiable-stored'
  validations:
  - $enforcerreport

# Attributes
attributes:
  v1:
  - name: CPULoad
    description: Total CPU utilization of the Defender as a percentage of vCPUs.
    type: float
    exposed: true
    example_value: 10

  - name: enforcerID
    description: ID of the Defender.
    type: string
    exposed: true
    example_value: xxx-xxx-xxx-xxx

  - name: memory
    description: Total resident memory used by the Defender in bytes.
    type: integer
    exposed: true
    example_value: 10000

  - name: name
    description: Name of the Defender.
    type: string
    exposed: true
    required: true
    example_value: aporeto-enforcerd-xxx

  - name: namespace
    description: Namespace of the Defender.
    type: string
    exposed: true
    required: true
    example_value: /my/ns

  - name: processes
    description: Number of active processes of the Defender.
    type: integer
    exposed: true
    example_value: 10

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"
