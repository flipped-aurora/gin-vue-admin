var noop = require('./noop');
var nextTick = require('./nextTick');
var restArgs = require('./restArgs');
exports = function(tasks, cb) {
    cb = cb || noop;
    var current = 0;
    var taskCb = restArgs(function(err, args) {
        if (++current >= tasks.length || err) {
            args.unshift(err);
            nextTick(function() {
                cb.apply(null, args);
            });
        } else {
            args.push(taskCb);
            tasks[current].apply(null, args);
        }
    });
    if (tasks.length) {
        tasks[0](taskCb);
    } else {
        nextTick(function() {
            cb();
        });
    }
};

module.exports = exports;
