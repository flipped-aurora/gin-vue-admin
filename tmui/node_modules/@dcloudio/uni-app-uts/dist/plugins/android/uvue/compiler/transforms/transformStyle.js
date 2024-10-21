"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformStyle = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const uni_nvue_styler_1 = require("@dcloudio/uni-nvue-styler");
const errors_1 = require("../errors");
const shared_1 = require("@vue/shared");
// Verify that the template style is in compliance with specifications
const transformStyle = (node, context) => {
    if (node.type === compiler_core_1.NodeTypes.ELEMENT) {
        node.props.forEach((p, i) => {
            if (p.type === compiler_core_1.NodeTypes.ATTRIBUTE && p.name === 'style' && p.value) {
                // 静态 style 编译成对象，减少运行时解析
                const styleObjectStr = parseStyleString2ObjectString(p.value.content);
                node.props[i] = createVBindStyleExp(styleObjectStr, p.loc);
                (0, uni_nvue_styler_1.parse)(p.value.content, {
                    logLevel: 'WARNING',
                    map: true,
                    ts: true,
                    noCode: true,
                    type: 'uvue',
                    platform: process.env.UNI_UTS_PLATFORM,
                }).then(({ messages }) => {
                    messages.forEach((message) => {
                        context.onWarn((0, errors_1.createCompilerError)(100, p.loc, {
                            100: message.text,
                        }));
                    });
                });
            }
        });
    }
};
exports.transformStyle = transformStyle;
function parseStyleString2ObjectString(styleString) {
    const styleObject = (0, shared_1.parseStringStyle)(styleString);
    return JSON.stringify(styleObject);
}
function createVBindStyleExp(styleObjectStr, loc) {
    return {
        type: compiler_core_1.NodeTypes.DIRECTIVE,
        name: 'bind',
        arg: {
            type: compiler_core_1.NodeTypes.SIMPLE_EXPRESSION,
            content: 'style',
            isStatic: true,
            constType: 3,
            loc: {
                end: loc.end,
                start: loc.start,
                source: 'style',
            },
        },
        exp: {
            constType: 0,
            content: styleObjectStr,
            isStatic: false,
            type: compiler_core_1.NodeTypes.SIMPLE_EXPRESSION,
            loc: {
                end: loc.end,
                start: loc.start,
                source: styleObjectStr,
            },
        },
        modifiers: [],
        loc: {
            end: loc.end,
            start: loc.start,
            source: `:style="${styleObjectStr}"`,
        },
    };
}
