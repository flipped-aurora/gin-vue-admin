import {BrowserProvider, Signer} from 'ethers';
export interface Userinfo {
    wallet?:myWallet;
    jwtToken: string;
    cliUser?: CliUser;
    cliLoad?: CliLoad;
    mainorder?:Mainorder;
}

export interface myWallet {
    address: string;
    bnb: string;
    usdt: string;
    signer?: Signer;
    provider?: BrowserProvider;
    usdtContract?: any;
    mengContract?: any;
}

export interface CliUser {
    address: string; // 用户地址
    nickname: string; // 用户昵称
    avatarurl: string; // 远程头像
    avatarlocal: string; // 本地头像
    parent: string; // 上级地址
    pullnum?: number; // 直推个数
    teamnum?: number; // 团队总数
    mypull: any[]; // 我的直推
    myparents: any[];
    status: string; // 当前状态
    desc: string; // 文本备注
    desnum?: number; // 金额备注
    truepull:any[];
    truenum: number;
    waitline:number;
}

interface CliLoad {
    address: string; // 用户地址
    loadtimes?: number; // 登录次数
    lasttime?: Date; // 最近登录
    loadip: string; // 登录 IP
    loadaddr: string; // 登录地址
    usdt?: number; // USD 余额
    status: string; // 当前状态
    desc: string; // 文本备注
    desnum?: number; // 金额备注
}
interface Mainorder {
    address: string;
    amount: number;
    desc: string;
    descnum: number;
}