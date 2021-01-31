all: test

.PHONY: test
test: libinjection
	@echo "Run test cases"
	go test -v ./...

.PHONY: deps
deps:
	@echo "Check libinjection dependency"
	@if ! test -d vendor/libinjection; then \
		git clone https://github.com/client9/libinjection vendor/libinjection; \
	fi

.PHONY: libinjection
libinjection: deps clean
	@echo "Build libinjection dependency"
	@mkdir libinjection
	@cp vendor/libinjection/src/libinjection*.h ./libinjection
	@cp vendor/libinjection/src/libinjection*.c ./libinjection
	gcc -std=c99 -Wall -Werror -fpic -c libinjection/libinjection_sqli.c -o libinjection/libinjection_sqli.o 
	gcc -std=c99 -Wall -Werror -fpic -c libinjection/libinjection_xss.c -o libinjection/libinjection_xss.o
	gcc -std=c99 -Wall -Werror -fpic -c libinjection/libinjection_html5.c -o libinjection/libinjection_html5.o
	ar rcs libinjection/libinjection.a libinjection/*.o

.PHONY: clean
clean:
	@rm -rf libinjection

.PHONY: example
example: libinjection
	go build -a -ldflags '-extldflags "-static"' -o example/main example/main.go
