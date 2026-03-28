<script setup>
import Modal from '../ui/Modal.vue';
import { ref, watch, computed, defineProps, defineEmits } from 'vue';
import { 
    Listbox,
    ListboxButton,
    ListboxOptions,
    ListboxOption, 
} from '@headlessui/vue';

const props = defineProps({
    modelValue: Boolean,
    data: {
        type: Object,
        default: () => ({
            id: null,
            origin: '',
            destination: '',
            distance: ''
        })
    },
    mode: { type: String, default: 'create' },
    companies: { type: Array, default: () => [] }
})

const emits = defineEmits(['update:modelValue', 'submit'])
const showModal = computed({
    get: () => props.modelValue,
    set: (val) => emits('update:modelValue', val)
})


const form = ref({ ...props.data })

watch(() => props.data, (val) => {
    form.value = { ...val }
}, { immediate: true })

const handleSubmit = () => {
    emits('submit', { ...form.value })
    showModal.value = false
}
</script>