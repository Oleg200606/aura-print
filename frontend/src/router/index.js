import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import("@/views/HomeView.vue"),
  },
  {
    path: "/catalog",
    name: "Catalog",
    component: () => import("@/views/CatalogView.vue"),
  },
  {
    path: "/admin",
    name: "Admin",
    component: () => import("@/views/AdminView.vue"),
  },
  {
    path: "/about",
    name: "About",
    component: () => import("@/views/AboutView.vue"),
  },
  {
    path: "/contact",
    name: "Contact",
    component: () => import("@/views/ContactView.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
