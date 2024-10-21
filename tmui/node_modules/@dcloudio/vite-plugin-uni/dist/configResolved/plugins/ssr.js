"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniSSRPlugin = void 0;
const debug_1 = __importDefault(require("debug"));
const crypto_1 = __importDefault(require("crypto"));
const estree_walker_1 = require("estree-walker");
const pluginutils_1 = require("@rollup/pluginutils");
const magic_string_1 = __importDefault(require("magic-string"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const debugSSR = (0, debug_1.default)('uni:ssr');
const KEYED_FUNC_RE = /(ssrRef|shallowSsrRef)/;
function uniSSRPlugin(config, options) {
    const filter = (0, pluginutils_1.createFilter)(options.include, options.exclude);
    return {
        name: 'uni:ssr:ref',
        transform(code, id) {
            if (!filter(id))
                return null;
            if (!KEYED_FUNC_RE.test(code)) {
                return;
            }
            debugSSR('try', id);
            const ast = this.parse(code);
            const s = new magic_string_1.default(code);
            (0, estree_walker_1.walk)(ast, {
                enter(node) {
                    if (!(0, uni_cli_shared_1.isCallExpression)(node)) {
                        return;
                    }
                    const { callee, arguments: args } = node;
                    if (args.length !== 1) {
                        return;
                    }
                    const name = (0, uni_cli_shared_1.isIdentifier)(callee)
                        ? callee.name
                        : (0, uni_cli_shared_1.isMemberExpression)(callee) && (0, uni_cli_shared_1.isIdentifier)(callee.property)
                            ? callee.property.name
                            : '';
                    if (name !== 'ssrRef' && name !== 'shallowSsrRef') {
                        return;
                    }
                    const { end } = node;
                    const key = id + '-' + node.end;
                    debugSSR(key, name);
                    s.appendLeft(end - 1, ", '" + createKey(`${id}-${end}`) + "'");
                },
            });
            return {
                code: s.toString(),
                map: (0, uni_cli_shared_1.withSourcemap)(config) ? s.generateMap().toString() : null,
            };
        },
    };
}
exports.uniSSRPlugin = uniSSRPlugin;
function createKey(source) {
    const hash = crypto_1.default.createHash('md5');
    hash.update(source);
    return hash.digest('base64').toString();
}
