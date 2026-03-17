package workdef

type Workflow struct {
	WorkflowId int
}

func Start() {
	println("Work Definition processing")

	/*
		- Accept workflow definitions from users
		- Represent workflows as **DAG**
		- Validate workflow structure
		   - Detect cycles in workflows
		- Store workflow metadata
		- Store task dependency relationships
	*/
}
