protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/product.proto

protoc --go_out=plugins=grpc:./proto --go_opt=paths=source_relative ./product/product.proto



protoc -I ecommerce \
ecommerce/product_info.proto \
--go_out=plugins=grpc:./ecommerce
