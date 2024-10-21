"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniH5Plugin = void 0;
const handleHotUpdate_1 = require("./handleHotUpdate");
const transformIndexHtml_1 = require("./transformIndexHtml");
const configureServer_1 = require("./configureServer");
const uni_1 = require("./uni");
const config_1 = require("./config");
function uniH5Plugin() {
    const configOptions = {
        resolvedConfig: null,
    };
    return {
        name: 'uni:h5',
        uni: (0, uni_1.createUni)(),
        config: (0, config_1.createConfig)(configOptions),
        configResolved(config) {
            configOptions.resolvedConfig = config;
        },
        configureServer: (0, configureServer_1.createConfigureServer)(),
        handleHotUpdate: (0, handleHotUpdate_1.createHandleHotUpdate)(),
        transformIndexHtml: (0, transformIndexHtml_1.createTransformIndexHtml)(),
    };
}
exports.uniH5Plugin = uniH5Plugin;
