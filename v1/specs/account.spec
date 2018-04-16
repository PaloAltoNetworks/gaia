# Model
model:
  rest_name: account
  resource_name: accounts
  entity_name: Account
  package: vince
  description: |-
    This api allows to view and manage basic information about your account like
    your name, password, enable 2 factor authentication.
  get: true
  update: true
  delete: true
  extends:
  - '@identifiable-pk-stored'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: LDAPAddress
    description: LDAPAddress holds the account authentication account's private ldap
      server.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: LDAPBaseDN
    description: LDAPBaseDN holds the base DN to use to ldap queries.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: LDAPBindDN
    description: LDAPBindDN holds the account's internal LDAP bind dn for querying
      auth.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: LDAPBindPassword
    description: LDAPBindPassword holds the password to the LDAPBindDN.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: LDAPBindSearchFilter
    description: |-
      LDAPBindSearchFilter holds filter to be used to uniquely search a user. For
      Windows based systems, value may be 'sAMAccountName={USERNAME}'. For Linux and
      other systems, value may be 'uid={USERNAME}'.
    type: string
    exposed: true
    stored: true
    default_value: uid={USERNAME}
    filterable: true
    format: free
    orderable: true

  - name: LDAPCertificateAuthority
    description: |-
      LDAPCertificateAuthority contains the optional certificate author ity that will
      be used to connect to the LDAP server. It is not needed if the TLS certificate
      of the LDAP is issued from a public truster CA.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: LDAPConnSecurityProtocol
    description: LDAPConnProtocol holds the connection type for the LDAP provider.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - TLS
    - InbandTLS
    default_value: InbandTLS
    filterable: true
    orderable: true

  - name: LDAPEnabled
    description: LDAPEnabled triggers if the account uses it's own LDAP for authentication.
    type: boolean
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: OTPEnabled
    description: Set to enable or disable two factor authentication.
    type: boolean
    exposed: true
    stored: true

  - name: OTPQRCode
    description: Returns the base64 encoded QRCode for setting up 2 factor auth.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    format: free
    transient: true

  - name: OTPSecret
    description: Stores the 2 factor secret.
    type: string
    stored: true
    format: free

  - name: accessEnabled
    description: AccessEnabled defines if the account holder should have access to
      the systems.
    type: boolean
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: activationExpiration
    description: ActivationExpiration contains the expiration date of the activation
      token.
    type: time
    stored: true
    autogenerated: true

  - name: activationToken
    description: ActivationToken contains the activation token.
    type: string
    stored: true
    autogenerated: true
    format: free

  - name: associatedAPIAuthPolicyID
    description: AssociatedAPIAuthPolicyID holds the ID of the associated API auth
      policy.
    type: string
    stored: true
    format: free

  - name: associatedAWSPolicies
    description: AssociatedAWSPolicies contains a map of associated AWS Enforcerd
      Policies.
    type: external
    subtype: associated_policies
    stored: true

  - name: associatedNamespaceID
    description: AssociatedNamespaceID contains the ID of the associated namespace.
    type: string
    stored: true
    format: free

  - name: associatedPlanKey
    description: AssociatedPlanKey contains the plan key that is associated to this
      account.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    default_value: aporeto.plan.free
    format: free

  - name: associatedQuotaPolicies
    description: AssociatedQuotaPolicies contains a mapping to the associated quota
      pollicies.
    type: external
    subtype: associated_policies
    stored: true

  - name: company
    description: Company of the account user.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: email
    description: Email of the account holder.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: user@aporeto.com
    filterable: true
    format: free
    orderable: true

  - name: firstName
    description: First Name of the account user.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: lastName
    description: Last Name of the account user.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: name
    description: Name of the account.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    allowed_chars: ^[^\*\=]*$
    example_value: acme
    filterable: true
    format: free
    orderable: true

  - name: password
    description: Password for the account.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: reCAPTCHAKey
    description: ReCAPTCHAKey contains the capcha validation if reCAPTCH is enabled.
    type: string
    exposed: true
    creation_only: true
    format: free

  - name: resetPasswordExpiration
    description: ResetPasswordExpiration contains the expiration time for reseting
      the password.
    type: time
    stored: true
    autogenerated: true

  - name: resetPasswordToken
    description: ResetPasswordToken contains the token to use for resetting password.
    type: string
    stored: true
    autogenerated: true
    format: free

  - name: status
    description: Status of the account.
    type: enum
    exposed: true
    stored: true
    read_only: true
    allowed_choices:
    - Active
    - Disabled
    - Invited
    - Pending
    autogenerated: true
    default_value: Pending
    filterable: true
    orderable: true
