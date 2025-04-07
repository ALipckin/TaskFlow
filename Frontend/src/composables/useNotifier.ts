import { ref } from 'vue'

const snackbar = ref(false)
const text = ref('')
const color = ref<'success' | 'info' | 'warning' | 'error'>('info')

export function useNotifier() {
  function notify(message: string, type: typeof color.value = 'info') {
    text.value = message
    color.value = type
    snackbar.value = true
  }

  return {
    snackbar,
    text,
    color,
    notify
  }
}
