var create = require('./create');
exports = function(Class, SuperClass) {
    Class.prototype = create(SuperClass.prototype);
};

module.exports = exports;
