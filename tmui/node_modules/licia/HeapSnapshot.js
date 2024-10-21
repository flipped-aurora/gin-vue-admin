var Class = require('./Class');
var toBool = require('./toBool');
var camelCase = require('./camelCase');
var LinkedList = require('./LinkedList');
var isStr = require('./isStr');
var each = require('./each');
var map = require('./map');
exports = Class({
    initialize: function HeapSnapshot(profile) {
        if (isStr(profile)) {
            profile = JSON.parse(profile);
        }
        this.nodes = new LinkedList();
        this.edges = new LinkedList();
        var snapshot = profile.snapshot;
        var meta = snapshot.meta;
        this.nodeFields = map(meta.node_fields, camelCase);
        this.nodeTypes = meta.node_types[this.nodeFields.indexOf('type')];
        this.edgeFields = map(meta.edge_fields, camelCase);
        this.edgeTypes = meta.edge_types[this.edgeFields.indexOf('type')];
        this._init(profile);
    },
    _init: function(profile) {
        var _this = this;
        var nodes = profile.nodes,
            edges = profile.edges,
            strings = profile.strings;
        var nodeFields = this.nodeFields,
            edgeFields = this.edgeFields;
        var curEdgeIdx = 0;
        var nodeFieldCount = nodeFields.length;
        var edgeFieldCount = edgeFields.length;
        var nodeMap = {};
        for (var i = 0, len = nodes.length; i < len; i += nodeFieldCount) {
            var node = new Node(this);
            node.init(nodes.slice(i, i + nodeFieldCount), strings);
            this.nodes.push(node);
            nodeMap[i] = node;
        }
        this.nodes.forEach(function(node) {
            var edgeCount = node.edgeCount;
            delete node.edgeCount;
            var maxEdgeIdx = curEdgeIdx + edgeCount * edgeFieldCount;
            for (var _i = curEdgeIdx; _i < maxEdgeIdx; _i += edgeFieldCount) {
                var edge = new Edge(_this, node);
                edge.init(
                    edges.slice(_i, _i + edgeFieldCount),
                    strings,
                    nodeMap
                );
                _this.edges.push(edge);
            }
            curEdgeIdx = maxEdgeIdx;
        });
    }
});
var Node = Class({
    initialize: function Node(heapSnapshot) {
        this._heapSnapshot = heapSnapshot;
    },
    init: function(fields, strings) {
        var _this2 = this;
        var heapSnapshot = this._heapSnapshot;
        var nodeFields = heapSnapshot.nodeFields,
            nodeTypes = heapSnapshot.nodeTypes;
        each(nodeFields, function(field, idx) {
            var val = fields[idx];
            switch (field) {
                case 'name':
                    val = strings[val];
                    break;
                case 'detachedness':
                    val = toBool(val);
                    break;
                case 'type':
                    val = nodeTypes[val];
                    break;
            }
            _this2[field] = val;
        });
    }
});
var Edge = Class({
    initialize: function Edge(heapSnapshot, fromNode) {
        this._heapSnapshot = heapSnapshot;
        this.fromNode = fromNode;
    },
    init: function(fields, strings, nodeMap) {
        var _this3 = this;
        var heapSnapshot = this._heapSnapshot;
        var edgeFields = heapSnapshot.edgeFields,
            edgeTypes = heapSnapshot.edgeTypes;
        each(edgeFields, function(field, idx) {
            var val = fields[idx];
            switch (field) {
                case 'nameOrIndex':
                    val = strings[val];
                    break;
                case 'type':
                    val = edgeTypes[val];
                    break;
                case 'toNode':
                    val = nodeMap[val];
                    break;
            }
            _this3[field] = val;
        });
    }
});

module.exports = exports;
