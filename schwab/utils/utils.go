package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/samjtro/go-trade/utils"
)

func init() {
	err := utils.LoadConfig()
	utils.Check(err)
}

type AccessTokenResponse struct {
	expires_in    int
	token_type    string
	scope         string
	refresh_token string
	access_token  string
	id_token      string
}

type TOKEN struct {
	RefreshExpiration time.Time
	Refresh           string
	BearerExpiration  time.Time
	Bearer            string
}

func readDB() TOKEN {
	var tokens TOKEN
	body, err := os.ReadFile(fmt.Sprintf("%s/.foo/bar.json", utils.HomeDir()))
	utils.Check(err)

	err = json.Unmarshal(body, &tokens)
	utils.Check(err)

	return tokens
}

// Credit: https://go.dev/play/p/C2sZRYC15XN
func getStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)

	if s == -1 {
		return
	}

	s += len(start)
	e := strings.Index(str[s:], end)

	if e == -1 {
		return
	}

	return str[s : s+e]
}

// Credit: https://gist.github.com/hyg/9c4afcd91fe24316cbf0
func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	utils.Check(err)
}
