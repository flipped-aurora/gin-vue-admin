import type * as ts from 'typescript/lib/tsserverlibrary';
import type { VueCompilerOptions, TextRange } from '../types';
export interface ScriptSetupRanges extends ReturnType<typeof parseScriptSetupRanges> {
}
export declare function parseScriptSetupRanges(ts: typeof import('typescript/lib/tsserverlibrary'), ast: ts.SourceFile, vueCompilerOptions: VueCompilerOptions): {
    leadingCommentEndOffset: number;
    importSectionEndOffset: number;
    bindings: TextRange[];
    props: {
        name?: string | undefined;
        define?: (TextRange & {
            arg?: TextRange | undefined;
            typeArg?: TextRange | undefined;
        } & {
            statement: TextRange;
        }) | undefined;
        withDefaults?: (TextRange & {
            arg?: TextRange | undefined;
        }) | undefined;
    };
    slots: {
        name?: string | undefined;
        define?: (TextRange & {
            arg?: TextRange | undefined;
            typeArg?: TextRange | undefined;
        }) | undefined;
    };
    emits: {
        name?: string | undefined;
        define?: (TextRange & {
            arg?: TextRange | undefined;
            typeArg?: TextRange | undefined;
        }) | undefined;
    };
    expose: {
        name?: string | undefined;
        define?: (TextRange & {
            arg?: TextRange | undefined;
            typeArg?: TextRange | undefined;
        }) | undefined;
    };
    defineProp: {
        name: TextRange | undefined;
        nameIsString: boolean;
        type: TextRange | undefined;
        defaultValue: TextRange | undefined;
        required: boolean;
    }[];
};
export declare function parseBindingRanges(ts: typeof import('typescript/lib/tsserverlibrary'), sourceFile: ts.SourceFile): TextRange[];
export declare function findBindingVars(ts: typeof import('typescript/lib/tsserverlibrary'), left: ts.BindingName, sourceFile: ts.SourceFile): TextRange[];
export declare function getStartEnd(node: ts.Node, sourceFile: ts.SourceFile): {
    start: number;
    end: number;
};
//# sourceMappingURL=scriptSetupRanges.d.ts.map