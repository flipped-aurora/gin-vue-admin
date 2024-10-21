declare const container: {
    cpuNum(): number;
    cpuUsage(period?: number): Promise<number>;
    cpuLoad(period?: number): Promise<number>;
    memUsage(): number;
    memLoad(): number;
};

export = container;
