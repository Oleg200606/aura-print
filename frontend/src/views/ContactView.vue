<template>
  <div class="contact">
    <div class="container">
      <div class="header-section">
        <h1 class="text-center">Свяжитесь с нами</h1>
        <p class="text-center subtitle">Мы готовы ответить на все ваши вопросы</p>
      </div>

      <div class="contact-content">
        <div class="contact-info">
          <div class="info-card">
            <h2>Контактная информация</h2>

            <a :href="'mailto:' + EMAIL_ADDRESS" class="contact-item">
              <div class="icon-wrapper">
                <span class="icon">📧</span>
              </div>
              <div class="contact-text">
                <h3>Email</h3>
                <p class="clickable">auraprint@mail.ru</p>
              </div>
            </a>

            <a class="contact-item" :href="'tel:' + PHONE_NUMBER">
              <div class="icon-wrapper">
                <span class="icon">📞</span>
              </div>
              <div class="contact-text">
                <h3>Телефон</h3>
                <p class="clickable">+7 (995) 505-40-01</p>
              </div>
            </a>

            <a class="contact-item" :href="MAP_URL">
              <div class="icon-wrapper">
                <span class="icon">📍</span>
              </div>
              <div class="contact-text">
                <h3>Адрес</h3>
                <p class="clickable">г. Москва, пр-кт Волгоградский 32к31</p>
              </div>
            </a>

            <div class="contact-item">
              <div class="icon-wrapper">
                <span class="icon">🕒</span>
              </div>
              <div class="contact-text">
                <h3>Режим работы</h3>
                <p>Пн-Пт: 9:00 - 18:00</p>
                <p>Сб-Вс: выходные</p>
              </div>
            </div>
          </div>
        </div>

        <div class="contact-form">
          <div class="form-card">
            <div class="form-header">
              <h2>Отправить сообщение</h2>
              <p>Заполните форму и мы свяжемся с вами в течение часа</p>
            </div>

            <form @submit.prevent="submitForm" class="form-body">
              <div class="form-group">
                <label>Имя *</label>
                <input
                  v-model="form.name"
                  type="text"
                  required
                  placeholder="Введите ваше имя"
                  class="form-input"
                  :class="{ error: errors.name }"
                />
                <span v-if="errors.name" class="error-message">{{ errors.name }}</span>
              </div>

              <div class="form-row">
                <div class="form-group">
                  <label>Email *</label>
                  <input
                    v-model="form.email"
                    type="email"
                    required
                    placeholder="your@email.com"
                    class="form-input"
                    :class="{ error: errors.email }"
                  />
                  <span v-if="errors.email" class="error-message">{{ errors.email }}</span>
                </div>

                <div class="form-group">
                  <label>Телефон</label>
                  <input
                    v-model="form.phone"
                    type="tel"
                    placeholder="+7 (___) ___-__-__"
                    class="form-input"
                  />
                </div>
              </div>

              <div class="form-group">
                <label>Сообщение *</label>
                <textarea
                  v-model="form.message"
                  required
                  placeholder="Опишите ваш вопрос или проект..."
                  rows="5"
                  class="form-textarea"
                  :class="{ error: errors.message }"
                ></textarea>
                <span v-if="errors.message" class="error-message">{{ errors.message }}</span>
              </div>

              <!-- Уведомления -->
              <div v-if="notification.message" :class="['notification', notification.type]">
                {{ notification.message }}
              </div>

              <button
                type="submit"
                class="btn btn-primary"
                :class="{ 'btn-loading': isLoading }"
                :disabled="isLoading"
              >
                <span v-if="!isLoading">Отправить сообщение</span>
                <span v-else>Отправка...</span>
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const form = ref({
  name: "",
  email: "",
  phone: "",
  message: "",
});

const isLoading = ref(false);
const errors = ref({});
const notification = ref({ message: "", type: "" });

const PHONE_NUMBER = "+79955054001"; // Убираем все нецифровые символы
const EMAIL_ADDRESS = "auraprint@mail.ru";

const OFFICE_ADDRESS = "г. Москва, пр-кт Волгоградский 32к31";
const MAP_URL = `https://yandex.ru/maps/?text=${encodeURIComponent(OFFICE_ADDRESS)}`;

const showNotification = (message, type = "success") => {
  notification.value = { message, type };
  setTimeout(() => {
    notification.value = { message: "", type: "" };
  }, 5000);
};

const validateForm = () => {
  errors.value = {};

  if (!form.value.name.trim()) {
    errors.value.name = "Имя обязательно для заполнения";
  }

  if (!form.value.email.trim()) {
    errors.value.email = "Email обязателен для заполнения";
  } else if (!/^\S+@\S+\.\S+$/.test(form.value.email)) {
    errors.value.email = "Введите корректный email";
  }

  if (!form.value.message.trim()) {
    errors.value.message = "Сообщение обязательно для заполнения";
  }

  return Object.keys(errors.value).length === 0;
};

const submitForm = async () => {
  if (!validateForm()) {
    showNotification("Пожалуйста, исправьте ошибки в форме", "error");
    return;
  }

  isLoading.value = true;

  try {
    const response = await fetch("/api/contact", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(form.value),
    });

    const result = await response.json();

    if (result.success) {
      showNotification(result.message, "success");
      form.value = { name: "", email: "", phone: "", message: "" };
    } else {
      showNotification(result.message || "Произошла ошибка при отправке", "error");
    }
  } catch (error) {
    console.error("Ошибка:", error);
    showNotification("Ошибка соединения с сервером", "error");
  } finally {
    isLoading.value = false;
  }
};
</script>

<style scoped>
.contact {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
  padding: 4rem 0;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

.header-section {
  margin-bottom: 4rem;
}

.header-section h1 {
  font-size: 3rem;
  font-weight: 700;
  color: white;
  margin-bottom: 1rem;
  text-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.subtitle {
  font-size: 1.2rem;
  color: rgba(255, 255, 255, 0.9);
}

.contact-content {
  display: grid;
  grid-template-columns: 1fr 1.2fr;
  gap: 3rem;
  align-items: start;
}

/* Стили для блока контактной информации */
.contact-info .info-card {
  background: white;
  padding: 2.5rem;
  border-radius: 20px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(10px);
}

.contact-info h2 {
  margin-bottom: 2rem;
  color: #2c3e50;
  font-size: 1.5rem;
  font-weight: 600;
}

.contact-item {
  display: flex;
  align-items: flex-start;
  margin-bottom: 2rem;
  padding: 1rem;
  border-radius: 12px;
  transition: all 0.3s ease;
  cursor: pointer;
}

.contact-item:hover {
  background: #f8f9fa;
  transform: translateX(5px);
}

.contact-item:active {
  transform: translateX(2px);
}

.clickable {
  color: #667eea;
  font-weight: 500;
  transition: color 0.3s ease;
  &:hover {
    padding-left: 0;
  }
}

.contact-item:hover .clickable {
  color: #764ba2;
}

.icon-wrapper {
  margin-right: 1rem;
  flex-shrink: 0;
}

.icon {
  font-size: 1.5rem;
  display: block;
}

.contact-text h3 {
  margin-bottom: 0.5rem;
  color: #34495e;
  font-size: 1.1rem;
  font-weight: 600;
}

.contact-text p {
  color: #7f8c8d;
  margin: 0;
  line-height: 1.5;
}

/* Стили для формы */
.contact-form .form-card {
  background: white;
  padding: 2.5rem;
  border-radius: 20px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
}

.form-header {
  margin-bottom: 2rem;
  text-align: center;
}

.form-header h2 {
  color: #2c3e50;
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
}

.form-header p {
  color: #7f8c8d;
  margin: 0;
}

.form-body {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-group label {
  margin-bottom: 0.5rem;
  color: #2c3e50;
  font-weight: 500;
  font-size: 0.9rem;
}

.form-input,
.form-textarea {
  padding: 1rem;
  border: 2px solid #e9ecef;
  border-radius: 10px;
  font-size: 1rem;
  transition: all 0.3s ease;
  background: #f8f9fa;
}

.form-input:focus,
.form-textarea:focus {
  outline: none;
  border-color: #667eea;
  background: white;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.form-textarea {
  resize: vertical;
  min-height: 120px;
  font-family: inherit;
}

.btn {
  padding: 1rem 2rem;
  border: none;
  border-radius: 10px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  text-align: center;
  margin-top: 1rem;
}

.btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.btn-primary:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 10px 20px rgba(102, 126, 234, 0.3);
}

.btn-primary:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.btn-loading {
  position: relative;
  color: transparent;
}

.btn-loading::after {
  content: "";
  position: absolute;
  width: 20px;
  height: 20px;
  top: 50%;
  left: 50%;
  margin-left: -10px;
  margin-top: -10px;
  border: 2px solid transparent;
  border-top: 2px solid white;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}

/* Адаптивность */
@media (max-width: 768px) {
  .contact-content {
    grid-template-columns: 1fr;
    gap: 2rem;
  }

  .header-section h1 {
    font-size: 2rem;
  }

  .form-row {
    grid-template-columns: 1fr;
  }

  .contact-info .info-card,
  .contact-form .form-card {
    padding: 1.5rem;
  }

  .contact-item {
    padding: 0.75rem;
  }
}

@media (max-width: 480px) {
  .contact {
    padding: 2rem 0;
  }

  .header-section h1 {
    font-size: 1.75rem;
  }

  .contact-info .info-card,
  .contact-form .form-card {
    padding: 1rem;
  }
}

.error {
  border-color: #e74c3c !important;
  background-color: #fdf2f2 !important;
}

.error-message {
  color: #e74c3c;
  font-size: 0.8rem;
  margin-top: 0.25rem;
  display: block;
}

.notification {
  padding: 1rem;
  border-radius: 10px;
  margin: 1rem 0;
  text-align: center;
  font-weight: 500;
}

.notification.success {
  background-color: #d4edda;
  color: #155724;
  border: 1px solid #c3e6cb;
}

.notification.error {
  background-color: #f8d7da;
  color: #721c24;
  border: 1px solid #f5c6cb;
}
</style>
