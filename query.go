package dstask

// main task data structures

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// when referring to tasks by ID, NON_RESOLVED_STATUSES must be loaded exclusively --
// even if the filter is set to show issues that have only some statuses.
type Query struct {
	Cmd           string
	IDs           []int
	Tags          []string
	AntiTags      []string
	Project       string
	AntiProjects  []string
	Priority      string
	Template      int
	Text          string
	IgnoreContext bool
	// any words after the note operator: /
	Note string
}

// reconstruct args string
func (query Query) String() string {
	var args []string

	addArg := func(prefix, value string) {
		if value != "" {
			args = append(args, prefix+value)
		}
	}

	for _, id := range query.IDs {
		args = append(args, strconv.Itoa(id))
	}

	for _, tag := range query.Tags {
		addArg("+", tag)
	}
	for _, tag := range query.AntiTags {
		addArg("-", tag)
	}

	addArg("project:", query.Project)

	for _, project := range query.AntiProjects {
		addArg("-project:", project)
	}

	addArg("", query.Priority)

	if query.Template > 0 {
		args = append(args, fmt.Sprintf("template:%v", query.Template))
	}

	addArg("\"", query.Text+"\"")

	return strings.Join(args, " ")
}

func (query Query) PrintContextDescription() {
	var envVarNotification string
	if os.Getenv("DSTASK_CONTEXT") != "" {
		envVarNotification = " (set by DSTASK_CONTEXT)"
	}
	if query.String() != "" {
		fmt.Printf("\033[33mActive context%s: %s\033[0m\n", envVarNotification, query)
	}
}

// HasOperators returns true if the query has positive or negative projects/tags,
// priorities, template
func (query Query) HasOperators() bool {
	return (len(query.Tags) > 0 ||
		len(query.AntiTags) > 0 ||
		query.Project != "" ||
		len(query.AntiProjects) > 0 ||
		query.Priority != "" ||
		query.Template > 0)
}

// ParseQuery parses the raw command line typed by the user.
func ParseQuery(args ...string) Query {
	var (
		cmd                string
		ids                []int
		tags               []string
		antiTags           []string
		project            string
		antiProjects       []string
		priority           string
		template           int
		words              []string
		notes              []string
		ignoreContext      bool
		notesModeActivated bool
		IDsExhausted       bool
	)

	for _, item := range args {
		lcItem := strings.ToLower(item)

		switch {
		case notesModeActivated:
			notes = append(notes, item)

		case cmd == "" && StrSliceContains(ALL_CMDS, lcItem):
			cmd = lcItem

		case !IDsExhausted:
			if id, err := strconv.ParseInt(item, 10, 64); err == nil {
				ids = append(ids, int(id))
				continue
			}
			IDsExhausted = true

		case item == IGNORE_CONTEXT_KEYWORD:
			ignoreContext = true

		case item == NOTE_MODE_KEYWORD:
			notesModeActivated = true

		case project == "" && strings.HasPrefix(lcItem, "project:"):
			project = lcItem[8:]

		case project == "" && strings.HasPrefix(lcItem, "+project:"):
			project = lcItem[9:]

		case strings.HasPrefix(lcItem, "-project:"):
			antiProjects = append(antiProjects, lcItem[9:])

		case strings.HasPrefix(lcItem, "template:"):
			if t, err := strconv.ParseInt(lcItem[9:], 10, 64); err == nil {
				template = int(t)
			}

		case len(item) > 1 && lcItem[0] == '+':
			tags = append(tags, lcItem[1:])

		case len(item) > 1 && lcItem[0] == '-':
			antiTags = append(antiTags, lcItem[1:])

		case priority == "" && IsValidPriority(item):
			priority = item

		default:
			words = append(words, item)
		}
	}

	return Query{
		Cmd:           cmd,
		IDs:           ids,
		Tags:          tags,
		AntiTags:      antiTags,
		Project:       project,
		AntiProjects:  antiProjects,
		Priority:      priority,
		Template:      template,
		Text:          strings.Join(words, " "),
		Note:          strings.Join(notes, " "),
		IgnoreContext: ignoreContext,
	}
}

// Merge applies a context to a new task. Returns new Query, does not mutate.
func (query *Query) Merge(q2 Query) Query {
	// dereference to make a copy of this query
	q := *query

	for _, tag := range q2.Tags {
		if !StrSliceContains(q.Tags, tag) {
			q.Tags = append(q.Tags, tag)
		}
	}

	for _, tag := range q2.AntiTags {
		if !StrSliceContains(q.AntiTags, tag) {
			q.AntiTags = append(q.AntiTags, tag)
		}
	}

	if q2.Project != "" {
		if q.Project != "" && q.Project != q2.Project {
			ExitFail("Could not apply q2, project conflict")
		} else {
			q.Project = q2.Project
		}
	}

	if q2.Priority != "" {
		if q.Priority != "" {
			ExitFail("Could not apply q2, priority conflict")
		} else {
			q.Priority = q2.Priority
		}
	}

	return q
}
