# Model
model:
  rest_name: firewalllog
  resource_name: firewalllogs
  entity_name: FirewallLog
  package: ngfw ## ???
  group: ngfw/firewall ## ???
  description: |-
    TODO
  
  # Attributes
attributes:
  v1:
  - name: logType
    description: The type of log.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Traffic
    - Threat
    - Decryption
    default_value: Traffic

  - name: destinationType
    description: The destination of type of the log.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - S3
    - Cloudwatch
    - Kinesis
    default_value: S3

  - name: destination
    description: The log destination
    type: string
    exposed: true
    stored: true

