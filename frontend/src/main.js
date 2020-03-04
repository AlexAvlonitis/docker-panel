import '@coreui/coreui/dist/css/coreui.min.css'
import '@coreui/coreui/dist/js/coreui.bundle.min.js'
import CoreuiVue from '@coreui/vue';
import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false
Vue.config.performance = true
Vue.use(CoreuiVue)

new Vue({
  render: h => h(App)
}).$mount('#app')
