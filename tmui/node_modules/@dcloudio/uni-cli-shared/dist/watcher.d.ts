/// <reference types="node" />
import { type FSWatcher, type WatchOptions } from 'chokidar';
type FileTransform = (source: Buffer, filename: string) => void | string;
export interface FileWatcherOptions {
    src: string | string[];
    dest: string;
    transform?: FileTransform;
}
export declare class FileWatcher {
    private src;
    private dest;
    private transform?;
    private watcher;
    private onChange?;
    constructor({ src, dest, transform }: FileWatcherOptions);
    watch(watchOptions: WatchOptions & {
        cwd: string;
        readyTimeout?: number;
    }, onReady?: (watcher: FSWatcher) => void, onChange?: () => void): FSWatcher;
    add(paths: string | ReadonlyArray<string>): FSWatcher;
    unwatch(paths: string | ReadonlyArray<string>): FSWatcher;
    close(): Promise<void>;
    copy(from: string): void;
    remove(from: string): void;
    info(type: 'close' | 'copy' | 'remove' | 'add' | 'unwatch', msg?: string | unknown): void;
    from(from: string): string;
    to(from: string): string;
}
export {};
