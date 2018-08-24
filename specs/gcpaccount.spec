# Model
model:
  rest_name: gcpaccount
  resource_name: gcpaccounts
  entity_name: GCPAccount
  package: vince
  description: |-
    Allows to bind an GCP account to your Aporeto account to allow auto registration
    of enforcers running on Google Cloud. We ask for the credential for the sole
    purpose of validating you are the owner the project ID. We don't store this
    credentials and we do not send any api call to your GCP project.
  aliases:
  - gcp
  - gcpaccs
  - gcpacc
  get:
    description: Retrieves the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@identifiable-pk-stored'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: JSONCredentialFile
    description: |-
      Credential file to use to validate your account. You can create a temporary
      service account with read only permissions. We do not store this file.
    type: string
    exposed: true
    required: true
    creation_only: true
    example_value: '{<credential file content>}'

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

  - name: projectID
    description: projectID contains your verified accound id.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    filterable: true
    orderable: true
