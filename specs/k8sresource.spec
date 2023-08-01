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
  - clusterID
- - namespace
  - clusterID
  - k8sNamespace
  - kind
- - namespace
  - clusterID
  - k8sNamespace
  - name
- - namespace
  - uid
