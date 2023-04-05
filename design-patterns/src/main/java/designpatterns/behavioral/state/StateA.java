package designpatterns.behavioral.state;

public class StateA implements State {
	private final Context context;

	public StateA( final Context context ) {
		this.context = context;
	}

	@Override
	public void operationA() {
		context.changeState( new StateB( context ) );
	}

	@Override
	public void operationB() {

	}
}
