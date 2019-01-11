import Vue from 'vue'
import App from './App.vue'
import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
import bAlert from 'bootstrap-vue/es/components/alert/alert'
import bModal from 'bootstrap-vue/es/components/modal/modal';
import vBModal from 'bootstrap-vue/es/directives/modal/modal';
import bButton from 'bootstrap-vue/es/components/button/button';
import bFormInput from 'bootstrap-vue/es/components/form-input/form-input';
import bFormTextarea from 'bootstrap-vue/es/components/form-textarea/form-textarea';

import { library } from '@fortawesome/fontawesome-svg-core'
import { faCoffee } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

Vue.component('b-form-textarea', bFormTextarea);
Vue.component('b-form-input', bFormInput);
Vue.component('b-button', bButton);
Vue.directive('b-modal', vBModal);
Vue.component('b-modal', bModal);
Vue.component('b-alert', bAlert)

library.add(faCoffee)

Vue.component('font-awesome-icon', FontAwesomeIcon)

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
