# Build Astro
FROM node:18 AS build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Lightweight static server
FROM node:18-slim
WORKDIR /app
COPY --from=build /app/dist ./dist
EXPOSE 4321
CMD ["npx", "serve", "dist", "-l", "4321"]
