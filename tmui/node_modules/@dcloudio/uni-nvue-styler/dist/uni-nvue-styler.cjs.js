'use strict';

var parseCSSFont = require('parse-css-font');
var postcss = require('postcss');

function _interopDefault (e) { return e && e.__esModule ? e : { default: e }; }

var parseCSSFont__default = /*#__PURE__*/_interopDefault(parseCSSFont);
var postcss__default = /*#__PURE__*/_interopDefault(postcss);

const COMBINATORS_RE = /^((?:(?:\.[A-Za-z0-9_\-]+)+[\+\~\> ])*)((?:\.[A-Za-z0-9_\-\:]+)+)$/;
function createDecl(prop, value, important, raws, source) {
    const decl = {
        type: 'decl',
        prop,
        value: value.toString(),
        raws,
        source,
    };
    if (important) {
        decl.important = true;
    }
    return decl;
}
const NUM_REGEXP = /^[-]?\d*\.?\d+$/;
const LENGTH_REGEXP = /^[-+]?\d*\.?\d+(\S*)$/;
const SUPPORTED_VALUES_REGEXP = /supported values are: ([^)]+)/;
const SUPPORT_CSS_UNIT = ['px', 'pt', 'wx', 'upx', 'rpx'];
const isNumber = (val) => typeof val === 'number';
const cacheStringFunction$1 = (fn) => {
    const cache = Object.create(null);
    return ((str) => {
        const hit = cache[str];
        return hit || (cache[str] = fn(str));
    });
};
const hyphenateRE$1 = /([A-Z])/g;
const hyphenateStyleProperty = cacheStringFunction$1((str) => str
    .replace(hyphenateRE$1, (_, m) => {
    if (typeof m === 'string') {
        return '-' + m.toLowerCase();
    }
    return m;
})
    .toLowerCase());
function autofixedReason(v, result) {
    return 'NOTE: property value `' + v + '` is autofixed to `' + result + '`';
}
function validReason(k, v) {
    return ('ERROR: property value `' +
        v +
        '` is not valid for `' +
        hyphenateStyleProperty(k) +
        '`');
}
function defaultValueReason(k, v) {
    return ('NOTE: property value `' +
        v +
        '` is the DEFAULT value for `' +
        hyphenateStyleProperty(k) +
        '` (could be removed)');
}
function supportedEnumReason(k, v, items) {
    return ('ERROR: property value `' +
        v +
        '` is not supported for `' +
        hyphenateStyleProperty(k) +
        '` (supported values are: `' +
        items.join('`|`') +
        '`)');
}
function supportedValueWithTipsReason(k, v, tips) {
    return ('ERROR: property value `' +
        v +
        '` is not supported for `' +
        hyphenateStyleProperty(k) +
        '` ' +
        tips);
}
function supportedUnitWithAutofixedReason(unit, v, result) {
    return ('NOTE: unit `' +
        unit +
        '` is not supported and property value `' +
        v +
        '` is autofixed to `' +
        result +
        '`');
}
function compatibilityReason(k) {
    return ('NOTE: the ' +
        hyphenateStyleProperty(k) +
        ' property may have compatibility problem on native');
}
function supportedPropertyReason(k) {
    return ('WARNING: `' +
        hyphenateStyleProperty(k) +
        '` is not a standard property name (may not be supported)');
}
function getSupportedPlatforms(uniPlatform) {
    const supportedPlatforms = [];
    if (uniPlatform?.app?.android?.unixVer !== 'x') {
        supportedPlatforms.push('app-android');
    }
    if (uniPlatform?.app.ios?.unixVer !== 'x') {
        supportedPlatforms.push('app-ios');
    }
    return supportedPlatforms;
}
function normalizeReasons(reasons, k, v) {
    let enums = [];
    for (let i = 0; i < reasons.length; i++) {
        const reason = reasons[i];
        if (SUPPORTED_VALUES_REGEXP.test(reason)) {
            const match = reason.match(SUPPORTED_VALUES_REGEXP);
            if (match) {
                const values = match[1]
                    .split('|')
                    .map((item) => item.replace(/`/g, ''))
                    .filter(Boolean);
                enums.push(...values);
                reasons.splice(i, 1);
                i--;
            }
        }
    }
    if (enums.length > 0) {
        enums = [...new Set(enums)];
        reasons.push(supportedEnumReason(k, v, enums));
    }
    return reasons;
}

const backgroundColor = 'background-color' ;
const backgroundImage = 'background-image' ;
const transformBackground = (decl) => {
    const { value, important, raws, source } = decl;
    if (/^#?\S+$/.test(value) || /^rgba?(.+)$/.test(value)) {
        return [createDecl(backgroundColor, value, important, raws, source)];
    }
    else if (/^linear-gradient(.+)$/.test(value)) {
        return [createDecl(backgroundImage, value, important, raws, source)];
    }
    return [decl];
};

const borderWidth = '-width' ;
const borderStyle = '-style' ;
const borderColor = '-color' ;
function createTransformBorder(options) {
    return (decl) => {
        const { prop, value, important, raws, source } = decl;
        const splitResult = value.replace(/\s*,\s*/g, ',').split(/\s+/);
        const result = [
            /^[\d\.]+\S*|^(thin|medium|thick)$/,
            /^(solid|dashed|dotted|none)$/,
            /\S+/,
        ].map((item) => {
            const index = splitResult.findIndex((str) => item.test(str));
            return index < 0 ? null : splitResult.splice(index, 1)[0];
        });
        if (splitResult.length) {
            return [decl];
        }
        return [
            createDecl(prop + borderWidth, (result[0] || (options.type === 'uvue' ? 'medium' : '0')).trim(), important, raws, source),
            createDecl(prop + borderStyle, (result[1] || (options.type === 'uvue' ? 'none' : 'solid')).trim(), important, raws, source),
            createDecl(prop + borderColor, (result[2] || '#000000').trim(), important, raws, source),
        ];
    };
}

/**
* @vue/shared v3.4.21
* (c) 2018-present Yuxi (Evan) You and Vue contributors
* @license MIT
**/

!!(process.env.NODE_ENV !== "production") ? Object.freeze({}) : {};
!!(process.env.NODE_ENV !== "production") ? Object.freeze([]) : [];
const extend = Object.assign;
const hasOwnProperty = Object.prototype.hasOwnProperty;
const hasOwn = (val, key) => hasOwnProperty.call(val, key);
const isFunction = (val) => typeof val === "function";
const isString = (val) => typeof val === "string";
const cacheStringFunction = (fn) => {
  const cache = /* @__PURE__ */ Object.create(null);
  return (str) => {
    const hit = cache[str];
    return hit || (cache[str] = fn(str));
  };
};
const camelizeRE = /-(\w)/g;
const camelize = cacheStringFunction((str) => {
  return str.replace(camelizeRE, (_, c) => c ? c.toUpperCase() : "");
});
const hyphenateRE = /\B([A-Z])/g;
const hyphenate = cacheStringFunction(
  (str) => str.replace(hyphenateRE, "-$1").toLowerCase()
);

const borderTop = 'border-top-' ;
const borderRight = 'border-right-' ;
const borderBottom = 'border-bottom-' ;
const borderLeft = 'border-left-' ;
const transformBorderColor = (decl) => {
    const { prop, value, important, raws, source } = decl;
    let property = hyphenate(prop).split('-')[1];
    const splitResult = value.replace(/\s*,\s*/g, ',').split(/\s+/);
    switch (splitResult.length) {
        case 1:
            return [decl];
        case 2:
            splitResult.push(splitResult[0], splitResult[1]);
            break;
        case 3:
            splitResult.push(splitResult[1]);
            break;
    }
    return [
        createDecl(borderTop + property, splitResult[0], important, raws, source),
        createDecl(borderRight + property, splitResult[1], important, raws, source),
        createDecl(borderBottom + property, splitResult[2], important, raws, source),
        createDecl(borderLeft + property, splitResult[3], important, raws, source),
    ];
};

const borderTopLeftRadius = 'border-top-left-radius'
    ;
const borderTopRightRadius = 'border-top-right-radius'
    ;
const borderBottomRightRadius = 'border-bottom-right-radius'
    ;
const borderBottomLeftRadius = 'border-bottom-left-radius'
    ;
const transformBorderRadius = (decl) => {
    const { value, important, raws, source } = decl;
    const splitResult = value.split(/\s+/);
    if (value.includes('/')) {
        return [decl];
    }
    switch (splitResult.length) {
        case 1:
            return [decl];
        case 2:
            splitResult.push(splitResult[0], splitResult[1]);
            break;
        case 3:
            splitResult.push(splitResult[1]);
            break;
    }
    return [
        createDecl(borderTopLeftRadius, splitResult[0], important, raws, source),
        createDecl(borderTopRightRadius, splitResult[1], important, raws, source),
        createDecl(borderBottomRightRadius, splitResult[2], important, raws, source),
        createDecl(borderBottomLeftRadius, splitResult[3], important, raws, source),
    ];
};

const transformBorderStyle = transformBorderColor;

const transformBorderWidth = transformBorderColor;

const flexDirection = 'flex-direction' ;
const flexWrap = 'flex-wrap' ;
const transformFlexFlow = (decl) => {
    const { value, important, raws, source } = decl;
    const splitResult = value.split(/\s+/);
    const result = [
        /^(column|column-reverse|row|row-reverse)$/,
        /^(nowrap|wrap|wrap-reverse)$/,
    ].map((item) => {
        const index = splitResult.findIndex((str) => item.test(str));
        return index < 0 ? null : splitResult.splice(index, 1)[0];
    });
    if (splitResult.length) {
        return [decl];
    }
    return [
        createDecl(flexDirection, result[0] || 'column', important, raws, source),
        createDecl(flexWrap, result[1] || 'nowrap', important, raws, source),
    ];
};

const transformFont = (decl) => {
    const { value, important, raws, source } = decl;
    const result = [];
    const font = parseCSSFont__default.default(value);
    if (font.system) {
        return result;
    }
    const { style, weight, size, lineHeight, family } = font;
    if (style) {
        result.push(createDecl('font-style', style, important, raws, source));
    }
    if (weight) {
        result.push(createDecl('font-weight', weight, important, raws, source));
    }
    if (size) {
        result.push(createDecl('font-size', size, important, raws, source));
    }
    if (lineHeight) {
        result.push(createDecl('line-height', lineHeight, important, raws, source));
    }
    if (family) {
        result.push(createDecl('font-family', serialize(family), important, raws, source));
    }
    return result;
};
function serialize(family) {
    return family.map((f) => (f.includes(' ') ? `"${f}"` : f)).join(', ');
}

const top = '-top' ;
const right = '-right' ;
const bottom = '-bottom' ;
const left = '-left' ;
const createTransformBox = (type) => {
    return (decl) => {
        const { value, important, raws, source } = decl;
        const splitResult = value.split(/\s+/);
        switch (splitResult.length) {
            case 1:
                splitResult.push(splitResult[0], splitResult[0], splitResult[0]);
                break;
            case 2:
                splitResult.push(splitResult[0], splitResult[1]);
                break;
            case 3:
                splitResult.push(splitResult[1]);
                break;
        }
        return [
            createDecl(type + top, splitResult[0], important, raws, source),
            createDecl(type + right, splitResult[1], important, raws, source),
            createDecl(type + bottom, splitResult[2], important, raws, source),
            createDecl(type + left, splitResult[3], important, raws, source),
        ];
    };
};
const transformMargin = createTransformBox('margin');

const transformPadding = createTransformBox('padding');

const transitionProperty = 'transition-property'
    ;
const transitionDuration = 'transition-duration'
    ;
const transitionTimingFunction = 'transition-timing-function'
    ;
const transitionDelay = 'transition-delay' ;
const transformTransition = (decl) => {
    const { value, important, raws, source } = decl;
    const result = [];
    let match;
    // 针对 cubic-bezier 特殊处理
    // eg: cubic-bezier(0.42, 0, 1.0, 3) // (0.2,-2,0.8,2)
    if (decl.value.includes('cubic-bezier')) {
        const CHUNK_REGEXP = /^(\S*)?\s*(\d*\.?\d+(?:ms|s)?)?\s*((\S*)|cubic-bezier\(.*\))?\s*(\d*\.?\d+(?:ms|s)?)?$/;
        match = value.match(CHUNK_REGEXP);
    }
    else {
        const CHUNK_REGEXP = /^(\S*)?\s*(\d*\.?\d+(?:ms|s)?)?\s*(\S*)?\s*(\d*\.?\d+(?:ms|s)?)?$/;
        match = value.match(CHUNK_REGEXP);
    }
    if (!match) {
        return result;
    }
    match[1] &&
        result.push(createDecl(transitionProperty, match[1], important, raws, source));
    match[2] &&
        result.push(createDecl(transitionDuration, match[2], important, raws, source));
    match[3] &&
        result.push(createDecl(transitionTimingFunction, match[3], important, raws, source));
    match[4] &&
        result.push(createDecl(transitionDelay, match[4], important, raws, source));
    return result;
};

function getDeclTransforms(options) {
    const transformBorder = createTransformBorder(options);
    const styleMap = {
        transition: transformTransition,
        border: transformBorder,
        background: transformBackground,
        borderTop: transformBorder,
        borderRight: transformBorder,
        borderBottom: transformBorder,
        borderLeft: transformBorder,
        borderStyle: transformBorderStyle,
        borderWidth: transformBorderWidth,
        borderColor: transformBorderColor,
        borderRadius: transformBorderRadius,
        // uvue已经支持这些简写属性，不需要展开
        // margin,padding继续展开，确保样式的优先级
        margin: transformMargin,
        padding: transformPadding,
        /* eslint-disable no-restricted-syntax */
        ...(options.type !== 'uvue'
            ? {
                flexFlow: transformFlexFlow,
            }
            : {}),
    };
    let result = {};
    {
        styleMap.font = transformFont;
        for (const property in styleMap) {
            result[hyphenateStyleProperty(property)] = styleMap[property];
        }
    }
    return result;
}
let DeclTransforms;
const expanded = Symbol('expanded');
function expand(options) {
    const plugin = {
        postcssPlugin: 'nvue:expand',
        Declaration(decl) {
            if (decl[expanded]) {
                return;
            }
            if (!DeclTransforms) {
                DeclTransforms = getDeclTransforms(options);
            }
            const transform = DeclTransforms[decl.prop];
            if (transform) {
                const res = transform(decl);
                const isSame = res.length === 1 && res[0] === decl;
                if (!isSame) {
                    decl.replaceWith(res);
                }
            }
            decl[expanded] = true;
        },
    };
    return plugin;
}

const normalizeColor = (v) => {
    v = (v || '').toString();
    if (v.match(/^#[0-9a-fA-F]{6}$/)) {
        return { value: v };
    }
    if (v.match(/^#[0-9a-fA-F]{3}$/)) {
        return {
            value: '#' + v[1] + v[1] + v[2] + v[2] + v[3] + v[3],
            reason: function reason(k, v, result) {
                return autofixedReason(v, result);
            },
        };
    }
    if (EXTENDED_COLOR_KEYWORDS[v]) {
        return {
            value: EXTENDED_COLOR_KEYWORDS[v],
            reason: function reason(k, v, result) {
                return autofixedReason(v, result);
            },
        };
    }
    let arrColor, r, g, b, a;
    const RGB_REGEXP = /^rgb\(\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*\)$/gi;
    const RGBA_REGEXP = /^rgba\(\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d*\.?\d+)\s*\)$/gi;
    if ((arrColor = RGB_REGEXP.exec(v))) {
        r = parseInt(arrColor[1]);
        g = parseInt(arrColor[2]);
        b = parseInt(arrColor[3]);
        if (r >= 0 && r <= 255 && g >= 0 && g <= 255 && b >= 0 && b <= 255) {
            return { value: 'rgb(' + [r, g, b].join(',') + ')' };
        }
    }
    if ((arrColor = RGBA_REGEXP.exec(v))) {
        r = parseInt(arrColor[1]);
        g = parseInt(arrColor[2]);
        b = parseInt(arrColor[3]);
        a = parseFloat(arrColor[4]);
        if (r >= 0 &&
            r <= 255 &&
            g >= 0 &&
            g <= 255 &&
            b >= 0 &&
            b <= 255 &&
            a >= 0 &&
            a <= 1) {
            return { value: 'rgba(' + [r, g, b, a].join(',') + ')' };
        }
    }
    if (v === 'transparent') {
        return { value: 'rgba(0,0,0,0)' };
    }
    return {
        value: null,
        reason(k, v, result) {
            return validReason(k, v);
        },
    };
};
// http://www.w3.org/TR/css3-color/#svg-color
const EXTENDED_COLOR_KEYWORDS = {
    aliceblue: '#F0F8FF',
    antiquewhite: '#FAEBD7',
    aqua: '#00FFFF',
    aquamarine: '#7FFFD4',
    azure: '#F0FFFF',
    beige: '#F5F5DC',
    bisque: '#FFE4C4',
    black: '#000000',
    blanchedalmond: '#FFEBCD',
    blue: '#0000FF',
    blueviolet: '#8A2BE2',
    brown: '#A52A2A',
    burlywood: '#DEB887',
    cadetblue: '#5F9EA0',
    chartreuse: '#7FFF00',
    chocolate: '#D2691E',
    coral: '#FF7F50',
    cornflowerblue: '#6495ED',
    cornsilk: '#FFF8DC',
    crimson: '#DC143C',
    cyan: '#00FFFF',
    darkblue: '#00008B',
    darkcyan: '#008B8B',
    darkgoldenrod: '#B8860B',
    darkgray: '#A9A9A9',
    darkgreen: '#006400',
    darkgrey: '#A9A9A9',
    darkkhaki: '#BDB76B',
    darkmagenta: '#8B008B',
    darkolivegreen: '#556B2F',
    darkorange: '#FF8C00',
    darkorchid: '#9932CC',
    darkred: '#8B0000',
    darksalmon: '#E9967A',
    darkseagreen: '#8FBC8F',
    darkslateblue: '#483D8B',
    darkslategray: '#2F4F4F',
    darkslategrey: '#2F4F4F',
    darkturquoise: '#00CED1',
    darkviolet: '#9400D3',
    deeppink: '#FF1493',
    deepskyblue: '#00BFFF',
    dimgray: '#696969',
    dimgrey: '#696969',
    dodgerblue: '#1E90FF',
    firebrick: '#B22222',
    floralwhite: '#FFFAF0',
    forestgreen: '#228B22',
    fuchsia: '#FF00FF',
    gainsboro: '#DCDCDC',
    ghostwhite: '#F8F8FF',
    gold: '#FFD700',
    goldenrod: '#DAA520',
    gray: '#808080',
    green: '#008000',
    greenyellow: '#ADFF2F',
    grey: '#808080',
    honeydew: '#F0FFF0',
    hotpink: '#FF69B4',
    indianred: '#CD5C5C',
    indigo: '#4B0082',
    ivory: '#FFFFF0',
    khaki: '#F0E68C',
    lavender: '#E6E6FA',
    lavenderblush: '#FFF0F5',
    lawngreen: '#7CFC00',
    lemonchiffon: '#FFFACD',
    lightblue: '#ADD8E6',
    lightcoral: '#F08080',
    lightcyan: '#E0FFFF',
    lightgoldenrodyellow: '#FAFAD2',
    lightgray: '#D3D3D3',
    lightgreen: '#90EE90',
    lightgrey: '#D3D3D3',
    lightpink: '#FFB6C1',
    lightsalmon: '#FFA07A',
    lightseagreen: '#20B2AA',
    lightskyblue: '#87CEFA',
    lightslategray: '#778899',
    lightslategrey: '#778899',
    lightsteelblue: '#B0C4DE',
    lightyellow: '#FFFFE0',
    lime: '#00FF00',
    limegreen: '#32CD32',
    linen: '#FAF0E6',
    magenta: '#FF00FF',
    maroon: '#800000',
    mediumaquamarine: '#66CDAA',
    mediumblue: '#0000CD',
    mediumorchid: '#BA55D3',
    mediumpurple: '#9370DB',
    mediumseagreen: '#3CB371',
    mediumslateblue: '#7B68EE',
    mediumspringgreen: '#00FA9A',
    mediumturquoise: '#48D1CC',
    mediumvioletred: '#C71585',
    midnightblue: '#191970',
    mintcream: '#F5FFFA',
    mistyrose: '#FFE4E1',
    moccasin: '#FFE4B5',
    navajowhite: '#FFDEAD',
    navy: '#000080',
    oldlace: '#FDF5E6',
    olive: '#808000',
    olivedrab: '#6B8E23',
    orange: '#FFA500',
    orangered: '#FF4500',
    orchid: '#DA70D6',
    palegoldenrod: '#EEE8AA',
    palegreen: '#98FB98',
    paleturquoise: '#AFEEEE',
    palevioletred: '#DB7093',
    papayawhip: '#FFEFD5',
    peachpuff: '#FFDAB9',
    peru: '#CD853F',
    pink: '#FFC0CB',
    plum: '#DDA0DD',
    powderblue: '#B0E0E6',
    purple: '#800080',
    red: '#FF0000',
    rosybrown: '#BC8F8F',
    royalblue: '#4169E1',
    saddlebrown: '#8B4513',
    salmon: '#FA8072',
    sandybrown: '#F4A460',
    seagreen: '#2E8B57',
    seashell: '#FFF5EE',
    sienna: '#A0522D',
    silver: '#C0C0C0',
    skyblue: '#87CEEB',
    slateblue: '#6A5ACD',
    slategray: '#708090',
    slategrey: '#708090',
    snow: '#FFFAFA',
    springgreen: '#00FF7F',
    steelblue: '#4682B4',
    tan: '#D2B48C',
    teal: '#008080',
    thistle: '#D8BFD8',
    tomato: '#FF6347',
    turquoise: '#40E0D0',
    violet: '#EE82EE',
    wheat: '#F5DEB3',
    white: '#FFFFFF',
    whitesmoke: '#F5F5F5',
    yellow: '#FFFF00',
    yellowgreen: '#9ACD32',
};

function createEnumNormalize(items) {
    return (v) => {
        const index = items.indexOf(v);
        if (index > 0) {
            return { value: v };
        }
        if (index === 0) {
            return {
                value: v,
                reason: function reason(k, v, result) {
                    return defaultValueReason(k, v);
                },
            };
        }
        return {
            value: null,
            reason: function reason(k, v, result) {
                return supportedEnumReason(k, v, items);
            },
        };
    };
}
function createEnumNormalizeWithPlatform(items) {
    return (v, { platform }) => {
        const property = items.find((item) => item.name === v);
        const supportedEnum = items
            .filter((item) => {
            const supportedPlatforms = getSupportedPlatforms(item.uniPlatform);
            return supportedPlatforms.includes(platform);
        })
            .map((item) => item.name);
        if (property) {
            const supportedPlatforms = getSupportedPlatforms(property.uniPlatform);
            // TODO 未跨平台支持的属性特殊提示
            if (!supportedPlatforms.includes(platform)) {
                return {
                    value: null,
                    reason: function reason(k, v, result) {
                        return supportedEnumReason(k, v, supportedEnum);
                    },
                };
            }
            return { value: v };
        }
        return {
            value: null,
            reason: function reason(k, v, result) {
                return supportedEnumReason(k, v, supportedEnum);
            },
        };
    };
}

const normalizeFlexWrap = (v) => {
    const values = ['nowrap', 'wrap', 'wrap-reverse'];
    const index = values.indexOf(v);
    if (index > 0) {
        return {
            value: v,
            reason(k, v, result) {
                return compatibilityReason(k);
            },
        };
    }
    if (index === 0) {
        return {
            value: v,
            reason: function reason(k, v, result) {
                return defaultValueReason(k, v);
            },
        };
    }
    return {
        value: null,
        reason(k, v, result) {
            return supportedEnumReason(k, v, values);
        },
    };
};

const normalizeInteger = (v) => {
    v = (v || '').toString();
    if (v.match(/^[-+]?\d+$/)) {
        return { value: parseInt(v, 10) };
    }
    return {
        value: null,
        reason: function reason(k, v, result) {
            return supportedEnumReason(k, v, ['integer']);
        },
    };
};

function createNormalizeLength({ removePx, property, } = {}) {
    return (v, options) => {
        v = (v || '').toString();
        // css 变量
        if (v.includes('CSS_VAR_')) {
            return { value: v };
        }
        const match = v.match(LENGTH_REGEXP);
        if (match) {
            var unit = match[1];
            const uvue = options.type === 'uvue';
            if (uvue) {
                if (!unit || (unit === 'px' && removePx)) {
                    return { value: parseFloat(v) };
                }
                else if (unit === 'px' ||
                    unit === 'rpx' ||
                    // 只有line-height支持em单位
                    (unit === 'em' && property === 'line-height')) {
                    return { value: v };
                }
            }
            else {
                // nvue
                if (!unit || unit === 'px') {
                    return { value: parseFloat(v) };
                }
                if (SUPPORT_CSS_UNIT.includes(unit)) {
                    return { value: v };
                }
                else {
                    return {
                        value: parseFloat(v),
                        reason(k, v, result) {
                            return supportedUnitWithAutofixedReason(unit, v, result);
                        },
                    };
                }
            }
        }
        return {
            value: null,
            reason(k, v, result) {
                return supportedEnumReason(k, v, ['number', 'pixel']);
            },
        };
    };
}
const normalizeLength = createNormalizeLength({
    removePx: true,
});
const normalizeLengthWithOptions = createNormalizeLength;
const normalizePercent = (v) => {
    v = (v || '').toString();
    const match = v.match(LENGTH_REGEXP);
    if (match) {
        var unit = match[1];
        if (unit === '%') {
            return { value: v };
        }
    }
    return {
        value: null,
        reason(k, v, result) {
            return supportedEnumReason(k, v, ['percent']);
        },
    };
};

const normalizeNumber = (v) => {
    v = (v || '').toString();
    var match = v.match(LENGTH_REGEXP);
    if (match && !match[1]) {
        return { value: parseFloat(v) };
    }
    return {
        value: null,
        reason: function reason(k, v, result) {
            return supportedEnumReason(k, v, ['number']);
        },
    };
};

const normalizeString = (v) => {
    v = (v || '').toString().replace(/["']/g, '');
    return {
        value: v,
    };
};

const normalizeShorthandLength = (v, options) => {
    v = (v || '').toString();
    let value = [];
    let reason = [];
    const results = v.split(/\s+/).map((v) => normalizeLength(v, options));
    for (let i = 0; i < results.length; ++i) {
        const res = results[i];
        if (res.value === null) {
            return res;
        }
        value.push(res.value);
        reason.push(res.reason);
    }
    return {
        value: value.join(' '),
        reason: function (k, v, result) {
            return reason
                .map(function (res) {
                if (isFunction(res)) {
                    return res(k, v, result);
                }
            })
                .join('\n');
        },
    };
};

const normalizeTransform = (v) => {
    return { value: v };
};

const normalizeInterval = (v, options) => {
    v = (v || 0).toString();
    let match, num;
    if ((match = v.match(/^\d*\.?\d+(ms|s)?$/))) {
        const uvue = options.type === 'uvue';
        if (uvue) {
            // uvue 需要单位
            if (match[1]) {
                return { value: v };
            }
        }
        else {
            num = parseFloat(match[0]);
            if (!match[1]) {
                return { value: parseInt(num + '') };
            }
            if (match[1] === 's') {
                num *= 1000;
            }
            return {
                value: parseInt(num + ''),
                reason(k, v, result) {
                    return autofixedReason(v, result);
                },
            };
        }
    }
    return {
        value: null,
        reason(k, v, result) {
            return supportedEnumReason(k, v, ['number of seconds', 'milliseconds']);
        },
    };
};

const normalizeTimingFunction = (v) => {
    v = (v || '').toString();
    if (v.match(/^linear|ease|ease-in|ease-out|ease-in-out$/)) {
        return { value: v };
    }
    let match;
    if ((match = v.match(/^cubic-bezier\(\s*(.*)\s*,\s*(.*)\s*,\s*(.*)\s*,\s*(.*)\s*\)$/))) {
        if (match[1].match(NUM_REGEXP) &&
            match[2].match(NUM_REGEXP) &&
            match[3].match(NUM_REGEXP) &&
            match[4].match(NUM_REGEXP)) {
            const ret = [
                parseFloat(match[1]),
                parseFloat(match[2]),
                parseFloat(match[3]),
                parseFloat(match[4]),
            ].join(',');
            return { value: 'cubic-bezier(' + ret + ')' };
        }
    }
    return {
        value: null,
        reason(k, v, result) {
            return supportedEnumReason(k, v, [
                'linear',
                'ease',
                'ease-in',
                'ease-out',
                'ease-in-out',
                'cubic-bezier(n,n,n,n)',
            ]);
        },
    };
};

function createCombinedNormalize(normalizes) {
    return (v, options) => {
        const reasons = [];
        for (let i = 0; i < normalizes.length; i++) {
            const result = normalizes[i](v, options);
            if (result.value !== null) {
                return result;
            }
            if (result.reason) {
                reasons.push(result.reason);
            }
        }
        return {
            value: null,
            reason(k, v, result) {
                return normalizeReasons(reasons.map((reason) => reason(k, v, result)), k, v).join('\n');
            },
        };
    };
}

const normalizeGradient = (v) => {
    v = (v || '').toString();
    if (/^linear-gradient(.+)$/s.test(v)) {
        return { value: v };
    }
    return {
        // 枚举里会做reason提示
        value: null,
    };
};
const normalizeUrl = (v) => {
    v = (v || '').toString();
    if (/^url(.+)$/s.test(v)) {
        return { value: v };
    }
    return {
        value: null,
    };
};

function normalizePlatform(normalize, uniPlatform) {
    return (v, options, declInfo) => {
        // platform 未定义时候忽略
        const currentPlatform = options.platform;
        const supportedPlatforms = getSupportedPlatforms(uniPlatform);
        // TODO 未跨平台支持的属性特殊提示
        if (!supportedPlatforms.includes(currentPlatform)) {
            return {
                value: v,
                reason(k, v, result) {
                    return supportedPropertyReason(k);
                },
            };
        }
        return normalize(v, options, declInfo);
    };
}

function normalizeShorthandProperty(normalize) {
    return (v, options) => {
        v = (v || '').toString();
        const value = [];
        const reasons = [];
        const results = v.split(/\s+/).map((v) => normalize(v, options));
        for (let i = 0; i < results.length; ++i) {
            const res = results[i];
            if (res.value === null) {
                return res;
            }
            if (res.reason) {
                reasons.push(res.reason);
            }
            value.push(res.value);
        }
        return {
            value: value.length === 1 ? value[0] : value.join(' '),
            reason: function (k, v, result) {
                return reasons.map((reason) => reason(k, v, result)).join('\n');
            },
        };
    };
}

function normalizeFontFace(normalize) {
    return (v, options, declInfo) => {
        if (declInfo?.atRule === 'font-face') {
            return {
                value: null,
                reason(k, v, result) {
                    const items = ['font-family', 'src'];
                    const name = '@' + declInfo.atRule;
                    return ('ERROR: property `' +
                        hyphenateStyleProperty(k) +
                        '` is not supported for `' +
                        name +
                        '` (supported properties are: `' +
                        items.join('`|`') +
                        '`)');
                },
            };
        }
        return normalize(v, options, declInfo);
    };
}
// 只有@font-face下的src属性才支持
const normalizeSrc = (v, options, declInfo) => {
    if (declInfo?.atRule === 'font-face') {
        return { value: v };
    }
    return {
        value: null,
        reason(k, v, result) {
            return supportedPropertyReason(k);
        },
    };
};

const normalizeFlexFlow = (v) => {
    v = (v || '').toString();
    const values = v.split(/\s+/);
    // flex-flow 需要定义每一个属性值
    if (values.length === 1) {
        return {
            value: v,
            reason(k, v, result) {
                return supportedValueWithTipsReason(k, v, '(both property values must be explicitly defined)');
            },
        };
    }
    return {
        value: v,
    };
};

// 从 property.ts 中移动到 map 里，避免循环依赖
const normalizeProperty = (v, options) => {
    v = (v || '').toString();
    v = v
        .split(/\s*,\s*/)
        .map(camelize)
        .join(',');
    if (v.split(/\s*,\s*/).every((p) => {
        return !!getNormalizeMap(options)[p];
    })) {
        return { value: v };
    }
    return {
        value: null,
        reason: function reason(k, v, result) {
            return supportedEnumReason(k, v, ['css property']);
        },
    };
};
const normalizeDefault = (v) => {
    return { value: v };
};
const NVUE_PROP_NAME_GROUPS = {
    boxModel: {
        display: createEnumNormalize(['flex']),
        width: normalizeLength,
        height: normalizeLength,
        overflow: createEnumNormalize(['hidden']),
        padding: normalizeShorthandLength,
        paddingLeft: normalizeLength,
        paddingRight: normalizeLength,
        paddingTop: normalizeLength,
        paddingBottom: normalizeLength,
        margin: normalizeShorthandLength,
        marginLeft: normalizeLength,
        marginRight: normalizeLength,
        marginTop: normalizeLength,
        marginBottom: normalizeLength,
        borderWidth: normalizeLength,
        borderLeftWidth: normalizeLength,
        borderTopWidth: normalizeLength,
        borderRightWidth: normalizeLength,
        borderBottomWidth: normalizeLength,
        borderColor: normalizeColor,
        borderLeftColor: normalizeColor,
        borderTopColor: normalizeColor,
        borderRightColor: normalizeColor,
        borderBottomColor: normalizeColor,
        borderStyle: createEnumNormalize(['dotted', 'dashed', 'solid']),
        borderTopStyle: createEnumNormalize(['dotted', 'dashed', 'solid']),
        borderRightStyle: createEnumNormalize(['dotted', 'dashed', 'solid']),
        borderBottomStyle: createEnumNormalize(['dotted', 'dashed', 'solid']),
        borderLeftStyle: createEnumNormalize(['dotted', 'dashed', 'solid']),
        borderRadius: normalizeLength,
        borderBottomLeftRadius: normalizeLength,
        borderBottomRightRadius: normalizeLength,
        borderTopLeftRadius: normalizeLength,
        borderTopRightRadius: normalizeLength,
    },
    flexbox: {
        flex: normalizeNumber,
        flexWrap: normalizeFlexWrap,
        flexDirection: createEnumNormalize([
            'column',
            'row',
            'column-reverse',
            'row-reverse',
        ]),
        justifyContent: createEnumNormalize([
            'flex-start',
            'flex-end',
            'center',
            'space-between',
            'space-around',
        ]),
        alignItems: createEnumNormalize([
            'stretch',
            'flex-start',
            'flex-end',
            'center',
        ]),
    },
    position: {
        position: createEnumNormalize(['relative', 'absolute', 'sticky', 'fixed']),
        top: normalizeLength,
        bottom: normalizeLength,
        left: normalizeLength,
        right: normalizeLength,
        zIndex: normalizeInteger,
    },
    common: {
        opacity: normalizeNumber,
        boxShadow: normalizeDefault,
        backgroundColor: normalizeColor,
        backgroundImage: normalizeDefault,
    },
    text: {
        lines: normalizeInteger,
        color: normalizeColor,
        fontSize: normalizeLength,
        fontStyle: createEnumNormalize(['normal', 'italic']),
        fontFamily: normalizeDefault,
        fontWeight: createEnumNormalize([
            'normal',
            'bold',
            '100',
            '200',
            '300',
            '400',
            '500',
            '600',
            '700',
            '800',
            '900',
        ]),
        textDecoration: createEnumNormalize(['none', 'underline', 'line-through']),
        textAlign: createEnumNormalize(['left', 'center', 'right']),
        textOverflow: createEnumNormalize(['clip', 'ellipsis', 'unset', 'fade']),
        lineHeight: normalizeLength,
    },
    transition: {
        transitionProperty: normalizeProperty,
        transitionDuration: normalizeInterval,
        transitionDelay: normalizeInterval,
        transitionTimingFunction: normalizeTimingFunction,
    },
    transform: {
        transform: normalizeTransform,
        transformOrigin: normalizeTransform, // fixed by xxxxxx
    },
    customized: {
        itemSize: normalizeLength,
        itemColor: normalizeColor,
        itemSelectedColor: normalizeColor,
        textColor: normalizeColor,
        timeColor: normalizeColor,
        textHighlightColor: normalizeColor,
    },
};
const uvueNormalizeMap = {
    transform: normalizeTransform,
    fontFamily: normalizeString,
    textDecoration: normalizeDefault,
    boxShadow: normalizeDefault,
    transitionProperty: normalizeProperty,
    transitionTimingFunction: normalizeTimingFunction,
};
const restrictionMap = {
    ["length" /* Restriction.LENGTH */]: normalizeLength,
    ["percentage" /* Restriction.PERCENTAGE */]: normalizePercent,
    ["number" /* Restriction.NUMBER */]: normalizeNumber,
    ["number(0-1)" /* Restriction.NUMBER_0_1 */]: normalizeNumber,
    ["integer" /* Restriction.INTEGER */]: normalizeInteger,
    ["color" /* Restriction.COLOR */]: normalizeColor,
    ["time" /* Restriction.TIME */]: normalizeInterval,
    ["property" /* Restriction.PROPERTY */]: normalizeProperty,
    ["timing-function" /* Restriction.TIMING_FUNCTION */]: normalizeTimingFunction,
    ["gradient" /* Restriction.GRADIENT */]: normalizeGradient,
    ["url" /* Restriction.URL */]: normalizeUrl,
};
// @font-face下不支持的属性
const invalidFontFaceProperties = ['fontWeight', 'fontStyle', 'fontVariant'];
function getUVueNormalizeMap() {
    const result = {
        src: normalizeSrc,
    };
    let cssJson;
    try {
        // eslint-disable-next-line no-restricted-globals
        cssJson = require('../lib/css.json');
    }
    catch (e) {
        // 单元测试环境，源码目录
        // eslint-disable-next-line no-restricted-globals
        cssJson = require('../../lib/css.json');
    }
    const { properties } = cssJson;
    for (let i = 0; i < properties.length; i++) {
        const property = properties[i];
        const prop = camelize(property.name);
        let normalize;
        if (uvueNormalizeMap[prop]) {
            normalize = uvueNormalizeMap[prop];
        }
        else {
            const normalizes = getNormalizes(property);
            if (normalizes.length > 1) {
                normalize = createCombinedNormalize(normalizes);
            }
            else if (normalizes.length === 1) {
                normalize = normalizes[0];
            }
            else {
                normalize = normalizeDefault;
            }
            // 简写属性
            if (property.shorthand) {
                normalize = normalizeShorthandProperty(normalize);
            }
            // 处理@font-face下不支持的属性
            if (invalidFontFaceProperties.includes(prop)) {
                normalize = normalizeFontFace(normalize);
            }
            // 校验flexFlow属性值的个数，先临时写死，后续考虑根据css.json动态判断
            if (prop === 'flexFlow') {
                normalize = createCombinedNormalize([normalizeFlexFlow, normalize]);
            }
        }
        result[prop] = normalizePlatform(normalize, property.uniPlatform);
    }
    return result;
}
function getNormalizes(property) {
    const normalizes = [];
    const { restrictions } = property;
    restrictions.forEach((restriction) => {
        let normalize = restrictionMap[restriction];
        if (normalize) {
            if (restriction === "length" /* Restriction.LENGTH */) {
                // 如果同时有number和length，例如line-height: 1.5, line-height: 16px，则不能移除px
                normalize = normalizeLengthWithOptions({
                    removePx: !restrictions.includes("number" /* Restriction.NUMBER */),
                    property: property.name,
                });
            }
            normalizes.push(normalize);
        }
    });
    // enum
    if (property?.values?.length) {
        normalizes.push(createEnumNormalizeWithPlatform(property.values));
    }
    return normalizes;
}
let normalizeMap;
function getNormalizeMap(options) {
    if (normalizeMap) {
        return normalizeMap;
    }
    const uvue = options.type === 'uvue';
    if (uvue) {
        normalizeMap = getUVueNormalizeMap();
    }
    else {
        normalizeMap = Object.keys(NVUE_PROP_NAME_GROUPS).reduce((res, name) => {
            const group = NVUE_PROP_NAME_GROUPS[name];
            Object.keys(group).forEach((prop) => {
                res[prop] = group[prop];
            });
            return res;
        }, {});
    }
    return normalizeMap;
}

const normalized = Symbol('normalized');
function normalize(opts = {}) {
    if (!hasOwn(opts, 'logLevel')) {
        opts.logLevel = 'WARNING';
    }
    const plugin = {
        postcssPlugin: `${opts.type || 'nvue'}:normalize`,
        Declaration: createDeclarationProcessor(opts),
    };
    {
        plugin.Rule = createRuleProcessor(opts);
    }
    return plugin;
}
function createRuleProcessor(opts = {}) {
    return (rule, helper) => {
        if (rule[normalized]) {
            return;
        }
        rule.selector = rule.selectors
            .map((selector) => {
            selector = selector
                .replace(/\s*([\+\~\>])\s*/g, '$1')
                .replace(/\s+/, ' ');
            if (COMBINATORS_RE.test(selector)) {
                return selector;
            }
            let type = opts.type || 'nvue';
            rule.warn(helper.result, 'ERROR: Selector `' +
                selector +
                '` is not supported. ' +
                type +
                ' only support classname selector');
            return '';
        })
            .filter(Boolean)
            .join(', ');
        if (!rule.selector) {
            rule.remove();
        }
        rule[normalized] = true;
    };
}
function createDeclarationProcessor(options) {
    return (decl, helper) => {
        if (decl[normalized]) {
            return;
        }
        decl.prop = camelize(decl.prop);
        const { value, log } = normalizeDecl(decl, options);
        if (isString(value) || isNumber(value)) {
            decl.value = value;
        }
        if (log && log.reason && helper) {
            const { reason } = log;
            let needLog = false;
            if (options.logLevel === 'NOTE') {
                needLog = true;
            }
            else if (options.logLevel === 'ERROR') {
                if (reason.startsWith('ERROR:')) {
                    needLog = true;
                }
            }
            else {
                if (!reason.startsWith('NOTE:')) {
                    needLog = true;
                }
            }
            needLog && decl.warn(helper.result, reason);
        }
        if (value === null) {
            decl.remove();
        }
        decl[normalized] = true;
    };
}
function normalizeDecl(decl, options) {
    let { prop: name, value } = decl;
    let result, log;
    const normalize = getNormalizeMap(options)[name];
    if (options.type === 'uvue') {
        if (hasCssVar(value)) {
            return {
                value: normalizeCssVar(value),
                log,
            };
        }
    }
    if (isFunction(normalize)) {
        if (!isFunction(value)) {
            result = normalize(value, options, {
                atRule: decl.parent?.type === 'atrule' ? decl.parent.name : '',
            });
        }
        else {
            result = { value: value };
        }
        if (result.reason) {
            log = { reason: result.reason(name, value, result.value) };
        }
    }
    else {
        // ensure number type, no `px`
        if (isString(value)) {
            const match = value.match(LENGTH_REGEXP);
            if (match && (!match[1] || SUPPORT_CSS_UNIT.indexOf(match[1]) === -1)) {
                value = parseFloat(value);
            }
        }
        result = { value: value };
        log = {
            reason: 'WARNING: `' +
                hyphenateStyleProperty(name) +
                '` is not a standard property name (may not be supported)',
        };
    }
    return {
        value: result.value,
        log,
    };
}
function hasCssVar(value) {
    return value.includes('var(');
}
function normalizeCssVar(value) {
    return value
        .replaceAll(`var(--window-top)`, `CSS_VAR_WINDOW_TOP`)
        .replaceAll(`var(--window-bottom)`, `CSS_VAR_WINDOW_BOTTOM`)
        .replaceAll(`var(--status-bar-height)`, `CSS_VAR_STATUS_BAR_HEIGHT`);
}

function objectifier(node, { trim } = { trim: false }) {
    if (!node) {
        return {};
    }
    const context = {
        'FONT-FACE': [],
        TRANSITION: {},
    };
    const result = transform(node, context);
    if (trim) {
        trimObj(result);
    }
    if (context['FONT-FACE'].length) {
        result['@FONT-FACE'] = context['FONT-FACE'];
    }
    if (Object.keys(context.TRANSITION).length) {
        result['@TRANSITION'] = context.TRANSITION;
    }
    return result;
}
function trimObj(obj) {
    Object.keys(obj).forEach((name) => {
        const value = obj[name];
        if (Object.keys(value).length === 1 && hasOwn(value, '')) {
            obj[name] = value[''];
        }
    });
}
function transform(node, context) {
    const result = {};
    node.each((child) => {
        if (child.type === 'atrule') {
            const body = transform(child, context);
            const fontFamily = body.fontFamily;
            if (fontFamily && '"\''.indexOf(fontFamily[0]) > -1) {
                body.fontFamily = fontFamily.slice(1, fontFamily.length - 1);
            }
            context['FONT-FACE'].push(body);
        }
        else if (child.type === 'rule') {
            const body = transform(child, context);
            child.selectors.forEach((selector) => {
                transformSelector(selector, body, result, context);
            });
        }
        else if (child.type === 'decl') {
            if (child.important) {
                result['!' + child.prop] = child.value;
                // !important的值域优先级高，故删除非!important的值域
                delete result[child.prop];
            }
            else {
                if (!hasOwn(result, '!' + child.prop)) {
                    result[child.prop] = child.value;
                }
            }
        }
    });
    return result;
}
function transformSelector(selector, body, result, context) {
    const res = selector.match(COMBINATORS_RE);
    if (!res) {
        return;
    }
    let parentSelector = res[1];
    let curSelector = res[2].substring(1);
    // .a.b => a.b
    const dotIndex = curSelector.indexOf('.');
    if (dotIndex > -1) {
        parentSelector += curSelector.substring(dotIndex);
        curSelector = curSelector.substring(0, dotIndex);
    }
    const pseudoIndex = curSelector.indexOf(':');
    if (pseudoIndex > -1) {
        const pseudoClass = curSelector.slice(pseudoIndex);
        curSelector = curSelector.slice(0, pseudoIndex);
        Object.keys(body).forEach(function (name) {
            body[name + pseudoClass] = body[name];
            delete body[name];
        });
    }
    transition(curSelector, body, context);
    if (!Object.keys(body).length) {
        return;
    }
    result = (result[curSelector] || (result[curSelector] = {}));
    if (result[parentSelector]) {
        // clone
        result[parentSelector] = processImportant(extend({}, result[parentSelector], body));
    }
    else {
        result[parentSelector] = body;
    }
}
/**
 * 处理 important 属性，如果某个属性是 important，需要将非 important 的该属性移除掉
 * @param body
 */
function processImportant(body) {
    Object.keys(body).forEach((name) => {
        if (name.startsWith('!')) {
            delete body[name.substring(1)];
        }
    });
    return body;
}
function transition(className, body, { TRANSITION }) {
    Object.keys(body).forEach((prop) => {
        if (prop.indexOf('transition') === 0 && prop !== 'transition') {
            const realProp = prop.replace('transition', '');
            TRANSITION[className] = TRANSITION[className] || {};
            TRANSITION[className][realProp[0].toLowerCase() + realProp.slice(1)] =
                body[prop];
        }
    });
}

async function parse(input, options = {}) {
    const { root, messages } = await postcss__default.default([
        expand(options),
        normalize(options),
    ])
        .process(input, {
        from: options.filename || 'foo.css',
    })
        .catch((err) => {
        return {
            root: null,
            messages: [
                {
                    type: 'error',
                    text: err.message,
                },
            ],
        };
    });
    if (options.noCode === true) {
        return { code: '', messages };
    }
    const obj = root ? objectifier(root, { trim: !!options.trim }) : {};
    if (options.map || options.mapOf) {
        return {
            code: mapToInitStringChunk(objToMap(obj), options.ts, true, options.mapOf, options.chunk),
            messages,
        };
    }
    let code = JSON.stringify(obj);
    if (options.type === 'uvue') {
        // TODO 暂时仅简易转换 CSS 变量
        code = code.replace(/\:\s*"(.+?)"/g, function (str, p1) {
            return isExpr(p1) ? `:${p1}` : str;
        });
    }
    return { code, messages };
}
function mapToInitStringChunk(map, ts = false, isRoot = false, mapOf = '', chunk = 0) {
    if (!chunk) {
        return mapToInitString(map, ts, isRoot, mapOf);
    }
    const chunks = [];
    let chunkMap = new Map();
    let chunkCount = 0;
    for (const [key, value] of map) {
        if (chunkCount === chunk) {
            chunks.push(mapToInitString(chunkMap, ts, isRoot, mapOf));
            chunkMap = new Map();
            chunkCount = 0;
        }
        chunkMap.set(key, value);
        chunkCount++;
    }
    if (chunkCount) {
        chunks.push(mapToInitString(chunkMap, ts, isRoot, mapOf));
    }
    return `[${chunks.join(',')}]`;
}
function mapToInitString(map, ts = false, isRoot = false, mapOf = '') {
    const entries = [];
    for (let [key, value] of map) {
        if (value instanceof Map) {
            // trim
            if (isRoot && !(value.values().next().value instanceof Map)) {
                entries.push(`["${key}", padStyleMapOf(${mapToInitString(value, ts, false, mapOf)})]`);
            }
            else {
                entries.push(`["${key}", ${mapToInitString(value, ts, false, mapOf)}]`);
            }
        }
        else {
            entries.push(`["${key}", ${isString(value) && isExpr(value) ? value : JSON.stringify(value)}]`);
        }
    }
    if (mapOf) {
        return `${mapOf}([${entries.join(', ')}])`;
    }
    return `new Map${ts
        ? isRoot
            ? '<string, Map<string, Map<string, any>>>'
            : '<string, any>'
        : ''}([${entries.join(', ')}])`;
}
function objToMap(obj) {
    const map = new Map();
    for (const key in obj) {
        const value = obj[key];
        if (typeof value === 'object') {
            map.set(key, objToMap(value));
        }
        else {
            map.set(key, value);
        }
    }
    return map;
}
function isExpr(value) {
    const v = value.slice(0, 5);
    return /* CSS_VAR_ */ v === 'CSS_V' || v === 'calc(';
}

exports.expand = expand;
exports.normalize = normalize;
exports.objectifier = objectifier;
exports.parse = parse;
