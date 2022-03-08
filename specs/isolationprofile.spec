# Model
model:
  rest_name: isolationprofile
  resource_name: isolationprofiles
  entity_name: IsolationProfile
  package: squall
  group: policy/processingunits
  description: |-
    Defines system call rules, system call actions, and other capabilities on a
    processing unit.
  aliases:
  - ip
  get:
    description: Retrieves the profile with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the profile with the given ID.
  delete:
    description: Deletes the profile with the given ID.
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: capabilitiesActions
    description: The capabilities that should be added to or removed from the processing
      unit.
    type: external
    exposed: true
    subtype: _cap_map
    stored: true
    orderable: true

  - name: defaultSyscallAction
    description: |-
      The default action applied to all system calls of this profile.
      Default is `Allow`.
    type: external
    exposed: true
    subtype: _syscall_action
    stored: true

  - name: syscallRules
    description: |-
      A list of system call rules that identify actions for particular
      system calls.
    type: external
    exposed: true
    subtype: _syscall_rules
    stored: true
    orderable: true

  - name: targetArchitectures
    description: The processor architectures that the profile supports. Default `all`.
    type: external
    exposed: true
    subtype: _arch_list
    stored: true
    orderable: true
