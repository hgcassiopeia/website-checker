FROM node:alpine
WORKDIR /usr/app

COPY ./package.json ./yarn.lock ./
RUN yarn install
COPY ./ /usr/app

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

CMD ["yarn", "dev"]