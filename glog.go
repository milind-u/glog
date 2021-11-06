// Package glog is a Google-Style Logger, wrapper around github.com/golang/glog with some additional functions,
// but not all functions in github.com/golang/glog.
package glog

import (
  "flag"
  "fmt"
  "math"

  g "github.com/golang/glog"
)

// Severity is a logging severity
type Severity uint8

const (
  // InfoSeverity is the Severity for Info logs
  InfoSeverity Severity = iota
  // WarningSeverity is the  Severity for Warning logs
  WarningSeverity
  // ErrorSeverity is the Severity for Error logs
  ErrorSeverity
  // FatalSeverity is the Severity for Fatal logs
  FatalSeverity
)

var (
  Info   = g.Info
  Infof  = g.Infof
  Infoln = g.Infoln

  Warning   = g.Warning
  Warningf  = g.Warningf
  Warningln = g.Warningln

  Error   = g.Error
  Errorf  = g.Errorf
  Errorln = g.Errorln

  Fatal   = g.Fatal
  Fatalf  = g.Fatalf
  Fatalln = g.Fatalln

  Exit   = g.Exit
  Exitf  = g.Exitf
  Exitln = g.Exitln
)

// SetSeverity sets the minimum logging severity to be the given severity
func SetSeverity(severity Severity) {
  WarningIf(flag.Set("stderrthreshold", fmt.Sprint(severity)))
  flag.Parse()
}

func logIf(err error, logFunc func(v ...interface{})) {
  if err != nil {
    logFunc(err)
  }
}

// InfoIf calls Infoln with err if err != nil
func InfoIf(err error) {
  logIf(err, Infoln)
}

// WarningIf calls Warningln with err if err != nil
func WarningIf(err error) {
  logIf(err, Warningln)
}

// ErrorIf calls Errorln with err if err != nil
func ErrorIf(err error) {
  logIf(err, Errorln)
}

// FatalIf calls Fatalln with err if err != nil
func FatalIf(err error) {
  logIf(err, Fatalln)
}

// ExitIf calls Exitln with err if err != nil
func ExitIf(err error) {
  logIf(err, Exitln)
}

func varargsToStr(v ...interface{}) string {
  s := fmt.Sprint(v)
  s = s[2 : len(s)-2] // remove "[[" and "]]"
  return s
}

// Check calls Fatalln with v if !cond
func Check(cond bool, v ...interface{}) {
  if !cond {
    Fatalln("Check Failed: ", varargsToStr(v))
  }
}

func checkOperator(a, b interface{}, cond bool, operator string, v ...interface{}) {
  if !cond {
    Fatalf("Check Failed: expected \n%v %v %v\n%v\n", a, operator, b, varargsToStr(v))
  }
}

// CheckEq checks if a == b and calls Fatalf with a, b, and v if a != b
func CheckEq(a, b interface{}, v ...interface{}) {
  checkOperator(a, b, a == b, "==", v...)
}

// CheckNe checks if a != b and calls Fatalf with a, b, and v if a == b
func CheckNe(a, b interface{}, v ...interface{}) {
  checkOperator(a, b, a != b, "!=", v...)
}

// CheckLt checks if a < b and calls Fatalf with a, b, and v if a >= b
func CheckLt(a, b float64, v ...interface{}) {
  checkOperator(a, b, a < b, "<", v...)
}

// CheckGt checks if a > b and calls Fatalf with a, b, and v if a <= b
func CheckGt(a, b float64, v ...interface{}) {
  checkOperator(a, b, a > b, ">", v...)
}

// CheckLe checks if a <= b and calls Fatalf with a, b, and v if a > b
func CheckLe(a, b float64, v ...interface{}) {
  checkOperator(a, b, a <= b, "<=", v...)
}

// CheckGe checks if a >= b and calls Fatalf with a, b, and v if a < b
func CheckGe(a, b float64, v ...interface{}) {
  checkOperator(a, b, a >= b, ">=", v...)
}

// CheckNear checks if |a - b| <= threshold and calls Fatalf with a, b, and v if not
func CheckNear(a, b, threshold float64, v ...interface{}) {
  if diff := math.Abs(a - b); diff > threshold {
    Fatalf("Check Failed:\nthe difference between %v and %v, %v, exceeds the threshold %v\n%v\n",
      a, b, diff, threshold, varargsToStr(v))
  }
}
