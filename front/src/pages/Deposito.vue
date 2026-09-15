<script setup lang="ts">
import { ref } from 'vue'
import { useTransaction } from '../mods/transactions/composables/useTransactions'

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits(['close', 'success'])

const valor = ref('')
const carregando = ref(false)
const errorMessage = ref('')

const { createDeposit } = useTransaction()

/**
 * Formata o valor digitado para moeda brasileira.
 *
 * Exemplos:
 * 5      -> R$ 5,00
 * 50     -> R$ 50,00
 * 500    -> R$ 500,00
 * 1234   -> R$ 1.234,00
 */
const formatarMoeda = (event: Event) => {
  const input = event.target as HTMLInputElement

  // Mantém somente números
  const apenasNumeros = input.value.replace(/\D/g, '')

  if (!apenasNumeros) {
    valor.value = ''
    return
  }

  // Converte para centavos
  const numero = Number(apenasNumeros) / 100

  valor.value = numero.toLocaleString('pt-BR', {
    style: 'currency',
    currency: 'BRL',
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  })
}

/**
 * Converte:
 *
 * R$ 50,00      -> 50
 * R$ 1.234,56   -> 1234.56
 * R$ 10.000,00  -> 10000
 */
const obterValorNumerico = (): number => {
  if (!valor.value) {
    return 0
  }

  return Number(
    valor.value
      .replace('R$', '')
      .replace(/\s/g, '')
      .replace(/\./g, '')
      .replace(',', '.')
  )
}

const fecharModal = () => {
  valor.value = ''
  errorMessage.value = ''

  emit('close')
}

const realizarDeposito = async () => {
  const numericValue = obterValorNumerico()

  console.log('Valor formatado:', valor.value)
  console.log('Valor enviado para API:', numericValue)

  if (!Number.isFinite(numericValue) || numericValue <= 0) {
    errorMessage.value =
      'Por favor, insira um valor válido para o depósito.'
    return
  }

  carregando.value = true
  errorMessage.value = ''

  try {
    const result = await createDeposit({
      amount: numericValue
    })

    if (result && result.error) {
      errorMessage.value = result.error
      return
    }

    console.log('Depósito efetuado com sucesso:', numericValue)

    alert(
      `Solicitação de depósito de R$ ${numericValue
        .toFixed(2)
        .replace('.', ',')} realizada com sucesso! Dirija-se ao caixa para concluir.`
    )

    emit('success', numericValue)

    fecharModal()
  } catch (error) {
    console.error('Erro ao realizar depósito:', error)

    errorMessage.value =
      'Ocorreu um erro ao realizar o depósito. Tente novamente.'
  } finally {
    carregando.value = false
  }
}
</script>

<template>
  <div
    v-if="isOpen"
    class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/80 backdrop-blur-sm transition-opacity"
  >
    <div
      class="bg-gray-900 border border-gray-800 rounded-2xl p-6 md:p-8 max-w-md w-full shadow-2xl relative space-y-6"
    >
      <!-- Cabeçalho -->
      <div class="flex items-center justify-between">
        <h3 class="text-xl font-bold text-white">
          Realizar Depósito
        </h3>

        <button
          type="button"
          @click="fecharModal"
          class="text-gray-400 hover:text-white transition-colors p-1 rounded-lg hover:bg-gray-800 cursor-pointer"
        >
          <svg
            class="w-5 h-5"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </button>
      </div>

      <!-- Erro -->
      <div
        v-if="errorMessage"
        class="text-sm text-yellow-200 text-center bg-yellow-900/80 border border-yellow-700 p-3 rounded-xl"
        role="alert"
      >
        {{ errorMessage }}
      </div>

      <!-- Formulário -->
      <form
        @submit.prevent="realizarDeposito"
        class="space-y-4"
      >
        <div class="space-y-2">
          <label
            class="text-xs font-semibold text-gray-400 uppercase tracking-wider block"
          >
            Quanto você deseja depositar?
          </label>

          <input
            :value="valor"
            @input="formatarMoeda"
            type="text"
            inputmode="numeric"
            autocomplete="off"
            placeholder="R$ 0,00"
            required
            class="w-full bg-gray-950 border border-gray-800 rounded-xl py-3.5 px-4 text-white placeholder-gray-600 focus:outline-none focus:border-orange-500 font-bold text-lg transition-all"
          />
        </div>

        <p class="text-xs text-gray-500">
          Informe o valor desejado e dirija-se a um caixa para
          entregar o valor e confirmar a operação.
        </p>

        <!-- Botões -->
        <div class="flex gap-3 pt-2">
          <button
            type="button"
            @click="fecharModal"
            :disabled="carregando"
            class="flex-1 px-4 py-3 rounded-xl bg-gray-800 hover:bg-gray-750 text-gray-300 font-semibold text-sm border border-gray-700 transition-all cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed"
          >
            Cancelar
          </button>

          <button
            type="submit"
            :disabled="carregando"
            class="flex-1 px-4 py-3 rounded-xl bg-orange-600 hover:bg-orange-500 text-white font-semibold text-sm transition-all shadow-lg shadow-orange-950/50 cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {{ carregando ? 'Processando...' : 'Confirmar' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>
