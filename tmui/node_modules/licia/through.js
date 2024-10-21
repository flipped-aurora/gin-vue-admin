const isFn = require('./isFn');
const extend = require('./extend');
const inherits = require('./inherits');

const Transform = require('stream').Transform;

exports = through(function(opts, transform, flush) {
    const t = new Transform(opts);

    t._transform = transform;
    if (flush) t._flush = flush;

    return t;
});

exports.obj = through(function(opts, transform, flush) {
    const t = new Transform(
        extend(
            {
                objectMode: true,
                highWaterMark: 16
            },
            opts
        )
    );

    t._transform = transform;
    if (flush) t._flush = flush;

    return t;
});

exports.ctor = through(function(opts, transform, flush) {
    function Through(override) {
        if (!(this instanceof Through)) return new Through(override);

        Transform.call(this, extend(opts, override));
    }

    inherits(Through, Transform);

    const proto = Through.prototype;
    proto._transform = transform;
    if (flush) proto._flush = flush;

    return Through;
});

function through(streamFactory) {
    return function(opts, transform, flush) {
        if (isFn(opts)) {
            flush = transform;
            transform = opts;
            opts = {};
        }

        if (!isFn(transform)) transform = defTransform;
        if (!isFn(flush)) flush = null;

        return streamFactory(opts, transform, flush);
    };
}

function defTransform(chunk, enc, cb) {
    cb(null, chunk);
}

module.exports = exports;
