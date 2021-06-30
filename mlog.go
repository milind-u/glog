// Google-Style Logger, wrapper around github.com/golang/glog with some additional functions, but not all functions
// in github.com/golang/glog.
package mlog

import (
  "flag"
  "fmt"

  "github.com/golang/glog"
)

// Level is a logging level, and is different from the glog.Level levels for glog.V logs
type Level int8

const (
  // LInfo is the Level for Info logs
  LInfo Level = iota
  // LWarning is the  Level for Warning logs
  LWarning
  // LError is the Level for Error logs
  LError
  // LFatal is the Level for Fatal logs
  LFatal
)

var (

  // Info is equivalent to glog.Info
  Info = glog.Info

  // Infof is equivalent to glog.Infof
  Infof = glog.Infof

  // Infoln is equivalent to glog.Infoln
  Infoln = glog.Infoln

  // Warning is equivalent to glog.Warning
  Warning = glog.Warning

  // Warningf is equivalent to glog.Warningf
  Warningf = glog.Warningf

  // Warningln is equivalent to glog.Warningln
  Warningln = glog.Warningln

  // Error is equivalent to glog.Error
  Error = glog.Error

  // Errorf is equivalent to glog.Errorf
  Errorf = glog.Errorf

  // Errorln is equivalent to glog.Errorln
  Errorln = glog.Errorln

  // Fatal is equivalent to glog.Fatal
  Fatal = glog.Fatal

  // Fatalf is equivalent to glog.Fatalf
  Fatalf = glog.Fatalf

  // Fatalln is equivalent to glog.Fatalln
  Fatalln = glog.Fatalln

  // Exit is equivalent to glog.Exit
  Exit = glog.Exit

  // Exitf is equivalent to glog.Exitf
  Exitf = glog.Exitf

  // Exitln is equivalent to glog.Exitln
  Exitln = glog.Exitln
)

// SetLevel sets the minimum logging level to be the given level
func SetLevel(level Level) {
  WarningIf(flag.Set("stderrthreshold", fmt.Sprint(level)))
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
  s := fmt.Sprintf("%v", v)
  s = s[1 : len(s)-1] // remove '[' and ']'
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
