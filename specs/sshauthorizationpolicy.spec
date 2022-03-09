# Model
model:
  rest_name: sshauthorizationpolicy
  resource_name: sshauthorizationpolicies
  entity_name: SSHAuthorizationPolicy
  package: squall
  group: policy/ssh
  description: |-
    An SSH authorization allows you to define the permissions for the owner
    of a OpenSSH certificate issued by a Microsegmentation certificate authority.
    You can
    define if a user with some claims can connect to an `sshd` server managed by
    an instance of `enforcerd` according to its tags, what permissions he has and
    for how long delivered certificates are valid.
  aliases:
  - sshpol
  - sshpols
  get:
    description: Retrieves the SSH authorization with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the SSH authorization with the given ID.
  delete:
    description: Deletes the SSH authorization with the given ID.
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
  - '@fallback'
  - '@schedulable'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: authorizedSubnets
    description: |-
      If set, the SSH authorization will only be valid if the request comes from one
      the declared subnets.
    type: list
    exposed: true
    subtype: string
    validations:
    - $optionalcidrs

  - name: expirationTime
    description: If set the SSH authorization will be automatically deleted after
      the given time.
    type: time
    exposed: true
    stored: true
    getter: true
    setter: true

  - name: extensions
    description: |-
      The list of permissions to apply to the OpenSSH certificate. You can check the
      list of
      standard extensions at
      <https://github.com/openssh/openssh-portable/blob/38e83e4f219c752ebb1560633b73f06f0392018b/PROTOCOL.certkeys#L281>.
    type: list
    exposed: true
    subtype: string

  - name: forceCommand
    description: |-
      Specify a single command that the user can issue on the remote host. This can be
      useful
      for issuing single-purpose certificates; ensuring that users stay in their home
      directories
      (`internal-sftp`); and restricting users to a bash shell (`/bin/bash`),
      preventing them
      from running arbitrary and unlogged commands such as `scp`, `rsync`, `-essh`,
      and `sftp`.
      Refer to the [FreeBSD
      documentation](https://www.freebsd.org/cgi/man.cgi?sshd_config(5))
      for more information.
    type: string
    exposed: true

  - name: object
    description: |-
      Contains the tag expression identifying the enforcers on the hosts the `subject`
      is allowed to access.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

  - name: principals
    description: |-
      On systems without an enforcer, you must provide the name of the Linux user.
      Otherwise, Microsegmentation will automatically populate this field and adding a
      value here is optional and not used during the authorization. However, the value
      becomes a tag associated with the SSH processing unit, which could be useful.
    type: list
    exposed: true
    subtype: string

  - name: requireSystemAccountMatching
    description: If selected, the system account will be used to log into the resource.
    type: boolean
    exposed: true
    stored: true

  - name: subject
    description: |-
      Contains the tag expression that identifies the user or group of users that
      should be
      allowed to access the remote hosts. If the user authenticates against an OIDC
      provider,
      these tags correspond to claims in the ID token.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

  - name: validity
    description: Set the validity of the delivered SSH certificate.
    type: string
    exposed: true
    default_value: 1h
    validations:
    - $timeDuration
