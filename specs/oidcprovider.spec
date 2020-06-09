# Model
model:
  rest_name: oidcprovider
  resource_name: oidcproviders
  entity_name: OIDCProvider
  package: cactuar
  group: core/authentication
  description: |-
    Allows you to declare a generic OpenID Connect (OIDC) provider that can be used in
    exchange
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
  - '@identifiable-stored'
  - '@timeable'
  - '@named'

# Attributes
attributes:
  v1:
  - name: certificateAuthority
    description: |-
      Set the CA to use to contact the OIDC server. This is useful when you are using
      a custom OIDC provider that doesn't use a trusted CA. Most of the
      time, you can leave this property empty.
    type: string
    exposed: true
    stored: true
    example_value: |-
      -----BEGIN CERTIFICATE-----
      MIIBczCCARigAwIBAgIRALD3Vz81Pq10g7n4eAkOsCYwCgYIKoZIzj0EAwIwJjEN
      MAsGA1UEChMEQWNtZTEVMBMGA1UEAxMMQWNtZSBSb290IENBMB4XDTE4MDExNzA2
      NTM1MloXDTI3MTEyNjA2NTM1MlowGDEWMBQGA1UEAxMNY2xhaXJlLWNsaWVudDBZ
      MBMGByqGSM49AgEGCCqGSM49AwEHA0IABOmzPJj+t25T148eQH5gVrZ7nHwckF5O
      evJQ3CjSEMesjZ/u7cW8IBfXlxZKHxl91IEbbB3svci4c8pycUNZ2kujNTAzMA4G
      A1UdDwEB/wQEAwIHgDATBgNVHSUEDDAKBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAA
      MAoGCCqGSM49BAMCA0kAMEYCIQCjAAmkQpTua0HR4q6jnePaFBp/JMXwTXTxzbV6
      peGbBQIhAP+1OR8GFnn2PlacwHqWXHwkvy6CLPVikvgtwEdB6jH8
      -----END CERTIFICATE-----
    validations:
    - $pem

  - name: clientID
    description: Unique client ID.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 6195189841830-0644ee9d89ef0644ee9d89examle.apps.googleusercontent.com

  - name: clientSecret
    description: Client secret associated with the client ID.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: Ytgbfjtj4652jHDFGls99jF
    encrypted: true

  - name: default
    description: |-
      If set, this will be the default OIDC provider. There can be only one default
      provider in your account. When logging in with OIDC, if no provider name is
      given, the default will be used.
    type: boolean
    exposed: true
    stored: true

  - name: endpoint
    description: |-
      OIDC [discovery
      endpoint](https://openid.net/specs/openid-connect-discovery-1_0.html#IssuerDiscovery).
    type: string
    exposed: true
    stored: true
    required: true
    example_value: https://accounts.google.com

  - name: parentID
    description: Contains the parent Segment account ID.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    filterable: true

  - name: parentName
    description: Contains the name of the parent Segment account.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    filterable: true

  - name: scopes
    description: List of scopes to allow.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - email
    - profile

  - name: subjects
    description: List of claims that will provide the subject.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - email
    - profile
