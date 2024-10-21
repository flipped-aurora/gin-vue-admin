var Url = require('./Url');
exports = function(name, url) {
    return new Url(url).query[name];
};

module.exports = exports;
