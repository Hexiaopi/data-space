import { createApp } from 'vue'
import App from '@/App.vue'
// CSS common style sheet
import "@/style/common.scss"
import "@/style/element.scss"
// element css
import 'element-plus/dist/index.css'
// element plus
import ElementPlus from 'element-plus'
// element icons
import * as Icons from '@element-plus/icons-vue'
// custom directives
import directives from "@/directives/index"
// vue Router
import router from '@/router'
// vue Store
import store from '@/store'


const app = createApp(App)

// register the element Icons component
for (const [key, component] of Object.entries(Icons)) {
    app.component(key, component)
}
app.use(ElementPlus).use(directives).use(router).use(store).mount('#app')
