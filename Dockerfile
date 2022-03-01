# build stage
FROM golang:alpine AS builder

ARG APP_NAME

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY pb ./pb
COPY api ./api
RUN go build -o app ./cmd/${APP_NAME}

# build frontend
FROM node:alpine AS builder-f2e

WORKDIR /src

ENV NODE_OPTIONS=--openssl-legacy-provider
COPY web/package.json web/yarn.lock ./
RUN yarn install

ARG DEPLOY_TO=uat

COPY ./web/src ./src
COPY ./web/index.html ./index.html
COPY ./web/vite.config.ts ./vite.config.ts
COPY ./web/tailwind.config.js ./tailwind.config.js
COPY ./web/postcss.config.js ./postcss.config.js
COPY ./web/tsconfig.json ./tsconfig.json
COPY ./web/tsconfig.node.json ./tsconfig.node.json
COPY ./web/.env.${DEPLOY_TO} ./.env
RUN yarn build

# final stage
FROM alpine:3

LABEL maintainer.name="blackhorseya"
LABEL maintainer.email="blackhorseya@gmail.com"

WORKDIR /app

COPY --from=builder /src/app ./
COPY --from=builder-f2e /src/dist ./web/build

ENTRYPOINT ./app