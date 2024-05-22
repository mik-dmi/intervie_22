import {createRouter, createWebHistory} from 'vue-router';
import DockerInfoView from "@/components/DockerInfoView.vue";
import Layout from  "@/components/Layout.vue";
import ContainersList from "@/components/ContainersList.vue";


const routes = [
    {
        path: '/',
        component: Layout,
        children: [
            {
                path: '/',
                name: 'home',
                component: DockerInfoView
            },
            {
                path: '/containers_list',
                name: 'containers_list',
                component: ContainersList
            }
        ]
    }
];
const router = createRouter({
    history: createWebHistory('/'),
    routes,
});
  
  export default router;