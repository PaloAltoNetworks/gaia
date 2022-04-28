# Model
model:
  rest_name: metricsqueryresult
  resource_name: metricsqueryresult
  entity_name: MetricsQueryResult
  package: jenova
  group: visualization/metricsquery
  description: Represent the result of a metrics query.
  detached: true

# Attributes
attributes:
  v1:
  - name: data
    description: The data of the query.
    type: external
    exposed: true
    subtype: map[string]interface{}

  - name: status
    description: The status of the query.
    type: string
    exposed: true
