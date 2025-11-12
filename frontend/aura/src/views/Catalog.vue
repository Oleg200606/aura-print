<template>
    <div class="catalog">
        <div class="container">
            <h1 class="text-center mb-4">Наша продукция</h1>

            <!-- Admin Controls -->
            <div v-if="$store.state.isAdmin" class="admin-controls mb-4">
                <button @click="showProductForm = true" class="btn btn-primary">
                    Добавить товар
                </button>
            </div>

            <!-- Product Form Modal -->
            <div v-if="showProductForm" class="modal">
                <div class="modal-content">
                    <h3>Добавить новый товар</h3>
                    <form @submit.prevent="addProduct">
                        <div class="form-group">
                            <label>Название:</label>
                            <input v-model="newProduct.name" type="text" required>
                        </div>
                        <div class="form-group">
                            <label>Описание:</label>
                            <textarea v-model="newProduct.description" required></textarea>
                        </div>
                        <div class="form-group">
                            <label>Ссылка на изображение:</label>
                            <input v-model="newProduct.image" type="url" required>
                        </div>
                        <div class="form-group">
                            <label>Категория:</label>
                            <input v-model="newProduct.category" type="text" required>
                        </div>
                        <div class="form-actions">
                            <button type="button" @click="showProductForm = false" class="btn">Отмена</button>
                            <button type="submit" class="btn btn-primary">Добавить</button>
                        </div>
                    </form>
                </div>
            </div>

            <!-- Products Grid -->
            <div class="grid grid-3">
                <div v-for="product in products" :key="product.id" class="product-card card">
                    <img :src="product.image_url" :alt="product.name" class="product-image">
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

        const newProduct = ref({
            name: '',
            description: '',
            image: '',
            category: ''
        })

        onMounted(() => {
            store.dispatch('fetchProducts')
        })

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
                    newProduct.value = { name: '', description: '', image: '', category: '' }
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
            newProduct,
            addProduct,
            deleteProduct
        }
    }
}
</script>

<style scoped>
.product-card {
    display: flex;
    flex-direction: column;
}

.product-image {
    width: 100%;
    height: 250px;
    object-fit: cover;
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
}

.product-category {
    background: #f8f9fa;
    padding: 0.25rem 0.5rem;
    border-radius: 15px;
    font-size: 0.8rem;
    color: #6c757d;
    margin-top: auto;
    margin-bottom: 1rem;
    align-self: flex-start;
}

.btn-sm {
    padding: 0.5rem 1rem;
    font-size: 0.9rem;
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
    border-radius: 10px;
    width: 90%;
    max-width: 500px;
}

.form-group {
    margin-bottom: 1rem;
}

.form-group label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: bold;
}

.form-group input,
.form-group textarea {
    width: 100%;
    padding: 0.5rem;
    border: 1px solid #ddd;
    border-radius: 5px;
    font-size: 1rem;
}

.form-group textarea {
    height: 100px;
    resize: vertical;
}

.form-actions {
    display: flex;
    gap: 1rem;
    justify-content: flex-end;
    margin-top: 1rem;
}

.empty-state {
    text-align: center;
    padding: 4rem 2rem;
    color: #6c757d;
}
</style>