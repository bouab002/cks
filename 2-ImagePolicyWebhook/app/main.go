package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// AdmissionReview represents the webhook request and response format
type AdmissionReview struct {
	APIVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Request    *AdmissionRequest      `json:"request,omitempty"`
	Response   *AdmissionResponse     `json:"response,omitempty"`
}

type AdmissionRequest struct {
	UID    string `json:"uid"`
	Object struct {
		Spec struct {
			Containers []struct {
				Image string `json:"image"`
			} `json:"containers"`
		} `json:"spec"`
	} `json:"object"`
}

type AdmissionResponse struct {
	UID     string  `json:"uid"`
	Allowed bool    `json:"allowed"`
	Status  *Status `json:"status,omitempty"`
}

type Status struct {
	Message string `json:"message"`
}

// validateHandler processes the AdmissionReview request
func validateHandler(w http.ResponseWriter, r *http.Request) {
	var admissionReview AdmissionReview
	allowedRegistries := []string{"registry.example.com", "docker.io"}

	// Decode the AdmissionReview request
	if err := json.NewDecoder(r.Body).Decode(&admissionReview); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse request: %v", err), http.StatusBadRequest)
		return
	}

	pod := admissionReview.Request.Object
	allowed := true
	var message string

	// Validate each container's image registry
	for _, container := range pod.Spec.Containers {
		image := container.Image
		registry := strings.Split(image, "/")[0]

		isAllowed := false
		for _, allowedRegistry := range allowedRegistries {
			if registry == allowedRegistry {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			allowed = false
			message = fmt.Sprintf("Image %s is not allowed. Allowed registries: %v", image, allowedRegistries)
			break
		}
	}

	// Build the response
	admissionReview.Response = &AdmissionResponse{
		UID:     admissionReview.Request.UID,
		Allowed: allowed,
	}

	if !allowed {
		admissionReview.Response.Status = &Status{
			Message: message,
		}
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(admissionReview)
}

func main() {
	http.HandleFunc("/validate", validateHandler)

	// Start HTTPS server with TLS certificates
	log.Println("Starting webhook server on :8080")
	if err := http.ListenAndServeTLS(":8080", "/app/certs/tls.crt", "/app/certs/tls.key", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

