# Model
model:
  rest_name: enforcersecret
  resource_name: enforcersecrets
  entity_name: EnforcerSecret
  package: squall
  group: internal/x509
  description: Manages private keys.
  get:
    description: Retrieves enforcer secrets information.

# Attributes
attributes:
  v1:
  - name: syslogCertificate
    description: Syslog public key in PEM format.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: |-
      -----BEGIN CERTIFICATE-----
      MIIBKjCB0aADAgECAhEAugDV416p70g1FGz9A91H4DAKBggqhkjOPQQDAjAOMQww
      CgYDVQQDEwNjdWwwHhcNMjEwNzA2MTYzMjI4WhcNMjIwNzAxMTYzMjI4WjAOMQww
      CgYDVQQDEwNjdWwwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQHNsAal/jSVlW6
      k42JKaMpwrgFLqHGTjUsZdexNraTO2qGSb6x9X7yTDI4b17AjgYMeBP/qvBfcYQi
      52CVnp5xoxAwDjAMBgNVHRMBAf8EAjAAMAoGCCqGSM49BAMCA0gAMEUCICZcMM89
      fJz9LMdpz/A1RpBEC0Xx4I/8JWrHroVYPOG7AiEA37iWAAi1DYBboyxevbCc6kNa
      Su0pBE163iBEKUew0/s=
      -----END CERTIFICATE-----
    validations:
    - $pem

  - name: syslogCertificateKey
    description: Syslog private key in PEM format.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: |-
      -----BEGIN EC PRIVATE KEY-----
      MHcCAQEEIHjrDddXMasytnPC2+7m2BnkgEX7a6Vk1G13dQsASpAhoAoGCCqGSM49
      AwEHoUQDQgAEBzbAGpf40lZVupONiSmjKcK4BS6hxk41LGXXsTa2kztqhkm+sfV+
      8kwyOG9ewI4GDHgT/6rwX3GEIudglZ6ecQ==
      -----END EC PRIVATE KEY-----
    encrypted: true
    validations:
    - $pem
