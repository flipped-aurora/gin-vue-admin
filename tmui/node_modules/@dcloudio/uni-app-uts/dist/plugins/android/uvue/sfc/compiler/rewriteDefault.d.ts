import MagicString from 'magic-string';
import type { ParserPlugin } from '@babel/parser';
import type { Statement } from '@babel/types';
export declare function rewriteDefault(input: string, as: string, define: string, parserPlugins?: ParserPlugin[]): string;
/**
 * Utility for rewriting `export default` in a script block into a variable
 * declaration so that we can inject things into it
 */
export declare function rewriteDefaultAST(ast: Statement[], s: MagicString, as: string, define: string): void;
export declare function hasDefaultExport(ast: Statement[]): boolean;
