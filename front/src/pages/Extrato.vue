<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useTransaction } from '../mods/transactions/composables/useTransactions' // Ajuste o caminho conforme a estrutura do seu projeto

const { getExtrato } = useTransaction()
const transacoes = ref<any[]>([])
const carregando = ref(false)

const carregarExtrato = async () => {
  carregando.value = true
  const response = await getExtrato()
  carregando.value = false

  if (response && response.transactions) {
    transacoes.value = response.transactions
  } else if (Array.isArray(response)) {
    transacoes.value = response
  }
}

onMounted(() => {
  carregarExtrato()
})

// Função para formatar o valor monetário em Real (R$)
const formatarValor = (valor: number) => {
  return new Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(valor)
}

// Função para formatar a data vinda do banco
const formatarData = (dataStr: string) => {
  if (!dataStr) return '-'
  const data = new Date(dataStr)
  return data.toLocaleDateString('pt-BR', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

// Estilização dinâmica baseada no status da transação
const getStatusClass = (status: string) => {
  switch (status?.toLowerCase()) {
    case 'pendente':
      return 'bg-yellow-900/50 text-yellow-300 border border-yellow-700/50'
    case 'concluido':
    case 'aprovado':
      return 'bg-green-900/50 text-green-300 border border-green-700/50'
    case 'cancelado':
    case 'rejeitado':
      return 'bg-red-900/50 text-red-300 border border-red-700/50'
    default:
      return 'bg-gray-800 text-gray-400'
  }
}
</script>

<template>
  <div class="bg-gray-900 border border-gray-800 rounded-2xl p-6 md:p-8 shadow-lg space-y-4">
    <div class="flex items-center justify-between">
      <h3 class="text-lg font-bold text-gray-200">Extrato Recente</h3>
      <span class="text-xs text-gray-400">Últimas movimentações</span>
    </div>

    <div class="overflow-x-auto">
      <table class="w-full text-left text-sm text-gray-300">
        <thead class="bg-gray-800/50 text-xs uppercase text-gray-400 border-b border-gray-800">
          <tr>
            <th class="py-3 px-4 rounded-l-xl">Descrição</th>
            <th class="py-3 px-4">Data</th>
            <th class="py-3 px-4">Valor</th>
            <th class="py-3 px-4 rounded-r-xl text-right">Status</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-800">
          
          <!-- Estado de carregamento -->
          <tr v-if="carregando">
            <td colspan="4" class="py-6 px-4 text-center text-gray-500">
              Carregando extrato...
            </td>
          </tr>

          <!-- Lista de transações retornadas do banco -->
          <template v-else-if="transacoes.length > 0">
            <tr v-for="tx in transacoes" :key="tx.ID || tx.id" class="hover:bg-gray-800/30 transition-colors">
              <td class="py-4 px-4 font-medium text-white capitalize">
                {{ tx.Type || tx.type || 'Transação' }}
              </td>
              <td class="py-4 px-4 text-gray-400">
                {{ formatarData(tx.CreatedAt || tx.created_at) }}
              </td>
              <td class="py-4 px-4 font-semibold text-white">
                {{ formatarValor(tx.Amount || tx.amount) }}
              </td>
              <td class="py-4 px-4 text-right">
                <span :class="['px-2.5 py-1 text-xs font-semibold rounded-lg capitalize', getStatusClass(tx.Status || tx.status)]">
                  {{ tx.Status || tx.status || 'Pendente' }}
                </span>
              </td>
            </tr>
          </template>

          <!-- Estado vazio (caso não haja nenhuma transação) -->
          <tr v-else class="hover:bg-gray-800/30 transition-colors">
            <td class="py-4 px-4 font-medium text-white">Nenhuma movimentação recente</td>
            <td class="py-4 px-4 text-gray-500">-</td>
            <td class="py-4 px-4 text-gray-500">-</td>
            <td class="py-4 px-4 text-right">
              <span class="px-2.5 py-1 text-xs font-semibold bg-gray-800 text-gray-400 rounded-lg">Vazio</span>
            </td>
          </tr>

        </tbody>
      </table>
    </div>
  </div>
</template>