package proxy

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"slices"
)

func GetServiceFromRoute(route string) (*Service, error) {
	index := slices.IndexFunc(Configuration.Services, func(service Service) bool {
		return service.Path == route
	})

	if index == -1 {
		return nil, errors.New("service not found")
	}

	return &Configuration.Services[index], nil
}

func CreateRequests(service *Service, path string, r *http.Request) (*[]*http.Request, error) {
	active, err := http.NewRequest(r.Method, service.Target.Active+"/"+path, r.Body)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("cannot create the active request")
	}

	preview, err := http.NewRequest(r.Method, service.Target.Preview+"/"+path, r.Body)

	// TODO maybe we can only send active and still proxy
	if err != nil {
		return nil, errors.New("cannot create the preview request")
	}

	// Set the headers
	copyHeaders(active, r)
	copyHeaders(preview, r)

	// add the Shadow header to the preview
	preview.Header.Add("X-Shadow-Request", "true")
	// TODO add X-FORWARDED-FOR to both

	// add tracking
	trackingNumber := "XSR-" + uuid.New().String()
	active.Header.Add("X-Shadow-Tracking", trackingNumber)
	preview.Header.Add("X-Shadow-Tracking", trackingNumber)

	return &[]*http.Request{active, preview}, nil
}
