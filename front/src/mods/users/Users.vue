<template>
  <div class="min-h-screen bg-gray-950 text-gray-100 flex flex-col">
 

    <main class="flex-1 p-6 md:p-8 max-w-7xl w-full mx-auto">
      <div class="bg-gray-900 border border-gray-800 rounded-2xl shadow-xl p-6">
        <h2 class="text-xl font-bold text-white mb-4">Usuários</h2>
        
        <div class="overflow-x-auto">
          <table class="table w-full border border-gray-800">
            <thead>
              <tr class="text-gray-400 border-b border-gray-800">
                <th class="bg-gray-900">ID</th>
                <th class="bg-gray-900">Nome</th>
                <th class="bg-gray-900">Telefone</th>
                <th class="bg-gray-900">E-mail</th>
              </tr>
            </thead>
            <tbody>
             
              <tr 
                v-for="user in users.users" 
                :key="user.ID" 
                @click="selectUser(user.ID)" 
                class="cursor-pointer hover:bg-gray-800/60 border-b border-gray-800/50 transition-colors"
              >
              
                <td class="font-medium text-gray-300">{{ user.ID }}</td>
                <td class="text-white font-semibold">{{ user.name }}</td>
                <td class="text-gray-300">{{ user.cellphone }}</td>
                <td class="text-gray-300">{{ user.email }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import useStore from "../../helpers/stores/store"
import { onMounted } from 'vue';

let { users } = useStore()

defineProps({
  selectUser: {
    default: () => {},
    required: true,
    type: Function
  }
})

onMounted(async () => {
  await users.getAllUsers()
})
</script>