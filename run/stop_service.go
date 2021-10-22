package run

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/joaomarcuslf/sucellus/definitions"
	"github.com/joaomarcuslf/sucellus/models"
	"github.com/joaomarcuslf/sucellus/repositories"
)

func StopService(ctx context.Context, connection definitions.DatabaseClient, id string) error {
	repository := repositories.NewServiceRepository(connection)

	var aux models.Service

	path := ".services"

	cmd := exec.Command("ls", "-al")
	output, _ := cmd.CombinedOutput()

	content := string(output)

	fmt.Println("Checking if .services exists")

	if !strings.Contains(content, path) {
		cmd := exec.Command("mkdir", path)
		cmd.CombinedOutput()
		return nil
	}

	fmt.Println("Getting service:", id)

	s, err := repository.Get(ctx, id)

	if err != nil {
		return err
	}

	aux = s.(models.Service)

	path = fmt.Sprintf(".services/%s", aux.UName)

	fmt.Printf("Running stop-service for %s\n", aux.UName)
	cmd = exec.Command("make", "stop-service")
	cmd.Dir = path
	output, _ = cmd.CombinedOutput()

	aux.Status = "STOPPED"
	repository.Set(ctx, aux.ID, aux)

	content = string(output)

	fmt.Println(content)

	fmt.Println("Finishing StopService")
	return nil
}
