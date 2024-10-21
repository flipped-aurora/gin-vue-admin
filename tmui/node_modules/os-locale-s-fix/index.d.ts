//
// Copyright (c) Sindre Sorhus <sindresorhus@gmail.com> (https://sindresorhus.com)
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
//
/*!
 - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
  Copyright (C) 2020 jeffy-g <hirotom1107@gmail.com>
  Released under the MIT license
  https://opensource.org/licenses/mit-license.php
 - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
*/
export declare interface LocaleDetectorOptions {
    /**
     * Set to `false` to avoid spawning subprocesses and instead only resolve the locale from environment variables.
     * 
     * @default true
     */
    readonly spawn?: boolean;
    /**
     * @default true
     */
    readonly cache?: boolean;
}
export declare interface LocaleDetectorBase {
    /**
     * Get the system [locale](https://en.wikipedia.org/wiki/Locale_(computer_software)).
     * 
     * @returns The locale.
     * 
     * @example
     * ```
     * import { osLocale } = require("os-locale-s");
         * 
     * (async () => {
     *     console.log(await osLocale());
     *     //=> 'en-US'
     * })();
     * ```
     */
    (options?: LocaleDetectorOptions): Promise<string>;
    /**
     *
     */
    readonly version: string;
}

export declare interface LocaleDetector extends LocaleDetectorBase {
    /**
     * Synchronously get the system [locale](https://en.wikipedia.org/wiki/Locale_(computer_software)).
     * 
     * @returns The locale.
     */
    sync(options?: LocaleDetectorOptions): string;
}

export declare const osLocale: LocaleDetector;

export as namespace NsOsLocale;
