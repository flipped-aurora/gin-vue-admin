declare function defineProp<T>(
    obj: T,
    prop: string,
    descriptor: PropertyDescriptor
): T;
declare function defineProp<T>(obj: T, descriptor: PropertyDescriptorMap): T;

export = defineProp;
