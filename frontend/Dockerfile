# Ебучий базовый образ
FROM node:24

# Рабочая папка, куда все это говно пойдет
WORKDIR /app

# Копируем package.json, ебнем зависимости
COPY package*.json ./
RUN npm install

# Копируем весь код, ебашим сборку
COPY . .
RUN npm run build

# Переходим на nginx для раздачи статики
FROM nginx:stable-alpine

# Копируем собранное говно в nginx
COPY --from=0 /app/dist /usr/share/nginx/html

# Открываем порт 80
EXPOSE 80

# Запускаем эту блядскую шарманку
CMD ["nginx", "-g", "daemon off;"]