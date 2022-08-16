# Model
model:
  rest_name: awsfirewallattachment
  resource_name: awsfirewallattachments
  entity_name: AWSFirewallAttachment
  package: ngfw ## ???
  group: ngfw/firewall ## ???
  description: |-
    TODO
  
  # Attributes
attributes:
  v1:
  - name: vpcID
    description: The AWS vpcID.
    type: string
    exposed: true
    stored: true
    
  - name: endpoints
    description: The list of AWS endpoints.
    type: list
    subtype: awsfirewallendpoint
    extensions:
      refMode: pointer
    exposed: true
    stored: true


