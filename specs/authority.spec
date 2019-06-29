# Model
model:
  rest_name: authority
  resource_name: authorities
  entity_name: Authority
  package: barret
  group: internal/x509
  description: Authority represents a certificate authority.
  aliases:
  - ca
  private: true
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@identifiable-stored'
  - '@zoned'

# Indexes
indexes:
- - commonName

# Attributes
attributes:
  v1:
  - name: Algorithm
    description: Algorithm defines the the signing algorithm to be used.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - ECDSA
    - RSA
    default_value: ECDSA

  - name: certificate
    description: PEM encoded certificate data.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: commonName
    description: CommonName contains the common name of the certificate.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: my ca

  - name: expirationDate
    description: Date of expiration of the issued certificate.
    type: time
    exposed: true
    stored: true
    creation_only: true

  - name: key
    description: Encrypted private key of the Authority.
    type: string
    stored: true
    autogenerated: true

  - name: serialNumber
    description: serialNumber of the certificate.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: type
    description: Type of signing authority can be a CA or a JWT signing certificate.
    type: enum
    exposed: true
    stored: true
    read_only: true
    allowed_choices:
    - CA
    - TokenSigning
    autogenerated: true
    default_value: CA
