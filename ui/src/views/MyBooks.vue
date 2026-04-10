<!-- views/MyBooks.vue -->
<template>
  <div>
    <h1 class="text-3xl font-bold mb-8">My Borrowed Books</h1>
    
    <div class="space-y-6">
      <div v-for="item in borrowed" :key="item.id"
           class="bg-white dark:bg-gray-800 rounded-3xl p-6 flex gap-6 items-center">
        <div class="w-20 h-28 bg-gray-200 dark:bg-gray-700 rounded-2xl flex-shrink-0 flex items-center justify-center text-5xl">
          📖
        </div>
        <div class="flex-1">
          <h3 class="font-semibold">{{ item.title }}</h3>
          <p class="text-sm text-gray-500">{{ item.author }}</p>
          <p class="text-xs mt-3">
            Issued: {{ item.issueDate }} &nbsp; | &nbsp; 
            Due: <span :class="item.overdue ? 'text-red-600 font-medium' : ''">{{ item.dueDate }}</span>
          </p>
        </div>
        
        <div class="text-right">
          <button @click="returnBook(item)"
                  class="px-6 py-3 bg-amber-600 text-white rounded-2xl text-sm">
            Return Book
          </button>
          <p v-if="item.overdue" class="text-red-600 text-xs mt-2">Overdue • Fine: ₹{{ item.fine }}</p>
        </div>
      </div>
    </div>
  </div>
</template>