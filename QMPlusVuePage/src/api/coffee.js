import service from '@/utils/request'


export const getCoffeeList = (data) => {
    return service({
        url: "/base/coffee",
        method: 'post',
        data: data
    })
}

export const addCoffee = (data) => {
    return service({
        url: "/coffee/add",
        method: 'post',
        data: data
    })
}

export const updateCoffee = (data) => {
    return service({
        url: "/coffee/update",
        method: 'patch',
        data: data
    })
}

export const delCoffee = (data) => {
    return service({
        url: `/coffee/delete`,
        method: 'delete',
        data: data
    })
}

export const getCoffeeTypeList = (data) => {
    return service({
        url: "/base/coffeetype",
        method: 'post',
        data
    })
}

export const addCoffeeType = (data) => {
    return service({
        url: "/coffeetype/add",
        method: 'post',
        data
    })
}

export const updateCoffeeType = (data) => {
    return service({
        url: "/coffeetype/update",
        method: 'patch',
        data
    })
}

export const delCoffeeType = (data) => {
    return service({
        url: `/coffeetype/delete`,
        method: 'delete',
        data: data
    })
}

export const changeCoffeeType = (data) => {
    return service({
        url: "/coffee/type",
        method: 'post',
        data: data
    })
}

export const getCoffeeById = (data) => {
    return service({
        url: "/coffee/getbyid",
        method: 'post',
        data: data
    })
}

export const getCoffeeTypeById = (data) => {
    return service({
        url: "/coffeetype/getbycode",
        method: 'post',
        data: data
    })
}

export const getCoffeeSpecByCoffeeId = (data) => {
    return service({
        url: "/coffeespec/getbycoffeeid",
        method: 'post',
        data: data
    })
}

export const getCoffeeSpecDetail = (data) => {
    return service({
        url: "/coffeespec/detail",
        method: 'post',
        data: data
    })
}
