"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniManifestJsonPlugin = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const shared_1 = require("@vue/shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function findUserProjectConfigFile(inputDir, config) {
    for (let i = 0; i < config.length; i++) {
        const projectFilename = path_1.default.resolve(inputDir, config[i]);
        // 根目录包含指定文件，则直接拷贝
        if (fs_1.default.existsSync(projectFilename)) {
            return projectFilename;
        }
    }
}
function uniManifestJsonPlugin(options) {
    return (0, uni_cli_shared_1.defineUniManifestJsonPlugin)((opts) => {
        const inputDir = process.env.UNI_INPUT_DIR;
        const platform = process.env.UNI_PLATFORM;
        let projectJson;
        let userProjectFilename;
        let projectSource;
        if (options.project) {
            userProjectFilename = findUserProjectConfigFile(inputDir, options.project.config);
        }
        return {
            name: 'uni:mp-manifest-json',
            enforce: 'pre',
            transform(code, id) {
                if (!opts.filter(id)) {
                    return;
                }
                this.addWatchFile(path_1.default.resolve(inputDir, 'manifest.json'));
                (0, uni_cli_shared_1.getLocaleFiles)(path_1.default.resolve(inputDir, 'locale')).forEach((filepath) => {
                    this.addWatchFile(filepath);
                });
                if (options.project) {
                    // 根目录包含指定文件，则直接拷贝
                    if (userProjectFilename) {
                        this.addWatchFile(userProjectFilename);
                        projectJson = (0, uni_cli_shared_1.parseJson)(fs_1.default.readFileSync(userProjectFilename, 'utf8'));
                    }
                    else {
                        const template = options.project.source;
                        if ((0, shared_1.hasOwn)(template, 'appid')) {
                            let projectName = path_1.default.basename(inputDir);
                            if (projectName === 'src') {
                                projectName = path_1.default.basename(path_1.default.dirname(inputDir));
                            }
                            template.projectname = projectName;
                            // TODO condition
                            if (process.env.UNI_AUTOMATOR_WS_ENDPOINT) {
                                if (!template.setting) {
                                    template.setting = {};
                                }
                                template.setting.urlCheck = false;
                            }
                            projectJson = (0, uni_cli_shared_1.parseMiniProgramProjectJson)(code, platform, {
                                template,
                                pagesJson: (0, uni_cli_shared_1.parsePagesJsonOnce)(inputDir, platform),
                            });
                        }
                        else {
                            // 无需解析，直接拷贝，如 quickapp-webview
                            projectJson = template;
                        }
                    }
                }
                return {
                    code: '',
                    map: { mappings: '' },
                };
            },
            generateBundle() {
                if (projectJson && options.project) {
                    const { filename, normalize } = options.project;
                    const source = JSON.stringify(normalize ? normalize(projectJson) : projectJson, null, 2);
                    if (projectSource !== source) {
                        projectSource = source;
                        this.emitFile({
                            fileName: filename,
                            type: 'asset',
                            source,
                        });
                    }
                }
            },
        };
    });
}
exports.uniManifestJsonPlugin = uniManifestJsonPlugin;
