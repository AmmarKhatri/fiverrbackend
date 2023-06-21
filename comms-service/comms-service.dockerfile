FROM node:18-alpine

WORKDIR /app

COPY package.json .
COPY package-lock.json .

RUN npm ci --only=production

COPY . .

CMD [ "node", "index.js" ]
