import './src/style.css';
import {createApp} from 'vue';
import App from './components/App.vue';
import router from './router'



const app = createApp(App).use(router).mount('#app');


