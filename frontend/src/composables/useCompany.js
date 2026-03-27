import { ref } from "vue";
import { getDataCompany, editDataCompany, postDataCompany, detailDataCompany, deleteDataCompany } from "@/services/auth";

export function useCompany() {
    const company = ref([])
    const search = ref('')

    const fetchCompany = async () => {
        const res = await getDataCompany({ search: search.value })
        company.value = res.data.data
    }

    const createCompany = (data) => postDataCompany(data)
    const updateCompany = (id, data) => editDataCompany(id, data)
    const removeCompany = (id) => deleteDataCompany(id)
    const detailCompany = (id) => detailDataCompany(id)

    return {
        company,
        search,
        fetchCompany,
        createCompany,
        updateCompany,
        removeCompany,
        detailCompany,
    }
}