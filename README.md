docker build -t boutique/product-catalog-service:v0.0.1 .

docker run -d --network=host \
  -e "DB_HOST=127.0.0.1" \
  --name product-catalog-service boutique/product-catalog-service:v0.0.1


docker build -t boutique/product-catalog-service:v0.0.1 -f Dockerfile.production .

docker run -d -p 8080:8080 -e "PORT=8080" -e "DB_HOST=host.docker.internal" --user 1001 02cdf5d7f7ce