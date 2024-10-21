import { hyphenate, capitalize } from '@vue/shared';

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

const backgroundColor = 'backgroundColor';
const backgroundImage = 'backgroundImage';
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

const borderWidth = 'Width';
const borderStyle = 'Style';
const borderColor = 'Color';
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

const borderTop = 'borderTop';
const borderRight = 'borderRight';
const borderBottom = 'borderBottom';
const borderLeft = 'borderLeft';
const transformBorderColor = (decl) => {
    const { prop, value, important, raws, source } = decl;
    let property = hyphenate(prop).split('-')[1];
    {
        property = capitalize(property);
    }
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

const borderTopLeftRadius = 'borderTopLeftRadius';
const borderTopRightRadius = 'borderTopRightRadius';
const borderBottomRightRadius = 'borderBottomRightRadius';
const borderBottomLeftRadius = 'borderBottomLeftRadius';
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

const flexDirection = 'flexDirection';
const flexWrap = 'flexWrap';
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

const top = 'Top';
const right = 'Right';
const bottom = 'Bottom';
const left = 'Left';
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

const transitionProperty = 'transitionProperty';
const transitionDuration = 'transitionDuration';
const transitionTimingFunction = 'transitionTimingFunction';
const transitionDelay = 'transitionDelay';
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
        result = styleMap;
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

export { expand };
