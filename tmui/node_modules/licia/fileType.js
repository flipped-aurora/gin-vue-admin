function _slicedToArray(r, e) {
    return (
        _arrayWithHoles(r) ||
        _iterableToArrayLimit(r, e) ||
        _unsupportedIterableToArray(r, e) ||
        _nonIterableRest()
    );
}
function _nonIterableRest() {
    throw new TypeError(
        'Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.'
    );
}
function _unsupportedIterableToArray(r, a) {
    if (r) {
        if ('string' == typeof r) return _arrayLikeToArray(r, a);
        var t = {}.toString.call(r).slice(8, -1);
        return (
            'Object' === t && r.constructor && (t = r.constructor.name),
            'Map' === t || 'Set' === t
                ? Array.from(r)
                : 'Arguments' === t ||
                  /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t)
                ? _arrayLikeToArray(r, a)
                : void 0
        );
    }
}
function _arrayLikeToArray(r, a) {
    (null == a || a > r.length) && (a = r.length);
    for (var e = 0, n = Array(a); e < a; e++) n[e] = r[e];
    return n;
}
function _iterableToArrayLimit(r, l) {
    var t =
        null == r
            ? null
            : ('undefined' != typeof Symbol && r[Symbol.iterator]) ||
              r['@@iterator'];
    if (null != t) {
        var e,
            n,
            i,
            u,
            a = [],
            f = !0,
            o = !1;
        try {
            if (((i = (t = t.call(r)).next), 0 === l)) {
                if (Object(t) !== t) return;
                f = !1;
            } else
                for (
                    ;
                    !(f = (e = i.call(t)).done) &&
                    (a.push(e.value), a.length !== l);
                    f = !0
                );
        } catch (r) {
            (o = !0), (n = r);
        } finally {
            try {
                if (
                    !f &&
                    null != t.return &&
                    ((u = t.return()), Object(u) !== u)
                )
                    return;
            } finally {
                if (o) throw n;
            }
        }
        return a;
    }
}
function _arrayWithHoles(r) {
    if (Array.isArray(r)) return r;
}

var type = require('./type');
var mime = require('./mime');
var isFn = require('./isFn');
exports = function(input) {
    if (type(input) !== 'uint8array') {
        input = new Uint8Array(input);
    }
    for (var i = 0, len = types.length; i < len; i++) {
        var _type = types[i];
        var _type2 = _slicedToArray(_type, 3),
            ext = _type2[0],
            magic = _type2[1],
            offset = _type2[2];
        if (isFn(magic)) {
            if (magic(input)) {
                return {
                    ext: ext,
                    mime: mime(ext)
                };
            }
        } else if (check(input, magic, offset)) {
            return {
                ext: ext,
                mime: mime(ext)
            };
        }
    }
};
var types = [
    ['jpg', [0xff, 0xd8, 0xff]],
    ['png', [0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a]],
    ['gif', [0x47, 0x49, 0x46]],
    ['webp', [0x57, 0x45, 0x42, 0x50], 8],
    ['bmp', [0x42, 0x4d]],
    ['gz', [0x1f, 0x8b, 0x8]],
    [
        'zip',
        function(input) {
            return (
                check(input, [0x50, 0x4b]) &&
                (input[2] === 0x3 || input[2] === 0x5 || input[2] === 0x7) &&
                (input[3] === 0x4 || input[3] === 0x6 || input[3] === 0x8)
            );
        }
    ],
    [
        'rar',
        function(input) {
            return (
                check(input, [0x52, 0x61, 0x72, 0x21, 0x1a, 0x7]) &&
                (input[6] === 0x0 || input[6] === 0x1)
            );
        }
    ],
    ['pdf', [0x25, 0x50, 0x44, 0x46]],
    ['exe', [0x4d, 0x5a]]
];
function check(input, magic) {
    var offset =
        arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : 0;
    for (var i = 0, len = magic.length; i < len; i++) {
        if (input[offset + i] !== magic[i]) {
            return false;
        }
    }
    return true;
}

module.exports = exports;
