FROM  node:21.1.0-slim as builder
WORKDIR /root
COPY . .
RUN yarn global add pnpm
CMD ["cd","web"]
RUN yarn && yarn build


FROM nginx:stable-alpine3.17-slim
WORKDIR /opt/www/html/
COPY --from=builder /root/dist /usr/share/nginx/html
COPY --from=builder /root/deploy/my.conf  /etc/nginx/conf.d/my.conf

CMD ["nginx","-g","daemon off;"]
