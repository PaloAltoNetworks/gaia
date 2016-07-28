{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "KeyValuePair stores the value of the tag in `<key>=<value>` format",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": "free",
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "keyValuePair",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": true,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": false,
            "uniqueScope": null
        }
    ],
    "children": [
        {
            "bulk_create": false,
            "bulk_delete": false,
            "bulk_update": false,
            "create": false,
            "delete": false,
            "deprecated": false,
            "get": false,
            "relationship": "member",
            "rest_name": "tag",
            "update": true
        }
    ],
    "model": {
        "create": null,
        "delete": true,
        "description": "Tag to be attached to an entity.",
        "entity_name": "Tag",
        "extends": [
            "@base",
            "@identifiable-nopk-stored"
        ],
        "get": true,
        "package": "Policies",
        "resource_name": "tags",
        "rest_name": "tag",
        "root": null,
        "update": true
    }
}