# Model
model:
  rest_name: app
  resource_name: apps
  entity_name: App
  package: highwind
  description: App represents an application that can be installed.
  extends:
  - '@described'
  - '@named'

# Attributes
attributes:
  v1:
  - name: beta
    description: Beta indicates if the app is in a beta version.
    type: boolean
    exposed: true
    read_only: true

  - name: categoryID
    description: CategoryID of the app.
    type: string
    exposed: true
    read_only: true

  - name: icon
    description: Icon contains a base64 image for the app.
    type: string
    exposed: true
    read_only: true

  - name: longDescription
    description: LongDescription contains a more detailed description of the app.
    type: string
    exposed: true

  - name: title
    description: Title represents the title of the app.
    type: string
    exposed: true

  - name: versionParameters
    description: VersionParameters contains parameters for each available version.
    type: external
    exposed: true
    subtype: app_versionparameters
