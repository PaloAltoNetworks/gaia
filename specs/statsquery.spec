# Model
model:
  rest_name: statsquery
  resource_name: statsqueries
  entity_name: StatsQuery
  package: jenova
  group: visualization/statsquery
  description: |-
    Retrieves time-series data stored by the Microsegmentation Console.
    Allows different types of queries that are all protected within
    the namespace of the user.
  aliases:
  - sq

# Attributes
attributes:
  v1:
  - name: descending
    description: If set, the results will be order by time from the most recent to
      the oldest.
    type: boolean
    exposed: true

  - name: fields
    description: |-
      List of fields to extract. If you don't pass anything, all available fields will
      be returned. It is also possible to use a function like `sum(value)`.
    type: list
    exposed: true
    subtype: string

  - name: filter
    description: Apply a filter to the query.
    type: string
    exposed: true

  - name: groups
    description: |-
      Group results by the provided values. Note that not all fields can be used to
      group the results.
    type: list
    exposed: true
    subtype: string

  - name: limit
    description: Limits the number of results. `-1` means no limit.
    type: integer
    exposed: true
    default_value: -1

  - name: measurement
    description: Name of the measurement.
    type: enum
    exposed: true
    allowed_choices:
    - Flows
    - Audit
    - Enforcers
    - Files
    - EventLogs
    - Packets
    - EnforcerTraces
    - Counters
    - Accesses
    - DNSLookups
    - PingReports
    - ConnectionException
    default_value: Flows

  - name: offset
    description: Offsets the results. -1 means no offset.
    type: integer
    exposed: true
    default_value: -1

  - name: results
    description: Contains the result of the query.
    type: refList
    exposed: true
    subtype: timeseriesqueryresults
    read_only: true
    autogenerated: true
    extensions:
      refMode: pointer
