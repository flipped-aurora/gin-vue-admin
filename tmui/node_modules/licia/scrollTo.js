var Tween = require('./Tween');
var defaults = require('./defaults');
var noop = require('./noop');
var isNum = require('./isNum');
var $offset = require('./$offset');
exports = function(target, options) {
    options = options || {};
    defaults(options, defOpts);
    if (!isNum(target)) target = $offset(target).top;
    new Tween({
        y: window.pageYOffset
    })
        .on('update', function(target) {
            window.scroll(0, target.y);
        })
        .on('end', function() {
            options.callback();
        })
        .to(
            {
                y: target - options.tolerance
            },
            options.duration,
            options.easing
        )
        .play();
};
var defOpts = {
    tolerance: 0,
    duration: 800,
    easing: 'outQuart',
    callback: noop
};

module.exports = exports;
