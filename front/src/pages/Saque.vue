<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/70 backdrop-blur-sm animate-fadeIn">
    
    <!-- Caixa do Modal -->
    <div class="bg-gray-900 border border-gray-800 text-white rounded-2xl w-full max-w-md p-6 md:p-8 shadow-2xl relative overflow-hidden">
      
      <!-- Detalhe decorativo de luz -->
      <div class="absolute -top-24 -right-24 w-48 h-48 bg-orange-500/10 rounded-full blur-3xl pointer-events-none"></div>

      <!-- Cabeçalho -->
      <div class="flex items-center justify-between mb-6">
        <div>
          <span class="text-xs font-semibold text-orange-400 uppercase tracking-wider">Retirada Física</span>
          <h3 class="text-xl md:text-2xl font-bold text-white mt-0.5">Saque no Caixa</h3>
        </div>
        <button 
          @click="$emit('close')" 
          class="text-gray-400 hover:text-white bg-gray-800 hover:bg-gray-700 p-2 rounded-xl transition-all cursor-pointer"
        >
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <!-- Alerta de Erro / Feedback se houver -->
      <div 
        v-if="errorMessage" 
        class="mb-4 text-sm text-yellow-200 text-center bg-yellow-900/80 border border-yellow-700 p-3 rounded-xl"
        role="alert"
      >
        {{ errorMessage }}
      </div>

      <!-- Aviso sobre ir ao local físico -->
      <div class="mb-5 p-3 rounded-xl bg-orange-500/10 border border-orange-500/20 text-xs text-orange-300 space-y-1">
        <p class="font-semibold flex items-center gap-1.5">
          <svg class="w-4 h-4 shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          Instruções para a retirada:
        </p>
        <p class="text-gray-300">Informe o valor desejado e dirija-se a um caixa para receber o valor.</p>
      </div>

      <!-- Formulário de Saque -->
      <form @submit.prevent="handleSaque" class="space-y-4">
        
        <div>
          <label for="valorSaque" class="block text-xs font-medium text-gray-300 mb-1">Valor do Saque (R$)</label>
          <input 
            v-model="valor" 
            type="text" 
            id="valorSaque" 
            placeholder="0,00"
            class="w-full px-4 py-3 rounded-xl bg-gray-800 border border-gray-700 text-sm text-gray-100 focus:bg-gray-900 focus:outline-none focus:border-orange-500 focus:ring-1 focus:ring-orange-500 transition-all"
          >
        </div>

        <!-- Botões de Ação -->
        <div class="pt-4 flex items-center gap-3">
          <button 
            type="button" 
            @click="$emit('close')" 
            class="w-1/2 py-3 px-4 rounded-xl bg-gray-800 hover:bg-gray-700 text-gray-300 font-semibold text-sm transition-all border border-gray-700 text-center cursor-pointer"
          >
            Cancelar
          </button>
          
          <button 
            type="submit" 
            class="w-1/2 py-3 px-4 rounded-xl bg-gradient-to-r from-orange-500 to-orange-600 hover:from-orange-600 hover:to-orange-700 text-white font-semibold text-sm shadow-lg shadow-orange-500/30 transition-all cursor-pointer active:scale-[0.99]"
          >
            Confirmar Saque
          </button>
        </div>

      </form>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

defineProps({
  isOpen: {
    type: Boolean,
    required: true
  }
})

const emit = defineEmits(['close', 'success'])

const valor = ref('')
const errorMessage = ref('')

function handleSaque() {
  if (!valor.value) {
    errorMessage.value = 'Por favor, informe o valor do saque.'
    return
  }
  
  errorMessage.value = ''
  emit('success', valor.value)
  emit('close')
}
</script>