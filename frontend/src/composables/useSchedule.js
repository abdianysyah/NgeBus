import { ref } from "vue";
import { getDataSchedule, editDataSchedule, postDataSchedule, deleteDataSchedule } from "@/services/auth";

export function useSchedule() {
    const schedule = ref([])
    const fetchDataSchedule = async () => {
        try {
            const res = await getDataSchedule()
            schedule.value = Array.isArray(res.data.data) ? res.data.data : []
        } catch (error) {
            schedule.value = []
            console.error(error)
        }
    }

    const createSchedule = (data) => postDataSchedule(data)
    const updateSchedule = (id, data) => editDataSchedule(id, data)
    const removeSchedule = (id) => deleteDataSchedule(id)

    return {
        schedule,
        fetchDataSchedule,
        createSchedule,
        updateSchedule,
        removeSchedule,
    }
}