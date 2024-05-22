<script setup>
import {onMounted, ref} from 'vue';



import Table from './ui/table/Table.vue';
import TableHead from './ui/table/TableHead.vue';
import TableBody from './ui/table/TableBody.vue';
import TableCaption from './ui/table/TableCaption.vue';
import TableCell from './ui/table/TableCell.vue';
import TableRow from './ui/table/TableRow.vue';
import TableHeader from './ui/table/TableHeader.vue';
const info = ref();

onMounted(async () => {
    // Fetch docker daemon info from backend
    const response = await fetch('/server/info', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    });


    info.value = await response.json();

    const responseContainers = await fetch('/server/containers', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const containersList= await responseContainers.json()
    console.log("Here contianers: \n")
    console.log(containersList)
})

</script>
<template >
  <section class =" flex justify-center   items-center self-start pt-[10rem] ">
      
     
      <div class="h-fit  "> 
      
      <h1 class="text-[3rem] font-semibold	pb-2" >Docker info</h1>
      <Table v-if="info" class="table-fixed	" theme="dark">
        <TableCaption>Docker info</TableCaption>
        <TableHeader>
        <TableRow>
          <TableHead>OS</TableHead>
          <TableHead>Docker version</TableHead>
          <TableHead>Hostname</TableHead>
          <TableHead>CPU</TableHead>
          <TableHead>Memory</TableHead>
        
        </TableRow>

      </TableHeader>

      <TableBody>
        <TableRow>
          
          <TableCell>{{info.OperatingSystem}}</TableCell>
          <TableCell>{{info.ServerVersion}}</TableCell>
          <TableCell>{{info.Name}}</TableCell>
          <TableCell>{{info.NCPU}} cores</TableCell>
          <TableCell>{{info.MemTotal}} bytes</TableCell>
        </TableRow>

      </TableBody>

      </Table>
      <span v-else>Loading...</span>
 
  </div> 
    
  </section>
</template>
