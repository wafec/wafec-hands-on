package designpatterns.behavioral.observer;

import java.util.ArrayList;
import java.util.List;

public class PublisherImpl implements Publisher {

	private final List<Subscriber> subscribers;

	public PublisherImpl() {
		subscribers = new ArrayList<>();
	}

	@Override
	public void subscribe( final Subscriber subscriber ) {
		if ( !subscribers.contains( subscriber ) ) {
			subscribers.add( subscriber );
		}
	}

	@Override
	public void unsubscribe( final Subscriber subscriber ) {
		subscribers.remove( subscriber );
	}

	@Override
	public void notifySubscribers( final Event event ) {
		for( final Subscriber subscriber : subscribers ) {
			subscriber.handleEvent( event );
		}
	}
}
