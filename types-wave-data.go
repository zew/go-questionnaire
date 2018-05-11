package main

import "github.com/zew/go-questionaire/qst"

type tWave struct {
	WaveId string `json:"wave_id,omitempty"`
	qst.QuestionaireT
}
