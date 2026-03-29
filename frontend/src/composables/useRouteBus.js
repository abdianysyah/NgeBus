import { ref } from "vue";
import { getDataRoute, postDataRoute, editDataRoute, deleteDataRoute } from "@/services/auth";


export function useRouteBus() {
    const rute = ref([])
    const search = ref('')

    const fetchDataRoute = async () => {
        const res = await getDataRoute({
            search: search.value
        })
        rute.value = res.data.data
    }

    const createRoute = (data) => postDataRoute(data)
    const updateRoute = (id, data) => editDataRoute(id, data)
    const removeRoute = (id) => deleteDataRoute(id)

    return {
        rute,
        search,
        fetchDataRoute,
        createRoute,
        updateRoute,
        removeRoute
    }
}