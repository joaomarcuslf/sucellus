package run

import (
	"fmt"
	"os/exec"

	"github.com/joaomarcuslf/sucellus/definitions"
	"github.com/joaomarcuslf/sucellus/models"
)

func DeleteService(repository definitions.Repository, service models.Service) error {
	path := ".services"

	path += "/" + service.UName

	fmt.Printf("Running stop-service for %s\n", service.UName)
	cmd := exec.Command("make", "stop-service")
	cmd.Dir = path
	cmd.CombinedOutput()

	fmt.Printf("Running delete-service for %s\n", service.UName)
	cmd = exec.Command("make", "delete-service")
	cmd.Dir = path
	cmd.CombinedOutput()

	fmt.Println("Clearing folder content")
	cmd = exec.Command("rm", "-rf", service.UName)
	cmd.Dir = ".services"
	cmd.CombinedOutput()

	fmt.Println("Finishing DeleteService")
	return nil
}
