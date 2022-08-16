# Model
model:
  rest_name: firewallattachment
  resource_name: firewallattachments
  entity_name: FirewallAttachment
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
    subtype: firewallendpoint
    extensions:
      refMode: pointer
    exposed: true
    stored: true


