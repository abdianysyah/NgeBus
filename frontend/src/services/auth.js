import api from "./api";

export const register = (data) => {
    return api.post("/register", data)
}

export const login = (data) => {
    return api.post("/login", data)
}

export const countData = (data) => {
    return api.get("/admin/dashboard", {
        params: data
    })
}

export const getDataBus = (params) => {
    return api.get("/bus", {
        params: params
    })
}

export const addBus = (data) => {
    return api.post("/bus", data)
}

export const editBus = (id, data) => {
    return api.put(`/bus/${id}`, data)
}

export const deleteBus = (id) => {
    return api.delete(`/bus/${id}`)
}

export const getDataRoute = (data) => {
    return api.get("/route", data)
}

export const postDataRoute = (data) => {
    return api.post("/route", data)
}

export const deleteDataRoute = (id) => {
    return api.delete(`/route/${id}`)
} 

export const editDataRoute = (id, data) => {
    return api.put(`/route/${id}`, data)
}

export const getDataCompany = (params) => {
    return api.get("/company", { params })
}

export const postDataCompany = (data) => {
    return api.post("/company", data)
}

export const editDataCompany = (id, data) => {
    return api.put(`/company/${id}`, data)
}

export const deleteDataCompany = (id) => {
    return api.delete(`/company/${id}`)
}

export const detailDataCompany = (id) => {
    return api.get(`/company/${id}`)
}
