{
    "attributes": [
        {
            "description": "LDAPAddress holds the account authentication account's private ldap server.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "LDAPAddress",
            "orderable": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "LDAPBaseDN holds the base DN to use to ldap queries.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "LDAPBaseDN",
            "orderable": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "LDAPBindDN holds the account's internal LDAP bind dn for querying auth.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "LDAPBindDN",
            "orderable": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "LDAPBindPassword holds the password to the LDAPBindDN. ",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "LDAPBindPassword",
            "orderable": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "LDAPCertificateAuthority contains the optional certificate authority that will be used to connect to the LDAP server. It is not needed if the TLS certificate of the LDAP is issued from a public truster CA.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "LDAPCertificateAuthority",
            "orderable": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "LDAPEnabled triggers if the account uses it's own LDAP for authentication.",
            "exposed": true,
            "filterable": true,
            "name": "LDAPEnabled",
            "orderable": true,
            "stored": true,
            "type": "boolean"
        },
        {
            "description": "Set to enable or disable two factor authentication.",
            "exposed": true,
            "name": "OTPEnabled",
            "stored": true,
            "type": "boolean"
        },
        {
            "autogenerated": true,
            "description": "Returns the base64 encoded QRCode for setting up 2 factor auth.",
            "exposed": true,
            "format": "free",
            "name": "OTPQRCode",
            "read_only": true,
            "transient": true,
            "type": "string"
        },
        {
            "description": "Stores the 2 factor secret",
            "format": "free",
            "name": "OTPSecret",
            "stored": true,
            "type": "string"
        },
        {
            "description": "AccessEnabled defines if the account holder should have access to the systems.",
            "exposed": true,
            "filterable": true,
            "name": "accessEnabled",
            "orderable": true,
            "stored": true,
            "type": "boolean"
        },
        {
            "autogenerated": true,
            "description": "ActivationExpiration contains the expiration date of the activation token.",
            "name": "activationExpiration",
            "stored": true,
            "type": "time"
        },
        {
            "autogenerated": true,
            "description": "ActivationToken contains the activation token.",
            "format": "free",
            "name": "activationToken",
            "stored": true,
            "type": "string"
        },
        {
            "description": "AssociatedAPIAuthPolicyID holds the ID of the associated API auth policy.",
            "format": "free",
            "name": "associatedAPIAuthPolicyID",
            "stored": true,
            "type": "string"
        },
        {
            "description": "AssociatedAWSPolicies contains a map of associated AWS Enforcerd Policies.",
            "name": "associatedAWSPolicies",
            "stored": true,
            "subtype": "associated_policies",
            "type": "external"
        },
        {
            "description": "AssociatedNamespaceID contains the ID of the associated namespace.",
            "format": "free",
            "name": "associatedNamespaceID",
            "stored": true,
            "type": "string"
        },
        {
            "autogenerated": true,
            "default_value": "aporeto.plan.free",
            "description": "AssociatedPlanKey contains the plan key that is associated to this account.",
            "exposed": true,
            "format": "free",
            "name": "associatedPlanKey",
            "read_only": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "AssociatedQuotaPolicies contains a mapping to the associated quota pollicies.",
            "name": "associatedQuotaPolicies",
            "stored": true,
            "subtype": "associated_policies",
            "type": "external"
        },
        {
            "description": "Company of the account user.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "company",
            "orderable": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "Email of the account holder.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "email",
            "orderable": true,
            "required": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "First Name of the account user.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "firstName",
            "orderable": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "Last Name of the account user.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "lastName",
            "orderable": true,
            "stored": true,
            "type": "string"
        },
        {
            "allowed_chars": "^[^\\*\\=]*$",
            "creation_only": true,
            "description": "Name of the account.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "name",
            "orderable": true,
            "required": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "Password for the account.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "password",
            "orderable": true,
            "stored": true,
            "type": "string"
        },
        {
            "creation_only": true,
            "description": "ReCAPTCHAKey contains the capcha validation if reCAPTCH is enabled.",
            "exposed": true,
            "format": "free",
            "name": "reCAPTCHAKey",
            "type": "string"
        },
        {
            "autogenerated": true,
            "description": "ResetPasswordExpiration contains the expiration time for reseting the password.",
            "name": "resetPasswordExpiration",
            "stored": true,
            "type": "time"
        },
        {
            "autogenerated": true,
            "description": "ResetPasswordToken contains the token to use for resetting password.",
            "format": "free",
            "name": "resetPasswordToken",
            "stored": true,
            "type": "string"
        },
        {
            "allowed_choices": [
                "Active",
                "Disabled",
                "Invited",
                "Pending"
            ],
            "autogenerated": true,
            "default_value": "Pending",
            "description": "Status of the account.",
            "exposed": true,
            "filterable": true,
            "name": "status",
            "orderable": true,
            "read_only": true,
            "stored": true,
            "type": "enum"
        }
    ],
    "model": {
        "delete": true,
        "get": true,
        "update": true,
        "description": "Manage your Account.",
        "entity_name": "Account",
        "extends": [
            "@identifiable-pk-stored",
            "@timeable"
        ],
        "package": "vince",
        "resource_name": "accounts",
        "rest_name": "account"
    }
}
