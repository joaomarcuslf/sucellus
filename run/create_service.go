package run

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/joaomarcuslf/sucellus/definitions"
	"github.com/joaomarcuslf/sucellus/models"
	dockerfile "github.com/joaomarcuslf/sucellus/templates/dockerfile"
	makefile "github.com/joaomarcuslf/sucellus/templates/makefile"
)

func CreateService(ctx context.Context, repository definitions.Repository, service models.Service) error {
	var aux models.Service

	dw := dockerfile.NewGoTemplate()
	mw := makefile.NewGoTemplate()

	cmd := exec.Command("ls", "-al")
	output, _ := cmd.CombinedOutput()

	content := string(output)

	path := ".services"

	fmt.Println("Checking if .services exists")

	if !strings.Contains(content, path) {
		cmd = exec.Command("mkdir", path)
		cmd.CombinedOutput()
	}

	fmt.Println("Creating service folder:", service.UName)

	cmd = exec.Command("mkdir", service.UName)
	cmd.Dir = path
	cmd.CombinedOutput()

	path = fmt.Sprintf("%s/%s", path, service.UName)

	f, err := os.Create(path + "/Dockerfile")

	if err != nil {
		return err
	}

	fmt.Printf("Writing %s/Dockerfile\n", service.UName)

	f.WriteString(dw.Execute(service.Port, service.UName, service.Url))
	f.Close()

	f, err = os.Create(path + "/Makefile")

	if err != nil {
		return err
	}

	fmt.Printf("Writing %s/Makefile\n", service.UName)
	f.WriteString(mw.Execute(service.Port, service.UName, service.Url))
	f.Close()

	if err != nil {
		return err
	}

	aux.Status = "CREATED"
	repository.Set(ctx, service.ID, aux)

	var envFile string

	for key := range service.EnvVars {
		envFile += fmt.Sprintf("%s=%s\n", key, service.EnvVars[key])
	}

	fmt.Printf("Writing %s/.env\n", service.UName)
	envFile += fmt.Sprintf("PORT=%d\n", service.Port)

	f, err = os.Create(path + "/.env")

	if err != nil {
		return err
	}

	_, err = f.WriteString(envFile)
	f.Close()

	if err != nil {
		return err
	}

	fmt.Printf("Running build-service for %s\n", service.UName)
	cmd = exec.Command("make", "build-service")
	cmd.Dir = path
	cmd.CombinedOutput()

	aux.Status = "BUILT"
	repository.Set(ctx, service.ID, aux)

	fmt.Printf("Running start-service for %s\n", service.UName)
	cmd = exec.Command("make", "start-service")
	cmd.Dir = path
	cmd.CombinedOutput()

	aux.Status = "RUNNING"
	repository.Set(ctx, service.ID, aux)

	content = string(output)

	fmt.Printf("Service %s is running\n", service.UName)

	fmt.Println("Finishing StartService")
	return nil
}
