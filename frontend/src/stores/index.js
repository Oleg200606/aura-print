import { createStore } from "vuex";

import { DEFAULT_IMAGE_URL } from "@/data";
export default createStore({
  state: {
    isAdmin: localStorage.getItem("isAdmin") === "true" || false,
    products: [],
    news: [],
  },
  mutations: {
    SET_ADMIN(state, status) {
      state.isAdmin = status;
      localStorage.setItem("isAdmin", status);
    },
    SET_PRODUCTS(state, products) {
      state.products = products;
    },
    SET_NEWS(state, news) {
      state.news = news;
    },
    ADD_PRODUCT(state, product) {
      state.products.push(product);
    },
    ADD_NEWS(state, newsItem) {
      state.news.unshift(newsItem);
    },
    DELETE_PRODUCT(state, productId) {
      state.products = state.products.filter((p) => p.id !== productId);
    },
    DELETE_NEWS(state, newsId) {
      state.news = state.news.filter((n) => n.id !== newsId);
    },
  },
  actions: {
    async login({ commit }, credentials) {
      try {
        console.log("📤 Sending login request to /api/admin/login");
        const response = await fetch("/api/admin/login", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(credentials),
        });
        if (!response.ok) {
          console.error("📥 Login response status:", response.status, response.statusText);
          return {
            success: false,
            error: "response status is " + response.status + " " + response.statusText,
          };
        } else {
          console.log("📥 Login response status:", response.status);
        }
        const data = await response.json();

        console.log("📥 Login response data:", data);

        if (data.success) {
          commit("SET_ADMIN", true);
          return { success: true };
        } else {
          return { success: false, error: data.error || "Invalid credentials" };
        }
      } catch (error) {
        console.error("❌ Login request failed:", error);
        return { success: false, error: "login failed: " + error.message };
      }
    },

    async fetchProducts({ commit }) {
      try {
        const response = await fetch("/api/products");
        if (!response.ok) throw new Error("Network response was not ok");
        const products = await response.json();
        commit("SET_PRODUCTS", products);
      } catch (error) {
        console.error("Failed to fetch products:", error);
      }
    },

    async fetchNews({ state, commit }) {
      try {
        const response = await fetch("/api/news");
        if (!response.ok) throw new Error("Network response was not ok");
        const news = await response.json();
        for (const newsItem of news) {
          newsItem.imageUrl = () => {
            const backHost = import.meta.env.BACKEND_HOST;
            if (!newsItem.image_url || !backHost) {
              return DEFAULT_IMAGE_URL;
            }

            return backHost + newsItem.image_url;
          };
        }
        state.news = news;
      } catch (error) {
        console.error("Failed to fetch news:", error);
      }
    },
  },
});
