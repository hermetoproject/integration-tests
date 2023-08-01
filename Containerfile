FROM docker.io/node:20

COPY . /src
WORKDIR /src

RUN yarn install && yarn build

CMD ["yarn", "berryscary", "--name", "King Arthur"]
