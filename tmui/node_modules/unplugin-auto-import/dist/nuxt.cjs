"use strict";Object.defineProperty(exports, "__esModule", {value: true});

var _chunkERF3N54Scjs = require('./chunk-ERF3N54S.cjs');
require('./chunk-6UOGCAOY.cjs');

// src/nuxt.ts
var _kit = require('@nuxt/kit');
var nuxt_default = _kit.defineNuxtModule.call(void 0, {
  setup(options) {
    options.exclude = options.exclude || [/[\\/]node_modules[\\/]/, /[\\/]\.git[\\/]/, /[\\/]\.nuxt[\\/]/];
    _kit.addWebpackPlugin.call(void 0, _chunkERF3N54Scjs.unplugin_default.webpack(options));
    _kit.addVitePlugin.call(void 0, _chunkERF3N54Scjs.unplugin_default.vite(options));
  }
});


module.exports = nuxt_default;
exports.default = module.exports;