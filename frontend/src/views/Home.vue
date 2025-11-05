<template>
    <div class="home">
        <!-- Герой секция -->
        <section class="hero">
            <div class="hero-content">
                <h1 class="hero-title">Aura Print</h1>
                <p class="hero-subtitle">Создаём мерч, который рассказывает вашу историю</p>
                <button class="cta-button" @click="$router.push('/catalog')">Смотреть каталог</button>
            </div>
            <div class="hero-visual">
                <div class="floating-card card-1">🎨</div>
                <div class="floating-card card-2">👕</div>
                <div class="floating-card card-3">🎁</div>
            </div>
        </section>

        <!-- Новости и обновления -->
        <section class="news-section">
            <h2>Новости и обновления</h2>
            <div class="news-grid">
                <div v-for="item in news" :key="item.id" class="news-card">
                    <img :src="item.image" :alt="item.title" class="news-image">
                    <div class="news-content">
                        <h3>{{ item.title }}</h3>
                        <p>{{ item.content }}</p>
                        <span class="news-date">{{ item.date }}</span>
                    </div>
                </div>
            </div>

            <!-- Админ панель для новостей -->
            <AdminPanel @news-added="addNews" />
        </section>

        <!-- Форма обратной связи -->
        <section class="contact-section">
            <h2>Свяжитесь с нами</h2>
            <ContactForm />
        </section>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import AdminPanel from '../components/AdminPanel.vue'
import ContactForm from '../components/ContactForm.vue'
import { useAdminStore } from '../stores/admin'

const adminStore = useAdminStore()
const news = ref([
    {
        id: 1,
        title: 'Запуск нового производства',
        content: 'Расширили мощности для больших заказов',
        date: '15.12.2024',
        image: '/api/placeholder/300/200'
    }
])

const addNews = (newsItem) => {
    news.value.unshift(newsItem)
}
</script>

<style scoped>
.hero {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 4rem;
    align-items: center;
    margin-bottom: 4rem;
}

.hero-title {
    font-size: 3.5rem;
    font-weight: 700;
    background: linear-gradient(135deg, #667eea, #764ba2);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    margin-bottom: 1rem;
}

.hero-subtitle {
    font-size: 1.25rem;
    color: #6b7280;
    margin-bottom: 2rem;
}

.cta-button {
    background: linear-gradient(135deg, #667eea, #764ba2);
    color: white;
    border: none;
    padding: 1rem 2rem;
    border-radius: 0.5rem;
    font-size: 1.1rem;
    font-weight: 600;
    cursor: pointer;
    transition: transform 0.3s;
}

.cta-button:hover {
    transform: translateY(-2px);
}

.hero-visual {
    position: relative;
    height: 300px;
}

.floating-card {
    position: absolute;
    background: white;
    border-radius: 1rem;
    padding: 2rem;
    font-size: 3rem;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
    animation: float 3s ease-in-out infinite;
}

.card-1 {
    top: 0;
    left: 0;
    animation-delay: 0s;
}

.card-2 {
    top: 50px;
    right: 0;
    animation-delay: 1s;
}

.card-3 {
    bottom: 0;
    left: 50%;
    transform: translateX(-50%);
    animation-delay: 2s;
}

@keyframes float {

    0%,
    100% {
        transform: translateY(0px);
    }

    50% {
        transform: translateY(-20px);
    }
}

.news-section {
    margin-bottom: 4rem;
}

.news-section h2 {
    text-align: center;
    margin-bottom: 2rem;
    font-size: 2.5rem;
}

.news-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 2rem;
    margin-bottom: 3rem;
}

.news-card {
    background: white;
    border-radius: 1rem;
    overflow: hidden;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
    transition: transform 0.3s;
}

.news-card:hover {
    transform: translateY(-5px);
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
    margin-bottom: 0.5rem;
    color: #2d3748;
}

.news-date {
    color: #9ca3af;
    font-size: 0.875rem;
}

.contact-section h2 {
    text-align: center;
    margin-bottom: 2rem;
    font-size: 2.5rem;
}
</style>