
# execute from service dir
protoc -I ./ordermanagement/client/ecommerce \
./ordermanagement/client/ecommerce/order.proto \
--go_out=plugins=grpc:.


protoc -I ./ordermanagement/service/ecommerce \
./ordermanagement/service/ecommerce/order.proto \
--go_out=plugins=grpc:.
