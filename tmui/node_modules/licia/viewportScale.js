var meta = require('./meta');
var clamp = require('./clamp');
var trim = require('./trim');
var each = require('./each');
var map = require('./map');
var isNaN = require('./isNaN');
exports = function() {
    var viewport = meta('viewport');
    if (!viewport) return 1;
    viewport = map(viewport.split(','), function(val) {
        return trim(val);
    });
    var minScale = 0.25,
        maxScale = 5,
        initialScale = 1;
    each(viewport, function(val) {
        val = val.split('=');
        var key = val[0];
        val = val[1];
        if (key === 'initial-scale') initialScale = +val;
        if (key === 'maximum-scale') maxScale = +val;
        if (key === 'minimum-scale') minScale = +val;
    });
    var ret = clamp(initialScale, minScale, maxScale);

    if (isNaN(ret)) return 1;
    return ret;
};

module.exports = exports;
