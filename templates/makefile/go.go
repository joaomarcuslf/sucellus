package makefile

import "fmt"

type GoTemplate struct{}

func NewGoTemplate() *GoTemplate {
	return &GoTemplate{}
}

func (t *GoTemplate) Execute(port int, appname, url string) string {
	return fmt.Sprintf(`GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

build-service:
	docker build . -t sucellus/%s

start-service:
	docker run --name %s \
	--rm -it -p %d:%d \
	-d sucellus/%s


stop-service:
	docker stop %s`, appname, appname, port, port, appname, appname)
}
