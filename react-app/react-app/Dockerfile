# Build stage 
FROM node:14 AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install 
COPY . . 
RUN npm run build 

# Run stage 
FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80 
