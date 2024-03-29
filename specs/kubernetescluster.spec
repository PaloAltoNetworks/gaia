# Model
model:
  rest_name: kubernetescluster
  resource_name: kubernetesclusters
  entity_name: KubernetesCluster
  package: squall
  group: core/processingunit
  description: Used to represent an instance of a Kubernetes API server.
  aliases:
  - k8scluster
  - k8sclusters
  get:
    description: Retrieve the Kubernetes cluster with the given ID.
  update:
    description: Update the Kubernetes cluster with the given ID.
  delete:
    description: Delete the Kubernetes cluster with the given ID.
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
  - name: APIServerServiceFQDNs
    description: |-
      Contains the FQDNs used by the API server. They will be used to populate the
      Certificate DNS SANs field.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: APIServerServiceIPs
    description: |-
      Contains the IPs used by the API server. They will be used to populate the
      Certificate IP SANs field.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: APIServerServiceName
    description: |-
      Kubernetes service name in the format <service name>.<service name
      namespace>.svc will be set in the certificate CommonName field.
    type: string
    exposed: true
    stored: true
    validations:
    - $apiserverservicename

  - name: APIServerVersions
    description: API versions supported by the API server.
    type: list
    exposed: true
    subtype: string
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
