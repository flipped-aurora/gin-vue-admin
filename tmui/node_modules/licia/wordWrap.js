var map = require('./map');
var reduce = require('./reduce');
var concat = require('./concat');
var last = require('./last');
var trim = require('./trim');
exports = function(txt, width) {
    var lines = txt.split('\n');
    return map(lines, function(line) {
        return wrap(line, width);
    }).join('\n');
};
var regWordBoundary = /(\S+\s+)/;
function wrap(txt, width) {
    var chunks = reduce(
        txt.split(regWordBoundary),
        function(chunks, word) {
            if (trim(word) === '') return chunks;
            if (word.length > width) {
                chunks = concat(
                    chunks,
                    word.match(new RegExp('.{1,'.concat(width, '}'), 'g'))
                );
            } else {
                chunks.push(word);
            }
            return chunks;
        },
        []
    );
    var lines = reduce(
        chunks,
        function(lines, chunk) {
            var lastLine = last(lines);
            if (lastLine.length + chunk.length > width) {
                if (trim(lastLine) === '') {
                    lines.pop();
                }
                lines.push(chunk);
            } else {
                lines[lines.length - 1] = lastLine + chunk;
            }
            return lines;
        },
        [chunks.shift()]
    );
    return lines.join('\n');
}

module.exports = exports;
