package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/SapphicCode/bdevault"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(zerolog.NewConsoleWriter())

	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}

	vaultToken := flag.String("token", os.Getenv("VAULT_TOKEN"), "Vault token")
	vaultAddress := flag.String("address", os.Getenv("VAULT_ADDR"), "Vault address")
	vaultMount := flag.String("mount", "secrets", "Vault mount to write to")
	vaultPrefix := flag.String("prefix", "bitlocker/", "Vault mount keyspace prefix")
	vaultKey := flag.String("key", strings.ToLower(hostname), "Vault key to write into")

	flag.Parse()

	// fetch keys
	keyMap, err := bdevault.GetRecoveryKeys()
	if err != nil {
		logger.Panic().Err(err).Msg("Failed to get recovery keys.")
	}
	logger.Info().Int("keys", len(keyMap)).Msgf("Found %d recovery keys to submit.", len(keyMap))
	dataMap := map[string]map[string]string{
		"data": keyMap,
	}

	// figure out address
	addr := fmt.Sprintf("%s/v1/%s/data/%s%s", *vaultAddress, *vaultMount, *vaultPrefix, *vaultKey)
	// set up body
	data := new(bytes.Buffer)
	encoder := json.NewEncoder(data)
	if err := encoder.Encode(dataMap); err != nil {
		logger.Panic().Err(err).Msg("Couldn't encode JSON.")
	}
	// assemble request
	request, err := http.NewRequest("POST", addr, data)
	if err != nil {
		logger.Panic().Err(err).Msg("Couldn't create HTTP request.")
	}
	request.Header["X-Vault-Token"] = []string{*vaultToken}
	// fire
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		logger.Panic().Err(err).Msg("Couldn't send HTTP request.")
	}
	logger.Info().Msgf("Vault Status: %s\n", resp.Status)
}
