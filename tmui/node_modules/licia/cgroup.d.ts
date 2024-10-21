declare const cgroup: {
    cpu: {
        stat(): {
            usage: number;
        };
        max(): number;
    };
    cpuset: {
        cpus(): {
            effective: number[];
        };
    };
    memory: {
        max(): number;
        current(): number;
    };
    version(): number;
};

export = cgroup;
