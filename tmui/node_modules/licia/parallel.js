var noop = require('./noop');
var each = require('./each');
var nextTick = require('./nextTick');
exports = function(tasks, cb) {
    cb = cb || noop;
    var results = [];
    var pending = tasks.length;
    if (!pending) return done(null);
    each(tasks, function(task, i) {
        task(function(err, result) {
            taskCb(i, err, result);
        });
    });
    function taskCb(i, err, result) {
        results[i] = result;
        if (--pending === 0 || err) done(err);
    }
    function done(err) {
        nextTick(function() {
            cb(err, results);
            cb = noop;
        });
    }
};

module.exports = exports;
