package cron

import (
	"errors"
	"time"
)

/**
 *  To better handle the job exceptions ,it is really recommended to set panics
 *  inside your job function and resolve errors outside the cron function.
 */

/**
 * Refer to the java runable interface, the run method has no arguments
 */
func Basic(job func(), runtime time.Duration) {

}

/**
 * !warning:
 */
func WithOvertime(job func(), runtime time.Duration, overtime time.Duration) {
	t := time.NewTimer(runtime)
	defer t.Stop()
	end :=
		true
	start := false
	for range t.C {
		if !end {
			panic(errors.New("task timeout"))
		}
		/**
		 * The timer and task must be in different threads,otherwise the timer will be blocked
		 * when the task is blocked.
		 */
		go func() {
			if !start {
				t.Reset(overtime)
				start = true
				end = false
				job()
				end = true
				start = false
				t.Reset(runtime)
			}
		}()
	}
}
