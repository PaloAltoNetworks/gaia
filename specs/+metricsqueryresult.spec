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
    omit_empty: true

  - name: error
    description: The error message from the query attempt.
    type: string
    exposed: true
    omit_empty: true

  - name: errorType
    description: The type of error that occurred.
    type: string
    exposed: true
    omit_empty: true

  - name: status
    description: The status of the query.
    type: string
    exposed: true
