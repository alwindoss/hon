<!-- views/AddBook.vue -->
<template>
  <div class="max-w-2xl mx-auto bg-white dark:bg-gray-800 rounded-3xl p-10">
    <h1 class="text-3xl font-bold mb-10">Add New Book</h1>
    
    <div class="grid grid-cols-2 gap-6">
      <div class="col-span-2">
        <label>Title</label>
        <input v-model="form.title" class="w-full px-5 py-4 rounded-2xl bg-gray-100 dark:bg-gray-700 border-0">
      </div>
      <div>
        <label>Author(s)</label>
        <input v-model="form.author" class="w-full px-5 py-4 rounded-2xl bg-gray-100 dark:bg-gray-700 border-0">
      </div>
      <div>
        <label>ISBN</label>
        <input v-model="form.isbn" class="w-full px-5 py-4 rounded-2xl bg-gray-100 dark:bg-gray-700 border-0">
      </div>
      <!-- More fields: genre, copies, year, etc. -->
    </div>
    
    <button class="mt-10 w-full py-5 bg-emerald-600 text-white rounded-2xl text-lg font-semibold">
      Save Book
    </button>
  </div>
</template>