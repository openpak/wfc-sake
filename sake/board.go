package sake

import (
	"context"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"openpak/wfc-common/common"
	"openpak/wfc-common/logging"
	"openpak/wfc-common/mkwboard"
)

// board is the Mario Kart Wii view of the leaderboards core; nil when WFC_LEADERBOARDS_ADDR is
// not set (the upload then only fills the game's ghost table).
var board atomic.Pointer[mkwboard.Board]

const reconcileEvery = 10 * time.Minute

// StartBoard connects to the leaderboards core and the account core, declares the title and
// imports the times the game's table holds that the core does not (the one-way migration, also
// the catch-up after a submit lost while the core was down). It returns at once; the work runs
// until ctx ends.
func StartBoard(ctx context.Context, cfg common.Config) {
	addr, key := os.Getenv("WFC_LEADERBOARDS_ADDR"), os.Getenv("WFC_LEADERBOARDS_KEY")
	if addr == "" {
		logging.Notice("SAKE", "WFC_LEADERBOARDS_ADDR is not set: Mario Kart Wii times are not sent to the leaderboards core")
		return
	}
	if cfg.CoreAddress == "" {
		logging.Error("SAKE", "WFC_CORE_ADDRESS is not set: cannot map consoles to accounts for the leaderboards core")
		return
	}
	lb, _, err := mkwboard.Dial(addr, key)
	if err != nil {
		logging.Error("SAKE", "leaderboards dial:", err)
		return
	}
	links, _, err := mkwboard.DialCore(cfg.CoreAddress, cfg.CoreKey)
	if err != nil {
		logging.Error("SAKE", "account core dial:", err)
		return
	}
	b := &mkwboard.Board{LB: lb, Links: links, CourseName: courseName}
	board.Store(b)
	go func() {
		declared := false
		for {
			if !declared {
				if err := b.Declare(ctx); err != nil {
					logging.Error("SAKE", "leaderboards declare:", err)
				} else {
					declared = true
				}
			}
			if declared {
				st, err := b.Reconcile(ctx, mkwboard.DBSource{DB: &db})
				if err != nil {
					logging.Error("SAKE", "leaderboards reconcile:", err)
				} else {
					logging.Info("SAKE", "leaderboards reconcile: seen", st.Seen, "imported", st.Imported, "present", st.Present, "unlinked", st.Unlinked, "failed", st.Failed)
				}
			}
			select {
			case <-ctx.Done():
				return
			case <-time.After(reconcileEvery):
			}
		}
	}()
}

func courseName(course int) string { return common.MarioKartWiiCourseId(course).Name() }

// submitTime files an accepted upload on the leaderboards core. A failure never fails the
// upload: the game's table has the time and the next reconcile run sends it.
func submitTime(ctx context.Context, region common.MarioKartWiiLeaderboardRegionId, course common.MarioKartWiiCourseId, score, pid int, playerInfo string) {
	b := board.Load()
	if b == nil {
		return
	}
	uid, ok, err := db.MKWUserOfProfile(ctx, int64(pid))
	if err != nil || !ok {
		return
	}
	err = b.Submit(ctx, mkwboard.Time{UserID: uid, PID: int64(pid), Region: int(region), Course: int(course), Score: score, PlayerInfo: playerInfo, At: time.Now()})
	switch err {
	case nil:
	case mkwboard.ErrNoLink:
		logging.Info("SAKE", "leaderboards: profile", strconv.Itoa(pid), "has no linked account; time kept in the ghost table")
	default:
		logging.Error("SAKE", "leaderboards submit:", err)
	}
}

