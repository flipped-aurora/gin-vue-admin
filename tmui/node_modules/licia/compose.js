var restArgs = require('./restArgs');
exports = restArgs(function(fnList) {
    return function() {
        var i = fnList.length - 1;
        var result = fnList[i].apply(this, arguments);
        while (i--) result = fnList[i].call(this, result);
        return result;
    };
});

module.exports = exports;
