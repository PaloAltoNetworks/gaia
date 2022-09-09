# Model
model:
  rest_name: ngfwriskassessmentdashboard
  resource_name: ngfwriskassessmentdashboards
  entity_name: NGFWRiskAssessmentDashboard
  package: placeholder
  group: placeholder
  description: Defines Next-Generation Firewall (NGFW) risk asssessment dashboard
    metrics.
  get:
    description: Retrieves NGFW risk assessment dashboard metrics for the specified
      cloud account and time filter.
    global_parameters:
    - $propagatable
  extends:
  - '@identifiable-stored'
  - '@namespaced'
  - '@zoned'
  - '@migratable'

# Indexes
indexes:
- - namespace

# Attributes
attributes:
  v1:
  - name: assetsAtRiskTrendline
    description: "Array of number of workloads contacting suspicious domains, involved
      in\nanomalous activities, \nor internet facing with high / critical vulnerabilities
      detected in provided\nrange."
    type: list
    exposed: true
    subtype: integer
    example_value:
    - 1
    - 2
    - 1
    - 2

  - name: cloudAccountID
    description: The specified cloud account to provide risk asssessment metrics for.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 12345678

  - name: customDateFilterEnd
    description: The end of the custom date range when CustomDate is used for dateRangeFilter.
    type: time
    example_value: "2018-06-14T00:00:00.000000000Z"

  - name: customDateFilterStart
    description: The start of the custom date range when CustomDate is used for dateRangeFilter.
    type: time
    example_value: "2018-06-14T00:00:00.000000000Z"

  - name: dateRangeFilter
    description: The specified date range filter type.
    type: enum
    required: true
    allowed_choices:
    - Last24Hours
    - LastWeek
    - LastMonth
    - Last90Days
    - CustomDate
    default_value: Last24Hours

  - name: recommendedFirewallDeploymentsTrendline
    description: Array of number of recommended firewall deployments in provided range.
    type: list
    exposed: true
    subtype: integer
    example_value:
    - 1
    - 2
    - 1
    - 2

  - name: riskScoreConfidenceLevel
    description: Latest risk score confidence level in provided range.
    type: enum
    exposed: true
    allowed_choices:
    - LowConfidence
    - MediumConfidence
    - HighConfidence
    - VeryHighConfidence
    default_value: LowConfidence

  - name: riskScoreTrendLine
    description: Array of risk score values over provided range.
    type: list
    exposed: true
    subtype: integer
    example_value:
    - 2
    - 3
    - 2
    - 8

  - name: threatsTrendline
    description: Array of number of CVEs on internet exposed assets detected in provided
      range.
    type: list
    exposed: true
    subtype: integer
    example_value:
    - 1
    - 2
    - 1
    - 2


  # WIP
  
  # - name: topMostAccessedDomains
  #   description: The most accessed domains and the assets that contacted them in provided range.
  #   type: external
  #   subtype: map[string]map[int][]string
  #   exposed: true
  #   example_value:
  #     domain.com:
  #       123:
  #         - asset1
  #         - asset2
  #     domain2.com:
  #       99:
  #         - asset2
  #         - asset3

  # - name: topSuspiciousDomains
  #   description: The most accessed suspicious domains and the assets that contacted them in provided range.
  #   type: external
  #   subtype: map[string]map[int][]string
  #   exposed: true
  #   example_value:
  #     domain.com:
  #       123:
  #         - asset1
  #         - asset2
  #     domain2.com:
  #       99:
  #         - asset2
  #         - asset3

  # - name: topWorkloadsContactingSuspiciousDomains
  #   description: |-
  #     The workloads with greatest number of connections to suspicious domains in provided range.

  # - name: topWorkloadsInvolvedWithAnomalousActivities

  # - name: topVPCsAtRisk

  # - name: cloudNetworkGraph
  # description: A visual graph of how threats can reach vulnerabilities on client's network.
  # type: TBD
  # exposed: true
  # example_value: