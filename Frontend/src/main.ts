import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia';
import App from './App.vue'
import router from './router'

import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import { faUser, faPlus, faBell, faBellSlash} from '@fortawesome/free-solid-svg-icons'

library.add(faUser, faPlus, faBell, faBellSlash)

const vuetify = createVuetify({
  components,
  directives,
})

const app = createApp(App)
const pinia = createPinia();
app.component('font-awesome-icon', FontAwesomeIcon)
app.use(vuetify)
app.use(pinia);
app.use(router)

app.mount('#app')
