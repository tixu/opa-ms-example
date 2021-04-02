FROM node as builder
#RUN mkdir -p
COPY cpq cpq
WORKDIR /cpq
RUN yarn install
RUN yarn build 



FROM nginx:1.16.0-alpine
COPY --from=builder cpq/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]