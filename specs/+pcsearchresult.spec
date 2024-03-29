# Model
model:
  rest_name: pcsearchresult
  resource_name: pcsearchresults
  entity_name: PCSearchResult
  package: karl
  group: core/rql
  description: Represents the result data of RQL search.

# Attributes
attributes:
  v1:
  - name: internalEdges
    description: The edges of the map connecting internal endpoints.
    type: refMap
    exposed: true
    subtype: cloudgraphedge
    read_only: true
    autogenerated: true
    extensions:
      refMode: pointer

  - name: items
    description: The payload of the search result.
    type: ref
    exposed: true
    subtype: reportsquery
    read_only: true
    extensions:
      refMode: pointer

  - name: nextPageToken
    description: The pagination token for next page.
    type: string
    exposed: true
    read_only: true

  - name: nodes
    description: Refers to the nodes of the map.
    type: refMap
    exposed: true
    subtype: cloudgraphnode
    read_only: true
    autogenerated: true
    extensions:
      refMode: pointer

  - name: paths
    description: |-
      The set of paths from the given source to the given destination. Only provided
      when the query type is NetworkPath.
    type: external
    exposed: true
    subtype: '[][]cloudpathelement'
    read_only: true
    autogenerated: true
    extensions:
      refMode: pointer

  - name: publicEdges
    description: The edges of the map connecting public endpoints.
    type: refMap
    exposed: true
    subtype: cloudgraphedge
    read_only: true
    autogenerated: true
    extensions:
      refMode: pointer

  - name: sourceDestinationMap
    description: |-
      The set of destinations that have been discovered based on the query and their
      associated verdicts.
    type: external
    exposed: true
    subtype: map[string]map[string]cloudnetworkquerydestination
    read_only: true
    autogenerated: true
    extensions:
      refMode: pointer

  - name: totalRows
    description: The total number of result items.
    type: integer
    exposed: true
    read_only: true
