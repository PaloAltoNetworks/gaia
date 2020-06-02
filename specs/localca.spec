# Model
model:
  rest_name: localca
  resource_name: localcas
  entity_name: LocalCA
  package: squall
  group: core/namespace
  description: |-
    Can be used to retrieve or renew the local and SSH certificate authorities of
    the namespace.

# Attributes
attributes:
  v1:
  - name: Certificate
    description: The certificate authority used by the namespace.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    transient: true

  - name: CertificateRenew
    description: Set to `true` to renew the certificate authority of the namespace.
    type: boolean
    exposed: true
    read_only: true
    autogenerated: true
    transient: true

  - name: SSHCertificate
    description: The SSH certificate authority used by the namespace.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    transient: true

  - name: SSHCertificateRenew
    description: Set to `true` to renew the SSH certificate authority of the namespace.
    type: boolean
    exposed: true
    read_only: true
    autogenerated: true
    transient: true
