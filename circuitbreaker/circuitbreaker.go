// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package circuitbreaker

import "time"

type Circuit struct {
	numErrors int
	commands  chan commandData
	resetTime time.Duration
}

type commandData struct {
	action int
	result chan<- interface{}
}

const (
	OK = iota
	ERROR
	GET_NUMBERS_OF_ERRORS
	RESET
	IS_RESETTING
	RESET_FINISHED
)

func NewCircuit(numErrors int, resetTime time.Duration) *Circuit {
	circuit := Circuit{
		numErrors: numErrors,
		resetTime: resetTime,
		commands:  make(chan commandData),
	}
	go circuit.run()
	return &circuit
}

func (crt Circuit) run() {
	numErrors := 0
	resetting := false
	lastError := time.Now()

	for command := range crt.commands {
		switch command.action {
		case GET_NUMBERS_OF_ERRORS:
			command.result <- numErrors
		case IS_RESETTING:
			command.result <- resetting
		case OK:
			numErrors = 0
		case ERROR:
			now := time.Now()
			if now.Sub(lastError) >= crt.resetTime {
				numErrors = 1
			} else {
				numErrors++
			}
			lastError = now

			if crt.numErrors == numErrors {
				resetting = true
				go crt.waitUntilReset(time.After(crt.resetTime))
			}
		case RESET_FINISHED:
			numErrors = 0
			resetting = false
		}
	}
}

func (crt Circuit) IsClosed() bool {
	return crt.replyTo(GET_NUMBERS_OF_ERRORS) != crt.numErrors
}

func (crt Circuit) IsOpen() bool {
	return !crt.IsClosed()
}

func (crt Circuit) replyTo(action int) interface{} {
	reply := make(chan interface{})
	crt.commands <- commandData{action: action, result: reply}
	return <-reply
}

func (crt Circuit) Ok() {
	crt.newStatus(OK)
}

func (crt Circuit) Error() {
	crt.newStatus(ERROR)
}

func (crt Circuit) newStatus(status int) {
	if !crt.isResetting() {
		crt.commands <- commandData{action: status}
	}
}

func (crt Circuit) isResetting() bool {
	return crt.replyTo(IS_RESETTING).(bool)
}

func (crt Circuit) waitUntilReset(resetTime <-chan time.Time) {
	select {
	case <-resetTime:
		crt.commands <- commandData{action: RESET_FINISHED}
	}
}
