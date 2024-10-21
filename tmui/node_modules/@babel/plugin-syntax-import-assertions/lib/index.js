"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = void 0;
var _helperPluginUtils = require("@babel/helper-plugin-utils");
var _default = exports.default = (0, _helperPluginUtils.declare)(api => {
  api.assertVersion(7);
  return {
    name: "syntax-import-assertions",
    manipulateOptions(opts, {
      plugins
    }) {
      for (let i = 0; i < plugins.length; i++) {
        const plugin = plugins[i];
        if (plugin === "importAttributes") {
          plugins[i] = ["importAttributes", {
            deprecatedAssertSyntax: true
          }];
          return;
        }
        if (Array.isArray(plugin) && plugin[0] === "importAttributes") {
          if (plugin.length < 2) plugins[i].push({});
          plugin[1].deprecatedAssertSyntax = true;
          return;
        }
      }
      plugins.push("importAssertions");
    }
  };
});

//# sourceMappingURL=index.js.map
