# Model
model:
  rest_name: k8sasset
  resource_name: k8sassets
  entity_name: K8sAsset
  package: pandemona
  group: pcn/infrastructure
  description: Represents a read-only K8s resource such as a service.
  extends:
  - '@identifiable-stored'
  - '@migratable'
  - '@namespaced'
  - '@zoned'
  - '@timeable'
  - '@k8sresourceattrs'
