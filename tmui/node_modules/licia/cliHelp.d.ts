declare namespace cliHelp {
    interface IOption {
        name: string;
        shorthand?: string;
        desc: string;
    }
    interface ICommand {
        name: string;
        desc: string;
        usage: string | string[];
        options?: IOption[];
    }
    interface IData {
        name: string;
        usage: string | string[];
        commands: ICommand[];
    }
}
declare function cliHelp(data: cliHelp.IData | cliHelp.ICommand): string;

export = cliHelp;
