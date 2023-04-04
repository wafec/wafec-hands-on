package designpatterns.behavioral.mediator;

import java.util.ArrayList;
import java.util.List;

public class MediatorImpl implements Mediator {
	private final ComponentA componentA;
	private final ComponentB componentB;

	private final List<Integer> valuesList;

	public MediatorImpl( final ComponentA componentA, final ComponentB componentB ) {
		this.componentA = componentA;
		this.componentB = componentB;
		this.valuesList = new ArrayList<>();
	}

	@Override
	public void notifyMediator( final Component sender ) {
		if ( sender.equals( componentA )) {
			reactToComponentA();
		} else if ( sender.equals( componentB ) ) {
			reactToComponentB();
		}
	}

	private void reactToComponentA(){
		valuesList.add( componentA.getValue() );
	}

	private void reactToComponentB() {
		valuesList.add( componentB.getValue() );
	}

	public List<Integer> getValuesList() {
		return new ArrayList<>( valuesList );
	}
}
