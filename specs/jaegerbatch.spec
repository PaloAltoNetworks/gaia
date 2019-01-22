# Model
model:
  rest_name: jaegerbatch
  resource_name: jaegerbatchs
  entity_name: Jaegerbatch
  package: meister
  description: |-
    A jaegerbatch is a batch of jaeger spans. This is used by external service to
    post jaeger span in our private jaeger services.
  aliases:
  - sp

# Attributes
attributes:
  v1:
  - name: batch
    description: Represents a jaeger batch.
    type: object
    exposed: true
    stored: true
    creation_only: true
