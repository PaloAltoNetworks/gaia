# Model
model:
  rest_name: policygraph
  resource_name: policygraphs
  entity_name: PolicyGraph
  package: yeul
  group: visualization/depmaps
  description: |-
    Returns a data structure representing the policy graph of all selected
    processing units and their possible connectivity based on the current policies
    associated with the namespace. Users can define a selector of processing units
    in which they are interested or define the identity tags of a virtual processing
    unit that is not yet activated.
  aliases:
  - polgraph

# Attributes
attributes:
  v1:
  - name: PUIdentity
    description: |-
      The set of tags that a future-activated processing unit will have for which
      the user wants to evaluate policies and understand its connectivity options.
    type: list
    exposed: true
    subtype: string

  - name: dependencyMap
    description: |-
      Contains the output of the policy evaluation. It is the same type of
      dependency map as created by other APIs.
    type: ref
    exposed: true
    subtype: dependencymap
    extensions:
      refMode: pointer

  - name: policyType
    description: |-
      Identifies the type of policy that should be analyzed: `Authorization`(default),
      `Infrastructure`, or `Combined`.
    type: enum
    exposed: true
    allowed_choices:
    - Authorization
    - Infrastructure
    - Combined
    default_value: Authorization

  - name: selectors
    description: |-
      Contains the tag expression that a processing unit must match in order
      to evaluate policy for it.
    type: external
    exposed: true
    subtype: '[][]string'
    example_value:
    - - $identity=processingunit
    validations:
    - $tagsExpression
