FROM node:18-alpine as development
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY . .
RUN yarn codegen
CMD ["yarn", "dev"]

FROM node:18-alpine as build
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY . .
RUN yarn codegen
RUN yarn build


FROM node:18-alpine as production
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile --production
COPY --from=build /app/.next ./.next
COPY --from=build /app/public ./public
COPY --from=build /app/next.config.js ./next.config.js
COPY --from=build /app/LICENSE ./LICENSE
COPY --from=build /app/README.md ./README.md
CMD [ "yarn" , "start" ]
