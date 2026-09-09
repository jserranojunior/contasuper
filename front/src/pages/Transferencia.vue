<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/70 backdrop-blur-sm animate-fadeIn">
    
    <!-- Caixa do Modal -->
    <div class="bg-gray-900 border border-gray-800 text-white rounded-2xl w-full max-w-lg p-6 md:p-8 shadow-2xl relative overflow-hidden max-h-[90vh] overflow-y-auto">
      
      <!-- Detalhe decorativo de luz -->
      <div class="absolute -top-24 -right-24 w-48 h-48 bg-orange-500/10 rounded-full blur-3xl pointer-events-none"></div>

      <!-- Cabeçalho -->
      <div class="flex items-center justify-between mb-6">
        <div>
          <span class="text-xs font-semibold text-orange-400 uppercase tracking-wider">Serviços Financeiros</span>
          <h3 class="text-xl md:text-2xl font-bold text-white mt-0.5">Transferência entre Contas</h3>
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

      <!-- Formulário -->
      <form @submit.prevent="handleTransferencia" class="space-y-4">
        
        <div>
          <label for="contaDestino" class="block text-xs font-medium text-gray-300 mb-1">Número da Conta ou ID de Destino</label>
          <input 
            v-model="destinoConta"
            type="text" 
            id="contaDestino" 
            placeholder="Ex: 12345-6"
            class="w-full px-4 py-3 rounded-xl bg-gray-800 border border-gray-700 text-sm text-gray-100 focus:bg-gray-900 focus:outline-none focus:border-orange-500 focus:ring-1 focus:ring-orange-500 transition-all"
          >
        </div>

        <div>
          <label for="valorTransferencia" class="block text-xs font-medium text-gray-300 mb-1">Valor da Transferência (R$)</label>
          <input 
            v-model="valor" 
            type="text" 
            id="valorTransferencia" 
            placeholder="0,00"
            class="w-full px-4 py-3 rounded-xl bg-gray-800 border border-gray-700 text-sm text-gray-100 focus:bg-gray-900 focus:outline-none focus:border-orange-500 focus:ring-1 focus:ring-orange-500 transition-all"
          >
        </div>

        <div>
          <label for="descricao" class="block text-xs font-medium text-gray-300 mb-1">Descrição (Opcional)</label>
          <input 
            v-model="descricao" 
            type="text" 
            id="descricao" 
            placeholder="Ex: Pagamento, Aluguel..."
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
            Confirmar Envio
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

const destinoConta = ref('')
const valor = ref('')
const descricao = ref('')
const errorMessage = ref('')

function handleTransferencia() {
  if (!destinoConta.value) {
    errorMessage.value = 'Por favor, informe a conta de destino.'
    return
  }
  if (!valor.value) {
    errorMessage.value = 'Por favor, informe o valor.'
    return
  }
  
  errorMessage.value = ''
  emit('success', { valor: valor.value, destino: destinoConta.value })
  emit('close')
}
</script>