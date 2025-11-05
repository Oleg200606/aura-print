<template>
    <div class="catalog">
        <div class="catalog-header">
            <h1>Наш каталог</h1>
            <p>Примеры нашей продукции для брендов и компаний</p>
        </div>

        <!-- Админ панель для товаров -->
        <AdminProductPanel @product-added="addProduct" />

        <div class="products-grid">
            <div v-for="product in products" :key="product.id" class="product-card">
                <img :src="product.image" :alt="product.name" class="product-image">
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
import { ref } from 'vue'
import AdminProductPanel from '../components/AdminProductPanel.vue'

const products = ref([
    {
        id: 1,
        name: 'Футболка премиум',
        description: 'Качественная хлопковая футболка для вашего бренда',
        image: '/api/placeholder/300/300',
        tags: ['одежда', 'футболки']
    },
    {
        id: 2,
        name: 'Эко-сумка',
        description: 'Прочная экологичная сумка с принтом',
        image: '/api/placeholder/300/300',
        tags: ['аксессуары', 'эко']
    }
])

const addProduct = (product) => {
    products.value.push({
        ...product,
        id: Date.now()
    })
}
</script>

<style scoped>
.catalog-header {
    text-align: center;
    margin-bottom: 3rem;
}

.catalog-header h1 {
    font-size: 3rem;
    margin-bottom: 1rem;
    background: linear-gradient(135deg, #667eea, #764ba2);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}

.products-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: 2rem;
}

.product-card {
    background: white;
    border-radius: 1rem;
    overflow: hidden;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
    transition: all 0.3s;
}

.product-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.15);
}

.product-image {
    width: 100%;
    height: 250px;
    object-fit: cover;
}

.product-info {
    padding: 1.5rem;
}

.product-info h3 {
    margin-bottom: 0.5rem;
    color: #2d3748;
}

.product-info p {
    color: #6b7280;
    margin-bottom: 1rem;
}

.product-tags {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;
}

.tag {
    background: #e5e7eb;
    color: #4b5563;
    padding: 0.25rem 0.75rem;
    border-radius: 1rem;
    font-size: 0.875rem;
}
</style>