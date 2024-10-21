const reduce = require('./reduce');

exports = function(...streams) {
    reduce(streams, (from, to) => from.pipe(to));
};

module.exports = exports;
