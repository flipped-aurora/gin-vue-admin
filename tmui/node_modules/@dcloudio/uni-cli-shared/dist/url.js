"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.decodeBase64Url = exports.encodeBase64Url = void 0;
const base64url_1 = __importDefault(require("base64url"));
function encodeBase64Url(str) {
    return base64url_1.default.encode(str);
}
exports.encodeBase64Url = encodeBase64Url;
function decodeBase64Url(str) {
    return base64url_1.default.decode(str);
}
exports.decodeBase64Url = decodeBase64Url;
