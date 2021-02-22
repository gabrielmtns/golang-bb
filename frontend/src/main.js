import Vue from 'vue'
import App from './App.vue'
import('./assets/styles/index.css');
import router from './router'
import VueFormulate from '@braid/vue-formulate'
import VueToastr from 'vue-toastr'
import VModal from 'vue-js-modal'


Vue.use(VueToastr, {
  defaultPosition: "toast-bottom-right"
})
Vue.use(VueFormulate)
Vue.use(VModal)

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
