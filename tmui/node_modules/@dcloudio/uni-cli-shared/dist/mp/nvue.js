"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.genNVueCssCode = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const nvue_1 = require("../json/app/manifest/nvue");
function genNVueCssCode(manifestJson) {
    let nvueCssCode = fs_1.default.readFileSync(path_1.default.resolve(__dirname, '../../lib/nvue.css'), 'utf8');
    const flexDirection = (0, nvue_1.getNVueFlexDirection)(manifestJson);
    if (flexDirection !== 'column') {
        nvueCssCode = nvueCssCode.replace('column', flexDirection);
    }
    return nvueCssCode;
}
exports.genNVueCssCode = genNVueCssCode;
