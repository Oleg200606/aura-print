<template>
    <div class="admin-product-panel" v-if="isAdmin">
        <h3>Добавить товар</h3>
        <form @submit.prevent="addProductItem" class="admin-form">
            <input v-model="productForm.name" placeholder="Название товара" required>
            <textarea v-model="productForm.description" placeholder="Описание" required></textarea>
            <input v-model="productForm.image" placeholder="URL изображения" required>
            <input v-model="productForm.tags" placeholder="Теги (через запятую)">
            <button type="submit">Добавить товар</button>
        </form>
    </div>
</template>

<script setup>
import { ref } from 'vue'

const emit = defineEmits(['product-added'])
const isAdmin = ref(true)

const productForm = ref({
    name: '',
    description: '',
    image: '',
    tags: ''
})

const addProductItem = () => {
    const product = {
        ...productForm.value,
        tags: productForm.value.tags.split(',').map(tag => tag.trim())
    }

    emit('product-added', product)
    productForm.value = { name: '', description: '', image: '', tags: '' }
}
</script>