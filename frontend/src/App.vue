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
          <span class="logo-text">Купи-Принт</span>
          <span class="logo-dot">.ру</span>
          <span class="logo-subtitle">Профессиональная печать</span>
        </router-link>
        <div class="nav-menu">
          <router-link to="/" class="nav-link" @click="closeMobileMenu">Главная</router-link>
          <router-link to="/catalog" class="nav-link" @click="closeMobileMenu">Каталог</router-link>
          <router-link to="/about" class="nav-link" @click="closeMobileMenu">О нас</router-link>
          <router-link to="/contact" class="nav-link" @click="closeMobileMenu">Контакты</router-link>
          <router-link v-if="isAdmin" to="/admin" class="nav-link" @click="closeMobileMenu">Админ</router-link>
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
          <h3>Купи-Принт.ру</h3>
          <p>Профессиональная печать и корпоративный мерч</p>
          <div class="social-links">
            <a href="#" class="social-link">📘</a>
            <a href="#" class="social-link">📷</a>
            <a href="#" class="social-link">💬</a>
          </div>
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
          <p>Разработка дизайна</p>
        </div>
      </div>
      <div class="footer-bottom">
        <p>&copy; 2025 Купи-Принт.ру. Все права защищены.</p>
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

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html {
  font-size: 16px;
}

body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  line-height: 1.6;
  color: var(--text-dark);
  overflow-x: hidden;
  background: var(--light);
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
  background: rgba(26, 31, 43, 0.95);
  backdrop-filter: blur(10px);
  padding: 1rem 0;
  box-shadow: 0 2px 20px rgba(0, 0, 0, 0.1);
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
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
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.logo-text {
  font-size: 1.5rem;
  font-weight: bold;
  background: linear-gradient(135deg, var(--primary), var(--accent));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.logo-dot {
  font-size: 1.5rem;
  font-weight: bold;
  color: var(--secondary);
}

.logo-subtitle {
  font-size: 0.7rem;
  opacity: 0.8;
  margin-left: 0.5rem;
  font-weight: 300;
}

.nav-menu {
  display: flex;
  gap: 1.5rem;
  align-items: center;
}

.nav-link {
  color: white;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border-radius: 8px;
  transition: all 0.3s ease;
  border: none;
  background: none;
  cursor: pointer;
  font-size: 0.95rem;
  white-space: nowrap;
  font-weight: 500;
}

.nav-link:hover {
  background: rgba(255, 255, 255, 0.1);
  transform: translateY(-2px);
}

.nav-link.router-link-active {
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  color: white;
}

.admin-link {
  background: rgba(255, 209, 102, 0.2);
  border: 1px solid var(--accent);
  font-weight: bold;
}

.admin-link:hover {
  background: var(--accent);
  color: var(--dark);
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
  padding-top: 80px;
}

/* Footer */
.footer {
  background: var(--dark);
  color: white;
  padding: 3rem 0 0;
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr;
  gap: 2rem;
  padding: 0 1rem 2rem;
}

.footer-section h3 {
  margin-bottom: 1rem;
  color: var(--primary);
  font-size: 1.5rem;
}

.footer-section h4 {
  margin-bottom: 1rem;
  color: var(--secondary);
  font-size: 1.2rem;
}

.footer-section p {
  margin-bottom: 0.5rem;
  transition: all 0.3s ease;
  opacity: 0.9;
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
  color: var(--accent);
  background-color: rgba(255, 255, 255, 0.1);
  padding-left: 0.5rem;
}

.social-links {
  display: flex;
  gap: 1rem;
  margin-top: 1rem;
}

.social-link {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 50%;
  text-decoration: none;
  font-size: 1.2rem;
  transition: all 0.3s ease;
}

.social-link:hover {
  background: var(--primary);
  transform: translateY(-2px);
}

.footer-bottom {
  text-align: center;
  padding: 1.5rem;
  margin-top: 2rem;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
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
    background: linear-gradient(135deg, var(--dark) 0%, #2D3748 100%);
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

  .social-links {
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
  background: linear-gradient(135deg, var(--primary), var(--primary-dark));
  color: white;
}

.btn-primary:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(255, 107, 53, 0.3);
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

  .nav-link:hover {
    transform: none;
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