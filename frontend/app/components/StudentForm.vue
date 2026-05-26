<template>
  <HeadlessTransitionRoot appear :show="show" as="template">
    <HeadlessDialog as="div" @close="$emit('close')" class="relative z-50">
      <HeadlessTransitionChild
        as="template"
        enter="duration-300 ease-out"
        enterFrom="opacity-0"
        enterTo="opacity-100"
        leave="duration-200 ease-in"
        leaveFrom="opacity-100"
        leaveTo="opacity-0"
      >
        <div class="fixed inset-0 bg-black/80 backdrop-blur-sm" />
      </HeadlessTransitionChild>

      <div class="fixed inset-0 overflow-y-auto">
        <div class="flex min-h-full items-center justify-center p-4 text-center">
          <HeadlessTransitionChild
            as="template"
            enter="duration-300 ease-out"
            enterFrom="opacity-0 scale-95"
            enterTo="opacity-100 scale-100"
            leave="duration-200 ease-in"
            leaveFrom="opacity-100 scale-100"
            leaveTo="opacity-0 scale-95"
          >
            <HeadlessDialogPanel class="w-full max-w-2xl transform bg-[color:var(--color-surface)] border border-[color:var(--color-border)] border-l-4 border-l-[color:var(--color-accent)] p-8 shadow-[--shadow-lg] text-left align-middle transition-all max-h-[90vh] overflow-y-auto">
              <div class="mb-6">
                <HeadlessDialogTitle as="h3" class="text-2xl font-bold text-[color:var(--color-heading)] tracking-tight">
                  Ubah Data Siswa
                </HeadlessDialogTitle>
                <p class="text-xs uppercase tracking-wider text-[color:var(--color-muted)] mt-1">
                  Lengkapi semua kolom formulir di bawah ini
                </p>
              </div>

              <div v-if="errorMessage" class="mb-4 bg-[color:var(--color-bg)] border border-[color:var(--color-error)] p-3 text-sm text-[color:var(--color-error)] font-medium flex items-center space-x-2">
                <AlertCircleIcon class="w-4 h-4 shrink-0" />
                <span>{{ errorMessage }}</span>
              </div>

              <form @submit.prevent="handleSubmit" class="space-y-4">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div class="space-y-1.5">
                    <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">NISN</label>
                    <input v-model="formData.nisn" type="text" class="input w-full font-mono text-[color:var(--color-accent)]" required placeholder="Contoh: 10001" />
                  </div>
                  <div class="space-y-1.5">
                    <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Nama Lengkap</label>
                    <input v-model="formData.name" type="text" class="input w-full" required placeholder="Contoh: Ahmad Fauzi" />
                  </div>
                </div>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div class="space-y-1.5">
                    <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Tempat Lahir</label>
                    <input v-model="formData.place_of_birth" type="text" class="input w-full" placeholder="Contoh: Jakarta" />
                  </div>
                  <div class="space-y-1.5">
                    <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Tanggal Lahir</label>
                    <input v-model="formData.date_of_birth" type="date" class="input w-full" />
                  </div>
                </div>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div class="space-y-1.5">
                    <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Kelas</label>
                    <input v-model="formData.class" type="text" class="input w-full" required placeholder="Contoh: X-A atau XI-B" />
                  </div>
                  <div class="space-y-1.5">
                    <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Gender</label>
                    <select v-model="formData.gender" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                      <option value="" disabled>Pilih Gender</option>
                      <option value="Laki-laki">Laki-laki</option>
                      <option value="Perempuan">Perempuan</option>
                    </select>
                  </div>
                </div>

                <div class="space-y-1.5">
                  <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Nomor Telepon</label>
                  <input v-model="formData.phone" type="text" class="input w-full font-mono" placeholder="Contoh: 081234567890" />
                </div>

                <div class="space-y-1.5">
                  <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Alamat Rumah</label>
                  <textarea v-model="formData.address" class="input w-full h-20 resize-none" placeholder="Alamat lengkap tempat tinggal siswa"></textarea>
                </div>

                <div class="flex items-center justify-end space-x-3 pt-4 border-t border-[color:var(--color-border)]">
                  <button type="button" @click="$emit('close')" class="button-ghost text-xs font-semibold uppercase tracking-wider">
                    Batal
                  </button>
                  <button type="submit" :disabled="isLoading" class="button-primary text-xs font-semibold uppercase tracking-wider flex items-center space-x-1">
                    <span v-if="isLoading" class="animate-spin rounded-full h-3 w-3 border-2 border-[color:var(--color-accent-fg)] border-t-transparent mr-1"></span>
                    <span>Simpan</span>
                  </button>
                </div>
              </form>
            </HeadlessDialogPanel>
          </HeadlessTransitionChild>
        </div>
      </div>
    </HeadlessDialog>
  </HeadlessTransitionRoot>
</template>

<script setup lang="ts">
import { ref, watch, computed } from "vue"
import { AlertCircle as AlertCircleIcon } from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

const props = defineProps<{
  show: boolean
  student: any | null
}>()

const emit = defineEmits(["close", "saved"])

const api = useApi()
const isLoading = ref(false)
const errorMessage = ref("")

const isEdit = computed(() => !!props.student)

const formData = ref({
  nisn: "",
  name: "",
  place_of_birth: "",
  date_of_birth: "",
  class: "",
  gender: "",
  address: "",
  phone: ""
})

watch(
  () => props.student,
  (newVal) => {
    if (newVal) {
      formData.value = {
        nisn: newVal.nisn || "",
        name: newVal.name || "",
        place_of_birth: newVal.place_of_birth || "",
        date_of_birth: newVal.date_of_birth || "",
        class: newVal.class || "",
        gender: newVal.gender || "",
        address: newVal.address || "",
        phone: newVal.phone || ""
      }
    } else {
      formData.value = {
        nisn: "",
        name: "",
        place_of_birth: "",
        date_of_birth: "",
        class: "",
        gender: "",
        address: "",
        phone: ""
      }
    }
    errorMessage.value = ""
  },
  { immediate: true }
)

async function handleSubmit() {
  if (!formData.value.nisn || !formData.value.name || !formData.value.class || !formData.value.gender) {
    errorMessage.value = "Kolom NISN, Nama, Kelas, dan Gender wajib diisi"
    return
  }

  isLoading.value = true
  errorMessage.value = ""

  try {
    let res: any
    if (isEdit.value && props.student) {
      res = await api.put(`/api/students/${props.student.id}`, formData.value)
    } else {
      res = await api.post("/api/students", formData.value)
    }

    if (res.success && res.data) {
      emit("saved", res.data)
      emit("close")
    } else {
      errorMessage.value = res.message || "Gagal menyimpan data siswa"
    }
  } catch (error: any) {
    if (error.response && error.response._data) {
      errorMessage.value = error.response._data.message || "Gagal memproses data"
    } else {
      errorMessage.value = "Koneksi ke server terputus"
    }
  } finally {
    isLoading.value = false
  }
}
</script>
