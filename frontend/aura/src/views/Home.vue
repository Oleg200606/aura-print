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
                    <div class="service-icon">👕</div>
                    <h3>Футболки</h3>
                    <p>Качественная печать на футболках любого размера и цвета</p>
                </div>
                <div class="service-card card">
                    <div class="service-icon">☕</div>
                    <h3>Кружки</h3>
                    <p>Сублимационная печать на кружках с вашим дизайном</p>
                </div>
                <div class="service-card card">
                    <div class="service-icon">🎁</div>
                    <h3>Корпоративный мерч</h3>
                    <p>Полный цикл производства корпоративной продукции</p>
                </div>
            </div>
        </section>

        <!-- News Section -->
        <section class="container news-section">
            <h2 class="text-center mb-4">Последние новости</h2>
            <div class="grid grid-2">
                <div v-for="item in news" :key="item.id" class="news-card card">
                    <img :src="news.image_url" :alt="news.title" class="news-image">
                    <div class="news-content">
                        <h3>{{ news.title }}</h3>
                        <p>{{ item.content }}</p>
                        <small>{{ formatDate(item.created_at) }}</small>
                    </div>
                </div>
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
            return new Date(dateString).toLocaleDateString('ru-RU')
        }

        return {
            news,
            formatDate
        }
    }
}
</script>

<style scoped>
.hero {
    background: linear-gradient(135deg, rgba(102, 126, 234, 0.8), rgba(118, 75, 162, 0.8)),
        url('https://images.unsplash.com/photo-1560472354-b33ff0c44a43?ixlib=rb-4.0.3&auto=format&fit=crop&w=1920&q=80');
    background-size: cover;
    background-position: center;
    color: white;
    text-align: center;
    padding: 8rem 2rem;
}

.hero-content {
    max-width: 800px;
    margin: 0 auto;
}

.hero-title {
    font-size: 4rem;
    margin-bottom: 1rem;
    font-weight: bold;
}

.hero-subtitle {
    font-size: 1.5rem;
    margin-bottom: 2rem;
    opacity: 0.9;
}

.service-card {
    text-align: center;
    padding: 2rem;
}

.service-icon {
    font-size: 3rem;
    margin-bottom: 1rem;
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
}

.news-image {
    width: 100%;
    height: 200px;
    object-fit: cover;
}

.news-content {
    padding: 1.5rem;
}

.news-content h3 {
    margin-bottom: 1rem;
    color: #2c3e50;
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
</style>