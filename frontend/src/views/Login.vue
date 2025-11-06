<template>
    <div class="login-page">
        <div class="login-container">
            <div class="login-card">
                <div class="login-header">
                    <h1>Aura Print Admin</h1>
                    <p>Войдите в панель управления</p>
                </div>

                <form @submit.prevent="handleLogin" class="login-form">
                    <div class="form-group">
                        <label for="username">Логин</label>
                        <input id="username" v-model="loginForm.username" type="text" required
                            placeholder="Введите логин" :disabled="loading">
                    </div>

                    <div class="form-group">
                        <label for="password">Пароль</label>
                        <input id="password" v-model="loginForm.password" type="password" required
                            placeholder="Введите пароль" :disabled="loading">
                    </div>

                    <button type="submit" class="login-button" :disabled="loading">
                        <span v-if="loading">Вход...</span>
                        <span v-else>Войти</span>
                    </button>

                    <div v-if="error" class="error-message">
                        {{ error }}
                    </div>

                    <!-- Добавьте для отладки -->
                    <div class="debug-info">
                        <p><small>Запрос отправляется на: {{ apiUrl }}</small></p>
                        <p><small>Тестовые данные: admin / admin123</small></p>
                    </div>
                </form>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const loginForm = ref({
    username: 'admin',
    password: 'admin123'
})

const loading = ref(false)
const error = ref('')

const apiUrl = computed(() => {
    return `${window.location.origin}/api/login`
})

const handleLogin = async () => {
    loading.value = true
    error.value = ''

    try {
        await authStore.login(loginForm.value.username, loginForm.value.password)
        router.push('/admin')
    } catch (err) {
        error.value = err.message
        console.error('Login error details:', err)
    } finally {
        loading.value = false
    }
}
</script>

<style scoped>
.debug-info {
    margin-top: 1rem;
    padding: 0.75rem;
    background: #f3f4f6;
    border-radius: 0.5rem;
    font-size: 0.75rem;
    color: #6b7280;
}

.debug-info p {
    margin: 0.25rem 0;
}
</style>