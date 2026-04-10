<!-- layouts/AppLayout.vue -->
<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 flex">
    <!-- Sidebar -->
    <aside class="w-64 bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700 hidden md:block">
      <div class="p-6">
        <h1 class="text-2xl font-bold text-emerald-600">📚 AptLibrary</h1>
        <p class="text-xs text-gray-500 mt-1">100-House Community Library</p>
      </div>
      
      <nav class="px-3">
        <a v-for="item in menu" :key="item.name"
           :href="item.link"
           class="flex items-center gap-3 px-4 py-3 rounded-xl text-sm font-medium hover:bg-emerald-50 dark:hover:bg-gray-700 mb-1"
           :class="{ 'bg-emerald-50 text-emerald-700 dark:bg-gray-700': $route.path === item.link }">
          {{ item.icon }} {{ item.name }}
        </a>
      </nav>
    </aside>

    <!-- Main Content -->
    <div class="flex-1 flex flex-col">
      <!-- Top Navbar -->
      <header class="bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700 px-6 py-4 flex items-center justify-between">
        <div class="flex items-center gap-4">
          <button class="md:hidden">☰</button>
          <input type="text" placeholder="Search books..." 
                 class="w-80 bg-gray-100 dark:bg-gray-700 border-0 rounded-2xl px-5 py-2 text-sm focus:ring-2 focus:ring-emerald-500">
        </div>
        
        <div class="flex items-center gap-6">
          <div class="text-sm">
            Welcome, <span class="font-semibold">{{ userName }}</span>
          </div>
          <div class="w-9 h-9 bg-emerald-100 dark:bg-emerald-900 rounded-2xl flex items-center justify-center text-xl cursor-pointer">
            👤
          </div>
        </div>
      </header>

      <!-- Page Content -->
      <main class="flex-1 p-6 overflow-auto">
        <slot />
      </main>
    </div>
  </div>
</template>