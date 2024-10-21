var isFn = require('./isFn');
var loadImg = require('./loadImg');
var noop = require('./noop');
var defaults = require('./defaults');
var createUrl = require('./createUrl');
var isStr = require('./isStr');
exports = function(file, options, cb) {
    if (isFn(options)) {
        cb = options;
        options = {};
    }
    cb = cb || noop;
    options = options || {};
    defaults(options, defOptions);
    options.mimeType = options.mimeType || file.type;
    if (isStr(file)) {
        options.isUrl = true;
    } else {
        file = createUrl(file);
    }
    loadImg(file, function(err, img) {
        if (err) return cb(err);
        compress(img, options, cb);
    });
};
function compress(img, options, cb) {
    var canvas = document.createElement('canvas');
    var ctx = canvas.getContext('2d');
    var width = img.width;
    var height = img.height;
    var ratio = width / height;
    var maxWidth = options.maxWidth;
    var maxHeight = options.maxHeight;
    if (options.width || options.height) {
        if (options.width) {
            width = options.width;
            height = width / ratio;
        } else if (options.height) {
            height = options.height;
            width = height * ratio;
        }
    } else {
        if (width > maxWidth) {
            width = maxWidth;
            height = width / ratio;
        }
        if (height > maxHeight) {
            height = maxHeight;
            width = height * ratio;
        }
    }
    width = floor(width);
    height = floor(height);
    canvas.width = width;
    canvas.height = height;
    ctx.drawImage(img, 0, 0, width, height);
    if (URL && options.isUrl) URL.revokeObjectURL(img.src);
    if (canvas.toBlob) {
        try {
            canvas.toBlob(
                function(file) {
                    cb(null, file);
                },
                options.mimeType,
                options.quality
            );
        } catch (e) {
            cb(e);
        }
    } else {
        cb(new Error('Canvas toBlob is not supported'));
    }
}
var defOptions = {
    maxWidth: Infinity,
    maxHeight: Infinity,
    quality: 0.8
};
var floor = Math.floor;

module.exports = exports;
