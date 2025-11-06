<template>
    <div class="catalog">
        <div class="catalog-header">
            <h1>Наш каталог</h1>
            <p>Примеры нашей продукции для брендов и компаний</p>
        </div>

        <div class="products-grid">
            <div v-for="product in products" :key="product.id" class="product-card">
                <img :src="product.image_url" :alt="product.name" class="product-image">
                <div class="product-info">
                    <h3>{{ product.name }}</h3>
                    <p>{{ product.description }}</p>
                    <div class="product-tags">
                        <span v-for="tag in product.tags" :key="tag" class="tag">{{ tag }}</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { apiService } from '@/services/api'

const products = ref([])

const loadProducts = async () => {
    try {
        products.value = await apiService.getProducts()
    } catch (error) {
        console.error('Error loading products:', error)
    }
}

onMounted(() => {
    loadProducts()
})
</script>