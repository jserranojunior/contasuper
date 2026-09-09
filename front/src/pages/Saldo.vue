<template>

      <div class="bg-gray-900 border border-gray-800 p-6 md:p-8 rounded-2xl shadow-lg relative overflow-hidden">
        <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
          <div>
            <div class="flex items-center gap-2 mb-2">
              <span class="text-xs font-semibold text-gray-400 uppercase tracking-wider">Saldo Disponível</span>
              <!-- Botão do Olho -->
              <button 
                @click="showBalance = !showBalance" 
                class="text-gray-400 hover:text-white transition-colors focus:outline-none cursor-pointer p-1"
                :aria-label="showBalance ? 'Ocultar saldo' : 'Mostrar saldo'"
              >
                <!-- Olho Aberto -->
                <svg v-if="showBalance" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
                <!-- Olho Fechado -->
                <svg v-else class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.98 8.223A10.477 10.477 0 001.934 12C3.322 16.173 7.258 19.18 11.9 19.18c1.71 0 3.324-.366 4.767-1.02m2.418-1.574C20.578 15.13 21.652 13.67 22.066 12c-1.388-4.173-5.324-7.18-9.966-7.18-1.42 0-2.775.26-4.02.732M3 3l18 18" />
                </svg>
              </button>
            </div>
            
            <div class="text-3xl md:text-4xl font-extrabold text-white tracking-tight">
              <span v-if="showBalance">R$ 0,00</span>
              <span v-else class="tracking-widest text-gray-600 select-none">••••••••</span>
            </div>
          </div>

          <div class="text-left sm:text-right">
            <span class="text-xs font-medium text-orange-400">Conta Ativa e Verificada</span>
          </div>
        </div>
      </div>
</template>

<script setup lang="ts">
import {ref} from 'vue'
// Estado da visibilidade do saldo
const showBalance = ref(false)

</script>