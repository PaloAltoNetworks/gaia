{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Colums contains the list of tags key you want to use for generating tabulated data.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "columns",
            "orderable": false,
            "primary_key": null,
            "read_only": null,
            "required": true,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": "string",
            "transient": null,
            "type": "list",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Data contains the tabulated data.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "data",
            "orderable": false,
            "primary_key": null,
            "read_only": true,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": "tabulated_data",
            "transient": null,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "TargetIdentity contains the request target identity data to tabulate.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "targetIdentity",
            "orderable": false,
            "primary_key": null,
            "read_only": null,
            "required": true,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "aliases": [
            "tab",
            "tabs",
            "tables",
            "table"
        ],
        "create": null,
        "delete": false,
        "description": "Tabulate API allows you to retrieve a custom table view for any identity using any tags you like as columns.",
        "entity_name": "Tabulate",
        "extends": [],
        "get": false,
        "package": null,
        "resource_name": "tabulate",
        "rest_name": "tabulate",
        "root": null,
        "update": false
    }
}