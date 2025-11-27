<template>
  <div class="home">
    <!-- Hero Section -->
    <section class="hero">
      <div class="hero-background">
        <div class="hero-shapes">
          <div class="shape shape-1"></div>
          <div class="shape shape-2"></div>
          <div class="shape shape-3"></div>
        </div>
      </div>

      <div class="hero-content">
        <div class="hero-badge">
          🎉 Новый бренд! Теперь мы - Купи-Принт.ру
        </div>
        <h1 class="hero-title">
          <span class="title-main">Купи-Принт</span>
          <span class="title-dot">.ру</span>
        </h1>
        <p class="hero-subtitle">Печать, которая вдохновляет</p>
        <p class="hero-description">
          Современная печать на любой поверхности. От одежды до сувениров -
          создаем то, что подчеркнет ваш стиль
        </p>
        <div class="hero-buttons">
          <router-link to="/catalog" class="btn btn-primary">
            🛍️ Смотреть каталог
          </router-link>
          <router-link to="/contact" class="btn btn-secondary">
            💬 Бесплатная консультация
          </router-link>
        </div>
      </div>

      <div class="hero-visual">
        <div class="floating-products">
          <div class="product-preview preview-1">👕</div>
          <div class="product-preview preview-2">☕</div>
          <div class="product-preview preview-3">🎁</div>
        </div>
      </div>

      <div class="hero-stats">
        <div class="stat-item">
          <div class="stat-number">500+</div>
          <div class="stat-label">Выполненных проектов</div>
        </div>
        <div class="stat-item">
          <div class="stat-number">3+</div>
          <div class="stat-label">Года на рынке</div>
        </div>
        <div class="stat-item">
          <div class="stat-number">98%</div>
          <div class="stat-label">Довольных клиентов</div>
        </div>
      </div>
    </section>

    <!-- Services Section -->
    <section class="container">
      <div class="section-header">
        <h2>Наши услуги</h2>
        <p>Полный цикл производства от дизайна до готовой продукции</p>
      </div>
      <div class="grid grid-3">
        <div class="service-card card">
          <div class="service-image">
            <img src="https://images.unsplash.com/photo-1586363104862-3a5e2ab60d99?w=500&h=300&fit=crop"
              alt="Печать на одежде" loading="lazy" />
          </div>
          <div class="service-content">
            <h3>Печать на одежде</h3>
            <p>Футболки, худи, поло и другая одежда с качественной печатью</p>
            <ul class="service-features">
              <li>Шелкография и термоперенос</li>
              <li>Любые размеры и цвета</li>
              <li>Собственное производство</li>
            </ul>
          </div>
        </div>
        <div class="service-card card">
          <div class="service-image">
            <img src="https://images.unsplash.com/photo-1549465227-3d42e0078eed?w=500&h=300&fit=crop"
              alt="Сувенирная продукция" loading="lazy" />
          </div>
          <div class="service-content">
            <h3>Сувенирная продукция</h3>
            <p>Кружки, брелоки, чехлы и другие промо-материалы</p>
            <ul class="service-features">
              <li>Сублимационная печать</li>
              <li>Лазерная гравировка</li>
              <li>Широкий ассортимент</li>
            </ul>
          </div>
        </div>
        <div class="service-card card">
          <div class="service-image">
            <img src="https://images.unsplash.com/photo-1552664730-d307ca884978?w=500&h=300&fit=crop"
              alt="Корпоративный мерч" loading="lazy" />
          </div>
          <div class="service-content">
            <h3>Корпоративный мерч</h3>
            <p>Комплексные решения для бизнеса и мероприятий</p>
            <ul class="service-features">
              <li>Брендирование под ключ</li>
              <li>Разработка дизайна</li>
              <li>Срочное производство</li>
            </ul>
          </div>
        </div>
      </div>
    </section>

    <!-- Advantages Section -->
    <section class="advantages-section">
      <div class="container">
        <div class="section-header">
          <h2>Почему выбирают нас</h2>
          <p>Мы обеспечиваем качество на каждом этапе работы</p>
        </div>
        <div class="grid grid-4">
          <div class="advantage-card">
            <div class="advantage-icon">⚡</div>
            <h4>Срочное производство</h4>
            <p>Выполняем заказы любой сложности в кратчайшие сроки</p>
          </div>
          <div class="advantage-card">
            <div class="advantage-icon">🎨</div>
            <h4>Индивидуальный дизайн</h4>
            <p>Разработаем уникальный дизайн или воплотим ваши идеи</p>
          </div>
          <div class="advantage-card">
            <div class="advantage-icon">🏭</div>
            <h4>Собственное производство</h4>
            <p>Контролируем каждый этап от печати до упаковки</p>
          </div>
          <div class="advantage-card">
            <div class="advantage-icon">💎</div>
            <h4>Гарантия качества</h4>
            <p>Используем только проверенные материалы и технологии</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Process Section -->
    <section class="container">
      <div class="section-header">
        <h2>Как мы работаем</h2>
        <p>Простой и понятный процесс от заявки до получения заказа</p>
      </div>
      <div class="process-steps">
        <div class="process-step">
          <div class="step-number">1</div>
          <div class="step-content">
            <h4>Заявка и консультация</h4>
            <p>Обсуждаем ваши задачи и предлагаем оптимальное решение</p>
          </div>
        </div>
        <div class="process-step">
          <div class="step-number">2</div>
          <div class="step-content">
            <h4>Дизайн и согласование</h4>
            <p>Разрабатываем макет и утверждаем его с вами</p>
          </div>
        </div>
        <div class="process-step">
          <div class="step-number">3</div>
          <div class="step-content">
            <h4>Производство</h4>
            <p>Изготавливаем продукцию с контролем качества</p>
          </div>
        </div>
        <div class="process-step">
          <div class="step-number">4</div>
          <div class="step-content">
            <h4>Доставка</h4>
            <p>Передаем готовый заказ удобным для вас способом</p>
          </div>
        </div>
      </div>
    </section>

    <!-- News Section -->
    <section class="news-section">
      <div class="container">
        <div class="section-header">
          <h2>Последние новости</h2>
          <p>Будьте в курсе наших новинок и акций</p>
        </div>

        <div v-if="news.length > 0" class="news-grid">
          <div v-for="item in news" :key="item.id" class="news-item">
            <div class="news-image-container">
              <img :src="item.imageUrl()" :alt="item.title" class="news-image" @error="handleImageError"
                loading="lazy" />
              <div class="news-overlay">
                <div class="news-date-badge">
                  {{ formatDate(item.created_at) }}
                </div>
              </div>
            </div>
            <div class="news-info">
              <h3 class="news-title">{{ item.title }}</h3>
              <p class="news-excerpt">{{ truncateText(item.content, 120) }}</p>
            </div>
          </div>
        </div>

        <div v-else class="empty-news">
          <div class="empty-icon">📰</div>
          <h3>Пока нет новостей</h3>
          <p>Следите за обновлениями, скоро здесь появятся интересные материалы</p>
        </div>
      </div>
    </section>

    <!-- CTA Section -->
    <section class="cta-section">
      <div class="container">
        <div class="cta-content">
          <h2>Начните свой проект сегодня</h2>
          <p>Получите бесплатную консультацию и расчет стоимости вашего заказа</p>
          <div class="cta-buttons">
            <router-link to="/contact" class="btn btn-primary">Обсудить проект</router-link>
            <a href="tel:+79955054001" class="btn btn-secondary">Позвонить нам</a>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { computed, onMounted } from "vue";
import { useStore } from "vuex";
import { DEFAULT_IMAGE_URL } from "@/data";

const store = useStore();
const news = computed(() => store.state.news);

onMounted(() => {
  store.dispatch("fetchNews");
});

const formatDate = (dateString) => {
  if (!dateString) return "Дата не указана";
  const date = new Date(dateString);
  return date.toLocaleDateString("ru-RU", {
    day: "numeric",
    month: "long",
    year: "numeric",
  });
};

const truncateText = (text, maxLength) => {
  if (!text) return "";
  if (text.length <= maxLength) return text;
  return text.substring(0, maxLength) + "...";
};

const handleImageError = (event) => {
  event.target.src = DEFAULT_IMAGE_URL;
};
</script>

<style scoped>
/* Новая цветовая схема */
:root {
  --primary: #FF6B35;
  --primary-dark: #E55A2B;
  --secondary: #2EC4B6;
  --secondary-dark: #25A99A;
  --dark: #1A1F2B;
  --light: #F8F9FA;
  --accent: #FFD166;
  --text-dark: #2D3748;
  --text-light: #718096;
}

/* Hero Section - полностью переработан */
.hero {
  background: linear-gradient(135deg, var(--dark) 0%, #2D3748 100%);
  color: white;
  padding: 4rem 1rem;
  position: relative;
  overflow: hidden;
  min-height: 80vh;
  display: flex;
  align-items: center;
}

.hero-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 1;
}

.hero-shapes .shape {
  position: absolute;
  border-radius: 50%;
  background: radial-gradient(circle, var(--primary) 0%, transparent 70%);
  opacity: 0.1;
}

.shape-1 {
  width: 300px;
  height: 300px;
  top: -100px;
  right: -100px;
}

.shape-2 {
  width: 200px;
  height: 200px;
  bottom: 100px;
  left: -50px;
  background: radial-gradient(circle, var(--secondary) 0%, transparent 70%);
}

.shape-3 {
  width: 150px;
  height: 150px;
  top: 50%;
  right: 20%;
  background: radial-gradient(circle, var(--accent) 0%, transparent 70%);
}

.hero-content {
  position: relative;
  z-index: 2;
  max-width: 1200px;
  margin: 0 auto;
  text-align: center;
}

.hero-badge {
  display: inline-block;
  background: linear-gradient(135deg, var(--primary), var(--accent));
  padding: 0.5rem 1.5rem;
  border-radius: 25px;
  font-size: 0.9rem;
  font-weight: 600;
  margin-bottom: 2rem;
  animation: pulse 2s infinite;
}

@keyframes pulse {

  0%,
  100% {
    transform: scale(1);
  }

  50% {
    transform: scale(1.05);
  }
}

.hero-title {
  font-size: 4rem;
  font-weight: 800;
  margin-bottom: 1rem;
  line-height: 1.1;
}

.title-main {
  background: linear-gradient(135deg, var(--primary), var(--accent));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.title-dot {
  color: var(--secondary);
}

.hero-subtitle {
  font-size: 1.5rem;
  margin-bottom: 1.5rem;
  color: var(--secondary);
  font-weight: 300;
}

.hero-description {
  font-size: 1.2rem;
  margin-bottom: 3rem;
  opacity: 0.9;
  max-width: 600px;
  margin-left: auto;
  margin-right: auto;
  line-height: 1.6;
}

.hero-buttons {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
  margin-bottom: 3rem;
}

/* Floating products animation */
.hero-visual {
  position: relative;
  z-index: 2;
}

.floating-products {
  position: absolute;
  right: 10%;
  top: 50%;
  transform: translateY(-50%);
  z-index: 2;
}

.product-preview {
  width: 80px;
  height: 80px;
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  position: absolute;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.preview-1 {
  top: -50px;
  left: 0;
  animation: float 3s ease-in-out infinite;
}

.preview-2 {
  top: 50px;
  left: 100px;
  animation: float 3s ease-in-out infinite 1s;
}

.preview-3 {
  top: 150px;
  left: 50px;
  animation: float 3s ease-in-out infinite 2s;
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

.hero-stats {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 2rem;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  padding: 2rem;
  border-radius: 20px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  max-width: 600px;
  margin: 0 auto;
  position: relative;
  z-index: 2;
}

.stat-item {
  text-align: center;
}

.stat-number {
  font-size: 2.5rem;
  font-weight: bold;
  margin-bottom: 0.5rem;
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.stat-label {
  font-size: 0.9rem;
  opacity: 0.9;
}

/* Обновленные кнопки */
.btn {
  padding: 1rem 2rem;
  border: none;
  border-radius: 12px;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
}

.btn-primary {
  background: linear-gradient(135deg, var(--primary), var(--primary-dark));
  color: white;
  box-shadow: 0 8px 25px rgba(255, 107, 53, 0.3);
}

.btn-primary:hover {
  transform: translateY(-3px);
  box-shadow: 0 15px 35px rgba(255, 107, 53, 0.4);
}

.btn-secondary {
  background: transparent;
  color: white;
  border: 2px solid var(--secondary);
  box-shadow: 0 8px 25px rgba(46, 196, 182, 0.2);
}

.btn-secondary:hover {
  background: var(--secondary);
  transform: translateY(-3px);
  box-shadow: 0 15px 35px rgba(46, 196, 182, 0.3);
}

/* Section Header */
.section-header {
  text-align: center;
  margin-bottom: 3rem;
}

.section-header h2 {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  color: var(--dark);
  line-height: 1.2;
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.section-header p {
  font-size: 1.2rem;
  color: var(--text-light);
  line-height: 1.4;
}

/* Grid System */
.grid {
  display: grid;
  gap: 2rem;
}

.grid-3 {
  grid-template-columns: 1fr;
}

.grid-4 {
  grid-template-columns: 1fr;
}

/* Cards */
.card {
  background: white;
  border-radius: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  transition: all 0.3s ease;
  border: 1px solid rgba(0, 0, 0, 0.05);
}

.service-card {
  text-align: center;
  padding: 0;
  overflow: hidden;
}

.service-card:hover {
  transform: translateY(-10px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
}

.service-image {
  width: 100%;
  height: 250px;
  overflow: hidden;
}

.service-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.5s ease;
}

.service-card:hover .service-image img {
  transform: scale(1.1);
}

.service-content {
  padding: 2rem;
}

.service-card h3 {
  margin-bottom: 1rem;
  color: var(--dark);
  font-size: 1.5rem;
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.service-features {
  text-align: left;
  list-style: none;
  padding: 0;
  margin-top: 1rem;
}

.service-features li {
  padding: 0.5rem 0;
  color: var(--text-light);
  position: relative;
  padding-left: 1.5rem;
  font-size: 0.95rem;
}

.service-features li:before {
  content: "✓";
  color: var(--primary);
  font-weight: bold;
  position: absolute;
  left: 0;
}

/* Advantages */
.advantages-section {
  background: linear-gradient(135deg, var(--light) 0%, #FFFFFF 100%);
  padding: 5rem 0;
}

.advantage-card {
  text-align: center;
  padding: 2rem;
  background: white;
  border-radius: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.08);
  border: 1px solid rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.advantage-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.12);
}

.advantage-icon {
  font-size: 3rem;
  margin-bottom: 1.5rem;
  display: block;
}

.advantage-card h4 {
  margin-bottom: 1rem;
  color: var(--dark);
  font-size: 1.3rem;
}

.advantage-card p {
  color: var(--text-light);
  line-height: 1.5;
  font-size: 1rem;
}

/* Process Steps */
.process-steps {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.5rem;
}

.process-step {
  display: flex;
  align-items: flex-start;
  gap: 1.5rem;
  padding: 2rem;
  background: white;
  border-radius: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease;
}

.process-step:hover {
  transform: translateX(10px);
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.12);
}

.step-number {
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  color: white;
  width: 50px;
  height: 50px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem;
  font-weight: bold;
  flex-shrink: 0;
}

.step-content h4 {
  margin-bottom: 0.5rem;
  color: var(--dark);
  font-size: 1.3rem;
}

.step-content p {
  color: var(--text-light);
  line-height: 1.5;
  font-size: 1rem;
}

/* News Section */
.news-section {
  background: linear-gradient(135deg, var(--light) 0%, #FFFFFF 100%);
  padding: 5rem 0;
}

.news-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 2rem;
}

.news-item {
  background: white;
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.news-item:hover {
  transform: translateY(-5px);
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.15);
}

.news-image-container {
  position: relative;
  width: 100%;
  height: 250px;
  overflow: hidden;
  background: var(--light);
}

.news-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.news-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(to bottom, rgba(0, 0, 0, 0.1), rgba(0, 0, 0, 0.2));
  display: flex;
  align-items: flex-start;
  justify-content: flex-start;
  padding: 1.5rem;
}

.news-date-badge {
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
}

.news-info {
  padding: 1.5rem;
}

.news-title {
  font-size: 1.3rem;
  font-weight: 700;
  color: var(--dark);
  margin-bottom: 0.75rem;
  line-height: 1.3;
}

.news-excerpt {
  color: var(--text-light);
  line-height: 1.5;
  font-size: 1rem;
}

/* CTA Section */
.cta-section {
  background: linear-gradient(135deg, var(--primary) 0%, var(--secondary) 100%);
  color: white;
  padding: 5rem 0;
  text-align: center;
  position: relative;
  overflow: hidden;
}

.cta-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: url('data:image/svg+xml,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1000 100" fill="%23ffffff" opacity="0.1"><polygon points="1000,100 1000,0 0,100"/></svg>');
  background-size: cover;
}

.cta-content {
  position: relative;
  z-index: 2;
}

.cta-content h2 {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  line-height: 1.2;
}

.cta-content p {
  font-size: 1.2rem;
  margin-bottom: 2rem;
  opacity: 0.9;
  line-height: 1.4;
}

.cta-buttons {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

/* Container */
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

/* Empty States */
.empty-news {
  text-align: center;
  padding: 4rem 2rem;
  background: white;
  border-radius: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.08);
}

.empty-icon {
  font-size: 4rem;
  margin-bottom: 1.5rem;
}

.empty-news h3 {
  margin-bottom: 1rem;
  color: var(--dark);
  font-size: 1.5rem;
}

.empty-news p {
  color: var(--text-light);
  line-height: 1.5;
  font-size: 1.1rem;
}

/* Tablet Styles */
@media (min-width: 768px) {
  .hero {
    padding: 6rem 2rem;
  }

  .hero-title {
    font-size: 5rem;
  }

  .hero-buttons {
    flex-direction: row;
    justify-content: center;
  }

  .hero-stats {
    grid-template-columns: repeat(3, 1fr);
  }

  .grid-3 {
    grid-template-columns: repeat(2, 1fr);
  }

  .grid-4 {
    grid-template-columns: repeat(2, 1fr);
  }

  .cta-buttons {
    flex-direction: row;
    justify-content: center;
  }

  .news-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .process-steps {
    grid-template-columns: repeat(2, 1fr);
  }
}

/* Desktop Styles */
@media (min-width: 1024px) {
  .hero {
    padding: 8rem 2rem;
  }

  .hero-title {
    font-size: 6rem;
  }

  .grid-3 {
    grid-template-columns: repeat(3, 1fr);
  }

  .grid-4 {
    grid-template-columns: repeat(4, 1fr);
  }

  .news-grid {
    grid-template-columns: repeat(3, 1fr);
  }

  .container {
    padding: 0 2rem;
  }
}

/* Large Desktop */
@media (min-width: 1200px) {
  .news-grid {
    grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  }
}

/* Touch Improvements */
@media (hover: none) and (pointer: coarse) {
  .btn:hover {
    transform: none;
  }

  .card:hover {
    transform: none;
  }

  .service-card:hover .service-image img {
    transform: none;
  }
}

/* Reduced motion for accessibility */
@media (prefers-reduced-motion: reduce) {
  * {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
}

/* Анимации появления */
.service-card,
.advantage-card,
.news-item,
.process-step {
  animation: fadeInUp 0.6s ease-out;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>