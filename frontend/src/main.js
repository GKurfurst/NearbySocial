import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from "./router/router.js";
import { createPinia } from 'pinia'
import axios from 'axios'

const app = createApp(App)
app.use(router)
app.use(ElementPlus)
app.mount('#app')
const pinia = createPinia()
app.use(pinia)