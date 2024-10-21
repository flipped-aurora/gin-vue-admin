var splitCase = require('./splitCase');
exports = function(str) {
    var arr = splitCase(str);
    var ret = arr[0];
    arr.shift();
    arr.forEach(capitalize, arr);
    ret += arr.join('');
    return ret;
};
function capitalize(val, idx) {
    this[idx] = val.replace(/\w/, function(match) {
        return match.toUpperCase();
    });
}

module.exports = exports;
