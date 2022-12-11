package templatemethod

func ExampleTemplate() {
	var t Templater
	t = NewConcreate1()
	t.Template()
	t = NewConcreate2()
	t.Template()
	// Output:
	// concreate1 step1
	// concreate1 step2
	// concreate1 step3
	// concreate2 step1
	// default step2
	// concreate2 step3
}
