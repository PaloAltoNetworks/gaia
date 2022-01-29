# Model
model:
  rest_name: uiparameter
  resource_name: uiparameters
  entity_name: UIParameter
  package: ignis
  group: core/workflow
  description: Represents a parameter that will be shown in the web interface.
  detached: true
  validations:
  - $uiparameters

# Attributes
attributes:
  v1:
  - name: advanced
    description: A value of `true` designates the parameter as advanced.
    type: boolean
    exposed: true
    stored: true

  - name: allowedChoices
    description: Lists all the choices in case of an enum.
    type: external
    exposed: true
    subtype: map[string]string
    stored: true

  - name: allowedValues
    description: List of values that can be used.
    type: list
    exposed: true
    subtype: object
    stored: true

  - name: defaultValue
    description: Default value of the parameter.
    type: object
    exposed: true
    stored: true

  - name: description
    description: Description of the parameter.
    type: string
    exposed: true
    stored: true

  - name: key
    description: Key identifying the parameter.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: unique_key

  - name: longDescription
    description: Long explanation of the parameter.
    type: string
    exposed: true
    stored: true

  - name: name
    description: Name of the parameter.
    type: string
    exposed: true
    stored: true

  - name: optional
    description: A value of `true` designates the parameter as optional.
    type: boolean
    exposed: true
    stored: true

  - name: subtype
    description: The subtype of a list parameter.
    type: string
    exposed: true
    stored: true

  - name: type
    description: The datatype of the parameter.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Boolean
    - Checkbox
    - CVSSThreshold
    - DangerMessage
    - Duration
    - Enum
    - Endpoint
    - FileDrop
    - Float
    - FloatSlice
    - InfoMessage
    - Integer
    - IntegerSlice
    - JSON
    - List
    - Message
    - Namespace
    - Password
    - String
    - StringSlice
    - Switch
    - TagsExpression
    - Title
    - WarningMessage
    - AutomationAction
    example_value: String

  - name: validationFunction
    description: A function that validates the parameter.
    type: string
    exposed: true
    stored: true

  - name: value
    description: Value of the parameter.
    type: object
    exposed: true
    stored: true
    deprecated: true

  - name: visibilityCondition
    description: |-
      A logical expression consisting of one or more
      [UIParameterVisibility](#uiparametervisibility)
      conditions linked together using AND or OR operators. If the expression
      evaluates to true
      the parameter is displayed to the user.
    type: external
    exposed: true
    subtype: uiparametersexpression
    stored: true

  - name: width
    description: Width of the parameter.
    type: string
    exposed: true
    stored: true
    default_value: 100%
