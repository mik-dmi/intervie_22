<template>
  <footer class="pt-[6rem] flex flex-col   bg-[#0a192b]">
      <div class="mx-auto flex  justify-center align-middle px-4 pb-[4rem]">
          <div class="flex flex-col justify-center align-middle pr-10">
              <h2 class="text-white text-4xl font-bold">Get in Touch</h2>
              <div class = "text-white py-[0.7rem]   flex flex-col"> 
                  <span>Full Stack Interview Challenge to join Kinexon in Munich</span>
                  <span>It was developed by Miguel Caridade</span>
              </div>
              <ul class="flex justify-start align-middle gap-4">
                  <li class="group list-item border-white  cursor-pointer border-2 rounded-full p-2 hover:scale-105 ease-in-out duration-75 hover:border-secondary">
                      
                      <a href="https://www.example.com" target="_blank">
                          <Linkedin color="#ffffff" class="group-hover:stroke-secondary	ease-in-out duration-75" />
                      </a>
                  </li>
                  <li class="group list-item cursor-pointer border-white border-2 rounded-full p-2 hover:scale-105 ease-in-out duration-75 hover:border-secondary">
                      <a href="https://www.example.com" target="_blank">
                          <Github color="#ffffff" class="group-hover:stroke-secondary ease-in-out duration-75	"/>
                      </a>
                  </li>
                  <li class="group list-item cursor-pointer border-white border-2 rounded-full p-2 hover:scale-105 ease-in-out duration-75 hover:border-secondary">
                      <a href="https://www.example.com" target="_blank">
                          <Mail color="#ffffff" class="group-hover:stroke-secondary ease-in-out duration-75	"/>
                      </a>
                  </li>
                  <li class="group list-item cursor-pointer border-white border-2 rounded-full p-2 hover:scale-105 ease-in-out duration-100 hover:border-secondary"> 
                      <a href="https://www.example.com" target="_blank">   
                          <BriefcaseBusiness color="#ffffff" class="group-hover:stroke-secondary	ease-in-out duration-100" />
                      </a>
                  </li>          
              </ul>       
          </div >
          <div class="flex  gap-10 pl-10 ">
              <div>
                  <div class="rounded-tl-2xl  rounded-br-2xl w-[13rem] h-[13rem] absolute	bg-secondary bg-opacity-[0.04] " />
                  <div class="rounded-tl-2xl hover:translate-x-[-4px] hover:translate-y-[-4px]  hover:bg-white ease-in-out duration-300 hover:scale-[1.02] flex flex-col group cursor-pointer gap-6 justify-center text-white align-middle rounded-br-2xl w-[13rem] h-[13rem] relative bottom-4 right-4 	bg-[#17262B]">
                  
                      <Mail color="#ffffff" size= "32" class="group-hover:stroke-secondary   group-hover:scale-110  ease-in-out duration-75	self-center"/>
                      <span class="text-[0.9rem] text-center group-hover:text-secondary group-hover:scale-110 ease-in-out duratio cursor-pointer	">miguel.caridade@gmail</span>
                   
                  </div>
              </div>
              <div>
                  <div class="rounded-tl-2xl rounded-br-2xl w-[13rem] h-[13rem] absolute	bg-secondary bg-opacity-[0.04] "/>
                  <div class="rounded-tl-2xl hover:translate-x-[-4px] hover:translate-y-[-4px]  hover:bg-white ease-in-out duration-300 hover:scale-[1.02] flex flex-col group cursor-pointer gap-6 justify-center text-white align-middle rounded-br-2xl w-[13rem] h-[13rem] relative bottom-4 right-4 	bg-[#17262B]">
                    <svg class=" ease-in-out   group-hover:scale-110 duration-75 self-center" width="25" height="25" viewBox="1362 3.95 328 357.52" fill="none" xmlns="http://www.w3.org/2000/svg" style="max-width: 100%; vertical-align: bottom; font-weight: bolder;">
                      <path fill-rule="evenodd" class="fill-white group-hover:fill-secondary ease-in-out duration-75" clip-rule="evenodd" d="M1555.89 198.121H1497.23L1362.29 361.473H1414.49C1420.66 361.473 1426.44 358.776 1430.48 354.153L1511.87 259.956C1515.53 255.622 1518.61 251.48 1521.02 247.531C1523.62 243.293 1525.45 238.574 1526.51 233.758C1527.76 238.959 1529.59 243.582 1532 247.531C1534.41 251.48 1537.49 255.622 1541.15 259.956L1622.54 354.153C1626.58 358.776 1632.36 361.473 1638.52 361.473H1690.73L1555.79 198.121"></path>
                      <path fill-rule="evenodd" class="fill-white group-hover:fill-secondary ease-in-out duration-75" clip-rule="evenodd" d="M1633.04 3.94897C1626.87 3.94897 1621.09 6.64582 1617.05 11.269L1541.25 99.0127C1537.59 103.347 1534.5 107.488 1532.1 111.437C1529.5 115.675 1527.67 120.395 1526.61 125.211C1525.55 120.298 1523.62 115.675 1521.12 111.437C1518.32 107.103 1515.34 102.962 1511.97 99.0127L1435.01 9.92055C1431.74 6.16424 1427.02 3.94897 1422.01 3.94897H1367.78L1497.42 160.847H1556.08L1685.53 3.94897H1633.04Z"></path>
                    </svg>
                    <span class="text-[0.9rem] text-center group-hover:text-secondary group-hover:scale-110 ease-in-out duration-75 cursor-pointer	">www.kinexon.com</span>
                     
                  </div>
              </div>
          </div>

      </div>
      <ul class=" text-white mx-auto w-full  sm:px-10 lg:max-w-7xl justify-center flex sm:gap-10 gap-5 tracking-wider sm:text-[1rem] text-[0.8rem] pb-2 pt-[0.25] border-t-[1.5px] mt-[4px] 	border-[#363f47]">
        <router-link :to="{ name: 'home' }">
          <span class=" hover:underline underline-offset-4 font-medium duration-75 hover:text-secondary">Home</span>
        </router-link>
        <router-link :to="{ name: 'containers_list' }">
          <span class=" hover:underline underline-offset-4 font-medium duration-75 hover:text-secondary">Container</span>
        </router-link>
        <router-link :to="{ name: 'home' }">
          <span class=" hover:underline underline-offset-4 font-medium duration-75 hover:text-secondary">About</span>
        </router-link>
      </ul>

      
  </footer>
</template>


<script setup>
                  
import { Instagram , Linkedin, Github  , Mail ,BriefcaseBusiness } from 'lucide-vue-next';


</script>

