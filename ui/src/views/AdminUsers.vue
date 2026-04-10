<!-- views/AdminUsers.vue -->
<template>
  <div>
    <div class="flex justify-between mb-8">
      <h1 class="text-3xl font-bold">Members ({{ users.length }})</h1>
      <button @click="showAddModal = true"
              class="px-6 py-3 bg-emerald-600 text-white rounded-2xl flex items-center gap-2">
        + Add New Member
      </button>
    </div>

    <div class="bg-white dark:bg-gray-800 rounded-3xl overflow-hidden">
      <table class="w-full">
        <thead class="bg-gray-50 dark:bg-gray-700">
          <tr>
            <th class="text-left p-6">House No.</th>
            <th class="text-left p-6">Name</th>
            <th class="text-left p-6">Phone</th>
            <th class="text-left p-6">Books Borrowed</th>
            <th class="text-left p-6">Status</th>
            <th class="p-6"></th>
          </tr>
        </thead>
        <tbody class="divide-y">
          <!-- v-for users here -->
        </tbody>
      </table>
    </div>
  </div>
</template>