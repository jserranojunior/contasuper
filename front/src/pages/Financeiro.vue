<template>
  
  <div class="min-h-screen bg-gray-950 text-gray-100 flex flex-col">
    <!-- Header / Navbar -->


    <!-- Conteúdo Principal -->
    <main class="flex-1 p-6 md:p-8 max-w-7xl w-full mx-auto space-y-6">

        <div class="bg-gradient-to-r from-gray-900 via-gray-900 to-orange-950/30 border border-gray-800 rounded-2xl p-6 md:p-8 shadow-xl">
          <h2 class="text-2xl md:text-3xl font-bold text-white mb-2">
            Bem-vindo de volta!
          </h2>
          <p class="text-sm text-gray-400">
            Gerencie suas informações e navegue utilizando as opções do menu acima.
          </p>
        </div>
      
 

      <!-- Grid de Indicadores / Métricas -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        
        <div class="bg-gray-900 border border-gray-800 p-6 rounded-2xl shadow-lg">
          <p class="text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Status da Conta</p>
          <div class="flex items-center justify-between">
            <span class="text-2xl font-bold text-emerald-400">Ativa</span>
            <span class="px-2.5 py-1 text-xs font-medium bg-emerald-950 text-emerald-300 border border-emerald-800 rounded-lg">Verificado</span>
          </div>
        </div>

        <div class="bg-gray-900 border border-gray-800 p-6 rounded-2xl shadow-lg">
          <p class="text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Permissões (ACL)</p>
          <div class="flex items-center justify-between">
            <span class="text-2xl font-bold text-white">Habilitadas</span>
            <span class="px-2.5 py-1 text-xs font-medium bg-orange-950 text-orange-300 border border-orange-800 rounded-lg">
             <div>{{ auth && auth.auth && auth.auth.user && auth.auth.user.role_id === 1 ? 'Administrador' : 'Cliente'   }}</div>

             
            </span>
          </div>
        </div>

        

        <div class="bg-gray-900 border border-gray-800 p-6 rounded-2xl shadow-lg sm:col-span-2 lg:col-span-1">
          <p class="text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Último Acesso</p>
          <span class="text-2xl font-bold text-gray-200">Hoje</span>
        </div>

      </div>

      <!-- Seção de Ações Rápidas -->
      <div class="bg-gray-900 border border-gray-800 rounded-2xl p-6 shadow-lg space-y-4">
        <h3 class="text-lg font-bold text-gray-200">Ações Rápidas</h3>
        
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
          <button 
            @click="router.push('/financeiro')" 
            class="p-4 rounded-xl bg-gray-800/60 hover:bg-gray-800 border border-gray-700/60 hover:border-orange-500/50 text-left transition-all group cursor-pointer"
          >
            <h4 class="font-semibold text-orange-400 group-hover:text-orange-300 mb-1">Módulo Financeiro</h4>
            <p class="text-xs text-gray-400">Acesse seus lançamentos e relatórios.</p>
          </button>

          <button 
            class="p-4 rounded-xl bg-gray-800/60 hover:bg-gray-800 border border-gray-700/60 hover:border-orange-500/50 text-left transition-all group cursor-pointer"
          >
            <h4 class="font-semibold text-orange-400 group-hover:text-orange-300 mb-1">Perfil do Usuário</h4>
            <p class="text-xs text-gray-400">Atualize suas credenciais e e-mail.</p>
          </button>

          <button 
            class="p-4 rounded-xl bg-gray-800/60 hover:bg-gray-800 border border-gray-700/60 hover:border-orange-500/50 text-left transition-all group cursor-pointer"
          >
            <h4 class="font-semibold text-orange-400 group-hover:text-orange-300 mb-1">Suporte</h4>
            <p class="text-xs text-gray-400">Entre em contato com a equipe de atendimento.</p>
          </button>
        </div>
      </div>

    </main>
  </div>
</template>

<script setup lang="ts">
import useStore from "../helpers/stores/store"
import { onMounted } from "vue"
import { onBeforeMount, watch } from "@vue/runtime-core";


const { auth, router } = useStore()


onMounted(async () => {
  // Verifica se o usuário está logado ao carregar a tela
  if (auth && typeof auth.isLogged === 'function') {
    const logged = await auth.isLogged()
    if (!logged) {
      router.push({ path: '/login' })
    }
  }
})


let { acl} = useStore()
function redirectPageTo(url:string){ 
    console.log("Redirecionando")
    router.push({ path: url })              
  }
 function sair(){
      auth.Logout()
      acl.clearRoutesEnableWithUserAcls()
      redirectPageTo("/login")
  }

onBeforeMount(()=>{
  auth.isLogged()
})
</script>