# Model
model:
  rest_name: suggestedpolicy
  resource_name: suggestedpolicies
  entity_name: SuggestedPolicy
  package: jenova
  group: visualization/depmaps
  description: Allows you to obtain network ruleset policy suggestions.
  aliases:
  - sugpol
  - sugpols
  - sugg
  - suggs

# Attributes
attributes:
  v1:
  - name: policy
    description: The suggested network policy.
    type: ref
    exposed: true
    subtype: networkrulesetpolicy
    extensions:
      refMode: pointer
