import type { Program } from '@babel/types';
import { type ParserPlugin } from '@babel/parser';
export declare function parseProgram(code: string, importer: string, { babelParserPlugins }: {
    babelParserPlugins?: ParserPlugin[];
}): Program;
