declare function getPort(
    port?: number | number[],
    host?: string
): Promise<number>;

export = getPort;
