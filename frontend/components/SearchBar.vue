<script setup>
import { Search } from 'lucide-vue-next'
import Input from './ui/input/Input.vue';

const emit = defineEmits( ['search'])


const search = (e) =>{
    emit('search', e.target.value)
}
</script>

<template>
    <div class="relative max-w-sm items-center w-fit lg:w-[20rem] mr-4 ">
        <Input @input="search" id="search" type="text" placeholder="Search for Container Image..." class="pl-10 text-[1.1rem] focus:border-secondary focus:border-[0.2rem]" />
        <span class="absolute start-0 inset-y-0 flex items-center justify-center px-2">
          <Search class="size-6 text-muted-foreground" />
        </span>
    </div>


</template>