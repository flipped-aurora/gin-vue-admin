"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformStyle = exports.genJsStylesCode = exports.genStyle = void 0;
const descriptorCache_1 = require("../descriptorCache");
function genStyle(_, { className }) {
    return `/*${className}Styles*/`;
}
exports.genStyle = genStyle;
async function genJsStylesCode(descriptor, pluginContext) {
    let stylesCode = ``;
    if (descriptor.styles.length) {
        for (let i = 0; i < descriptor.styles.length; i++) {
            const style = descriptor.styles[i];
            if (style.src) {
                await linkSrcToDescriptor(style.src, descriptor, pluginContext);
            }
            const src = style.src || descriptor.filename;
            // do not include module in default query, since we use it to indicate
            // that the module needs to export the modules json
            const attrsQuery = attrsToQuery(style.attrs, 'css');
            const srcQuery = style.src ? '&src=true' : '';
            const query = `?vue&type=style&index=${i}${srcQuery}`;
            const styleRequest = src + query + attrsQuery;
            stylesCode += `\nimport ${JSON.stringify(styleRequest)}`;
        }
    }
    return stylesCode;
}
exports.genJsStylesCode = genJsStylesCode;
/**
 * For blocks with src imports, it is important to link the imported file
 * with its owner SFC descriptor so that we can get the information about
 * the owner SFC when compiling that file in the transform phase.
 */
async function linkSrcToDescriptor(src, descriptor, pluginContext) {
    const srcFile = (await pluginContext.resolve(src, descriptor.filename))?.id || src;
    // #1812 if the src points to a dep file, the resolved id may contain a
    // version query.
    (0, descriptorCache_1.setSrcDescriptor)(srcFile.replace(/\?.*$/, ''), descriptor);
}
// these are built-in query parameters so should be ignored
// if the user happen to add them as attrs
const ignoreList = ['id', 'index', 'src', 'type', 'lang', 'module', 'scoped'];
function attrsToQuery(attrs, langFallback, forceLangFallback = false) {
    let query = ``;
    for (const name in attrs) {
        const value = attrs[name];
        if (!ignoreList.includes(name)) {
            query += `&${encodeURIComponent(name)}${value ? `=${encodeURIComponent(value)}` : ``}`;
        }
    }
    if (langFallback || attrs.lang) {
        query +=
            `lang` in attrs
                ? forceLangFallback
                    ? `&lang.${langFallback}`
                    : `&lang.${attrs.lang}`
                : `&lang.${langFallback}`;
    }
    return query;
}
async function transformStyle(code, descriptor, index, options, pluginContext, filename) {
    const block = descriptor.styles[index];
    // vite already handles pre-processors and CSS module so this is only
    // applying SFC-specific transforms like scoped mode and CSS vars rewrite (v-bind(var))
    const result = await options.compiler.compileStyleAsync({
        filename: descriptor.filename,
        id: `data-v-${descriptor.id}`,
        isProd: true,
        source: code,
    });
    if (result.errors.length) {
        result.errors.forEach((error) => {
            if (error.line && error.column) {
                error.loc = {
                    file: descriptor.filename,
                    line: error.line + block.loc.start.line,
                    column: error.column,
                };
            }
            pluginContext.error(error);
        });
        return null;
    }
    return {
        code: result.code,
        map: null,
    };
}
exports.transformStyle = transformStyle;
