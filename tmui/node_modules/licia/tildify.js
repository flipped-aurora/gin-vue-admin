const startWith = require('./startWith');

const path = require('path');
const os = require('os');

const homeDir = os.homedir();

exports = function(p) {
    p = path.normalize(p) + path.sep;
    if (startWith(p, homeDir)) {
        p = p.replace(homeDir + path.sep, `~${path.sep}`);
    }

    return p.slice(0, -1);
};

module.exports = exports;
