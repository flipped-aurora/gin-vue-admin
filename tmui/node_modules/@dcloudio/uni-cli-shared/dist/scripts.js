"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseScripts = void 0;
const shared_1 = require("@vue/shared");
const fs_1 = __importDefault(require("fs"));
function parseScripts(name, pkgPath) {
    if (!fs_1.default.existsSync(pkgPath)) {
        return;
    }
    const pkg = JSON.parse(fs_1.default.readFileSync(pkgPath, 'utf8'));
    const scripts = pkg['uni-app']?.scripts || {};
    const options = scripts[name];
    if (!options) {
        return;
    }
    if (!options.env?.UNI_PLATFORM) {
        console.error(`package.json->uni-app->scripts->${name}->env->UNI_PLATFORM is required`);
        process.exit(0);
    }
    const { UNI_PLATFORM, ...define } = options.env;
    const context = options.define || {};
    // 补充当前编译环境未定义的其他编译环境 define
    Object.keys(scripts).forEach((scriptName) => {
        if (scriptName !== name) {
            const scriptDefine = scripts[scriptName].define || {};
            Object.keys(scriptDefine).forEach((key) => {
                if (!(0, shared_1.hasOwn)(context, key)) {
                    context[key] = false;
                }
            });
        }
    });
    return {
        name: name,
        platform: UNI_PLATFORM,
        define,
        context,
    };
}
exports.parseScripts = parseScripts;
