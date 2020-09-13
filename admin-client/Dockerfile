# build stage
FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
ENV VUE_APP_ClientID $VUE_APP_ClientID
ENV VUE_APP_ClientSecret $VUE_APP_ClientSecret
ENV VUE_APP_Scopes $VUE_APP_Scopes
ENV VUE_APP_RedirectURL $VUE_APP_RedirectURL
ENV VUE_APP_AUTH_URL $VUE_APP_AUTH_URL
ENV VUE_APP_AuthURL $VUE_APP_AuthURL
ENV VUE_APP_TokenURL $VUE_APP_TokenURL
ENV VUE_APP_ECOMMERCE_API_URL $VUE_APP_ECOMMERCE_API_URL

RUN npm run build

# production stage
FROM nginx:stable-alpine as production-stage
COPY nginx.conf /etc/nginx/conf.d/configfile.template

ENV PORT 8080
ENV HOST 0.0.0.0

RUN sh -c "envsubst '\$PORT'  < /etc/nginx/conf.d/configfile.template > /etc/nginx/conf.d/default.conf"
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]