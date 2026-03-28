import { ref } from "vue";
import { getDataBus, editDataBus, postDataBus, deleteDataBus, detailDataBus } from "@/services/auth";

export function useBus() {
    const bus = ref([])
    const search = ref('')

    const fetchDataBus = async () => {
        const res = await getDataBus({ search: search.value })
        bus.value = res.data.data
    }

    const createBus = (data) => postDataBus(data)
    const updateBus = (id, data) => editDataBus(id, data)
    const removeBus = (id) => deleteDataBus(id)
    const detailBus = (id) => detailDataBus(id)

    return {
        bus,
        search,
        fetchDataBus,
        createBus,
        updateBus,
        removeBus,
        detailBus,
    }
}