package HelloWorld

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

//This is added
var log = logger.GetLogger("activity-helloworld")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// get the data from context

	name := context.GetInput("name").(string)

	salutation := context.GetInput("salutation").(string)

	log.Debugf("The Flogo engine says [%s] to [%s]", salutation, name)

	context.SetOutput("result", "The Flogo engine says "+salutation+" to "+name)

	return true, nil
}
