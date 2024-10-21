var each = require('./each');
var arrToMap = require('./arrToMap');
var startWith = require('./startWith');
var endWith = require('./endWith');
var extend = require('./extend');

var noPromiseMethods = arrToMap([
    'stopRecord',
    'getRecorderManager',
    'pauseVoice',
    'stopVoice',
    'pauseBackgroundAudio',
    'stopBackgroundAudio',
    'getBackgroundAudioManager',
    'createAudioContext',
    'createInnerAudioContext',
    'createVideoContext',
    'createCameraContext',
    'createMapContext',
    'canIUse',
    'startAccelerometer',
    'stopAccelerometer',
    'startCompass',
    'stopCompass',
    'onBLECharacteristicValueChange',
    'onBLEConnectionStateChange',
    'hideToast',
    'hideLoading',
    'showNavigationBarLoading',
    'hideNavigationBarLoading',
    'navigateBack',
    'createAnimation',
    'pageScrollTo',
    'createSelectorQuery',
    'createCanvasContext',
    'createContext',
    'drawCanvas',
    'hideKeyboard',
    'stopPullDownRefresh',
    'arrayBufferToBase64',
    'base64ToArrayBuffer'
]);
function needToPromisify(name) {
    return (
        !noPromiseMethods[name] &&
        !startWith(name, 'on') &&
        !endWith(name, 'Sync')
    );
}
each(wx, function(fn, name) {
    if (!needToPromisify(name)) return;
    exports[name] = function(obj) {
        return new Promise(function(resolve, reject) {
            fn.call(
                wx,
                extend(obj, {
                    success: function(res) {
                        resolve(res);
                    },
                    fail: function(res) {
                        reject(res);
                    }
                })
            );
        });
    };
});

module.exports = exports;
