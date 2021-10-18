package run

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/joaomarcuslf/sucellus/definitions"
	"github.com/joaomarcuslf/sucellus/models"
	"github.com/joaomarcuslf/sucellus/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

func StopServices(ctx context.Context, connection definitions.DatabaseClient) error {
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

	fmt.Println("Getting services list")

	services, err := repository.Query(ctx, bson.M{})

	if err != nil {
		return err
	}

	for _, s := range services {
		aux = s

		path = fmt.Sprintf(".services/%s", s.UName)

		fmt.Printf("Running stop-service for %s\n", s.UName)
		cmd := exec.Command("make", "stop-service")
		cmd.Dir = path
		output, _ := cmd.CombinedOutput()

		aux.Status = "STOPPED"
		repository.Set(ctx, s.ID, aux)

		content := string(output)

		fmt.Println(content)
	}

	fmt.Println("Finishing StopServices")
	return nil
}
