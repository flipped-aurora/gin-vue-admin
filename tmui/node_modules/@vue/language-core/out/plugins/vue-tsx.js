"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.tsCodegen = void 0;
const computeds_1 = require("computeds");
const script_1 = require("../generators/script");
const template_1 = require("../generators/template");
const scriptRanges_1 = require("../parsers/scriptRanges");
const scriptSetupRanges_1 = require("../parsers/scriptSetupRanges");
const language_core_1 = require("@volar/language-core");
const muggle = require("muggle-string");
const templateFormatReg = /^\.template_format\.ts$/;
const templateStyleCssReg = /^\.template_style\.css$/;
exports.tsCodegen = new WeakMap();
const plugin = (ctx) => {
    return {
        version: 1,
        requiredCompilerOptions: [
            'noPropertyAccessFromIndexSignature',
            'exactOptionalPropertyTypes',
        ],
        getEmbeddedFileNames(fileName, sfc) {
            const tsx = useTsx(fileName, sfc);
            const fileNames = [];
            if (['js', 'ts', 'jsx', 'tsx'].includes(tsx.lang())) {
                fileNames.push(fileName + '.' + tsx.lang());
            }
            if (sfc.template) {
                fileNames.push(fileName + '.template_format.ts');
                fileNames.push(fileName + '.template_style.css');
            }
            return fileNames;
        },
        resolveEmbeddedFile(fileName, sfc, embeddedFile) {
            const _tsx = useTsx(fileName, sfc);
            const suffix = embeddedFile.fileName.replace(fileName, '');
            if (suffix === '.' + _tsx.lang()) {
                embeddedFile.kind = language_core_1.FileKind.TypeScriptHostFile;
                embeddedFile.capabilities = {
                    ...language_core_1.FileCapabilities.full,
                    foldingRange: false,
                    documentFormatting: false,
                    documentSymbol: false,
                };
                const tsx = _tsx.generatedScript();
                if (tsx) {
                    const [content, contentStacks] = ctx.codegenStack ? muggle.track([...tsx.codes], [...tsx.codeStacks]) : [[...tsx.codes], [...tsx.codeStacks]];
                    embeddedFile.content = content;
                    embeddedFile.contentStacks = contentStacks;
                    embeddedFile.mirrorBehaviorMappings = [...tsx.mirrorBehaviorMappings];
                }
            }
            else if (suffix.match(templateFormatReg)) {
                embeddedFile.parentFileName = fileName + '.template.' + sfc.template?.lang;
                embeddedFile.kind = language_core_1.FileKind.TextFile;
                embeddedFile.capabilities = {
                    ...language_core_1.FileCapabilities.full,
                    diagnostic: false,
                    foldingRange: false,
                    codeAction: false,
                    inlayHint: false,
                };
                const template = _tsx.generatedTemplate();
                if (template) {
                    const [content, contentStacks] = ctx.codegenStack
                        ? muggle.track([...template.formatCodes], [...template.formatCodeStacks])
                        : [[...template.formatCodes], [...template.formatCodeStacks]];
                    embeddedFile.content = content;
                    embeddedFile.contentStacks = contentStacks;
                }
                for (const style of sfc.styles) {
                    embeddedFile.content.push('\n\n');
                    for (const cssVar of style.cssVars) {
                        embeddedFile.content.push('(');
                        embeddedFile.content.push([
                            cssVar.text,
                            style.name,
                            cssVar.offset,
                            {},
                        ]);
                        embeddedFile.content.push(');\n');
                    }
                }
            }
            else if (suffix.match(templateStyleCssReg)) {
                embeddedFile.parentFileName = fileName + '.template.' + sfc.template?.lang;
                const template = _tsx.generatedTemplate();
                if (template) {
                    const [content, contentStacks] = ctx.codegenStack
                        ? muggle.track([...template.cssCodes], [...template.cssCodeStacks])
                        : [[...template.cssCodes], [...template.cssCodeStacks]];
                    embeddedFile.content = content;
                    embeddedFile.contentStacks = contentStacks;
                }
                // for color pickers support
                embeddedFile.capabilities.documentSymbol = true;
            }
        },
    };
    function useTsx(fileName, sfc) {
        if (!exports.tsCodegen.has(sfc)) {
            exports.tsCodegen.set(sfc, createTsx(fileName, sfc, ctx));
        }
        return exports.tsCodegen.get(sfc);
    }
};
exports.default = plugin;
function createTsx(fileName, _sfc, { vueCompilerOptions, compilerOptions, codegenStack, modules }) {
    const ts = modules.typescript;
    const lang = (0, computeds_1.computed)(() => {
        return !_sfc.script && !_sfc.scriptSetup ? 'ts'
            : _sfc.scriptSetup && _sfc.scriptSetup.lang !== 'js' ? _sfc.scriptSetup.lang
                : _sfc.script && _sfc.script.lang !== 'js' ? _sfc.script.lang
                    : 'js';
    });
    const scriptRanges = (0, computeds_1.computed)(() => _sfc.script
        ? (0, scriptRanges_1.parseScriptRanges)(ts, _sfc.script.ast, !!_sfc.scriptSetup, false)
        : undefined);
    const scriptSetupRanges = (0, computeds_1.computed)(() => _sfc.scriptSetup
        ? (0, scriptSetupRanges_1.parseScriptSetupRanges)(ts, _sfc.scriptSetup.ast, vueCompilerOptions)
        : undefined);
    const shouldGenerateScopedClasses = (0, computeds_1.computed)(() => {
        const option = vueCompilerOptions.experimentalResolveStyleCssClasses;
        return _sfc.styles.some(s => {
            return option === 'always' || (option === 'scoped' && s.scoped);
        });
    });
    const stylesScopedClasses = (0, computeds_1.computedSet)(() => {
        const classes = new Set();
        if (!shouldGenerateScopedClasses()) {
            return classes;
        }
        for (const style of _sfc.styles) {
            const option = vueCompilerOptions.experimentalResolveStyleCssClasses;
            if (option === 'always' || (option === 'scoped' && style.scoped)) {
                for (const className of style.classNames) {
                    classes.add(className.text.substring(1));
                }
            }
        }
        return classes;
    });
    const generatedTemplate = (0, computeds_1.computed)(() => {
        if (!_sfc.template)
            return;
        return (0, template_1.generate)(ts, compilerOptions, vueCompilerOptions, _sfc.template, shouldGenerateScopedClasses(), stylesScopedClasses(), hasScriptSetupSlots(), slotsAssignName(), propsAssignName(), codegenStack);
    });
    const hasScriptSetupSlots = (0, computeds_1.computed)(() => !!scriptSetupRanges()?.slots.define);
    const slotsAssignName = (0, computeds_1.computed)(() => scriptSetupRanges()?.slots.name);
    const propsAssignName = (0, computeds_1.computed)(() => scriptSetupRanges()?.props.name);
    const generatedScript = (0, computeds_1.computed)(() => (0, script_1.generate)(ts, fileName, _sfc.script, _sfc.scriptSetup, _sfc.styles, lang(), scriptRanges(), scriptSetupRanges(), generatedTemplate(), compilerOptions, vueCompilerOptions, codegenStack));
    return {
        scriptRanges,
        scriptSetupRanges,
        lang,
        generatedScript,
        generatedTemplate,
    };
}
//# sourceMappingURL=vue-tsx.js.map