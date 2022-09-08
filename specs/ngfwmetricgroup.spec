# Model
model:
  rest_name: ngfwmetricgroup
  resource_name: ngfwmetricgroups
  entity_name: NGFWMetricGroup
  package: placeholder
  group: placeholder
  description: Defines the Next-Generation Firewall (NGFW) group of metrics.
  get:
    description: Retrieves the NGFW group of metrics with the given ID.
    global_parameters:
    - $propagatable
  extends:
  - '@identifiable-stored'
  - '@namespaced'
  - '@zoned'
  - '@migratable'

# Attributes
attributes:
  v1:
  - name: countCreditsConsumed
    description: The number of Prisma Cloud credits consumed.
    type: integer
    exposed: true
    stored: true
    omit_empty: true

  - name: countInstancesAdvancedLicense
    description: The number of instances with advanced licensing.
    type: integer
    exposed: true
    stored: true
    omit_empty: true

  - name: countInstancesBasicLicense
    description: The number of instances with basic licensing.
    type: integer
    exposed: true
    stored: true
    omit_empty: true

  - name: countInstancesNGFWMode
    description: The number of instances in NGFW mode.
    type: integer
    exposed: true
    stored: true
    omit_empty: true

  - name: countInstancesTAPLicense
    description: The number of instances with TAP licensing.
    type: integer
    exposed: true
    stored: true
    omit_empty: true

  - name: countInstancesTAPMode
    description: The number of instances in TAP mode.
    type: integer
    exposed: true
    stored: true
    omit_empty: true

  - name: countPoliciesAutoRemedyEnabled
    description: The number of NGFW policies with auto-remediation enabled.
    type: integer
    exposed: true
    stored: true
    omit_empty: true

  - name: countTenantsNGFWEnabled
    description: The number of tenants with NGFW enabled.
    type: integer
    exposed: true
    stored: true
    omit_empty: true

  - name: countThreatsBlocked
    description: The number of threats blocked.
    type: integer
    exposed: true
    stored: true
    omit_empty: true

  - name: countThreatsDetected
    description: The number of threats detected.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
