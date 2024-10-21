import { type ElementNode, ElementTypes, type ExpressionNode, type SimpleExpressionNode } from '@vue/compiler-core';
import { type NodeTransform, type TransformContext } from '../transform';
import { type Expression, type Identifier, type Pattern, type RestElement } from '@babel/types';
import type { CodegenScope, CodegenVForScope } from '../options';
export type VForOptions = Omit<ForParseResult, 'tagType'> & {
    sourceExpr?: Expression;
    sourceAlias: string;
    sourceCode: string;
    valueCode: string;
    valueExpr: Identifier | Pattern | RestElement;
    valueAlias: string;
    keyCode: string;
    keyExpr: Identifier | Pattern | RestElement;
    keyAlias: string;
    indexCode: string;
    indexExpr: Identifier | Pattern | RestElement;
    indexAlias: string;
    node: ElementNode;
};
export type ForElementNode = ElementNode & {
    vFor: VForOptions & {
        source: ExpressionNode;
    };
};
export declare function isForElementNode(node: unknown): node is ForElementNode;
export declare const transformFor: NodeTransform;
export declare function parseVForScope(currentScope: CodegenScope): CodegenVForScope | undefined;
export interface ForParseResult {
    source: ExpressionNode;
    value: ExpressionNode;
    key: ExpressionNode;
    index: ExpressionNode;
    tagType: ElementTypes;
}
export declare function parseForExpression(input: SimpleExpressionNode, context: TransformContext): ForParseResult | undefined;
export declare function createForLoopParams({ value, key, index }: ForParseResult, memoArgs?: ExpressionNode[]): ExpressionNode[];
export declare function createVForArrowFunctionExpression({ valueExpr, keyExpr, indexExpr, properties, }: CodegenVForScope): import("@babel/types").ArrowFunctionExpression;
