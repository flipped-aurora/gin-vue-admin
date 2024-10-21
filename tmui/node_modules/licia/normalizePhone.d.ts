declare function normalizePhone(
    phone: string,
    options: {
        countryCode: number;
        trunkPrefix?: boolean;
    }
): string;

export = normalizePhone;
