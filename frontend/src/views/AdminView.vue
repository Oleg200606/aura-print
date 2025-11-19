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
              <input
                v-model="credentials.password"
                type="password"
                required
                autocomplete="current-password"
              />
            </div>
            <button type="submit" class="btn btn-primary">Войти</button>
            <p v-if="loginError" class="error-message">{{ loginError }}</p>
          </form>
        </div>
      </div>

      <!-- Admin Panel -->
      <div v-else class="admin-panel">
        <h1>Панель администратора</h1>

        <!-- News Management -->
        <section class="admin-section">
          <h2>Управление новостями</h2>
          <button @click="showNewsForm = true" class="btn btn-primary mb-4">
            Добавить новость
          </button>

          <!-- News Form Modal -->
          <div v-if="showNewsForm" class="modal">
            <div class="modal-content">
              <h3>Добавить новость</h3>
              <form @submit.prevent="addNews">
                <div class="form-group">
                  <label>Заголовок:</label>
                  <input v-model="newNews.title" type="text" required />
                </div>
                <div class="form-group">
                  <label>Содержание:</label>
                  <textarea v-model="newNews.content" required></textarea>
                </div>
                <!-- В Admin.vue в форме добавления новостей -->
                <div class="form-group">
                  <label>Изображение:</label>
                  <input
                    type="file"
                    @change="handleNewsImageSelect"
                    accept="image/*"
                    class="form-input"
                  />
                  <small>Или укажите ссылку:</small>
                  <input
                    v-model="newNews.image_url"
                    type="url"
                    placeholder="https://example.com/image.jpg"
                    class="form-input mt-1"
                  />
                </div>
                <div class="form-actions">
                  <button type="button" @click="showNewsForm = false" class="btn">Отмена</button>
                  <button type="submit" class="btn btn-primary">Добавить</button>
                </div>
              </form>
            </div>
          </div>

          <!-- News List -->
          <div class="grid grid-2">
            <div v-for="item in news" :key="item.id" class="news-admin-card card">
              <img :src="item.image" :alt="item.title" class="news-image" />
              <div class="news-content">
                <h3>{{ item.title }}</h3>
                <p>{{ item.content }}</p>
                <small>{{ formatDate(item.created_at) }}</small>
                <button @click="deleteNews(item.id)" class="btn btn-danger btn-sm">Удалить</button>
              </div>
            </div>
          </div>
        </section>

        <!-- Products Management -->
        <section class="admin-section">
          <h2>Управление товарами</h2>
          <p>
            Управление товарами доступно на странице
            <router-link to="/catalog">Каталог</router-link>
          </p>
        </section>
      </div>
    </div>
  </div>
</template>

<!-- Admin.vue - обновленная функция login -->
<script setup>
import { ref, computed, onMounted } from "vue";
import { useStore } from "vuex";

const store = useStore();
const credentials = ref({ username: "", password: "" });
const loginError = ref("");
const showNewsForm = ref(false);
const newsImageFile = ref(null);
import { useSound } from "@vueuse/sound";
import FailedAlert from "../assets/idi_nahui.ogg";

const newNews = ref({
  title: "",
  content: "",
  image_url: "",
});

const news = computed(() => store.state.news);

onMounted(() => {
  store.dispatch("fetchNews");
});

const playFailedAlert = useSound(FailedAlert);

const login = async () => {
  const onError = (error) => {
    console.error("❌ Ошибка логина:", error);
    const { play } = playFailedAlert;
    play();
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
    const response = await fetch("http://localhost:8081/api/admin/upload/image", {
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
      const response = await fetch(`http://localhost:8081/api/admin/news/${newsId}`, {
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
.login-section {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 60vh;
}

.login-card {
  width: 100%;
  max-width: 400px;
  padding: 2rem;
}

.admin-section {
  margin-bottom: 3rem;
  padding: 2rem;
  background: #f8f9fa;
  border-radius: 10px;
}

.news-admin-card {
  display: flex;
  flex-direction: column;
}

.news-admin-card .news-content {
  padding: 1.5rem;
  flex-grow: 1;
  display: flex;
  flex-direction: column;
}

.news-admin-card .btn {
  margin-top: auto;
  align-self: flex-start;
}

.error-message {
  color: #e74c3c;
  margin-top: 1rem;
  text-align: center;
}
</style>
