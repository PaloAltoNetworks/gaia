# Model
model:
  rest_name: pccprovider
  resource_name: pccproviders
  entity_name: PCCProvider
  package: cactuar
  group: core/authentication
  description: |-
    Allows you to declare a trusted Prisma Cloud Compute (PCC) authentication
    provider. Microsegmentation will accept JSON web tokens (JWT) from the specified
    PCC provider.
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
      Set the CA to use to contact the PCC Console in case it uses a non widely
      trusted
      certificate authority.
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

  - name: default
    description: |-
      If set, this will be the default PCC provider. There can be only one default
      provider in your account. When logging in with PCC, if no provider name is
      given, the default will be used.
    type: boolean
    exposed: true
    stored: true

  - name: endpoint
    description: The URL of the PCC service. It must use HTTPS.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: https://my.pcc.acme.com
