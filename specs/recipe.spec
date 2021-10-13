# Model
model:
  rest_name: recipe
  resource_name: recipes
  entity_name: Recipe
  package: ignis
  group: core/workflow
  description: Defines a list of steps that make up a workflow.
  aliases:
  - rcp
  get:
    description: Retrieves the recipe with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the recipe with the given ID.
  delete:
    description: Deletes the recipe with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@named'
  - '@metadatable'
  - '@propagated'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: deploymentMode
    description: |-
      Defines the deployment mode of the recipe. If `Unrestricted`, the recipe can be
      deployed multiple times in the current namespace and below. If `Unique`,only one
      deployment is allowed in the current namespace and its child namespaces. If
      `NamespaceUnique`, only one deployment is allowed in the current namespace.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Unrestricted
    - Unique
    - NamespaceUnique
    default_value: Unrestricted

  - name: icon
    description: Contains a base64-encoded image for the recipe.
    type: string
    exposed: true
    stored: true

  - name: key
    description: The unique key of the recipe.
    type: string
    exposed: true
    stored: true
    read_only: true

  - name: label
    description: Defines the recipe.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    default_value: magicpanda

  - name: longDescription
    description: Provides a long description of the recipe.
    type: string
    exposed: true
    stored: true

  - name: options
    description: Options of the recipe.
    type: ref
    exposed: true
    subtype: recipeoptions
    stored: true
    extensions:
      refMode: pointer

  - name: steps
    description: Contains all the steps with parameters to follow for the recipe.
    type: refList
    exposed: true
    subtype: uistep
    stored: true
    extensions:
      refMode: pointer

  - name: successfullMessage
    description: A string message presented upon success (optional).
    type: string
    exposed: true
    stored: true

  - name: targetIdentities
    description: Contains the list of identities the recipes will try to create.
    type: list
    exposed: true
    subtype: string
    stored: true
    required: true
    example_value:
    - processingunit
    - enforcer

  - name: template
    description: Template of the recipe to import.
    type: string
    exposed: true
    stored: true

  - name: templateHash
    description: A hash of the template.
    type: string
    exposed: true
    stored: true
    read_only: true

# Relations
relations:
- rest_name: importreference
  get:
    description: Returns the list of import references that depend on a recipe.
  create:
    description: Create an import request for the given recipe.
