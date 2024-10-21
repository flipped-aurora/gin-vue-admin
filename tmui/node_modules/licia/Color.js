var Class = require('./Class');
var isStr = require('./isStr');
var clamp = require('./clamp');
var rgbToHsl = require('./rgbToHsl');
var hslToRgb = require('./hslToRgb');
var hex = require('./hex');
exports = Class(
    {
        initialize: function Color(color) {
            if (isStr(color)) color = exports.parse(color);
            this.model = color.model;
            this.val = color.val;
        },
        toRgb: function() {
            var val = this.val;
            if (this.model === 'hsl') val = hslToRgb(val);
            var prefix = 'rgba';
            if (val[3] === 1) {
                prefix = 'rgb';
                val = val.slice(0, 3);
            }
            return prefix + '(' + val.join(', ') + ')';
        },
        toHex: function() {
            var val = this.val;
            if (this.model === 'hsl') val = hslToRgb(val);
            var ret = hex.encode(val.slice(0, 3));
            if (ret[0] === ret[1] && ret[2] === ret[3] && ret[4] === ret[5]) {
                ret = ret[0] + ret[2] + ret[5];
            }
            return '#' + ret;
        },
        toHsl: function() {
            var val = this.val;
            if (this.model === 'rgb') val = rgbToHsl(val);
            var prefix = 'hsla';
            if (val[3] === 1) {
                prefix = 'hsl';
                val = val.slice(0, 3);
            }
            val[1] = val[1] + '%';
            val[2] = val[2] + '%';
            return prefix + '(' + val.join(', ') + ')';
        }
    },
    {
        parse: function(colorStr) {
            var i, match;
            var val = [0, 0, 0, 1],
                model = 'rgb';

            if ((match = colorStr.match(regHexAbbr))) {
                match = match[1];
                for (i = 0; i < 3; i++) {
                    val[i] = parseInt(match[i] + match[i], 16);
                }
            } else if ((match = colorStr.match(regHex))) {
                match = match[1];
                for (i = 0; i < 3; i++) {
                    var i2 = i * 2;
                    val[i] = parseInt(match.slice(i2, i2 + 2), 16);
                }
            } else if ((match = colorStr.match(regRgba))) {
                for (i = 0; i < 3; i++) {
                    val[i] = parseInt(match[i + 1], 0);
                }
                if (match[4]) val[3] = parseFloat(match[4]);
            } else if ((match = colorStr.match(regRgbaPer))) {
                for (i = 0; i < 3; i++) {
                    val[i] = Math.round(parseFloat(match[i + 1]) * 2.55);
                }
                if (match[4]) val[3] = parseFloat(match[4]);
            } else if ((match = colorStr.match(regHsla))) {
                model = 'hsl';
                val = [
                    ((parseFloat(match[1]) % 360) + 360) % 360,
                    clamp(parseFloat(match[2]), 0, 100),
                    clamp(parseFloat(match[3]), 0, 100),
                    clamp(parseFloat(match[4]), 0, 1)
                ];
            }
            return {
                val: val,
                model: model
            };
        }
    }
);
var regHexAbbr = /^#([a-fA-F0-9]{3})$/;
var regHex = /^#([a-fA-F0-9]{6})$/;
var regRgba = /^rgba?\(\s*([+-]?\d+)\s*,\s*([+-]?\d+)\s*,\s*([+-]?\d+)\s*(?:,\s*([+-]?[\d.]+)\s*)?\)$/;
var regRgbaPer = /^rgba?\(\s*([+-]?[\d.]+)%\s*,\s*([+-]?[\d.]+)%\s*,\s*([+-]?[\d.]+)%\s*(?:,\s*([+-]?[\d.]+)\s*)?\)$/;
var regHsla = /^hsla?\(\s*([+-]?\d*[.]?\d+)(?:deg)?\s*,\s*([+-]?[\d.]+)%\s*,\s*([+-]?[\d.]+)%\s*(?:,\s*([+-]?[\d.]+)\s*)?\)$/;

module.exports = exports;
