<template>
    <div class="about">
        <div class="about-hero">
            <h1>О компании Aura Print</h1>
            <p>Мы создаём мерч, который вдохновляет и объединяет</p>
        </div>

        <div class="about-content">
            <section class="about-section">
                <h2>Наша миссия</h2>
                <p>Мы помогаем компаниям выражать свою уникальность через качественный мерч. От идеи до реализации - мы
                    сопровождаем каждый этап создания продукции.</p>
            </section>

            <section class="about-section">
                <h2>Наши возможности</h2>
                <div class="capabilities-grid">
                    <div class="capability-card">
                        <div class="capability-icon">🏭</div>
                        <h3>Собственное производство</h3>
                        <p>Современное оборудование для печати и вышивки</p>
                    </div>
                    <div class="capability-card">
                        <div class="capability-icon">🛒</div>
                        <h3>Продажи на маркетплейсах</h3>
                        <p>Wildberries, Ozon и другие площадки</p>
                    </div>
                    <div class="capability-card">
                        <div class="capability-icon">🎨</div>
                        <h3>Дизайн услуги</h3>
                        <p>Разработка уникального дизайна для вашего бренда</p>
                    </div>
                </div>
            </section>

            <section class="about-section">
                <h2>Свяжитесь с нами</h2>
                <ContactForm />
            </section>
        </div>
    </div>
</template>

<script setup>
import ContactForm from '../components/ContactForm.vue'
</script>

<style scoped>
.about-hero {
    text-align: center;
    margin-bottom: 4rem;
    padding: 4rem 0;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    border-radius: 1rem;
}

.about-hero h1 {
    font-size: 3rem;
    margin-bottom: 1rem;
}

.about-content {
    max-width: 800px;
    margin: 0 auto;
}

.about-section {
    margin-bottom: 4rem;
}

.about-section h2 {
    font-size: 2rem;
    margin-bottom: 1.5rem;
    color: #2d3748;
}

.capabilities-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 2rem;
}

.capability-card {
    text-align: center;
    padding: 2rem;
    background: white;
    border-radius: 1rem;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.capability-icon {
    font-size: 3rem;
    margin-bottom: 1rem;
}

.capability-card h3 {
    margin-bottom: 1rem;
    color: #2d3748;
}
</style>