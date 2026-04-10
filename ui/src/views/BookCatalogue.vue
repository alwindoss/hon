<!-- views/BookCatalogue.vue -->
<template>
  <div>
    <div class="flex justify-between items-center mb-8">
      <h1 class="text-3xl font-bold">Book Catalogue</h1>
      <div v-if="isAdmin" class="flex gap-3">
        <button @click="addNewBook" 
                class="px-6 py-3 bg-emerald-600 text-white rounded-2xl flex items-center gap-2">
          + Add New Book
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
      <div v-for="book in books" :key="book.id"
           class="bg-white dark:bg-gray-800 rounded-3xl overflow-hidden shadow hover:shadow-xl transition">
        <div class="h-64 bg-gray-200 dark:bg-gray-700 flex items-center justify-center text-8xl">
          {{ book.cover || '📖' }}
        </div>
        <div class="p-5">
          <h3 class="font-semibold line-clamp-2">{{ book.title }}</h3>
          <p class="text-sm text-gray-500">{{ book.author }}</p>
          
          <div class="flex justify-between items-center mt-6">
            <span :class="book.available ? 'text-emerald-600' : 'text-red-500'" 
                  class="font-medium text-sm">
              {{ book.available ? 'Available' : 'Borrowed' }}
            </span>
            <button v-if="book.available && !isAdmin"
                    @click="borrowBook(book)"
                    class="px-5 py-2 bg-emerald-600 text-white text-sm rounded-2xl">
              Borrow
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>