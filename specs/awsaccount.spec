# Model
model:
  rest_name: awsaccount
  resource_name: awsaccounts
  entity_name: AWSAccount
  package: vince
  group: core/authentication
  description: |-
    Allows to bind an AWS account to your Aporeto account to allow auto registration
    of enforcers running on EC2.
  aliases:
  - aws
  - awsaccs
  - awsacc
  get:
    description: Retrieves the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@identifiable-stored'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: accessKeyID
    description: |-
      AccessKeyID contains the aws access key ID. This is used to retrieve your
      account id, and it is not stored.
    type: string
    exposed: true
    required: true
    creation_only: true
    example_value: AKIAIOSFODNN7EXAMPLE

  - name: accessToken
    description: |-
      AccessToken contains your aws access token. It is used to retrieve your account
      id, and it not stored.
    type: string
    exposed: true
    creation_only: true

  - name: accountID
    description: accountID contains your verified accound id.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    filterable: true
    orderable: true

  - name: parentID
    description: ParentID contains the parent Vince account ID.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    filterable: true

  - name: parentName
    description: ParentName contains the name of the Vince parent Account.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    filterable: true

  - name: region
    description: Region contains your the region where your AWS account is located.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: us-west-2
    filterable: true
    orderable: true

  - name: secretAccessKey
    description: |-
      secretAccessKey contains the secret key. It is used to retrieve your account id,
      and it is not stored.
    type: string
    exposed: true
    required: true
    creation_only: true
    example_value: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
