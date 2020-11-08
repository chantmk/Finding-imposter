import Vue from "vue";
import VueRouter from "vue-router";
import Index from "../views/Index.vue";
import Main from "../views/Main.vue";

Vue.use(VueRouter);

const routes = [
    {
        path: "/main",
        component: Main,
    },
    {
        path: "/",
        component: Index,
    },
  
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
