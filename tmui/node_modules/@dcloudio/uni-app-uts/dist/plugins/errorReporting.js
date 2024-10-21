"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const os_1 = __importDefault(require("os"));
const https_1 = __importDefault(require("https"));
const crypto_1 = __importDefault(require("crypto"));
const SERVER_HOST = 'uacer.dcloud.net.cn';
const SERVER_PATH = '/http/error-report-x';
const EXCLUDE_ERROR_LIST = [
    'uni-app-compiler',
    'Error: ENOENT: no such file or directory',
];
function getMacHash() {
    let mac = '';
    const network = os_1.default.networkInterfaces();
    for (const key in network) {
        const array = network[key];
        for (let i = 0; i < array.length; i++) {
            const item = array[i];
            if (!item.family || (item.mac && item.mac === '00:00:00:00:00:00')) {
                continue;
            }
            if (
            // Node < v18
            typeof item.family === 'string' &&
                (item.family === 'IPv4' || item.family === 'IPv6')) {
                mac = item.mac;
                break;
            }
            else if (
            // Node >= v18
            typeof item.family === 'number' &&
                (item.family === 4 || item.family === 6)) {
                mac = item.mac;
                break;
            }
        }
    }
    return crypto_1.default.createHash('md5').update(mac).digest('hex');
}
const CacheList = [];
function shouldReport(err = '') {
    try {
        const errMsg = err.toString();
        const errorIndex = EXCLUDE_ERROR_LIST.findIndex((item) => errMsg.includes(item));
        if (errorIndex >= 0) {
            return false;
        }
        // 目前简单的上报逻辑为：错误信息中包含@dcloudio包名
        if (errMsg.includes('@dcloudio') &&
            !errMsg.includes('Errors compiling template')) {
            return true;
        }
    }
    catch (e) { }
    return false;
}
function report(type, err) {
    if (!shouldReport(err)) {
        return;
    }
    if (typeof err === 'object') {
        try {
            err = err.toString();
        }
        catch (e) { }
    }
    const UNI_INPUT_DIR_REG = new RegExp(process.env.UNI_INPUT_DIR, 'ig');
    const UNI_CLI_CONTEXT_REG = new RegExp(process.env.UNI_CLI_CONTEXT, 'ig');
    err = err.replace(UNI_INPUT_DIR_REG, 'UNI_INPUT_DIR');
    err = err.replace(UNI_CLI_CONTEXT_REG, 'UNI_CLI_CONTEXT');
    const data = JSON.stringify({
        di: getMacHash(),
        np: process.platform,
        nv: process.version,
        cp: process.env.UNI_PLATFORM,
        cv: process.env.UNI_COMPILER_VERSION,
        hx: process.env.UNI_COMPILER_VERSION_TYPE,
        et: type,
        em: err,
    });
    const dataHash = crypto_1.default.createHash('md5').update(data).digest('hex');
    if (CacheList.includes(dataHash)) {
        return;
    }
    CacheList.push(dataHash);
    setTimeout(() => {
        const req = https_1.default.request({
            hostname: SERVER_HOST,
            port: 443,
            path: SERVER_PATH,
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Content-Length': data.length,
            },
        });
        req.write(data);
        req.end();
    }, 10);
}
global.__error_reporting__ = report;
process
    .on('unhandledRejection', (reason, p) => {
    global.__error_reporting__('unhandledRejection', reason);
    console.log(reason);
})
    .on('uncaughtException', (err) => {
    global.__error_reporting__('uncaughtException', err.stack);
    console.log(err);
});
