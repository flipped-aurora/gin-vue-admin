import type { TextRange } from '../types';
import type * as ts from 'typescript/lib/tsserverlibrary';
export interface ScriptRanges extends ReturnType<typeof parseScriptRanges> {
}
export declare function parseScriptRanges(ts: typeof import('typescript/lib/tsserverlibrary'), ast: ts.SourceFile, hasScriptSetup: boolean, withNode: boolean): {
    exportDefault: (TextRange & {
        expression: TextRange;
        args: TextRange;
        argsNode: ts.ObjectLiteralExpression | undefined;
        componentsOption: TextRange | undefined;
        componentsOptionNode: ts.ObjectLiteralExpression | undefined;
        nameOption: TextRange | undefined;
    }) | undefined;
    bindings: TextRange[];
};
//# sourceMappingURL=scriptRanges.d.ts.map