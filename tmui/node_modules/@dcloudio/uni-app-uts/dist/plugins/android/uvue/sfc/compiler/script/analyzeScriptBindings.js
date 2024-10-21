"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getObjectOrArrayExpressionKeys = exports.analyzeScriptBindings = void 0;
const compiler_dom_1 = require("@vue/compiler-dom");
const utils_1 = require("./utils");
/**
 * Analyze bindings in normal `<script>`
 * Note that `compileScriptSetup` already analyzes bindings as part of its
 * compilation process so this should only be used on single `<script>` SFCs.
 */
function analyzeScriptBindings(ast) {
    for (const node of ast) {
        if (node.type === 'ExportDefaultDeclaration' &&
            node.declaration.type === 'ObjectExpression') {
            return analyzeBindingsFromOptions(node.declaration);
        }
    }
    return {};
}
exports.analyzeScriptBindings = analyzeScriptBindings;
function analyzeBindingsFromOptions(node) {
    const bindings = {};
    // #3270, #3275
    // mark non-script-setup so we don't resolve components/directives from these
    Object.defineProperty(bindings, '__isScriptSetup', {
        enumerable: false,
        value: false,
    });
    for (const property of node.properties) {
        if (property.type === 'ObjectProperty' &&
            !property.computed &&
            property.key.type === 'Identifier') {
            // props
            if (property.key.name === 'props') {
                // props: ['foo']
                // props: { foo: ... }
                for (const key of getObjectOrArrayExpressionKeys(property.value)) {
                    bindings[key] = compiler_dom_1.BindingTypes.PROPS;
                }
            }
            // inject
            else if (property.key.name === 'inject') {
                // inject: ['foo']
                // inject: { foo: {} }
                for (const key of getObjectOrArrayExpressionKeys(property.value)) {
                    bindings[key] = compiler_dom_1.BindingTypes.OPTIONS;
                }
            }
            // computed & methods
            else if (property.value.type === 'ObjectExpression' &&
                (property.key.name === 'computed' || property.key.name === 'methods')) {
                // methods: { foo() {} }
                // computed: { foo() {} }
                for (const key of getObjectExpressionKeys(property.value)) {
                    bindings[key] = compiler_dom_1.BindingTypes.OPTIONS;
                }
            }
        }
        // setup & data
        else if (property.type === 'ObjectMethod' &&
            property.key.type === 'Identifier' &&
            (property.key.name === 'setup' || property.key.name === 'data')) {
            for (const bodyItem of property.body.body) {
                // setup() {
                //   return {
                //     foo: null
                //   }
                // }
                if (bodyItem.type === 'ReturnStatement' &&
                    bodyItem.argument &&
                    bodyItem.argument.type === 'ObjectExpression') {
                    for (const key of getObjectExpressionKeys(bodyItem.argument)) {
                        bindings[key] =
                            property.key.name === 'setup'
                                ? compiler_dom_1.BindingTypes.SETUP_MAYBE_REF
                                : compiler_dom_1.BindingTypes.DATA;
                    }
                }
            }
        }
    }
    return bindings;
}
function getObjectExpressionKeys(node) {
    const keys = [];
    for (const prop of node.properties) {
        if (prop.type === 'SpreadElement')
            continue;
        const key = (0, utils_1.resolveObjectKey)(prop.key, prop.computed);
        if (key)
            keys.push(String(key));
    }
    return keys;
}
function getArrayExpressionKeys(node) {
    const keys = [];
    for (const element of node.elements) {
        if (element && element.type === 'StringLiteral') {
            keys.push(element.value);
        }
    }
    return keys;
}
function getObjectOrArrayExpressionKeys(value) {
    if (value.type === 'ArrayExpression') {
        return getArrayExpressionKeys(value);
    }
    if (value.type === 'ObjectExpression') {
        return getObjectExpressionKeys(value);
    }
    return [];
}
exports.getObjectOrArrayExpressionKeys = getObjectOrArrayExpressionKeys;
