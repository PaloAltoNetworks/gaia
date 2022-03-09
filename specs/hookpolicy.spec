# Model
model:
  rest_name: hookpolicy
  resource_name: hookpolicies
  entity_name: HookPolicy
  package: squall
  group: policy/hooks
  description: |-
    Allows you to define hooks to the write operations in squall. Hooks are sent
    to an external Rufus server that will do the processing and eventually return a
    modified version of the object before we save it.
  aliases:
  - hook
  - hooks
  - hookpol
  - hookpols
  get:
    description: Retrieves the hook with the given ID.
  update:
    description: Updates the hook with the given ID.
  delete:
    description: Deletes the hook with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@namespaced'
  - '@described'
  - '@disabled'
  - '@identifiable-not-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@hidden'
  - '@fallback'
  - '@timeable'
  validations:
  - $hookpolicy

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: certificateAuthority
    description: Contains the PEM block of the certificate authority used by the remote
      endpoint.
    type: string
    exposed: true
    stored: true
    example_value: |-
      -----BEGIN CERTIFICATE-----
      MIIBbjCCARSgAwIBAgIRANRbvVzTzBZOvMCb8BiKCLowCgYIKoZIzj0EAwIwJjEN
      MAsGA1UEChMEQWNtZTEVMBMGA1UEAxMMQWNtZSBSb290IENBMB4XDTE4MDExNTE4
      NDgwN1oXDTI3MTEyNDE4NDgwN1owJjENMAsGA1UEChMEQWNtZTEVMBMGA1UEAxMM
      QWNtZSBSb290IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJ/80HR51+vau
      7XH7zS7b8ABA0e/TdBOg1NznbnXdXil1tDvWloWuH5+/bbaiEg54wksJHFXaukw8
      jhTLU7zT56MjMCEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wCgYI
      KoZIzj0EAwIDSAAwRQIhALwAZh2KLFFC1qfb5CqFHExlXS0PUltax9PvQCN9P0vl
      AiBl7/st9u/JpERjJgirxJxOgKNlV6pq9ti75EfQtZZcQA==
      -----END CERTIFICATE-----
    orderable: true
    validations:
    - $pem

  - name: clientCertificate
    description: |-
      Contains the client certificate that will be used to connect
      to the remote endpoint. If provided, the private key associated with this
      certificate must also be configured.
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
    orderable: true
    validations:
    - $pem

  - name: clientCertificateKey
    description: |-
      Contains the key associated with the `clientCertificate`. It must be provided
      only
      when `clientCertificate` has been configured.
    type: string
    exposed: true
    stored: true
    example_value: |-
      -----BEGIN EC PRIVATE KEY-----
      MHcCAQEEIGOXJI/123456789oamOu4tQAIKFdbyvkIJg9GME0mHzoAoGCCqGSM49
      AwEHoUQDQgAE6bM8mP123456789AfmBWtnucfByQXk568lDcKNIQx6yNn+7txbwg
      F9eXFkofGX3UgRtsHe123456789xQ1naSw==
      -----END EC PRIVATE KEY-----
    orderable: true
    secret: true
    transient: true
    encrypted: true
    validations:
    - $pem

  - name: continueOnError
    description: |-
      If set to `true` and `mode` is in `Pre`, the request will be honored even if
      calling the hook fails.
    type: boolean
    exposed: true
    stored: true

  - name: endpoint
    description: Contains the full address of the remote processor endpoint.
    type: string
    exposed: true
    stored: true
    example_value: https://hooks.hookserver.com/remoteprocessors
    orderable: true

  - name: endpointType
    description: Defines the type of endpoint for the hook.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - URL
    - Automation
    default_value: URL
    orderable: true

  - name: expirationTime
    description: If set the hook will be automatically deleted after the given time.
    type: time
    exposed: true
    stored: true
    getter: true
    setter: true

  - name: mode
    description: Defines the type of hook.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Both
    - Post
    - Pre
    default_value: Pre
    orderable: true

  - name: selectors
    description: |-
      A tag or tag expression that identifies the automation that must be run in
      case no endpoint is provided.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    example_value:
    - - automation:name=myautomation
    validations:
    - $tagsExpression

  - name: subject
    description: |-
      Contains the tag expression that an object must match in order to trigger the
      hook.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    example_value:
    - - $identity=processingunit
    validations:
    - $tagsExpression

  - name: triggerOperations
    description: |-
      Select on which operation(s) you want to the hook to trigger. An empty list.
      Only
      means all operations. You can only set any combination of `create`, `update` or
      `delete`. Any other value will trigger a validation error.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $writeoperations
