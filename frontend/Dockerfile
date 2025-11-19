# Ебучий базовый образ
FROM node:24 AS builder

# Рабочая папка, куда все это говно пойдет
WORKDIR /app

# Копируем package.json, ебнем зависимости
COPY package*.json ./
RUN npm install
EXPOSE 5173
# Копируем весь код, ебашим сборку
COPY . .
CMD npm run dev

# # Переходим на nginx для раздачи статики
# FROM nginx:stable-alpine

# # Копируем собранное говно в nginx
# COPY --from=builder /app/dist /usr/share/nginx/html

# # Открываем порт 80
# EXPOSE 80

# # Запускаем эту блядскую шарманку
# CMD ["nginx", "-g", "daemon off;"]
