# Model
model:
  rest_name: metricsquery
  resource_name: metricsquery
  entity_name: MetricsQuery
  package: jenova
  group: visualization/metrics
  description: |-
    Prometheus compatible endpoint to evaluate a Prometheus query expression at a
    single instant or over a range of time. This can be used to retrieve back
    Aporeto specific metrics for a given namespace. All queries are protected within
    the namespace of the caller.
  aliases:
  - mq

# Attributes
attributes:
  v1:
  - name: query
    description: Prometheus expression query string.
    type: string
    exposed: true
    required: true
    example_value: flows{namespace=~"/mycompany.*"}

  - name: time
    description: Evaluation timestamp <rfc3339 | unix_timestamp>.
    type: string
    exposed: true
    example_value: "2015-07-01T20:11:00.781Z"
