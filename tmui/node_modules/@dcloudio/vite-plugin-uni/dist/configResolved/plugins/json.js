"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniJsonPlugin = void 0;
const jsonc_parser_1 = require("jsonc-parser");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const jsonExtRE = /\.json($|\?)(?!commonjs-proxy)/;
const SPECIAL_QUERY_RE = /[\?&](?:worker|sharedworker|raw|url)\b/;
function uniJsonPlugin(options) {
    return {
        name: 'uni:json',
        transform(code, id) {
            if (!jsonExtRE.test(id))
                return null;
            if (SPECIAL_QUERY_RE.test(id))
                return null;
            if (id.endsWith('.json.js'))
                return null;
            // preprocess
            if (code.includes('#endif')) {
                code = (0, uni_cli_shared_1.preJs)(code);
            }
            let jsonObj = (0, jsonc_parser_1.parse)(code);
            if ((0, uni_cli_shared_1.isUniAppLocaleFile)(id)) {
                jsonObj = jsonObj.common || {};
            }
            return {
                code: JSON.stringify(jsonObj),
                map: null,
            };
        },
    };
}
exports.uniJsonPlugin = uniJsonPlugin;
