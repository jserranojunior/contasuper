<template>
  <!-- Header Combinado -->
  <header class="bg-gray-900 border-b border-gray-800 px-4 md:px-6 py-4 flex flex-wrap items-center justify-between gap-4 w-full">
    
    <!-- Lado Esquerdo: Botão Drawer + Título / Slot -->
    <div class="flex items-center gap-3">

      
      <h1 class="text-xl font-bold text-white tracking-wide">
        Conta<span class="text-brand-orange"> Super</span>
      </h1>
      <slot name="title"></slot>
    </div>

    <!-- Centro / Direita: Links de Navegação com ACL e Botões de Sessão -->
    <div class="flex items-center flex-wrap gap-2 md:gap-4">
      <nav class="flex flex-wrap items-center gap-1.5 my-auto">
        <!-- Links de Visitante -->
        <router-link class="px-3 py-1.5 text-xs font-semibold rounded-lg bg-gray-800 hover:bg-gray-700 text-gray-200 transition-colors" v-if="!auth.checkStateToken()" :to="'/'">Home</router-link>
        <router-link class="px-3 py-1.5 text-xs font-semibold rounded-lg bg-gray-800 hover:bg-gray-700 text-gray-200 transition-colors" :to="'/login'" v-if="!auth.checkStateToken()">Login</router-link>
<!--         <router-link class="px-3 py-1.5 text-xs font-semibold rounded-lg bg-brand-orange hover:bg-brand-orange-hover text-white transition-colors" :to="'/cadastro'" v-if="!auth.checkStateToken()">Cadastro</router-link> -->
        
        <!-- Links Autenticados -->
        <router-link class="px-3 py-1.5 text-xs font-semibold rounded-lg bg-gray-800 hover:bg-gray-700 text-gray-200 transition-colors" :to="'/financeiro'" v-if="acl.checkIfExisteRoutes(4) && auth.checkStateToken()">Financeiro</router-link>

        <router-link class="px-3 py-1.5 text-xs font-semibold rounded-lg bg-gray-800 hover:bg-gray-700 text-gray-200 transition-colors" :to="'/usuarios'" v-if="acl.checkIfExisteRoutes(5) && auth.checkStateToken()">Usuários</router-link>
      </nav>

      <!-- Informações do Usuário & Sair -->
      <div class="flex items-center gap-3 pl-2 border-l border-gray-800" v-if="auth.checkStateToken()">
        <div class="text-right hidden sm:block">
          <p class="text-xs font-semibold text-gray-200">
            {{ auth.fields.email || 'Usuário Logado' }}
          </p>
          <span class="text-[10px] text-brand-orange font-medium">Sessão Ativa</span>
        </div>

        <button 
          @click="sair"
          class="px-3 py-1.5 text-xs font-semibold rounded-xl bg-gray-800 hover:bg-red-950/50 hover:text-red-400 border border-gray-700 hover:border-red-800 text-gray-300 transition-all cursor-pointer"
       v-if="auth.checkStateToken()"
          >
          Sair
        </button>
      </div>
    </div>

  </header>
</template>

<script setup lang="ts">
import useStore from "../helpers/stores/store"

const { auth, router, acl } = useStore()

function sair() {
  if (auth && typeof auth.Logout === 'function') {
    auth.Logout()
  }
  router.push({ path: '/login' })
}
</script>