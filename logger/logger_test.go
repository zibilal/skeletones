package logger

import (
	"testing"
	"os"
)

func TestLogger(t *testing.T) {
	t.Log("Testing logger IS_DEBUG is not set")
	{
		Info("info1", "info2", "info3")
		Debug("debug1", "debug2", "debug3")
		Warn("warning1", "warning2", "warning3")
		Error("error1", "error2", "error3")
	}

	t.Log("Testing logger IS_DEBUG = 1")
	{
		os.Setenv("IS_DEBUG", "1")
		Info("info1", "info2", "info3")
		Debug("debug1", "debug2", "debug3")
		Warn("warning1", "warning2", "warning3")
		Error("error1", "error2", "error3")
	}

	t.Log("Testing logger IS_DEBUG = 0")
	{
		os.Setenv("IS_DEBUG", "0")
		Info("info1", "info2", "info3")
		Debug("debug1", "debug2", "debug3")
		Warn("warning1", "warning2", "warning3")
		Error("error1", "error2", "error3")

		os.Setenv("IS_DEBUG", "1")
	}
}
