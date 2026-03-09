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