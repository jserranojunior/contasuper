import routes from "./mods/rotas/index";
import { createApp } from 'vue';
import App from './App.vue';
import './assets/index.css';

// Importação atualizada da Maska v2/v3
import { vMaska } from 'maska/vue';

const app = createApp(App);

// Registra a diretiva globalmente v-maska
app.directive('maska', vMaska);
app.use(routes);
app.mount('#app');