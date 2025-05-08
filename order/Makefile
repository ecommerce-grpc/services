SERVICE_PORT=4001

run-service:
	DATA_SOURCE_URL="root:verysecretpassword@tcp(127.0.0.1:3306)/order" \
	APPLICATION_PORT="4001" \
	PAYMENT_SERVICE_URL="localhost:4002" \
	ENV="development" \
	go run cmd/main.go

request:
	grpcurl -d '{"user_id": 334,"order_items": [{"product_code": "sku22", "unit_price": 0.22, "quantity": 1}]}' \
	 -plaintext localhost:$(SERVICE_PORT) Order/Create