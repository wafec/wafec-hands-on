package designpatterns.observer;

import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.never;
import static org.mockito.Mockito.verify;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

public class PublisherImplTest {
	private Publisher testObj;

	private static final String TEST_DATA = "test_data";

	@BeforeEach
	public void setUp() {
		testObj = new PublisherImpl();
	}

	@Test
	void publishing() {
		final Subscriber subscriber = mock( Subscriber.class );
		final Event inputEvent = Event.builder().data( TEST_DATA ).build();
		testObj.subscribe( subscriber );

		testObj.notifySubscribers( inputEvent );

		verify( subscriber ).handleEvent( inputEvent );
	}

	@Test
	void unsubscribing() {
		final Subscriber subscribed = mock( Subscriber.class );
		final Subscriber unsubscribed = mock( Subscriber.class );
		final Event inputEvent = Event.builder().data( TEST_DATA ).build();
		testObj.subscribe( subscribed );
		testObj.subscribe( unsubscribed );

		testObj.unsubscribe( unsubscribed );
		testObj.notifySubscribers( inputEvent );

		verify( subscribed ).handleEvent( inputEvent );
		verify( unsubscribed, never() ).handleEvent( inputEvent );
	}
}
