# Model
model:
  rest_name: accessreport
  resource_name: accessreports
  entity_name: AccessReport
  package: zack
  group: policy/access
  description: Represents any access made by the user.
  extends:
  - '@identifiable-stored'
  - '@zoned-monotonic'
  - '@migratable'

# Ordering
default_order:
- :no-inherit
- timestamp

# Indexes
indexes:
- - namespace
  - timestamp

# Attributes
attributes:
  v1:
  - name: action
    description: Action applied to the access.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Accept
    - Reject
    example_value: Accept
    omit_empty: true
    extensions:
      bson_name: a

  - name: claimHash
    description: Hash of the claims used to communicate.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: b

  - name: enforcerID
    description: Identifier of the enforcer.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: c

  - name: enforcerNamespace
    description: Namespace of the enforcer.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/namespace
    omit_empty: true
    extensions:
      bson_name: d

  - name: namespace
    description: |-
      The namespace that this access report belongs to. This field is not exposed.
      Therefore, any value set by the user will be ignored.
    type: string
    stored: true
    example_value: /my/namespace
    filterable: true
    omit_empty: true
    extensions:
      bson_name: k

  - name: processingUnitID
    description: ID of the processing unit of the report.
    type: string
    exposed: true
    stored: true
    example_value: xxx-xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: e

  - name: processingUnitName
    description: Name of the processing unit of the report.
    type: string
    exposed: true
    stored: true
    example_value: pu1
    omit_empty: true
    extensions:
      bson_name: f

  - name: processingUnitNamespace
    description: Namespace of the processing unit of the report.
    type: string
    exposed: true
    stored: true
    example_value: /my/ns
    omit_empty: true
    extensions:
      bson_name: g

  - name: reason
    description: |-
      This field is only set if `action` is set to `Reject`. It specifies the reason
      for the rejection.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: h

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
    stored: true
    orderable: true
    omit_empty: true
    extensions:
      bson_name: i

  - name: type
    description: Type of the report.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - SSHLogin
    - SSHLogout
    - SudoEnter
    - SudoExit
    example_value: SSHLogin
    omit_empty: true
    extensions:
      bson_name: j
