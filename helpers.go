package bdevault

import (
	"errors"
	"os/exec"
	"regexp"
)

// RecoveryKeyRegex matches BitLocker recovery keys
var RecoveryKeyRegex = regexp.MustCompile(`(?:\d{6}-){7}\d{6}`)

// GetRecoveryKey gets a recovery key for a specific Windows drive letter
func GetRecoveryKey(driveLetter string) (string, error) {
	cmd := exec.Command("manage-bde.exe", "-protectors", "-get", driveLetter, "-Type", "RecoveryPassword")

	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), err
	}
	return string(RecoveryKeyRegex.Find(out)), nil
}

// GetRecoveryKeys extracts all recovery keys for encrypted drives
func GetRecoveryKeys() (map[string]string, error) {
	volumeRegex := regexp.MustCompile(`Volume (\w:)`)

	cmd := exec.Command("manage-bde.exe", "-status")

	bdeStatus, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	driveKeys := make(map[string]string)
	for _, match := range volumeRegex.FindAllStringSubmatch(string(bdeStatus), 32) {
		letter := match[1]
		key, err := GetRecoveryKey(letter)
		if err == nil {
			driveKeys[letter] = key
		}
	}

	if len(driveKeys) == 0 {
		return driveKeys, errors.New("No keys found")
	}

	return driveKeys, nil
}
