GO=go
GOBUILD= go build
APPNAME=zombiewalk
APP_BUILD_FOLDER=build/zombiewalk

.PHONEY: build 
bld: 
		cd cmd/ && $(GOBUILD) -o ../$(APP_BUILD_FOLDER)

build-linux: clean ## Prepare a build for a linux environment
	CGO_ENABLED=0 $(GOBUILD) -a -installsuffix cgo -o $(APP_BUILD_FOLDER)
	redis-server &
	./$(APPNAME)

clean: 
	rm $(APP_BUILD_FOLDER)

run: bld
	$(APP_BUILD_FOLDER)

test: bld
	go test ./...