package designpatterns.behavioral.state;

public class StateB implements State {
	private final Context context;

	public StateB( final Context context ) {
		this.context = context;
	}

	@Override
	public void operationA() {

	}

	@Override
	public void operationB() {
		context.changeState( new StateA( context ) );
	}
}
