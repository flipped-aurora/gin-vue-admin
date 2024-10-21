const isWindows = require('./isWindows');

const childProcess = require('child_process');

exports = function(target) {
    let cmd;
    const args = [];

    if (isWindows) {
        cmd = 'cmd';
        args.push('/c', 'start', '""', '/b');
    } else {
        cmd = 'open';
    }

    args.push(target);

    const cp = childProcess.spawn(cmd, args);
    cp.unref();

    return cp;
};

module.exports = exports;
