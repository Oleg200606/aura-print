<template>
  <div class="catalog">
    <div class="container">
      <h1 class="text-center mb-4">Наша продукция</h1>

      <!-- Admin Controls -->
      <div v-if="$store.state.isAdmin" class="admin-controls mb-4">
        <button @click="showProductForm = true" class="btn btn-primary">Добавить товар</button>
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
              <div
                class="file-upload-area"
                @click="triggerFileInput"
                @drop="handleDrop"
                @dragover.prevent
                @dragenter.prevent
                :class="{ 'drag-over': isDragOver }"
              >
                <div class="upload-placeholder">
                  <div class="upload-icon">📁</div>
                  <p v-if="!selectedFile">Перетащите изображение сюда или кликните для выбора</p>
                  <p v-else class="file-selected">Выбран файл: {{ selectedFile.name }}</p>
                  <small>Поддерживаемые форматы: JPG, PNG, GIF</small>
                </div>
                <input
                  type="file"
                  ref="fileInput"
                  @change="handleFileSelect"
                  accept="image/*"
                  style="display: none"
                />
              </div>
            </div>
            <div class="form-group">
              <label>Категория:</label>
              <input v-model="newProduct.category" type="text" required class="form-input" />
            </div>
            <div class="form-actions">
              <button type="button" @click="cancelProductForm" class="btn">Отмена</button>
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
            <img
              :src="getImageUrl(product.image_url)"
              :alt="product.name"
              class="product-image"
              @error="handleImageError"
            />
          </div>
          <div class="product-content">
            <h3>{{ product.name }}</h3>
            <p>{{ product.description }}</p>
            <div class="product-category">{{ product.category }}</div>
            <button
              v-if="$store.state.isAdmin"
              @click="deleteProduct(product.id)"
              class="btn btn-danger btn-sm"
            >
              Удалить
            </button>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="products.length === 0" class="empty-state">
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

const newProduct = ref({
  name: "",
  description: "",
  image_url: "",
  category: "",
});

onMounted(() => {
  store.dispatch("fetchProducts");
});

const DEFAULT_IMAGE_URL = "http://placekitten.com/800/600";
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
.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 2rem;
}

.product-card {
  display: flex;
  flex-direction: column;
  height: 100%;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.product-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.15);
}

.product-image-container {
  width: 100%;
  height: 250px;
  overflow: hidden;
  background: #f8f9fa;
}

.product-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.product-card:hover .product-image {
  transform: scale(1.05);
}

.product-content {
  padding: 1.5rem;
  flex-grow: 1;
  display: flex;
  flex-direction: column;
}

.product-content h3 {
  margin-bottom: 1rem;
  color: #2c3e50;
  font-size: 1.25rem;
}

.product-content p {
  flex-grow: 1;
  margin-bottom: 1rem;
  color: #666;
  line-height: 1.5;
}

.product-category {
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 500;
  margin-bottom: 1rem;
  align-self: flex-start;
}

.btn-sm {
  padding: 0.5rem 1rem;
  font-size: 0.9rem;
}

.admin-controls {
  display: flex;
  justify-content: center;
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
  border-radius: 15px;
  width: 90%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 600;
  color: #2c3e50;
}

.form-input {
  width: 100%;
  padding: 0.75rem;
  border: 2px solid #e9ecef;
  border-radius: 8px;
  font-size: 1rem;
  transition: border-color 0.3s ease;
}

.form-input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

/* Стили для области загрузки файлов */
.file-upload-area {
  border: 2px dashed #dee2e6;
  border-radius: 8px;
  padding: 2rem;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  background: #f8f9fa;
}

.file-upload-area:hover {
  border-color: #667eea;
  background: #f0f2ff;
}

.file-upload-area.drag-over {
  border-color: #667eea;
  background: #e6ebff;
}

.upload-placeholder .upload-icon {
  font-size: 2rem;
  margin-bottom: 1rem;
}

.file-selected {
  color: #28a745;
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
  color: #6c757d;
}

.empty-state h3 {
  margin-bottom: 1rem;
  color: #495057;
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
}
</style>
