import { type AttributeNode, type ComponentNode, type DirectiveNode, type ElementNode, type ExpressionNode, type Position, type RootNode, type SourceLocation, type TemplateChildNode, type TransformContext } from '@vue/compiler-core';
export declare const VUE_REF = "r";
export declare const VUE_REF_IN_FOR = "r-i-f";
export declare function isVueSfcFile(id: string): boolean;
export declare function isUserComponent(node: RootNode | TemplateChildNode, context: {
    isBuiltInComponent: TransformContext['isBuiltInComponent'];
}): node is ComponentNode;
export declare function createAttributeNode(name: string, content: string): AttributeNode;
export declare function addStaticClass(node: ElementNode, clazz: string): string | number | undefined;
export declare function createDirectiveNode(name: string, arg: string, exp?: string | ExpressionNode): DirectiveNode;
export declare function createOnDirectiveNode(name: string, value: string): DirectiveNode;
export declare function createBindDirectiveNode(name: string, value: string | ExpressionNode): DirectiveNode;
export declare function createUniVueTransformAssetUrls(base: string): {
    base: string;
    includeAbsolute: boolean;
    tags: {
        audio: string[];
        video: string[];
        img: string[];
        image: string[];
        'cover-image': string[];
        'v-uni-audio': string[];
        'v-uni-video': string[];
        'v-uni-image': string[];
        'v-uni-cover-image': string[];
        'u-image': string[];
        'u-video': string[];
    };
};
export declare function getBaseNodeTransforms(base: string): import("@vue/compiler-core").NodeTransform[];
export declare function renameProp(name: string, prop?: DirectiveNode | AttributeNode): void;
export declare function getInnerRange(loc: SourceLocation, offset: number, length: number): SourceLocation;
export declare function advancePositionWithClone(pos: Position, source: string, numberOfCharacters?: number): Position;
export declare function advancePositionWithMutation(pos: Position, source: string, numberOfCharacters?: number): Position;
