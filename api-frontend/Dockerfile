FROM node:18-alpine

WORKDIR /app

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

RUN echo ${AUTH0_BASE_URL} 

RUN yarn build

CMD ["yarn", "start"]
