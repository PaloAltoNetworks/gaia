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
  - '@identifiable-stored'

# Attributes
attributes:
  v1:
  - name: annotations
    description: Stores additional information about an entity.
    type: external
    exposed: true
    subtype: map[string][]string
    stored: true

  - name: namespace
    description: Namespace associated with the cluster.
    type: string
    exposed: true
    stored: true
    read_only: true

  - name: apiversions
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
