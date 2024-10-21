"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createNonAppGenerateBundle = exports.uniSubpackagePlugin = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function uniSubpackagePlugin({ style: { extname }, }) {
    return {
        name: 'uni:mp-subpackage',
        enforce: 'post',
        generateBundle: createNonAppGenerateBundle(extname),
    };
}
exports.uniSubpackagePlugin = uniSubpackagePlugin;
function createNonAppGenerateBundle(extname) {
    return function generateBundle(_, bundle) {
        ;
        ['project.config.json', 'app.json'].forEach((name) => {
            delete bundle[name];
        });
        const appJsFile = 'app.js';
        const appCssFile = 'app' + extname;
        Object.keys(bundle).forEach((name) => {
            if (!(0, uni_cli_shared_1.isMiniProgramPageFile)(name)) {
                return;
            }
            // 仅页面级 wxss 需要补充 app.wxss
            if (name.endsWith(extname)) {
                const cssFile = bundle[name];
                cssFile.source =
                    `@import "${(0, uni_cli_shared_1.relativeFile)(name, appCssFile)}";\n` +
                        cssFile.source.toString();
            }
            else if (name.endsWith('.js')) {
                const jsFile = bundle[name];
                jsFile.code =
                    `require('${(0, uni_cli_shared_1.relativeFile)(name, appJsFile)}');\n` + jsFile.code;
            }
        });
    };
}
exports.createNonAppGenerateBundle = createNonAppGenerateBundle;
