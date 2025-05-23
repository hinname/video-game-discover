FROM node:22.15.0-bookworm

WORKDIR /app

RUN npx playwright install --with-deps

COPY package*.json .

RUN npm ci

COPY . .

COPY entrypoint.sh /docker-entrypoint/entrypoint.sh

EXPOSE 3000

ENTRYPOINT [ "/docker-entrypoint/entrypoint.sh" ]