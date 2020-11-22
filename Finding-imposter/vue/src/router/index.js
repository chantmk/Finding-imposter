import Vue from "vue";
import VueRouter from "vue-router";
import Index from "../views/Index.vue";
import Main from "../views/Main.vue";
import Doctor from "../views/Doctor.vue";
import Admin from "../views/Admin.vue";

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
  {
    path: "/doctor",
    component: Doctor,
  },
  {
    path: "/admin",
    component: Admin,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
