<template>
  <div class="catalog">
    <div class="container">
      <div class="header-section">
        <h1 class="text-center">Наша продукция</h1>
        <p class="text-center subtitle">Широкий ассортимент качественной печатной продукции</p>
      </div>

      <!-- Admin Controls -->
      <div v-if="$store.state.isAdmin" class="admin-controls mb-4">
        <button @click="showProductForm = true" class="btn btn-primary">
          ➕ Добавить товар
        </button>
      </div>

      <!-- Product Form Modal -->
      <div v-if="showProductForm" class="modal">
        <div class="modal-content">
          <h3>Добавить новый товар</h3>
          <form @submit.prevent="addProduct">
            <div class="form-group">
              <label>Название:</label>
              <input v-model="newProduct.name" type="text" required class="form-input" />
            </div>
            <div class="form-group">
              <label>Описание:</label>
              <textarea v-model="newProduct.description" required class="form-input"></textarea>
            </div>
            <div class="form-group">
              <label>Изображение товара:</label>
              <div class="file-upload-area" @click="triggerFileInput" @drop="handleDrop" @dragover.prevent
                @dragenter.prevent :class="{ 'drag-over': isDragOver }">
                <div class="upload-placeholder">
                  <div class="upload-icon">📁</div>
                  <p v-if="!selectedFile">Перетащите изображение сюда или кликните для выбора</p>
                  <p v-else class="file-selected">Выбран файл: {{ selectedFile.name }}</p>
                  <small>Поддерживаемые форматы: JPG, PNG, GIF</small>
                </div>
                <input type="file" ref="fileInput" @change="handleFileSelect" accept="image/*" style="display: none" />
              </div>
            </div>
            <div class="form-group">
              <label>Категория:</label>
              <input v-model="newProduct.category" type="text" required class="form-input" />
            </div>
            <div class="form-actions">
              <button type="button" @click="cancelProductForm" class="btn btn-secondary">
                Отмена
              </button>
              <button type="submit" class="btn btn-primary" :disabled="uploading">
                {{ uploading ? "Загрузка..." : "Добавить товар" }}
              </button>
            </div>
          </form>
        </div>
      </div>

      <!-- Products Grid -->
      <div class="products-grid">
        <div v-for="product in products" :key="product.id" class="product-card card">
          <div class="product-image-container">
            <img :src="getImageUrl(product.image_url)" :alt="product.name" class="product-image"
              @error="handleImageError" />
            <div class="product-overlay">
              <div class="product-actions">
                <button class="btn btn-sm btn-primary">Подробнее</button>
              </div>
            </div>
          </div>
          <div class="product-content">
            <h3>{{ product.name }}</h3>
            <p>{{ product.description }}</p>
            <div class="product-category">{{ product.category }}</div>
            <button v-if="$store.state.isAdmin" @click="deleteProduct(product.id)" class="btn btn-danger btn-sm">
              🗑️ Удалить
            </button>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="products.length === 0" class="empty-state">
        <div class="empty-icon">📦</div>
        <h3>Пока нет товаров</h3>
        <p>Администратор может добавить товары через панель управления</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref, onMounted } from "vue";
import { useStore } from "vuex";

const store = useStore();
const products = computed(() => store.state.products);
const showProductForm = ref(false);
const uploading = ref(false);
const selectedFile = ref(null);
const isDragOver = ref(false);
const fileInput = ref(null);

import { DEFAULT_IMAGE_URL } from "@/data";

const newProduct = ref({
  name: "",
  description: "",
  image_url: "",
  category: "",
});

onMounted(() => {
  store.dispatch("fetchProducts");
});

const getImageUrl = (imagePath) => {
  console.log(imagePath);
  if (!imagePath) return DEFAULT_IMAGE_URL;
  if (imagePath.startsWith("http")) return imagePath;
  return import.meta.env.BACKEND_HOST + `/api${imagePath}`;
};

/**
 *
 * @param event {Event}
 */
const handleImageError = (event) => {
  console.error("Error while loading image " + event.target.src);
  if (event.target.src === DEFAULT_IMAGE_URL) {
    return;
  }
  console.error("Setting default image", DEFAULT_IMAGE_URL);
  event.target.src = DEFAULT_IMAGE_URL;
};

const triggerFileInput = () => {
  fileInput.value.click();
};

const handleFileSelect = (event) => {
  const file = event.target.files[0];
  if (file && file.type.startsWith("image/")) {
    selectedFile.value = file;
  } else {
    alert("Пожалуйста, выберите файл изображения (JPG, PNG, GIF)");
  }
};

const handleDrop = (event) => {
  event.preventDefault();
  isDragOver.value = false;

  const files = event.dataTransfer.files;
  if (files.length > 0) {
    const file = files[0];
    if (file.type.startsWith("image/")) {
      selectedFile.value = file;
    } else {
      alert("Пожалуйста, перетащите файл изображения (JPG, PNG, GIF)");
    }
  }
};

const uploadImage = async () => {
  if (!selectedFile.value) {
    alert("Пожалуйста, выберите изображение для товара");
    return null;
  }

  uploading.value = true;
  const formData = new FormData();
  formData.append("image", selectedFile.value);
  formData.append("folder", "products");

  try {
    const response = await fetch("/api/admin/upload/image", {
      method: "POST",
      body: formData,
    });

    const result = await response.json();

    if (result.success) {
      return result.data.url;
    } else {
      alert("Ошибка загрузки изображения: " + result.message);
      return null;
    }
  } catch (error) {
    console.error("Upload failed:", error);
    alert("Ошибка загрузки изображения");
    return null;
  } finally {
    uploading.value = false;
  }
};

const addProduct = async () => {
  if (!selectedFile.value) {
    alert("Пожалуйста, выберите изображение для товара");
    return;
  }

  // Сначала загружаем изображение
  const imageUrl = await uploadImage();

  if (!imageUrl) {
    return; // Ошибка уже обработана в uploadImage
  }

  // Затем добавляем товар с полученным URL изображения
  try {
    const productData = {
      ...newProduct.value,
      image_url: imageUrl,
    };

    const response = await fetch("/api/admin/products", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(productData),
    });

    if (response.ok) {
      const product = await response.json();
      store.commit("ADD_PRODUCT", product);
      cancelProductForm();
    } else {
      alert("Ошибка при добавлении товара");
    }
  } catch (error) {
    console.error("Failed to add product:", error);
    alert("Ошибка при добавлении товара");
  }
};

const cancelProductForm = () => {
  showProductForm.value = false;
  newProduct.value = { name: "", description: "", image_url: "", category: "" };
  selectedFile.value = null;
  isDragOver.value = false;
};

const deleteProduct = async (productId) => {
  if (confirm("Вы уверены, что хотите удалить этот товар?")) {
    try {
      const response = await fetch(`/api/admin/products/${productId}`, {
        method: "DELETE",
      });

      if (response.ok) {
        store.commit("DELETE_PRODUCT", productId);
      }
    } catch (error) {
      console.error("Failed to delete product:", error);
    }
  }
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

.catalog {
  background: linear-gradient(135deg, var(--light) 0%, #FFFFFF 100%);
  min-height: 100vh;
  padding: 6rem 0 2rem;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

.header-section {
  margin-bottom: 3rem;
}

.header-section h1 {
  font-size: 3rem;
  font-weight: 700;
  color: var(--dark);
  margin-bottom: 1rem;
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.subtitle {
  font-size: 1.2rem;
  color: var(--text-light);
}

.admin-controls {
  display: flex;
  justify-content: center;
  margin-bottom: 2rem;
}

.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 2rem;
}

.product-card {
  display: flex;
  flex-direction: column;
  height: 100%;
  transition: all 0.3s ease;
  background: white;
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(0, 0, 0, 0.05);
}

.product-card:hover {
  transform: translateY(-10px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
}

.product-image-container {
  width: 100%;
  height: 250px;
  overflow: hidden;
  background: var(--light);
  position: relative;
}

.product-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.5s ease;
}

.product-card:hover .product-image {
  transform: scale(1.1);
}

.product-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(to bottom, rgba(0, 0, 0, 0.1), rgba(0, 0, 0, 0.3));
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.product-card:hover .product-overlay {
  opacity: 1;
}

.product-actions {
  transform: translateY(20px);
  transition: transform 0.3s ease;
}

.product-card:hover .product-actions {
  transform: translateY(0);
}

.product-content {
  padding: 1.5rem;
  flex-grow: 1;
  display: flex;
  flex-direction: column;
}

.product-content h3 {
  margin-bottom: 1rem;
  color: var(--dark);
  font-size: 1.25rem;
  font-weight: 600;
}

.product-content p {
  flex-grow: 1;
  margin-bottom: 1rem;
  color: var(--text-light);
  line-height: 1.5;
}

.product-category {
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 500;
  margin-bottom: 1rem;
  align-self: flex-start;
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 10px;
  font-size: 0.9rem;
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

/* Стили для области загрузки файлов */
.file-upload-area {
  border: 2px dashed #dee2e6;
  border-radius: 8px;
  padding: 2rem;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  background: var(--light);
}

.file-upload-area:hover {
  border-color: var(--primary);
  background: #f0f2ff;
}

.file-upload-area.drag-over {
  border-color: var(--primary);
  background: #e6ebff;
}

.upload-placeholder .upload-icon {
  font-size: 2rem;
  margin-bottom: 1rem;
}

.file-selected {
  color: var(--secondary);
  font-weight: 600;
}

.form-actions {
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
  margin-top: 2rem;
}

.empty-state {
  text-align: center;
  padding: 4rem 2rem;
  background: white;
  border-radius: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.08);
}

.empty-icon {
  font-size: 4rem;
  margin-bottom: 1.5rem;
}

.empty-state h3 {
  margin-bottom: 1rem;
  color: var(--dark);
  font-size: 1.5rem;
}

.empty-state p {
  color: var(--text-light);
  line-height: 1.5;
}

/* Адаптивность */
@media (max-width: 768px) {
  .products-grid {
    grid-template-columns: 1fr;
  }

  .modal-content {
    margin: 1rem;
    padding: 1.5rem;
  }

  .file-upload-area {
    padding: 1.5rem;
  }

  .header-section h1 {
    font-size: 2rem;
  }
}

@media (max-width: 480px) {
  .catalog {
    padding: 5rem 0 2rem;
  }

  .header-section h1 {
    font-size: 1.75rem;
  }
}
</style>