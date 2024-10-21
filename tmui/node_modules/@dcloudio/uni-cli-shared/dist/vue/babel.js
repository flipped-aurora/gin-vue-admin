"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformUniH5Jsx = void 0;
const uni_shared_1 = require("@dcloudio/uni-shared");
function transformUniH5Jsx({ types }) {
    return {
        name: 'babel-plugin-uni-h5-jsx',
        visitor: {
            JSXOpeningElement({ node: { name } }) {
                if (types.isJSXIdentifier(name) && (0, uni_shared_1.isBuiltInComponent)(name.name)) {
                    name.name = 'v-uni-' + name.name;
                }
            },
        },
    };
}
exports.transformUniH5Jsx = transformUniH5Jsx;
