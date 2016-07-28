{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "edges are the edges of the map",
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
            "name": "edges",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": true,
            "setter": false,
            "stored": false,
            "subtype": "edge_list",
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
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "nodes refers to the nodes of the map",
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
            "name": "nodes",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": true,
            "setter": false,
            "stored": false,
            "subtype": "node_list",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "create": false,
        "delete": false,
        "description": "dependencymap creates a map of dependencies.",
        "entity_name": "DependencyMap",
        "extends": [],
        "get": false,
        "package": "Visualization",
        "resource_name": "dependencymaps",
        "rest_name": "dependencymap",
        "root": null,
        "update": false
    }
}