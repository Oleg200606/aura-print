<template>
  <div id="app">
    <!-- Mobile Menu Button -->
    <div class="mobile-menu-toggle" @click="toggleMobileMenu">
      <span></span>
      <span></span>
      <span></span>
    </div>

    <nav class="navbar" :class="{ 'mobile-open': isMobileMenuOpen }">
      <div class="nav-container">
        <router-link to="/" class="nav-logo" @click="closeMobileMenu">
          <span class="logo-text">AuraPrint</span>
          <span class="logo-subtitle">Маркетинговое оборудование</span>
        </router-link>
        <div class="nav-menu">
          <router-link to="/" class="nav-link" @click="closeMobileMenu">Главная</router-link>
          <router-link to="/catalog" class="nav-link" @click="closeMobileMenu">Каталог</router-link>
          <router-link to="/about" class="nav-link" @click="closeMobileMenu">О нас</router-link>
          <router-link to="/contact" class="nav-link" @click="closeMobileMenu"
            >Контакты</router-link
          >
          <router-link v-if="isAdmin" to="/admin" class="nav-link" @click="closeMobileMenu"
            >Админ</router-link
          >
          <button v-if="isAdmin" @click="logout" class="nav-link admin-link">Выйти</button>
        </div>
      </div>
    </nav>

    <!-- Mobile Menu Overlay -->
    <div v-if="isMobileMenuOpen" class="mobile-overlay" @click="closeMobileMenu"></div>

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
          <p class="clickable" @click="sendEmail">📧 Email: auraprint@mail.ru</p>
          <p class="clickable" @click="makeCall">📞 Телефон: +7 (995) 505-40-01</p>
          <p class="clickable" @click="openMap">📍 Адрес: г. Москва, пр-кт Волгоградский 32к31</p>
        </div>
        <div class="footer-section">
          <h4>Услуги</h4>
          <p>Печать на одежде</p>
          <p>Сувенирная продукция</p>
          <p>Корпоративный мерч</p>
        </div>
      </div>
      <div class="footer-bottom">
        <p>&copy; 2025 AuraPrint. Все права защищены.</p>
        <router-link to="/admin" class="hidden-admin-link">.</router-link>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { computed, ref } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";

const store = useStore();
const router = useRouter();
const isMobileMenuOpen = ref(false);

const isAdmin = computed(() => store.state.isAdmin);

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value;
};

const closeMobileMenu = () => {
  isMobileMenuOpen.value = false;
};

const sendEmail = () => {
  const email = "auraprint@mail.ru";
  const subject = "Вопрос от клиента";
  const body = "Здравствуйте! Хочу узнать о ваших услугах.";

  const mailtoLink = `mailto:${email}?subject=${encodeURIComponent(
    subject
  )}&body=${encodeURIComponent(body)}`;
  window.location.href = mailtoLink;
};

const makeCall = () => {
  const phoneNumber = "+79955054001";
  window.location.href = `tel:${phoneNumber}`;
};

const openMap = () => {
  const address = "г. Москва, пр-кт Волгоградский 32к31";
  const yandexMapsUrl = `https://yandex.ru/maps/?text=${encodeURIComponent(address)}`;
  window.open(yandexMapsUrl, "_blank");
};

const logout = () => {
  store.commit("SET_ADMIN", false);
  router.push("/");
  closeMobileMenu();
};
</script>

<style>
/* Mobile First Styles */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html {
  font-size: 16px;
}

body {
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  line-height: 1.6;
  color: #333;
  overflow-x: hidden;
}

/* Mobile Menu Toggle */
.mobile-menu-toggle {
  display: none;
  flex-direction: column;
  justify-content: space-between;
  width: 30px;
  height: 21px;
  position: fixed;
  top: 1rem;
  right: 1rem;
  z-index: 1001;
  background: transparent;
  border: none;
  cursor: pointer;
}

.mobile-menu-toggle span {
  display: block;
  height: 3px;
  width: 100%;
  background-color: white;
  border-radius: 3px;
  transition: all 0.3s ease;
}

.navbar.mobile-open .mobile-menu-toggle span:nth-child(1) {
  transform: rotate(45deg) translate(6px, 6px);
}

.navbar.mobile-open .mobile-menu-toggle span:nth-child(2) {
  opacity: 0;
}

.navbar.mobile-open .mobile-menu-toggle span:nth-child(3) {
  transform: rotate(-45deg) translate(6px, -6px);
}

/* Navigation */
.navbar {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 1rem 0;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  position: relative;
  z-index: 1000;
}

.nav-container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 1rem;
}

.nav-logo {
  text-decoration: none;
  color: white;
  z-index: 1002;
}

.logo-text {
  font-size: 1.5rem;
  font-weight: bold;
  display: block;
}

.logo-subtitle {
  font-size: 0.7rem;
  opacity: 0.9;
  display: block;
}

.nav-menu {
  display: flex;
  gap: 1rem;
  align-items: center;
}

.nav-link {
  color: white;
  text-decoration: none;
  padding: 0.5rem 0.8rem;
  border-radius: 5px;
  transition: background-color 0.3s;
  border: none;
  background: none;
  cursor: pointer;
  font-size: 0.9rem;
  white-space: nowrap;
}

.nav-link:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

.admin-link {
  background-color: rgba(255, 255, 255, 0.2);
  font-weight: bold;
}

/* Mobile Overlay */
.mobile-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  z-index: 998;
}

.main-content {
  min-height: calc(100vh - 200px);
  padding-top: 0;
}

/* Footer */
.footer {
  background: #2c3e50;
  color: white;
  padding: 2rem 0 0;
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr;
  gap: 2rem;
  padding: 0 1rem;
}

.footer-section h3,
.footer-section h4 {
  margin-bottom: 1rem;
}

.footer-section p {
  margin-bottom: 0.5rem;
  transition: all 0.3s ease;
}

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

.footer-bottom {
  text-align: center;
  padding: 1rem;
  margin-top: 2rem;
  border-top: 1px solid #34495e;
  position: relative;
}

.hidden-admin-link {
  position: absolute;
  bottom: 5px;
  right: 5px;
  width: 8px;
  height: 8px;
  background-color: rgba(255, 255, 255, 0.1);
  border-radius: 50%;
  text-decoration: none;
  color: transparent;
  font-size: 1px;
  opacity: 0.3;
  transition: all 0.3s ease;
  cursor: pointer;
}

/* Mobile Styles */
@media (max-width: 768px) {
  .mobile-menu-toggle {
    display: flex;
  }

  .nav-menu {
    position: fixed;
    top: 0;
    right: -100%;
    width: 80%;
    max-width: 300px;
    height: 100vh;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    flex-direction: column;
    justify-content: flex-start;
    align-items: flex-start;
    padding: 5rem 2rem 2rem;
    transition: right 0.3s ease;
    z-index: 1000;
    gap: 0;
  }

  .navbar.mobile-open .nav-menu {
    right: 0;
  }

  .nav-link {
    width: 100%;
    padding: 1rem 0;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
    font-size: 1.1rem;
  }

  .nav-link:last-child {
    border-bottom: none;
  }

  .footer-content {
    text-align: center;
  }

  .clickable {
    justify-content: center;
  }
}

/* Tablet Styles */
@media (min-width: 769px) and (max-width: 1024px) {
  .nav-container {
    padding: 0 1.5rem;
  }

  .nav-menu {
    gap: 1.5rem;
  }

  .footer-content {
    grid-template-columns: repeat(2, 1fr);
  }
}

/* Desktop Styles */
@media (min-width: 1025px) {
  .nav-container {
    padding: 0 2rem;
  }

  .nav-menu {
    gap: 2rem;
  }

  .footer-content {
    grid-template-columns: repeat(3, 1fr);
  }
}

/* Large Screens */
@media (min-width: 1200px) {
  .nav-container {
    padding: 0 2rem;
  }
}

/* Utility Classes */
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

.card {
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  transition: all 0.3s ease;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

.text-center {
  text-align: center;
}

.mb-4 {
  margin-bottom: 2rem;
}

/* Touch device improvements */
@media (hover: none) and (pointer: coarse) {
  .btn:hover {
    transform: none;
  }

  .card:hover {
    transform: none;
  }

  .clickable:hover {
    padding-left: 0;
    background-color: transparent;
  }
}

/* Reduced motion */
@media (prefers-reduced-motion: reduce) {
  * {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
}
</style>
