# Model
model:
  rest_name: k8sresource
  resource_name: k8sresources
  entity_name: K8sResource
  package: pandemona
  group: pcn/infrastructure
  description: Represents a K8s resource such as a service.
  private: true
  extends:
  - '@identifiable-stored'
  - '@migratable'
  - '@namespaced'
  - '@zoned'
  - '@timeable'
  - '@k8sresourceattrs'

# Indexes
indexes:
- - namespace
  - k8sID
- - namespace
  - clusterID
  - k8sNamespace
  - kind
  - labels
- - namespace
  - uid
- - prismaID
- - prismaID
  - kind
  - cloudProvider
  - accountID
  - clusterID
  - k8sNamespace
- - prismaID
  - kind
  - cloudProvider
  - clusterID
  - k8sNamespace
- - prismaID
  - kind
  - cloudProvider
  - accountID
  - k8sNamespace
- - prismaID
  - kind
  - cloudProvider
  - clusterID
  - k8sNamespace
  - name
- - prismaID
  - kind
  - cloudProvider
  - clusterID
  - name
- - prismaID
  - kind
  - cloudProvider
  - name
- - prismaID
  - kind
  - cloudProvider
  - uid
- - prismaID
  - kind
  - cloudProvider
