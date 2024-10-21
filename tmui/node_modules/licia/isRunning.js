exports = function(pid) {
    try {
        return process.kill(pid, 0);
    } catch (e) {
        return e.code === 'EPERM';
    }
};

module.exports = exports;
