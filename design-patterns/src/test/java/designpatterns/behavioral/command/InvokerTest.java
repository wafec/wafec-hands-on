package designpatterns.behavioral.command;

import static org.mockito.Mockito.verify;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

@ExtendWith( MockitoExtension.class )
public class InvokerTest {
	private Invoker testObj;

	@Mock
	private HelloNotifier helloNotifier;

	@BeforeEach
	public void setUp() {

		testObj = new Invoker( new HelloCommand( helloNotifier ) );
	}

	@Test
	void invokingCommand() {
		final String expectedHelloMessage = "Hello!";

		testObj.executeCommand();

		verify( helloNotifier ).sendHello( expectedHelloMessage );
	}
}
