services:
  api:
    image: "${API_IMAGE}"
    container_name: donely-api
    restart: always
    env_file:
      - "${API_ENV_FILE}"
    expose:
      - "8000"
    networks:
      - app

  web:
    image: "${WEB_IMAGE}"
    container_name: donely-web
    restart: always
    depends_on:
      - api
    environment:
      APP_DOMAIN: "${APP_DOMAIN}"
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - caddy_data:/data
      - caddy_config:/config
    networks:
      - app

networks:
  app:

volumes:
  caddy_data:
  caddy_config:
