# Model
model:
  rest_name: cnssearch
  resource_name: cnssearches
  entity_name: CNSSearch
  package: karl
  group: core/rql
  description: Provide search results for Prisma Cloud's investigate page.

# Attributes
attributes:
  v1:
  - name: ID
    exposed_name: id
    description: ID of the search request.
    type: string
    exposed: true
    omit_empty: true

  - name: data
    description: The payload of the search results.
    type: ref
    exposed: true
    subtype: pcsearchresults
    extensions:
      refMode: pointer

  - name: description
    description: Description of the search.
    type: string
    exposed: true
    omit_empty: true

  - name: endAbsolute
    description: Absolute end time of search, in UNIX time.
    type: integer
    exposed: true
    required: true
    default_value: 0

  - name: limit
    description: The number of items to fetch.
    type: integer
    exposed: true
    default_value: 100
    omit_empty: true

  - name: name
    description: Name of the RQL search request. Should set to be empty.
    type: string
    exposed: true
    omit_empty: true

  - name: pageToken
    description: Represents the token to fetch next page.
    type: string
    exposed: true
    omit_empty: true

  - name: query
    description: The RQL query.
    type: string
    exposed: true
    required: true
    example_value: network dns where id == 1

  - name: saved
    description: Indicates if the search has been saved.
    type: boolean
    exposed: true
    omit_empty: true

  - name: searchType
    description: Type of search request. Should set to be network.
    type: string
    exposed: true
    omit_empty: true

  - name: startAbsolute
    description: Absolute start time of search, in UNIX time.
    type: integer
    exposed: true
    required: true
    default_value: 0
