# Model
model:
  rest_name: export
  resource_name: export
  entity_name: Export
  package: yuna
  group: core
  description: |-
    Allows you to obtain a JSON object containing policies and other objects from a
    given namespace. You can then import this JSON object into a different
    namespace.

# Attributes
attributes:
  v1:
  - name: APIVersion
    description: Version of the Microsegmentation Console API used for the exported data.
    type: integer
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: data
    description: List of all exported data.
    type: external
    exposed: true
    subtype: map[string][]map[string]interface{}
    stored: true
    autogenerated: true

  - name: identities
    description: The list of identities to export.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - externalnetworks
    - networkaccesspolicies

  - name: label
    description: |-
      Allows you to define a unique label for this export. When importing the
      content of the export, this label will be added as a tag that will be used to
      recognize imported object in a later import.
    type: string
    exposed: true
    stored: true
    example_value: my-import-name
