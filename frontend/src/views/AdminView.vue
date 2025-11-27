<template>
  <div class="admin">
    <div class="container">
      <!-- Login Form -->
      <div v-if="!$store.state.isAdmin" class="login-section">
        <div class="login-card card">
          <h2>Вход для администратора</h2>
          <form @submit.prevent="login">
            <div class="form-group">
              <label>Логин:</label>
              <input v-model="credentials.username" type="text" required autocomplete="username" />
            </div>
            <div class="form-group">
              <label>Пароль:</label>
              <input v-model="credentials.password" type="password" required autocomplete="current-password" />
            </div>
            <button type="submit" class="btn btn-primary">Войти</button>
            <p v-if="loginError" class="error-message">{{ loginError }}</p>
          </form>
        </div>
      </div>

      <!-- Admin Panel -->
      <div v-else class="admin-panel">
        <div class="admin-header">
          <h1>Панель администратора</h1>
          <p>Управление контентом сайта Купи-Принт.ру</p>
        </div>

        <!-- News Management -->
        <section class="admin-section">
          <div class="section-header">
            <h2>Управление новостями</h2>
            <button @click="showNewsForm = true" class="btn btn-primary">
              📝 Добавить новость
            </button>
          </div>

          <!-- News Form Modal -->
          <div v-if="showNewsForm" class="modal">
            <div class="modal-content">
              <h3>Добавить новость</h3>
              <form @submit.prevent="addNews">
                <div class="form-group">
                  <label>Заголовок:</label>
                  <input v-model="newNews.title" type="text" required class="form-input" />
                </div>
                <div class="form-group">
                  <label>Содержание:</label>
                  <textarea v-model="newNews.content" required class="form-input" rows="4"></textarea>
                </div>
                <div class="form-group">
                  <label>Изображение:</label>
                  <input type="file" @change="handleNewsImageSelect" accept="image/*" class="form-input" />
                  <small>Или укажите ссылку:</small>
                  <input v-model="newNews.image_url" type="url" placeholder="https://example.com/image.jpg"
                    class="form-input mt-1" />
                </div>
                <div class="form-actions">
                  <button type="button" @click="showNewsForm = false" class="btn btn-secondary">
                    Отмена
                  </button>
                  <button type="submit" class="btn btn-primary">Добавить</button>
                </div>
              </form>
            </div>
          </div>

          <!-- News List -->
          <div class="news-admin-grid">
            <div v-for="item in news" :key="item.id" class="news-admin-card card">
              <div class="news-admin-image">
                <img :src="item.imageUrl()" :alt="item.title" />
              </div>
              <div class="news-admin-content">
                <h3>{{ item.title }}</h3>
                <p>{{ item.content }}</p>
                <div class="news-admin-meta">
                  <small>{{ formatDate(item.created_at) }}</small>
                  <button @click="deleteNews(item.id)" class="btn btn-danger btn-sm">
                    🗑️ Удалить
                  </button>
                </div>
              </div>
            </div>
          </div>

          <div v-if="news.length === 0" class="empty-state">
            <div class="empty-icon">📰</div>
            <h3>Новостей пока нет</h3>
            <p>Добавьте первую новость, используя кнопку выше</p>
          </div>
        </section>

        <!-- Products Management -->
        <section class="admin-section">
          <div class="section-header">
            <h2>Управление товарами</h2>
            <p>
              Управление товарами доступно на странице
              <router-link to="/catalog" class="admin-link">Каталог</router-link>
            </p>
          </div>
        </section>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { useStore } from "vuex";

const store = useStore();
const credentials = ref({ username: "", password: "" });
const loginError = ref("");
const showNewsForm = ref(false);
const newsImageFile = ref(null);

const newNews = ref({
  title: "",
  content: "",
  image_url: "",
});

const news = computed(() => store.state.news);

onMounted(() => {
  store.dispatch("fetchNews");
});

const login = async () => {
  const onError = (error) => {
    console.error("❌ Ошибка логина:", error);
  };

  console.log("🔄 Sending login request:", credentials.value);

  // Валидация на клиенте
  if (!credentials.value.username || !credentials.value.password) {
    loginError.value = "Пожалуйста, заполните все поля";
    onError("не все поля заполнены");
    return;
  }

  try {
    const result = await store.dispatch("login", credentials.value);
    if (!result.success) {
      loginError.value = result.error || "Ошибка авторизации";
      console.log("❌ Login failed:", result.error);
      onError("не удалось авторизоваться: " + result.error);
    } else {
      console.log("✅ Login successful!");
      loginError.value = "";
    }
  } catch (error) {
    console.error("Login error:", error);
    loginError.value = "Ошибка соединения с сервером";
    onError(error);
  }
};

const handleNewsImageSelect = (event) => {
  newsImageFile.value = event.target.files[0];
};

const uploadNewsImage = async () => {
  if (!newsImageFile.value) return null;

  const formData = new FormData();
  formData.append("image", newsImageFile.value);
  formData.append("folder", "news");

  try {
    const response = await fetch("/api/admin/upload/image", {
      method: "POST",
      body: formData,
    });

    const result = await response.json();
    return result.success ? result.data.url : null;
  } catch (error) {
    console.error("News image upload failed:", error);
    return null;
  }
};

const addNews = async () => {
  // Сначала загружаем изображение если есть
  if (newsImageFile.value) {
    const imageUrl = await uploadNewsImage();
    if (imageUrl) {
      newNews.value.image_url = imageUrl;
    }
  }

  // Затем создаем новость
  try {
    const response = await fetch("/api/admin/news", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newNews.value),
    });

    if (response.ok) {
      const newsItem = await response.json();
      store.commit("ADD_NEWS", newsItem);
      showNewsForm.value = false;
      newNews.value = { title: "", content: "", image_url: "" };
      newsImageFile.value = null;
    } else {
      console.error("Failed to add news:", response.statusText);
    }
  } catch (error) {
    console.error("Failed to add news:", error);
  }
};

const deleteNews = async (newsId) => {
  if (confirm("Вы уверены, что хотите удалить эту новость?")) {
    try {
      const response = await fetch(`/api/admin/news/${newsId}`, {
        method: "DELETE",
      });

      if (response.ok) {
        store.commit("DELETE_NEWS", newsId);
      } else {
        console.error("Failed to delete news:", response.statusText);
      }
    } catch (error) {
      console.error("Failed to delete news:", error);
    }
  }
};

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString("ru-RU");
};
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

.admin {
  background: linear-gradient(135deg, var(--light) 0%, #FFFFFF 100%);
  min-height: 100vh;
  padding: 6rem 0 2rem;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

.login-section {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 60vh;
}

.login-card {
  width: 100%;
  max-width: 400px;
  padding: 2.5rem;
  background: white;
  border-radius: 20px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
}

.login-card h2 {
  text-align: center;
  margin-bottom: 2rem;
  color: var(--dark);
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 600;
  color: var(--dark);
}

.form-input {
  width: 100%;
  padding: 0.75rem;
  border: 2px solid #e9ecef;
  border-radius: 8px;
  font-size: 1rem;
  transition: border-color 0.3s ease;
  background: var(--light);
}

.form-input:focus {
  outline: none;
  border-color: var(--primary);
  box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.1);
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 10px;
  font-size: 1rem;
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
  box-shadow: 0 4px 15px rgba(255, 107, 53, 0.3);
}

.btn-primary:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(255, 107, 53, 0.4);
}

.btn-secondary {
  background: transparent;
  color: var(--text-light);
  border: 2px solid var(--text-light);
}

.btn-secondary:hover {
  background: var(--text-light);
  color: white;
}

.btn-danger {
  background: linear-gradient(135deg, #e74c3c, #c0392b);
  color: white;
}

.btn-sm {
  padding: 0.5rem 1rem;
  font-size: 0.8rem;
}

.error-message {
  color: #e74c3c;
  margin-top: 1rem;
  text-align: center;
  font-weight: 500;
}

.admin-panel {
  max-width: 1000px;
  margin: 0 auto;
}

.admin-header {
  text-align: center;
  margin-bottom: 3rem;
}

.admin-header h1 {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  color: var(--dark);
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.admin-header p {
  font-size: 1.1rem;
  color: var(--text-light);
}

.admin-section {
  background: white;
  padding: 2.5rem;
  border-radius: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.08);
  margin-bottom: 2rem;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  flex-wrap: wrap;
  gap: 1rem;
}

.section-header h2 {
  color: var(--dark);
  font-size: 1.5rem;
  font-weight: 600;
}

.admin-link {
  color: var(--primary);
  text-decoration: none;
  font-weight: 600;
}

.admin-link:hover {
  text-decoration: underline;
}

.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  padding: 2rem;
  border-radius: 20px;
  width: 90%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
}

.modal-content h3 {
  margin-bottom: 1.5rem;
  color: var(--dark);
  text-align: center;
}

.form-actions {
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
  margin-top: 2rem;
}

.news-admin-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.5rem;
}

.news-admin-card {
  display: flex;
  gap: 1.5rem;
  padding: 1.5rem;
  transition: transform 0.3s ease;
}

.news-admin-card:hover {
  transform: translateY(-2px);
}

.news-admin-image {
  flex-shrink: 0;
  width: 120px;
  height: 120px;
  border-radius: 10px;
  overflow: hidden;
}

.news-admin-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.news-admin-content {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.news-admin-content h3 {
  margin-bottom: 0.5rem;
  color: var(--dark);
  font-size: 1.2rem;
}

.news-admin-content p {
  flex: 1;
  margin-bottom: 1rem;
  color: var(--text-light);
  line-height: 1.5;
}

.news-admin-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.news-admin-meta small {
  color: var(--text-light);
}

.empty-state {
  text-align: center;
  padding: 3rem 2rem;
  color: var(--text-light);
}

.empty-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.empty-state h3 {
  margin-bottom: 1rem;
  color: var(--dark);
}

/* Адаптивность */
@media (max-width: 768px) {
  .admin {
    padding: 5rem 0 2rem;
  }

  .section-header {
    flex-direction: column;
    align-items: stretch;
  }

  .news-admin-card {
    flex-direction: column;
  }

  .news-admin-image {
    width: 100%;
    height: 200px;
  }

  .news-admin-meta {
    flex-direction: column;
    gap: 1rem;
    align-items: flex-start;
  }

  .admin-section {
    padding: 1.5rem;
  }

  .login-card {
    margin: 1rem;
    padding: 1.5rem;
  }
}

@media (max-width: 480px) {
  .admin-header h1 {
    font-size: 2rem;
  }

  .modal-content {
    margin: 1rem;
    padding: 1.5rem;
  }
}
</style>