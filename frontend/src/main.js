import '@coreui/coreui/dist/css/coreui.min.css'
import '@coreui/coreui/dist/js/coreui.bundle.min.js'
import CoreuiVue from '@coreui/vue';
import Vue from 'vue'
import App from './App.vue'
import { iconsSet as icons } from './assets/icons/icons.js'
import store from './store'
import router from './router'

Vue.config.productionTip = false
Vue.config.performance = true
Vue.use(CoreuiVue)

new Vue({
  el: '#app',
  store,
  router,
  icons,
  template: '<App/>',
  components: {
    App
  }
})
