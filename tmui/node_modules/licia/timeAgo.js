var isDate = require('./isDate');
var toInt = require('./toInt');
exports = function(date, now) {
    if (!isDate(date)) date = new Date(date);
    now = now || new Date();
    if (!isDate(now)) now = new Date(now);
    var diff = (now - date) / 1000;
    var i = 0;
    var ago = diff > 0;
    diff = Math.abs(diff);
    while (diff >= secArr[i] && i < secArrLen) {
        diff /= secArr[i];
        i++;
    }
    diff = toInt(diff);
    i *= 2;
    if (diff > (i === 0 ? 9 : 1)) i += 1;
    return format(diff, i, ago);
};
var secArr = [60, 60, 24, 7, 365 / 7 / 12, 12];
var secArrLen = secArr.length;
function format(diff, i, ago) {
    return exports.i18n[i][ago ? 0 : 1].replace('%s', diff);
}
exports.i18n = [
    ['just now', 'right now'],
    ['%s seconds ago', 'in %s seconds'],
    ['1 minute ago', 'in 1 minute'],
    ['%s minutes ago', 'in %s minutes'],
    ['1 hour ago', 'in 1 hour'],
    ['%s hours ago', 'in %s hours'],
    ['1 day ago', 'in 1 day'],
    ['%s days ago', 'in %s days'],
    ['1 week ago', 'in 1 week'],
    ['%s weeks ago', 'in %s weeks'],
    ['1 month ago', 'in 1 month'],
    ['%s months ago', 'in %s months'],
    ['1 year ago', 'in 1 year'],
    ['%s years ago', 'in %s years']
];

module.exports = exports;
