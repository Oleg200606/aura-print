import { useAuthStore } from '@/stores/auth'

class ApiService {
    constructor() {
        this.baseURL = '/api'
    }

    async request(endpoint, options = {}) {
        // НЕ используем useAuthStore() здесь, т.к. это не Vue компонент
        const url = `${this.baseURL}${endpoint}`

        const config = {
            headers: {
                'Content-Type': 'application/json',
                ...options.headers,
            },
            ...options,
        }

        // Получаем токен из localStorage напрямую
        const token = localStorage.getItem('token')
        if (token && !endpoint.includes('/login')) {
            config.headers.Authorization = `Bearer ${token}`
        }

        try {
            const response = await fetch(url, config)

            if (response.status === 401) {
                // Если 401, очищаем localStorage и перенаправляем на логин
                localStorage.removeItem('token')
                localStorage.removeItem('user')
                window.location.href = '/login'
                throw new Error('Authentication required')
            }

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`)
            }

            return await response.json()
        } catch (error) {
            console.error('API request failed:', error)
            throw error
        }
    }

    // Исправленный метод login
    async login(username, password) {
        return this.request('/login', {
            method: 'POST',
            body: JSON.stringify({ username, password }),
        })
    }

    // Продукты
    async getProducts() {
        return this.request('/products')
    }

    async createProduct(product) {
        return this.request('/admin/products', {
            method: 'POST',
            body: JSON.stringify(product),
        })
    }

    async updateProduct(id, product) {
        return this.request(`/admin/products/${id}`, {
            method: 'PUT',
            body: JSON.stringify(product),
        })
    }

    async deleteProduct(id) {
        return this.request(`/admin/products/${id}`, {
            method: 'DELETE',
        })
    }

    // Новости
    async getNews() {
        return this.request('/news')
    }

    async createNews(news) {
        return this.request('/admin/news', {
            method: 'POST',
            body: JSON.stringify(news),
        })
    }

    async updateNews(id, news) {
        return this.request(`/admin/news/${id}`, {
            method: 'PUT',
            body: JSON.stringify(news),
        })
    }

    async deleteNews(id) {
        return this.request(`/admin/news/${id}`, {
            method: 'DELETE',
        })
    }

    // Контакты
    async submitContact(contact) {
        return this.request('/contact', {
            method: 'POST',
            body: JSON.stringify(contact),
        })
    }
}

export const apiService = new ApiService()