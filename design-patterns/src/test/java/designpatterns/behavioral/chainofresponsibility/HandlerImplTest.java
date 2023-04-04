package designpatterns.behavioral.chainofresponsibility;

import static org.mockito.Mockito.never;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

@ExtendWith( MockitoExtension.class )
public class HandlerImplTest {
	private HandlerImpl testObj;

	@Mock
	private Handler nextHandler;
	@Mock
	private RequestProcessor requestProcessor;

	private static final String TEST_DATA = "test_data";

	@BeforeEach
	public void setUp() {
		testObj = new HandlerImpl( nextHandler, requestProcessor );
	}

	@Test
	void chaining() {
		final Request request = Request.builder().data( TEST_DATA ).build();
		final ProcessorResult processorResult = ProcessorResult.builder().shouldCallNextHandler( true ).build();
		when( requestProcessor.processRequest( request ) ).thenReturn( processorResult );

		testObj.execute( request );

		verify( nextHandler ).execute( request );
	}

	@Test
	void notCallingNextHandler() {
		final Request request = Request.builder().data( TEST_DATA ).build();
		final ProcessorResult processorResult = ProcessorResult.builder().shouldCallNextHandler( false ).build();
		when( requestProcessor.processRequest( request ) ).thenReturn( processorResult );

		testObj.execute( request );

		verify( nextHandler, never() ).execute( request );
	}
}
