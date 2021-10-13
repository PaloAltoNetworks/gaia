# Model
model:
  rest_name: pollaccount
  resource_name: pollaccounts
  entity_name: PollAccount
  package: pandemona
  group: pcn/infrastructure
  description: |-
    Initiates a poll for a particular account. Data are stored in the current
    namespace.

# Attributes
attributes:
  v1:
  - name: accountID
    description: The ID of the account.
    type: string
    exposed: true
    required: true
    example_value: 912303033

  - name: authorizationRegion
    description: The region to use for authorization.
    type: string
    exposed: true
    required: true
    example_value: us-east-1

  - name: cloudType
    description: The cloud type for the account.
    type: enum
    exposed: true
    allowed_choices:
    - AWS
    - GCP
    default_value: AWS

  - name: name
    description: The name of the account.
    type: string
    exposed: true
    required: true
    example_value: account-foo

  - name: role
    description: The role that it should use to poll the account.
    type: string
    exposed: true
    required: true
    example_value: ec2-read

  - name: targetRegions
    description: Limit polling to these regions only. If empty, all regions will be polled.
    type: list
    exposed: true
    subtype: string
    example_value:
    - us-east-1
    - us-east-2
