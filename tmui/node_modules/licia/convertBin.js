var isStr = require('./isStr');
var base64 = require('./base64');
var isArrBuffer = require('./isArrBuffer');
var isArr = require('./isArr');
var isBuffer = require('./isBuffer');
var type = require('./type');
var lowerCase = require('./lowerCase');
exports = function(bin, t) {
    var result;
    t = lowerCase(t);
    if (isStr(bin)) {
        result = new Uint8Array(base64.decode(bin));
    } else if (isArrBuffer(bin)) {
        bin = bin.slice(0);
        result = new Uint8Array(bin);
    } else if (isArr(bin)) {
        result = new Uint8Array(bin);
    } else if (type(bin) === 'uint8array') {
        result = bin.slice(0);
    } else if (isBuffer(bin)) {
        result = new Uint8Array(bin.length);
        for (var i = 0; i < bin.length; i++) {
            result[i] = bin[i];
        }
    }
    if (result) {
        switch (t) {
            case 'base64':
                result = base64.encode(result);
                break;
            case 'arraybuffer':
                result = result.buffer;
                break;
            case 'array':
                result = [].slice.call(result);
                break;
            case 'buffer':
                result = Buffer.from(result);
                break;
            case 'blob':
                result = new Blob([result.buffer]);
                break;
        }
    }
    return result;
};
exports.blobToArrBuffer = function(blob) {
    return new Promise(function(resolve, reject) {
        var fileReader = new FileReader();
        fileReader.onload = function(e) {
            resolve(e.target.result);
        };
        fileReader.onerror = function(err) {
            reject(err);
        };
        fileReader.readAsArrayBuffer(blob);
    });
};

module.exports = exports;
