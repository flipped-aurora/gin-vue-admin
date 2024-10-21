var isFn = require('./isFn');
var has = require('./has');
var keys = require('./keys');
exports = function(a, b) {
    return eq(a, b);
};
function deepEq(a, b, aStack, bStack) {
    var className = toString.call(a);
    if (className !== toString.call(b)) return false;
    switch (className) {
        case '[object RegExp]':
        case '[object String]':
            return '' + a === '' + b;
        case '[object Number]':
            if (+a !== +a) return +b !== +b;
            return +a === 0 ? 1 / +a === 1 / b : +a === +b;
        case '[object Date]':
        case '[object Boolean]':
            return +a === +b;
    }
    var areArrays = className === '[object Array]';
    if (!areArrays) {
        if (typeof a != 'object' || typeof b != 'object') return false;
        var aCtor = a.constructor;
        var bCtor = b.constructor;
        if (
            aCtor !== bCtor &&
            !(
                isFn(aCtor) &&
                aCtor instanceof aCtor &&
                isFn(bCtor) &&
                bCtor instanceof bCtor
            ) &&
            'constructor' in a &&
            'constructor' in b
        )
            return false;
    }
    aStack = aStack || [];
    bStack = bStack || [];
    var length = aStack.length;
    while (length--) if (aStack[length] === a) return bStack[length] === b;
    aStack.push(a);
    bStack.push(b);
    if (areArrays) {
        length = a.length;
        if (length !== b.length) return false;
        while (length--)
            if (!eq(a[length], b[length], aStack, bStack)) return false;
    } else {
        var _keys = keys(a);
        var key;
        length = _keys.length;
        if (keys(b).length !== length) return false;
        while (length--) {
            key = _keys[length];
            if (!(has(b, key) && eq(a[key], b[key], aStack, bStack)))
                return false;
        }
    }
    aStack.pop();
    bStack.pop();
    return true;
}
function eq(a, b, aStack, bStack) {
    if (a === b) return a !== 0 || 1 / a === 1 / b;
    if (a == null || b == null) return a === b;
    if (a !== a) return b !== b;
    var type = typeof a;
    if (type !== 'function' && type !== 'object' && typeof b != 'object')
        return false;
    return deepEq(a, b, aStack, bStack);
}

module.exports = exports;
