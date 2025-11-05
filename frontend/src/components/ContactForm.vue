<template>
    <form @submit.prevent="submitForm" class="contact-form">
        <div class="form-group">
            <input v-model="form.name" placeholder="Ваше имя" required>
        </div>
        <div class="form-group">
            <input v-model="form.email" type="email" placeholder="Email" required>
        </div>
        <div class="form-group">
            <input v-model="form.phone" placeholder="Телефон">
        </div>
        <div class="form-group">
            <textarea v-model="form.message" placeholder="Сообщение" rows="5" required></textarea>
        </div>
        <button type="submit" class="submit-button">Отправить сообщение</button>
    </form>
</template>

<script setup>
import { ref } from 'vue'

const form = ref({
    name: '',
    email: '',
    phone: '',
    message: ''
})

const submitForm = async () => {
    // Здесь будет отправка на бэкенд
    console.log('Форма отправлена:', form.value)
    alert('Сообщение отправлено!')
    form.value = { name: '', email: '', phone: '', message: '' }
}
</script>

<style scoped>
.contact-form {
    max-width: 500px;
    margin: 0 auto;
}

.form-group {
    margin-bottom: 1.5rem;
}

.form-group input,
.form-group textarea {
    width: 100%;
    padding: 1rem;
    border: 1px solid #e5e7eb;
    border-radius: 0.5rem;
    font-size: 1rem;
    transition: border-color 0.3s;
}

.form-group input:focus,
.form-group textarea:focus {
    outline: none;
    border-color: #667eea;
}

.submit-button {
    width: 100%;
    background: linear-gradient(135deg, #667eea, #764ba2);
    color: white;
    border: none;
    padding: 1rem;
    border-radius: 0.5rem;
    font-size: 1.1rem;
    cursor: pointer;
    transition: transform 0.3s;
}

.submit-button:hover {
    transform: translateY(-2px);
}
</style>