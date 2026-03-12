JAEGGER_RUNNING = $(docker ps --filter name=jaegger)

run-database:
	CONTAINERS=$(docker ps --filter ancestor=mysql --filter=payment)
	if [[ $(CONTAINERS) == "" ]]; then
		docker run -d -p 3307:3306 -e MYSQL_ROOT_PASSWORD=verysecretpassword -e MYSQL_DATABASE=payment mysql
	fi

# run-jaegger:
	# echo $JAEGGER_RUNNING
	# ifeq  $(JAEGGER_RUNNING) ""
	# 	docker run -d --rm --name jaeger \
	# 	-p 16686:16686 \
	# 	-p 4317:4317 \
	# 	-p 4318:4318 \
	# 	-p 5778:5778 \
	# 	-p 9411:9411 \
	# 	jaegertracing/jaeger:2.5.0
    # endif

run-service:
	DATA_SOURCE_URL="root:verysecretpassword@tcp(127.0.0.1:3307)/payment" \
	APPLICATION_PORT="4002" \
	ENV="development" \
	TRACER_BACKEND="http://localhost:16686/api/traces" \
	go run cmd/main.go