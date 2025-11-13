<template>
    <div class="home">
        <!-- Hero Section -->
        <section class="hero">
            <div class="hero-content">
                <h1 class="hero-title">AuraPrint</h1>
                <p class="hero-subtitle">Производство маркетингового оборудования и корпоративного мерча</p>
                <router-link to="/catalog" class="btn btn-primary">Смотреть каталог</router-link>
            </div>
        </section>

        <!-- Services Section -->
        <section class="container">
            <h2 class="text-center mb-4">Наши услуги</h2>
            <div class="grid grid-3">
                <div class="service-card card">
                    <div class="service-image">
                        <img src="https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=400" alt="Футболки">
                    </div>
                    <div class="service-content">
                        <h3>Футболки</h3>
                        <p>Качественная печать на футболках любого размера и цвета</p>
                    </div>
                </div>
                <div class="service-card card">
                    <div class="service-image">
                        <img src="https://images.unsplash.com/photo-1544787219-7f47ccb76574?w=400" alt="Кружки">
                    </div>
                    <div class="service-content">
                        <h3>Кружки</h3>
                        <p>Сублимационная печать на кружках с вашим дизайном</p>
                    </div>
                </div>
                <div class="service-card card">
                    <div class="service-image">
                        <img src="https://images.unsplash.com/photo-1560472354-b33ff0c44a43?w=400"
                            alt="Корпоративный мерч">
                    </div>
                    <div class="service-content">
                        <h3>Корпоративный мерч</h3>
                        <p>Полный цикл производства корпоративной продукции</p>
                    </div>
                </div>
            </div>
        </section>

        <!-- News Section -->
        <section class="container news-section">
            <h2 class="text-center mb-4">Последние новости</h2>
            <div class="grid grid-2">
                <div v-for="item in news" :key="item.id" class="news-card card">
                    <div class="news-image-container">
                        <img :src="getImageUrl(item.image_url)" :alt="item.title" class="news-image"
                            @error="handleImageError">
                    </div>
                    <div class="news-content">
                        <h3>{{ item.title }}</h3>
                        <p>{{ item.content }}</p>
                        <small>{{ formatDate(item.created_at) }}</small>
                    </div>
                </div>
            </div>

            <!-- Если новостей нет -->
            <div v-if="news.length === 0" class="empty-news">
                <p>Пока нет новостей. Зайдите позже!</p>
            </div>
        </section>

        <!-- CTA Section -->
        <section class="cta-section">
            <div class="container">
                <h2>Готовы начать проект?</h2>
                <p>Свяжитесь с нами для получения бесплатной консультации</p>
                <router-link to="/contact" class="btn btn-primary">Связаться с нами</router-link>
            </div>
        </section>
    </div>
</template>

<script>
import { computed, onMounted } from 'vue'
import { useStore } from 'vuex'

export default {
    name: 'Home',
    setup() {
        const store = useStore()
        const news = computed(() => store.state.news)

        onMounted(() => {
            store.dispatch('fetchNews')
        })

        const formatDate = (dateString) => {
            if (!dateString) return 'Дата не указана'
            return new Date(dateString).toLocaleDateString('ru-RU')
        }

        const getImageUrl = (imagePath) => {
            console.log('Image path:', imagePath) // Для дебага

            if (!imagePath) {
                return 'https://images.unsplash.com/photo-1551288049-bebda4e38f71?w=400&h=250&fit=crop'
            }

            // Если это полный URL
            if (imagePath.startsWith('http')) {
                return imagePath
            }

            // Если это локальный путь
            if (imagePath.startsWith('/')) {
                return `http://localhost:8081${imagePath}`
            }

            // По умолчанию
            return `http://localhost:8081/uploads/news/${imagePath}`
        }

        const handleImageError = (event) => {
            console.log('Image load error:', event)
            event.target.src = 'https://images.unsplash.com/photo-1551288049-bebda4e38f71?w=400&h=250&fit=crop'
        }

        return {
            news,
            formatDate,
            getImageUrl,
            handleImageError
        }
    }
}
</script>

<style scoped>
.hero {
    background: linear-gradient(135deg, rgba(102, 126, 234, 0.9), rgba(118, 75, 162, 0.9)),
        url('https://images.unsplash.com/photo-1560472354-b33ff0c44a43?w=1200');
    background-size: cover;
    background-position: center;
    color: white;
    text-align: center;
    padding: 8rem 2rem;
    position: relative;
}

.hero-content {
    max-width: 800px;
    margin: 0 auto;
    position: relative;
    z-index: 2;
}

.hero-title {
    font-size: 4rem;
    margin-bottom: 1rem;
    font-weight: bold;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
}

.hero-subtitle {
    font-size: 1.5rem;
    margin-bottom: 2rem;
    opacity: 0.95;
}

.service-card {
    text-align: center;
    padding: 0;
    overflow: hidden;
}

.service-image {
    width: 100%;
    height: 200px;
    overflow: hidden;
}

.service-image img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform 0.3s ease;
}

.service-card:hover .service-image img {
    transform: scale(1.05);
}

.service-content {
    padding: 1.5rem;
}

.service-card h3 {
    margin-bottom: 1rem;
    color: #2c3e50;
}

.news-section {
    background: #f8f9fa;
    padding: 4rem 0;
}

.news-card {
    display: flex;
    flex-direction: column;
    height: 100%;
}

.news-image-container {
    width: 100%;
    height: 250px;
    overflow: hidden;
    background: #f0f0f0;
}

.news-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform 0.3s ease;
}

.news-card:hover .news-image {
    transform: scale(1.05);
}

.news-content {
    padding: 1.5rem;
    flex-grow: 1;
    display: flex;
    flex-direction: column;
}

.news-content h3 {
    margin-bottom: 1rem;
    color: #2c3e50;
    font-size: 1.25rem;
}

.news-content p {
    flex-grow: 1;
    margin-bottom: 1rem;
    color: #666;
    line-height: 1.5;
}

.news-content small {
    color: #999;
    font-size: 0.9rem;
}

.empty-news {
    text-align: center;
    padding: 2rem;
    color: #666;
    background: white;
    border-radius: 10px;
}

.cta-section {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    text-align: center;
    padding: 4rem 2rem;
}

.cta-section h2 {
    font-size: 2.5rem;
    margin-bottom: 1rem;
}

.cta-section p {
    font-size: 1.2rem;
    margin-bottom: 2rem;
    opacity: 0.9;
}

/* Адаптивность */
@media (max-width: 768px) {
    .hero-title {
        font-size: 2.5rem;
    }

    .hero-subtitle {
        font-size: 1.2rem;
    }

    .grid-3 {
        grid-template-columns: 1fr;
    }
}
</style>