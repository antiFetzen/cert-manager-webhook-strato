package python_strato_certbot

import (
	"cert-manager-webhook-strato/internal/python-strato-certbot/data"
	"cert-manager-webhook-strato/internal/types"
	"github.com/cert-manager/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"github.com/kluctl/go-embed-python/embed_util"
	"github.com/kluctl/go-embed-python/python"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func setEnv(name string, value string) {
	err := os.Setenv(name, value)
	if err != nil {
		panic(err)
	}
}

func setEnvironmentVariables(cfg types.StratoDNSProviderConfig, ch *v1alpha1.ChallengeRequest) {
	// provide environment variables for python `strato-certbot.auth-hook`
	//AUTHENTICATION
	setEnv("STRATO_AUTH_ENV_ENABLE", "TRUE")
	setEnv("STRATO_USERNAME", cfg.Username)
	setEnv("STRATO_PASSWORD", cfg.Password)
	setEnv("STRATO_TOTP_SECRET", cfg.TotpSecret)
	setEnv("STRATO_TOTP_DEVICENAME", cfg.TotpDeviceName)
	setEnv("STRATO_WAITING_TIME", strconv.Itoa(cfg.WaitingTime))

	//ACME-CHALLENGE
	setEnv("CERTBOT_VALIDATION", ch.Key)
	// remove trailing period
	setEnv("CERTBOT_DOMAIN", regexp.MustCompile(`\.$`).ReplaceAllString(ch.ResolvedZone, ""))
}

func executePython(fileName string) {
	tmpDir := filepath.Join(os.TempDir(), "cert-manager-webhook-strato")

	ep, err := python.NewEmbeddedPythonWithTmpDir(tmpDir+"-python", true)
	if err != nil {
		panic(err)
	}

	lib, err := embed_util.NewEmbeddedFilesWithTmpDir(data.Data, tmpDir+"-strato-certbot", true)
	if err != nil {
		panic(err)
	}

	ep.AddPythonPath(lib.GetExtractedPath())

	cmd, err := ep.PythonCmd("internal/python-strato-certbot/lib/strato-certbot/" + fileName)
	if err != nil {
		panic(err)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

func Present(cfg types.StratoDNSProviderConfig, ch *v1alpha1.ChallengeRequest) {
	setEnvironmentVariables(cfg, ch)

	executePython("auth-hook.py")
}

func CleanUp(cfg types.StratoDNSProviderConfig, ch *v1alpha1.ChallengeRequest) {
	setEnvironmentVariables(cfg, ch)

	executePython("cleanup-hook.py")
}
