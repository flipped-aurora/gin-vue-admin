"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.nvueOutDir = exports.templateDir = void 0;
const path_1 = __importDefault(require("path"));
exports.templateDir = path_1.default.resolve(__dirname, '../lib/template/');
function nvueOutDir() {
    return path_1.default.join(process.env.UNI_OUTPUT_DIR, '../.nvue');
}
exports.nvueOutDir = nvueOutDir;
