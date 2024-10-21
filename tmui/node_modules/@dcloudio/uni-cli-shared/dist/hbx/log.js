"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.errorFormatter = exports.removeWarnFormatter = exports.removeInfoFormatter = exports.h5ServeFormatter = exports.formatAtFilename = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const picocolors_1 = __importDefault(require("picocolors"));
const compiler_core_1 = require("@vue/compiler-core");
const shared_1 = require("@vue/shared");
const utils_1 = require("../utils");
const constants_1 = require("../constants");
const ast_1 = require("../vite/utils/ast");
const utils_2 = require("../vite/plugins/vitejs/utils");
const SIGNAL_H5_LOCAL = ' ➜  Local:';
const SIGNAL_H5_NETWORK = ' ➜  Network:';
const networkLogs = [];
const ZERO_WIDTH_CHAR = {
    NOTE: '',
    WARNING: '\u200B',
    ERROR: '\u200C',
    backup0: '\u200D',
    backup1: '\u200E',
    backup2: '\u200F',
    backup3: '\uFEFF',
};
function overridedConsole(name, oldFn, char) {
    console[name] = function (...args) {
        oldFn.apply(this, args.map((arg) => {
            let item;
            if (typeof arg !== 'object') {
                item = `${char}${arg}${char}`;
            }
            else {
                item = `${char}${JSON.stringify(arg)}${char}`;
            }
            return item;
        }));
    };
}
if (typeof console !== 'undefined') {
    overridedConsole('warn', console.log, ZERO_WIDTH_CHAR.WARNING);
    // overridedConsole('error', console.error, ZERO_WIDTH_CHAR.ERROR)
}
function formatAtFilename(filename, line, column) {
    const file = path_1.default.relative(process.env.UNI_INPUT_DIR, filename.replace('\x00', '').split('?')[0]);
    return `at ${picocolors_1.default.cyan((0, utils_1.normalizePath)(file === 'pages-json-uts' ? 'pages.json' : file) +
        ':' +
        (line || 1) +
        ':' +
        (column || 0))}`;
}
exports.formatAtFilename = formatAtFilename;
exports.h5ServeFormatter = {
    test(msg) {
        return msg.includes(SIGNAL_H5_LOCAL) || msg.includes(SIGNAL_H5_NETWORK);
    },
    format(msg) {
        if (msg.includes(SIGNAL_H5_NETWORK)) {
            networkLogs.push(msg.replace('➜ ', '*'));
            process.nextTick(() => {
                if (networkLogs.length) {
                    // 延迟打印所有 network,仅最后一个 network 替换 ➜ 为 -，通知 hbx
                    const len = networkLogs.length - 1;
                    networkLogs[len] = networkLogs[len].replace('* Network', '- Network');
                    console.log(networkLogs.join('\n'));
                    networkLogs.length = 0;
                }
            });
            return '';
        }
        if (msg.includes(SIGNAL_H5_LOCAL)) {
            return msg.replace('➜ ', '-');
        }
        return msg.replace('➜ ', '*');
    },
};
const REMOVED_MSGS = [
    'build started...',
    (msg) => {
        return /built in [0-9]+ms\./.test(msg);
    },
    'watching for file changes...',
];
exports.removeInfoFormatter = {
    test(msg) {
        return !!REMOVED_MSGS.find((m) => ((0, shared_1.isString)(m) ? msg.includes(m) : m(msg)));
    },
    format() {
        return '';
    },
};
const REMOVED_WARN_MSGS = [];
exports.removeWarnFormatter = {
    test(msg) {
        return !!REMOVED_WARN_MSGS.find((m) => msg.includes(m));
    },
    format() {
        return '';
    },
};
exports.errorFormatter = {
    test(_, opts) {
        return !!(opts && opts.error);
    },
    format(_, opts) {
        return buildErrorMessage(opts.error, [], false);
    },
};
function buildErrorMessage(err, args = [], includeStack = true) {
    if (err.plugin) {
        const otherMsgs = [];
        if (err.message.includes(': [plugin ')) {
            const messages = err.message.split(': [plugin ');
            err.message = messages[0];
            messages.slice(1).forEach((msg) => {
                otherMsgs.push(`[plugin:${msg}`);
            });
        }
        args.push(`${picocolors_1.default.magenta('[plugin:' + err.plugin + ']')} ${picocolors_1.default.red(err.message)}`);
        args.push(...otherMsgs);
        if (err.loc &&
            err.hook === 'transform' &&
            err.plugin === 'rollup-plugin-dynamic-import-variables' &&
            err.id &&
            constants_1.EXTNAME_VUE_RE.test(err.id)) {
            try {
                const ast = (0, ast_1.parseVue)(fs_1.default.readFileSync(err.id, 'utf8'), []);
                const scriptNode = ast.children.find((node) => node.type === compiler_core_1.NodeTypes.ELEMENT && node.tag === 'script');
                if (scriptNode) {
                    const scriptLoc = scriptNode.loc;
                    args.push(picocolors_1.default.yellow(pad((0, utils_2.generateCodeFrame)(scriptLoc.source, err.loc))));
                    // correct error location
                    err.loc.line = scriptLoc.start.line + err.loc.line - 1;
                }
            }
            catch (e) { }
        }
    }
    else {
        args.push(picocolors_1.default.red(err.message));
    }
    if (err.id) {
        args.push(formatAtFilename(err.id, err.loc?.line, err.loc?.column));
    }
    if (err.frame) {
        args.push(picocolors_1.default.yellow(pad(err.frame)));
    }
    if (includeStack && err.stack) {
        args.push(pad(cleanStack(err.stack)));
    }
    return args.join('\n');
}
function cleanStack(stack) {
    return stack
        .split(/\n/g)
        .filter((l) => /^\s*at/.test(l))
        .join('\n');
}
const splitRE = /\r?\n/;
function pad(source, n = 2) {
    const lines = source.split(splitRE);
    return lines.map((l) => ` `.repeat(n) + l).join(`\n`);
}
