var defaults = require('./defaults');
var random = require('./random');
var Color = require('./Color');
var seedRandom = require('./seedRandom');
var isFn = require('./isFn');
exports = function() {
    var options =
        arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {};
    defaults(options, defOpts);
    var count = options.count;
    var randomH = options.randomH,
        randomL = options.randomL,
        randomS = options.randomS;
    if (!isFn(randomH)) {
        var seed = options.seed || random(0, 100000);
        randomH = seedRandom(seed, 0, 360, false);
        randomL = seedRandom(seed + 1, 0, 1);
        randomS = seedRandom(seed + 2, 0, 1);
    }
    if (count > 1) {
        var colors = [];
        for (var i = 0; i < count; i++) {
            colors.push(
                exports(
                    defaults(
                        {
                            count: 1,
                            randomH: randomH,
                            randomL: randomL,
                            randomS: randomS
                        },
                        options
                    )
                )
            );
        }
        return colors;
    }
    var hue = options.hue || randomH();
    var lightness = options.lightness || randomL().toFixed(2);
    var saturation = options.saturation || randomS().toFixed(2);
    var color = new Color({
        val: [hue, Math.round(saturation * 100), Math.round(lightness * 100)],
        model: 'hsl'
    });
    switch (options.format) {
        case 'hsl':
            return color.toHsl();
        case 'rgb':
            return color.toRgb();
        default:
            return color.toHex();
    }
};
var defOpts = {
    count: 1,
    format: 'hex'
};

module.exports = exports;
