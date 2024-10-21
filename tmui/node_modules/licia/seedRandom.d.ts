declare function seedRandom(
    seed: number,
    min?: number,
    max?: number,
    floating?: boolean
): () => number;

export = seedRandom;
