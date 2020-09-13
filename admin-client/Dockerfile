# build stage
FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# production stage
FROM nginx:stable-alpine as production-stage
COPY nginx.conf /etc/nginx/conf.d/configfile.template

ENV PORT 8080
ENV HOST 0.0.0.0
ENV VUE_APP_ClientID 4
ENV VUE_APP_ClientSecret c
ENV VUE_APP_Scopes all
ENV VUE_APP_RedirectURL http://127.0.0.1:8080
ENV VUE_APP_AUTH_URL http://127.0.0.1:9096/
ENV VUE_APP_AuthURL http://127.0.0.1:9096/auth/authorize
ENV VUE_APP_TokenURL http://127.0.0.1:9096/auth/token
ENV VUE_APP_ECOMMERCE_API_URL http://localhost:8081
ENV TEST_ENV qwertyuy

RUN sh -c "envsubst '\$PORT'  < /etc/nginx/conf.d/configfile.template > /etc/nginx/conf.d/default.conf"
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]