package main

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestQuery(t *testing.T) {

	expression := "fields @timestamp, @message\n    | filter @message not like 'example'\n    | sort @timestamp asc\n    | limit 100"

	q := QueryDetails{}
	q.Add("end", "0", false)
	q.Add("start", "-3600", false)
	q.Add("timeType", "RELATIVE", true)
	q.Add("unit", "seconds", true)
	q.Add("editorString", expression, true)
	q.Add("isLiveTrail", "false", false)
	q.Add("source", "/application/sample1", true)
	q.Add("source", "/application/sample2", true)

	t.Run("PrimitiveEscape", func(t *testing.T) {
		assert.DeepEqual(t, q, QueryDetails{
			"end":          []string{"0"},
			"start":        []string{"-3600"},
			"timeType":     []string{"'RELATIVE"},
			"unit":         []string{"'seconds"},
			"editorString": []string{"'fields*20*40timestamp*2C*20*40message*0A*20*20*20*20*7C*20filter*20*40message*20not*20like*20'example'*0A*20*20*20*20*7C*20sort*20*40timestamp*20asc*0A*20*20*20*20*7C*20limit*20100"},
			"isLiveTrail":  []string{"false"},
			"source":       []string{"'*2Fapplication*2Fsample1", "'*2Fapplication*2Fsample2"},
		})
	})

	assert.DeepEqual(t, q.Encode(), "$3FqueryDetail$3D$257E$2528editorString$257E$2527fields*20*40timestamp*2C*20*40message*0A*20*20*20*20*7C*20filter*20*40message*20not*20like*20$2527example$2527*0A*20*20*20*20*7C*20sort*20*40timestamp*20asc*0A*20*20*20*20*7C*20limit*20100$257Eend$257E0$257EisLiveTrail$257Efalse$257Esource$257E$2528$257E$2527*2Fapplication*2Fsample1$257E$2527*2Fapplication*2Fsample2$2529$257Estart$257E-3600$257EtimeType$257E$2527RELATIVE$257Eunit$257E$2527seconds$2529")
}

func TestEscape(t *testing.T) {
	input := "~(end~0~start~-3600~timeType~'RELATIVE~unit~'seconds~editorString~'fields*20*40timestamp*2C*20*40message*0A*20*20*20*20*7C*20filter*20*40message*20not*20like*20'example'*0A*20*20*20*20*7C*20sort*20*40timestamp*20asc*0A*20*20*20*20*7C*20limit*20100~isLiveTrail~false~source~(~'*2Fapplication*2Fsample1~'*2Fapplication*2Fsample2))"
	want := "%7E%28end%7E0%7Estart%7E-3600%7EtimeType%7E%27RELATIVE%7Eunit%7E%27seconds%7EeditorString%7E%27fields*20*40timestamp*2C*20*40message*0A*20*20*20*20*7C*20filter*20*40message*20not*20like*20%27example%27*0A*20*20*20*20*7C*20sort*20*40timestamp*20asc*0A*20*20*20*20*7C*20limit*20100%7EisLiveTrail%7Efalse%7Esource%7E%28%7E%27*2Fapplication*2Fsample1%7E%27*2Fapplication*2Fsample2%29%29"
	assert.DeepEqual(t, want, Escape(input))
}

func TestFinal(t *testing.T) {
	input := "?queryDetail=%7E%28end%7E0%7Estart%7E-3600%7EtimeType%7E%27RELATIVE%7Eunit%7E%27seconds%7EeditorString%7E%27fields*20*40timestamp*2C*20*40message*0A*20*20*20*20*7C*20filter*20*40message*20not*20like*20%27example%27*0A*20*20*20*20*7C*20sort*20*40timestamp*20asc*0A*20*20*20*20*7C*20limit*20100%7EisLiveTrail%7Efalse%7Esource%7E%28%7E%27*2Fapplication*2Fsample1%7E%27*2Fapplication*2Fsample2%29%29"
	want := "$3FqueryDetail$3D$257E$2528end$257E0$257Estart$257E-3600$257EtimeType$257E$2527RELATIVE$257Eunit$257E$2527seconds$257EeditorString$257E$2527fields*20*40timestamp*2C*20*40message*0A*20*20*20*20*7C*20filter*20*40message*20not*20like*20$2527example$2527*0A*20*20*20*20*7C*20sort*20*40timestamp*20asc*0A*20*20*20*20*7C*20limit*20100$257EisLiveTrail$257Efalse$257Esource$257E$2528$257E$2527*2Fapplication*2Fsample1$257E$2527*2Fapplication*2Fsample2$2529$2529"
	input = QueryEscape(input)
	input = strings.ReplaceAll(input, "%", "$")
	assert.DeepEqual(t, input, want)
}
