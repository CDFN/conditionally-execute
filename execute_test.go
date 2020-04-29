package condexec

import "testing"

func TestConditionallyExecute(t *testing.T){
	executed := false
	ce := New(false).OnFalse(func(){
		executed = false
		t.Log("onFalse executed")
	}).OnTrue(func() {
		executed = true
		t.Log("onTrue executed")
	}).Execute()
	if executed != ce.condition{
		t.Errorf("Expected %v, got %v", ce.condition, executed)
	}
}