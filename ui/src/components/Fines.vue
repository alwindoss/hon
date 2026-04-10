<!-- views/Fines.vue -->
<template>
  <div>
    <h1 class="text-3xl font-bold mb-8">Fines & Payments</h1>
    
    <div class="bg-white dark:bg-gray-800 rounded-3xl p-8">
      <div class="text-5xl font-bold text-red-600 mb-2">₹{{ totalFine }}</div>
      <p class="text-gray-500">Total pending fine</p>
      
      <div class="mt-10 space-y-6">
        <!-- List of individual fines with "Mark as Paid" button for admin -->
      </div>
    </div>
  </div>
</template>