/*!
  * core-base v9.2.2
  * (c) 2022 kazuya kawaguchi
  * Released under the MIT License.
  */
import { isObject, isString, isFunction, isNumber, isPlainObject, toDisplayString, isArray, format, isBoolean, assign, isRegExp, warn, escapeHtml, inBrowser, mark, measure, isEmptyObject, generateCodeFrame, generateFormatCacheKey, isDate, getGlobalThis } from '@intlify/shared';
import { defaultOnError, baseCompile, CompileErrorCodes, createCompileError } from '@intlify/message-compiler';
export { CompileErrorCodes, createCompileError } from '@intlify/message-compiler';
import { IntlifyDevToolsHooks } from '@intlify/devtools-if';

const pathStateMachine =  [];
pathStateMachine[0 /* BEFORE_PATH */] = {
    ["w" /* WORKSPACE */]: [0 /* BEFORE_PATH */],
    ["i" /* IDENT */]: [3 /* IN_IDENT */, 0 /* APPEND */],
    ["[" /* LEFT_BRACKET */]: [4 /* IN_SUB_PATH */],
    ["o" /* END_OF_FAIL */]: [7 /* AFTER_PATH */]
};
pathStateMachine[1 /* IN_PATH */] = {
    ["w" /* WORKSPACE */]: [1 /* IN_PATH */],
    ["." /* DOT */]: [2 /* BEFORE_IDENT */],
    ["[" /* LEFT_BRACKET */]: [4 /* IN_SUB_PATH */],
    ["o" /* END_OF_FAIL */]: [7 /* AFTER_PATH */]
};
pathStateMachine[2 /* BEFORE_IDENT */] = {
    ["w" /* WORKSPACE */]: [2 /* BEFORE_IDENT */],
    ["i" /* IDENT */]: [3 /* IN_IDENT */, 0 /* APPEND */],
    ["0" /* ZERO */]: [3 /* IN_IDENT */, 0 /* APPEND */]
};
pathStateMachine[3 /* IN_IDENT */] = {
    ["i" /* IDENT */]: [3 /* IN_IDENT */, 0 /* APPEND */],
    ["0" /* ZERO */]: [3 /* IN_IDENT */, 0 /* APPEND */],
    ["w" /* WORKSPACE */]: [1 /* IN_PATH */, 1 /* PUSH */],
    ["." /* DOT */]: [2 /* BEFORE_IDENT */, 1 /* PUSH */],
    ["[" /* LEFT_BRACKET */]: [4 /* IN_SUB_PATH */, 1 /* PUSH */],
    ["o" /* END_OF_FAIL */]: [7 /* AFTER_PATH */, 1 /* PUSH */]
};
pathStateMachine[4 /* IN_SUB_PATH */] = {
    ["'" /* SINGLE_QUOTE */]: [5 /* IN_SINGLE_QUOTE */, 0 /* APPEND */],
    ["\"" /* DOUBLE_QUOTE */]: [6 /* IN_DOUBLE_QUOTE */, 0 /* APPEND */],
    ["[" /* LEFT_BRACKET */]: [
        4 /* IN_SUB_PATH */,
        2 /* INC_SUB_PATH_DEPTH */
    ],
    ["]" /* RIGHT_BRACKET */]: [1 /* IN_PATH */, 3 /* PUSH_SUB_PATH */],
    ["o" /* END_OF_FAIL */]: 8 /* ERROR */,
    ["l" /* ELSE */]: [4 /* IN_SUB_PATH */, 0 /* APPEND */]
};
pathStateMachine[5 /* IN_SINGLE_QUOTE */] = {
    ["'" /* SINGLE_QUOTE */]: [4 /* IN_SUB_PATH */, 0 /* APPEND */],
    ["o" /* END_OF_FAIL */]: 8 /* ERROR */,
    ["l" /* ELSE */]: [5 /* IN_SINGLE_QUOTE */, 0 /* APPEND */]
};
pathStateMachine[6 /* IN_DOUBLE_QUOTE */] = {
    ["\"" /* DOUBLE_QUOTE */]: [4 /* IN_SUB_PATH */, 0 /* APPEND */],
    ["o" /* END_OF_FAIL */]: 8 /* ERROR */,
    ["l" /* ELSE */]: [6 /* IN_DOUBLE_QUOTE */, 0 /* APPEND */]
};
/**
 * Check if an expression is a literal value.
 */
const literalValueRE = /^\s?(?:true|false|-?[\d.]+|'[^']*'|"[^"]*")\s?$/;
function isLiteral(exp) {
    return literalValueRE.test(exp);
}
/**
 * Strip quotes from a string
 */
function stripQuotes(str) {
    const a = str.charCodeAt(0);
    const b = str.charCodeAt(str.length - 1);
    return a === b && (a === 0x22 || a === 0x27) ? str.slice(1, -1) : str;
}
/**
 * Determine the type of a character in a keypath.
 */
function getPathCharType(ch) {
    if (ch === undefined || ch === null) {
        return "o" /* END_OF_FAIL */;
    }
    const code = ch.charCodeAt(0);
    switch (code) {
        case 0x5b: // [
        case 0x5d: // ]
        case 0x2e: // .
        case 0x22: // "
        case 0x27: // '
            return ch;
        case 0x5f: // _
        case 0x24: // $
        case 0x2d: // -
            return "i" /* IDENT */;
        case 0x09: // Tab (HT)
        case 0x0a: // Newline (LF)
        case 0x0d: // Return (CR)
        case 0xa0: // No-break space (NBSP)
        case 0xfeff: // Byte Order Mark (BOM)
        case 0x2028: // Line Separator (LS)
        case 0x2029: // Paragraph Separator (PS)
            return "w" /* WORKSPACE */;
    }
    return "i" /* IDENT */;
}
/**
 * Format a subPath, return its plain form if it is
 * a literal string or number. Otherwise prepend the
 * dynamic indicator (*).
 */
function formatSubPath(path) {
    const trimmed = path.trim();
    // invalid leading 0
    if (path.charAt(0) === '0' && isNaN(parseInt(path))) {
        return false;
    }
    return isLiteral(trimmed)
        ? stripQuotes(trimmed)
        : "*" /* ASTARISK */ + trimmed;
}
/**
 * Parse a string path into an array of segments
 */
function parse(path) {
    const keys = [];
    let index = -1;
    let mode = 0 /* BEFORE_PATH */;
    let subPathDepth = 0;
    let c;
    let key; // eslint-disable-line
    let newChar;
    let type;
    let transition;
    let action;
    let typeMap;
    const actions = [];
    actions[0 /* APPEND */] = () => {
        if (key === undefined) {
            key = newChar;
        }
        else {
            key += newChar;
        }
    };
    actions[1 /* PUSH */] = () => {
        if (key !== undefined) {
            keys.push(key);
            key = undefined;
        }
    };
    actions[2 /* INC_SUB_PATH_DEPTH */] = () => {
        actions[0 /* APPEND */]();
        subPathDepth++;
    };
    actions[3 /* PUSH_SUB_PATH */] = () => {
        if (subPathDepth > 0) {
            subPathDepth--;
            mode = 4 /* IN_SUB_PATH */;
            actions[0 /* APPEND */]();
        }
        else {
            subPathDepth = 0;
            if (key === undefined) {
                return false;
            }
            key = formatSubPath(key);
            if (key === false) {
                return false;
            }
            else {
                actions[1 /* PUSH */]();
            }
        }
    };
    function maybeUnescapeQuote() {
        const nextChar = path[index + 1];
        if ((mode === 5 /* IN_SINGLE_QUOTE */ &&
            nextChar === "'" /* SINGLE_QUOTE */) ||
            (mode === 6 /* IN_DOUBLE_QUOTE */ &&
                nextChar === "\"" /* DOUBLE_QUOTE */)) {
            index++;
            newChar = '\\' + nextChar;
            actions[0 /* APPEND */]();
            return true;
        }
    }
    while (mode !== null) {
        index++;
        c = path[index];
        if (c === '\\' && maybeUnescapeQuote()) {
            continue;
        }
        type = getPathCharType(c);
        typeMap = pathStateMachine[mode];
        transition = typeMap[type] || typeMap["l" /* ELSE */] || 8 /* ERROR */;
        // check parse error
        if (transition === 8 /* ERROR */) {
            return;
        }
        mode = transition[0];
        if (transition[1] !== undefined) {
            action = actions[transition[1]];
            if (action) {
                newChar = c;
                if (action() === false) {
                    return;
                }
            }
        }
        // check parse finish
        if (mode === 7 /* AFTER_PATH */) {
            return keys;
        }
    }
}
// path token cache
const cache = new Map();
/**
 * key-value message resolver
 *
 * @remarks
 * Resolves messages with the key-value structure. Note that messages with a hierarchical structure such as objects cannot be resolved
 *
 * @param obj - A target object to be resolved with path
 * @param path - A {@link Path | path} to resolve the value of message
 *
 * @returns A resolved {@link PathValue | path value}
 *
 * @VueI18nGeneral
 */
function resolveWithKeyValue(obj, path) {
    return isObject(obj) ? obj[path] : null;
}
/**
 * message resolver
 *
 * @remarks
 * Resolves messages. messages with a hierarchical structure such as objects can be resolved. This resolver is used in VueI18n as default.
 *
 * @param obj - A target object to be resolved with path
 * @param path - A {@link Path | path} to resolve the value of message
 *
 * @returns A resolved {@link PathValue | path value}
 *
 * @VueI18nGeneral
 */
function resolveValue(obj, path) {
    // check object
    if (!isObject(obj)) {
        return null;
    }
    // parse path
    let hit = cache.get(path);
    if (!hit) {
        hit = parse(path);
        if (hit) {
            cache.set(path, hit);
        }
    }
    // check hit
    if (!hit) {
        return null;
    }
    // resolve path value
    const len = hit.length;
    let last = obj;
    let i = 0;
    while (i < len) {
        const val = last[hit[i]];
        if (val === undefined) {
            return null;
        }
        last = val;
        i++;
    }
    return last;
}

const DEFAULT_MODIFIER = (str) => str;
const DEFAULT_MESSAGE = (ctx) => ''; // eslint-disable-line
const DEFAULT_MESSAGE_DATA_TYPE = 'text';
const DEFAULT_NORMALIZE = (values) => values.length === 0 ? '' : values.join('');
const DEFAULT_INTERPOLATE = toDisplayString;
function pluralDefault(choice, choicesLength) {
    choice = Math.abs(choice);
    if (choicesLength === 2) {
        // prettier-ignore
        return choice
            ? choice > 1
                ? 1
                : 0
            : 1;
    }
    return choice ? Math.min(choice, 2) : 0;
}
function getPluralIndex(options) {
    // prettier-ignore
    const index = isNumber(options.pluralIndex)
        ? options.pluralIndex
        : -1;
    // prettier-ignore
    return options.named && (isNumber(options.named.count) || isNumber(options.named.n))
        ? isNumber(options.named.count)
            ? options.named.count
            : isNumber(options.named.n)
                ? options.named.n
                : index
        : index;
}
function normalizeNamed(pluralIndex, props) {
    if (!props.count) {
        props.count = pluralIndex;
    }
    if (!props.n) {
        props.n = pluralIndex;
    }
}
function createMessageContext(options = {}) {
    const locale = options.locale;
    const pluralIndex = getPluralIndex(options);
    const pluralRule = isObject(options.pluralRules) &&
        isString(locale) &&
        isFunction(options.pluralRules[locale])
        ? options.pluralRules[locale]
        : pluralDefault;
    const orgPluralRule = isObject(options.pluralRules) &&
        isString(locale) &&
        isFunction(options.pluralRules[locale])
        ? pluralDefault
        : undefined;
    const plural = (messages) => {
        return messages[pluralRule(pluralIndex, messages.length, orgPluralRule)];
    };
    const _list = options.list || [];
    const list = (index) => _list[index];
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const _named = options.named || {};
    isNumber(options.pluralIndex) && normalizeNamed(pluralIndex, _named);
    const named = (key) => _named[key];
    function message(key) {
        // prettier-ignore
        const msg = isFunction(options.messages)
            ? options.messages(key)
            : isObject(options.messages)
                ? options.messages[key]
                : false;
        return !msg
            ? options.parent
                ? options.parent.message(key) // resolve from parent messages
                : DEFAULT_MESSAGE
            : msg;
    }
    const _modifier = (name) => options.modifiers
        ? options.modifiers[name]
        : DEFAULT_MODIFIER;
    const normalize = isPlainObject(options.processor) && isFunction(options.processor.normalize)
        ? options.processor.normalize
        : DEFAULT_NORMALIZE;
    const interpolate = isPlainObject(options.processor) &&
        isFunction(options.processor.interpolate)
        ? options.processor.interpolate
        : DEFAULT_INTERPOLATE;
    const type = isPlainObject(options.processor) && isString(options.processor.type)
        ? options.processor.type
        : DEFAULT_MESSAGE_DATA_TYPE;
    const linked = (key, ...args) => {
        const [arg1, arg2] = args;
        let type = 'text';
        let modifier = '';
        if (args.length === 1) {
            if (isObject(arg1)) {
                modifier = arg1.modifier || modifier;
                type = arg1.type || type;
            }
            else if (isString(arg1)) {
                modifier = arg1 || modifier;
            }
        }
        else if (args.length === 2) {
            if (isString(arg1)) {
                modifier = arg1 || modifier;
            }
            if (isString(arg2)) {
                type = arg2 || type;
            }
        }
        let msg = message(key)(ctx);
        // The message in vnode resolved with linked are returned as an array by processor.nomalize
        if (type === 'vnode' && isArray(msg) && modifier) {
            msg = msg[0];
        }
        return modifier ? _modifier(modifier)(msg, type) : msg;
    };
    const ctx = {
        ["list" /* LIST */]: list,
        ["named" /* NAMED */]: named,
        ["plural" /* PLURAL */]: plural,
        ["linked" /* LINKED */]: linked,
        ["message" /* MESSAGE */]: message,
        ["type" /* TYPE */]: type,
        ["interpolate" /* INTERPOLATE */]: interpolate,
        ["normalize" /* NORMALIZE */]: normalize
    };
    return ctx;
}

let devtools = null;
function setDevToolsHook(hook) {
    devtools = hook;
}
function getDevToolsHook() {
    return devtools;
}
function initI18nDevTools(i18n, version, meta) {
    // TODO: queue if devtools is undefined
    devtools &&
        devtools.emit(IntlifyDevToolsHooks.I18nInit, {
            timestamp: Date.now(),
            i18n,
            version,
            meta
        });
}
const translateDevTools = /* #__PURE__*/ createDevToolsHook(IntlifyDevToolsHooks.FunctionTranslate);
function createDevToolsHook(hook) {
    return (payloads) => devtools && devtools.emit(hook, payloads);
}

const CoreWarnCodes = {
    NOT_FOUND_KEY: 1,
    FALLBACK_TO_TRANSLATE: 2,
    CANNOT_FORMAT_NUMBER: 3,
    FALLBACK_TO_NUMBER_FORMAT: 4,
    CANNOT_FORMAT_DATE: 5,
    FALLBACK_TO_DATE_FORMAT: 6,
    __EXTEND_POINT__: 7
};
/** @internal */
const warnMessages = {
    [CoreWarnCodes.NOT_FOUND_KEY]: `Not found '{key}' key in '{locale}' locale messages.`,
    [CoreWarnCodes.FALLBACK_TO_TRANSLATE]: `Fall back to translate '{key}' key with '{target}' locale.`,
    [CoreWarnCodes.CANNOT_FORMAT_NUMBER]: `Cannot format a number value due to not supported Intl.NumberFormat.`,
    [CoreWarnCodes.FALLBACK_TO_NUMBER_FORMAT]: `Fall back to number format '{key}' key with '{target}' locale.`,
    [CoreWarnCodes.CANNOT_FORMAT_DATE]: `Cannot format a date value due to not supported Intl.DateTimeFormat.`,
    [CoreWarnCodes.FALLBACK_TO_DATE_FORMAT]: `Fall back to datetime format '{key}' key with '{target}' locale.`
};
function getWarnMessage(code, ...args) {
    return format(warnMessages[code], ...args);
}

/**
 * Fallback with simple implemenation
 *
 * @remarks
 * A fallback locale function implemented with a simple fallback algorithm.
 *
 * Basically, it returns the value as specified in the `fallbackLocale` props, and is processed with the fallback inside intlify.
 *
 * @param ctx - A {@link CoreContext | context}
 * @param fallback - A {@link FallbackLocale | fallback locale}
 * @param start - A starting {@link Locale | locale}
 *
 * @returns Fallback locales
 *
 * @VueI18nGeneral
 */
function fallbackWithSimple(ctx, fallback, start // eslint-disable-line @typescript-eslint/no-unused-vars
) {
    // prettier-ignore
    return [...new Set([
            start,
            ...(isArray(fallback)
                ? fallback
                : isObject(fallback)
                    ? Object.keys(fallback)
                    : isString(fallback)
                        ? [fallback]
                        : [start])
        ])];
}
/**
 * Fallback with locale chain
 *
 * @remarks
 * A fallback locale function implemented with a fallback chain algorithm. It's used in VueI18n as default.
 *
 * @param ctx - A {@link CoreContext | context}
 * @param fallback - A {@link FallbackLocale | fallback locale}
 * @param start - A starting {@link Locale | locale}
 *
 * @returns Fallback locales
 *
 * @VueI18nSee [Fallbacking](../guide/essentials/fallback)
 *
 * @VueI18nGeneral
 */
function fallbackWithLocaleChain(ctx, fallback, start) {
    const startLocale = isString(start) ? start : DEFAULT_LOCALE;
    const context = ctx;
    if (!context.__localeChainCache) {
        context.__localeChainCache = new Map();
    }
    let chain = context.__localeChainCache.get(startLocale);
    if (!chain) {
        chain = [];
        // first block defined by start
        let block = [start];
        // while any intervening block found
        while (isArray(block)) {
            block = appendBlockToChain(chain, block, fallback);
        }
        // prettier-ignore
        // last block defined by default
        const defaults = isArray(fallback) || !isPlainObject(fallback)
            ? fallback
            : fallback['default']
                ? fallback['default']
                : null;
        // convert defaults to array
        block = isString(defaults) ? [defaults] : defaults;
        if (isArray(block)) {
            appendBlockToChain(chain, block, false);
        }
        context.__localeChainCache.set(startLocale, chain);
    }
    return chain;
}
function appendBlockToChain(chain, block, blocks) {
    let follow = true;
    for (let i = 0; i < block.length && isBoolean(follow); i++) {
        const locale = block[i];
        if (isString(locale)) {
            follow = appendLocaleToChain(chain, block[i], blocks);
        }
    }
    return follow;
}
function appendLocaleToChain(chain, locale, blocks) {
    let follow;
    const tokens = locale.split('-');
    do {
        const target = tokens.join('-');
        follow = appendItemToChain(chain, target, blocks);
        tokens.splice(-1, 1);
    } while (tokens.length && follow === true);
    return follow;
}
function appendItemToChain(chain, target, blocks) {
    let follow = false;
    if (!chain.includes(target)) {
        follow = true;
        if (target) {
            follow = target[target.length - 1] !== '!';
            const locale = target.replace(/!/g, '');
            chain.push(locale);
            if ((isArray(blocks) || isPlainObject(blocks)) &&
                blocks[locale] // eslint-disable-line @typescript-eslint/no-explicit-any
            ) {
                // eslint-disable-next-line @typescript-eslint/no-explicit-any
                follow = blocks[locale];
            }
        }
    }
    return follow;
}

/* eslint-disable @typescript-eslint/no-explicit-any */
/**
 * Intlify core-base version
 * @internal
 */
const VERSION = '9.2.2';
const NOT_REOSLVED = -1;
const DEFAULT_LOCALE = 'en-US';
const MISSING_RESOLVE_VALUE = '';
const capitalize = (str) => `${str.charAt(0).toLocaleUpperCase()}${str.substr(1)}`;
function getDefaultLinkedModifiers() {
    return {
        upper: (val, type) => {
            // prettier-ignore
            return type === 'text' && isString(val)
                ? val.toUpperCase()
                : type === 'vnode' && isObject(val) && '__v_isVNode' in val
                    ? val.children.toUpperCase()
                    : val;
        },
        lower: (val, type) => {
            // prettier-ignore
            return type === 'text' && isString(val)
                ? val.toLowerCase()
                : type === 'vnode' && isObject(val) && '__v_isVNode' in val
                    ? val.children.toLowerCase()
                    : val;
        },
        capitalize: (val, type) => {
            // prettier-ignore
            return (type === 'text' && isString(val)
                ? capitalize(val)
                : type === 'vnode' && isObject(val) && '__v_isVNode' in val
                    ? capitalize(val.children)
                    : val);
        }
    };
}
let _compiler;
function registerMessageCompiler(compiler) {
    _compiler = compiler;
}
let _resolver;
/**
 * Register the message resolver
 *
 * @param resolver - A {@link MessageResolver} function
 *
 * @VueI18nGeneral
 */
function registerMessageResolver(resolver) {
    _resolver = resolver;
}
let _fallbacker;
/**
 * Register the locale fallbacker
 *
 * @param fallbacker - A {@link LocaleFallbacker} function
 *
 * @VueI18nGeneral
 */
function registerLocaleFallbacker(fallbacker) {
    _fallbacker = fallbacker;
}
// Additional Meta for Intlify DevTools
let _additionalMeta = null;
const setAdditionalMeta =  (meta) => {
    _additionalMeta = meta;
};
const getAdditionalMeta =  () => _additionalMeta;
let _fallbackContext = null;
const setFallbackContext = (context) => {
    _fallbackContext = context;
};
const getFallbackContext = () => _fallbackContext;
// ID for CoreContext
let _cid = 0;
function createCoreContext(options = {}) {
    // setup options
    const version = isString(options.version) ? options.version : VERSION;
    const locale = isString(options.locale) ? options.locale : DEFAULT_LOCALE;
    const fallbackLocale = isArray(options.fallbackLocale) ||
        isPlainObject(options.fallbackLocale) ||
        isString(options.fallbackLocale) ||
        options.fallbackLocale === false
        ? options.fallbackLocale
        : locale;
    const messages = isPlainObject(options.messages)
        ? options.messages
        : { [locale]: {} };
    const datetimeFormats = isPlainObject(options.datetimeFormats)
            ? options.datetimeFormats
            : { [locale]: {} }
        ;
    const numberFormats = isPlainObject(options.numberFormats)
            ? options.numberFormats
            : { [locale]: {} }
        ;
    const modifiers = assign({}, options.modifiers || {}, getDefaultLinkedModifiers());
    const pluralRules = options.pluralRules || {};
    const missing = isFunction(options.missing) ? options.missing : null;
    const missingWarn = isBoolean(options.missingWarn) || isRegExp(options.missingWarn)
        ? options.missingWarn
        : true;
    const fallbackWarn = isBoolean(options.fallbackWarn) || isRegExp(options.fallbackWarn)
        ? options.fallbackWarn
        : true;
    const fallbackFormat = !!options.fallbackFormat;
    const unresolving = !!options.unresolving;
    const postTranslation = isFunction(options.postTranslation)
        ? options.postTranslation
        : null;
    const processor = isPlainObject(options.processor) ? options.processor : null;
    const warnHtmlMessage = isBoolean(options.warnHtmlMessage)
        ? options.warnHtmlMessage
        : true;
    const escapeParameter = !!options.escapeParameter;
    const messageCompiler = isFunction(options.messageCompiler)
        ? options.messageCompiler
        : _compiler;
    const messageResolver = isFunction(options.messageResolver)
        ? options.messageResolver
        : _resolver || resolveWithKeyValue;
    const localeFallbacker = isFunction(options.localeFallbacker)
        ? options.localeFallbacker
        : _fallbacker || fallbackWithSimple;
    const fallbackContext = isObject(options.fallbackContext)
        ? options.fallbackContext
        : undefined;
    const onWarn = isFunction(options.onWarn) ? options.onWarn : warn;
    // setup internal options
    const internalOptions = options;
    const __datetimeFormatters = isObject(internalOptions.__datetimeFormatters)
            ? internalOptions.__datetimeFormatters
            : new Map()
        ;
    const __numberFormatters = isObject(internalOptions.__numberFormatters)
            ? internalOptions.__numberFormatters
            : new Map()
        ;
    const __meta = isObject(internalOptions.__meta) ? internalOptions.__meta : {};
    _cid++;
    const context = {
        version,
        cid: _cid,
        locale,
        fallbackLocale,
        messages,
        modifiers,
        pluralRules,
        missing,
        missingWarn,
        fallbackWarn,
        fallbackFormat,
        unresolving,
        postTranslation,
        processor,
        warnHtmlMessage,
        escapeParameter,
        messageCompiler,
        messageResolver,
        localeFallbacker,
        fallbackContext,
        onWarn,
        __meta
    };
    {
        context.datetimeFormats = datetimeFormats;
        context.numberFormats = numberFormats;
        context.__datetimeFormatters = __datetimeFormatters;
        context.__numberFormatters = __numberFormatters;
    }
    // for vue-devtools timeline event
    if ((process.env.NODE_ENV !== 'production')) {
        context.__v_emitter =
            internalOptions.__v_emitter != null
                ? internalOptions.__v_emitter
                : undefined;
    }
    // NOTE: experimental !!
    if ((process.env.NODE_ENV !== 'production') || __INTLIFY_PROD_DEVTOOLS__) {
        initI18nDevTools(context, version, __meta);
    }
    return context;
}
/** @internal */
function isTranslateFallbackWarn(fallback, key) {
    return fallback instanceof RegExp ? fallback.test(key) : fallback;
}
/** @internal */
function isTranslateMissingWarn(missing, key) {
    return missing instanceof RegExp ? missing.test(key) : missing;
}
/** @internal */
function handleMissing(context, key, locale, missingWarn, type) {
    const { missing, onWarn } = context;
    // for vue-devtools timeline event
    if ((process.env.NODE_ENV !== 'production')) {
        const emitter = context.__v_emitter;
        if (emitter) {
            emitter.emit("missing" /* MISSING */, {
                locale,
                key,
                type,
                groupId: `${type}:${key}`
            });
        }
    }
    if (missing !== null) {
        const ret = missing(context, locale, key, type);
        return isString(ret) ? ret : key;
    }
    else {
        if ((process.env.NODE_ENV !== 'production') && isTranslateMissingWarn(missingWarn, key)) {
            onWarn(getWarnMessage(CoreWarnCodes.NOT_FOUND_KEY, { key, locale }));
        }
        return key;
    }
}
/** @internal */
function updateFallbackLocale(ctx, locale, fallback) {
    const context = ctx;
    context.__localeChainCache = new Map();
    ctx.localeFallbacker(ctx, fallback, locale);
}
/* eslint-enable @typescript-eslint/no-explicit-any */

const RE_HTML_TAG = /<\/?[\w\s="/.':;#-\/]+>/;
const WARN_MESSAGE = `Detected HTML in '{source}' message. Recommend not using HTML messages to avoid XSS.`;
function checkHtmlMessage(source, options) {
    const warnHtmlMessage = isBoolean(options.warnHtmlMessage)
        ? options.warnHtmlMessage
        : true;
    if (warnHtmlMessage && RE_HTML_TAG.test(source)) {
        warn(format(WARN_MESSAGE, { source }));
    }
}
const defaultOnCacheKey = (source) => source;
let compileCache = Object.create(null);
function clearCompileCache() {
    compileCache = Object.create(null);
}
function compileToFunction(source, options = {}) {
    {
        // check HTML message
        (process.env.NODE_ENV !== 'production') && checkHtmlMessage(source, options);
        // check caches
        const onCacheKey = options.onCacheKey || defaultOnCacheKey;
        const key = onCacheKey(source);
        const cached = compileCache[key];
        if (cached) {
            return cached;
        }
        // compile error detecting
        let occurred = false;
        const onError = options.onError || defaultOnError;
        options.onError = (err) => {
            occurred = true;
            onError(err);
        };
        // compile
        const { code } = baseCompile(source, options);
        // evaluate function
        const msg = new Function(`return ${code}`)();
        // if occurred compile error, don't cache
        return !occurred ? (compileCache[key] = msg) : msg;
    }
}

let code = CompileErrorCodes.__EXTEND_POINT__;
const inc = () => ++code;
const CoreErrorCodes = {
    INVALID_ARGUMENT: code,
    INVALID_DATE_ARGUMENT: inc(),
    INVALID_ISO_DATE_ARGUMENT: inc(),
    __EXTEND_POINT__: inc() // 18
};
function createCoreError(code) {
    return createCompileError(code, null, (process.env.NODE_ENV !== 'production') ? { messages: errorMessages } : undefined);
}
/** @internal */
const errorMessages = {
    [CoreErrorCodes.INVALID_ARGUMENT]: 'Invalid arguments',
    [CoreErrorCodes.INVALID_DATE_ARGUMENT]: 'The date provided is an invalid Date object.' +
        'Make sure your Date represents a valid date.',
    [CoreErrorCodes.INVALID_ISO_DATE_ARGUMENT]: 'The argument provided is not a valid ISO date string'
};

const NOOP_MESSAGE_FUNCTION = () => '';
const isMessageFunction = (val) => isFunction(val);
// implementation of `translate` function
function translate(context, ...args) {
    const { fallbackFormat, postTranslation, unresolving, messageCompiler, fallbackLocale, messages } = context;
    const [key, options] = parseTranslateArgs(...args);
    const missingWarn = isBoolean(options.missingWarn)
        ? options.missingWarn
        : context.missingWarn;
    const fallbackWarn = isBoolean(options.fallbackWarn)
        ? options.fallbackWarn
        : context.fallbackWarn;
    const escapeParameter = isBoolean(options.escapeParameter)
        ? options.escapeParameter
        : context.escapeParameter;
    const resolvedMessage = !!options.resolvedMessage;
    // prettier-ignore
    const defaultMsgOrKey = isString(options.default) || isBoolean(options.default) // default by function option
        ? !isBoolean(options.default)
            ? options.default
            : (!messageCompiler ? () => key : key)
        : fallbackFormat // default by `fallbackFormat` option
            ? (!messageCompiler ? () => key : key)
            : '';
    const enableDefaultMsg = fallbackFormat || defaultMsgOrKey !== '';
    const locale = isString(options.locale) ? options.locale : context.locale;
    // escape params
    escapeParameter && escapeParams(options);
    // resolve message format
    // eslint-disable-next-line prefer-const
    let [formatScope, targetLocale, message] = !resolvedMessage
        ? resolveMessageFormat(context, key, locale, fallbackLocale, fallbackWarn, missingWarn)
        : [
            key,
            locale,
            messages[locale] || {}
        ];
    // NOTE:
    //  Fix to work around `ssrTransfrom` bug in Vite.
    //  https://github.com/vitejs/vite/issues/4306
    //  To get around this, use temporary variables.
    //  https://github.com/nuxt/framework/issues/1461#issuecomment-954606243
    let format = formatScope;
    // if you use default message, set it as message format!
    let cacheBaseKey = key;
    if (!resolvedMessage &&
        !(isString(format) || isMessageFunction(format))) {
        if (enableDefaultMsg) {
            format = defaultMsgOrKey;
            cacheBaseKey = format;
        }
    }
    // checking message format and target locale
    if (!resolvedMessage &&
        (!(isString(format) || isMessageFunction(format)) ||
            !isString(targetLocale))) {
        return unresolving ? NOT_REOSLVED : key;
    }
    if ((process.env.NODE_ENV !== 'production') && isString(format) && context.messageCompiler == null) {
        warn(`The message format compilation is not supported in this build. ` +
            `Because message compiler isn't included. ` +
            `You need to pre-compilation all message format. ` +
            `So translate function return '${key}'.`);
        return key;
    }
    // setup compile error detecting
    let occurred = false;
    const errorDetector = () => {
        occurred = true;
    };
    // compile message format
    const msg = !isMessageFunction(format)
        ? compileMessageFormat(context, key, targetLocale, format, cacheBaseKey, errorDetector)
        : format;
    // if occurred compile error, return the message format
    if (occurred) {
        return format;
    }
    // evaluate message with context
    const ctxOptions = getMessageContextOptions(context, targetLocale, message, options);
    const msgContext = createMessageContext(ctxOptions);
    const messaged = evaluateMessage(context, msg, msgContext);
    // if use post translation option, proceed it with handler
    const ret = postTranslation
        ? postTranslation(messaged, key)
        : messaged;
    // NOTE: experimental !!
    if ((process.env.NODE_ENV !== 'production') || __INTLIFY_PROD_DEVTOOLS__) {
        // prettier-ignore
        const payloads = {
            timestamp: Date.now(),
            key: isString(key)
                ? key
                : isMessageFunction(format)
                    ? format.key
                    : '',
            locale: targetLocale || (isMessageFunction(format)
                ? format.locale
                : ''),
            format: isString(format)
                ? format
                : isMessageFunction(format)
                    ? format.source
                    : '',
            message: ret
        };
        payloads.meta = assign({}, context.__meta, getAdditionalMeta() || {});
        translateDevTools(payloads);
    }
    return ret;
}
function escapeParams(options) {
    if (isArray(options.list)) {
        options.list = options.list.map(item => isString(item) ? escapeHtml(item) : item);
    }
    else if (isObject(options.named)) {
        Object.keys(options.named).forEach(key => {
            if (isString(options.named[key])) {
                options.named[key] = escapeHtml(options.named[key]);
            }
        });
    }
}
function resolveMessageFormat(context, key, locale, fallbackLocale, fallbackWarn, missingWarn) {
    const { messages, onWarn, messageResolver: resolveValue, localeFallbacker } = context;
    const locales = localeFallbacker(context, fallbackLocale, locale); // eslint-disable-line @typescript-eslint/no-explicit-any
    let message = {};
    let targetLocale;
    let format = null;
    let from = locale;
    let to = null;
    const type = 'translate';
    for (let i = 0; i < locales.length; i++) {
        targetLocale = to = locales[i];
        if ((process.env.NODE_ENV !== 'production') &&
            locale !== targetLocale &&
            isTranslateFallbackWarn(fallbackWarn, key)) {
            onWarn(getWarnMessage(CoreWarnCodes.FALLBACK_TO_TRANSLATE, {
                key,
                target: targetLocale
            }));
        }
        // for vue-devtools timeline event
        if ((process.env.NODE_ENV !== 'production') && locale !== targetLocale) {
            const emitter = context.__v_emitter;
            if (emitter) {
                emitter.emit("fallback" /* FALBACK */, {
                    type,
                    key,
                    from,
                    to,
                    groupId: `${type}:${key}`
                });
            }
        }
        message =
            messages[targetLocale] || {};
        // for vue-devtools timeline event
        let start = null;
        let startTag;
        let endTag;
        if ((process.env.NODE_ENV !== 'production') && inBrowser) {
            start = window.performance.now();
            startTag = 'intlify-message-resolve-start';
            endTag = 'intlify-message-resolve-end';
            mark && mark(startTag);
        }
        if ((format = resolveValue(message, key)) === null) {
            // if null, resolve with object key path
            format = message[key]; // eslint-disable-line @typescript-eslint/no-explicit-any
        }
        // for vue-devtools timeline event
        if ((process.env.NODE_ENV !== 'production') && inBrowser) {
            const end = window.performance.now();
            const emitter = context.__v_emitter;
            if (emitter && start && format) {
                emitter.emit("message-resolve" /* MESSAGE_RESOLVE */, {
                    type: "message-resolve" /* MESSAGE_RESOLVE */,
                    key,
                    message: format,
                    time: end - start,
                    groupId: `${type}:${key}`
                });
            }
            if (startTag && endTag && mark && measure) {
                mark(endTag);
                measure('intlify message resolve', startTag, endTag);
            }
        }
        if (isString(format) || isFunction(format))
            break;
        const missingRet = handleMissing(context, // eslint-disable-line @typescript-eslint/no-explicit-any
        key, targetLocale, missingWarn, type);
        if (missingRet !== key) {
            format = missingRet;
        }
        from = to;
    }
    return [format, targetLocale, message];
}
function compileMessageFormat(context, key, targetLocale, format, cacheBaseKey, errorDetector) {
    const { messageCompiler, warnHtmlMessage } = context;
    if (isMessageFunction(format)) {
        const msg = format;
        msg.locale = msg.locale || targetLocale;
        msg.key = msg.key || key;
        return msg;
    }
    if (messageCompiler == null) {
        const msg = (() => format);
        msg.locale = targetLocale;
        msg.key = key;
        return msg;
    }
    // for vue-devtools timeline event
    let start = null;
    let startTag;
    let endTag;
    if ((process.env.NODE_ENV !== 'production') && inBrowser) {
        start = window.performance.now();
        startTag = 'intlify-message-compilation-start';
        endTag = 'intlify-message-compilation-end';
        mark && mark(startTag);
    }
    const msg = messageCompiler(format, getCompileOptions(context, targetLocale, cacheBaseKey, format, warnHtmlMessage, errorDetector));
    // for vue-devtools timeline event
    if ((process.env.NODE_ENV !== 'production') && inBrowser) {
        const end = window.performance.now();
        const emitter = context.__v_emitter;
        if (emitter && start) {
            emitter.emit("message-compilation" /* MESSAGE_COMPILATION */, {
                type: "message-compilation" /* MESSAGE_COMPILATION */,
                message: format,
                time: end - start,
                groupId: `${'translate'}:${key}`
            });
        }
        if (startTag && endTag && mark && measure) {
            mark(endTag);
            measure('intlify message compilation', startTag, endTag);
        }
    }
    msg.locale = targetLocale;
    msg.key = key;
    msg.source = format;
    return msg;
}
function evaluateMessage(context, msg, msgCtx) {
    // for vue-devtools timeline event
    let start = null;
    let startTag;
    let endTag;
    if ((process.env.NODE_ENV !== 'production') && inBrowser) {
        start = window.performance.now();
        startTag = 'intlify-message-evaluation-start';
        endTag = 'intlify-message-evaluation-end';
        mark && mark(startTag);
    }
    const messaged = msg(msgCtx);
    // for vue-devtools timeline event
    if ((process.env.NODE_ENV !== 'production') && inBrowser) {
        const end = window.performance.now();
        const emitter = context.__v_emitter;
        if (emitter && start) {
            emitter.emit("message-evaluation" /* MESSAGE_EVALUATION */, {
                type: "message-evaluation" /* MESSAGE_EVALUATION */,
                value: messaged,
                time: end - start,
                groupId: `${'translate'}:${msg.key}`
            });
        }
        if (startTag && endTag && mark && measure) {
            mark(endTag);
            measure('intlify message evaluation', startTag, endTag);
        }
    }
    return messaged;
}
/** @internal */
function parseTranslateArgs(...args) {
    const [arg1, arg2, arg3] = args;
    const options = {};
    if (!isString(arg1) && !isNumber(arg1) && !isMessageFunction(arg1)) {
        throw createCoreError(CoreErrorCodes.INVALID_ARGUMENT);
    }
    // prettier-ignore
    const key = isNumber(arg1)
        ? String(arg1)
        : isMessageFunction(arg1)
            ? arg1
            : arg1;
    if (isNumber(arg2)) {
        options.plural = arg2;
    }
    else if (isString(arg2)) {
        options.default = arg2;
    }
    else if (isPlainObject(arg2) && !isEmptyObject(arg2)) {
        options.named = arg2;
    }
    else if (isArray(arg2)) {
        options.list = arg2;
    }
    if (isNumber(arg3)) {
        options.plural = arg3;
    }
    else if (isString(arg3)) {
        options.default = arg3;
    }
    else if (isPlainObject(arg3)) {
        assign(options, arg3);
    }
    return [key, options];
}
function getCompileOptions(context, locale, key, source, warnHtmlMessage, errorDetector) {
    return {
        warnHtmlMessage,
        onError: (err) => {
            errorDetector && errorDetector(err);
            if ((process.env.NODE_ENV !== 'production')) {
                const message = `Message compilation error: ${err.message}`;
                const codeFrame = err.location &&
                    generateCodeFrame(source, err.location.start.offset, err.location.end.offset);
                const emitter = context.__v_emitter;
                if (emitter) {
                    emitter.emit("compile-error" /* COMPILE_ERROR */, {
                        message: source,
                        error: err.message,
                        start: err.location && err.location.start.offset,
                        end: err.location && err.location.end.offset,
                        groupId: `${'translate'}:${key}`
                    });
                }
                console.error(codeFrame ? `${message}\n${codeFrame}` : message);
            }
            else {
                throw err;
            }
        },
        onCacheKey: (source) => generateFormatCacheKey(locale, key, source)
    };
}
function getMessageContextOptions(context, locale, message, options) {
    const { modifiers, pluralRules, messageResolver: resolveValue, fallbackLocale, fallbackWarn, missingWarn, fallbackContext } = context;
    const resolveMessage = (key) => {
        let val = resolveValue(message, key);
        // fallback to root context
        if (val == null && fallbackContext) {
            const [, , message] = resolveMessageFormat(fallbackContext, key, locale, fallbackLocale, fallbackWarn, missingWarn);
            val = resolveValue(message, key);
        }
        if (isString(val)) {
            let occurred = false;
            const errorDetector = () => {
                occurred = true;
            };
            const msg = compileMessageFormat(context, key, locale, val, key, errorDetector);
            return !occurred
                ? msg
                : NOOP_MESSAGE_FUNCTION;
        }
        else if (isMessageFunction(val)) {
            return val;
        }
        else {
            // TODO: should be implemented warning message
            return NOOP_MESSAGE_FUNCTION;
        }
    };
    const ctxOptions = {
        locale,
        modifiers,
        pluralRules,
        messages: resolveMessage
    };
    if (context.processor) {
        ctxOptions.processor = context.processor;
    }
    if (options.list) {
        ctxOptions.list = options.list;
    }
    if (options.named) {
        ctxOptions.named = options.named;
    }
    if (isNumber(options.plural)) {
        ctxOptions.pluralIndex = options.plural;
    }
    return ctxOptions;
}

const intlDefined = typeof Intl !== 'undefined';
const Availabilities = {
    dateTimeFormat: intlDefined && typeof Intl.DateTimeFormat !== 'undefined',
    numberFormat: intlDefined && typeof Intl.NumberFormat !== 'undefined'
};

// implementation of `datetime` function
function datetime(context, ...args) {
    const { datetimeFormats, unresolving, fallbackLocale, onWarn, localeFallbacker } = context;
    const { __datetimeFormatters } = context;
    if ((process.env.NODE_ENV !== 'production') && !Availabilities.dateTimeFormat) {
        onWarn(getWarnMessage(CoreWarnCodes.CANNOT_FORMAT_DATE));
        return MISSING_RESOLVE_VALUE;
    }
    const [key, value, options, overrides] = parseDateTimeArgs(...args);
    const missingWarn = isBoolean(options.missingWarn)
        ? options.missingWarn
        : context.missingWarn;
    const fallbackWarn = isBoolean(options.fallbackWarn)
        ? options.fallbackWarn
        : context.fallbackWarn;
    const part = !!options.part;
    const locale = isString(options.locale) ? options.locale : context.locale;
    const locales = localeFallbacker(context, // eslint-disable-line @typescript-eslint/no-explicit-any
    fallbackLocale, locale);
    if (!isString(key) || key === '') {
        return new Intl.DateTimeFormat(locale, overrides).format(value);
    }
    // resolve format
    let datetimeFormat = {};
    let targetLocale;
    let format = null;
    let from = locale;
    let to = null;
    const type = 'datetime format';
    for (let i = 0; i < locales.length; i++) {
        targetLocale = to = locales[i];
        if ((process.env.NODE_ENV !== 'production') &&
            locale !== targetLocale &&
            isTranslateFallbackWarn(fallbackWarn, key)) {
            onWarn(getWarnMessage(CoreWarnCodes.FALLBACK_TO_DATE_FORMAT, {
                key,
                target: targetLocale
            }));
        }
        // for vue-devtools timeline event
        if ((process.env.NODE_ENV !== 'production') && locale !== targetLocale) {
            const emitter = context.__v_emitter;
            if (emitter) {
                emitter.emit("fallback" /* FALBACK */, {
                    type,
                    key,
                    from,
                    to,
                    groupId: `${type}:${key}`
                });
            }
        }
        datetimeFormat =
            datetimeFormats[targetLocale] || {};
        format = datetimeFormat[key];
        if (isPlainObject(format))
            break;
        handleMissing(context, key, targetLocale, missingWarn, type); // eslint-disable-line @typescript-eslint/no-explicit-any
        from = to;
    }
    // checking format and target locale
    if (!isPlainObject(format) || !isString(targetLocale)) {
        return unresolving ? NOT_REOSLVED : key;
    }
    let id = `${targetLocale}__${key}`;
    if (!isEmptyObject(overrides)) {
        id = `${id}__${JSON.stringify(overrides)}`;
    }
    let formatter = __datetimeFormatters.get(id);
    if (!formatter) {
        formatter = new Intl.DateTimeFormat(targetLocale, assign({}, format, overrides));
        __datetimeFormatters.set(id, formatter);
    }
    return !part ? formatter.format(value) : formatter.formatToParts(value);
}
/** @internal */
const DATETIME_FORMAT_OPTIONS_KEYS = [
    'localeMatcher',
    'weekday',
    'era',
    'year',
    'month',
    'day',
    'hour',
    'minute',
    'second',
    'timeZoneName',
    'formatMatcher',
    'hour12',
    'timeZone',
    'dateStyle',
    'timeStyle',
    'calendar',
    'dayPeriod',
    'numberingSystem',
    'hourCycle',
    'fractionalSecondDigits'
];
/** @internal */
function parseDateTimeArgs(...args) {
    const [arg1, arg2, arg3, arg4] = args;
    const options = {};
    let overrides = {};
    let value;
    if (isString(arg1)) {
        // Only allow ISO strings - other date formats are often supported,
        // but may cause different results in different browsers.
        const matches = arg1.match(/(\d{4}-\d{2}-\d{2})(T|\s)?(.*)/);
        if (!matches) {
            throw createCoreError(CoreErrorCodes.INVALID_ISO_DATE_ARGUMENT);
        }
        // Some browsers can not parse the iso datetime separated by space,
        // this is a compromise solution by replace the 'T'/' ' with 'T'
        const dateTime = matches[3]
            ? matches[3].trim().startsWith('T')
                ? `${matches[1].trim()}${matches[3].trim()}`
                : `${matches[1].trim()}T${matches[3].trim()}`
            : matches[1].trim();
        value = new Date(dateTime);
        try {
            // This will fail if the date is not valid
            value.toISOString();
        }
        catch (e) {
            throw createCoreError(CoreErrorCodes.INVALID_ISO_DATE_ARGUMENT);
        }
    }
    else if (isDate(arg1)) {
        if (isNaN(arg1.getTime())) {
            throw createCoreError(CoreErrorCodes.INVALID_DATE_ARGUMENT);
        }
        value = arg1;
    }
    else if (isNumber(arg1)) {
        value = arg1;
    }
    else {
        throw createCoreError(CoreErrorCodes.INVALID_ARGUMENT);
    }
    if (isString(arg2)) {
        options.key = arg2;
    }
    else if (isPlainObject(arg2)) {
        Object.keys(arg2).forEach(key => {
            if (DATETIME_FORMAT_OPTIONS_KEYS.includes(key)) {
                overrides[key] = arg2[key];
            }
            else {
                options[key] = arg2[key];
            }
        });
    }
    if (isString(arg3)) {
        options.locale = arg3;
    }
    else if (isPlainObject(arg3)) {
        overrides = arg3;
    }
    if (isPlainObject(arg4)) {
        overrides = arg4;
    }
    return [options.key || '', value, options, overrides];
}
/** @internal */
function clearDateTimeFormat(ctx, locale, format) {
    const context = ctx;
    for (const key in format) {
        const id = `${locale}__${key}`;
        if (!context.__datetimeFormatters.has(id)) {
            continue;
        }
        context.__datetimeFormatters.delete(id);
    }
}

// implementation of `number` function
function number(context, ...args) {
    const { numberFormats, unresolving, fallbackLocale, onWarn, localeFallbacker } = context;
    const { __numberFormatters } = context;
    if ((process.env.NODE_ENV !== 'production') && !Availabilities.numberFormat) {
        onWarn(getWarnMessage(CoreWarnCodes.CANNOT_FORMAT_NUMBER));
        return MISSING_RESOLVE_VALUE;
    }
    const [key, value, options, overrides] = parseNumberArgs(...args);
    const missingWarn = isBoolean(options.missingWarn)
        ? options.missingWarn
        : context.missingWarn;
    const fallbackWarn = isBoolean(options.fallbackWarn)
        ? options.fallbackWarn
        : context.fallbackWarn;
    const part = !!options.part;
    const locale = isString(options.locale) ? options.locale : context.locale;
    const locales = localeFallbacker(context, // eslint-disable-line @typescript-eslint/no-explicit-any
    fallbackLocale, locale);
    if (!isString(key) || key === '') {
        return new Intl.NumberFormat(locale, overrides).format(value);
    }
    // resolve format
    let numberFormat = {};
    let targetLocale;
    let format = null;
    let from = locale;
    let to = null;
    const type = 'number format';
    for (let i = 0; i < locales.length; i++) {
        targetLocale = to = locales[i];
        if ((process.env.NODE_ENV !== 'production') &&
            locale !== targetLocale &&
            isTranslateFallbackWarn(fallbackWarn, key)) {
            onWarn(getWarnMessage(CoreWarnCodes.FALLBACK_TO_NUMBER_FORMAT, {
                key,
                target: targetLocale
            }));
        }
        // for vue-devtools timeline event
        if ((process.env.NODE_ENV !== 'production') && locale !== targetLocale) {
            const emitter = context.__v_emitter;
            if (emitter) {
                emitter.emit("fallback" /* FALBACK */, {
                    type,
                    key,
                    from,
                    to,
                    groupId: `${type}:${key}`
                });
            }
        }
        numberFormat =
            numberFormats[targetLocale] || {};
        format = numberFormat[key];
        if (isPlainObject(format))
            break;
        handleMissing(context, key, targetLocale, missingWarn, type); // eslint-disable-line @typescript-eslint/no-explicit-any
        from = to;
    }
    // checking format and target locale
    if (!isPlainObject(format) || !isString(targetLocale)) {
        return unresolving ? NOT_REOSLVED : key;
    }
    let id = `${targetLocale}__${key}`;
    if (!isEmptyObject(overrides)) {
        id = `${id}__${JSON.stringify(overrides)}`;
    }
    let formatter = __numberFormatters.get(id);
    if (!formatter) {
        formatter = new Intl.NumberFormat(targetLocale, assign({}, format, overrides));
        __numberFormatters.set(id, formatter);
    }
    return !part ? formatter.format(value) : formatter.formatToParts(value);
}
/** @internal */
const NUMBER_FORMAT_OPTIONS_KEYS = [
    'localeMatcher',
    'style',
    'currency',
    'currencyDisplay',
    'currencySign',
    'useGrouping',
    'minimumIntegerDigits',
    'minimumFractionDigits',
    'maximumFractionDigits',
    'minimumSignificantDigits',
    'maximumSignificantDigits',
    'compactDisplay',
    'notation',
    'signDisplay',
    'unit',
    'unitDisplay',
    'roundingMode',
    'roundingPriority',
    'roundingIncrement',
    'trailingZeroDisplay'
];
/** @internal */
function parseNumberArgs(...args) {
    const [arg1, arg2, arg3, arg4] = args;
    const options = {};
    let overrides = {};
    if (!isNumber(arg1)) {
        throw createCoreError(CoreErrorCodes.INVALID_ARGUMENT);
    }
    const value = arg1;
    if (isString(arg2)) {
        options.key = arg2;
    }
    else if (isPlainObject(arg2)) {
        Object.keys(arg2).forEach(key => {
            if (NUMBER_FORMAT_OPTIONS_KEYS.includes(key)) {
                overrides[key] = arg2[key];
            }
            else {
                options[key] = arg2[key];
            }
        });
    }
    if (isString(arg3)) {
        options.locale = arg3;
    }
    else if (isPlainObject(arg3)) {
        overrides = arg3;
    }
    if (isPlainObject(arg4)) {
        overrides = arg4;
    }
    return [options.key || '', value, options, overrides];
}
/** @internal */
function clearNumberFormat(ctx, locale, format) {
    const context = ctx;
    for (const key in format) {
        const id = `${locale}__${key}`;
        if (!context.__numberFormatters.has(id)) {
            continue;
        }
        context.__numberFormatters.delete(id);
    }
}

// TODO: we could not exports for Node native ES Moudles yet...
{
    if (typeof __INTLIFY_PROD_DEVTOOLS__ !== 'boolean') {
        getGlobalThis().__INTLIFY_PROD_DEVTOOLS__ = false;
    }
}

export { CoreErrorCodes, CoreWarnCodes, DATETIME_FORMAT_OPTIONS_KEYS, DEFAULT_LOCALE, DEFAULT_MESSAGE_DATA_TYPE, MISSING_RESOLVE_VALUE, NOT_REOSLVED, NUMBER_FORMAT_OPTIONS_KEYS, VERSION, clearCompileCache, clearDateTimeFormat, clearNumberFormat, compileToFunction, createCoreContext, createCoreError, createMessageContext, datetime, fallbackWithLocaleChain, fallbackWithSimple, getAdditionalMeta, getDevToolsHook, getFallbackContext, getWarnMessage, handleMissing, initI18nDevTools, isMessageFunction, isTranslateFallbackWarn, isTranslateMissingWarn, number, parse, parseDateTimeArgs, parseNumberArgs, parseTranslateArgs, registerLocaleFallbacker, registerMessageCompiler, registerMessageResolver, resolveValue, resolveWithKeyValue, setAdditionalMeta, setDevToolsHook, setFallbackContext, translate, translateDevTools, updateFallbackLocale };
