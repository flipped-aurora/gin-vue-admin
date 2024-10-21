exports = function(xpath) {
    var ret = [];
    var nodesSnapshot = document.evaluate(
        xpath,
        document,
        null,
        XPathResult.ORDERED_NODE_SNAPSHOT_TYPE,
        null
    );
    for (var i = 0; i < nodesSnapshot.snapshotLength; i++) {
        ret.push(nodesSnapshot.snapshotItem(i));
    }
    return ret;
};

module.exports = exports;
