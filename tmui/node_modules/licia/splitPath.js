exports = function(path) {
    var match = path.match(regSplit);
    return {
        dir: match[1],
        name: match[2],
        ext: match[3]
    };
};
var regSplit = /^([\s\S]*?)((?:\.{1,2}|[^\\/]+?|)(\.[^./\\]*|))(?:[\\/]*)$/;

module.exports = exports;
