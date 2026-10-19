package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	ver "github.com/vault-thirteen/auxie/Versioneer/classes/Versioneer"
	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"

	cm "github.com/vault-thirteen/SPA/pkg/common/message"
	server "github.com/vault-thirteen/SPA/pkg/proxy"
	"github.com/vault-thirteen/SPA/pkg/proxy/settings"
)

const AppName = "Proxy Server"

func main() {
	showIntro()

	cla, err := readCLA()
	mustBeNoError(err)
	if cla.IsDefaultFile() {
		cio.LogMessage(cm.Msg_UsingDefaultConfigurationFile)
	}

	var stn *settings.Settings
	stn, err = settings.NewSettingsFromFile(cla.ConfigurationFilePath)
	mustBeNoError(err)

	cio.LogMessage(cm.Msg_ServerStarting)
	var srv *server.Server
	srv, err = server.NewServer(stn)
	mustBeNoError(err)

	err = srv.Start()
	mustBeNoError(err)

	switch srv.GetWorkMode() {
	case settings.ServerModeIdHttp:
		cio.PrintMessage(cm.Msg_HttpServer, srv.GetListenDsn())
	case settings.ServerModeIdHttps:
		cio.PrintMessage(cm.Msg_HttpsServer, srv.GetListenDsn())
	}

	serverMustBeStopped := srv.GetStopChannel()
	waitForQuitSignalFromOS(serverMustBeStopped)
	<-*serverMustBeStopped

	cio.LogMessage(cm.Msg_ServerStopping)
	err = srv.Stop()
	if err != nil {
		errors.LogError(err)
	}
	cio.LogMessage(cm.Msg_ServerStopped)
	time.Sleep(time.Second)
}

func mustBeNoError(err error) {
	if err != nil {
		errors.LogError(err)
		os.Exit(1)
	}
}

func showIntro() {
	versioneer, err := ver.New(false)
	mustBeNoError(err)
	versioneer.ShowIntroText(AppName)
	versioneer.ShowComponentsInfoText()
	cio.PrintMessage()
}

func waitForQuitSignalFromOS(serverMustBeStopped *chan bool) {
	osSignals := make(chan os.Signal, 16)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for sig := range osSignals {
			switch sig {
			case syscall.SIGINT,
				syscall.SIGTERM:
				cio.LogMessage(cm.Msg_QuitSignal, sig)
				*serverMustBeStopped <- true
			}
		}
	}()
}
