<template>
    <div class="admin-page">
        <div class="admin-header">
            <h1>Панель управления</h1>
            <button @click="handleLogout" class="logout-button">Выйти</button>
        </div>

        <div class="admin-content">
            <!-- Управление товарами -->
            <section class="admin-section">
                <h2>Управление товарами</h2>

                <!-- Форма добавления товара -->
                <div class="admin-card">
                    <h3>Добавить товар</h3>
                    <form @submit.prevent="addProduct" class="admin-form">
                        <div class="form-row">
                            <input v-model="productForm.name" placeholder="Название товара" required>
                            <input v-model="productForm.image_url" placeholder="URL изображения" required>
                        </div>
                        <textarea v-model="productForm.description" placeholder="Описание товара" required></textarea>
                        <input v-model="productForm.tagsInput" placeholder="Теги (через запятую)">
                        <button type="submit" :disabled="productLoading">
                            {{ productLoading ? 'Добавление...' : 'Добавить товар' }}
                        </button>
                    </form>
                </div>

                <!-- Список товаров -->
                <div class="admin-card">
                    <h3>Список товаров</h3>
                    <div class="items-list">
                        <div v-for="product in products" :key="product.id" class="item-card">
                            <img :src="product.image_url" :alt="product.name" class="item-image">
                            <div class="item-info">
                                <h4>{{ product.name }}</h4>
                                <p>{{ product.description }}</p>
                                <div class="item-tags">
                                    <span v-for="tag in product.tags" :key="tag" class="tag">{{ tag }}</span>
                                </div>
                            </div>
                            <div class="item-actions">
                                <button @click="editProduct(product)" class="edit-button">Редактировать</button>
                                <button @click="deleteProduct(product.id)" class="delete-button">Удалить</button>
                            </div>
                        </div>
                    </div>
                </div>
            </section>

            <!-- Управление новостями -->
            <section class="admin-section">
                <h2>Управление новостями</h2>

                <!-- Форма добавления новости -->
                <div class="admin-card">
                    <h3>Добавить новость</h3>
                    <form @submit.prevent="addNews" class="admin-form">
                        <input v-model="newsForm.title" placeholder="Заголовок новости" required>
                        <textarea v-model="newsForm.content" placeholder="Содержание новости" required
                            rows="4"></textarea>
                        <input v-model="newsForm.image_url" placeholder="URL изображения">
                        <button type="submit" :disabled="newsLoading">
                            {{ newsLoading ? 'Добавление...' : 'Добавить новость' }}
                        </button>
                    </form>
                </div>

                <!-- Список новостей -->
                <div class="admin-card">
                    <h3>Список новостей</h3>
                    <div class="items-list">
                        <div v-for="newsItem in news" :key="newsItem.id" class="item-card">
                            <img v-if="newsItem.image_url" :src="newsItem.image_url" :alt="newsItem.title"
                                class="item-image">
                            <div class="item-info">
                                <h4>{{ newsItem.title }}</h4>
                                <p>{{ newsItem.content }}</p>
                                <span class="item-date">{{ formatDate(newsItem.created_at) }}</span>
                            </div>
                            <div class="item-actions">
                                <button @click="editNews(newsItem)" class="edit-button">Редактировать</button>
                                <button @click="deleteNews(newsItem.id)" class="delete-button">Удалить</button>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { apiService } from '@/services/api'

const router = useRouter()
const authStore = useAuthStore()

// Проверка аутентификации
if (!authStore.isAuthenticated) {
    router.push('/login')
}

const products = ref([])
const news = ref([])

const productForm = ref({
    name: '',
    description: '',
    image_url: '',
    tagsInput: ''
})

const newsForm = ref({
    title: '',
    content: '',
    image_url: ''
})

const productLoading = ref(false)
const newsLoading = ref(false)

// Загрузка данных
const loadData = async () => {
    try {
        products.value = await apiService.getProducts()
        news.value = await apiService.getNews()
    } catch (error) {
        console.error('Error loading data:', error)
    }
}

// Работа с товарами
const addProduct = async () => {
    productLoading.value = true
    try {
        const productData = {
            ...productForm.value,
            tags: productForm.value.tagsInput.split(',').map(tag => tag.trim()).filter(tag => tag)
        }

        await apiService.createProduct(productData)
        await loadData()
        productForm.value = { name: '', description: '', image_url: '', tagsInput: '' }
    } catch (error) {
        console.error('Error adding product:', error)
    } finally {
        productLoading.value = false
    }
}

const deleteProduct = async (id) => {
    if (!confirm('Вы уверены, что хотите удалить этот товар?')) return

    try {
        await apiService.deleteProduct(id)
        await loadData()
    } catch (error) {
        console.error('Error deleting product:', error)
    }
}

// Работа с новостями
const addNews = async () => {
    newsLoading.value = true
    try {
        await apiService.createNews(newsForm.value)
        await loadData()
        newsForm.value = { title: '', content: '', image_url: '' }
    } catch (error) {
        console.error('Error adding news:', error)
    } finally {
        newsLoading.value = false
    }
}

const deleteNews = async (id) => {
    if (!confirm('Вы уверены, что хотите удалить эту новость?')) return

    try {
        await apiService.deleteNews(id)
        await loadData()
    } catch (error) {
        console.error('Error deleting news:', error)
    }
}

const editProduct = (product) => {
    // Реализация редактирования
    console.log('Edit product:', product)
}

const editNews = (newsItem) => {
    // Реализация редактирования
    console.log('Edit news:', newsItem)
}

const handleLogout = () => {
    authStore.logout()
}

const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString('ru-RU')
}

onMounted(() => {
    loadData()
})
</script>

<style scoped>
.admin-page {
    min-height: 100vh;
    background: #f8fafc;
    padding: 2rem;
}

.admin-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 3rem;
    padding-bottom: 1rem;
    border-bottom: 1px solid #e5e7eb;
}

.admin-header h1 {
    color: #2d3748;
    font-size: 2.5rem;
}

.logout-button {
    background: #ef4444;
    color: white;
    border: none;
    padding: 0.75rem 1.5rem;
    border-radius: 0.5rem;
    cursor: pointer;
    font-weight: 500;
}

.logout-button:hover {
    background: #dc2626;
}

.admin-content {
    display: grid;
    gap: 3rem;
    max-width: 1200px;
    margin: 0 auto;
}

.admin-section h2 {
    color: #2d3748;
    margin-bottom: 1.5rem;
    font-size: 1.8rem;
}

.admin-card {
    background: white;
    padding: 2rem;
    border-radius: 1rem;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
    margin-bottom: 2rem;
}

.admin-card h3 {
    color: #374151;
    margin-bottom: 1.5rem;
    font-size: 1.3rem;
}

.admin-form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

.form-row {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 1rem;
}

.admin-form input,
.admin-form textarea {
    padding: 0.75rem;
    border: 1px solid #d1d5db;
    border-radius: 0.5rem;
    font-size: 1rem;
}

.admin-form button {
    background: linear-gradient(135deg, #667eea, #764ba2);
    color: white;
    border: none;
    padding: 0.75rem;
    border-radius: 0.5rem;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
}

.admin-form button:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}

.items-list {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

.item-card {
    display: grid;
    grid-template-columns: 100px 1fr auto;
    gap: 1rem;
    padding: 1rem;
    border: 1px solid #e5e7eb;
    border-radius: 0.5rem;
    align-items: start;
}

.item-image {
    width: 100px;
    height: 100px;
    object-fit: cover;
    border-radius: 0.5rem;
}

.item-info {
    flex: 1;
}

.item-info h4 {
    margin-bottom: 0.5rem;
    color: #2d3748;
}

.item-info p {
    color: #6b7280;
    margin-bottom: 0.5rem;
}

.item-tags {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;
}

.tag {
    background: #e5e7eb;
    color: #4b5563;
    padding: 0.25rem 0.5rem;
    border-radius: 1rem;
    font-size: 0.75rem;
}

.item-date {
    color: #9ca3af;
    font-size: 0.875rem;
}

.item-actions {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
}

.edit-button {
    background: #3b82f6;
    color: white;
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 0.25rem;
    cursor: pointer;
    font-size: 0.875rem;
}

.delete-button {
    background: #ef4444;
    color: white;
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 0.25rem;
    cursor: pointer;
    font-size: 0.875rem;
}
</style>