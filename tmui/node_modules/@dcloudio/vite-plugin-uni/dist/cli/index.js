"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const cac_1 = require("cac");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const utils_1 = require("./utils");
const action_1 = require("./action");
(0, uni_cli_shared_1.fixBinaryPath)();
const cli = (0, cac_1.cac)('uni');
cli
    .option('-c, --config <file>', `[string] use specified config file`)
    .option('-p, --platform [platform]', '[string] ' + utils_1.PLATFORMS.join(' | '), {
    default: 'h5',
})
    .option('--base <path>', `[string] public base path (default: /)`)
    .option('-ssr', '[boolean] server-side rendering', {
    default: false,
})
    .option('-l, --logLevel <level>', `[string] silent | error | warn | all`)
    .option('--clearScreen', `[boolean] allow/disable clear screen when logging`)
    .option('-d, --debug [feat]', `[string | boolean] show debug logs`)
    .option('-f, --filter <filter>', `[string] filter debug logs`)
    .option('-m, --mode <mode>', `[string] set env mode`)
    .option('--minify [minifier]', `[boolean | "terser" | "esbuild"] enable/disable minification, ` +
    `or specify minifier to use (default: terser)`)
    .option('--autoHost [autoHost]', `[string] specify automator hostname`)
    .option('--autoPort [autoPort]', `[number] specify automator port`)
    .option('--devtools', `[boolean] enable devtools`)
    .option('--devtoolsHost [devtoolsHost]', `[string] specify devtools hostname`)
    .option('--devtoolsPort [devtoolsPort]', `[number] specify devtools port`)
    .option('--subpackage [subpackage]', `[string] specify subpackage to build`)
    .option('--plugin', `[boolean] build plugin`);
cli
    .command('')
    .alias('dev')
    .option('--host [host]', `[string] specify hostname`)
    .option('--port <port>', `[number] specify port`)
    .option('--https', `[boolean] use TLS + HTTP/2`)
    .option('--open [path]', `[boolean | string] open browser on startup`)
    .option('--cors', `[boolean] enable CORS`)
    .option('--strictPort', `[boolean] exit if specified port is already in use`)
    .option('--force', `[boolean] force the optimizer to ignore the cache and re-bundle`)
    .action(action_1.runDev);
cli
    .command('build')
    .option('--outDir <dir>', `[string] output directory (default: dist)`)
    .option('--assetsInlineLimit <number>', `[number] static asset base64 inline threshold in bytes (default: 4096)`)
    .option('--sourcemap', `[boolean] output source maps for build (default: false)`)
    .option('--manifest', `[boolean] emit build manifest json`)
    .option('--ssrManifest', `[boolean] emit ssr manifest json`)
    .option('--emptyOutDir', `[boolean] force empty outDir when it's outside of root`, {
    default: true,
})
    .option('-w, --watch', `[boolean] rebuilds when modules have changed on disk`)
    .action(action_1.runBuild);
cli.help();
cli.version(require('../../package.json').version);
cli.parse();
