<!-- components/BorrowModal.vue -->
<template>
  <div class="fixed inset-0 bg-black/60 flex items-center justify-center z-50">
    <div class="bg-white dark:bg-gray-800 rounded-3xl max-w-lg w-full p-8">
      <h2 class="text-2xl font-bold mb-6">Borrow Book</h2>
      <p class="mb-8">You are borrowing <strong>"{{ selectedBook.title }}"</strong></p>
      
      <div class="space-y-4">
        <div>Due Date: <strong>{{ dueDate }}</strong> (14 days)</div>
        <button @click="confirmBorrow"
                class="w-full py-4 bg-emerald-600 text-white rounded-2xl font-semibold">
          Confirm Borrow
        </button>
      </div>
    </div>
  </div>
</template>