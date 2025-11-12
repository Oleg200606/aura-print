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
          <router-link v-if="!isAdmin" to="/admin" class="nav-link admin-link">Админ</router-link>
          <button v-else @click="logout" class="nav-link admin-link">Выйти</button>
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
          <p>Email: info@auraprint.ru</p>
          <p>Телефон: +7 (999) 123-45-67</p>
        </div>
        <div class="footer-section">
          <h4>Услуги</h4>
          <p>Футболки</p>
          <p>Кружки</p>
          <p>Корпоративный мерч</p>
        </div>
      </div>
      <div class="footer-bottom">
        <p>&copy; 2024 AuraPrint. Все права защищены.</p>
      </div>
    </footer>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'

export default {
  name: 'App',
  setup() {
    const store = useStore()
    const router = useRouter()

    const isAdmin = computed(() => store.state.isAdmin)

    const logout = () => {
      store.commit('SET_ADMIN', false)
      router.push('/')
    }

    return {
      isAdmin,
      logout
    }
  }
}
</script>

<style>
/* Стили остаются такими же */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
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
</style>