

<template>
    
    <main class="min-h-[100vh] flex flex-col justify-between">
        <NavBar />
        <div class="mx-auto max-w-2xl px-8  sm:px-10 lg:max-w-7xl lg:min-h-[70vh] min-h-[50vh] flex w-full">
            <router-view />
        </div>
        <Footer />
    </main>
    
  </template>
  
  <script setup>
  import NavBar from "./NavBar.vue";
  import Footer from "./Footer.vue";
  </script>