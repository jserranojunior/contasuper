<template>
  <div class="min-h-screen bg-gray-950 text-gray-100 flex flex-col">
    <!-- Conteúdo Principal -->
    <main class="flex-1 p-6 md:p-8 max-w-7xl w-full mx-auto space-y-6">

      <!-- Saudação / Banner Superior -->
      <div class="bg-gradient-to-r from-gray-900 via-gray-900 to-orange-950/30 border border-gray-800 rounded-2xl p-6 md:p-8 shadow-xl">
        <h2 class="text-2xl md:text-3xl font-bold text-white mb-2">
          Painel Financeiro
        </h2>
        <p class="text-sm text-gray-400">
          Acompanhe o seu saldo, realize transferências, depósitos ou saques com rapidez e segurança.
        </p>
      </div>
      
      <!-- Card de Saldo com Visibilidade (Olho) -->
       <Saldo></Saldo>

      <!-- Seção de Ações Rápidas (Depósito, Transferência e Saque) -->
      <div class="bg-gray-900 border border-gray-800 rounded-2xl p-6 md:p-8 shadow-lg space-y-4">
        <h3 class="text-lg font-bold text-gray-200">Ações Financeiras</h3>
        
        <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
          
          <!-- Botão Depósito (Abre o Modal de Depósito) -->
          <button 
            @click="isDepositoOpen = true" 
            class="p-5 rounded-xl bg-gray-800 hover:bg-gray-750 border border-gray-700 hover:border-orange-500/60 text-left transition-all duration-200 group cursor-pointer hover:scale-[1.02] active:scale-95"
          >
            <div class="flex items-center justify-between mb-3">
              <div class="w-10 h-10 rounded-lg bg-orange-500/10 flex items-center justify-center text-orange-400 group-hover:bg-orange-500 group-hover:text-white transition-all">
                <svg class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                </svg>
              </div>
              <span class="text-xs uppercase tracking-wider font-medium text-gray-400">Adicionar</span>
            </div>
            <h4 class="font-bold text-white text-lg mb-0.5 group-hover:text-orange-400 transition-colors">Depósito</h4>
            <p class="text-xs text-gray-400">Adicione fundos à sua conta rapidamente.</p>
          </button>

          <!-- Botão Transferência (Abre o Modal de Transferência) -->
          <button 
            @click="isTransferenciaOpen = true" 
            class="p-5 rounded-xl bg-gray-800 hover:bg-gray-750 border border-gray-700 hover:border-orange-500/60 text-left transition-all duration-200 group cursor-pointer hover:scale-[1.02] active:scale-95"
          >
            <div class="flex items-center justify-between mb-3">
              <div class="w-10 h-10 rounded-lg bg-orange-500/10 flex items-center justify-center text-orange-400 group-hover:bg-orange-500 group-hover:text-white transition-all">
                <svg class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4" />
                </svg>
              </div>
              <span class="text-xs uppercase tracking-wider font-medium text-gray-400">Enviar</span>
            </div>
            <h4 class="font-bold text-white text-lg mb-0.5 group-hover:text-orange-400 transition-colors">Transferência</h4>
            <p class="text-xs text-gray-400">Envie valores para outras contas ou bancos.</p>
          </button>

          <!-- Botão Saque (Abre o Modal de Saque) -->
          <button 
            @click="isSaqueOpen = true" 
            class="p-5 rounded-xl bg-gray-800 hover:bg-gray-750 border border-gray-700 hover:border-orange-500/60 text-left transition-all duration-200 group cursor-pointer hover:scale-[1.02] active:scale-95"
          >
            <div class="flex items-center justify-between mb-3">
              <div class="w-10 h-10 rounded-lg bg-orange-500/10 flex items-center justify-center text-orange-400 group-hover:bg-orange-500 group-hover:text-white transition-all">
                <svg class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" />
                </svg>
              </div>
              <span class="text-xs uppercase tracking-wider font-medium text-gray-400">Retirar</span>
            </div>
            <h4 class="font-bold text-white text-lg mb-0.5 group-hover:text-orange-400 transition-colors">Saque</h4>
            <p class="text-xs text-gray-400">Realize saques do seu saldo disponível.</p>
          </button>

        </div>
      </div>

      <!-- Tabela Recente de Extrato / Transações -->
  <Extrato></Extrato>

    </main>

    <!-- Modais Integrados -->
    <Deposito 
      :is-open="isDepositoOpen" 
      @close="isDepositoOpen = false" 
      @success="handleDepositoSuccess" 
    />

    <Transferencia 
      :is-open="isTransferenciaOpen" 
      @close="isTransferenciaOpen = false" 
      @success="handleTransferenciaSuccess" 
    />

    <Saque 
      :is-open="isSaqueOpen" 
      @close="isSaqueOpen = false" 
      @success="handleSaqueSuccess" 
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import useStore from "../helpers/stores/store"
import Deposito from './Deposito.vue'
import Transferencia from './Transferencia.vue'
import Saque from './Saque.vue'
import Saldo from './Saldo.vue'
import Extrato from './Extrato.vue'


const { auth, router } = useStore()

// Estados dos Modais
const isDepositoOpen = ref(false)
const isTransferenciaOpen = ref(false)
const isSaqueOpen = ref(false)



onMounted(async () => {
  if (auth && typeof auth.isLogged === 'function') {
    const logged = await auth.isLogged()
    if (!logged) {
      router.push({ path: '/login' })
    }
  }
})

function handleDepositoSuccess(valor: string) {
  console.log('Depósito efetuado com sucesso:', valor)
}

function handleTransferenciaSuccess(dados: any) {
  console.log('Transferência efetuada com sucesso:', dados)
}

function handleSaqueSuccess(valor: string) {
  console.log('Saque efetuado com sucesso:', valor)
}
</script>