echo "worked"
echo "worked"
echo "worked"
echo "worked"
echo "worked"
echo "wor2k22seaddasd2"

package callback

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GhostEventOnSuccess struct {
	Database string `json:"database"`
	Table string `json:"table"`
	DDL string `json:"ddl"`
	CopiedRows string `json:"copied_rows"`
	ElapsedCopySeconds string `json:"elapsed_copy_seconds"`
}

type GhostEvent struct {
	Database string `json:"database"`
	Table string `json:"table"`
	DDL string `json:"ddl"`
	ElapsedSeconds string `json:"elapsed_copy_seconds"`
}

//TODO Some parameters declared temporarily. Could be read from config later.
const (
	owner = "migroscomtr"
	repo = "migros-db-schema"
	ghostOnSuccessMessage = "Schema migration has successfully completed on `%s.%s`\nCopied Rows `%s`. Elapsed Time `%s` seconds.\nChange Statement `%s`."
	ghostOnFailureMessage = "Schema migration has FAILED on `%s.%s`\nElapsed Time `%s` seconds.\nChange Statement `%s`. Logs:\n`%s`"
	ghostOnStartupMessage = "Schema migration is started to running on `%s.%s`"
	ghostOnBeginPostponedMessage = "GH-OST cutover postponed on `%s.%s`"
	ghostStatusMessage = "Status:\n`%s`"
	ghostUnpostponeMessage = "`%s`"
)

//After migration triggered successfully with !migrate, all pull request information will be written in a db. We will get the pr number from there we need to comment.
func (cb *CallbackServer) handleGhostOnSuccessHook(w http.ResponseWriter, r *http.Request) {
	var p GhostEventOnSuccess

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		cb.logger.Error(err)
		return
	}
	comment := fmt.Sprintf(ghostOnSuccessMessage, p.Database, p.Table, p.CopiedRows, p.ElapsedCopySeconds, p.DDL)

	fmt.Println(comment)

}

func (cb *CallbackServer) handleGhostOnFailureHook(w http.ResponseWriter, r *http.Request) {
	var p GhostEvent

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		cb.logger.Error(err)
		return
	}
	fmt.Println(p.Database)
}

func (cb *CallbackServer) handleGhostOnStartupHook(w http.ResponseWriter, r *http.Request) {
	var p GhostEvent

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		cb.logger.Error(err)
		return
	}
	fmt.Println(p.Database)
}

func (cb *CallbackServer) handleGhostOnBeginPostponedHook(w http.ResponseWriter, r *http.Request) {
	var p GhostEvent

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		cb.logger.Error(err)
		return
	}
	fmt.Println(p.Database)
}

//TODO These two functions are irrevelant with hooks. Could move to another package.

func (cb *CallbackServer) handleGhostStatus(w http.ResponseWriter, r *http.Request) {
	var p GhostEvent

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		cb.logger.Error(err)
		return
	}
	fmt.Println(p.Database)
}

func (cb *CallbackServer) handleGhostUnpostpone(w http.ResponseWriter, r *http.Request) {
	var p GhostEvent

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		cb.logger.Error(err)
		return
	}
	fmt.Println(p.Database)
}