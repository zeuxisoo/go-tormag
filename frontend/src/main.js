import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import App from './App.vue'
import router from './router'
import './registerServiceWorker'

Vue.use(VueAxios, axios)

Vue.filter('defaultEmpty', value => {
    return value.trim().length <= 0 ? "---" : value;
});

Vue.config.productionTip = false;

new Vue({
    router,
    render: h => h(App)
}).$mount('#app');
