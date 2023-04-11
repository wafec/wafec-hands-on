package designpatterns.behavioral.state;

class Context {
	private State currentState;

	public Context() {
		currentState = new StateA( this );
	}

	public void changeState( final State state ) {
		currentState = state;
	}

	public void operationA() {
		currentState.operationA();
	}

	public void operationB() {
		currentState.operationB();
	}

	public State getCurrentState() {
		return currentState;
	}
}
