FROM node:alpine
WORKDIR /usr/app

COPY ./package.json ./yarn.lock ./
RUN yarn install
COPY ./ /usr/app

CMD ["yarn", "dev"]