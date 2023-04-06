package designpatterns.behavioral.state;

/*
	Main idea: lets an object changing its state when reacting to changes.
	Pros:
		1. Single Responsibility: each state does a particular task
		2. Open/Closed Principle: a new state implementation won't require any change
		3. Conditionals are eliminated when using state classes
	Cons:
		1. Overkill for a few states which rarely change
 */
public interface State {
	void operationA();

	void operationB();

	default String getStateName() {
		return getClass().getSimpleName();
	}
}
