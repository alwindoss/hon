<!-- views/AdminDashboard.vue -->
<template>
  <div>
    <h1 class="text-3xl font-bold mb-8">Library Dashboard</h1>
    
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-white dark:bg-gray-800 p-6 rounded-3xl shadow">
        <div class="flex justify-between">
          <div>
            <p class="text-gray-500 text-sm">Total Books</p>
            <p class="text-4xl font-bold mt-2">1,248</p>
          </div>
          <div class="text-5xl opacity-20">📚</div>
        </div>
      </div>
      
      <div class="bg-white dark:bg-gray-800 p-6 rounded-3xl shadow">
        <div class="flex justify-between">
          <div>
            <p class="text-gray-500 text-sm">Currently Borrowed</p>
            <p class="text-4xl font-bold mt-2 text-amber-600">87</p>
          </div>
          <div class="text-5xl opacity-20">🔄</div>
        </div>
      </div>
      
      <div class="bg-white dark:bg-gray-800 p-6 rounded-3xl shadow">
        <div class="flex justify-between">
          <div>
            <p class="text-gray-500 text-sm">Overdue Books</p>
            <p class="text-4xl font-bold mt-2 text-red-600">12</p>
          </div>
          <div class="text-5xl opacity-20">⏰</div>
        </div>
      </div>
      
      <div class="bg-white dark:bg-gray-800 p-6 rounded-3xl shadow">
        <div class="flex justify-between">
          <div>
            <p class="text-gray-500 text-sm">Active Members</p>
            <p class="text-4xl font-bold mt-2">92</p>
          </div>
          <div class="text-5xl opacity-20">👥</div>
        </div>
      </div>
    </div>

    <!-- Recent Activity -->
    <div class="mt-10 bg-white dark:bg-gray-800 rounded-3xl p-6">
      <h2 class="font-semibold mb-4">Recent Transactions</h2>
      <table class="w-full">
        <thead>
          <tr class="border-b">
            <th class="text-left py-4">Member</th>
            <th class="text-left py-4">Book</th>
            <th class="text-left py-4">Action</th>
            <th class="text-left py-4">Date</th>
          </tr>
        </thead>
        <tbody>
          <!-- Populate with v-for -->
        </tbody>
      </table>
    </div>
  </div>
</template>