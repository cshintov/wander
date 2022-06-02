package nomad

import (
	"fmt"
	"wander/components/page"
	"wander/formatter"
	"wander/message"

	tea "github.com/charmbracelet/bubbletea"
)

func FetchJobSpec(url, token, jobID, jobNamespace string) tea.Cmd {
	return func() tea.Msg {
		params := map[string]string{
			"namespace": jobNamespace,
		}
		fullPath := fmt.Sprintf("%s%s%s", url, "/v1/job/", jobID)
		body, err := get(fullPath, token, params)
		if err != nil {
			return message.ErrMsg{Err: err}
		}

		pretty := formatter.PrettyJsonStringAsLines(string(body))

		var jobSpecPageData []page.Row
		for _, row := range pretty {
			jobSpecPageData = append(jobSpecPageData, page.Row{Key: "", Row: row})
		}

		return PageLoadedMsg{
			Page:        JobSpecPage,
			TableHeader: []string{},
			AllPageData: jobSpecPageData,
		}
	}
}