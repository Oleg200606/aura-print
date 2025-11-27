<template>
  <div class="about">
    <div class="container">
      <div class="about-content">
        <div class="about-text">
          <div class="header-section">
            <h1>О компании Купи-Принт.ру</h1>
            <p class="subtitle">
              Профессиональная печать и корпоративный мерч с 2022 года
            </p>
          </div>

          <div class="about-section">
            <h2>Наша миссия</h2>
            <p>
              Мы помогаем бизнесу выразить свою уникальность через качественную
              полиграфическую продукцию и корпоративный мерч. Наша цель - создавать
              решения, которые работают на укрепление бренда и повышение лояльности клиентов.
            </p>
          </div>

          <div class="about-section">
            <h2>Наши преимущества</h2>
            <ul class="advantages-list">
              <li>
                <span class="advantage-icon">⚡</span>
                <div class="advantage-content">
                  <strong>Современное оборудование</strong>
                  <p>Используем новейшие технологии печати для безупречного качества</p>
                </div>
              </li>
              <li>
                <span class="advantage-icon">🎨</span>
                <div class="advantage-content">
                  <strong>Индивидуальный подход</strong>
                  <p>Разрабатываем уникальные дизайны под конкретные задачи клиента</p>
                </div>
              </li>
              <li>
                <span class="advantage-icon">🚀</span>
                <div class="advantage-content">
                  <strong>Быстрые сроки</strong>
                  <p>Выполняем срочные заказы без потери качества</p>
                </div>
              </li>
              <li>
                <span class="advantage-icon">💎</span>
                <div class="advantage-content">
                  <strong>Гарантия качества</strong>
                  <p>Контролируем каждый этап производства</p>
                </div>
              </li>
            </ul>
          </div>

          <div class="about-section">
            <h2>Наши ценности</h2>
            <div class="values-grid">
              <div class="value-item">
                <div class="value-icon">🤝</div>
                <h4>Клиентоориентированность</h4>
                <p>Ставим интересы клиента на первое место</p>
              </div>
              <div class="value-item">
                <div class="value-icon">🌟</div>
                <h4>Качество</h4>
                <p>Не compromiseим с качеством продукции</p>
              </div>
              <div class="value-item">
                <div class="value-icon">💡</div>
                <h4>Инновации</h4>
                <p>Постоянно развиваемся и внедряем новые технологии</p>
              </div>
              <div class="value-item">
                <div class="value-icon">📈</div>
                <h4>Результат</h4>
                <p>Нацелены на достижение целей наших клиентов</p>
              </div>
            </div>
          </div>
        </div>

        <div class="about-image">
          <div class="image-card">
            <img :src="DEFAULT_IMAGE_URL" alt="Наше производство" />
            <div class="image-overlay">
              <h3>Купи-Принт.ру</h3>
              <p>Качество, которому доверяют</p>
            </div>
          </div>

          <div class="stats-card">
            <div class="stat">
              <div class="stat-number">500+</div>
              <div class="stat-label">проектов</div>
            </div>
            <div class="stat">
              <div class="stat-number">3+</div>
              <div class="stat-label">года опыта</div>
            </div>
            <div class="stat">
              <div class="stat-number">98%</div>
              <div class="stat-label">довольных клиентов</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { DEFAULT_IMAGE_URL } from "@/data";
</script>

<style scoped>
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

.about {
  background: linear-gradient(135deg, var(--light) 0%, #FFFFFF 100%);
  min-height: 100vh;
  padding: 6rem 0 2rem;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

.about-content {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 4rem;
  align-items: start;
}

.header-section {
  margin-bottom: 3rem;
}

.header-section h1 {
  font-size: 3rem;
  font-weight: 700;
  margin-bottom: 1rem;
  color: var(--dark);
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.subtitle {
  font-size: 1.2rem;
  color: var(--text-light);
  line-height: 1.6;
}

.about-section {
  margin-bottom: 3rem;
}

.about-section h2 {
  font-size: 2rem;
  margin-bottom: 1.5rem;
  color: var(--dark);
  font-weight: 600;
}

.about-section p {
  font-size: 1.1rem;
  line-height: 1.7;
  color: var(--text-dark);
  margin-bottom: 1.5rem;
}

.advantages-list {
  list-style: none;
  padding: 0;
}

.advantages-list li {
  display: flex;
  align-items: flex-start;
  gap: 1rem;
  margin-bottom: 2rem;
  padding: 1.5rem;
  background: white;
  border-radius: 15px;
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.08);
  transition: transform 0.3s ease;
}

.advantages-list li:hover {
  transform: translateX(10px);
}

.advantage-icon {
  font-size: 2rem;
  flex-shrink: 0;
}

.advantage-content strong {
  display: block;
  font-size: 1.2rem;
  margin-bottom: 0.5rem;
  color: var(--dark);
}

.advantage-content p {
  margin: 0;
  color: var(--text-light);
  font-size: 1rem;
}

.values-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.5rem;
}

.value-item {
  text-align: center;
  padding: 2rem 1rem;
  background: white;
  border-radius: 15px;
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.08);
  transition: transform 0.3s ease;
}

.value-item:hover {
  transform: translateY(-5px);
}

.value-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.value-item h4 {
  margin-bottom: 1rem;
  color: var(--dark);
  font-size: 1.2rem;
}

.value-item p {
  color: var(--text-light);
  margin: 0;
  line-height: 1.5;
}

.about-image {
  position: sticky;
  top: 6rem;
}

.image-card {
  position: relative;
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
  margin-bottom: 2rem;
}

.image-card img {
  width: 100%;
  height: 300px;
  object-fit: cover;
}

.image-overlay {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  background: linear-gradient(transparent, rgba(0, 0, 0, 0.8));
  color: white;
  padding: 2rem;
}

.image-overlay h3 {
  font-size: 1.5rem;
  margin-bottom: 0.5rem;
  color: var(--accent);
}

.image-overlay p {
  margin: 0;
  opacity: 0.9;
}

.stats-card {
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  color: white;
  padding: 2rem;
  border-radius: 20px;
  text-align: center;
  box-shadow: 0 15px 35px rgba(255, 107, 53, 0.3);
}

.stat {
  margin-bottom: 1.5rem;
}

.stat:last-child {
  margin-bottom: 0;
}

.stat-number {
  font-size: 2.5rem;
  font-weight: bold;
  margin-bottom: 0.5rem;
}

.stat-label {
  font-size: 1rem;
  opacity: 0.9;
}

/* Адаптивность */
@media (max-width: 1024px) {
  .about-content {
    grid-template-columns: 1fr;
    gap: 3rem;
  }

  .about-image {
    position: static;
  }

  .values-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 768px) {
  .about {
    padding: 5rem 0 2rem;
  }

  .header-section h1 {
    font-size: 2rem;
  }

  .about-section h2 {
    font-size: 1.5rem;
  }

  .advantages-list li {
    flex-direction: column;
    text-align: center;
  }

  .image-card img {
    height: 250px;
  }
}

@media (max-width: 480px) {
  .header-section h1 {
    font-size: 1.75rem;
  }

  .stats-card {
    padding: 1.5rem;
  }

  .stat-number {
    font-size: 2rem;
  }
}
</style>