"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.addMiniProgramTemplateFilter = exports.clearMiniProgramTemplateFilter = exports.addMiniProgramTemplateFile = exports.clearMiniProgramTemplateFiles = exports.findMiniProgramTemplateFiles = void 0;
const path_1 = __importDefault(require("path"));
const uni_shared_1 = require("@dcloudio/uni-shared");
const utils_1 = require("../utils");
const templateFilesCache = new Map();
const templateFiltersCache = new Map();
function relativeFilterFilename(filename, filter) {
    if (!filter.src) {
        return '';
    }
    return ('./' +
        (0, utils_1.normalizeMiniProgramFilename)(path_1.default.relative(path_1.default.dirname(filename), filter.src)));
}
function findMiniProgramTemplateFiles(genFilter) {
    const files = Object.create(null);
    templateFilesCache.forEach((code, filename) => {
        if (!genFilter) {
            files[filename] = code;
        }
        else {
            const filters = getMiniProgramTemplateFilters(filename);
            if (filters && filters.length) {
                files[filename] =
                    filters
                        .map((filter) => genFilter(filter, relativeFilterFilename(filename, filter)))
                        .join(uni_shared_1.LINEFEED) +
                        uni_shared_1.LINEFEED +
                        code;
            }
            else {
                files[filename] = code;
            }
        }
    });
    return files;
}
exports.findMiniProgramTemplateFiles = findMiniProgramTemplateFiles;
function clearMiniProgramTemplateFiles() {
    templateFilesCache.clear();
}
exports.clearMiniProgramTemplateFiles = clearMiniProgramTemplateFiles;
function addMiniProgramTemplateFile(filename, code) {
    templateFilesCache.set(filename, code);
}
exports.addMiniProgramTemplateFile = addMiniProgramTemplateFile;
function getMiniProgramTemplateFilters(filename) {
    return templateFiltersCache.get(filename);
}
function clearMiniProgramTemplateFilter(filename) {
    templateFiltersCache.delete(filename);
}
exports.clearMiniProgramTemplateFilter = clearMiniProgramTemplateFilter;
function addMiniProgramTemplateFilter(filename, filter) {
    const filters = templateFiltersCache.get(filename);
    if (filters) {
        const filterIndex = filters.findIndex((f) => f.id === filter.id);
        if (filterIndex > -1) {
            filters.splice(filterIndex, 1, filter);
        }
        else {
            filters.push(filter);
        }
    }
    else {
        templateFiltersCache.set(filename, [filter]);
    }
}
exports.addMiniProgramTemplateFilter = addMiniProgramTemplateFilter;
