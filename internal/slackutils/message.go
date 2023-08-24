package slackutils

import (
	"fmt"

	"github.com/slack-go/slack"
)

type SendMessageParams struct {
	Endpoint string
	Script   string
	Failures int32
	Warnings int32
}

// Helper function to send a message.
func SendMessage(client *slack.Client, channel string, params SendMessageParams) (string, string, error) {
	blocks := []slack.Block{
		slack.NewHeaderBlock(slack.NewTextBlockObject(slack.PlainTextType, "OWASP ZAP Report", false, false)),
		slack.NewDividerBlock(),
	}

	if params.Failures > 0 {
		msg := fmt.Sprintf(":large_red_square: *Fails* were detected: %d", params.Failures)
		blocks = addContextBlock(blocks, "errors", msg)
	}

	if params.Warnings > 0 {
		msg := fmt.Sprintf(":large_yellow_square: *Warnings* were detected: %d", params.Warnings)
		blocks = addContextBlock(blocks, "warnings", msg)
	}

	blocks = addSectionBlock(blocks, "intro", slack.PlainTextType, "A routine security scan has been executed with the following parameters:")

	// Add the details related to this scan.
	blocks = addSectionBlock(blocks, "endpoint", slack.MarkdownType, fmt.Sprintf("*Endpoint:* `%s`", params.Endpoint))
	blocks = addSectionBlock(blocks, "script", slack.MarkdownType, fmt.Sprintf("*Script:* `%s`", params.Script))

	blocks = append(blocks, slack.NewDividerBlock())
	blocks = addSectionBlock(blocks, "find_more", slack.MarkdownType, "_The full report has been added to this thread._")

	return client.PostMessage(channel, slack.MsgOptionBlocks(blocks...))
}

// Helper function to add a context block to the message.
func addContextBlock(blocks []slack.Block, id, msg string) []slack.Block {
	block := slack.NewTextBlockObject(slack.MarkdownType, msg, false, false)
	blocks = append(blocks, slack.NewContextBlock(id, block))
	return blocks
}

// Helper function to add a section block to the message.
func addSectionBlock(blocks []slack.Block, id, msgType, msg string) []slack.Block {
	block := slack.NewTextBlockObject(msgType, msg, false, false)
	blocks = append(blocks, slack.NewSectionBlock(block, nil, nil, slack.SectionBlockOptionBlockID(id)))
	return blocks
}
