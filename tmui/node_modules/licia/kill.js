const isWindows = require('./isWindows');

const childProcess = require('child_process');

exports = function(pid) {
    try {
        let cmd = '';
        if (isWindows) {
            cmd = `taskkill /pid ${pid} /T /F`;
        } else {
            cmd = `kill ${pid} -9`;
        }
        childProcess.execSync(cmd);
    } catch (e) {
        /* eslint-disable no-empty */
    }
};

module.exports = exports;
