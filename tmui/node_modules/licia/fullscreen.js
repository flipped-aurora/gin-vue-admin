var Emitter = require('./Emitter');
var toBool = require('./toBool');
var fnMap = [
    [
        'requestFullscreen',
        'exitFullscreen',
        'fullscreenElement',
        'fullscreenEnabled',
        'fullscreenchange',
        'fullscreenerror'
    ],
    [
        'webkitRequestFullscreen',
        'webkitExitFullscreen',
        'webkitFullscreenElement',
        'webkitFullscreenEnabled',
        'webkitfullscreenchange',
        'webkitfullscreenerror'
    ],
    [
        'mozRequestFullScreen',
        'mozCancelFullScreen',
        'mozFullScreenElement',
        'mozFullScreenEnabled',
        'mozfullscreenchange',
        'mozfullscreenerror'
    ],
    [
        'msRequestFullscreen',
        'msExitFullscreen',
        'msFullscreenElement',
        'msFullscreenEnabled',
        'MSFullscreenChange',
        'MSFullscreenError'
    ]
];
var fn;
for (var i = 0, len = fnMap.length; i < len; i++) {
    fn = fnMap[i];
    if (fn[1] in document) {
        break;
    }
}
exports = {
    request: function() {
        var el =
            arguments.length > 0 && arguments[0] !== undefined
                ? arguments[0]
                : document.documentElement;
        el[fn[0]]();
    },
    exit: function() {
        document[fn[1]]();
    },
    toggle: function() {
        var el =
            arguments.length > 0 && arguments[0] !== undefined
                ? arguments[0]
                : document.documentElement;
        this.isActive() ? this.exit() : this.request(el);
    },
    isActive: function() {
        return toBool(this.getEl());
    },
    isEnabled: function() {
        return toBool(document[fn[3]]);
    },
    getEl: function() {
        return document[fn[2]];
    }
};
Emitter.mixin(exports);
document.addEventListener(fn[4], function() {
    exports.emit('change');
});
document.addEventListener(fn[5], function() {
    exports.emit('error');
});

module.exports = exports;
