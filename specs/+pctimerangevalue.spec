# Model
model:
  rest_name: pctimerangevalue
  resource_name: pctimerangevalues
  entity_name: PCTimeRangeValue
  package: karl
  group: core/rql
  description: Represents the time range value of rql search request.

# Attributes
attributes:
  v1:
  - name: amount
    description: The count of time durations.
    type: integer
    exposed: true

  - name: endTime
    description: the end time of the search, in Unix time format.
    type: integer
    exposed: true

  - name: startTime
    description: The start time of the search, in Unix time format.
    type: integer
    exposed: true

  - name: unit
    description: The unit of the time durations.
    type: string
    exposed: true
