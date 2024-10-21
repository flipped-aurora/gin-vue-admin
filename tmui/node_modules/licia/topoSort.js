exports = function(edges) {
    return sort(uniqueNodes(edges), edges);
};
function uniqueNodes(arr) {
    var ret = [];
    for (var i = 0, len = arr.length; i < len; i++) {
        var edge = arr[i];
        if (ret.indexOf(edge[0]) < 0) ret.push(edge[0]);
        if (ret.indexOf(edge[1]) < 0) ret.push(edge[1]);
    }
    return ret;
}
function sort(nodes, edges) {
    var cursor = nodes.length;
    var sorted = new Array(cursor);
    var visited = {};
    var i = cursor;
    while (i--) {
        if (!visited[i]) visit(nodes[i], i, []);
    }
    function visit(node, i, predecessors) {
        if (predecessors.indexOf(node) >= 0) {
            throw new Error('Cyclic dependency: ' + JSON.stringify(node));
        }
        if (visited[i]) return;
        visited[i] = true;
        var outgoing = edges.filter(function(edge) {
            return edge[0] === node;
        });

        if ((i = outgoing.length)) {
            var preds = predecessors.concat(node);
            do {
                var child = outgoing[--i][1];
                visit(child, nodes.indexOf(child), preds);
            } while (i);
        }
        sorted[--cursor] = node;
    }
    return sorted;
}

module.exports = exports;
