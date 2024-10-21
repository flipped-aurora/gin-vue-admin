export type Options = {
	browser?: boolean;
	conditions?: readonly string[];
	require?: boolean;
	unsafe?: false;
} | {
	conditions?: readonly string[];
	unsafe?: true;
}

export function resolve<T=any>(pkg: T, entry: string, options?: Options): string | void;

export type BrowserFiles = Record<string, string | false>;

export function legacy<T=any>(pkg: T, options: { browser: true, fields?: readonly string[] }): BrowserFiles | string | void;
export function legacy<T=any>(pkg: T, options: { browser: string, fields?: readonly string[] }): string | false | void;
export function legacy<T=any>(pkg: T, options: { browser: false, fields?: readonly string[] }): string | void;
export function legacy<T=any>(pkg: T, options?: {
	browser?: boolean | string;
	fields?: readonly string[];
}): BrowserFiles | string | false | void;
