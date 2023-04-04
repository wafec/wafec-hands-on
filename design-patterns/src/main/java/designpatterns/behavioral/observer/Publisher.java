package designpatterns.behavioral.observer;

/*
	Main idea: subscription mechanism to notify interested parties about changes that may happen to
	           the observed object.
	Pros:
		1. Open/Close Principle: the presence of a new Subscriber does not change the behavior of
						         the Publisher.
	Cons:
		1. Subscribers are notified in random order.
 */
public interface Publisher {
	void subscribe( Subscriber subscriber );
	void unsubscribe( Subscriber subscriber );
	void notifySubscribers( Event event );
}
