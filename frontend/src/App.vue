<template>
  <div id="app">
    <nav class="navbar">
      <div class="nav-container">
        <router-link to="/" class="nav-logo">
          <span class="logo-text">AuraPrint</span>
          <span class="logo-subtitle">Маркетинговое оборудование</span>
        </router-link>
        <div class="nav-menu">
          <router-link to="/" class="nav-link">Главная</router-link>
          <router-link to="/catalog" class="nav-link">Каталог</router-link>
          <router-link to="/about" class="nav-link">О нас</router-link>
          <router-link to="/contact" class="nav-link">Контакты</router-link>
          <router-link v-if="isAdmin" to="/admin" class="nav-link">Админ</router-link>
          <button v-if="isAdmin" @click="logout" class="nav-link admin-link">Выйти</button>
        </div>
      </div>
    </nav>

    <main class="main-content">
      <router-view />
    </main>

    <footer class="footer">
      <div class="footer-content">
        <div class="footer-section">
          <h3>AuraPrint</h3>
          <p>Производство маркетингового оборудования и корпоративного мерча</p>
        </div>
        <div class="footer-section">
          <h4>Контакты</h4>
          <p class="clickable" @click="sendEmail">📧 Email: info@auraprint.ru</p>
          <p class="clickable" @click="makeCall">📞 Телефон: +7 (995) 505-40-01</p>
          <p class="clickable" @click="openMap">📍 Адрес: г. Москва, пр-кт Волгоградский 32к31</p>
        </div>
        <div class="footer-section">
          <h4>Услуги</h4>
          <p>Футболки</p>
          <p>Кружки</p>
          <p>Корпоративный мерч</p>
        </div>
      </div>
      <div class="footer-bottom">
        <p>&copy; 2025 AuraPrint. Все права защищены.</p>
      </div>
    </footer>
  </div>
</template>

<script>
import { computed } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";

export default {
  name: "App",
  setup() {
    const store = useStore();
    const router = useRouter();

    const isAdmin = computed(() => store.state.isAdmin);

    // Функция для отправки email
    const sendEmail = () => {
      const email = "auraprint@mail.ru";
      const subject = "Вопрос от клиента";
      const body = "Здравствуйте! Хочу узнать о ваших услугах.";

      const mailtoLink = `mailto:${email}?subject=${encodeURIComponent(
        subject
      )}&body=${encodeURIComponent(body)}`;
      window.location.href = mailtoLink;
    };

    // Функция для звонка
    const makeCall = () => {
      const phoneNumber = "+79955054001";
      window.location.href = `tel:${phoneNumber}`;
    };

    // Функция для открытия карт
    const openMap = () => {
      const address = "г. Москва, пр-кт Волгоградский 32к31";
      const yandexMapsUrl = `https://yandex.ru/maps/?text=${encodeURIComponent(address)}`;
      window.open(yandexMapsUrl, "_blank");
    };

    const logout = () => {
      store.commit("SET_ADMIN", false);
      router.push("/");
    };

    return {
      isAdmin,
      logout,
      sendEmail,
      makeCall,
      openMap,
    };
  },
};
</script>

<style>
/* Стили остаются такими же */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  line-height: 1.6;
  color: #333;
}

.navbar {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 1rem 0;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.nav-container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 2rem;
}

.nav-logo {
  text-decoration: none;
  color: white;
}

.logo-text {
  font-size: 2rem;
  font-weight: bold;
  display: block;
}

.logo-subtitle {
  font-size: 0.8rem;
  opacity: 0.9;
}

.nav-menu {
  display: flex;
  gap: 2rem;
}

.nav-link {
  color: white;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border-radius: 5px;
  transition: background-color 0.3s;
  border: none;
  background: none;
  cursor: pointer;
  font-size: 1rem;
}

.nav-link:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

.admin-link {
  background-color: rgba(255, 255, 255, 0.2);
  font-weight: bold;
}

.main-content {
  min-height: calc(100vh - 200px);
}

.footer {
  background: #2c3e50;
  color: white;
  padding: 2rem 0 0;
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
  padding: 0 2rem;
}

.footer-section h3,
.footer-section h4 {
  margin-bottom: 1rem;
}

.footer-section p {
  margin-bottom: 0.5rem;
  transition: all 0.3s ease;
}

/* Стили для кликабельных контактов в футере */
.clickable {
  cursor: pointer;
  padding: 0.25rem 0;
  border-radius: 4px;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.clickable:hover {
  color: #667eea;
  background-color: rgba(255, 255, 255, 0.1);
  padding-left: 0.5rem;
}

.clickable:active {
  transform: translateY(1px);
}

.footer-bottom {
  text-align: center;
  padding: 1rem;
  margin-top: 2rem;
  border-top: 1px solid #34495e;
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 1rem;
  transition: all 0.3s;
  text-decoration: none;
  display: inline-block;
}

.btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);
}

.btn-danger {
  background: #e74c3c;
  color: white;
}

.btn-danger:hover {
  background: #c0392b;
}

.card {
  background: white;
  border-radius: 10px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  transition: transform 0.3s;
}

.card:hover {
  transform: translateY(-5px);
}

.grid {
  display: grid;
  gap: 2rem;
}

.grid-3 {
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
}

.grid-2 {
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.text-center {
  text-align: center;
}

.mb-4 {
  margin-bottom: 2rem;
}

/* Дополнительные глобальные стили для изображений */
.image-placeholder {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: bold;
}

.img-error {
  background: #f8f9fa;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6c757d;
}

/* Улучшенные карточки */
.card {
  background: white;
  border-radius: 15px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.08);
  overflow: hidden;
  transition: all 0.3s ease;
  border: 1px solid #f0f0f0;
}

.card:hover {
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.15);
  transform: translateY(-2px);
}

/* Улучшенные кнопки */
.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
  transition: all 0.3s ease;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}

.btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.btn-primary:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(102, 126, 234, 0.3);
}

/* Улучшенная сетка */
.grid {
  display: grid;
  gap: 2rem;
}

.grid-3 {
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
}

.grid-2 {
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
}

/* Адаптивность */
@media (max-width: 768px) {
  .grid-3,
  .grid-2 {
    grid-template-columns: 1fr;
  }

  .container {
    padding: 1rem;
  }

  .footer-content {
    grid-template-columns: 1fr;
    text-align: center;
  }

  .clickable {
    justify-content: center;
  }
}
</style>
