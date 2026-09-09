<script setup lang="ts">
import { ref } from 'vue'

// Prop para controlar a visibilidade do modal (caso queira controlar pelo componente pai)
const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits(['close', 'success'])

const valor = ref('')
const carregando = ref(false)

const fecharModal = () => {
  valor.value = ''
  emit('close')
}

const realizarDeposito = () => {
  if (!valor.value || Number(valor.value) <= 0) {
    alert('Por favor, insira um valor válido para o depósito.')
    return
  }

  carregando.value = (true)

  // Simulação de requisição/processamento
  setTimeout(() => {
    carregando.value = false
    alert(`Depósito de R$ ${valor.value} realizado com sucesso!`)
    emit('success', valor.value)
    fecharModal()
  }, 1000)
}
</script>

<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/80 backdrop-blur-sm transition-opacity">
    <div class="bg-gray-900 border border-gray-800 rounded-2xl p-6 md:p-8 max-w-md w-full shadow-2xl relative space-y-6">
      
      <!-- Cabeçalho do Modal -->
      <div class="flex items-center justify-between">
        <h3 class="text-xl font-bold text-white">Realizar Depósito</h3>
        <button 
          @click="fecharModal" 
          class="text-gray-400 hover:text-white transition-colors p-1 rounded-lg hover:bg-gray-800 cursor-pointer"
        >
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <!-- Formulário de Depósito -->
      <form @submit.prevent="realizarDeposito" class="space-y-4">
        <div class="space-y-2">
          <label class="text-xs font-semibold text-gray-400 uppercase tracking-wider block">
            Quanto você deseja depositar?
          </label>
          <div class="relative">
            <span class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-400 font-semibold text-lg">R$</span>
            <input 
              v-model="valor"
              type="number" 
              step="0.01"
              placeholder="0,00" 
              class="w-full bg-gray-950 border border-gray-800 rounded-xl py-3.5 pl-12 pr-4 text-white placeholder-gray-600 focus:outline-none focus:border-orange-500 font-bold text-lg transition-all"
              required
            />
          </div>
        </div>

        <p class="text-xs text-gray-500">
          O valor será creditado na sua conta após a validação.
        </p>

        <!-- Botões de Ação -->
        <div class="flex gap-3 pt-2">
          <button 
            type="button" 
            @click="fecharModal"
            class="flex-1 px-4 py-3 rounded-xl bg-gray-800 hover:bg-gray-750 text-gray-300 font-semibold text-sm border border-gray-700 transition-all cursor-pointer"
          >
            Cancelar
          </button>
          <button 
            type="submit" 
            :disabled="carregando"
            class="flex-1 px-4 py-3 rounded-xl bg-orange-600 hover:bg-orange-500 text-white font-semibold text-sm transition-all shadow-lg shadow-orange-950/50 cursor-pointer disabled:opacity-50"
          >
            {{ carregando ? 'Processando...' : 'Confirmar' }}
          </button>
        </div>
      </form>

    </div>
  </div>
</template>