# Model
model:
  rest_name: namespacetype
  resource_name: namespacetypes
  entity_name: NamespaceType
  package: squall
  group: core/namespace
  description: Returns the type of the specified namespace.

# Attributes
attributes:
  v1:
  - name: type
    description: |-
      The type defines the purpose of the namespace:
      - `Default`: A universal namespace that is capable of all actions and views.
      - `Tenant`: A namespace that houses a tenant (e.g. ACME).
      - `CloudAccount`: A child namespace of a tenant that houses a cloud provider
      account (e.g. aws-123, gcp-54).
      - `Group`: A child namespace of a cloud account that houses a managed group
      (e.g. marketing, app-234).
    type: enum
    exposed: true
    read_only: true
    autogenerated: true
    transient: true
    allowed_choices:
    - Default
    - Tenant
    - CloudAccount
    - Group
    default_value: Default
