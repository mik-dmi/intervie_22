<script setup >
import {computed, onMounted, ref, watch} from 'vue';

import Table from './ui/table/Table.vue';
import TableHead from './ui/table/TableHead.vue';
import TableBody from './ui/table/TableBody.vue';
import TableCaption from './ui/table/TableCaption.vue';
import TableCell from './ui/table/TableCell.vue';
import TableRow from './ui/table/TableRow.vue';
import TableHeader from './ui/table/TableHeader.vue';
import Button from './ui/button/Button.vue'
import SearchBar from './SearchBar.vue'

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'

const open = ref(false);
const selectedContainer = ref(null);
const searchFilter = ref("")

const containersData = ref([]);
                            
const containerStats = ref();

const eventSourceRef = ref(null);

const fetchContainersListData = async () => {
  try {
    const response = await fetch('/server/containers', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    });

    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    containersData.value = await response.json();
    console.log("Here containers: \n");
    console.log(containersData.value);
  } catch (error) {
    console.error('Failed to fetch containers:', error);
    throw error; 
  }
};

onMounted(fetchContainersListData);

const handleRefresh = async () => {
  try {
    await fetchContainersListData();
    console.log("Sucess Updating")
  
  } catch (error) {
  
    console.log("Error Updating")
    
  }
};
/*
const getStats = (container) =>{
  containerStats.value = container
  console.log( container)
}*/
const getStats = async (container) => {
  try {
    const eventSource = new EventSource(`/server/sse?containerId=${container.Id}`);
    
    eventSourceRef.value = eventSource;

    eventSource.onerror = (err) => {
      console.log("onerror", err);
    };
    
    eventSource.onmessage = (event )=> {
      try {
        const data = JSON.parse(event.data);
        containerStats.value = data
        console.log("Aqui ")
        console.log(containerStats.value.cpu_usage_percent);

       
        
        // Update your Vue component state with the new stats here
      } catch (error) {
        console.error('Failed to parse event data:', error);
      }
    };

  } catch (error) {
    console.error('Failed to fetch container stats:', error);
    throw error;
  }
};


const handleRowClick = (container) => {
      selectedContainer.value = container;
      open.value = true;
    };
//clean the containersStats for now, later will close the conection 
watch(open, (newVal) => {
      if (newVal && selectedContainer.value) {
        getStats(selectedContainer.value);
      }     else {
      if (eventSourceRef.value) {
        eventSourceRef.value.close();  // Close the EventSource connection
        eventSourceRef.value = null;  // Clear the ref
      }
      containerStats.value = null;
    }
})


watch(selectedContainer, (newContainerStats) => {
     console.log(`Stats is ${newContainerStats.cpu_usage_percent}`)
}); 



const filteredImages = computed(()=>{
  if(searchFilter.value != ""){
    return containersData.value.filter(item=> item.Image.includes(searchFilter.value))
  }
  return containersData.value
})
const handleSearch = (search)=>{
  searchFilter.value = search;
}

</script>
<template>
  
  <section class =" flex flex-col justify-center   items-center self-start py-[5rem] w-full">
    <h1 class="text-[3rem] font-semibold	pb-2 self-start	" >Docker Container List</h1>
    <div class="flex  self-start ">
      <SearchBar @search="handleSearch" />
      <Button @click="handleRefresh">
    Refresh</Button>
    </div>
    <Table v-if="containersData.length"  class=" w-full text-start text-sm font-light text-surface  rounded-t-lg m-5 mx-auto bg-secondary text-white ">      
      <TableCaption>Information about Docker Containers</TableCaption>
      <TableHeader  class="border-b border-neutral-200 font-medium hover:bg-secondary ">
      <TableRow class="text-left border-b border-secondary/60 text-[1.2rem]  hover:bg-secondary ">
        <TableHead class="px-4    py-3 text-white ">Container Id</TableHead>
        <TableHead class="px-4 py-3 text-white" >Container Names</TableHead>
        <TableHead class="px-4 py-3 text-white" >Conatianer Image</TableHead>
        <TableHead class="px-4 py-3 text-white" >Container Status</TableHead>
        <TableHead class="px-4 py-3 text-white " >CreationDate </TableHead>
      </TableRow>

      </TableHeader>
     
          <TableBody>
            
              
                <TableRow  class="bg-[#142a44] border-secondary/60 border-b hover:bg-[#2d496c] cursor-pointer "  v-for="(container) in filteredImages" :key="container.Id" @click="handleRowClick(container)"> 
                  
                  <TableCell class=" border-b border-secondary/60">{{container.Id}}</TableCell>
                  <TableCell class=" border-b border-secondary/60"><span v-for="(name, index) in container.Names" :key="index">{{ name }}</span></TableCell>
                  <TableCell class=" border-b border-secondary/60">{{container.Image}}</TableCell>
                  <TableCell class=" border-b border-secondary/60">{{container.Status}} </TableCell>
                  <TableCell class=" border-b border-secondary/60">{{container.CreationDate}} </TableCell>
                </TableRow>
          </TableBody>
        </Table>
    <div class="flex min-w-full max-h-full justify-center align-middle " v-else>
      <div class="border-gray-300 h-20 w-20 animate-spin rounded-full border-8 border-t-secondary mt-[4rem]" />
    </div>
        <Dialog v-model:open="open">  
        <DialogContent  class="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Container Resource Stats</DialogTitle>
            <DialogDescription>
              Resource usage statistics of the contianer 
            </DialogDescription>
          </DialogHeader>
          <div class="grid gap-4 py-4">
            <div class="grid gap-4 py-4" v-if=containerStats >
              <div><strong>CPU Usage:</strong> {{ containerStats.cpu_usage_percent }}%</div>
              <div><strong>Memory Usage:</strong> {{ containerStats.memory_usage_percent }}%</div>
              <div><strong>Network I/O:</strong> {{ containerStats.rx_network_bytes }}kB / {{ containerStats.tx_network_bytes }}kB</div>
            </div>
            <div class="flex min-w-full max-h-full justify-center align-middle " v-if="open && !containerStats" ><div class="border-gray-300 h-20 w-20 animate-spin rounded-full border-8 border-t-secondary  self-center" /></div>
          </div>
        </DialogContent>
      </Dialog>

 

    </section>


    <div>

    </div>
   
</template>

