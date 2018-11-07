# Model
model:
  rest_name: appcredential
  resource_name: appcredentials
  entity_name: AppCredential
  package: cactuar
  description: Create a credential for an application.
  aliases:
  - appcred
  - appcreds
  indexes:
  - - namespace
  - - namespace
    - normalizedTags
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@metadatable'
  - '@timeable'
  - '@named'
  - '@disabled'

# Attributes
attributes:
  v1:
  - name: certificate
    description: The string representation of the Certificate used by the application.
    type: string
    exposed: true
    stored: true
    read_only: true

  - name: certificateSN
    description: Link to the certificate created for this application.
    type: string
    stored: true

  - name: credentials
    description: The credential data.
    type: ref
    exposed: true
    subtype: credential
    read_only: true
    autogenerated: true
    orderable: true
    extensions:
      noInit: true
      refMode: pointer

  - name: email
    description: The email address that will receive a copy of the application credentials.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: regenerate
    description: Regenerates the credentials files and certificates.
    type: boolean
    exposed: true

  - name: roles
    description: List of roles to give the credentials.
    type: list
    exposed: true
    subtype: string
    stored: true
    required: true
    default_order: true
    example_value:
    - '@auth:role=enforcer'
    - '@auth:role=kubesquall'
