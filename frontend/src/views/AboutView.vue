<template>
  <div class="about">
    <div class="container">
      <div class="about-content">
        <div class="about-text">
          <h1>О компании AuraPrint</h1>
          <p>
            AuraPrint - это современная компания, специализирующаяся на производстве маркетингового
            оборудования и корпоративного мерча. Мы помогаем бизнесу создавать уникальную продукцию,
            которая укрепляет бренд и повышает лояльность клиентов.
          </p>

          <h2>Наши преимущества</h2>
          <ul class="advantages-list">
            <li>✅ Современное оборудование и технологии</li>
            <li>✅ Высокое качество продукции</li>
            <li>✅ Быстрые сроки выполнения заказов</li>
            <li>✅ Индивидуальный подход к каждому клиенту</li>
            <li>✅ Конкурентные цены</li>
          </ul>

          <h2>Наша миссия</h2>
          <p>
            Мы стремимся помочь компаниям выразить свою уникальность через качественную
            полиграфическую продукцию и корпоративный мерч, создавая решения, которые работают на
            укрепление бренда.
          </p>
        </div>

        <div class="about-image">
          <img :src="DEFAULT_IMAGE_URL" alt="Наше производство" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { DEFAULT_IMAGE_URL } from "@/data";
</script>

<style scoped>
.about-content {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 4rem;
  align-items: start;
}

.about-text h1 {
  font-size: 3rem;
  margin-bottom: 2rem;
  color: #2c3e50;
}

.about-text h2 {
  margin: 2rem 0 1rem;
  color: #34495e;
}

.advantages-list {
  list-style: none;
  padding: 0;
}

.advantages-list li {
  padding: 0.5rem 0;
  font-size: 1.1rem;
}

.about-image img {
  width: 100%;
  border-radius: 10px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
}

@media (max-width: 768px) {
  .about-content {
    grid-template-columns: 1fr;
  }
}
</style>
