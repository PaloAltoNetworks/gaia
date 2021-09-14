# Model
model:
  rest_name: k8sapiserver
  resource_name: k8sapiservers
  entity_name: K8SAPIServer
  package: squall
  group: core/processingunit
  description: Used to represent an instance of a Kubernetes API server.
  aliases:
  - k8scluster
  - k8sclusters
  get:
    description: Retrieve the Kubernetes api server with the given ID.
  update:
    description: Update the Kubernetes api server with the given ID.
  delete:
    description: Delete the Kubernetes api server with the given ID.
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
  - name: APIVersions
    description: API versions supported by the API server.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: K8SNamespace
    description: Kubernetes namespace.
    type: string
    exposed: true
    stored: true

  - name: externalIP
    description: Cluster external IP address.
    type: string
    exposed: true
    stored: true

  - name: internalIP
    description: Cluster internal IP address.
    type: string
    exposed: true
    stored: true
