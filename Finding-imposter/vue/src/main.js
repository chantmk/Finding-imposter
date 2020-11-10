import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import log from "./store/log";
import _ from "lodash";

Vue.config.productionTip = false;

Object.defineProperty(Vue.prototype, "$lodash", { value: _ });

const ComponentContext = require.context("./", true, /\.vue$/i, "lazy");
ComponentContext.keys().forEach((componentFilePath) => {
  const componentName = componentFilePath.split("/").pop().split(".")[0];
  Vue.component(componentName, () => ComponentContext(componentFilePath));
});

new Vue({
  router,
  store: {
    store,
    log
  },
  render: (h) => h(App),
}).$mount("#app");
