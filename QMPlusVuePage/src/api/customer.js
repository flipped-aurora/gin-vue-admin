import service from '@/utils/request'


export const getCustomerList = (data) => {
    return service({
        url: "/customer/list",
        method: 'post',
        data: data
    })
}

export const addCustomer = (data) => {
    return service({
        url: "/customer/add",
        method: 'post',
        data: data
    })
}

export const updateCustomer = (data) => {
    return service({
        url: "/customer/update",
        method: 'patch',
        data: data
    })
}

export const delCustomer = (data) => {
    return service({
        url: `/customer/delete`,
        method: 'delete',
        data: data
    })
}

export const getCustomerAddressList = (data) => {
    return service({
        url: "/customer/address",
        method: 'post',
        data
    })
}

export const addCustomerAddress = (data) => {
    return service({
        url: "/customerAddress/add",
        method: 'post',
        data
    })
}

export const updateCustomerAddress = (data) => {
    return service({
        url: "/customerAddress/update",
        method: 'patch',
        data
    })
}

export const delCustomerAddress = (data) => {
    return service({
        url: `/customerAddress/delete`,
        method: 'delete',
        data: data
    })
}

export const getCustomerById = (data) => {
    return service({
        url: `/customer/getbyid`,
        method: 'post',
        data: data
    })
}

export const getAddressById = (data) => {
    return service({
        url: "/customer/getbyid",
        method: 'post',
        data: data
    })
}