"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.initUniApp = void 0;
const nvue_1 = require("./nvue");
function initUniApp(manifestJson) {
    manifestJson.plus['uni-app'] = {
        control: 'uni-v3',
        vueVersion: '3',
        compilerVersion: process.env.UNI_COMPILER_VERSION,
        nvueCompiler: (0, nvue_1.getNVueCompiler)(manifestJson),
        renderer: 'auto',
        nvue: {
            'flex-direction': (0, nvue_1.getNVueFlexDirection)(manifestJson),
        },
        nvueLaunchMode: manifestJson.plus.nvueLaunchMode === 'fast' ? 'fast' : 'normal',
    };
    delete manifestJson.plus.nvueLaunchMode;
}
exports.initUniApp = initUniApp;
