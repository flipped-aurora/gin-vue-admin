import { PluginFunc, ConfigType, QUnitType } from 'dayjs/esm'

declare const plugin: PluginFunc
export = plugin

declare module 'dayjs/esm' {
  interface Dayjs {
    quarter(): number

    quarter(quarter: number): Dayjs

    add(value: number, unit: QUnitType): Dayjs

    subtract(value: number, unit: QUnitType): Dayjs

    startOf(unit: QUnitType): Dayjs

    endOf(unit: QUnitType): Dayjs

    isSame(date: ConfigType, unit?: QUnitType): boolean

    isBefore(date: ConfigType, unit?: QUnitType): boolean

    isAfter(date: ConfigType, unit?: QUnitType): boolean
  }
}
