"use strict";Object.defineProperty(exports, "__esModule", {value: true});

var _chunk23WSNQH3cjs = require('./chunk-23WSNQH3.cjs');
require('./chunk-6IWOTYMJ.cjs');
require('./chunk-OBGZSXTJ.cjs');
require('./chunk-ZBPRDZS4.cjs');

// src/nuxt.ts
var _kit = require('@nuxt/kit');
var nuxt_default = _kit.defineNuxtModule.call(void 0, {
  setup(options) {
    _kit.addWebpackPlugin.call(void 0, _chunk23WSNQH3cjs.unplugin_default.webpack(options));
    _kit.addVitePlugin.call(void 0, _chunk23WSNQH3cjs.unplugin_default.vite(options));
  }
});


exports.default = nuxt_default;
