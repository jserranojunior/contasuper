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
              <tr class="hover:bg-gray-800/30 transition-colors">
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