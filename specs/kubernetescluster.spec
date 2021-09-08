# Model
model:
  rest_name: kubernetescluster
  resource_name: kubernetesclusters
  entity_name: KubernetesCluster
  package: squall
  group: core/processingunit
  description: |-
    Used to represent an instance of a Kubernetes API server.
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: K8SNamespace
    description: Kubernetes namespace.
    type: string
    exposed: true
    stored: true

  - name: APIVersions
    description: API versions supported by the API server.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: internalIP
    description: Cluster internal IP address.
    type: string
    exposed: true
    stored: true

  - name: externalIP
    description: Cluster external IP address.
    type: string
    exposed: true
    stored: true
