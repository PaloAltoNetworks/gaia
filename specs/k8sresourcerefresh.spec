# Model
model:
  rest_name: k8sresourcerefresh
  resource_name: k8sresourcerefreshes
  entity_name: K8sResourceRefresh
  package: pandemona
  group: pcn/infrastructure
  description: |-
    When requested, k8s resources will start the process of pulling down all
    Kubernetes resources from PrismaCloud.

# Attributes
attributes:
  v1:
  - name: background
    description: Set to `true` to make the request run in the background.
    type: boolean
    exposed: true
