exports = function(rgb) {
    var r = rgb[0] / 255;
    var g = rgb[1] / 255;
    var b = rgb[2] / 255;
    var min = mMin(r, g, b);
    var max = mMax(r, g, b);
    var delta = max - min;
    var h;
    var s;
    if (max === min) {
        h = 0;
    } else if (r === max) {
        h = (g - b) / delta;
    } else if (g === max) {
        h = 2 + (b - r) / delta;
    } else {
        h = 4 + (r - g) / delta;
    }
    h = mMin(h * 60, 360);
    if (h < 0) h += 360;
    var l = (min + max) / 2;
    if (max === min) {
        s = 0;
    } else if (l <= 0.5) {
        s = delta / (max + min);
    } else {
        s = delta / (2 - max - min);
    }
    var ret = [round(h), round(s * 100), round(l * 100)];
    if (rgb[3]) ret[3] = rgb[3];
    return ret;
};
var mMin = Math.min;
var mMax = Math.max;
var round = Math.round;

module.exports = exports;
