protoc -I productinfo/service/ecommerce \
productinfo/service/ecommerce/product_info.proto \
--go_out=plugins=grpc:.
/service/ecommerce



protoc -I ecommerce \
ecommerce/product_info.proto \
--go_out=plugins=grpc:../../productinfo/service/ecommerce


# execute from service dir
protoc -I ecommerce \
ecommerce/product_info.proto \
--go_out=plugins=grpc:ecommerce