import { ref } from "vue";
import { getDataTicket, postDataTicket, editDataTicket, deleteDataTicket } from "@/services/auth";

export function useTicket() {
    const ticket = ref([])
    const fetchDataTicket = async (params) => {
        const res = await getDataTicket(params)
        ticket.value = res.data.data
    }

    const createTicket = (data) => postDataTicket(data)
    const updateTicket = (id, data) => editDataTicket(id, data)
    const removeTicket = (id) => deleteDataTicket(id)

    return {
        ticket,
        fetchDataTicket,
        createTicket,
        updateTicket,
        removeTicket,
    }
}