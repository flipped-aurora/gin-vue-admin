"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.emitFile = exports.getTemplateFiles = exports.getFilterFiles = void 0;
const path_1 = __importDefault(require("path"));
const debug_1 = __importDefault(require("debug"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const renderjs_1 = require("../plugins/renderjs");
const debugTemplate = (0, debug_1.default)('uni:mp-template');
function getFilterFiles(resolvedConfig, getModuleInfo) {
    const filters = Object.create(null);
    const filtersCache = (0, renderjs_1.getFiltersCache)(resolvedConfig);
    if (!filtersCache.length) {
        return filters;
    }
    const inputDir = process.env.UNI_INPUT_DIR;
    function addFilter(id, filter) {
        const templateFilename = (0, uni_cli_shared_1.removeExt)((0, uni_cli_shared_1.normalizeMiniProgramFilename)(id, inputDir));
        (0, uni_cli_shared_1.addMiniProgramTemplateFilter)(templateFilename, filter);
        const filterFilename = (0, uni_cli_shared_1.removeExt)((0, uni_cli_shared_1.normalizeMiniProgramFilename)(filter.id, inputDir));
        if (templateFilename !== filterFilename) {
            // 外链
            filter.src = filterFilename;
            filters[filterFilename] = filter;
        }
    }
    filtersCache.forEach((filter) => {
        const moduleInfo = getModuleInfo(filter.id);
        if (!moduleInfo) {
            return;
        }
        const { importers } = moduleInfo;
        if (!importers.length) {
            return;
        }
        importers.forEach((importer) => addFilter(importer, filter));
    });
    return filters;
}
exports.getFilterFiles = getFilterFiles;
function getTemplateFiles(template) {
    const files = (0, uni_cli_shared_1.findMiniProgramTemplateFiles)(template.filter?.generate);
    (0, uni_cli_shared_1.clearMiniProgramTemplateFiles)();
    return files;
}
exports.getTemplateFiles = getTemplateFiles;
const emitFile = (emittedFile) => {
    if (emittedFile.type === 'asset') {
        const filename = emittedFile.fileName;
        (0, uni_cli_shared_1.addMiniProgramTemplateFile)((0, uni_cli_shared_1.removeExt)((0, uni_cli_shared_1.normalizeMiniProgramFilename)(path_1.default.relative(process.env.UNI_INPUT_DIR, filename))), emittedFile.source.toString());
        debugTemplate(filename);
        return filename;
    }
    return '';
};
exports.emitFile = emitFile;
