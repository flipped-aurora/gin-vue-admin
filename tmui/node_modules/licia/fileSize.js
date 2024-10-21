exports = function(bytes) {
    if (bytes <= 0) return '0';
    var suffixIdx = Math.floor(Math.log(bytes) / Math.log(1024));
    var val = bytes / Math.pow(2, suffixIdx * 10);
    return +val.toFixed(2) + suffixList[suffixIdx];
};
var suffixList = ['', 'K', 'M', 'G', 'T'];

module.exports = exports;
