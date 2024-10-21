"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.errorMessages = exports.createCompilerError = exports.createDOMCompilerError = exports.defaultOnWarn = exports.defaultOnError = void 0;
function defaultOnError(error) {
    throw error;
}
exports.defaultOnError = defaultOnError;
function defaultOnWarn(msg) {
    console.warn(`[Vue warn] ${msg.message}`);
}
exports.defaultOnWarn = defaultOnWarn;
function createDOMCompilerError(code, loc) {
    return createCompilerError(code, loc);
}
exports.createDOMCompilerError = createDOMCompilerError;
function createCompilerError(code, loc, messages, additionalMessage) {
    const msg = (messages || exports.errorMessages)[code] + (additionalMessage || ``);
    const error = new SyntaxError(String(msg));
    error.code = code;
    error.loc = loc;
    return error;
}
exports.createCompilerError = createCompilerError;
exports.errorMessages = {
    // parse errors
    [0 /* ErrorCodes.ABRUPT_CLOSING_OF_EMPTY_COMMENT */]: 'Illegal comment.',
    [1 /* ErrorCodes.CDATA_IN_HTML_CONTENT */]: 'CDATA section is allowed only in XML context.',
    [2 /* ErrorCodes.DUPLICATE_ATTRIBUTE */]: 'Duplicate attribute.',
    [3 /* ErrorCodes.END_TAG_WITH_ATTRIBUTES */]: 'End tag cannot have attributes.',
    [4 /* ErrorCodes.END_TAG_WITH_TRAILING_SOLIDUS */]: "Illegal '/' in tags.",
    [5 /* ErrorCodes.EOF_BEFORE_TAG_NAME */]: 'Unexpected EOF in tag.',
    [6 /* ErrorCodes.EOF_IN_CDATA */]: 'Unexpected EOF in CDATA section.',
    [7 /* ErrorCodes.EOF_IN_COMMENT */]: 'Unexpected EOF in comment.',
    [8 /* ErrorCodes.EOF_IN_SCRIPT_HTML_COMMENT_LIKE_TEXT */]: 'Unexpected EOF in script.',
    [9 /* ErrorCodes.EOF_IN_TAG */]: 'Unexpected EOF in tag.',
    [10 /* ErrorCodes.INCORRECTLY_CLOSED_COMMENT */]: 'Incorrectly closed comment.',
    [11 /* ErrorCodes.INCORRECTLY_OPENED_COMMENT */]: 'Incorrectly opened comment.',
    [12 /* ErrorCodes.INVALID_FIRST_CHARACTER_OF_TAG_NAME */]: "Illegal tag name. Use '&lt;' to print '<'.",
    [13 /* ErrorCodes.MISSING_ATTRIBUTE_VALUE */]: 'Attribute value was expected.',
    [14 /* ErrorCodes.MISSING_END_TAG_NAME */]: 'End tag name was expected.',
    [15 /* ErrorCodes.MISSING_WHITESPACE_BETWEEN_ATTRIBUTES */]: 'Whitespace was expected.',
    [16 /* ErrorCodes.NESTED_COMMENT */]: "Unexpected '<!--' in comment.",
    [17 /* ErrorCodes.UNEXPECTED_CHARACTER_IN_ATTRIBUTE_NAME */]: 'Attribute name cannot contain U+0022 ("), U+0027 (\'), and U+003C (<).',
    [18 /* ErrorCodes.UNEXPECTED_CHARACTER_IN_UNQUOTED_ATTRIBUTE_VALUE */]: 'Unquoted attribute value cannot contain U+0022 ("), U+0027 (\'), U+003C (<), U+003D (=), and U+0060 (`).',
    [19 /* ErrorCodes.UNEXPECTED_EQUALS_SIGN_BEFORE_ATTRIBUTE_NAME */]: "Attribute name cannot start with '='.",
    [21 /* ErrorCodes.UNEXPECTED_QUESTION_MARK_INSTEAD_OF_TAG_NAME */]: "'<?' is allowed only in XML context.",
    [20 /* ErrorCodes.UNEXPECTED_NULL_CHARACTER */]: `Unexpected null character.`,
    [22 /* ErrorCodes.UNEXPECTED_SOLIDUS_IN_TAG */]: "Illegal '/' in tags.",
    // Vue-specific parse errors
    [23 /* ErrorCodes.X_INVALID_END_TAG */]: 'Invalid end tag.',
    [24 /* ErrorCodes.X_MISSING_END_TAG */]: 'Element is missing end tag.',
    [25 /* ErrorCodes.X_MISSING_INTERPOLATION_END */]: 'Interpolation end sign was not found.',
    [27 /* ErrorCodes.X_MISSING_DYNAMIC_DIRECTIVE_ARGUMENT_END */]: 'End bracket for dynamic directive argument was not found. ' +
        'Note that dynamic directive argument cannot contain spaces.',
    [26 /* ErrorCodes.X_MISSING_DIRECTIVE_NAME */]: 'Legal directive name was expected.',
    // transform errors
    [28 /* ErrorCodes.X_V_IF_NO_EXPRESSION */]: `v-if/v-else-if is missing expression.`,
    [29 /* ErrorCodes.X_V_IF_SAME_KEY */]: `v-if/else branches must use unique keys.`,
    [30 /* ErrorCodes.X_V_ELSE_NO_ADJACENT_IF */]: `v-else/v-else-if has no adjacent v-if or v-else-if.`,
    [31 /* ErrorCodes.X_V_FOR_NO_EXPRESSION */]: `v-for is missing expression.`,
    [32 /* ErrorCodes.X_V_FOR_MALFORMED_EXPRESSION */]: `v-for has invalid expression.`,
    [33 /* ErrorCodes.X_V_FOR_TEMPLATE_KEY_PLACEMENT */]: `<template v-for> key should be placed on the <template> tag.`,
    [34 /* ErrorCodes.X_V_BIND_NO_EXPRESSION */]: `v-bind is missing expression.`,
    [35 /* ErrorCodes.X_V_ON_NO_EXPRESSION */]: `v-on is missing expression.`,
    [36 /* ErrorCodes.X_V_SLOT_UNEXPECTED_DIRECTIVE_ON_SLOT_OUTLET */]: `Unexpected custom directive on <slot> outlet.`,
    [37 /* ErrorCodes.X_V_SLOT_MIXED_SLOT_USAGE */]: `Mixed v-slot usage on both the component and nested <template>. ` +
        `When there are multiple named slots, all slots should use <template> ` +
        `syntax to avoid scope ambiguity.`,
    [38 /* ErrorCodes.X_V_SLOT_DUPLICATE_SLOT_NAMES */]: `Duplicate slot names found. `,
    [39 /* ErrorCodes.X_V_SLOT_EXTRANEOUS_DEFAULT_SLOT_CHILDREN */]: `Extraneous children found when component already has explicitly named ` +
        `default slot. These children will be ignored.`,
    [40 /* ErrorCodes.X_V_SLOT_MISPLACED */]: `v-slot can only be used on components or <template> tags.`,
    [41 /* ErrorCodes.X_V_MODEL_NO_EXPRESSION */]: `v-model is missing expression.`,
    [42 /* ErrorCodes.X_V_MODEL_MALFORMED_EXPRESSION */]: `v-model value must be a valid JavaScript member expression.`,
    [43 /* ErrorCodes.X_V_MODEL_ON_SCOPE_VARIABLE */]: `v-model cannot be used on v-for or v-slot scope variables because they are not writable.`,
    [44 /* ErrorCodes.X_V_MODEL_ON_PROPS */]: `v-model cannot be used on a prop, because local prop bindings are not writable.\nUse a v-bind binding combined with a v-on listener that emits update:x event instead.`,
    [45 /* ErrorCodes.X_INVALID_EXPRESSION */]: `Error parsing JavaScript expression: `,
    [46 /* ErrorCodes.X_KEEP_ALIVE_INVALID_CHILDREN */]: `<KeepAlive> expects exactly one child component.`,
    [53 /* ErrorCodes.X_V_HTML_NO_EXPRESSION */]: `v-html is missing expression.`,
    [54 /* ErrorCodes.X_V_HTML_WITH_CHILDREN */]: `v-html will override element children.`,
    // generic errors
    [55 /* ErrorCodes.X_PREFIX_ID_NOT_SUPPORTED */]: `"prefixIdentifiers" option is not supported in this build of compiler.`,
    [56 /* ErrorCodes.X_MODULE_MODE_NOT_SUPPORTED */]: `ES module mode is not supported in this build of compiler.`,
    [57 /* ErrorCodes.X_CACHE_HANDLER_NOT_SUPPORTED */]: `"cacheHandlers" option is only supported when the "prefixIdentifiers" option is enabled.`,
    [58 /* ErrorCodes.X_SCOPE_ID_NOT_SUPPORTED */]: `"scopeId" option is only supported in module mode.`,
    // just to fulfill types
    [59 /* ErrorCodes.__EXTEND_POINT__ */]: ``,
};
