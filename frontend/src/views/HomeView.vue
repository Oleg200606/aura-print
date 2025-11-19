<template>
  <div class="home">
    <!-- Hero Section -->
    <section class="hero">
      <div class="hero-content">
        <h1 class="hero-title">AuraPrint</h1>
        <p class="hero-subtitle">Профессиональная печать и корпоративный мерч</p>
        <p class="hero-description">
          Современное оборудование, качественные материалы и индивидуальный подход к каждому клиенту
        </p>
        <div class="hero-buttons">
          <router-link to="/catalog" class="btn btn-primary">Смотреть каталог</router-link>
          <router-link to="/contact" class="btn btn-secondary">Получить консультацию</router-link>
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
            <img
              src="https://images.unsplash.com/photo-1586363104862-3a5e2ab60d99?w=500&h=300&fit=crop"
              alt="Печать на одежде"
              loading="lazy"
            />
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
            <img
              src="https://images.unsplash.com/photo-1549465227-3d42e0078eed?w=500&h=300&fit=crop"
              alt="Сувенирная продукция"
              loading="lazy"
            />
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
            <img
              src="https://images.unsplash.com/photo-1552664730-d307ca884978?w=500&h=300&fit=crop"
              alt="Корпоративный мерч"
              loading="lazy"
            />
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
              <img
                :src="getImageUrl(item.image_url)"
                :alt="item.title"
                class="news-image"
                @error="handleImageError"
                loading="lazy"
              />
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

<script>
import { computed, onMounted } from "vue";
import { useStore } from "vuex";

export default {
  name: "Home",
  setup() {
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

    const getImageUrl = (imagePath) => {
      if (!imagePath) {
        return "https://images.unsplash.com/photo-1551288049-bebda4e38f71?w=500&h=300&fit=crop";
      }

      if (imagePath.startsWith("http")) {
        return imagePath;
      }

      if (imagePath.startsWith("/")) {
        return import.meta.env.BACKEND_HOST + `${imagePath}`;
      }

      return import.meta.env.BACKEND_HOST + `/uploads/news/${imagePath}`;
    };

    const handleImageError = (event) => {
      event.target.src =
        "https://images.unsplash.com/photo-1551288049-bebda4e38f71?w=500&h=300&fit=crop";
    };

    return {
      news,
      formatDate,
      truncateText,
      getImageUrl,
      handleImageError,
    };
  },
};
</script>

<style scoped>
/* Mobile First Approach */
.hero {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.9), rgba(118, 75, 162, 0.9));
  color: white;
  text-align: center;
  padding: 2rem 1rem;
  position: relative;
  min-height: 60vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.hero-content {
  max-width: 800px;
  margin: 0 auto;
}

.hero-title {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  font-weight: bold;
  line-height: 1.2;
}

.hero-subtitle {
  font-size: 1.4rem;
  margin-bottom: 1rem;
  opacity: 0.95;
  font-weight: 300;
  line-height: 1.3;
}

.hero-description {
  font-size: 1.1rem;
  margin-bottom: 2rem;
  opacity: 0.9;
  line-height: 1.5;
}

.hero-buttons {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin-bottom: 2rem;
}

.hero-stats {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.5rem;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  padding: 1.5rem;
  border-radius: 15px;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.stat-item {
  text-align: center;
}

.stat-number {
  font-size: 2rem;
  font-weight: bold;
  margin-bottom: 0.5rem;
}

.stat-label {
  font-size: 0.9rem;
  opacity: 0.9;
}

.section-header {
  text-align: center;
  margin-bottom: 2rem;
}

.section-header h2 {
  font-size: 2rem;
  margin-bottom: 1rem;
  color: #2c3e50;
  line-height: 1.2;
}

.section-header p {
  font-size: 1.1rem;
  color: #666;
  line-height: 1.4;
}

/* Grid System */
.grid {
  display: grid;
  gap: 1.5rem;
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
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  transition: all 0.3s ease;
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

.service-content {
  padding: 1.5rem;
}

.service-card h3 {
  margin-bottom: 1rem;
  color: #2c3e50;
  font-size: 1.3rem;
}

.service-features {
  text-align: left;
  list-style: none;
  padding: 0;
  margin-top: 1rem;
}

.service-features li {
  padding: 0.4rem 0;
  color: #555;
  position: relative;
  padding-left: 1.5rem;
  font-size: 0.9rem;
}

.service-features li:before {
  content: "✓";
  color: #667eea;
  font-weight: bold;
  position: absolute;
  left: 0;
}

/* Advantages */
.advantages-section {
  background: #f8f9fa;
  padding: 3rem 0;
}

.advantage-card {
  text-align: center;
  padding: 1.5rem;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.advantage-icon {
  font-size: 2.5rem;
  margin-bottom: 1rem;
}

.advantage-card h4 {
  margin-bottom: 1rem;
  color: #2c3e50;
  font-size: 1.2rem;
}

.advantage-card p {
  color: #666;
  line-height: 1.5;
  font-size: 0.95rem;
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
  gap: 1rem;
  padding: 1.5rem;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.step-number {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.2rem;
  font-weight: bold;
  flex-shrink: 0;
}

.step-content h4 {
  margin-bottom: 0.5rem;
  color: #2c3e50;
  font-size: 1.1rem;
}

.step-content p {
  color: #666;
  line-height: 1.5;
  font-size: 0.95rem;
}

/* News Section */
.news-section {
  background: #f8f9fa;
  padding: 3rem 0;
}

.news-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.5rem;
}

.news-item {
  background: white;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.news-image-container {
  position: relative;
  width: 100%;
  height: 200px;
  overflow: hidden;
  background: #f8f9fa;
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
  padding: 1rem;
}

.news-date-badge {
  background: linear-gradient(135deg, #667eea, #764ba2);
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
  font-size: 1.2rem;
  font-weight: 700;
  color: #2c3e50;
  margin-bottom: 0.75rem;
  line-height: 1.3;
}

.news-excerpt {
  color: #666;
  line-height: 1.5;
  font-size: 0.9rem;
}

/* CTA Section */
.cta-section {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 3rem 0;
  text-align: center;
}

.cta-content h2 {
  font-size: 2rem;
  margin-bottom: 1rem;
  line-height: 1.2;
}

.cta-content p {
  font-size: 1.1rem;
  margin-bottom: 2rem;
  opacity: 0.9;
  line-height: 1.4;
}

.cta-buttons {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

/* Buttons */
.btn {
  padding: 1rem 1.5rem;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  text-align: center;
}

.btn-primary {
  background: white;
  color: #667eea;
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(255, 255, 255, 0.2);
}

.btn-secondary {
  background: transparent;
  color: white;
  border: 2px solid white;
}

.btn-secondary:hover {
  background: white;
  color: #667eea;
  transform: translateY(-2px);
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
  padding: 3rem 1rem;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.empty-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.empty-news h3 {
  margin-bottom: 1rem;
  color: #2c3e50;
}

.empty-news p {
  color: #666;
  line-height: 1.5;
}

/* Tablet Styles */
@media (min-width: 768px) {
  .hero {
    padding: 4rem 2rem;
  }

  .hero-title {
    font-size: 3rem;
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
}

/* Desktop Styles */
@media (min-width: 1024px) {
  .hero {
    padding: 6rem 2rem;
  }

  .hero-title {
    font-size: 3.5rem;
  }

  .grid-3 {
    grid-template-columns: repeat(3, 1fr);
  }

  .grid-4 {
    grid-template-columns: repeat(4, 1fr);
  }

  .process-steps {
    grid-template-columns: repeat(2, 1fr);
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
</style>
