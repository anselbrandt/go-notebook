FROM node:16 AS dependencies
WORKDIR /app
COPY package*.json ./
RUN npm install

FROM dependencies AS builder
COPY . .
RUN npm run build

FROM nginx
COPY --from=builder /app/build /usr/share/nginx/html
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]