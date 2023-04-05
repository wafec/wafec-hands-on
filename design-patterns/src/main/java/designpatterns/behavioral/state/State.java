package designpatterns.behavioral.state;

public interface State {
	void operationA();

	void operationB();

	default String getStateName() {
		return getClass().getSimpleName();
	}
}
