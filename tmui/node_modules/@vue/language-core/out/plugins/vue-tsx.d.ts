import { Sfc, VueLanguagePlugin } from '../types';
import * as muggle from 'muggle-string';
export declare const tsCodegen: WeakMap<Sfc, {
    scriptRanges: () => {
        exportDefault: (import("../types").TextRange & {
            expression: import("../types").TextRange;
            args: import("../types").TextRange;
            argsNode: import("typescript").ObjectLiteralExpression | undefined;
            componentsOption: import("../types").TextRange | undefined;
            componentsOptionNode: import("typescript").ObjectLiteralExpression | undefined;
            nameOption: import("../types").TextRange | undefined;
        }) | undefined;
        bindings: import("../types").TextRange[];
    } | undefined;
    scriptSetupRanges: () => {
        leadingCommentEndOffset: number;
        importSectionEndOffset: number;
        bindings: import("../types").TextRange[];
        props: {
            name?: string | undefined;
            define?: (import("../types").TextRange & {
                arg?: import("../types").TextRange | undefined;
                typeArg?: import("../types").TextRange | undefined;
            } & {
                statement: import("../types").TextRange;
            }) | undefined;
            withDefaults?: (import("../types").TextRange & {
                arg?: import("../types").TextRange | undefined;
            }) | undefined;
        };
        slots: {
            name?: string | undefined;
            define?: (import("../types").TextRange & {
                arg?: import("../types").TextRange | undefined;
                typeArg?: import("../types").TextRange | undefined;
            }) | undefined;
        };
        emits: {
            name?: string | undefined;
            define?: (import("../types").TextRange & {
                arg?: import("../types").TextRange | undefined;
                typeArg?: import("../types").TextRange | undefined;
            }) | undefined;
        };
        expose: {
            name?: string | undefined;
            define?: (import("../types").TextRange & {
                arg?: import("../types").TextRange | undefined;
                typeArg?: import("../types").TextRange | undefined;
            }) | undefined;
        };
        defineProp: {
            name: import("../types").TextRange | undefined;
            nameIsString: boolean;
            type: import("../types").TextRange | undefined;
            defaultValue: import("../types").TextRange | undefined;
            required: boolean;
        }[];
    } | undefined;
    lang: () => string;
    generatedScript: () => {
        codes: muggle.Segment<import("@volar/language-core").FileRangeCapabilities>[];
        codeStacks: muggle.StackNode[];
        mirrorBehaviorMappings: import("@volar/source-map").Mapping<[import("@volar/language-core").MirrorBehaviorCapabilities, import("@volar/language-core").MirrorBehaviorCapabilities]>[];
    };
    generatedTemplate: () => {
        codes: (string | [string, string | undefined, number | [number, number], import("@volar/language-core").FileRangeCapabilities])[];
        codeStacks: muggle.StackNode[];
        formatCodes: (string | [string, string | undefined, number | [number, number], import("@volar/language-core").FileRangeCapabilities])[];
        formatCodeStacks: muggle.StackNode[];
        cssCodes: (string | [string, string | undefined, number | [number, number], import("@volar/language-core").FileRangeCapabilities])[];
        cssCodeStacks: muggle.StackNode[];
        tagNames: Record<string, number[]>;
        accessedGlobalVariables: Set<string>;
        hasSlot: boolean;
    } | undefined;
}>;
declare const plugin: VueLanguagePlugin;
export default plugin;
//# sourceMappingURL=vue-tsx.d.ts.map