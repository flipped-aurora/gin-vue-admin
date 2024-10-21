"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getIsStaticFile = exports.createIsStaticFile = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const utils_1 = require("../utils");
const json_1 = require("../../../../json/json");
const uniModulesStaticRe = /^uni_modules\/[^/]+\/static\//;
function createIsStaticFile() {
    let subPackageStatics = [];
    const pagesFilename = path_1.default.join(process.env.UNI_INPUT_DIR, 'pages.json');
    if (fs_1.default.existsSync(pagesFilename)) {
        const pagesJson = (0, json_1.parseJson)(fs_1.default.readFileSync(pagesFilename, 'utf8'), true);
        subPackageStatics = (pagesJson.subPackages || pagesJson.subpackages || [])
            .filter((subPackage) => subPackage.root)
            .map((subPackage) => {
            return (0, utils_1.normalizePath)(path_1.default.join(subPackage.root, 'static')) + '/';
        });
    }
    return function isStaticFile(relativeFile) {
        if (path_1.default.isAbsolute(relativeFile)) {
            relativeFile = (0, utils_1.normalizePath)(path_1.default.relative(process.env.UNI_INPUT_DIR, relativeFile));
        }
        return (relativeFile.startsWith('static/') ||
            uniModulesStaticRe.test(relativeFile) ||
            subPackageStatics.some((s) => relativeFile.startsWith(s)));
    };
}
exports.createIsStaticFile = createIsStaticFile;
let isStaticFile;
function getIsStaticFile() {
    if (!isStaticFile) {
        isStaticFile = createIsStaticFile();
    }
    return isStaticFile;
}
exports.getIsStaticFile = getIsStaticFile;
