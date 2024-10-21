var memStorage = require('./memStorage');
exports = function(type) {
    type = type || 'local';
    var ret;
    switch (type) {
        case 'local':
            ret = window.localStorage;
            break;
        case 'session':
            ret = window.sessionStorage;
            break;
    }
    try {
        var x = 'test-localStorage-' + Date.now();
        ret.setItem(x, x);
        var y = ret.getItem(x);
        ret.removeItem(x);
        if (y !== x) throw new Error();
    } catch (e) {
        return memStorage;
    }
    return ret;
};

module.exports = exports;
