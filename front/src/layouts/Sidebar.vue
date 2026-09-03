<template>
  <div class="menu p-4 overflow-y-auto w-80 min-h-full text-gray-200 border-r border-gray-800 bg-gray-900 shadow-xl flex flex-col gap-1.5">
    
    <!-- Título do Menu Lateral -->
    <div class="px-3 py-2 mb-2 border-b border-gray-800">
      <h2 class="text-xs font-bold uppercase tracking-wider text-orange-500">Navegação</h2>
    </div>

    <!-- Links de Navegação -->
    <router-link 
      class="px-4 py-2.5 text-sm font-semibold rounded-xl bg-gray-800 hover:bg-gray-700 hover:text-white text-gray-300 transition-all border border-gray-700/50" 
      to="/"
    >
      Home
    </router-link>

    <router-link 
      class="px-4 py-2.5 text-sm font-semibold rounded-xl bg-gray-800 hover:bg-gray-700 hover:text-white text-gray-300 transition-all border border-gray-700/50" 
      to="/login" 
      v-if="!auth.checkStateToken()"
    >
      Login
    </router-link>

  <!--   <router-link 
      class="px-4 py-2.5 text-sm font-semibold rounded-xl bg-brand-orange hover:bg-brand-orange-hover text-white transition-all shadow-md shadow-orange-500/10" 
      to="/cadastro" 
      v-if="!auth.checkStateToken()"
    >
      Cadastro
    </router-link> -->

    <router-link 
      class="px-4 py-2.5 text-sm font-semibold rounded-xl bg-gray-800 hover:bg-gray-700 hover:text-white text-gray-300 transition-all border border-gray-700/50" 
      to="/financeiro" 
      v-if="acl.checkIfExisteRoutes(4) && auth.checkStateToken()"
    >
      Financeiro
    </router-link>

    <router-link 
      class="px-4 py-2.5 text-sm font-semibold rounded-xl bg-gray-800 hover:bg-gray-700 hover:text-white text-gray-300 transition-all border border-gray-700/50" 
      to="/financeiro" 
      v-if="auth.checkStateToken()"
    >
      Depósito
    </router-link>
        
    <router-link 
      class="px-4 py-2.5 text-sm font-semibold rounded-xl bg-gray-800 hover:bg-gray-700 hover:text-white text-gray-300 transition-all border border-gray-700/50" 
      to="/financeiro" 
      v-if="auth.checkStateToken()"
    >
      Extrato
    </router-link>

    <router-link 
      class="px-4 py-2.5 text-sm font-semibold rounded-xl bg-gray-800 hover:bg-gray-700 hover:text-white text-gray-300 transition-all border border-gray-700/50" 
      to="/financeiro" 
      v-if="auth.checkStateToken()"
    >
      Transferência
    </router-link>

    <router-link 
      class="px-4 py-2.5 text-sm font-semibold rounded-xl bg-gray-800 hover:bg-gray-700 hover:text-white text-gray-300 transition-all border border-gray-700/50" 
      to="/usuarios" 
      v-if="acl.checkIfExisteRoutes(5) && auth.checkStateToken()"
    >
      Usuários
    </router-link>

    <!-- Botão de Sair com visual de alerta/vermelho suave -->
    <button 
      class="mt-4 px-4 py-2.5 text-sm font-semibold rounded-xl bg-gray-800 hover:bg-red-950/50 hover:text-red-400 border border-gray-700 hover:border-red-800 text-gray-300 transition-all text-left cursor-pointer" 
      @click="sair()" 
      v-if="auth.checkStateToken()"
      
    >
      Sair
    </button> 
  </div>
</template>

<script lang="ts" setup>
import { onBeforeMount } from "@vue/runtime-core";
import useStore from "../helpers/stores/store"
let { auth, router, acl } = useStore()

function redirectPageTo(url: string) { 
    console.log("Redirecionando")
    router.push({ path: url })              
}

function sair() {
    auth.Logout()
    acl.clearRoutesEnableWithUserAcls()
    redirectPageTo("/login")
}

onBeforeMount(async () => {
  await auth.isLogged().then(async () => {
    await acl.getUserAcl().then(() => {
          acl.generateRoutesEnableWithUserAcls();
        })
  })
})
</script>