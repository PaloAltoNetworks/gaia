{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Identifier of an entity",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": "free",
            "getter": false,
            "identifier": true,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "ID",
            "orderable": true,
            "primary_key": true,
            "read_only": false,
            "required": false,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": true,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Annotation stores additional information about an entity",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "annotation",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": false,
            "stored": true,
            "subtype": "annotation",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "AssociatedTags are the list of tags attached to an entity",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": true,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "associatedTags",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": true,
            "stored": true,
            "subtype": "tags_list",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "CreatedAt is the time at which an entity was created",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "createdAt",
            "orderable": true,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": true,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "time",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Description of the entity",
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
            "name": "description",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Name of the entity",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": "free",
            "getter": true,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "name",
            "orderable": true,
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
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Namespace tag attached to an entity",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": "free",
            "getter": true,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "namespace",
            "orderable": true,
            "primary_key": true,
            "read_only": false,
            "required": false,
            "setter": true,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": true,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Owner who created the entity. It is a list of tags that identify it",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "owner",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": false,
            "stored": true,
            "subtype": "tag_list",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "Active",
                "Candidate",
                "Disabled"
            ],
            "autogenerated": true,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Status of an entity",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": null,
            "getter": true,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "status",
            "orderable": true,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": true,
            "stored": true,
            "subtype": "status_enum",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "UpdatedAt is the time at which an entity was updated",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "updatedAt",
            "orderable": true,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": true,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "time",
            "unique": false,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "create": false,
        "delete": false,
        "description": null,
        "entity_name": null,
        "extends": [],
        "get": false,
        "package": null,
        "resource_name": null,
        "rest_name": null,
        "root": false,
        "update": false
    }
}