package upload

import (
	"os"

	"github.com/jlaffaye/ftp"
	"github.com/morgulbrut/colorlog"
)

// FTP uploads a given file trough a FTP connection.
// ftpServer schould be formated as <server>:<port>
func FTP(ftpServer string, user string, pwd string, local string, remote string) {
	colorlog.Info("Uploading %s to %s on %s", local, remote, ftpServer)
	colorlog.Debug("Dialing")
	client, err := ftp.Dial(ftpServer)
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	colorlog.Debug("Login in using %s:%s", user, pwd)
	if err := client.Login(user, pwd); err != nil {
		colorlog.Fatal(err.Error())
	}

	file, err := os.Open(local)
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	colorlog.Debug("Transfering")
	client.Stor(remote, file)
}
