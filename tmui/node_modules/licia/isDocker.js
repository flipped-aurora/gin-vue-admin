const memoize = require('./memoize');
const contain = require('./contain');

const fs = require('fs');

exports = memoize(function() {
    try {
        fs.statSync('/.dockerenv');
        return true;
    } catch (e) {}

    return hasDocker('/proc/self/cgroup') || hasDocker('/proc/self/mountinfo');
});

function hasDocker(file) {
    try {
        return contain(fs.readFileSync(file, 'utf8'), 'docker');
    } catch (e) {
        return false;
    }
}

module.exports = exports;
