<template>
    <div class="catalog">
        <div class="container">
            <h1 class="text-center mb-4">Наша продукция</h1>

            <!-- Admin Controls -->
            <div v-if="$store.state.isAdmin" class="admin-controls mb-4">
                <button @click="showProductForm = true" class="btn btn-primary">
                    Добавить товар
                </button>

                <!-- Upload Image Button -->
                <button @click="showUploadForm = true" class="btn btn-secondary">
                    Загрузить изображение
                </button>
            </div>

            <!-- Image Upload Modal -->
            <div v-if="showUploadForm" class="modal">
                <div class="modal-content">
                    <h3>Загрузить изображение</h3>
                    <form @submit.prevent="uploadImage">
                        <div class="form-group">
                            <label>Выберите изображение:</label>
                            <input type="file" @change="handleFileSelect" accept="image/*" required>
                        </div>
                        <div class="form-group">
                            <label>Папка:</label>
                            <select v-model="uploadFolder" class="form-input">
                                <option value="products">Товары</option>
                                <option value="news">Новости</option>
                                <option value="images">Общие</option>
                            </select>
                        </div>
                        <div class="form-actions">
                            <button type="button" @click="showUploadForm = false" class="btn">Отмена</button>
                            <button type="submit" class="btn btn-primary" :disabled="uploading">
                                {{ uploading ? 'Загрузка...' : 'Загрузить' }}
                            </button>
                        </div>
                    </form>
                </div>
            </div>

            <!-- Product Form Modal -->
            <div v-if="showProductForm" class="modal">
                <div class="modal-content">
                    <h3>Добавить новый товар</h3>
                    <form @submit.prevent="addProduct">
                        <div class="form-group">
                            <label>Название:</label>
                            <input v-model="newProduct.name" type="text" required class="form-input">
                        </div>
                        <div class="form-group">
                            <label>Описание:</label>
                            <textarea v-model="newProduct.description" required class="form-input"></textarea>
                        </div>
                        <div class="form-group">
                            <label>Ссылка на изображение:</label>
                            <input v-model="newProduct.image_url" type="url" required class="form-input"
                                placeholder="/api/images/filename.jpg">
                            <small>Используйте загрузчик выше или укажите путь</small>
                        </div>
                        <div class="form-group">
                            <label>Категория:</label>
                            <input v-model="newProduct.category" type="text" required class="form-input">
                        </div>
                        <div class="form-actions">
                            <button type="button" @click="showProductForm = false" class="btn">Отмена</button>
                            <button type="submit" class="btn btn-primary">Добавить</button>
                        </div>
                    </form>
                </div>
            </div>

            <!-- Products Grid -->
            <div class="products-grid">
                <div v-for="product in products" :key="product.id" class="product-card card">
                    <div class="product-image-container">
                        <img :src="getImageUrl(product.image_url)" :alt="product.name" class="product-image"
                            @error="handleImageError">
                    </div>
                    <div class="product-content">
                        <h3>{{ product.name }}</h3>
                        <p>{{ product.description }}</p>
                        <div class="product-category">{{ product.category }}</div>
                        <button v-if="$store.state.isAdmin" @click="deleteProduct(product.id)"
                            class="btn btn-danger btn-sm">
                            Удалить
                        </button>
                    </div>
                </div>
            </div>

            <!-- Empty State -->
            <div v-if="products.length === 0" class="empty-state">
                <h3>Пока нет товаров</h3>
                <p>Администратор может добавить товары через панель управления</p>
            </div>
        </div>
    </div>
</template>

<script>
import { computed, ref, onMounted } from 'vue'
import { useStore } from 'vuex'

export default {
    name: 'Catalog',
    setup() {
        const store = useStore()
        const products = computed(() => store.state.products)
        const showProductForm = ref(false)
        const showUploadForm = ref(false)
        const uploading = ref(false)
        const selectedFile = ref(null)
        const uploadFolder = ref('products')

        const newProduct = ref({
            name: '',
            description: '',
            image_url: '',
            category: ''
        })

        onMounted(() => {
            store.dispatch('fetchProducts')
        })

        const getImageUrl = (imagePath) => {
            if (!imagePath) return '/api/images/default-product.jpg'
            if (imagePath.startsWith('http')) return imagePath
            return `/api${imagePath}`
        }

        const handleImageError = (event) => {
            event.target.src = '/api/images/default-product.jpg'
        }

        const handleFileSelect = (event) => {
            selectedFile.value = event.target.files[0]
        }

        const uploadImage = async () => {
            if (!selectedFile.value) return

            uploading.value = true
            const formData = new FormData()
            formData.append('image', selectedFile.value)
            formData.append('folder', uploadFolder.value)

            try {
                const response = await fetch('http://localhost:8081/api/admin/upload/image', {
                    method: 'POST',
                    body: formData
                })

                const result = await response.json()

                if (result.success) {
                    alert(`Изображение загружено! URL: ${result.data.url}`)
                    // Автоматически подставляем URL в форму товара
                    newProduct.value.image_url = result.data.url
                    showUploadForm.value = false
                    showProductForm.value = true
                } else {
                    alert('Ошибка загрузки: ' + result.message)
                }
            } catch (error) {
                console.error('Upload failed:', error)
                alert('Ошибка загрузки изображения')
            } finally {
                uploading.value = false
            }
        }

        const addProduct = async () => {
            try {
                const response = await fetch('http://localhost:8081/api/admin/products', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(newProduct.value)
                })

                if (response.ok) {
                    const product = await response.json()
                    store.commit('ADD_PRODUCT', product)
                    showProductForm.value = false
                    newProduct.value = { name: '', description: '', image_url: '', category: '' }
                }
            } catch (error) {
                console.error('Failed to add product:', error)
            }
        }

        const deleteProduct = async (productId) => {
            if (confirm('Вы уверены, что хотите удалить этот товар?')) {
                try {
                    const response = await fetch(`http://localhost:8081/api/admin/products/${productId}`, {
                        method: 'DELETE'
                    })

                    if (response.ok) {
                        store.commit('DELETE_PRODUCT', productId)
                    }
                } catch (error) {
                    console.error('Failed to delete product:', error)
                }
            }
        }

        return {
            products,
            showProductForm,
            showUploadForm,
            uploading,
            selectedFile,
            uploadFolder,
            newProduct,
            getImageUrl,
            handleImageError,
            handleFileSelect,
            uploadImage,
            addProduct,
            deleteProduct
        }
    }
}
</script>

<style scoped>
.products-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 2rem;
}

.product-card {
    display: flex;
    flex-direction: column;
    height: 100%;
    transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.product-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.15);
}

.product-image-container {
    width: 100%;
    height: 250px;
    overflow: hidden;
    background: #f8f9fa;
}

.product-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform 0.3s ease;
}

.product-card:hover .product-image {
    transform: scale(1.05);
}

.product-content {
    padding: 1.5rem;
    flex-grow: 1;
    display: flex;
    flex-direction: column;
}

.product-content h3 {
    margin-bottom: 1rem;
    color: #2c3e50;
    font-size: 1.25rem;
}

.product-content p {
    flex-grow: 1;
    margin-bottom: 1rem;
    color: #666;
    line-height: 1.5;
}

.product-category {
    background: linear-gradient(135deg, #667eea, #764ba2);
    color: white;
    padding: 0.5rem 1rem;
    border-radius: 20px;
    font-size: 0.8rem;
    font-weight: 500;
    margin-bottom: 1rem;
    align-self: flex-start;
}

.btn-sm {
    padding: 0.5rem 1rem;
    font-size: 0.9rem;
}

.btn-secondary {
    background: #6c757d;
    color: white;
    margin-left: 1rem;
}

.btn-secondary:hover {
    background: #5a6268;
}

.admin-controls {
    display: flex;
    justify-content: center;
    gap: 1rem;
    flex-wrap: wrap;
}

.modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
}

.modal-content {
    background: white;
    padding: 2rem;
    border-radius: 15px;
    width: 90%;
    max-width: 500px;
    max-height: 90vh;
    overflow-y: auto;
}

.form-group {
    margin-bottom: 1.5rem;
}

.form-group label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: 600;
    color: #2c3e50;
}

.form-input {
    width: 100%;
    padding: 0.75rem;
    border: 2px solid #e9ecef;
    border-radius: 8px;
    font-size: 1rem;
    transition: border-color 0.3s ease;
}

.form-input:focus {
    outline: none;
    border-color: #667eea;
    box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.form-actions {
    display: flex;
    gap: 1rem;
    justify-content: flex-end;
    margin-top: 2rem;
}

.empty-state {
    text-align: center;
    padding: 4rem 2rem;
    color: #6c757d;
}

.empty-state h3 {
    margin-bottom: 1rem;
    color: #495057;
}

/* Адаптивность */
@media (max-width: 768px) {
    .products-grid {
        grid-template-columns: 1fr;
    }

    .admin-controls {
        flex-direction: column;
        align-items: center;
    }

    .btn-secondary {
        margin-left: 0;
        margin-top: 1rem;
    }

    .modal-content {
        margin: 1rem;
        padding: 1.5rem;
    }
}
</style>