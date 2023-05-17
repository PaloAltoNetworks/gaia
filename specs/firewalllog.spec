# Model
model:
  rest_name: firewalllog
  resource_name: firewalllogs
  entity_name: FirewallLog
  package: pcfw
  group: core/log
  description: Represents a log line in a log query result.

# Attributes
attributes:
  v1:
  - name: XFFIP
    description: |-
      The IP address of the user who requested the web page or the IP address of the
      next to last device that the request traversed. If the request goes through one
      or more proxies, load balancers, or other upstream devices, the firewall
      displays the IP address of the most recent device.
    type: string
    exposed: true
    read_only: true

  - name: action
    description: Action taken for the session.
    type: string
    exposed: true
    read_only: true

  - name: app
    description: Application associated with the session.
    type: string
    exposed: true
    read_only: true

  - name: appCategory
    description: Application category associated with the session.
    type: string
    exposed: true
    read_only: true

  - name: appSubcategory
    description: Application subcategory associated with the session.
    type: string
    exposed: true
    read_only: true

  - name: bytesReceived
    description: Number of bytes in the server-to-client direction of the session.
    type: string
    exposed: true
    read_only: true

  - name: bytesSent
    description: Number of bytes in the client-to-server direction of the session.
    type: string
    exposed: true
    read_only: true

  - name: chainStatus
    description: Chain status of the session.
    type: string
    exposed: true
    read_only: true

  - name: cloudAccount
    description: Cloud account associated with the session.
    type: string
    exposed: true
    read_only: true

  - name: cloudRegion
    description: Cloud provider region associated with the session.
    type: string
    exposed: true
    read_only: true

  - name: cloudType
    description: Cloud provider associated with the session.
    type: string
    exposed: true
    read_only: true

  - name: destinationCountry
    description: |-
      Destination country or Internal region for private addresses. Maximum length is
      32 bytes.
    type: string
    exposed: true
    read_only: true

  - name: destinationIP
    description: Original session destination IP address.
    type: string
    exposed: true
    read_only: true

  - name: destinationPort
    description: Destination port utilized by the session.
    type: integer
    exposed: true
    read_only: true

  - name: firewallName
    description: Name of firewall that generated the log.
    type: string
    exposed: true
    read_only: true

  - name: operationMode
    description: Operation mode of firewall that generated the log.
    type: string
    exposed: true
    read_only: true

  - name: packetsReceived
    description: Number of server-to-client packets for the session.
    type: string
    exposed: true
    read_only: true

  - name: packetsSent
    description: Number of client-to-server packets for the session.
    type: string
    exposed: true
    read_only: true

  - name: protocol
    description: IP protocol associated with the session.
    type: string
    exposed: true
    read_only: true

  - name: proxyType
    description: Type of decryption proxy associated with the session.
    type: string
    exposed: true
    read_only: true

  - name: rule
    description: Name of the rule that the session matched.
    type: string
    exposed: true
    read_only: true

  - name: sessionEndReason
    description: |-
      The reason a session terminated. If the termination had multiple causes, this
      field displays only the highest priority reason.
    type: string
    exposed: true
    read_only: true

  - name: sni
    description: Server Name Indication (SNI) value for the session.
    type: string
    exposed: true
    read_only: true

  - name: sourceCountry
    description: |-
      Source country or Internal region for private addresses; maximum length is 32
      bytes.
    type: string
    exposed: true
    read_only: true

  - name: sourceIP
    description: Original session source IP address.
    type: string
    exposed: true
    read_only: true

  - name: sourcePort
    description: Source port utilized by the session.
    type: integer
    exposed: true
    read_only: true

  - name: threatCategory
    description: |-
      Describes threat categories used to classify different types of threat
      signatures.
    type: string
    exposed: true
    read_only: true

  - name: threatDirection
    description: Indicates the direction of the attack, client-to-server or server-to-client.
    type: string
    exposed: true
    read_only: true

  - name: threatName
    description: Name of the threat signature that triggered the session.
    type: string
    exposed: true
    read_only: true

  - name: threatSeverity
    description: |-
      Severity associated with the threat; values are informational, low, medium,
      high, critical.
    type: string
    exposed: true
    read_only: true

  - name: tlsVersion
    description: TLS version associated with the session.
    type: string
    exposed: true
    read_only: true

  - name: urlCategory
    description: URL category associated with the session (if applicable).
    type: string
    exposed: true
    read_only: true

  - name: vpcEndpoint
    description: Service endpoint associated with the session.
    type: string
    exposed: true
    read_only: true
