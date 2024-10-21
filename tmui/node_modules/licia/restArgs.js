exports = function(fn, startIdx) {
    startIdx = startIdx == null ? fn.length - 1 : +startIdx;
    return function() {
        var len = Math.max(arguments.length - startIdx, 0);
        var rest = new Array(len);
        var i;
        for (i = 0; i < len; i++) rest[i] = arguments[i + startIdx];

        switch (startIdx) {
            case 0:
                return fn.call(this, rest);
            case 1:
                return fn.call(this, arguments[0], rest);
            case 2:
                return fn.call(this, arguments[0], arguments[1], rest);
        }
        var args = new Array(startIdx + 1);
        for (i = 0; i < startIdx; i++) args[i] = arguments[i];
        args[startIdx] = rest;
        return fn.apply(this, args);
    };
};

module.exports = exports;
