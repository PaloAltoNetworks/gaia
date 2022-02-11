# Model
model:
  rest_name: dnslookupreport
  resource_name: dnslookupreports
  entity_name: DNSLookupReport
  package: zack
  group: policy/dns
  description: |-
    A DNS lookup report is used to report a DNS lookup that is happening on
    behalf of a processing unit. If the DNS server is on the standard UDP port 53
    then the enforcer can proxy the DNS traffic and make a report. The report
    indicate whether or not the lookup was successful.
  extends:
  - '@identifiable-stored'
  - '@zoned-monotonic'
  - '@migratable'
  - '@namespaced'
  validations:
  - $dnslookupreport

# Ordering
default_order:
- :no-inherit
- timestamp

# Indexes
indexes:
- - namespace
  - timestamp

# Attributes
attributes:
  v1:
  - name: action
    description: Action of the DNS request.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Accept
    - Fail
    - Reject
    - Resolve
    example_value: Accept
    omit_empty: true
    extensions:
      bson_name: a

  - name: enforcerDNSReportID
    description: The local ID set by enforcer, used to correllate with flow reports.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: m

  - name: enforcerID
    description: ID of the enforcer.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: b

  - name: enforcerNamespace
    description: Namespace of the enforcer.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/namespace
    omit_empty: true
    extensions:
      bson_name: c

  - name: processingUnitID
    description: ID of the PU.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: d

  - name: processingUnitNamespace
    description: Namespace of the PU. This is deprecated. Use `namespace` instead.
    type: string
    exposed: true
    stored: true
    deprecated: true
    example_value: /my/namespace
    omit_empty: true
    extensions:
      bson_name: e

  - name: reason
    description: |-
      This field is only set when the lookup fails. It specifies the reason for the
      failure.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: f

  - name: resolvedCNAMEs
    description: CNAME aliases.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true
    extensions:
      bson_name: k

  - name: resolvedIPs
    description: resolved IP addresses.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true
    validations:
    - $optionalcidroriplist
    extensions:
      bson_name: l

  - name: resolvedName
    description: name used for DNS resolution.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: www.google.com
    omit_empty: true
    extensions:
      bson_name: g

  - name: sourceIP
    description: Type of the source.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 10.0.0.1
    omit_empty: true
    extensions:
      bson_name: h

  - name: timestamp
    description: Time and date of the log.
    type: time
    exposed: true
    stored: true
    orderable: true
    omit_empty: true
    extensions:
      bson_name: i

  - name: value
    description: Number of times the client saw this activity.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 1
    omit_empty: true
    extensions:
      bson_name: j
