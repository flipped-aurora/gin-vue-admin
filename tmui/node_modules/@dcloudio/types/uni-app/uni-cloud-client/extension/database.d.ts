declare namespace UniCloudNamespace {
  enum UPDATE_COMMANDS_LITERAL {
    SET = 'set',
    REMOVE = 'remove',
    INC = 'inc',
    MUL = 'mul',
    PUSH = 'push',
    PULL = 'pull',
    PULL_ALL = 'pullAll',
    POP = 'pop',
    SHIFT = 'shift',
    UNSHIFT = 'unshift',
    ADD_TO_SET = 'addToSet',
    BIT = 'bit',
    RENAME = 'rename',
    MAX = 'max',
    MIN = 'min'
  }

  enum QUERY_COMMANDS_LITERAL {
    EQ = 'eq',
    NEQ = 'neq',
    GT = 'gt',
    GTE = 'gte',
    LT = 'lt',
    LTE = 'lte',
    IN = 'in',
    NIN = 'nin',
    ALL = 'all',
    ELEM_MATCH = 'elemMatch',
    EXISTS = 'exists',
    SIZE = 'size',
    MOD = 'mod',
    GEO_NEAR = 'geoNear',
    GEO_WITHIN = 'geoWithin',
    GEO_INTERSECTS = 'geoIntersects'
  }

  enum LOGIC_COMMANDS_LITERAL {
    AND = 'and',
    OR = 'or',
    NOT = 'not',
    NOR = 'nor'
  }

  interface InternalSymbol {
    for(target: any): InternalSymbol;
  }

  interface LogicCommand {
    fieldName: any;
    operator: any;
    operands: any;
    _internalType: InternalSymbol;
    _setFieldName(fieldName: string): LogicCommand;
    /**
     * 查询操作符，用于表示逻辑 "与" 的关系，表示需同时满足多个查询筛选条件
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=and](https://uniapp.dcloud.io/uniCloud/cf-database?id=and)
     */
    and(__expressions__: any): LogicCommand;
    /**
     * 查询操作符，用于表示逻辑 "或" 的关系，表示需同时满足多个查询筛选条件。或指令有两种用法，一是可以进行字段值的 “或” 操作，二是也可以进行跨字段的 “或” 操作。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=or](https://uniapp.dcloud.io/uniCloud/cf-database?id=or)
     */
    or(__expressions__: any): LogicCommand;
  }
  type GetTempQuery = object;

  /**
   * 时区，'Asia/Shanghai' 或 '+08:00' 格式均支持
   */
  type Timezone = string;

  interface DateToStringOptions {
    /**
     * 日期表达式
     */
    date: string;
    /**
     * 格式化表达式
     */
    format: string;
    /**
     * 时区表达式，指明运算结果的时区。它可以解析格式为 UTC Offset 或者 Olson Timezone Identifier 的字符串。
     * - Asia/Shanghai:
     */
    timezone?: Timezone;
    /**
     * 空值表达式，可选。当日期表达式返回空或者不存在的时候，会返回此表达式指明的值。
     */
    onNull?: string;
  }

  interface DateFromStringOptions {
    /**
     * 日期字符串
     */
    dateString: string;
    /**
     * 时区，Olson Timezone Identifier
     * - Asia/Shanghai:
     */
    timezone?: Timezone;
  }

  interface DateFromPartsOptions {
    /**
     * 年份
     */
    year?: number;
    /**
     * 月份
     */
    month?: number;
    /**
     * 日期
     */
    day?: number;
    /**
     * 小时
     */
    hour?: number;
    /**
     * 分钟
     */
    minute?: number;
    /**
     * 秒
     */
    second?: number;
    /**
     * 毫秒
     */
    millisecond?: number;
    /**
     * 时区，Olson Timezone Identifier
     * - Asia/Shanghai:
     */
    timezone?: Timezone;
    /**
     * ISO标准年份
     */
    isoWeekYear?: number;
    /**
     * ISO标准一年中的第几周
     */
    isoWeek?: number;
    /**
     * ISO标准一周中的第几天（周一：0-周日：7）
     */
    isoDayOfWeek?: number;
  }

  type DBDate = object;

  interface ZipOptions {
    /**
     * 一个二维数组（inputs 不可以是字段引用），其中每个元素的表达式（这个可以是字段引用）都可以解析为数组。如果其中任意一个表达式返回 null，inputs 也返回 null。如果其中任意一个表达式不是指向一个合法的字段 / 解析为数组 / 解析为 null，则返回错误。
     */
    inputs: any;
    /**
     * 决定输出数组的长度是否采用输入数组中的最长数组的长度。默认为 false，即输入数组中的最短的数组的长度即是输出数组的各个元素的长度。
     */
    useLongestLength: any;
    /**
     * 一个数组，用于指定在输入数组长度不一的情况下时采用的数组各元素默认值。指定这个字段则必须指定 useLongestLength，否则返回错误。如果 useLongestLength 是 true 但是 defaults 是空或没有指定，则 zip 用 null 做数组元素的缺省默认值。指定各元素默认值时 defaults 数组的长度必须是输入数组最大的长度。
     */
    defaults: any;
  }

  interface ReduceOptions {
    /**
     * 输入数组，可以是任意解析为数组的表达式
     */
    input: string | any[];
    /**
     * 初始值
     */
    initialValue: string;
    /**
     * 用来作用于每个元素的表达式，在 in 中有两个可用变量，value 是表示累计值的变量，this 是表示当前数组元素的变量
     */
    in: any;
  }

  interface MapOptions {
    /**
     * 一个可以解析为数组的表达式
     */
    input: string;
    /**
     * 可选，用于表示数组各个元素的变量，默认为 this
     */
    as: string;
    /**
     * 一个可以应用在给定数组的各个元素上的表达式，各个元素的名字由 as 参数决定（参数名需加 $$ 前缀，如 $$this）
     */
    in: any;
  }

  interface FilterOptions {
    /**
     * 一个可以解析为数组的表达式
     */
    input: string;
    /**
     * 可选，用于表示数组各个元素的变量，默认为 this
     */
    as: string;
    /**
     * 一个可以解析为布尔值的表达式，用于判断各个元素是否满足条件，各个元素的名字由 as 参数决定（参数名需加 $$ 前缀，如 $$this）
     */
    cond: boolean;
  }

  interface BranchedOptions {
    /**
     * 判断条件
     */
    case: any;
    /**
     * 条件为真时执行的操作
     */
    then: any;
  }

  interface LetOptions {
    /**
     * 定义多个变量，变量的值由 变量表达式 计算而来，并且被定义的变量只有在 in 中的 结果表达式 才可以访问。
     */
    vars: any;
    /**
     * 结果表达式中访问自定义变量时候，请在变量名前加上双美元符号( $$ )并用引号括起来，如：'$$price'
     */
    in: any;
  }

  interface SwitchOptions {
    /**
     * switch操作的分支部分
     */
    branches: BranchedOptions[];
    /**
     * switch的默认操作
     */
    default: any;
  }

  interface BucketOptions {
    /**
     * 一个用以决定分组的表达式，会应用在各个输入记录上。可以用 $ 前缀加上要用以分组的字段路径来作为表达式。除非用 default 指定了默认值，否则每个记录都需要包含指定的字段，且字段值必须在 boundaries 指定的范围之内。
     */
    groupBy: string;
    /**
     * 一个数组，每个元素分别是每组的下界。必须至少指定两个边界值。数组值必须是同类型递增的值
     */
    boundaries: any[];
    /**
     * 可选，指定之后，没有进入任何分组的记录将都进入一个默认分组，这个分组记录的 _id 即由 default 决定。default 的值必须小于 boundaries 中的最小值或大于等于其中的最大值。default 的值可以与 boundaries 元素值类型不同。
     */
    default: string;
    /**
     * 可选，用以决定输出记录除了 _id 外还要包含哪些字段，各个字段的值必须用累加器表达式指定。
     */
    output: any;
  }

  interface BucketAutoOptions {
    /**
     * 一个用以决定分组的表达式，会应用在各个输入记录上。可以用 $ 前缀加上要用以分组的字段路径来作为表达式。除非用 default 指定了默认值，否则每个记录都需要包含指定的字段，且字段值必须在 boundaries 指定的范围之内。
     */
    groupBy: string;
    /**
     * 一个用于指定划分组数的正整数
     */
    buckets: number;
    /**
     * 用于保证自动计算出的边界符合给定的规则。这个字段仅可在所有 groupBy 值都是数字并且没有 NaN 的情况下使用。
     * - R5: 边界规则
     * - R10: 边界规则
     * - R20: 边界规则
     * - R40: 边界规则
     * - R80: 边界规则
     * - 1-2-5: 边界规则
     * - E6: 边界规则
     * - E12: 边界规则
     * - E24: 边界规则
     * - E48: 边界规则
     * - E96: 边界规则
     * - E192: 边界规则
     * - POWERSOF: 边界规则
     */
    granularity: 'R5' | 'R10' | 'R20' | 'R40' | 'R80' | '1-2-5' | 'E6' | 'E12' | 'E24' | 'E48' | 'E96' | 'E192' | 'POWERSOF';
    /**
     * 用以决定输出记录除了 _id 外还要包含哪些字段，各个字段的值必须用累加器表达式指定
     */
    output: any;
  }

  interface LookupOptions {
    /**
     * 要进行连接的另外一个集合的名字
     */
    from: string;
    /**
     * 当前流水线的输入记录的字段名，该字段将被用于与 from 指定的集合的 foreignField 进行相等匹配。如果输入记录中没有该字段，则该字段的值在匹配时会被视作 null
     */
    localField: string;
    /**
     * 被连接集合的字段名，该字段会被用于与 localField 进行相等匹配。如果被连接集合的记录中没有该字段，该字段的值将在匹配时被视作 null
     */
    foreignField: string;
    /**
     * 指定连接匹配出的记录列表要存放的字段名，这个数组包含的是匹配出的来自 from 集合的记录。如果输入记录中本来就已有该字段，则该字段会被覆写
     */
    as: string;
  }

  interface LookupPipelineOptions {
    /**
     * 要进行连接的另外一个集合的名字
     */
    from: string;
    /**
     * 可选。指定在 pipeline 中可以使用的变量，变量的值可以引用输入记录的字段，比如 let: { userName: '$name' } 就代表将输入记录的 name 字段作为变量 userName 的值。在 pipeline 中无法直接访问输入记录的字段，必须通过 let 定义之后才能访问，访问的方式是在 expr 操作符中用 $$变量名 的方式访问，比如 $$userName。
     */
    let: any;
    /**
     * 指定要在被连接集合中运行的聚合操作。如果要返回整个集合，则该字段取值空数组 []。在 pipeline 中无法直接访问输入记录的字段，必须通过 let 定义之后才能访问，访问的方式是在 expr 操作符中用 $$变量名 的方式访问，比如 $$userName。
     */
    pipeline: any;
    /**
     * 指定连接匹配出的记录列表要存放的字段名，这个数组包含的是匹配出的来自 from 集合的记录。如果输入记录中本来就已有该字段，则该字段会被覆写
     */
    as: string;
  }

  interface ISerializedLineString {
    type: string;
    coordinates: any;
  }

  interface ISerializedPolygon {
    type: string;
    coordinates: any;
  }

  interface ISerializedMultiPoint {
    type: string;
    coordinates: any;
  }

  interface ISerializedMultiLineString {
    type: string;
    coordinates: any;
  }

  interface ISerializedMultiPolygon {
    type: string;
    coordinates: any;
  }

  interface UnwindOptions {
    /**
     * 想要拆分的数组的字段名，需要以 $ 开头。
     */
    path: string;
    /**
     * 传入一个新的字段名，数组索引会保存在这个新的字段上。新的字段名不能以 $ 开头。
     */
    includeArrayIndex: string;
    /**
     * 如果为 true，那么在 path 对应的字段为 null、空数组或者这个字段不存在时，依然会输出这个文档；如果为 false，unwind 将不会输出这些文档。默认为 false。
     */
    preserveNullAndEmptyArrays: boolean;
  }

  interface SampleOptions {
    /**
     * 返回记录的数量
     */
    size: number;
  }

  interface ReplaceRootOptions {
    /**
     * 新的根节点
     */
    newRoot: any;
  }

  interface GetTreeParam {
    /**
     * 第一层级条件
     */
    startWith: string;
    /**
     * 最大查询层级
     */
    limitLevel: number;
  }

  interface GetTreePathParam {
    /**
     * 要查询节点的条件
     */
    startWith: string;
    /**
     * 最大查询层级
     */
    limitLevel: number;
  }

  interface GetParam {
    /**
     * 是否获取单条
     */
    getOne?: boolean;
    /**
     * 是否返回总数
     */
    getCount?: boolean;
    /**
     * 查询树形结构选定节点的所有子节点
     */
    getTree?: GetTreeParam | boolean;
    /**
     * 查询选定节点在树形结构内的路径
     */
    getTreePath?: GetTreePathParam;
  }

  interface AggregateCommand {
    /**
     * 返回一个数字的绝对值。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=abs](https://uniapp.dcloud.io/uniCloud/cf-database?id=abs)
     */
    abs(absOptions: number | string): number;
    /**
     * 将数字相加或将数字加在日期上。如果数组中的其中一个值是日期，那么其他值将被视为毫秒数加在该日期上。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=add](https://uniapp.dcloud.io/uniCloud/cf-database?id=add)
     */
    add(addOptions: any[]): number;
    /**
     * 向上取整，返回大于或等于给定数字的最小整数。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=ceil](https://uniapp.dcloud.io/uniCloud/cf-database?id=ceil)
     */
    ceil(ceilOptions: number | string): number;
    /**
     * 传入被除数和除数，求商。参数形式为：[被除数,除数]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=divide](https://uniapp.dcloud.io/uniCloud/cf-database?id=divide)
     */
    divide(divideOptions: any[]): number;
    /**
     * 取 e（自然对数的底数，欧拉数） 的 n 次方
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=exp](https://uniapp.dcloud.io/uniCloud/cf-database?id=exp)
     */
    exp(expOptions: number | string): number;
    /**
     * 向下取整，返回大于或等于给定数字的最小整数。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=floor](https://uniapp.dcloud.io/uniCloud/cf-database?id=floor)
     */
    floor(floorOptions: number | string): number;
    /**
     * 计算给定数字在自然对数值。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=ln](https://uniapp.dcloud.io/uniCloud/cf-database?id=ln)
     */
    ln(lnOptions: number | string): number;
    /**
     * 计算给定数字在给定对数底下的 log 值。参数形式为：[真数,底数]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=log](https://uniapp.dcloud.io/uniCloud/cf-database?id=log)
     */
    log(logOptions: any[]): number;
    /**
     * 计算给定数字在对数底为 10 下的 log 值。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=log10](https://uniapp.dcloud.io/uniCloud/cf-database?id=log10)
     */
    log10(log10Options: number | string): number;
    /**
     * 取模运算，取数字取模后的值。参数形式为：[被除数,除数]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=mod](https://uniapp.dcloud.io/uniCloud/cf-database?id=mod)
     */
    mod(modOptions: any[]): number;
    /**
     * 取传入的数字参数相乘的结果。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=multiply](https://uniapp.dcloud.io/uniCloud/cf-database?id=multiply)
     */
    multiply(multiplyOptions: any[]): number;
    /**
     * 求给定基数的指数次幂。参数形式为：[底数,指数]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=pow](https://uniapp.dcloud.io/uniCloud/cf-database?id=pow)
     */
    pow(powOptions: any[]): number;
    /**
     * 求平方根。参数形式为：[被开方数]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=sqrt](https://uniapp.dcloud.io/uniCloud/cf-database?id=sqrt)
     */
    sqrt(sqrtOptions: any[]): number;
    /**
     * 将两个数字相减然后返回差值，或将两个日期相减然后返回相差的毫秒数，或将一个日期减去一个数字返回结果的日期。参数形式为：[被减数,减数]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=subtract](https://uniapp.dcloud.io/uniCloud/cf-database?id=subtract)
     */
    subtract(subtractOptions: any[]): number;
    /**
     * 将数字截断为整形。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=trunc](https://uniapp.dcloud.io/uniCloud/cf-database?id=trunc)
     */
    trunc(truncOptions: number | string): number;
    /**
     * 返回在指定数组下标的元素。参数形式为：[数组,下标]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=arrayelemat](https://uniapp.dcloud.io/uniCloud/cf-database?id=arrayelemat)
     */
    arrayElemAt(arrayElemAtOptions: any[]): any;
    /**
     * 将一个数组转换为对象。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=arraytoobject](https://uniapp.dcloud.io/uniCloud/cf-database?id=arraytoobject)
     */
    arrayToObject(arrayToObjectOptions: any[] | string): any;
    /**
     * 将多个数组拼接成一个数组。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=concatarrays](https://uniapp.dcloud.io/uniCloud/cf-database?id=concatarrays)
     */
    concatArrays(concatArraysOptions: any[]): any[];
    /**
     * 根据给定条件返回满足条件的数组的子集。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=filter](https://uniapp.dcloud.io/uniCloud/cf-database?id=filter)
     */
    filter(filterOptions: FilterOptions): any[];
    /**
     * 给定一个值和一个数组，如果值在数组中则返回 true，否则返回 false。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=in](https://uniapp.dcloud.io/uniCloud/cf-database?id=in)
     */
    in(inOptions: number | string): number;
    /**
     * 在数组中找出等于给定值的第一个元素的下标，如果找不到则返回 -1。参数形式为：[数组,查找值,起始索引,结束索引]，起始索引、结束索引为可选值。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=indexofarray](https://uniapp.dcloud.io/uniCloud/cf-database?id=indexofarray)
     */
    indexOfArray(indexOfArrayOptions: any[]): number;
    /**
     * 判断给定表达式是否是数组，返回布尔值。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=isarray](https://uniapp.dcloud.io/uniCloud/cf-database?id=isarray)
     */
    isArray(isArrayOptions: any[]): boolean;
    /**
     * 类似 JavaScript Array 上的 map 方法，将给定数组的每个元素按给定转换方法转换后得出新的数组。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=map](https://uniapp.dcloud.io/uniCloud/cf-database?id=map)
     */
    map(mapOptions: MapOptions): any[];
    /**
     * 将一个对象转换为数组。方法把对象的每个键值对都变成输出数组的一个元素，元素形如 { k: <key>, v: <value> }。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=objecttoarray](https://uniapp.dcloud.io/uniCloud/cf-database?id=objecttoarray)
     */
    objectToArray(objectToArrayOptions: any): any[];
    /**
     * 返回一组生成的序列数字。给定开始值、结束值、非零的步长，range 会返回从开始值开始逐步增长、步长为给定步长、但不包括结束值的序列。参数形式为：[起始值，结束值，步长]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=range](https://uniapp.dcloud.io/uniCloud/cf-database?id=range)
     */
    range(rangeOptions: any[]): any[];
    /**
     * 类似 JavaScript 的 reduce 方法，应用一个表达式于数组各个元素然后归一成一个元素。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=reduce](https://uniapp.dcloud.io/uniCloud/cf-database?id=reduce)
     */
    reduce(reduceOptions: ReduceOptions): any;
    /**
     * 返回给定数组的倒序形式。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=reversearray](https://uniapp.dcloud.io/uniCloud/cf-database?id=reversearray)
     */
    reverseArray(reverseArrayOptions: any[] | string): any[];
    /**
     * 返回数组长度。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=size](https://uniapp.dcloud.io/uniCloud/cf-database?id=size)
     */
    size(sizeOptions: any[] | string): number;
    /**
     * 类似 JavaScritp 的 slice 方法。返回给定数组的指定子集。参数形式：[数组,下标]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=slice](https://uniapp.dcloud.io/uniCloud/cf-database?id=slice)
     */
    slice(sliceOptions: any[]): any[];
    /**
     * 把二维数组的第二维数组中的相同序号的元素分别拼装成一个新的数组进而组装成一个新的二维数组。如可将 [ [ 1, 2, 3 ], [ "a", "b", "c" ] ] 转换成 [ [ 1, "a" ], [ 2, "b" ], [ 3, "c" ] ]。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=zip](https://uniapp.dcloud.io/uniCloud/cf-database?id=zip)
     */
    zip(zipOptions: ZipOptions): any[];
    /**
     * 给定多个表达式，and 仅在所有表达式都返回 true 时返回 true，否则返回 false 。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=and](https://uniapp.dcloud.io/uniCloud/cf-database?id=and)
     */
    and(andOptions: any[]): boolean;
    /**
     * 给定一个表达式，如果表达式返回 true，则 not 返回 false，否则返回 true。注意表达式不能为逻辑表达式（and、or、nor、not）。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=not](https://uniapp.dcloud.io/uniCloud/cf-database?id=not)
     */
    not(notOptions: string): boolean;
    /**
     * 给定多个表达式，如果任意一个表达式返回 true，则 or 返回 true，否则返回 false。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=or](https://uniapp.dcloud.io/uniCloud/cf-database?id=or)
     */
    or(orOptions: any[]): boolean;
    /**
     * 给定两个值，返回其比较值。如果第一个值小于第二个值，返回 -1 。如果第一个值大于第二个值，返回 1 。 如果两个值相等，返回 0 。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=cmp](https://uniapp.dcloud.io/uniCloud/cf-database?id=cmp)
     */
    cmp(cmpOptions: any[]): number;
    /**
     * 匹配两个值，如果相等则返回 true，否则返回 false。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=eq](https://uniapp.dcloud.io/uniCloud/cf-database?id=eq)
     */
    eq(eqOptions: any[]): boolean;
    /**
     * 匹配两个值，如果前者大于后者则返回 true，否则返回 false。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=gt](https://uniapp.dcloud.io/uniCloud/cf-database?id=gt)
     */
    gt(gtOptions: any[]): boolean;
    /**
     * 匹配两个值，如果前者大于或等于后者则返回 true，否则返回 false。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=gte](https://uniapp.dcloud.io/uniCloud/cf-database?id=gte)
     */
    gte(gteOptions: any[]): boolean;
    /**
     * 匹配两个值，如果前者小于后者则返回 true，否则返回 false。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=lt](https://uniapp.dcloud.io/uniCloud/cf-database?id=lt)
     */
    lt(ltOptions: any[]): boolean;
    /**
     * 匹配两个值，如果前者小于或等于后者则返回 true，否则返回 false。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=lte](https://uniapp.dcloud.io/uniCloud/cf-database?id=lte)
     */
    lte(lteOptions: any[]): boolean;
    /**
     * 匹配两个值，如果不相等则返回 true，否则返回 false。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=neq](https://uniapp.dcloud.io/uniCloud/cf-database?id=neq)
     */
    neq(neqOptions: any[]): boolean;
    /**
     * 计算布尔表达式，返回指定的两个值其中之一。参数形式为：[布尔表达式,真值,假值]，效果类似于 javascript 里的 condition?a:b
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=cond](https://uniapp.dcloud.io/uniCloud/cf-database?id=cond)
     */
    cond(condOptions: any[]): any;
    /**
     * 计算给定的表达式，如果表达式结果为 null、undefined 或者不存在，那么返回一个替代值；否则返回原值。参数形式为：[表达式,替代值]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=ifnull](https://uniapp.dcloud.io/uniCloud/cf-database?id=ifnull)
     */
    ifNull(ifNullOptions: any[]): any;
    /**
     * 根据给定的 switch-case-default 计算返回值。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=switch](https://uniapp.dcloud.io/uniCloud/cf-database?id=switch)
     */
    switch(switchOptions: SwitchOptions): boolean;
    /**
     * 给定日期的相关信息，构建并返回一个日期对象。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=datefromparts](https://uniapp.dcloud.io/uniCloud/cf-database?id=datefromparts)
     */
    dateFromParts(dateFromPartsOptions: DateFromPartsOptions): DBDate;
    /**
     * 将一个日期/时间字符串转换为日期对象。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=datefromstring](https://uniapp.dcloud.io/uniCloud/cf-database?id=datefromstring)
     */
    dateFromString(dateFromStringOptions: DateFromStringOptions): DBDate;
    /**
     * 根据指定的表达式将日期对象格式化为符合要求的字符串。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=datetostring](https://uniapp.dcloud.io/uniCloud/cf-database?id=datetostring)
     */
    dateToString(dateToStringOptions: DateToStringOptions): string;
    /**
     * 返回日期字段对应的天数（一个月中的哪一天），是一个介于 1 至 31 之间的数字。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=dayofmonth](https://uniapp.dcloud.io/uniCloud/cf-database?id=dayofmonth)
     */
    dayOfMonth(dayOfMonthOptions: string): number;
    /**
     * 返回日期字段对应的天数（一周中的第几天），是一个介于 1（周日）到 7（周六）之间的整数。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=dayofweek](https://uniapp.dcloud.io/uniCloud/cf-database?id=dayofweek)
     */
    dayOfWeek(dayOfWeekOptions: string): number;
    /**
     * 返回日期字段对应的天数（一年中的第几天），是一个介于 1 到 366 之间的整数。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=dayofyear](https://uniapp.dcloud.io/uniCloud/cf-database?id=dayofyear)
     */
    dayOfYear(dayOfYearOptions: string): number;
    /**
     * 返回日期字段对应的小时数，是一个介于 0 到 23 之间的整数。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=hour](https://uniapp.dcloud.io/uniCloud/cf-database?id=hour)
     */
    hour(hourOptions: string): number;
    /**
     * 返回日期字段对应的 ISO 8601 标准的天数（一周中的第几天），是一个介于 1（周一）到 7（周日）之间的整数。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=isodayofweek](https://uniapp.dcloud.io/uniCloud/cf-database?id=isodayofweek)
     */
    isoDayOfWeek(isoDayOfWeekOptions: string): number;
    /**
     * 返回日期字段对应的 ISO 8601 标准的周数（一年中的第几周），是一个介于 1 到 53 之间的整数。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=isoweek](https://uniapp.dcloud.io/uniCloud/cf-database?id=isoweek)
     */
    isoWeek(isoWeekOptions: string): void;
    /**
     * 返回日期字段对应的 ISO 8601 标准的天数（一年中的第几天）。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=isoweekyear](https://uniapp.dcloud.io/uniCloud/cf-database?id=isoweekyear)
     */
    isoWeekYear(isoWeekYearOptions: string): number;
    /**
     * 返回日期字段对应的毫秒数，是一个介于 0 到 999 之间的整数。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=millisecond](https://uniapp.dcloud.io/uniCloud/cf-database?id=millisecond)
     */
    millisecond(millisecondOptions: string): number;
    /**
     * 返回日期字段对应的分钟数，是一个介于 0 到 59 之间的整数。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=minute](https://uniapp.dcloud.io/uniCloud/cf-database?id=minute)
     */
    minute(minuteOptions: string): number;
    /**
     * 返回日期字段对应的月份，是一个介于 1 到 12 之间的整数。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=month](https://uniapp.dcloud.io/uniCloud/cf-database?id=month)
     */
    month(monthOptions: string): number;
    /**
     * 返回日期字段对应的秒数，是一个介于 0 到 59 之间的整数，在特殊情况下（闰秒）可能等于 60。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=second](https://uniapp.dcloud.io/uniCloud/cf-database?id=second)
     */
    second(secondOptions: string): number;
    /**
     * 返回日期字段对应的周数（一年中的第几周），是一个介于 0 到 53 之间的整数。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=week](https://uniapp.dcloud.io/uniCloud/cf-database?id=week)
     */
    week(weekOptions: string): number;
    /**
     * 返回日期字段对应的年份。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=year](https://uniapp.dcloud.io/uniCloud/cf-database?id=year)
     */
    year(yearOptions: string): number;
    /**
     * 直接返回一个值的字面量，不经过任何解析和处理。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=literal](https://uniapp.dcloud.io/uniCloud/cf-database?id=literal)
     */
    literal(literalOptions: string): string;
    /**
     * 输入一个数组，或者数组字段的表达式。如果数组中所有元素均为真值，那么返回 true，否则返回 false。空数组永远返回 true。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=allelementstrue](https://uniapp.dcloud.io/uniCloud/cf-database?id=allelementstrue)
     */
    allElementsTrue(allElementsTrueOptions: any[]): boolean;
    /**
     * 输入一个数组，或者数组字段的表达式。如果数组中任意一个元素为真值，那么返回 true，否则返回 false。空数组永远返回 false。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=anyelementtrue](https://uniapp.dcloud.io/uniCloud/cf-database?id=anyelementtrue)
     */
    anyElementTrue(anyElementTrueOptions: any[]): boolean;
    /**
     * 输入两个集合，输出只存在于第一个集合中的元素。。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=setdifference](https://uniapp.dcloud.io/uniCloud/cf-database?id=setdifference)
     */
    setDifference(setDifferenceOptions: any[]): any;
    /**
     * 输入两个集合，判断两个集合中包含的元素是否相同（不考虑顺序、去重）。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=setequals](https://uniapp.dcloud.io/uniCloud/cf-database?id=setequals)
     */
    setEquals(setEqualsOptions: any[]): boolean;
    /**
     * 输入两个集合，判断两个集合中包含的元素是否相同（不考虑顺序、去重）。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=setintersection](https://uniapp.dcloud.io/uniCloud/cf-database?id=setintersection)
     */
    setIntersection(setIntersectionOptions: any[]): any;
    /**
     * 输入两个集合，判断第一个集合是否是第二个集合的子集。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=setissubset](https://uniapp.dcloud.io/uniCloud/cf-database?id=setissubset)
     */
    setIsSubset(setIsSubsetOptions: any[]): any;
    /**
     * 输入两个集合，输出两个集合的并集。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=setunion](https://uniapp.dcloud.io/uniCloud/cf-database?id=setunion)
     */
    setUnion(setUnionOptions: any[]): any;
    /**
     * 连接字符串，返回拼接后的字符串。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=concat](https://uniapp.dcloud.io/uniCloud/cf-database?id=concat)
     */
    concat(concatOptions: any[]): string;
    /**
     * 在目标字符串中查找子字符串，并返回第一次出现的 UTF-8 的字节索引（从0开始）。如果不存在子字符串，返回 -1。参数形式为：[目标字符串表达式, 子字符串表达式, 开始位置, 结束位置]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=indexofbytes](https://uniapp.dcloud.io/uniCloud/cf-database?id=indexofbytes)
     */
    indexOfBytes(indexOfBytesOptions: any[]): number;
    /**
     * 在目标字符串中查找子字符串，并返回第一次出现的 UTF-8 的 code point 索引（从0开始）。如果不存在子字符串，返回 -1。参数形式为：[目标字符串表达式, 子字符串表达式, 开始位置, 结束位置]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=indexofcp](https://uniapp.dcloud.io/uniCloud/cf-database?id=indexofcp)
     */
    indexOfCP(indexOfCPOptions: any[]): number;
    /**
     * 按照分隔符分隔字符串，并且删除分隔符，返回子字符串组成的数组。如果字符串无法找到分隔符进行分隔，返回原字符串作为数组的唯一元素。参数形式为：[字符串,分隔符]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=split](https://uniapp.dcloud.io/uniCloud/cf-database?id=split)
     */
    split(splitOptions: string): any[];
    /**
     * 计算并返回指定字符串中 utf-8 编码的字节数量。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=strlenbytes](https://uniapp.dcloud.io/uniCloud/cf-database?id=strlenbytes)
     */
    strLenBytes(strLenBytesOptions: string): number;
    /**
     * 计算并返回指定字符串的UTF-8 code points 数量。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=strlencp](https://uniapp.dcloud.io/uniCloud/cf-database?id=strlencp)
     */
    strLenCP(strLenCPOptions: string): number;
    /**
     * 对两个字符串在不区分大小写的情况下进行大小比较，并返回比较的结果。如果第一个值小于第二个值，返回 -1 。如果第一个值大于第二个值，返回 1 。 如果两个值相等，返回 0 。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=strcasecmp](https://uniapp.dcloud.io/uniCloud/cf-database?id=strcasecmp)
     */
    strcasecmp(strcasecmpOptions: any[]): number;
    /**
     * 返回字符串从指定位置开始的指定长度的子字符串。它是 db.command.aggregate.substrBytes 的别名，更推荐使用后者。参数形式为：[字符串,起始位置,结束位置]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=substr](https://uniapp.dcloud.io/uniCloud/cf-database?id=substr)
     */
    substr(substrOptions: any[]): string;
    /**
     * 返回字符串从指定位置开始的指定长度的子字符串。子字符串是由字符串中指定的 UTF-8 字节索引的字符开始，长度为指定的字节数。参数形式为：[字符串,起始位置,结束位置]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=substrbytes](https://uniapp.dcloud.io/uniCloud/cf-database?id=substrbytes)
     */
    substrBytes(substrBytesOptions: any[]): string;
    /**
     * 返回字符串从指定位置开始的指定长度的子字符串。子字符串是由字符串中指定的 UTF-8 字节索引的字符开始，长度为指定的字节数。参数形式为：[字符串,起始位置,结束位置]
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=substrcp](https://uniapp.dcloud.io/uniCloud/cf-database?id=substrcp)
     */
    substrCP(substrCPOptions: any[]): string;
    /**
     * 将字符串转化为小写并返回。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=tolower](https://uniapp.dcloud.io/uniCloud/cf-database?id=tolower)
     */
    toLower(toLowerOptions: string): string;
    /**
     * 将字符串转化为大写并返回。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=toupper](https://uniapp.dcloud.io/uniCloud/cf-database?id=toupper)
     */
    toUpper(toUpperOptions: string): string;
    /**
     * 向数组中添加值，如果数组中已存在该值，不执行任何操作。它只能在 group 阶段中使用。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=addtoset](https://uniapp.dcloud.io/uniCloud/cf-database?id=addtoset)
     */
    addToSet(addToSetOptions: string): void;
    /**
     * 返回一组集合中，指定字段对应数据的平均值。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=avg](https://uniapp.dcloud.io/uniCloud/cf-database?id=avg)
     */
    avg(avgOptions: string): void;
    /**
     * 返回指定字段在一组集合的第一条记录对应的值。仅当这组集合是按照某种定义排序（ sort ）后，此操作才有意义。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=first](https://uniapp.dcloud.io/uniCloud/cf-database?id=first)
     */
    first(firstOptions: string): void;
    /**
     * 返回指定字段在一组集合的最后一条记录对应的值。仅当这组集合是按照某种定义排序（ sort ）后，此操作才有意义。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=last](https://uniapp.dcloud.io/uniCloud/cf-database?id=last)
     */
    last(lastOptions: string): void;
    /**
     * 返回一组数值的最大值。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=max](https://uniapp.dcloud.io/uniCloud/cf-database?id=max)
     */
    max(maxOptions: string): void;
    /**
     * 返回一组数值的最小值。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=min](https://uniapp.dcloud.io/uniCloud/cf-database?id=min)
     */
    min(minOptions: string): void;
    /**
     * 将多个文档合并为单个文档
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=mergeobjects](https://uniapp.dcloud.io/uniCloud/cf-database?id=mergeobjects)
     */
    mergeObjects(mergeObjectsOptions: string | any[]): void;
    /**
     * 在 group 阶段，返回一组中表达式指定列与对应的值，一起组成的数组。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=push-1](https://uniapp.dcloud.io/uniCloud/cf-database?id=push-1)
     */
    push(pushOptions: any): void;
    /**
     * 返回一组字段对应值的标准差。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=stddevpop](https://uniapp.dcloud.io/uniCloud/cf-database?id=stddevpop)
     */
    stdDevPop(stdDevPopOptions: string): void;
    /**
     * 计算输入值的样本标准偏差。如果输入值代表数据总体，或者不概括更多的数据，请改用 db.command.aggregate.stdDevPop 。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=stddevsamp](https://uniapp.dcloud.io/uniCloud/cf-database?id=stddevsamp)
     */
    stdDevSamp(stdDevSampOptions: string): void;
    /**
     * 计算并且返回一组字段所有数值的总和。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=sum](https://uniapp.dcloud.io/uniCloud/cf-database?id=sum)
     */
    sum(sumOptions: string): void;
    /**
     * 自定义变量，并且在指定表达式中使用，返回的结果是表达式的结果。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=let](https://uniapp.dcloud.io/uniCloud/cf-database?id=let)
     */
    let(letOptions: LetOptions): void;
  }

  interface ISerializedPoint {
    type: string;
    coordinates: any;
  }

  interface Point {
    /**
     * 纬度
     */
    latitude: number;
    /**
     * 经度
     */
    longitude: number;
    _internalType: any;
    parse(key: any): any;
    /**
     * 返回相应的 GeoJSON 结构的对象
     */
    toJSON(): any;
    /**
     * 转换成可读字符串
     */
    toReadableString(): string;
    validate(point: ISerializedPoint): boolean;
  }

  interface IGeoNearOptions {
    /**
     * 地理位置点 (Point)
     */
    geometry: Point;
    /**
     * 选填，最大距离，单位为米
     */
    maxDistance: number;
    /**
     * 选填，最小距离，单位为米
     */
    minDistance: number;
  }

  interface IGeoWithinOptions {
    /**
     * 地理信息结构，Polygon，MultiPolygon，或 { centerSphere }
     */
    geometry: any;
  }

  interface IGeoIntersectsOptions {
    /**
     * 地理信息结构，Point
     */
    geometry: any;
  }

  interface GeoType {
    /**
     * 数据库地理位置结构集
     */
    Point: any;
  }

  interface QueryCommand {
    operator: QUERY_COMMANDS_LITERAL;
    /**
     * 聚合操作符
     */
    aggregate: AggregateCommand;
    _setFieldName(fieldName: string): QueryCommand;
    /**
     * 用于设定字段等于指定值。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=set](https://uniapp.dcloud.io/uniCloud/cf-database?id=set)
     */
    set(): void;
    /**
     * 用于指示字段自增某个值。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=inc](https://uniapp.dcloud.io/uniCloud/cf-database?id=inc)
     */
    inc(): void;
    /**
     * 用于指示字段自乘某个值。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=mul](https://uniapp.dcloud.io/uniCloud/cf-database?id=mul)
     */
    mul(): void;
    /**
     * 向数组尾部追加元素，支持传入单个元素或数组
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=push](https://uniapp.dcloud.io/uniCloud/cf-database?id=push)
     */
    push(): void;
    /**
     * 删除数组尾部元素
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=pop](https://uniapp.dcloud.io/uniCloud/cf-database?id=pop)
     */
    pop(): void;
    /**
     * 向数组头部添加元素，支持传入单个元素或数组
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=unshift](https://uniapp.dcloud.io/uniCloud/cf-database?id=unshift)
     */
    unshift(): void;
    /**
     * 删除数组头部元素
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=set](https://uniapp.dcloud.io/uniCloud/cf-database?id=set)
     */
    shift(): void;
    /**
     * 查询筛选条件，表示字段等于某个值。eq 指令接受一个字面量 (literal)，可以是 number, boolean, string, object, array, Date
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=eq](https://uniapp.dcloud.io/uniCloud/cf-database?id=eq)
     */
    eq(val: any): LogicCommand;
    /**
     * 查询筛选条件，表示字段不等于某个值。eq 指令接受一个字面量 (literal)，可以是 number, boolean, string, object, array, Date。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=neq](https://uniapp.dcloud.io/uniCloud/cf-database?id=neq)
     */
    neq(val: any): LogicCommand;
    /**
     * 查询筛选操作符，表示需大于指定值。可以传入 Date 对象用于进行日期比较。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=gt](https://uniapp.dcloud.io/uniCloud/cf-database?id=gt)
     */
    gt(val: any): LogicCommand;
    /**
     * 查询筛选操作符，表示需大于或等于指定值。可以传入 Date 对象用于进行日期比较。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=gte](https://uniapp.dcloud.io/uniCloud/cf-database?id=gte)
     */
    gte(val: any): LogicCommand;
    /**
     * 查询筛选操作符，表示需小于指定值。可以传入 Date 对象用于进行日期比较。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=lt](https://uniapp.dcloud.io/uniCloud/cf-database?id=lt)
     */
    lt(val: any): LogicCommand;
    /**
     * 查询筛选操作符，表示需小于或等于指定值。可以传入 Date 对象用于进行日期比较。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=lte](https://uniapp.dcloud.io/uniCloud/cf-database?id=lte)
     */
    lte(val: any): LogicCommand;
    /**
     * 查询筛选操作符，表示要求值在给定的数组内。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=in](https://uniapp.dcloud.io/uniCloud/cf-database?id=in)
     */
    in(list: any): LogicCommand;
    /**
     * 查询筛选操作符，表示要求值不在给定的数组内。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=nin](https://uniapp.dcloud.io/uniCloud/cf-database?id=nin)
     */
    nin(list: any): LogicCommand;
    /**
     * 查询操作符，用于表示逻辑 "或" 的关系，表示需同时满足多个查询筛选条件。或指令有两种用法，一是可以进行字段值的 “或” 操作，二是也可以进行跨字段的 “或” 操作。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=or-1](https://uniapp.dcloud.io/uniCloud/cf-database?id=or-1)
     */
    or(val: any): LogicCommand;
    /**
     * 查询操作符，用于表示逻辑 "与" 的关系，表示需同时满足多个查询筛选条件。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=and-1](https://uniapp.dcloud.io/uniCloud/cf-database?id=and-1)
     */
    and(val: any): LogicCommand;
    /**
     * 按从近到远的顺序，找出字段值在给定点的附近的记录。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=geonear](https://uniapp.dcloud.io/uniCloud/cf-database?id=geonear)
     */
    geoNear(val: IGeoNearOptions): LogicCommand;
    /**
     * 找出字段值在指定区域内的记录，无排序。指定的区域必须是多边形（Polygon）或多边形集合（MultiPolygon）。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=geowithin](https://uniapp.dcloud.io/uniCloud/cf-database?id=geowithin)
     */
    geoWithin(val: IGeoWithinOptions): LogicCommand;
    /**
     * 找出给定的地理位置图形相交的记录
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=geointersects](https://uniapp.dcloud.io/uniCloud/cf-database?id=geointersects)
     */
    geoIntersects(val: IGeoIntersectsOptions): LogicCommand;
  }

  interface Query {
    /**
     * 获取记录数据，或获取根据查询条件筛选后的记录数据
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=query](https://uniapp.dcloud.io/uniCloud/cf-database?id=query)
     */
    get(GetParam?: GetParam): Promise<any>;
    /**
     * 延迟查询请求
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=multi-send](https://uniapp.dcloud.io/uniCloud/jql?id=multi-send)
     */
    getTemp(GetParam?: GetParam): GetTempQuery;
    /**
     * 统计匹配查询条件的记录的条数
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=count](https://uniapp.dcloud.io/uniCloud/cf-database?id=count)
     */
    count(): Promise<any>;
    /**
     * 指定查询条件，返回带新查询条件的新的集合引用
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=where](https://uniapp.dcloud.io/uniCloud/cf-database?id=where)
     */
    where(query: any): Query;
    /**
     * 指定查询排序条件
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=order-by](https://uniapp.dcloud.io/uniCloud/cf-database?id=order-by)
     */
    orderBy(fieldPath: string | string.DBFieldString, directionStr: 'desc' | 'asc'): Query;
    /**
     * 指定查询结果集数量上限
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=limit](https://uniapp.dcloud.io/uniCloud/cf-database?id=limit)
     */
    limit(limit: number): Query;
    /**
     * 指定查询返回结果时从指定序列后的结果开始返回，常用于分页
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=skip](https://uniapp.dcloud.io/uniCloud/cf-database?id=skip)
     */
    skip(offset: number): Query;
    /**
     * 更新多条记录
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=update](https://uniapp.dcloud.io/uniCloud/cf-database?id=update)
     */
    update(data: any): Promise<any>;
    /**
     * 指定返回结果中记录需返回的字段
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=field](https://uniapp.dcloud.io/uniCloud/cf-database?id=field)
     */
    field(projection: any): Query;
    /**
     * 指定要使用的foreignKey
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=lookup-foreign-key](https://uniapp.dcloud.io/uniCloud/jql?id=lookup-foreign-key)
     */
    foreignKey(foreignKeyName: string): Query;
    /**
     * 指定分组依据
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=groupby](https://uniapp.dcloud.io/uniCloud/jql?id=groupby)
     */
    groupBy(projection: string | string.DBFieldString): Query;
    /**
     * 指定统计指标
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=groupby](https://uniapp.dcloud.io/uniCloud/jql?id=groupby)
     */
    groupField(projection: string | string.DBFieldString): Query;
    /**
     * 删除多条记录。注意只支持通过匹配 where 语句来删除，不支持 skip 和 limit
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=remove](https://uniapp.dcloud.io/uniCloud/cf-database?id=remove)
     */
    remove(): Promise<any>;
  }

  interface AggregateReference {
    /**
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-add-fields](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-add-fields)
     */
    addFields(addFieldsOptions: any): AggregateReference;
    /**
     * 将输入记录根据给定的条件和边界划分成不同的组，每组即一个bucket
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-bucket](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-bucket)
     */
    bucket(bucketOptions: BucketOptions): AggregateReference;
    /**
     * 将输入记录根据给定的条件划分成不同的组，每组即一个 bucket。与 bucket 的其中一个不同之处在于无需指定 boundaries，bucketAuto 会自动尝试将记录尽可能平均的分散到每组中。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-bucket-auto](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-bucket-auto)
     */
    bucketAuto(bucketAutoOptions: BucketAutoOptions): AggregateReference;
    /**
     * 计算上一聚合阶段输入到本阶段的记录数，输出一个记录，其中指定字段的值为记录数。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-count](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-count)
     */
    count(countOptions: string): AggregateReference;
    /**
     * 将输入记录按给定表达式分组，输出时每个记录代表一个分组，每个记录的 _id 是区分不同组的 key。输出记录中也可以包括累计值，将输出字段设为累计值即会从该分组中计算累计值。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-group](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-group)
     */
    group(groupOptions: any): AggregateReference;
    /**
     * 限制输出到下一阶段的记录数。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-limit](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-limit)
     */
    limit(limitOptions: number): AggregateReference;
    /**
     * 联表查询。与同个数据库下的一个指定的集合做 left outer join(左外连接)。对该阶段的每一个输入记录，lookup 会在该记录中增加一个数组字段，该数组是被联表中满足匹配条件的记录列表。lookup 会将连接后的结果输出给下个阶段。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-lookup](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-lookup)
     */
    lookup(lookupOptions: LookupOptions | LookupPipelineOptions): AggregateReference;
    /**
     * 根据条件过滤文档，并且把符合条件的文档传递给下一个流水线阶段。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-match](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-match)
     */
    match(matchOptions: any): AggregateReference;
    /**
     * 把指定的字段传递给下一个流水线，指定的字段可以是某个已经存在的字段，也可以是计算出来的新字段。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-project](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-project)
     */
    project(projectOptions: any): AggregateReference;
    /**
     * 指定一个已有字段作为输出的根节点，也可以指定一个计算出的新字段作为根节点。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-replace-root](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-replace-root)
     */
    replaceRoot(replaceRootOptions: ReplaceRootOptions): AggregateReference;
    /**
     * 随机从文档中选取指定数量的记录。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-sample](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-sample)
     */
    sample(sampleOptions: SampleOptions): AggregateReference;
    /**
     * 指定一个正整数，跳过对应数量的文档，输出剩下的文档。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-skip](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-skip)
     */
    skip(skipOptions: number): AggregateReference;
    /**
     * 根据指定的字段，对输入的文档进行排序。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-sort](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-sort)
     */
    sort(sortOptions: any): AggregateReference;
    /**
     * 根据指定的字段，对输入的文档进行排序。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-sort-by-count](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-sort-by-count)
     */
    sortByCount(sortByCountOptions: string): AggregateReference;
    /**
     * 使用指定的数组字段中的每个元素，对文档进行拆分。拆分后，文档会从一个变为一个或多个，分别对应数组的每个元素。
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-unwind](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-unwind)
     */
    unwind(unwindOptions: UnwindOptions | string): void;
    /**
     * 标志聚合操作定义完成，发起实际聚合操作
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-end](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate-end)
     */
    end(): Promise<any>;
  }

  interface DocumentReference {
    /**
     * 文档ID
     */
    id: string;
    /**
     * 返回结果中记录需返回的字段
     */
    projection: any;
    /**
     * 替换更新一条记录
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=doc-set](https://uniapp.dcloud.io/uniCloud/cf-database?id=doc-set)
     */
    set(data: any): Promise<any>;
    /**
     * 更新一条记录
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=where-update](https://uniapp.dcloud.io/uniCloud/cf-database?id=where-update)
     */
    update(data: any): Promise<any>;
    /**
     * 删除一条记录
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=remove](https://uniapp.dcloud.io/uniCloud/cf-database?id=remove)
     */
    remove(): Promise<any>;
    /**
     * 获取记录数据，或获取根据查询条件筛选后的记录数据
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=query](https://uniapp.dcloud.io/uniCloud/cf-database?id=query)
     */
    get(GetParam?: GetParam): Promise<any>;
    /**
     * 获取记录数据，或获取根据查询条件筛选后的记录数据
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=multi-send](https://uniapp.dcloud.io/uniCloud/jql?id=multi-send)
     */
    getTemp(GetParam?: GetParam): GetTempQuery;
    /**
     * 指定返回结果中记录需返回的字段
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=field](https://uniapp.dcloud.io/uniCloud/cf-database?id=field)
     */
    field(projection: string | string.DBFieldString): DocumentReference;
    /**
     * 指定要使用的foreignKey
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=lookup-foreign-key](https://uniapp.dcloud.io/uniCloud/jql?id=lookup-foreign-key)
     */
    foreignKey(foreignKeyName: string): DocumentReference;
    /**
     * 指定返回结果中记录需返回的字段
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=groupby](https://uniapp.dcloud.io/uniCloud/jql?id=groupby)
     */
    preField(projection: string | string.DBFieldString): Query;
    /**
     * 指定分组依据
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=groupby](https://uniapp.dcloud.io/uniCloud/jql?id=groupby)
     */
    groupBy(projection: string | string.DBFieldString): Query;
    /**
     * 指定统计指标
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=groupby](https://uniapp.dcloud.io/uniCloud/jql?id=groupby)
     */
    groupField(projection: string | string.DBFieldString): Query;
  }

  interface CollectionReference {
    name: string;
    /**
     * 获取集合中指定记录的引用。方法接受一个 id 参数，指定需引用的记录的 _id
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=doc](https://uniapp.dcloud.io/uniCloud/cf-database?id=doc)
     */
    doc(docID: string): DocumentReference;
    /**
     * 新增记录，如果传入的记录对象没有 _id 字段，则由后台自动生成 _id；若指定了 _id，则不能与已有记录冲突
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=add](https://uniapp.dcloud.io/uniCloud/cf-database?id=add)
     */
    add(data: any): Promise<any>;
    /**
     * 获取记录数据，或获取根据查询条件筛选后的记录数据
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=query](https://uniapp.dcloud.io/uniCloud/cf-database?id=query)
     */
    get(GetParam?: GetParam): Promise<any>;
    /**
     * 延迟查询请求
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=multi-send](https://uniapp.dcloud.io/uniCloud/jql?id=multi-send)
     */
    getTemp(GetParam?: GetParam): GetTempQuery;
    /**
     * 统计匹配查询条件的记录的条数
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=count](https://uniapp.dcloud.io/uniCloud/cf-database?id=count)
     */
    count(): any;
    /**
     * 指定查询条件，返回带新查询条件的新的集合引用
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=where](https://uniapp.dcloud.io/uniCloud/cf-database?id=where)
     */
    where(query: any): Query;
    /**
     * 指定查询排序条件
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=order-by](https://uniapp.dcloud.io/uniCloud/cf-database?id=order-by)
     */
    orderBy(fieldPath: string | string.DBFieldString, directionStr: 'desc' | 'asc'): Query;
    /**
     * 指定查询结果集数量上限
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=limit](https://uniapp.dcloud.io/uniCloud/cf-database?id=limit)
     */
    limit(limit: number): Query;
    /**
     * 指定查询返回结果时从指定序列后的结果开始返回，常用于分页
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=skip](https://uniapp.dcloud.io/uniCloud/cf-database?id=skip)
     */
    skip(offset: number): Query;
    /**
     * 更新多条记录
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=update](https://uniapp.dcloud.io/uniCloud/cf-database?id=update)
     */
    update(data: any): Promise<any>;
    /**
     * 指定返回结果中记录需返回的字段
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=field](https://uniapp.dcloud.io/uniCloud/cf-database?id=field)
     */
    field(projection: string | string.DBFieldString): Query;
    /**
     * 指定要使用的foreignKey
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=lookup-foreign-key](https://uniapp.dcloud.io/uniCloud/jql?id=lookup-foreign-key)
     */
    foreignKey(foreignKeyName: string): Query;
    /**
     * 指定返回结果中记录需返回的字段
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=groupby](https://uniapp.dcloud.io/uniCloud/jql?id=groupby)
     */
    preField(projection: string | string.DBFieldString): Query;
    /**
     * 指定分组依据
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=groupby](https://uniapp.dcloud.io/uniCloud/jql?id=groupby)
     */
    groupBy(projection: string | string.DBFieldString): Query;
    /**
     * 指定统计指标
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=groupby](https://uniapp.dcloud.io/uniCloud/jql?id=groupby)
     */
    groupField(projection: string | string.DBFieldString): Query;
    /**
     * 删除多条记录。注意只支持通过匹配 where 语句来删除，不支持 skip 和 limit
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=remove](https://uniapp.dcloud.io/uniCloud/cf-database?id=remove)
     */
    remove(): Promise<any>;
    /**
     * 获取数据库集合的聚合操作实例
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate](https://uniapp.dcloud.io/uniCloud/cf-database?id=aggregate)
     */
    aggregate(): AggregateReference;
  }

  type DatabaseEventName = 'refreshToken' | 'error';
  interface Database {
    /**
     * 监听数据库事件
     *
     * 文档: [https://uniapp.dcloud.net.cn/uniCloud/clientdb.html#event](https://uniapp.dcloud.net.cn/uniCloud/clientdb.html#event)
     */
    on(event: DatabaseEventName, callback: (result: any) => any): void;
    /**
     * 数据库地理位置结构集
     */
    Geo: GeoType;
    /**
     * 数据库操作符
     */
    command: QueryCommand;
    /**
     * 正则表达式查询
     */
    RegExp: any;
    /**
     * 创建一个服务端当前时间的标记
     */
    serverDate: any;
    config: any;
    /**
     * 数据库集合引用
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-database?id=collection](https://uniapp.dcloud.io/uniCloud/cf-database?id=collection)
     */
    collection(...collections: (GetTempQuery | string | string.DBCollectionString)[]): CollectionReference;
    /**
     * 指定数据库操作需要执行的action
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=action](https://uniapp.dcloud.io/uniCloud/jql?id=action)
     */
    action(actionName: string): Database;
    /**
     * 获取云端环境变量
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=variable](https://uniapp.dcloud.io/uniCloud/jql?id=variable)
     */
    getCloudEnv(envStr: '$cloudEnv_uid' | '$cloudEnv_now' | '$cloudEnv_clientIP'): any;
    /**
     * 同时执行多次数据库查询操作
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/jql?id=multi-send](https://uniapp.dcloud.io/uniCloud/jql?id=multi-send)
     */
    multiSend(tempQuery1: GetTempQuery, tempQuery2?: GetTempQuery, tempQuery3?: GetTempQuery, tempQuery4?: GetTempQuery, tempQuery5?: GetTempQuery, tempQuery6?: GetTempQuery, tempQuery7?: GetTempQuery, tempQuery8?: GetTempQuery): Promise<any>;
  }

  interface UniCloud {
    /**
     * 获取数据库实例
     *
     * 文档: [https://uniapp.dcloud.net.cn/uniCloud/jql.html#jssdk](https://uniapp.dcloud.net.cn/uniCloud/jql.html#jssdk)
     */
    database(options?: UniCloudOptions): Database;
    /**
     * 获取数据库实例，返回结果与database方法不同，少了一层result，和云端databaseForJQL方法一致
     *
     * 文档: [https://uniapp.dcloud.net.cn/uniCloud/jql.html#jssdk](https://uniapp.dcloud.net.cn/uniCloud/jql.html#jssdk)
     */
    databaseForJQL(options?: UniCloudOptions): Database;
  }
}
