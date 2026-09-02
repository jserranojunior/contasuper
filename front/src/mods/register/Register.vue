<template>
  <div class="min-h-screen flex items-start justify-center p-4 pt-12 md:pt-20">
    <div class="bg-gray-900 text-white p-6 md:p-8 rounded-2xl shadow-2xl w-full max-w-md border border-gray-800">
      
      <!-- Mensagens de Alerta (Erro / Sucesso) -->
      <span v-if="register && register.fields">
        <div
          v-if="register.erro"
          class="mb-4 block text-xs font-medium text-red-200 text-center bg-red-950/80 border border-red-800 p-3 rounded-xl"
          role="alert"
        >
          {{ register.erro }}
        </div>
      </span>

      <h3 class="text-xl md:text-2xl font-bold text-gray-100 mb-6 text-center">
        Cadastro<sup class="text-xs font-normal">®</sup>
      </h3>

      <!-- Formulário de Cadastro -->
      <form @submit.prevent class="space-y-4">

        <!-- Nome Completo -->
        <div>
          <label for="nome" class="block text-xs font-medium text-gray-100 mb-1">Nome Completo</label>
          <input 
            v-model="register.fields.name" 
            type="text" 
            id="nome" 
            name="nome" 
            placeholder="Seu nome completo"
            class="w-full px-4 py-3 rounded-xl border border-gray-500 text-sm text-gray-100 focus:bg-gray-800 focus:outline-none focus:border-brand-orange focus:ring-1 focus:ring-brand-orange transition-all"
          >
        </div>

        <!-- E-mail -->
        <div>
          <label for="cadastroEmail" class="block text-xs font-medium text-gray-100 mb-1">E-mail</label>
          <input 
            v-model="register.fields.email" 
            type="email" 
            id="cadastroEmail" 
            name="cadastroEmail" 
            placeholder="seu.email@exemplo.com"
            class="w-full px-4 py-3 rounded-xl border border-gray-500 text-sm text-gray-100 focus:bg-gray-800 focus:outline-none focus:border-brand-orange focus:ring-1 focus:ring-brand-orange transition-all"
          >
        </div>

        <!-- Telefone -->
        <div>
          <label for="telefone" class="block text-xs font-medium text-gray-100 mb-1">Telefone</label>
          <div class="flex gap-2">
            <div class="dropdown w-auto">
              <label tabindex="0" class="btn btn-sm h-full bg-gray-800 border-gray-500 text-white hover:bg-gray-700 rounded-xl px-3">
                {{ register.phoneCode }}
              </label>
              <ul tabindex="0" class="dropdown-content menu p-2 shadow bg-gray-900 border border-gray-700 rounded-box w-40 z-20 text-white">
                <li @click="register.phoneCode = '+55'">
                  <a class="flex items-center gap-2 hover:bg-gray-800 rounded-lg">
                    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Flag_of_Brazil.svg/33px-Flag_of_Brazil.svg.png" class="w-5" />
                    +55
                  </a>
                </li>
                <li @click="register.phoneCode = '+591'">
                  <a class="flex items-center gap-2 hover:bg-gray-800 rounded-lg">
                    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/4/48/Flag_of_Bolivia.svg/33px-Flag_of_Bolivia.svg.png" class="w-5" />
                    +591
                  </a>
                </li>
              </ul>
            </div>
            <div class="w-full">
              <input
                id="cellphone"
                v-model="register.fields.cellphone"
                type="text"
                class="w-full px-4 py-3 rounded-xl border border-gray-500 text-sm text-gray-100 focus:bg-gray-800 focus:outline-none focus:border-brand-orange focus:ring-1 focus:ring-brand-orange transition-all"
                placeholder="(00) 00000-0000"
                v-maska="'(##) #####-####'"
              />
            </div>
          </div>
        </div>

        <!-- Data de Nascimento -->
        <div>
          <label for="dataNascimento" class="block text-xs font-medium text-gray-100 mb-1">Data de Nascimento</label>
          <input 
            v-maska="'##/##/####'" 
            v-model="register.fields.dtnascimento" 
            type="text" 
            id="dataNascimento" 
            name="dataNascimento" 
            class="w-full px-4 py-3 rounded-xl border border-gray-500 text-sm text-gray-100 focus:bg-gray-800 focus:outline-none focus:border-brand-orange focus:ring-1 focus:ring-brand-orange transition-all" 
            placeholder="dd/mm/aaaa"
          >
        </div>

        <!-- Senha -->
        <div>
          <label for="cadastroSenha" class="block text-xs font-medium text-gray-100 mb-1">Senha</label>
          <input 
            v-model="register.fields.password" 
            type="password" 
            id="cadastroSenha" 
            name="cadastroSenha" 
            placeholder="••••••••"
            class="w-full px-4 py-3 rounded-xl border border-gray-500 text-sm text-gray-100 focus:bg-gray-800 focus:outline-none focus:border-brand-orange focus:ring-1 focus:ring-brand-orange transition-all"
          >
        </div>

        <!-- Confirmação de Senha -->
        <div>
          <label for="confirmacaoSenha" class="block text-xs font-medium text-gray-100 mb-1">Confirmação de Senha</label>
          <input 
            v-model="register.confirmPassword" 
            type="password" 
            id="confirmacaoSenha" 
            name="confirmacaoSenha" 
            placeholder="••••••••"
            class="w-full px-4 py-3 rounded-xl border border-gray-500 text-sm text-gray-100 focus:bg-gray-800 focus:outline-none focus:border-brand-orange focus:ring-1 focus:ring-brand-orange transition-all"
          >
        </div>

      </form>

      <div class="mt-6 border-t border-gray-800 pt-4"></div>

      <!-- Botões de Ação (Acessar e Cadastrar) -->
      <div class="grid grid-cols-2 gap-3">
        <a href="/login" class="w-full">
          <button class="w-full py-3.5 px-4 rounded-xl bg-gray-800 hover:bg-gray-700 text-gray-200 font-semibold text-sm transition-all duration-200 border border-gray-700 text-center cursor-pointer">
            Acessar
          </button>
        </a>
        <button 
          type="button" 
          @click="cadastrar()" 
          class="w-full py-3.5 px-4 rounded-xl bg-gradient-to-r from-[#FF6B00] to-[#FF4500] hover:from-[#E66000] hover:to-[#E03D00] text-white font-semibold text-sm shadow-lg shadow-orange-500/30 transition-all duration-200 active:scale-[0.99] cursor-pointer"
        >
          Cadastrar
        </button>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import useStore from "../../helpers/stores/store";
import { onMounted } from "vue";

let { register, router, auth } = useStore();

function redirectPageTo(url: string) {
  router.push({ path: url });
}

function cadastrar() {
  register.Register().then(async (res: boolean) => {
    if (res) {
      await logar();
    }
  });
}

async function logar() {
  auth.fields.email = register.fields.email;
  auth.fields.password = register.fields.password;
  await auth.Login().then((res: boolean) => {
    if (res) {
      redirectPageTo("/financeiro");
    }
  });
}

onMounted(() => {
  document.addEventListener("keyup", (event) => {
    if (event.key == "Enter") {
      cadastrar();
    }
  });
});
</script>