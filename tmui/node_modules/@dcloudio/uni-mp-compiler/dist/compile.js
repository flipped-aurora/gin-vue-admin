"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.baseCompile = exports.getBaseTransformPreset = void 0;
const fs_1 = __importDefault(require("fs"));
const compiler_core_1 = require("@vue/compiler-core");
const shared_1 = require("@vue/shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const codegen_1 = require("./codegen");
const transform_1 = require("./transform");
const transformExpression_1 = require("./transforms/transformExpression");
const transformIdentifier_1 = require("./transforms/transformIdentifier");
const vIf_1 = require("./transforms/vIf");
const vFor_1 = require("./transforms/vFor");
const codegen_2 = require("./template/codegen");
const vOn_1 = require("./transforms/vOn");
const transformElement_1 = require("./transforms/transformElement");
const vBind_1 = require("./transforms/vBind");
const transformComponent_1 = require("./transforms/transformComponent");
const vSlot_1 = require("./transforms/vSlot");
const transformRoot_1 = require("./transforms/transformRoot");
const transformTag_1 = require("./transforms/transformTag");
const vHtml_1 = require("./transforms/vHtml");
const vText_1 = require("./transforms/vText");
const transformAttr_1 = require("./transforms/transformAttr");
function getBaseTransformPreset({ prefixIdentifiers, skipTransformIdentifier, }) {
    // order is important
    const nodeTransforms = [
        transformRoot_1.transformRoot,
        transformAttr_1.transformAttr,
        transformTag_1.transformTag,
        vHtml_1.transformHtml,
        vText_1.transformText,
        vIf_1.transformIf,
        vFor_1.transformFor,
        vSlot_1.transformSlot,
    ];
    if (!skipTransformIdentifier) {
        nodeTransforms.push(transformIdentifier_1.transformIdentifier);
    }
    nodeTransforms.push(transformElement_1.transformElement);
    nodeTransforms.push(transformComponent_1.transformComponent);
    if (prefixIdentifiers) {
        nodeTransforms.push(transformExpression_1.transformExpression);
    }
    return [
        nodeTransforms,
        { on: vOn_1.transformOn, bind: vBind_1.transformBind },
    ];
}
exports.getBaseTransformPreset = getBaseTransformPreset;
function baseCompile(template, options = {}) {
    const prefixIdentifiers = options.prefixIdentifiers === true || options.mode === 'module';
    const ast = ((0, shared_1.isString)(template)
        ? (0, compiler_core_1.baseParse)(template, { ...options, parseMode: 'html' })
        : template);
    const [nodeTransforms, directiveTransforms] = getBaseTransformPreset({
        prefixIdentifiers,
        skipTransformIdentifier: options.skipTransformIdentifier === true,
    });
    options.hashId = genHashId(options);
    if (options.filename) {
        if (!options.filters && options.miniProgram?.filter) {
            options.filters = parseFilters(options.miniProgram.filter.lang, options.filename);
        }
    }
    const context = (0, transform_1.transform)(ast, (0, shared_1.extend)({}, options, {
        prefixIdentifiers,
        nodeTransforms: [...nodeTransforms, ...(options.nodeTransforms || [])],
        directiveTransforms: (0, shared_1.extend)({}, directiveTransforms, options.directiveTransforms || {}),
    }));
    const result = (0, shared_1.extend)((0, codegen_1.generate)((0, shared_1.extend)(ast, {
        bindingComponents: context.bindingComponents,
    }), options), { ast });
    if (options.filename && options.miniProgram?.emitFile) {
        const { class: clazz, directive, emitFile, event, slot, lazyElement, component, } = options.miniProgram;
        (0, codegen_2.generate)(ast, {
            class: clazz,
            scopeId: options.scopeId,
            filename: options.filename,
            directive,
            emitFile,
            event,
            slot,
            lazyElement,
            component,
            isBuiltInComponent: context.isBuiltInComponent,
            isMiniProgramComponent: context.isMiniProgramComponent,
        });
    }
    return result;
}
exports.baseCompile = baseCompile;
function genHashId(options) {
    if (options.hashId) {
        return options.hashId;
    }
    if (options.scopeId) {
        return options.scopeId.replace('data-v-', '');
    }
    if (options.filename) {
        return (0, uni_cli_shared_1.hash)(options.filename);
    }
    return '';
}
function parseFilters(lang, filename) {
    filename = filename.split('?')[0];
    if (fs_1.default.existsSync(filename)) {
        return (0, uni_cli_shared_1.parseFilterNames)(lang, fs_1.default.readFileSync(filename, 'utf8'));
    }
    return [];
}
