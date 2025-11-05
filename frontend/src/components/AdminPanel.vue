<template>
    <div class="admin-panel" v-if="isAdmin">
        <h3>Добавить новость</h3>
        <form @submit.prevent="addNewsItem" class="admin-form">
            <input v-model="newsForm.title" placeholder="Заголовок" required>
            <textarea v-model="newsForm.content" placeholder="Содержание" required></textarea>
            <input v-model="newsForm.image" placeholder="URL изображения">
            <button type="submit">Добавить новость</button>
        </form>
    </div>
</template>

<script setup>
import { ref } from 'vue'

const emit = defineEmits(['news-added'])
const isAdmin = ref(true) // В реальном приложении - проверка авторизации

const newsForm = ref({
    title: '',
    content: '',
    image: ''
})

const addNewsItem = () => {
    const newsItem = {
        ...newsForm.value,
        id: Date.now(),
        date: new Date().toLocaleDateString('ru-RU')
    }

    emit('news-added', newsItem)
    newsForm.value = { title: '', content: '', image: '' }
}
</script>

<style scoped>
.admin-panel {
    background: #f8fafc;
    padding: 2rem;
    border-radius: 1rem;
    margin-top: 2rem;
}

.admin-form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    max-width: 400px;
}

.admin-form input,
.admin-form textarea {
    padding: 0.75rem;
    border: 1px solid #e5e7eb;
    border-radius: 0.5rem;
    font-size: 1rem;
}

.admin-form button {
    background: #10b981;
    color: white;
    border: none;
    padding: 0.75rem;
    border-radius: 0.5rem;
    cursor: pointer;
}
</style>