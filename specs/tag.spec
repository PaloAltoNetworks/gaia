{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Count represents the number of time the tag is used.",
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
            "name": "count",
            "orderable": false,
            "primary_key": null,
            "read_only": true,
            "required": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "integer",
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
            "description": "Namespace represents the namespace of the counted tag.",
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
            "name": "namespace",
            "orderable": false,
            "primary_key": true,
            "read_only": true,
            "required": false,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": "^[\\w\\d\\*\\$\\+\\.:,|@<>/-]+=[= \\w\\d\\*\\$\\+\\.:,|@<>/-]+$",
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": true,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Value represents the value of the tag.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": false,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "value",
            "orderable": false,
            "primary_key": true,
            "read_only": null,
            "required": true,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "create": null,
        "delete": false,
        "description": "A Tag represents a tags associated to an object.",
        "entity_name": "Tag",
        "extends": [
            "@identifiable-nopk-nostored"
        ],
        "get": false,
        "package": "Policies",
        "resource_name": "tags",
        "rest_name": "tag",
        "root": null,
        "update": false
    }
}
