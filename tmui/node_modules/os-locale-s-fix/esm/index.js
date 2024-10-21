/*!
 // Copyright (c) Sindre Sorhus <sindresorhus@gmail.com> (https://sindresorhus.com)
 // Released under the MIT license
 // https://opensource.org/licenses/mit-license.php
 */
/*!
 - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
  Copyright (C) 2020 jeffy-g <hirotom1107@gmail.com>
  Released under the MIT license
  https://opensource.org/licenses/mit-license.php
 - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
*/
import * as lcid from "lcid";
import { execFile, execFileSync, } from "child_process";
const defaultLocale = "en_US";
const getStdOut = async (command, args) => await new Promise((resolve) => {
    execFile(command, args, (err, stdout) => {
        resolve(err || stdout);
    });
});
const getStdOutSync = (command, args) => {
    try {
        return execFileSync(command, args);
    }
    catch (e) {
        return e;
    }
};
const normalise = (input) => input.replace(/_/, "-");
function validate(result, processor) {
    if (typeof result === "string" && result.length) {
        return processor ? processor(result) : result.trim();
    }
    return defaultLocale;
}
const getEnvLocale = (env = process.env) => env.LC_ALL || env.LC_MESSAGES || env.LANG || env.LANGUAGE || "";
function parseLocale(str) {
    const env = str.split("\n").reduce((env, definition) => {
        const [key, value] = definition.split("=");
        if (key && value) {
            env[key] = value.replace(/^"|"$/g, "");
        }
        return env;
    }, {});
    return getEnvLocale(env);
}
const purgeExtraToken = (str) => (str && str.replace(/[.:].*/, "")) || defaultLocale;
const getSupportedLocale = (locale, locales) => locales.includes(locale) ? locale : defaultLocale;
const [getAppleLocale, getAppleLocaleSync] = ((cmd0, args0, cmd1, args1) => {
    return [
        async () => {
            const results = await Promise.all([
                getStdOut(cmd0, args0).then(ret => validate(ret)),
                getStdOut(cmd1, args1).then(ret => validate(ret))
            ]);
            return getSupportedLocale(results[0], results[1]);
        },
        () => getSupportedLocale(validate(getStdOutSync(cmd0, args0)), validate(getStdOutSync(cmd1, args1)))
    ];
})("defaults", ["read", "-globalDomain", "AppleLocale"], "locale", ["-a"]);
const getUnixLocale = async () => purgeExtraToken(parseLocale(await getStdOut("locale").then(ret => validate(ret))));
const getUnixLocaleSync = () => purgeExtraToken(parseLocale(validate(getStdOutSync("locale"))));
const parseLCID = (result) => {
    const lcidCode = parseInt(result.replace("Locale", ""), 16);
    return lcid.from(lcidCode);
};
const [getWinLocale, getWinLocaleSync] = ((cmd0, args0) => {
    return [
        async () => validate(await getStdOut(cmd0, args0), parseLCID),
        () => validate(getStdOutSync(cmd0, args0), parseLCID)
    ];
})("wmic", ["os", "get", "locale"]);
let detector;
{
    const gettersSlot = [
        {
            win32: getWinLocaleSync,
            darwin: getAppleLocaleSync,
            linux: getUnixLocaleSync,
        }, {
            win32: getWinLocale,
            darwin: getAppleLocale,
            linux: getUnixLocale,
        }
    ];
    const typeString = {}.toString;
    const isPromise = (o) => typeString.call(o) === "[object Promise]";
    let cacheLocal;
    const base = (async) => (options = {}) => {
        options = { spawn: true, cache: true, ...options };
        const cache = options.cache;
        if (cache && (cacheLocal === null || cacheLocal === void 0 ? void 0 : cacheLocal.length)) {
            return (async ? Promise.resolve(cacheLocal) : cacheLocal);
        }
        const functions = gettersSlot[+(!!async)];
        let locale;
        const withCache = (l, mustPromise) => {
            l = normalise(l);
            cacheLocal = cache ? l : void 0;
            return (mustPromise ? Promise.resolve(l) : l);
        };
        const envLocale = getEnvLocale();
        if (envLocale || !options.spawn) {
            locale = purgeExtraToken(envLocale);
        }
        else {
            let { platform } = process;
            if (platform !== "win32" && platform !== "darwin") {
                platform = "linux";
            }
            locale = functions[platform]();
        }
        if (isPromise(locale)) {
            return locale.then(result => withCache(result));
        }
        else {
            return withCache(locale, async === true || void 0);
        }
    };
    detector = (base(true));
    detector.sync = base();
    Object.defineProperties(detector, {
        purge: {
            value: () => cacheLocal = void 0,
            enumerable: false,
        },
        version: {
            value: "v1.0.8",
            enumerable: true,
        },
    });
}
export const osLocale = Object.freeze(detector);