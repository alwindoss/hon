<!-- views/Login.vue -->
<template>
  <div class="min-h-screen bg-gradient-to-br from-emerald-50 to-white dark:from-gray-900 dark:to-gray-800 flex items-center justify-center">
    <div class="bg-white dark:bg-gray-800 rounded-3xl shadow-2xl w-full max-w-md p-10">
      <div class="text-center mb-10">
        <div class="text-6xl mb-4">📚</div>
        <h1 class="text-3xl font-bold text-gray-900 dark:text-white">AptLibrary</h1>
        <p class="text-gray-500 mt-2">Community Book Lending System</p>
      </div>

      <form @submit.prevent="login">
        <div class="space-y-6">
          <div>
            <label class="block text-sm font-medium mb-2">Email / House Number</label>
            <input v-model="email" type="text" 
                   class="w-full px-5 py-4 bg-gray-100 dark:bg-gray-700 border-0 rounded-2xl focus:ring-2 focus:ring-emerald-500 outline-none">
          </div>
          <div>
            <label class="block text-sm font-medium mb-2">Password</label>
            <input v-model="password" type="password" 
                   class="w-full px-5 py-4 bg-gray-100 dark:bg-gray-700 border-0 rounded-2xl focus:ring-2 focus:ring-emerald-500 outline-none">
          </div>

          <button type="submit"
                  class="w-full py-4 bg-emerald-600 hover:bg-emerald-700 text-white font-semibold rounded-2xl transition">
            Login
          </button>
        </div>
      </form>

      <p class="text-center text-xs text-gray-500 mt-8">
        Only admin can create new member accounts
      </p>
    </div>
  </div>
</template>