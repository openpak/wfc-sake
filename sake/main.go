package sake

import (
	"net/http"
	"openpak/wfc-common/common"
	"openpak/wfc-common/database"
	"openpak/wfc-common/logging"
)

var (
	db database.Connection
)

func StartServer(reload bool) {
	// Get config
	config := common.GetConfig()

	common.ReadGameList()

	// Start SQL
	db = database.Start(config)
	// Every accepted Mario Kart Wii ghost upload is also archived with its header metadata.
	if err := db.EnsureMarioKartWiiGhostArchive(); err != nil {
		logging.Error("SAKE", "ghost archive schema:", err)
	}
}

func Shutdown() {
	db.Close()
}

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("POST /SakeStorageServer/StorageServer.asmx", handleStorageRequest)
	mux.HandleFunc("GET /SakeFileServer/download.aspx", handleFileDownloadRequest)
	mux.HandleFunc("POST /SakeFileServer/upload.aspx", handleFileUploadRequest)
	mux.HandleFunc("GET /SakeFileServer/ghostdownload.aspx", handleMarioKartWiiGhostDownloadRequest)
}
