import { ref } from "vue";
import { getDataBus, editDataBus, postDataBus, deleteDataBus, detailDataBus } from "@/services/auth";

export function useBus() {
    const bus = ref([])
    const search = ref('')
    const statusBus = ref('')
    const classBus = ref('')

    const fetchDataBus = async () => {
        const res = await getDataBus({ 
            search: search.value,
            status: statusBus.value,
            bus_class: classBus.value
        })
        bus.value = res.data.data
    }

    const createBus = (data) => postDataBus(data)
    const updateBus = (id, data) => editDataBus(id, data)
    const removeBus = (id) => deleteDataBus(id)
    const detailBus = (id) => detailDataBus(id)

    return {
        bus,
        search,
        statusBus,
        classBus,
        fetchDataBus,
        createBus,
        updateBus,
        removeBus,
        detailBus,
    }
}