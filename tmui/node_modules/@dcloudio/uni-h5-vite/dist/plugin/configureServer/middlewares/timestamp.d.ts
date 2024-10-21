/// <reference types="node" />
import type { IncomingMessage, ServerResponse } from 'http';
import type { ViteDevServer } from 'vite';
import type { NextHandler } from './static';
export declare function uniTimestampMiddleware(server: ViteDevServer): (req: IncomingMessage, _: ServerResponse, next: NextHandler) => Promise<void>;
