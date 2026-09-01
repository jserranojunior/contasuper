<script setup lang="ts">
import { reactive, ref } from 'vue'
import emailjs from '@emailjs/browser'
import fundoContato from '../assets/img/fundo-card.jpg'

const form = reactive({
  nome: '',
  cpf: '',
  email: '',
  telefone: ''
})

const phone = '5511984526015'
const loading = ref(false)

const handleSubmit = async () => {
  loading.value = true

  try {
    // 1. Envio para o E-mail via EmailJS
    // Certifique-se de que as chaves do objeto abaixo correspondem às variáveis do seu template no EmailJS
   const templateParams = {
      nome: form.nome,
      cpf: form.cpf,
      email: form.email,
      telefone: form.telefone,
    }

    const SERVICE_ID = 'service_e5e9ut8'
    const TEMPLATE_ID = 'template_db0lps6'
    const PUBLIC_KEY = 'r5i_Z-wfSplLs325N'

    await emailjs.send(SERVICE_ID, TEMPLATE_ID, templateParams, PUBLIC_KEY)

    // 2. Montagem e Envio para o WhatsApp (código que você já tinha)
    const message = `Olá, gostaria de solicitar o minha Conta Super!

*Dados da solicitação:*
• *Nome:* ${form.nome}
• *CPF:* ${form.cpf}
• *E-mail:* ${form.email}
• *Telefone:* ${form.telefone}`

    const encodedMessage = encodeURIComponent(message)
    const isMobile = /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)
    
    const whatsappUrl = isMobile
      ? `https://api.whatsapp.com/send?phone=${phone}&text=${encodedMessage}`
      : `https://web.whatsapp.com/send?phone=${phone}&text=${encodedMessage}`

    window.open(whatsappUrl, '_blank')

    // Limpa o formulário
    form.nome = ''
    form.cpf = ''
    form.email = ''
    form.telefone = ''

  } catch (error) {
    console.error('Erro ao enviar e-mail:', error)
    alert('Ocorreu um erro ao processar sua solicitação por e-mail. Tente novamente.')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <section 
    class="bg-brand-dark bg-cover bg-center bg-no-repeat text-white p-12 px-6 md:px-14 rounded-[32px] max-w-6xl mx-auto my-4 shadow-2xl overflow-hidden relative"
    :style="{ backgroundImage: `url(${fundoContato})` }"
  >
    <div class="absolute inset-0 bg-black/20 rounded-[32px] pointer-events-none"></div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-center relative z-10">
      
      <div class="lg:col-span-6 space-y-6">
        <div>
          <h2 class="text-3xl md:text-5xl font-extrabold leading-[1.15] tracking-tight">
            Conta Super <span class="text-white"><sup class="text-xs md:text-sm font-normal">®</sup></span> <br />
            <span class="font-normal text-2xl md:text-3xl text-orange-100/90 block mt-1">Carteira Digital</span>
          </h2>
        </div>
      </div>

      <div class="lg:col-span-6 relative">
        <div class="bg-gray-900 text-white p-6 md:p-8 rounded-2xl shadow-2xl max-w-md mx-auto">
          
          <h3 class="text-xl md:text-2xl font-bold text-gray-100 mb-6 text-center">
            Seja Super<sup class="text-xs font-normal">®</sup>
          </h3>

          <form @submit.prevent="handleSubmit" class="space-y-4">
            
            <div>
              <label class="block text-xs font-medium text-gray-100 mb-1">Nome</label>
              <input 
                v-model="form.nome"
                type="text" 
                placeholder="Seu nome completo" 
                class="w-full px-4 py-3 rounded-xl border border-gray-500 text-sm text-gray-100 focus:bg-gray-800 focus:outline-none focus:border-brand-orange focus:ring-1 focus:ring-brand-orange transition-all"
                required
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-100 mb-1">CPF</label>
              <input 
                v-model="form.cpf"
                type="text" 
                placeholder="000.000.000-00" 
                class="w-full px-4 py-3 rounded-xl border border-gray-500 text-sm text-gray-100 focus:bg-gray-800 focus:outline-none focus:border-brand-orange focus:ring-1 focus:ring-brand-orange transition-all"
                required
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-100 mb-1">E-mail</label>
              <input 
                v-model="form.email"
                type="email" 
                placeholder="seu.email@exemplo.com" 
                class="w-full px-4 py-3 rounded-xl border border-gray-500 text-sm text-gray-100 focus:bg-gray-800 focus:outline-none focus:border-brand-orange focus:ring-1 focus:ring-brand-orange transition-all"
                required
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-100 mb-1">Telefone</label>
              <input 
                v-model="form.telefone"
                type="tel" 
                placeholder="(00) 00000-0000" 
                class="w-full px-4 py-3 rounded-xl border border-gray-500 text-sm text-gray-100 focus:bg-gray-800 focus:outline-none focus:border-brand-orange focus:ring-1 focus:ring-brand-orange transition-all"
                required
              />
            </div>

            <button 
              type="submit" 
              :disabled="loading"
              class="w-full py-3.5 px-6 mt-2 rounded-xl bg-gradient-to-r from-[#FF6B00] to-[#FF4500] hover:from-[#E66000] hover:to-[#E03D00] text-white font-semibold text-base shadow-lg shadow-orange-500/30 transition-all duration-200 active:scale-[0.99] disabled:opacity-50"
            >
              {{ loading ? 'Enviando...' : 'Enviar' }}
            </button>
          </form>

        </div>
      </div>

    </div>
  </section>
</template>