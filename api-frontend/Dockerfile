FROM node:18-alpine

RUN adduser --disabled-password --gecos "" runner \
    && usermod -aG sudo runner \
    && chown -R runner:runner /app

WORKDIR /app

USER runner

COPY package.json yarn.lock ./

RUN yarn install

COPY src ./src
COPY public ./public
COPY next.config.js .
COPY tsconfig.json .

ENV AUTH0_SECRET=${AUTH0_SECRET}
ENV AUTH0_BASE_URL=${AUTH0_BASE_URL}
ENV AUTH0_ISSUER_BASE_URL=${AUTH0_ISSUER_BASE_URL}
ENV AUTH0_CLIENT_ID=${AUTH0_CLIENT_ID}
ENV AUTH0_CLIENT_SECRET=${AUTH0_CLIENT_SECRET}

RUN yarn build

CMD ["yarn", "start"]
