package cron

const ACTION_TAG = `group:"executors"`

type Action interface {
	Execute()
}
