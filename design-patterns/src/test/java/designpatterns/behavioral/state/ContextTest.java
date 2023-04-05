package designpatterns.behavioral.state;

import org.assertj.core.api.Assertions;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

public class ContextTest {
	private Context testObj;

	@BeforeEach
	void setUp() {
		testObj = new Context();
	}

	@Test
	void stateChangesFromAToB() {
		testObj.operationA();

		Assertions.assertThat( testObj.getCurrentState().getStateName() ).isEqualTo( "StateB" );
	}

	@Test
	void stateChangesFromBToA() {
		testObj.operationA();
		testObj.operationB();

		Assertions.assertThat( testObj.getCurrentState().getStateName() ).isEqualTo( "StateA" );
	}
}
