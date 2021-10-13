# Model
model:
  rest_name: ldapprovider
  resource_name: ldapproviders
  entity_name: LDAPProvider
  package: cactuar
  group: core/authentication
  description: |-
    Allows you to declare a generic LDAP provider that can be used in exchange
    for a Midgard token.
  get:
    description: Retrieves the provider with the given ID.
  update:
    description: Updates the provider with the given ID.
  delete:
    description: Deletes the provider with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@timeable'
  - '@named'

# Attributes
attributes:
  v1:
  - name: address
    description: |-
      Contains the fully qualified domain name (FQDN) or IP address of the private
      LDAP server.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: ldap.company.com
    filterable: true
    orderable: true

  - name: baseDN
    description: |-
      Contains the base distinguished name (DN) to use for LDAP queries. Example:
      `dc=example,dc=com`.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: dc=universe,dc=io
    filterable: true
    orderable: true

  - name: bindDN
    description: |-
      Contains the DN to use to bind to the LDAP server. Example:
      `cn=admin,dc=example,dc=com`.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: cn=readonly,dc=universe,dc=io
    filterable: true
    orderable: true

  - name: bindPassword
    description: |-
      Contains the password to be used with the `bindDN` to authenticate to the LDAP
      server.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: s3cr3t
    orderable: true

  - name: bindSearchFilter
    description: |-
      The filter to use to locate the relevant user accounts. For Windows-based
      systems, the value may
      be `sAMAccountName={USERNAME}`. For Linux and other systems, the value may be
      `uid={USERNAME}`.
    type: string
    exposed: true
    stored: true
    default_value: uid={USERNAME}
    orderable: true

  - name: certificateAuthority
    description: |-
      Can be left empty if the LDAP server's certificate is signed by a public,
      trusted certificate
      authority. Otherwise, include the public key of the certificate authority that
      signed the
      LDAP server's certificate.
    type: string
    exposed: true
    stored: true
    example_value: |-
      -----BEGIN CERTIFICATE-----
      MIIBPzCB5qADAgECAhEAwbx3c+QW24ePXyD94geytzAKBggqhkjOPQQDAjAPMQ0w
      CwYDVQQDEwR0b3RvMB4XDTE5MDIyMjIzNDA1MFoXDTI4MTIzMTIzNDA1MFowDzEN
      MAsGA1UEAxMEdG90bzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJi6CwRDeKks
      Xb3pDEslmFGR7k9Aeh5RK+XmdqKKPGb3NQWEFPGolnqOR34iVuf7KSxTuzaaVWfu
      XEa94faUQEqjIzAhMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8EBTADAQH/MAoG
      CCqGSM49BAMCA0gAMEUCIQD+nL9RF9EvQXHyYuJ31Lz9yWd9hsK91stnpAs890gS
      /AIgQIKjBBpiyQNZZWso5H04qke9QYMVPegiQQufFFBj32c=
      -----END CERTIFICATE-----
    orderable: true

  - name: connSecurityProtocol
    description: |-
      Specifies the connection type for the LDAP provider. `TLS` or `InbandTLS`
      (default).
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - TLS
    - InbandTLS
    default_value: InbandTLS
    filterable: true
    orderable: true

  - name: default
    description: |-
      If set, this will be the default LDAP provider. There can be only one default
      provider in your account. When logging in with LDAP, if no provider name is
      given, the default will be used.
    type: boolean
    exposed: true
    stored: true

  - name: ignoredKeys
    description: |-
      A list of keys that must not be imported into a Microsegmentation authorization.
      If
      `includedKeys` is also set, and a key is in both lists, the key will be ignored.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: includedKeys
    description: |-
      A list of keys that must be imported into a Microsegmentation authorization. If
      `ignoredKeys` is also set, and a key is in both lists, the key will be ignored.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: subjectKey
    description: |-
      The key to be used to populate the subject of the Midgard token. If you want to
      use the user as a subject, for Windows-based systems you may use
      `sAMAccountName`.
      For Linux and other systems, you may wish to use `uid` (default). You can also
      use
      any alternate key.
    type: string
    exposed: true
    stored: true
    default_value: uid
    orderable: true
