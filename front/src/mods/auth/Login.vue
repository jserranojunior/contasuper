<template>
  <div class="min-h-screen flex items-start justify-center p-4 pt-12 md:pt-20">
    <div class="bg-gray-900 text-white p-6 md:p-8 rounded-2xl shadow-2xl w-full max-w-md border border-gray-800">
      
      <!-- Mensagens de Alerta (Erro / Sucesso) -->
      <div v-if="auth && auth.auth" class="mb-4">
        <div 
          v-if="auth.auth.erro"
          class="block text-xs font-medium text-red-200 text-center bg-red-950/80 border border-red-800 p-3 rounded-xl"
          role="alert"
        >
          {{ auth.auth.erro }}
        </div>
        <div 
          v-if="auth.checkStateToken()"
          class="block text-xs font-medium text-emerald-200 text-center bg-emerald-950/80 border border-emerald-800 p-3 rounded-xl"
          role="alert"
        >
          Logado com sucesso!
        </div>
      </div>

      <h3 class="text-xl md:text-2xl font-bold text-gray-100 mb-6 text-center">
        Acessar Conta<sup class="text-xs font-normal">®</sup>
      </h3>

      <!-- Formulário de Acesso -->
      <form @submit.prevent="logar" class="space-y-4">
        
        <!-- E-mail -->
        <div>
          <label for="email" class="block text-xs font-medium text-gray-100 mb-1">E-mail</label>
          <input 
            type="email" 
            id="email" 
            v-model="auth.fields.email" 
            placeholder="seu.email@exemplo.com"
            class="w-full px-4 py-3 rounded-xl border border-gray-500 text-sm text-gray-100 focus:bg-gray-800 focus:outline-none focus:border-brand-orange focus:ring-1 focus:ring-brand-orange transition-all"
            required
          />
        </div>

        <!-- Senha -->
        <div>
          <label for="senha" class="block text-xs font-medium text-gray-100 mb-1">Senha</label>
          <input 
            type="password" 
            id="senha" 
            v-model="auth.fields.password" 
            placeholder="••••••••"
            class="w-full px-4 py-3 rounded-xl border border-gray-500 text-sm text-gray-100 focus:bg-gray-800 focus:outline-none focus:border-brand-orange focus:ring-1 focus:ring-brand-orange transition-all"
            required
          />
        </div>

        <!-- Link Esqueci minha senha -->
        <!-- <div class="flex justify-end text-xs pt-1">
          <router-link to="/esqueci-senha" class="text-orange-200 hover:text-white transition-colors">
            Esqueceu a senha?
          </router-link>
        </div> -->

        <!-- Botão Acessar -->
        <button 
          type="submit" 
          class="w-full py-3.5 px-6 mt-2 rounded-xl bg-gradient-to-r from-[#FF6B00] to-[#FF4500] hover:from-[#E66000] hover:to-[#E03D00] text-white font-semibold text-base shadow-lg shadow-orange-500/30 transition-all duration-200 active:scale-[0.99] cursor-pointer"
        >
          Acessar
        </button>

        <!-- Botão/Link de Cadastro -->
        <div class="text-center pt-2 text-xs text-gray-400">
          Ainda não tem uma conta? 
          <router-link to="/" class="text-orange-200 hover:text-white font-semibold ml-1">
            Solicitar Acesso
          </router-link>
        </div>

      </form>

    </div>
  </div>
</template>

<script setup lang="ts">
import useStore from "../../helpers/stores/store"
import { onMounted } from "vue";

let { auth, router, acl } = useStore()

function redirectPageTo(url: string) { 
  router.push({ path: url })              
}

function logar() {
  auth.Login().then(async (res: boolean) => {
    if (res) {
      await acl.getUserAcl().then(() => {
        acl.generateRoutesEnableWithUserAcls();
      }).then(() => {
        redirectPageTo("/financeiro")
      });
    }
  })
}

onMounted(async () => {
  document.addEventListener('keyup', function (event) {
    if (event.key == "Enter") {
      logar();
    }
  });

  await auth.isLogged().then((res) => {
    if (res == true) {
      redirectPageTo("/financeiro")
    }
  })
})
</script>