# Model
model:
  rest_name: installedapp
  resource_name: installedapps
  entity_name: InstalledApp
  package: highwind
  group: integration/app
  description: Represents an installed application.
  aliases:
  - iapps
  - iapp
  get:
    description: Retrieves the application with the given ID.
  update:
    description: Updates the application with the given ID.
  delete:
    description: Deletes the application with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@identifiable-stored'
  - '@named'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: additionalConfiguration
    description: Additional configuration of the app is needed by the app itself.
    type: boolean
    exposed: true
    stored: true

  - name: appIdentifier
    description: AppIdentifier retains the identifier for the app.
    type: string
    stored: true

  - name: categoryID
    description: The category ID of the application.
    type: string
    exposed: true
    stored: true
    read_only: true
    orderable: true

  - name: checkPublicEndpoint
    description: |-
      If true, will look for the public endpoints and store them as annotations in the
      installed app.
    type: boolean
    exposed: true
    stored: true

  - name: currentVersion
    description: Version of the installed application.
    type: string
    exposed: true
    stored: true

  - name: deploymentCount
    description: DeploymentCount represents the number of expected deployment for
      this app.
    type: integer
    stored: true
    read_only: true

  - name: externalWindowButton
    description: Adds a button in the UI.
    type: external
    exposed: true
    subtype: map[string]string
    stored: true

  - name: parameters
    description: Contains the computed parameters to start the application.
    type: external
    exposed: true
    subtype: map[string]interface{}
    stored: true

  - name: status
    description: Status of the application.
    type: enum
    exposed: true
    stored: true
    read_only: true
    allowed_choices:
    - Unknown
    - Deploying
    - Initializing
    - Running
    - Undeploying
    - Error
    default_value: Unknown
    orderable: true

  - name: statusMessage
    description: Reason for the status of the application.
    type: string
    exposed: true
    stored: true
    read_only: true

# Relations
relations:
- rest_name: log
  get:
    description: Returns the logs for an application.
