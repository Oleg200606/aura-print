import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { apiService } from '@/services/api' // Импортируем apiService

export const useAuthStore = defineStore('auth', () => {
    const token = ref(localStorage.getItem('token'))
    const user = ref(JSON.parse(localStorage.getItem('user') || 'null'))
    const router = useRouter()

    const isAuthenticated = computed(() => !!token.value)

    const login = async (username, password) => {
        try {
            // Используем apiService вместо прямого fetch
            const data = await apiService.login(username, password)

            token.value = data.token
            user.value = data.user

            localStorage.setItem('token', data.token)
            localStorage.setItem('user', JSON.stringify(data.user))

            return true
        } catch (error) {
            console.error('Login error:', error)

            // Более информативные ошибки
            if (error.message.includes('404')) {
                throw new Error('Сервер недоступен. Проверьте, запущен ли бэкенд.')
            } else if (error.message.includes('Network Error')) {
                throw new Error('Ошибка сети. Проверьте подключение.')
            } else {
                throw new Error('Неверный логин или пароль')
            }
        }
    }

    const logout = () => {
        token.value = null
        user.value = null
        localStorage.removeItem('token')
        localStorage.removeItem('user')
        router.push('/login')
    }

    const getAuthHeaders = () => {
        return {
            'Authorization': `Bearer ${token.value}`,
            'Content-Type': 'application/json',
        }
    }

    return {
        token,
        user,
        isAuthenticated,
        login,
        logout,
        getAuthHeaders,
    }
})