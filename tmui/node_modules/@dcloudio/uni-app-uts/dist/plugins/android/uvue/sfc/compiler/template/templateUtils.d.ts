/// <reference types="node" />
import { type UrlWithStringQuery } from 'url';
export declare function isRelativeUrl(url: string): boolean;
export declare function isExternalUrl(url: string): boolean;
export declare function isDataUrl(url: string): boolean;
/**
 * Parses string url into URL object.
 */
export declare function parseUrl(url: string): UrlWithStringQuery;
