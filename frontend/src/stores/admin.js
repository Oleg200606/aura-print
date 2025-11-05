import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAdminStore = defineStore('admin', () => {
    const products = ref([])
    const news = ref([])

    const addProduct = (product) => {
        products.value.push(product)
    }

    const addNews = (newsItem) => {
        news.value.unshift(newsItem)
    }

    return {
        products,
        news,
        addProduct,
        addNews
    }
})