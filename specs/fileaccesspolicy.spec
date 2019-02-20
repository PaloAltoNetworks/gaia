# Model
model:
  rest_name: fileaccesspolicy
  resource_name: fileaccesspolicies
  entity_name: FileAccessPolicy
  package: squall
  group: policy/files
  description: |-
    A File Access Policy allows Processing Units to access various folder and files.
    It will use the tags of a File Path to know what is the path of the file or
    folder to allow access to. You can allow the Processing Unit to have any
    combination of read, write or execute.

    When a Processing Unit is Docker container, then it will police the volumes
    mount. executewon''t have any effect.

    File path are not supported yet for standard Linux processes.
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
  - '@disabled'
  - '@identifiable-not-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'
  - '@schedulable'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: allowsExecute
    description: AllowsExecute allows to execute the files.
    type: boolean
    exposed: true
    orderable: true

  - name: allowsRead
    description: AllowsRead allows to read the files.
    type: boolean
    exposed: true
    orderable: true

  - name: allowsWrite
    description: AllowsWrite allows to write the files.
    type: boolean
    exposed: true
    orderable: true

  - name: encryptionEnabled
    description: EncryptionEnabled will enable the automatic encryption.
    type: boolean
    exposed: true
    orderable: true

  - name: logsEnabled
    description: LogsEnabled will enable logging when this policy is used.
    type: boolean
    exposed: true
    orderable: true

  - name: object
    description: Object is the object of the policy.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true

  - name: subject
    description: Subject is the subject of the policy.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
