package main

import (
	"github.com/alecthomas/kingpin/v2"
	"github.com/slack-go/slack"

	"github.com/previousnext/zap-orb/internal/slackutils"
	"github.com/previousnext/zap-orb/internal/zaputils"
)

var (
	cliEndpoint     = kingpin.Flag("endpoint", "The endpoint which was targetted for the report.").Required().String()
	cliScriptName   = kingpin.Flag("script-name", "Name of the script which was executed.").Required().String()
	cliReportPath   = kingpin.Flag("report-path", "Path to the full report which will be attached to the Slack message.").Required().String()
	cliProgressPath = kingpin.Flag("progress-path", "Path to the progress file to determine how many fails, warnings and passes there are.").Required().String()
	cliSlackToken   = kingpin.Flag("slack-token", "Token used to authenticate with Slack").Envar("OWASP_ZAP_SLACK_TOKEN").Required().String()
	cliSlackChannel = kingpin.Arg("slack-channel", "Channel that will be notified.").Required().String()
)

func main() {
	kingpin.Parse()

	progress, err := zaputils.GetProgress(*cliProgressPath)

	client := slack.New(*cliSlackToken)

	params := slackutils.SendMessageParams{
		Endpoint: *cliEndpoint,
		Script:   *cliScriptName,
		Failures: progress.Fail,
		Warnings: progress.Warn,
	}

	channel, timestamp, err := slackutils.SendMessage(client, *cliSlackChannel, params)
	if err != nil {
		panic(err)
	}

	_, err = client.UploadFile(slack.FileUploadParameters{
		File: *cliReportPath,
		Channels: []string{
			channel,
		},
		ThreadTimestamp: timestamp,
	})
	if err != nil {
		panic(err)
	}
}
