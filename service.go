package daemon

import (
	"fmt"
	"os"
	// "os/signal"
	// "syscall"
)

// Service Daemon
type ServiceDaemon struct {
	Daemon
}

func (daemon *ServiceDaemon) Manage(service Service) (string, error) {

	usage := "Usage: myservice install | remove | start | stop | status"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return daemon.Install()
		case "remove":
			return daemon.Remove()
		case "start":
			return daemon.Start()
		case "stop":
			return daemon.Stop()
		case "status":
			return daemon.Status()
		default:
			return usage, nil
		}
	}

	process := service.GetProcess()
	process()

	// never happen, but need to complete code
	return daemon.Start()
}

// Services

type Service struct {
	Name        string
	Description string
	Port        int
	Process     func()
}

func (service Service) GetName() string {
	return service.Name
}

func (service Service) GetPort() int {
	return service.Port
}

func (service Service) GetDescription() string {
	return service.Description
}

func (service Service) GetProcess() func() {
	return service.Process
}

func (service Service) Initialize() {

	srv, err := New(service.GetName(), service.GetDescription())

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	service_daemon := &ServiceDaemon{srv}
	status, err := service_daemon.Manage(service)

	if err != nil {
		fmt.Println(status, "\nError: ", err)
		os.Exit(1)
	}

	fmt.Println(status)
}

func (service *Service) GetInfo(serviceName string) {
	redis_ := redis.Redis{
		Name:   "Penlook",
		Server: "localhost:6379",
	}.Connect()

	result, _ := redis.String(redis_.Do("GET", "service.yml"))

	decoder := json.NewDecoder(strings.NewReader(result))
	var data map[string]map[string]interface{}
	decoder.Decode(&data)

	service.Description = data[serviceName]["description"].(string)
	service.Port = data[serviceName]["port"].(string)
	service.Name = serviceName
}
