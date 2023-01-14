package main

import (
	"fmt"
	"net/http"
	"strings"
	"encoding/json"
)

// Profile struct to parse the user's profile from keycloak
type Profile struct {
    Name string `json:"name"`
    Sub string `json:"sub"`
    EmailVerified bool `json:"email_verified"`
    PreferredUsername string `json:"preferred_username"`
    GivenName string `json:"given_name"`
}

func authenticateUser(keycloakURL, realm, clientID, clientSecret string) (string, error) {
	// Prepare the request data
	data := fmt.Sprintf("grant_type=password&client_id=%s&client_secret=%s&username=golang-snippets-user&password=admin",
		clientID, clientSecret)

	// Create an http client and post the request
	client := &http.Client{}
	req, err := http.NewRequest("POST", keycloakURL+"/auth/realms/"+realm+"/protocol/openid-connect/token",
		strings.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Parse the response JSON
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	// Check for an error
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Error authenticating user: %s", result["error_description"])
	}

	// Get the access token from the response
	accessToken, _ := result["access_token"].(string)
	return accessToken, nil
}

func getUserProfile(keycloakURL, realm, token string) (*Profile, error) {
	// Create an http client and get the request
	client := &http.Client{}
	req, err := http.NewRequest("GET", keycloakURL+"/auth/realms/"+realm+"/protocol/openid-connect/userinfo", nil)
	if err != nil {
		return nil, err
	}

	// Add the Authorization header with the access token
	req.Header.Add("Authorization", "Bearer "+token)

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the response JSON
	var profile Profile
	json.NewDecoder(resp.Body).Decode(&profile)

	// Check for an error
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting user profile: %d", resp.StatusCode)
	}
	
	return &profile, nil
}

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## KeyCloak")

	// Keycloak server URL and Realm
	//http://localhost:8080/auth/realms/master/account/
	keycloakURL := "http://localhost:8080"
	realm := "master"

	// Client ID and secret - secret should be handled differently in real code
	clientID := "golang-snippets-client"
	clientSecret := "VT0jOTILMpZ7f1yWAcNJOhZwzkRRSQ7r"

	// Authenticate the user and get their access token
	token, err := authenticateUser(keycloakURL, realm, clientID, clientSecret)
	if err != nil {
		fmt.Println("Error authenticating user:", err)
		return
	}
	// Debug: fmt.Println("Token:", token)

	// Use the access token to get the user's profile
	profile, err := getUserProfile(keycloakURL, realm, token)
	if err != nil {
		fmt.Println("Error getting user profile:", err)
		return
	}

	// Print the user's name and show if email is verified
	fmt.Printf("Name: %s\n", profile.Name)
	fmt.Println("Is the email verified?", profile.EmailVerified)
}