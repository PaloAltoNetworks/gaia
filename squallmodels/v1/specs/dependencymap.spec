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
            "description": "edges are the edges of the map",
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
            "name": "edges",
            "orderable": false,
            "primary_key": false,
            "read_only": true,
            "required": true,
            "secret": false,
            "setter": false,
            "stored": false,
            "subtype": "graphedges_map",
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
            "description": "Groups provide information about the group values",
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
            "name": "groups",
            "orderable": false,
            "primary_key": false,
            "read_only": true,
            "required": true,
            "secret": false,
            "setter": false,
            "stored": false,
            "subtype": "graphgroups_map",
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
            "description": "nodes refers to the nodes of the map",
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
            "name": "nodes",
            "orderable": false,
            "primary_key": false,
            "read_only": true,
            "required": true,
            "secret": false,
            "setter": false,
            "stored": false,
            "subtype": "graphnodes_map",
            "transient": false,
            "type": "external",
            "unique": false,
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
            "description": "TagsStatistics provides statistics for relevant tags",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "tagsStatistics",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": null,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "aliases": [
            "depmap",
            "depmaps"
        ],
        "create": false,
        "delete": false,
        "description": "This api returns a data structure representing the graph of all processing units and their connections in a particular namespace, in a given time window. To pass the time window you can use the query parameters \"startAbsolute\", \"endAbsolute\", \"startRelative\", \"endRelative\".  For example \"https://squall.aporeto.com/dependencymaps?startAbsolute=1489132800000&endAbsolute=1489219200000\"",
        "entity_name": "DependencyMap",
        "extends": [
            "@identifiable-nopk-nostored"
        ],
        "get": false,
        "package": "Squall API",
        "resource_name": "dependencymaps",
        "rest_name": "dependencymap",
        "root": null,
        "update": false
    }
}