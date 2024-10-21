"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.createSSRServer = exports.createServer = void 0;
const fs_1 = __importDefault(require("fs"));
const os_1 = __importDefault(require("os"));
const path_1 = __importDefault(require("path"));
const picocolors_1 = __importDefault(require("picocolors"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const utils_1 = require("./utils");
function createLogger(level, options) {
    return Promise.resolve().then(() => __importStar(require('vite'))).then(({ createLogger }) => createLogger(level, options));
}
function createViteServer(inlineConfig) {
    return Promise.resolve().then(() => __importStar(require('vite'))).then(({ createServer }) => createServer(inlineConfig));
}
async function createServer(options) {
    const server = await createViteServer((0, utils_1.addConfigFile)({
        root: process.env.VITE_ROOT_DIR,
        configFile: options.config,
        base: options.base,
        mode: options.mode,
        logLevel: options.logLevel || 'info',
        clearScreen: options.clearScreen,
        server: (0, utils_1.cleanOptions)(options),
    }));
    await server.listen();
    const logger = server.config.logger;
    logger.info(picocolors_1.default.cyan(`\n  vite v${require('vite/package.json').version}`) +
        picocolors_1.default.green(` dev server running at:\n`), {
        clear: !server.config.logger.hasWarned,
    });
    server.printUrls();
    // printUrls 会在 nextTick 中输出
    process.nextTick(() => (0, utils_1.printStartupDuration)(logger));
    return server;
}
exports.createServer = createServer;
async function createSSRServer(options) {
    // 延迟加载
    const app = (await Promise.resolve().then(() => __importStar(require('express')))).default();
    /**
     * @type {import('vite').ViteDevServer}
     */
    const vite = await createViteServer((0, utils_1.addConfigFile)({
        // custom: don't include HTML middlewares
        appType: 'custom',
        root: process.env.VITE_ROOT_DIR,
        configFile: options.config,
        base: options.base,
        mode: options.mode,
        logLevel: options.logLevel || 'info',
        clearScreen: options.clearScreen,
        server: {
            middlewareMode: true,
            watch: {
                // During tests we edit the files too fast and sometimes chokidar
                // misses change events, so enforce polling for consistency
                usePolling: true,
                interval: 100,
            },
        },
    }));
    // use vite's connect instance as middleware
    app.use(vite.middlewares);
    app.use('*', async (req, res) => {
        try {
            const manifestJson = (0, uni_cli_shared_1.parseManifestJson)(process.env.UNI_INPUT_DIR);
            const h5 = (0, uni_cli_shared_1.getPlatformManifestJson)(manifestJson, 'h5');
            const base = (h5 && h5.router && h5.router.base) || '';
            const url = req.originalUrl.replace(base, '');
            const template = await vite.transformIndexHtml(url, fs_1.default.readFileSync(path_1.default.resolve(process.env.VITE_ROOT_DIR, 'index.html'), 'utf-8'));
            const render = (await vite.ssrLoadModule(path_1.default.resolve(process.env.UNI_INPUT_DIR, 'entry-server.js'))).render;
            const { title, headMeta, preloadLinks, appHtml, appContext } = await render(url);
            const icon = template.includes('rel="icon"')
                ? ''
                : '<link rel="icon" href="data:," />\n';
            const html = template
                .replace(/<title>(.*?)<\/title>/, `${icon}<title>${title}</title>`)
                .replace(`<!--head-meta-->`, headMeta)
                .replace(`<!--preload-links-->`, preloadLinks)
                .replace(`<!--app-html-->`, appHtml)
                .replace(`<!--app-context-->`, appContext);
            res.status(200).set({ 'Content-Type': 'text/html' }).end(html);
        }
        catch (e) {
            vite && vite.ssrFixStacktrace(e);
            res.status(500).end(e.stack);
        }
    });
    const logger = await createLogger(options.logLevel);
    const serverOptions = vite.config.server || {};
    let port = options.port || serverOptions.port || 5173;
    let hostname;
    if (options.host === 'localhost') {
        // Use a secure default
        hostname = '127.0.0.1';
    }
    else if (options.host === undefined || options.host === true) {
        // probably passed --host in the CLI, without arguments
        hostname = undefined; // undefined typically means 0.0.0.0 or :: (listen on all IPs)
    }
    else {
        hostname = options.host;
    }
    return new Promise((resolve, reject) => {
        const onSuccess = () => {
            printHttpServerUrls(server, vite.config);
            process.nextTick(() => (0, utils_1.printStartupDuration)(logger));
            resolve(vite);
        };
        const onError = (e) => {
            if (e.code === 'EADDRINUSE') {
                if (options.strictPort) {
                    server.off('error', onError);
                    reject(new Error(`Port ${port} is already in use`));
                }
                else {
                    logger.info(`Port ${port} is in use, trying another one...`);
                    server = app.listen(++port, hostname, onSuccess).on('error', onError);
                }
            }
            else {
                server.off('error', onError);
                reject(e);
            }
        };
        let server = app.listen(port, hostname, onSuccess).on('error', onError);
    });
}
exports.createSSRServer = createSSRServer;
function printHttpServerUrls(server, config) {
    printCommonServerUrls(server, config.server, config);
}
function printCommonServerUrls(server, options, config) {
    const address = server.address();
    const isAddressInfo = (x) => x?.address;
    if (isAddressInfo(address)) {
        const hostname = resolveHostname(options.host);
        const protocol = options.https ? 'https' : 'http';
        printServerUrls(hostname, protocol, address.port, config.base, config.logger.info);
    }
}
function printServerUrls(hostname, protocol, port, base, info) {
    if (hostname.host === '127.0.0.1') {
        const url = `${protocol}://${hostname.name}:${picocolors_1.default.bold(port)}${base}`;
        info(`  - Local: ${picocolors_1.default.cyan(url)}`);
        if (hostname.name !== '127.0.0.1') {
            info(`  - Network: ${picocolors_1.default.dim('use `--host` to expose')}`);
        }
    }
    else {
        Object.values(os_1.default.networkInterfaces())
            .flatMap((nInterface) => nInterface ?? [])
            .filter((detail) => detail &&
            detail.address &&
            // Node < v18
            ((typeof detail.family === 'string' && detail.family === 'IPv4') ||
                // Node >= v18
                // @ts-expect-error
                (typeof detail.family === 'number' && detail.family === 4)))
            .map((detail) => {
            const type = detail.address.includes('127.0.0.1')
                ? '  - Local:   '
                : '  * Network: ';
            const host = detail.address.replace('127.0.0.1', hostname.name);
            const url = `${protocol}://${host}:${picocolors_1.default.bold(port)}${base}`;
            return `${type} ${picocolors_1.default.cyan(url)}`;
        })
            .sort((msg1) => {
            return msg1.indexOf('- Local') > -1 ? -1 : 1;
        })
            .forEach((msg, index, arr) => {
            if (arr.length - 1 === index) {
                info(msg.replace('* Network', '- Network'));
            }
            else {
                info(msg);
            }
        });
    }
}
function resolveHostname(optionsHost) {
    let host;
    if (optionsHost === undefined || optionsHost === false) {
        // Use a secure default
        host = '127.0.0.1';
    }
    else if (optionsHost === true) {
        // If passed --host in the CLI without arguments
        host = undefined; // undefined typically means 0.0.0.0 or :: (listen on all IPs)
    }
    else {
        host = optionsHost;
    }
    // Set host name to localhost when possible, unless the user explicitly asked for '127.0.0.1'
    const name = (optionsHost !== '127.0.0.1' && host === '127.0.0.1') ||
        host === '0.0.0.0' ||
        host === '::' ||
        host === undefined
        ? 'localhost'
        : host;
    return { host, name };
}
